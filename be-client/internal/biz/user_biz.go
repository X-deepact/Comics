package biz

import (
	"be-client/internal/dto/request"
	"be-client/internal/dto/response"
	"be-client/internal/mapper"
	"be-client/internal/middleware"
	"be-client/internal/repository"
	"be-client/util"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"pkg-common/model"

	"gorm.io/gorm"
)

type UserBiz interface {
	UpdateUser(ctx context.Context, id int64, req *request.UserUpdateRequest) (response.UserResponse, error)
	ValidateUser(ctx context.Context, username string, hashPassword string) (*response.UserResponse, error)
	GetUserProfile(ctx context.Context, username string) (*response.ProfileResponse, error)
	RegisterUser(ctx context.Context, req *request.RegisterRequest) (response.ProfileResponse, error)
	ChangePassword(ctx context.Context, userId int64, input *request.ChangePasswordRequest) (response.UserResponse, error)
}

type userBiz struct {
	userMapper     mapper.UserMapper
	profileMapper  mapper.ProfileMapper
	userRepository repository.UserRepository
	profileRepo    repository.ProfileRepository
	tierrepo       repository.TierRepository
	auth           middleware.AuthenHandler
}

func NewUserBiz(repo repository.UserRepository,
	profileRepo repository.ProfileRepository,
	tierRepo repository.TierRepository,
	profileMapper mapper.ProfileMapper,
	userMapper mapper.UserMapper,
	auth middleware.AuthenHandler) UserBiz {
	return &userBiz{userRepository: repo, userMapper: userMapper,
		profileMapper: profileMapper,
		tierrepo:      tierRepo,
		auth:          auth,
		profileRepo:   profileRepo}
}

func (u *userBiz) UpdateUser(ctx context.Context, id int64, req *request.UserUpdateRequest) (response.UserResponse, error) {
	userModel, err := u.userRepository.GetUserById(ctx, id)
	if err != nil {
		return response.UserResponse{}, err
	}
	if req.FirstName != "" {
		userModel.FirstName = req.FirstName
	}
	if req.LastName != "" {
		userModel.LastName = req.LastName
	}
	if req.Phone != "" {
		userModel.Phone = &req.Phone
	}
	if req.Email != "" {
		userModel.Email = &req.Email
	}

	if req.FirstName != "" && req.LastName != "" {
		userModel.FullName = req.FirstName + " " + req.LastName
	}

	err = u.userRepository.UpdateUser(ctx, userModel)
	if err != nil {
		return response.UserResponse{}, err
	}
	return u.userMapper.ToResponse(userModel), nil
}

func (u *userBiz) ValidateUser(ctx context.Context, username string, hashPassword string) (*response.UserResponse, error) {
	userModel, err := u.userRepository.ValidateUserPasswordByUserName(ctx, username, hashPassword)
	if err != nil {
		return nil, err
	}
	if userModel.Id > 0 {
		resp := u.userMapper.ToResponse(userModel)
		return &resp, nil
	}
	return nil, nil
}

func (u *userBiz) GetUserProfile(ctx context.Context, username string) (*response.ProfileResponse, error) {
	profile, err := u.profileRepo.GetUserProfile(ctx, username)
	if err != nil {
		return nil, err
	}

	profileResponse := u.profileMapper.ToResponse(profile)
	tierModel, err := u.tierrepo.GetTierById(context.Background(), profile.TierId)
	if err == nil && tierModel != nil && tierModel.Id > 0 {
		profileResponse.Tier = u.profileMapper.ToPrfileTierResponse(tierModel)
	}

	return profileResponse, nil
}

func (u *userBiz) RegisterUser(ctx context.Context, req *request.RegisterRequest) (response.ProfileResponse, error) {
	userModel, err := u.userRepository.GetUserByUserName(ctx, req.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		slog.Error(fmt.Sprintf("get user by username error: %v", err))
		return response.ProfileResponse{}, err
	}

	if userModel != nil && userModel.Id > 0 {
		slog.Error("username already exists")
		return response.ProfileResponse{}, errors.New("username already exists")
	}

	userModel = &model.UserModel{
		Username:     req.Username,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		FullName:     req.FirstName + " " + req.LastName,
		Phone:        &req.Phone,
		Email:        &req.Email,
		HashPassword: req.HashedPassword,
	}
	err = u.userRepository.CreateUser(ctx, userModel)
	if err != nil {
		slog.Error(fmt.Sprintf("create user error: %v", err))
		return response.ProfileResponse{}, err
	}

	profileModel := &model.ProfileModel{
		Active:      true,
		UserId:      userModel.Id,
		DisplayName: userModel.FullName,
		Description: "User profile of " + userModel.FullName,
		TierId:      util.DefaultTierNewUser,
	}

	err = u.profileRepo.CreateProfile(ctx, profileModel)
	if err != nil {
		slog.Error(fmt.Sprintf("create profile error: %v", err))
		return response.ProfileResponse{}, err
	}
	profileResp := u.profileMapper.ToResponse(profileModel)
	profileResp.Username = userModel.Username
	tierModel, err := u.tierrepo.GetTierById(context.Background(), profileModel.TierId)
	if err == nil && tierModel != nil && tierModel.Id > 0 {
		slog.Error(fmt.Sprintf("get tier by id error: %v", err))
		profileResp.Tier = u.profileMapper.ToPrfileTierResponse(tierModel)
	}
	return *profileResp, nil
}

func (u *userBiz) ChangePassword(ctx context.Context, userId int64, input *request.ChangePasswordRequest) (response.UserResponse, error) {
	if userId <= 0 {
		return response.UserResponse{}, errors.New("user id is invalid")
	}
	user, err := u.userRepository.GetUserById(ctx, userId)
	if err != nil {
		return response.UserResponse{}, err
	}

	oldPasswordHash, err := u.auth.HashPassword(input.OldPassword)
	if err != nil {
		return response.UserResponse{}, err
	}
	if oldPasswordHash != user.HashPassword {
		return response.UserResponse{}, errors.New("old password is incorrect")
	}

	newPasswordHash, err := u.auth.HashPassword(input.NewPassword)
	if err != nil {
		return response.UserResponse{}, err
	}
	user.HashPassword = newPasswordHash
	err = u.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return response.UserResponse{}, err
	}
	return u.userMapper.ToResponse(user), nil
}

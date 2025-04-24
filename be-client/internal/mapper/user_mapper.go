package mapper

import (
	"be-client/internal/dto/request"
	"be-client/internal/dto/response"
	"pkg-common/model"
	"time"
)

type UserMapper interface {
	ToModel(req *request.UserUpdateRequest) *model.UserModel
	ToResponse(user *model.UserModel) response.UserResponse
}

type userMapper struct {
}

func NewUserMapper() UserMapper {
	return &userMapper{}
}

func (m *userMapper) ToModel(req *request.UserUpdateRequest) *model.UserModel {
	now := time.Now()
	return &model.UserModel{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		FullName:  req.FirstName + " " + req.LastName,
		Email:     &req.Email,
		Phone:     &req.Phone,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}

func (m *userMapper) ToResponse(user *model.UserModel) response.UserResponse {
	if user == nil {
		return response.UserResponse{}
	}
	resp := response.UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		FullName:  user.FullName,
		Email:     *user.Email,
		Phone:     *user.Phone,
	}
	if user.CreatedAt != nil {
		resp.CreatedAt = user.CreatedAt.UnixMilli()
	}
	if user.UpdatedAt != nil {
		resp.UpdatedAt = user.UpdatedAt.UnixMilli()
	}
	return resp
}

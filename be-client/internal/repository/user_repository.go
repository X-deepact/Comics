package repository

import (
	"be-client/internal/infra/database"
	"context"
	"errors"
	"pkg-common/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.UserModel) error
	UpdateUser(ctx context.Context, user *model.UserModel) error
	GetUserByUserName(ctx context.Context, username string) (*model.UserModel, error)
	GetUserById(context.Context, int64) (*model.UserModel, error)
	ValidateUser(ctx context.Context, userName, email, phone string) error
	ValidateUserPasswordByUserName(context.Context, string, string) (*model.UserModel, error)
}

type userRepository struct {
	db database.DBInterface
}

func NewUserRepository(db database.DBInterface) UserRepository {
	return &userRepository{db: db}
}
func (r *userRepository) CreateUser(ctx context.Context, user *model.UserModel) error {
	return r.db.Ctx(ctx).Create(user)
}

func (r *userRepository) UpdateUser(ctx context.Context, user *model.UserModel) error {
	return r.db.Ctx(ctx).Save(user)
}

func (r *userRepository) GetUserByUserName(ctx context.Context, username string) (*model.UserModel, error) {
	req := model.UserModel{}
	err := r.db.Ctx(ctx).Where("username = ?", username).First(&req)
	return &req, err
}

func (r *userRepository) ValidateUser(ctx context.Context, userName, email, phone string) error {
	req := model.UserModel{}
	err := r.db.Ctx(ctx).Where("username = ?", userName).Or("email = ?", email).Or("phone = ?", phone).First(&req)
	if err != nil {
		return err
	}

	if req.Id > 0 {
		return errors.New("username, email or phone already exists")
	}
	return nil
}

func (r *userRepository) ValidateUserPasswordByUserName(ctx context.Context, username, passwordHash string) (*model.UserModel, error) {
	req := model.UserModel{}
	err := r.db.Ctx(ctx).Where("username = ?", username).Where("hash_password = ?", passwordHash).First(&req)
	if err != nil {
		return nil, err
	}
	if req.Id == 0 {
		return nil, errors.New("username or password is incorrect")
	}
	return &req, nil
}

func (r *userRepository) GetUserById(ctx context.Context, id int64) (*model.UserModel, error) {
	var user model.UserModel
	if id <= 0 {
		return &user, errors.New("input invalid")
	}
	err := r.db.Where("id = ?", id).First(&user)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

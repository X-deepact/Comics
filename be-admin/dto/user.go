package dto

import (
	"comics-admin/val"
	"fmt"
	"pkg-common/model"
	"time"
)

type UserRequest struct {
	Username    string `json:"username" binding:"required,alphanum"`
	Password    string `json:"password" binding:"required,min=7"`
	Phone       string `json:"phone" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DisplayName string `json:"display_name" binding:"required"`
	Description string `json:"description"`
}

func (req *UserRequest) Validate() error {
	validators := []struct {
		name string
		fn   func() error
	}{
		{"username", func() error { return val.ValidateUsername(req.Username) }},
		//{"full name", func() error { return val.ValidateFullName(req.FullName) }},
		{"password", func() error { return val.ValidatePassword(req.Password) }},
		{"email", func() error { return val.ValidateEmail(req.Email) }},
	}

	for _, validator := range validators {
		if err := validator.fn(); err != nil {
			return fmt.Errorf("%s validation failed: %w", validator.name, err)
		}
	}

	return nil
}

type UserResponse struct {
	Id            int64      `json:"id"`
	Username      string     `json:"username"`
	Phone         string     `json:"phone"`
	Email         string     `json:"email"`
	Birthday      *time.Time `json:"birthday"`
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	FullName      string     `json:"full_name"`
	Active        bool       `json:"active"`
	RoleName      string     `json:"role_name"`
	DisplayName   string     `json:"display_name"`
	Description   string     `json:"description"`
	Avatar        string     `json:"avatar"`
	TierID        int64      `json:"tier_id"`
	TierCode      string     `json:"tier_code"`
	CreatedAt     string     `json:"created_at"`
	CreatedByName string     `json:"created_by_name"`
	UpdatedByName string     `json:"updated_by_name"`
	UpdatedAt     string     `json:"updated_at"`
}

type UserLoginResponse struct {
	UserResponse
	HashPassword string `json:"hash_password"`
}

type UserCreateRequest struct {
	Username    string `form:"username" binding:"required,alphanum"`
	Phone       string `form:"phone"`
	Email       string `form:"email" binding:"email"`
	Birthday    string `form:"birthday"`
	Password    string `form:"password" binding:"required,min=7"`
	FirstName   string `form:"first_name"`
	LastName    string `form:"last_name"`
	DisplayName string `form:"display_name" binding:"required"`
	Description string `form:"description"`
	TierId      int64  `form:"tier_id" binding:"required"`
}

type UserListRequest struct {
	Page        int    `form:"page" json:"page" binding:"required,min=1"`
	PageSize    int    `form:"page_size" json:"page_size" binding:"required,min=1,max=100"`
	Username    string `form:"username" json:"username"`
	Phone       string `form:"phone" json:"phone"`
	Email       string `form:"email" json:"email"`
	Name        string `form:"name" json:"name"`
	DisplayName string `form:"display_name" json:"display_name"`
	TierId      int64  `form:"tier_id" json:"tier_id"`
}

type UserUpdateRequest struct {
	ID          int64  `form:"id" binding:"required"`
	Username    string `form:"username" binding:"required,alphanum"`
	Phone       string `form:"phone"`
	Email       string `form:"email" binding:"email"`
	Birthday    string `form:"birthday"`
	FirstName   string `form:"first_name"`
	LastName    string `form:"last_name"`
	DisplayName string `form:"display_name" binding:"required"`
	Description string `form:"description"`
	TierId      int64  `form:"tier_id" binding:"required"`
}

type UserModel struct {
	model.UserModel
	Profile   model.ProfileModel    `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	UserRoles []model.UserRoleModel `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
}

type UserExistDto struct {
	IsUsername bool `json:"is_username"`
	IsPhone    bool `json:"is_phone"`
	IsEmail    bool `json:"is_email"`
}

type UserProfileUpdateRequest struct {
	Username    string `form:"username" binding:"required,alphanum"`
	Phone       string `form:"phone"`
	Email       string `form:"email" binding:"email"`
	Birthday    string `form:"birthday"`
	FirstName   string `form:"first_name"`
	LastName    string `form:"last_name"`
	DisplayName string `form:"display_name" binding:"required"`
	Description string `form:"description"`
}

type UserDetailDto struct {
	Id          int64      `json:"id"`
	Username    string     `json:"username"`
	Phone       string     `json:"phone"`
	Email       string     `json:"email"`
	Birthday    *time.Time `json:"birthday"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	FullName    string     `json:"full_name"`
	Active      bool       `json:"active"`
	RoleName    string     `json:"role_name"`
	DisplayName string     `json:"display_name"`
	Description string     `json:"description"`
	Avatar      string     `json:"avatar"`
	TierID      int64      `json:"tier_id"`
	TierCode    string     `json:"tier_code"`
}

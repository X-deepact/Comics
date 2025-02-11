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

type UserResponse struct {
	Username          string     `json:"username"`
	Email             string     `json:"email"`
	FullName          string     `json:"full_name"`
	Role              string     `json:"role"`
	Active            bool       `json:"active"`
	PasswordChangedAt time.Time  `json:"password_changed_at"`
	CreatedAt         *time.Time `json:"created_at"`
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

type UserModel struct {
	model.UserModel
	RoleName string `json:"role_name"`
}

package dto

import (
	"comics-admin/val"
	"fmt"
	"time"
)

type UserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=7"`
	Email    string `json:"email" binding:"required,email"`
	FullName string `json:"fullName" binding:"required"`
	Role     string `json:"role" `
	IsActive bool   `json:"isActive"`
}

type UserResponse struct {
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	FullName          string    `json:"full_name"`
	Role              string    `json:"role"`
	IsActive          bool      `json:"is_active"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func (req *UserRequest) Validate() error {
	validators := []struct {
		name string
		fn   func() error
	}{
		{"username", func() error { return val.ValidateUsername(req.Username) }},
		{"full name", func() error { return val.ValidateFullName(req.FullName) }},
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

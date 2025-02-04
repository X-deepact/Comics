package dto

import "time"

type LoginRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=7"`
}

type LoginResponse struct {
	AccessToken          string       `json:"access_token"`
	AccessTokenExpiresAt time.Time    `json:"access_token_expires_at"`
	User                 UserResponse `json:"user"`
}

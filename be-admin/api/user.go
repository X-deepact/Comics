package api

import (
	"comics-admin/db"
	"comics-admin/dto"
	config "comics-admin/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func newUserResponse(user *db.User) dto.UserResponse {
	return dto.UserResponse{
		Username:  user.Username,
		Email:     user.Email,
		FullName:  user.FullName,
		Role:      string(user.Role),
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
	}
}

func (s *Server) register(ctx *gin.Context) {
	var req dto.UserRequest

	req.IsActive = true

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashPassword, err := config.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := &db.User{
		Username: req.Username,
		Password: hashPassword,
		Email:    req.Email,
		FullName: req.FullName,
		Role:     db.UsersRole(req.Role),
		IsActive: req.IsActive,
	}

	err = s.store.CreateUser(arg)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := s.store.GetUser(arg.Username)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	res := newUserResponse(user)

	ctx.JSON(http.StatusOK, res)
}

func (s *Server) login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := s.store.GetUser(req.Username)

	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	err = config.CheckPassword(req.Password, user.Password)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, accessPayload, err := s.tokenMaker.CreateToken(req.Username, string(user.Role), s.config.Source.AccessTokenDuration)

	fmt.Printf("%s", accessToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := dto.LoginResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
		User:                 newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, res)
}

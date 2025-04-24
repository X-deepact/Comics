package handler

import (
	"be-client/internal/biz"
	"be-client/internal/dto/request"
	"be-client/internal/dto/response"
	"be-client/internal/middleware"
	"be-client/util"
	"context"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	RegisterUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
}

type authHandler struct {
	uc   biz.UserBiz
	auth middleware.AuthenHandler
}

func NewAuthHandler(uc biz.UserBiz, auth middleware.AuthenHandler) AuthHandler {
	return &authHandler{uc: uc, auth: auth}
}

// @Summary     Register new user
// @Description Create a new user account and return JWT tokens
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       request          body      request.RegisterRequest  true  "Registration credentials"
// @Success 200 {object} common.Response{data=response.RegisterResponse} "Successful response or Fail response. Code can be 'SUCCESS' or 'ERROR'."
// @Router      /v1/auth/register [post]
func (h *authHandler) RegisterUser(c *fiber.Ctx) error {
	req := &request.RegisterRequest{}
	if err := req.Bind(c); err != nil {
		return util.ResponseApi(c, nil, err)
	}

	hashPassword, _ := h.auth.HashPassword(req.Password)
	req.HashedPassword = hashPassword

	profile, err := h.uc.RegisterUser(context.Background(), req)
	if err != nil {
		return util.ResponseApi(c, nil, err)
	}

	token, err := h.auth.GenerateTokenPair(profile.Id, profile.Username)
	if err != nil {
		return util.ResponseApi(c, nil, err)
	}
	resp := response.RegisterResponse{
		Profile:      profile,
		Token:        token.AccessToken,
		RefreshToken: token.RefreshToken,
	}
	return util.ResponseApi(c, resp, nil)
}

// @Summary Login user with username and password
// @Description Login user with username and password in the system
//
//	@Tags Auth
//	@Accept json
//	@Produce json
//
// @Param request body request.LoginRequest true "Login request"
//
// @Success 200 {object} common.Response{data=response.LoginResponse} "Successful response or Fail response. Code can be 'SUCCESS' or 'ERROR'."
//
//	@Router /v1/auth/login [post]
func (h *authHandler) LoginUser(c *fiber.Ctx) error {
	req := &request.LoginRequest{}
	if err := req.Bind(c); err != nil {
		return util.ResponseApi(c, nil, err)
	}

	hashPassword, _ := h.auth.HashPassword(req.Password)
	user, err := h.uc.ValidateUser(context.Background(), req.Username, hashPassword)
	if err != nil {
		return util.ResponseApi(c, nil, err)
	}

	if user == nil {
		return util.ResponseApi(c, nil, errors.New(response.ErrNotFound.Error))
	}

	token, err := h.auth.GenerateTokenPair(user.Id, user.Username)
	if err != nil {
		return util.ResponseApi(c, nil, err)
	}

	resp := response.LoginResponse{
		User:         *user,
		Token:        token.AccessToken,
		RefreshToken: token.RefreshToken,
	}
	return util.ResponseApi(c, resp, nil)
}

// @Summary Refresh user token
// @Description Refresh user token
// @Tags Auth
// @Accept json
// @Produce json
// @Param     RefreshToken   header string true "Refresh token"
// @Success 200 {object} common.Response{data=middleware.TokenPair} "Successful response or Fail response. Code can be 'SUCCESS' or 'ERROR'."
// @Router /v1/auth/refresh [post]
func (h *authHandler) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Get(middleware.RefreshTokenHeader)
	if refreshToken == "" {
		return util.ResponseApi(c, nil, middleware.ErrMissingToken)
	}

	refreshToken = strings.TrimPrefix(refreshToken, middleware.Prefix+" ")
	claims, err := h.auth.ValidateRefreshToken(refreshToken)
	if err != nil {
		return util.ResponseApi(c, nil, err)
	}

	tokenPair, err := h.auth.GenerateAcessToken(claims.UserId, claims.UserName)
	if err != nil {
		return util.ResponseApi(c, nil, err)
	}
	tokenPair.RefreshToken = refreshToken
	return util.ResponseApi(c, tokenPair, nil)
}

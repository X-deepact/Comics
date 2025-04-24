package handler

import (
	"be-client/internal/biz"
	"be-client/internal/dto/request"
	"be-client/internal/middleware"
	"be-client/util"
	"context"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	UpdateUser(c *fiber.Ctx) error
	Profile(c *fiber.Ctx) error
	ChangePassword(c *fiber.Ctx) error
}

type userHandler struct {
	userUc biz.UserBiz
	auth   middleware.AuthenHandler
}

func NewUserHandler(uc biz.UserBiz, auth middleware.AuthenHandler) UserHandler {
	return &userHandler{userUc: uc, auth: auth}
}

// @Summary Create new user
// @Description Create a new user in the system
// @Tags Users
// @Accept json
// @Produce json
// @Param     Authorization header string true "Bearer authorization token"
// @Param request body request.UserUpdateRequest true "User creation request"
// @Success 200 {object} common.Response{data=response.UserResponse} "Successful response or Fail response. Code can be 'SUCCESS' or 'ERROR'."
// @Router /v1/user/update [put]
func (h *userHandler) UpdateUser(c *fiber.Ctx) error {
	req := &request.UserUpdateRequest{}
	if err := req.Bind(c); err != nil {
		return util.ResponseApi(c, nil, err)
	}

	id, err := h.auth.GetUserIdFromContext(c)
	if err != nil {
		return util.ResponseApi(c, nil, err)
	}

	resp, err := h.userUc.UpdateUser(context.Background(), id, req)
	if err != nil {
		return util.ResponseApi(c, nil, err)
	}
	return util.ResponseApi(c, resp, nil)
}

// @Summary Get user profile
// @Description Get user profile
// @Tags Users
// @Accept json
// @Param     Authorization  header string true "Bearer authorization token"
// @Produce json
// @Success 200 {object} common.Response{data=response.ProfileResponse} "Successful response or Fail response. Code can be 'SUCCESS' or 'ERROR'."
// @Router /v1/user/profile [get]
func (h *userHandler) Profile(c *fiber.Ctx) error {
	username, err := h.auth.GetUserNameFromContext(c)
	if err != nil {
		return util.ResponseApi(c, nil, err)
	}
	resp, err := h.userUc.GetUserProfile(context.Background(), username)
	if err != nil {
		return util.ResponseApi(c, nil, err)
	}
	return util.ResponseApi(c, resp, nil)
}

// @Summary Change user password
// @Description Change user password
// @Tags Users
// @Accept json
// @Produce json
// @Param     Authorization  header string true "Bearer authorization token"
// @Param request body request.ChangePasswordRequest true "Change password request"
// @Success 200 {object} common.Response{data=response.UserResponse} "Successful response or Fail response. Code can be 'SUCCESS' or 'ERROR'."
// @Router /v1/user/change-password [put]
func (h *userHandler) ChangePassword(c *fiber.Ctx) error {
	req := &request.ChangePasswordRequest{}
	if err := req.Bind(c); err != nil {
		return util.ResponseApi(c, nil, err)
	}

	userId, err := h.auth.GetUserIdFromContext(c)
	if err != nil {
		return util.ResponseApi(c, nil, err)
	}
	resp, err := h.userUc.ChangePassword(context.Background(), int64(userId), req)
	if err != nil {
		return util.ResponseApi(c, nil, err)
	}
	return util.ResponseApi(c, resp, nil)
}

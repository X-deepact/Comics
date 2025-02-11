package api

import (
	"comics-admin/dto"
	config "comics-admin/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pkg-common/model"
	"strings"
)

func newUserResponse(user *dto.UserModel) dto.UserResponse {
	return dto.UserResponse{
		Username:  user.Username,
		Email:     user.Email,
		FullName:  user.FullName,
		Active:    user.Active,
		CreatedAt: user.CreatedAt,
		Role:      user.RoleName,
	}
}

func (s *Server) userRouter() {
	group := s.router.Group("api/user")

	group.POST("/register", s.register)
	group.POST("/login", s.login)

	group.POST("/upload-avatar", s.saveAvatar)
}

func (s *Server) register(ctx *gin.Context) {
	var req dto.UserRequest
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

	arg := &model.UserModel{
		Username:     req.Username,
		HashPassword: hashPassword,
		Email:        req.Email,
		Phone:        req.Phone,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		FullName:     strings.Join(strings.Fields(req.FirstName+" "+req.LastName), " "),
		Active:       true,
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

	roleAdmin, err := s.store.GetRole(config.ADMIN)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userRole := &model.UserRoleModel{
		UserId: user.Id,
		RoleId: roleAdmin.Id,
	}

	err = s.store.CreateUserRole(userRole)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res := newUserResponse(user)

	ctx.JSON(http.StatusOK, res)
}

// @Summary Login
// @Description Authenticates the user and returns a JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param login body dto.LoginRequest true "User Login Data"
// @Success 200 {object} dto.LoginResponse "Login successful"
// @Failure 400  "Invalid request"
// @Router /api/user/login [post]
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
	err = config.CheckPassword(req.Password, user.HashPassword)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, accessPayload, err := s.tokenMaker.CreateToken(user.Id, req.Username, user.RoleName, s.config.Source.AccessTokenDuration)
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

func (s *Server) saveAvatar(ctx *gin.Context) {
	file, _ := ctx.FormFile("avatar")
	fileLink := ""

	if file != nil {
		fileName, err := config.SaveImage(file, s.config.FileStorage.AvatarFolder)
		if err != nil {
			config.BuildErrorResponse(ctx, http.StatusBadRequest, err, nil)
			return
		}

		fileLink = config.GetFileUrl(s.config.ApiFile.Url, s.config.FileStorage.RootFolder, s.config.FileStorage.AvatarFolder, fileName)
	}

	ctx.JSON(http.StatusOK, fileLink)
}

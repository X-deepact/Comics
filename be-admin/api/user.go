package api

import (
	"comics-admin/dto"
	config "comics-admin/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"pkg-common/model"
	"strconv"
	"strings"
	"time"
)

func (s *Server) userRouter() {
	group := s.router.Group("api/user")

	group.POST("/register", s.register)
	group.POST("/login", s.login)

	group.Use(s.authMiddleware(s.tokenMaker))
	group.POST("", s.createUser)
	group.GET("/:id", s.getUser)
	group.GET("", s.getUsers)
	group.PUT("", s.updateUser)
	group.DELETE("/:id", s.deleteUser)
	group.PUT("/:id/active", s.activeUser)
	group.GET("/profile", s.getProfile)
	group.PUT("/profile", s.updateProfile)
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

	user := model.UserModel{
		Username:     req.Username,
		HashPassword: hashPassword,
		Email:        req.Email,
		Phone:        req.Phone,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		FullName:     strings.Join(strings.Fields(req.FirstName+" "+req.LastName), " "),
		Active:       true,
	}

	profile := model.ProfileModel{
		Active:      true,
		DisplayName: req.DisplayName,
		Description: req.Description,
		UserId:      user.Id,
		TierId:      1,
	}

	//role
	roleAdmin, err := s.store.GetRole(config.USER)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userRole := model.UserRoleModel{
		UserId: user.Id,
		RoleId: roleAdmin.Id,
	}

	userSave := &dto.UserModel{
		UserModel: user,
		Profile:   profile,
		UserRoles: []model.UserRoleModel{userRole},
	}

	err = s.store.CreateUser(userSave)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := s.store.GetUser(userSave.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

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

	user, err := s.store.GetUserLogin(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "username or password is incorrect"})
			return
		}

		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	err = config.CheckPassword(req.Password, user.HashPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "username or password is incorrect"})
		return
	}

	accessToken, accessPayload, err := s.tokenMaker.CreateToken(user.Id, req.Username, user.RoleName, s.config.Source.AccessTokenDuration)
	fmt.Printf("%s", accessToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if user.Avatar != "" {
		user.Avatar = config.GetFileUrl(s.config.ApiFile.Url, s.config.FileStorage.RootFolder, s.config.FileStorage.AvatarFolder, user.Avatar)
	}

	res := dto.LoginResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
		User: dto.UserDetailDto{
			Id:          user.Id,
			Username:    user.Username,
			Phone:       user.Phone,
			Email:       user.Email,
			Birthday:    user.Birthday,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			FullName:    user.FullName,
			Active:      user.Active,
			RoleName:    user.RoleName,
			DisplayName: user.DisplayName,
			Description: user.Description,
			Avatar:      user.Avatar,
			TierID:      user.TierID,
			TierCode:    user.TierCode,
		},
	}
	ctx.JSON(http.StatusOK, res)
}

// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param username formData string true "Username (Alphanumeric)"
// @Param phone formData string false "Phone Number"
// @Param email formData string false "Email Address"
// @Param birthday formData string false "Birthday (YYYY-MM-DD)"
// @Param password formData string true "Password (Min: 7 chars)"
// @Param first_name formData string false "First Name"
// @Param last_name formData string false "Last Name"
// @Param display_name formData string true "Display Name"
// @Param description formData string false "Description"
// @Param avatar formData file false "Avatar File"
// @Param tier_id formData int64 true "Tier ID"
// @Security     BearerAuth
// @Success 200 {object} dto.UserDetailDto
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Failure 401 {object} dto.ErrorResponse "User not authenticated"
// @Router /api/user [post]
func (s *Server) createUser(ctx *gin.Context) {
	var req dto.UserCreateRequest

	// Bind form-data request
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userExist, err := s.store.CheckUserExist(req.Username, req.Phone, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if userExist.IsUsername || userExist.IsPhone || userExist.IsEmail {
		var list []string

		if userExist.IsUsername {
			list = append(list, "username")
		}
		if userExist.IsPhone {
			list = append(list, "phone")
		}
		if userExist.IsEmail {
			list = append(list, "email")
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"error": strings.Join(list, ", ") + " already exist"})
		return
	}

	// Retrieve avatar file
	file, _ := ctx.FormFile("avatar")

	// Extract user ID from context
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("user not authenticated")))
		return
	}

	userIDInt64, ok := userID.(int64)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("invalid user ID type")))
		return
	}

	//user
	var birthday *time.Time
	if req.Birthday != "" {
		bConvert, err := config.ConvertStringToDate(req.Birthday)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		birthday = &bConvert
	}

	hashPassword, err := config.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user := model.UserModel{
		Username:     req.Username,
		Phone:        req.Phone,
		Email:        req.Email,
		BirthDay:     birthday,
		HashPassword: hashPassword,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		FullName:     strings.Join(strings.Fields(req.FirstName+" "+req.LastName), " "),
		CreatedBy:    userIDInt64,
		UpdatedBy:    userIDInt64,
	}

	//profile
	fileName := ""
	if file != nil {
		fileName, err = config.SaveImage(file, s.config.FileStorage.AvatarFolder)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
	}

	profile := model.ProfileModel{
		Active:      true,
		DisplayName: req.DisplayName,
		Description: req.Description,
		Avatar:      fileName,
		UserId:      user.Id,
		TierId:      req.TierId,
		CreatedBy:   userIDInt64,
		UpdatedBy:   userIDInt64,
	}

	//role
	roleAdmin, err := s.store.GetRole(config.USER)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userRole := model.UserRoleModel{
		UserId: user.Id,
		RoleId: roleAdmin.Id,
	}

	userSave := &dto.UserModel{
		UserModel: user,
		Profile:   profile,
		UserRoles: []model.UserRoleModel{userRole},
	}

	err = s.store.CreateUser(userSave)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userRes, err := s.store.GetUser(userSave.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	if userRes.Avatar != "" {
		userRes.Avatar = config.GetFileUrl(s.config.ApiFile.Url, s.config.FileStorage.RootFolder, s.config.FileStorage.AvatarFolder, userRes.Avatar)
	}

	ctx.JSON(http.StatusOK, userRes)
}

// @Summary Get a user
// @Description Get a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Security     BearerAuth
// @Success 200 {object} dto.UserDetailDto
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/user/{id} [get]
func (s *Server) getUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := s.store.GetUser(int64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if user.Avatar != "" {
		user.Avatar = config.GetFileUrl(s.config.ApiFile.Url, s.config.FileStorage.RootFolder, s.config.FileStorage.AvatarFolder, user.Avatar)
	}

	ctx.JSON(http.StatusOK, user)
}

// @Summary List users
// @Description List all user
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param username query string false "Username"
// @Param phone query string false "Phone"
// @Param email query string false "Email"
// @Param name query string false "Name"
// @Param display_name query string false "Display name"
// @Param tier_id query integer false "Tier ID"
// @Security     BearerAuth
// @Success 200 {object} []dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/user [get]
func (s *Server) getUsers(ctx *gin.Context) {
	var req dto.UserListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	users, total, err := s.store.GetUsers(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	for _, user := range users {
		if user.Avatar != "" {
			user.Avatar = config.GetFileUrl(s.config.ApiFile.Url, s.config.FileStorage.RootFolder, s.config.FileStorage.AvatarFolder, user.Avatar)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"pagination": gin.H{
			"page":      req.Page,
			"page_size": req.PageSize,
			"total":     total,
		},
		"data": users,
	})
}

// @Summary Update a new user
// @Description Update a new user
// @Tags users
// @Accept json
// @Produce json
// @Param id formData integer true "ID"
// @Param username formData string true "Username (Alphanumeric)"
// @Param phone formData string false "Phone Number"
// @Param email formData string false "Email Address"
// @Param birthday formData string false "Birthday (YYYY-MM-DD)"
// @Param first_name formData string false "First Name"
// @Param last_name formData string false "Last Name"
// @Param display_name formData string true "Display Name"
// @Param description formData string false "Description"
// @Param avatar formData file false "Avatar File"
// @Param tier_id formData int64 true "Tier ID"
// @Security     BearerAuth
// @Success 200 {object} dto.UserDetailDto
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Failure 401 {object} dto.ErrorResponse "User not authenticated"
// @Router /api/user [put]
func (s *Server) updateUser(ctx *gin.Context) {
	var req dto.UserUpdateRequest

	// Bind form-data request
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userExist, err := s.store.CheckUserExistNotMe(req.ID, req.Username, req.Phone, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if userExist.IsUsername || userExist.IsPhone || userExist.IsEmail {
		var list []string

		if userExist.IsUsername {
			list = append(list, "username")
		}
		if userExist.IsPhone {
			list = append(list, "phone")
		}
		if userExist.IsEmail {
			list = append(list, "email")
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"error": strings.Join(list, ", ") + " already exist"})
		return
	}

	// Retrieve avatar file
	file, _ := ctx.FormFile("avatar")

	// Extract user ID from context
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("user not authenticated")))
		return
	}

	userIDInt64, ok := userID.(int64)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("invalid user ID type")))
		return
	}

	//user
	var birthday *time.Time
	if req.Birthday != "" {
		bConvert, err := config.ConvertStringToDate(req.Birthday)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		birthday = &bConvert
	}

	user, err := s.store.GetUserData(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	user.Username = req.Username
	user.Phone = req.Phone
	user.Email = req.Email
	user.BirthDay = birthday
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.FullName = strings.Join(strings.Fields(req.FirstName+" "+req.LastName), " ")
	user.UpdatedBy = userIDInt64

	//profile
	fileName := ""
	if file != nil {
		fileName, err = config.SaveImage(file, s.config.FileStorage.AvatarFolder)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		user.Profile.Avatar = fileName
	}

	user.Profile.DisplayName = req.DisplayName
	user.Profile.Description = req.Description
	user.Profile.TierId = req.TierId
	user.Profile.UpdatedBy = userIDInt64

	err = s.store.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	userRes, err := s.store.GetUser(user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	if userRes.Avatar != "" {
		userRes.Avatar = config.GetFileUrl(s.config.ApiFile.Url, s.config.FileStorage.RootFolder, s.config.FileStorage.AvatarFolder, userRes.Avatar)
	}

	ctx.JSON(http.StatusOK, userRes)
}

// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Security     BearerAuth
// @Success 200 {object} nil
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/user/{id} [delete]
func (s *Server) deleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("user not authenticated")))
		return
	}

	userIDInt64, ok := userID.(int64)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("invalid user ID type")))
		return
	}

	if err := s.store.DeleteUser(int64(id), userIDInt64); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User successfully deleted"})
}

// @Summary Activate/Deactivate a user
// @Description Activate/Deactivate a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Security     BearerAuth
// @Success 200 {object} nil
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/user/{id}/active [put]
func (s *Server) activeUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("user not authenticated")))
		return
	}

	userIDInt64, ok := userID.(int64)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("invalid user ID type")))
		return
	}

	if err := s.store.ActiveUser(int64(id), userIDInt64); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User successfully activated/deactivated"})
}

// @Summary Get profile
// @Description Get profile
// @Tags users
// @Accept json
// @Produce json
// @Security     BearerAuth
// @Success 200 {object} dto.UserDetailDto
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/user/profile [get]
func (s *Server) getProfile(ctx *gin.Context) {
	// Extract user ID from context
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("user not authenticated")))
		return
	}

	userIDInt64, ok := userID.(int64)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("invalid user ID type")))
		return
	}

	user, err := s.store.GetUser(userIDInt64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if user.Avatar != "" {
		user.Avatar = config.GetFileUrl(s.config.ApiFile.Url, s.config.FileStorage.RootFolder, s.config.FileStorage.AvatarFolder, user.Avatar)
	}

	ctx.JSON(http.StatusOK, user)
}

// @Summary Update profile
// @Description Update profile
// @Tags users
// @Accept json
// @Produce json
// @Param username formData string true "Username (Alphanumeric)"
// @Param phone formData string false "Phone Number"
// @Param email formData string false "Email Address"
// @Param birthday formData string false "Birthday (YYYY-MM-DD)"
// @Param first_name formData string false "First Name"
// @Param last_name formData string false "Last Name"
// @Param display_name formData string true "Display Name"
// @Param description formData string false "Description"
// @Param avatar formData file false "Avatar File"
// @Security     BearerAuth
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Failure 401 {object} dto.ErrorResponse "User not authenticated"
// @Router /api/user/profile [put]
func (s *Server) updateProfile(ctx *gin.Context) {
	var req dto.UserProfileUpdateRequest

	// Bind form-data request
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Extract user ID from context
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("user not authenticated")))
		return
	}

	userIDInt64, ok := userID.(int64)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("invalid user ID type")))
		return
	}

	userExist, err := s.store.CheckUserExistNotMe(userIDInt64, req.Username, req.Phone, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if userExist.IsUsername || userExist.IsPhone || userExist.IsEmail {
		var list []string

		if userExist.IsUsername {
			list = append(list, "username")
		}
		if userExist.IsPhone {
			list = append(list, "phone")
		}
		if userExist.IsEmail {
			list = append(list, "email")
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"error": strings.Join(list, ", ") + " already exist"})
		return
	}

	// Retrieve avatar file
	file, _ := ctx.FormFile("avatar")

	//user
	var birthday *time.Time
	if req.Birthday != "" {
		bConvert, err := config.ConvertStringToDate(req.Birthday)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		birthday = &bConvert
	}

	user, err := s.store.GetUserData(userIDInt64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	user.Username = req.Username
	user.Phone = req.Phone
	user.Email = req.Email
	user.BirthDay = birthday
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.FullName = strings.Join(strings.Fields(req.FirstName+" "+req.LastName), " ")
	user.UpdatedBy = userIDInt64

	//profile
	fileName := ""
	if file != nil {
		fileName, err = config.SaveImage(file, s.config.FileStorage.AvatarFolder)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		user.Profile.Avatar = fileName
	}

	user.Profile.DisplayName = req.DisplayName
	user.Profile.Description = req.Description
	user.Profile.UpdatedBy = userIDInt64

	err = s.store.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	userRes, err := s.store.GetUser(user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	if userRes.Avatar != "" {
		userRes.Avatar = config.GetFileUrl(s.config.ApiFile.Url, s.config.FileStorage.RootFolder, s.config.FileStorage.AvatarFolder, userRes.Avatar)
	}

	ctx.JSON(http.StatusOK, userRes)
}

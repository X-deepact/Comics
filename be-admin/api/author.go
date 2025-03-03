package api

import (
	"comics-admin/dto"
	config "comics-admin/util"
	"net/http"
	"pkg-common/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) authorRoutes() {
	group := s.router.Group("api/author").Use(s.authMiddleware(s.tokenMaker))

	group.POST("", s.CreateAuthor)
	group.GET("/:id", s.GetAuthorById)
	group.GET("", s.GetAuthors)
	group.PUT("", s.UpdateAuthorById)
	group.DELETE("/:id", s.DeleteAuthorById)
}

// @Summary Create author
// @Description Create a new author
// @Tags authors
// @Accept json
// @Produce json
// @Param     Authorization header string true "Bearer authorization token"
// @Param author body dto.AuthorRequest true "Author Request"
// @Security     BearerAuth
// @Success 200 {object} dto.AuthorResponse
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/author [post]
func (s *Server) CreateAuthor(c *gin.Context) {
	var req dto.AuthorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	birthDay, err := config.ConvertStringToDate(req.BirthDay)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	userId, err := s.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	userName, err := s.GetUsernameFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	now := time.Now()
	authorMode := model.AuthorModel{
		Name:      req.Name,
		Biography: req.Biography,
		BirthDay:  &birthDay,
		CreatedAt: &now,
		CreatedBy: userId,
		UpdatedAt: &now,
		UpdatedBy: userId,
	}

	err = s.store.CreateAuthor(&authorMode)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp := dto.AuthorResponse{
		Id:            authorMode.Id,
		Name:          authorMode.Name,
		Biography:     authorMode.Biography,
		BirthDay:      authorMode.BirthDay.Format(time.RFC3339),
		CreatedAt:     authorMode.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     authorMode.UpdatedAt.Format(time.RFC3339),
		CreatedByName: userName,
		UpdatedByName: userName,
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get GetAuthorById
// @Description Get a author by GetAuthorById
// @Tags authors
// @Accept json
// @Produce json
// @Param     Authorization header string true "Bearer authorization token"
// @Param id path int true "Author ID"
// @Security     BearerAuth
// @Success 200 {object} dto.AuthorResponse
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/author/{id} [get]
func (s *Server) GetAuthorById(c *gin.Context) {
	id := c.Param("id")
	if len(id) <= 0 {
		c.JSON(http.StatusBadRequest, errorResponseMessage("id request author invalid"))
		return
	}

	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	author, err := s.store.GetAuthorById(idInt64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	mapUserIdUserName, err := s.store.GetUserNamesByIds([]int64{author.CreatedBy, author.UpdatedBy})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp := dto.AuthorResponse{
		Id:        author.Id,
		Name:      author.Name,
		Biography: author.Biography,
		BirthDay:  author.BirthDay.Format(time.RFC3339),
		CreatedAt: author.CreatedAt.Format(time.RFC3339),
		UpdatedAt: author.UpdatedAt.Format(time.RFC3339),
	}
	if createdByName, ok := mapUserIdUserName[author.CreatedBy]; ok {
		resp.CreatedByName = createdByName
	}
	if updatedByName, ok := mapUserIdUserName[author.UpdatedBy]; ok {
		resp.UpdatedByName = updatedByName
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary List GetAuthors
// @Description List all authors
// @Tags authors
// @Accept json
// @Produce json
// @Param     Authorization header string true "Bearer authorization token"
// @Param name query string false "Name"
// @Param sort_by query string false "Sort By"
// @Param sort query string false "Sort"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Security     BearerAuth
// @Success 200 {object} []dto.AuthorResponse
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/author [get]
func (s *Server) GetAuthors(c *gin.Context) {
	var req dto.RequestQueryFilter
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	req.SortBy = config.GetSortBy(req.SortBy)
	req.Sort = config.GetSortOrder(req.Sort)

	authors, total, err := s.store.GetAuthors(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userIds := []int64{}
	for _, author := range authors {
		userIds = append(userIds, author.CreatedBy, author.UpdatedBy)
	}

	mapUserIdUserName, err := s.store.GetUserNamesByIds(userIds)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp := make([]dto.AuthorResponse, len(authors))
	for i, author := range authors {
		resp[i] = dto.AuthorResponse{
			Id:        author.Id,
			Name:      author.Name,
			Biography: author.Biography,
			BirthDay:  author.BirthDay.Format(time.RFC3339),
			CreatedAt: author.CreatedAt.Format(time.RFC3339),
			UpdatedAt: author.UpdatedAt.Format(time.RFC3339),
		}
		if createdByName, ok := mapUserIdUserName[author.CreatedBy]; ok {
			resp[i].CreatedByName = createdByName
		}
		if updatedByName, ok := mapUserIdUserName[author.UpdatedBy]; ok {
			resp[i].UpdatedByName = updatedByName
		}
	}
	ListResponse(c, req.Page, req.PageSize, int(total), resp)
}

// @Summary Update author by Id
// @Description Update an existing author
// @Tags authors
// @Accept json
// @Produce json
// @Param     Authorization header string true "Bearer authorization token"
// @Param author body dto.AuthorUpdateRequest true "Author Update Request"
// @Security     BearerAuth
// @Success 200 {object} dto.AuthorResponse
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/author [put]
func (s *Server) UpdateAuthorById(c *gin.Context) {
	var req dto.AuthorUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return

	}

	author, err := s.store.GetAuthorById(req.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	isUpdate := false
	if len(req.Name) > 0 {
		isUpdate = true
		author.Name = req.Name
	}

	if len(req.Biography) > 0 {
		isUpdate = true
		author.Biography = req.Biography
	}

	if len(req.BirthDay) > 0 {
		bConvert, err := config.ConvertStringToDate(req.BirthDay)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		isUpdate = true
		author.BirthDay = &bConvert
	}

	var resp *dto.AuthorResponse
	if isUpdate {
		now := time.Now()
		userId, err := s.GetUserIdFromContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		author.UpdatedAt = &now
		author.UpdatedBy = userId

		modelResponse, err := s.store.UpdateAuthor(author)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		resp = &dto.AuthorResponse{
			Id:        modelResponse.Id,
			Name:      modelResponse.Name,
			Biography: modelResponse.Biography,
			BirthDay:  modelResponse.BirthDay.Format(time.RFC3339),
			CreatedAt: modelResponse.CreatedAt.Format(time.RFC3339),
			UpdatedAt: modelResponse.UpdatedAt.Format(time.RFC3339),
		}
	}

	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	c.JSON(http.StatusInternalServerError, errorResponseMessage("error update author"))
}

// @Summary Delete author by Id
// @Description Delete an existing author
// @Tags authors
// @Accept json
// @Produce json
// @Param     Authorization header string true "Bearer authorization token"
// @Param id path int true "Author ID"
// @Security     BearerAuth
// @Success 200 {object} dto.AuthorResponse
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/author/{id} [delete]
func (a *Server) DeleteAuthorById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponseMessage("invalid id"))
		return
	}

	author, err := a.store.DeleteAuthorById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponseMessage(err.Error()))
		return
	}
	mapUserIdUserName, err := a.store.GetUserNamesByIds([]int64{author.CreatedBy, author.UpdatedBy})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	resp := dto.AuthorResponse{
		Id:        author.Id,
		Name:      author.Name,
		Biography: author.Biography,
		BirthDay:  author.BirthDay.Format(time.RFC3339),
		CreatedAt: author.CreatedAt.Format(time.RFC3339),
		UpdatedAt: author.UpdatedAt.Format(time.RFC3339),
	}
	if createdByName, ok := mapUserIdUserName[author.CreatedBy]; ok {
		resp.CreatedByName = createdByName
	}
	if updatedByName, ok := mapUserIdUserName[author.UpdatedBy]; ok {
		resp.UpdatedByName = updatedByName
	}
	c.JSON(http.StatusOK, resp)
}

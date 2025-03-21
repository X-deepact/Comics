package api

import (
	"comics-admin/dto"
	config "comics-admin/util"
	"errors"
	"pkg-common/common"
	"pkg-common/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) genreRouter() {
	group := s.router.Group("api/genre").Use(s.authMiddleware(s.tokenMaker))

	group.POST("", s.createGenre)
	group.GET("/:id", s.getGenre)
	group.GET("", s.getGenres)
	group.PUT("", s.updateGenre)
	group.DELETE("/:id", s.deleteGenre)
}

// @Summary Create a new genre
// @Description Create a new genre
// @Tags genres
// @Accept json
// @Produce json
// @Param genre body dto.GenreCreateRequest true "Genre Request"
// @Security     BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=dto.GenreResponse} "Genre created successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/genre [post]
func (s *Server) createGenre(ctx *gin.Context) {
	var req dto.GenreCreateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	userID, exists := ctx.Get("user_id")
	if !exists {
		config.BuildErrorResponse(ctx, errors.New("user not authenticated"), nil)
		return
	}

	userIDInt64, ok := userID.(int64)
	if !ok {
		config.BuildErrorResponse(ctx, errors.New("invalid user ID type"), nil)
		return
	}

	genre := model.GenreModel{
		Name:      req.Name,
		Position:  req.Position,
		Lang:      req.Language,
		CreatedBy: userIDInt64,
		UpdatedBy: userIDInt64,
	}

	if err := s.store.CreateGenre(&genre); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	genreDTO, err := s.store.GetGenre(genre.Id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, genreDTO)
}

// @Summary Get a genre
// @Description Get a genre by ID
// @Tags genres
// @Accept json
// @Produce json
// @Param id path int true "Genre ID"
// @Security     BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=dto.GenreResponse} "Genre retrieved successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/genre/{id} [get]
func (s *Server) getGenre(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	genre, err := s.store.GetGenre(int64(id))
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, genre)
}

// @Summary List genres
// @Description List all genre
// @Tags genres
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param page_size query int true "Page size (must be between 10 and 50)" min(10) max(50)
// @Param name query string false "Name"
// @Param language query string false "Language"
// @Param sort_by query string false "Sort by"
// @Param sort query string false "Sort order (asc, desc)"
// @Security     BearerAuth
// @Success 200 {object} dto.ListSuccessResponse{data=[]dto.GenreResponse} "List genres"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/genre [get]
func (s *Server) getGenres(ctx *gin.Context) {
	var req dto.GenreListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	genres, total, err := s.store.GetGenres(req)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildListResponse(ctx,
		&common.Pagination{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    int(total),
		},
		genres)
}

// @Summary Update a new genre
// @Description Update a new genre
// @Tags genres
// @Accept json
// @Produce json
// @Param genre body dto.GenreUpdateRequest true "Genre Request"
// @Security     BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=dto.GenreResponse} "Genre updated successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/genre [put]
func (s *Server) updateGenre(ctx *gin.Context) {
	var req dto.GenreUpdateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	// Extract user ID from context
	userID, exists := ctx.Get("user_id")
	if !exists {
		config.BuildErrorResponse(ctx, errors.New("user not authenticated"), nil)
		return
	}

	userIDInt64, ok := userID.(int64)
	if !ok {
		config.BuildErrorResponse(ctx, errors.New("invalid user ID type"), nil)
		return
	}

	genre := model.GenreModel{
		Id:        req.ID,
		Name:      req.Name,
		Position:  req.Position,
		Lang:      req.Language,
		UpdatedBy: userIDInt64,
	}

	if err := s.store.UpdateGenre(&genre); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	genreDTO, err := s.store.GetGenre(genre.Id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, genreDTO)
}

// @Summary Delete a genre
// @Description Delete a genre by ID
// @Tags genres
// @Accept json
// @Produce json
// @Param id path int true "Genre ID"
// @Security     BearerAuth
// @Success 200 {object} dto.SuccessResponse "Genre successfully deleted"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/genre/{id} [delete]
func (s *Server) deleteGenre(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	if err := s.store.DeleteGenre(int64(id)); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, nil)
}

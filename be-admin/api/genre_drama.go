package api

import (
	"comics-admin/dto"
	config "comics-admin/util"
	"errors"
	"pkg-common/common"
	"pkg-common/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) genreDramaRouter() {
	group := s.router.Group("api/genres-drama").Use(s.authMiddleware(s.tokenMaker))
	group.POST("", s.createGenreDrama)
	group.PUT("", s.updateGenreDrama)
	group.GET("/:id", s.getGenreDrama)
	group.DELETE("/:id", s.deleteGenreDrama)
	group.GET("", s.listGenreDramas)
}

// @Summary Create a new genre for drama
// @Description Create a new genre for drama with translations
// @Tags genres-drama
// @Accept json
// @Produce json
// @Param genre body dto.GenreDramaCreateRequest true "Genre Drama Request"
// @Security     BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=dto.GenreDramaResponse} "Genre created successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/genres-drama [post]
func (s *Server) createGenreDrama(ctx *gin.Context) {
	var req dto.GenreDramaCreateRequest

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

	now := time.Now()

	// Create genre
	genre := &model.GenreForShortDramaModel{
		Name:      req.Name,
		Position:  req.Position,
		CreatedBy: userIDInt64,
		UpdatedBy: userIDInt64,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	if err := s.store.CreateGenreDrama(genre); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	// Create translations in batch
	translations := make([]*model.GenreTranslationModel, len(req.Translations))
	for i, t := range req.Translations {
		translations[i] = &model.GenreTranslationModel{
			GenreId:        genre.Id,
			Language:       t.Language,
			TranslatedName: t.TranslatedName,
			CreatedBy:      userIDInt64,
			UpdatedBy:      userIDInt64,
			CreatedAt:      &now,
			UpdatedAt:      &now,
		}
	}

	if err := s.store.CreateGenreTranslations(translations); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	// Get complete genre data with translations
	genreDTO, err := s.store.GetGenreDrama(genre.Id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, genreDTO)
}

// @Summary Update a genre for drama
// @Description Update a genre for drama with translations
// @Tags genres-drama
// @Accept json
// @Produce json
// @Param genre body dto.GenreDramaUpdateRequest true "Genre Drama Update Request"
// @Security     BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=dto.GenreDramaResponse} "Genre updated successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/genres-drama [put]
func (s *Server) updateGenreDrama(ctx *gin.Context) {
	var req dto.GenreDramaUpdateRequest

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

	now := time.Now()

	// Update genre
	genre := &model.GenreForShortDramaModel{
		Id:        req.ID,
		Name:      req.Name,
		Position:  req.Position,
		UpdatedBy: userIDInt64,
		UpdatedAt: &now,
	}

	if err := s.store.UpdateGenreDrama(genre); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	// Delete old translations
	if err := s.store.DeleteGenreTranslations(req.ID); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	// Create new translations in batch
	translations := make([]*model.GenreTranslationModel, len(req.Translations))
	for i, t := range req.Translations {
		translations[i] = &model.GenreTranslationModel{
			GenreId:        req.ID,
			Language:       t.Language,
			TranslatedName: t.TranslatedName,
			CreatedBy:      userIDInt64,
			UpdatedBy:      userIDInt64,
			CreatedAt:      &now,
			UpdatedAt:      &now,
		}
	}

	if err := s.store.CreateGenreTranslations(translations); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	// Get updated genre data with translations
	genreDTO, err := s.store.GetGenreDrama(req.ID)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, genreDTO)
}

// @Summary Get genre details
// @Description Get genre details by ID with translations
// @Tags genres-drama
// @Accept json
// @Produce json
// @Param id path int true "Genre ID"
// @Security     BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=dto.GenreDramaResponse} "Genre details retrieved successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 404 {object} dto.ErrorResponse "Genre not found"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/genres-drama/{id} [get]
func (s *Server) getGenreDrama(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		config.BuildErrorResponse(ctx, errors.New("invalid genre ID"), nil)
		return
	}

	genre, err := s.store.GetGenreDrama(id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, genre)
}

// @Summary Delete a genre
// @Description Delete a genre and its translations
// @Tags genres-drama
// @Accept json
// @Produce json
// @Param id path int true "Genre ID"
// @Security     BearerAuth
// @Success 200 {object} dto.SuccessResponse "Genre deleted successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/genres-drama/{id} [delete]
func (s *Server) deleteGenreDrama(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		config.BuildErrorResponse(ctx, errors.New("invalid genre ID"), nil)
		return
	}

	if err := s.store.DeleteGenreDrama(id); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, nil)
}

// @Summary List genres
// @Description List all genres with pagination and filtering
// @Tags genres-drama
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param page_size query int true "Page size (must be between 10 and 50)" minimum(10) maximum(50)
// @Param name query string false "Filter by name"
// @Param language query string false "Filter by language"
// @Param sort_by query string false "Sort by field"
// @Param sort query string false "Sort direction (asc/desc)"
// @Security     BearerAuth
// @Success 200 {object} dto.ListSuccessResponse{data=[]dto.GenreDramaResponse} "List of genres"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/genres-drama [get]
func (s *Server) listGenreDramas(ctx *gin.Context) {
	var req dto.GenreDramaListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	genres, total, err := s.store.GetGenreDramas(req)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildListResponse(ctx, &common.Pagination{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    int(total),
	}, genres)
}

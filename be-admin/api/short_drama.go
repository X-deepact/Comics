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

func (s *Server) shortDramaRoutes() {
	group := s.router.Group("/api/dramas").Use(s.authMiddleware(s.tokenMaker))
	{
		group.GET("", s.getDramas)
		group.GET("/:id", s.getDrama)
		group.POST("", s.createDrama)
		group.PUT("", s.updateDrama)
		group.DELETE("/:id", s.deleteDrama)
		group.PUT("/active/:id", s.activeDrama)
	}
}

// @Summary Get list of dramas
// @Description Get list of dramas
// @Tags dramas
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param page_size query int false "Page Size"
// @Param sort_by query string false "Sort by"
// @Param sort query string false "Sort order"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=[]dto.ShortDramaResponse} "List of dramas"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/dramas [get]
func (s *Server) getDramas(c *gin.Context) {
	var req dto.ShortDramaListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		config.BuildErrorResponse(c, err, nil)
		return
	}

	dramas, total, err := s.store.GetDramas(req)
	if err != nil {
		config.BuildErrorResponse(c, err, nil)
		return
	}

	if len(dramas) == 0 {
		config.BuildErrorResponse(c, errors.New("no dramas"), nil)
		return
	}
	result := make([]*dto.ShortDramaResponse, len(dramas))
	for i, drama := range dramas {
		result[i] = &dto.ShortDramaResponse{
			ShortDrama:   *drama,
			Translations: s.mapDramaTranslations(drama.Id),
			Genres:       s.mapDramaGenres(drama.Id),
		}
		result[i].ShortDrama.Thumb = s.mapThumb(drama.Thumb)
	}

	config.BuildListResponse(c, &common.Pagination{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    int(total),
	}, result)
}

// @Summary Get a drama
// @Description Get a drama
// @Tags dramas
// @Accept json
// @Produce json
// @Param id path int true "Drama ID"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=dto.ShortDramaResponse} "Drama retrieved successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/dramas/{id} [get]
func (s *Server) getDrama(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	drama, err := s.store.GetDrama(int64(id))
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	translations, err := s.store.GetDramaTranslations(int64(id))
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	genres, err := s.store.GetDramaGenres(int64(id))
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	resp := dto.ShortDramaResponse{
		ShortDrama:   *drama,
		Translations: translations,
		Genres:       genres,
	}
	resp.ShortDrama.Thumb = s.mapThumb(drama.Thumb)
	config.BuildSuccessResponse(ctx, resp)
}

// @Summary Create a drama
// @Description Create a drama
// @Tags dramas
// @Accept json
// @Produce json
// @Param drama body dto.ShortDramaRequest true "Drama Request"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=dto.ShortDramaResponse} "Drama created successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/dramas [post]
func (s *Server) createDrama(ctx *gin.Context) {
	var req dto.ShortDramaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	userIDInt64, err := ExtractUserID(ctx)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	releaseDate, err := config.ConvertStringToDate(req.ReleaseDate)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	now := time.Now()
	drama := model.ShortDramaModel{
		ReleaseDate: &releaseDate,
		Thumb:       req.Thumb,
		CreatedBy:   userIDInt64,
		UpdatedBy:   userIDInt64,
		Active:      config.Bool(false),
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	if err := s.store.CreateDrama(&drama); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	translations := make([]*model.DramaTranslationModel, len(req.DramaTranslations))
	for i, translation := range req.DramaTranslations {
		dramaTranslation := model.DramaTranslationModel{
			DramaId:     drama.Id,
			Language:    translation.Language,
			Title:       translation.Title,
			Description: translation.Description,
			CreatedBy:   userIDInt64,
			UpdatedBy:   userIDInt64,
		}
		translations[i] = &dramaTranslation
	}
	if err := s.store.CreateDramaTranslations(translations); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	genres := make([]*model.DramaGenreModel, len(req.Genres))
	for i, genreID := range req.Genres {
		dramaGenre := model.DramaGenreModel{
			DramaId:   drama.Id,
			GenreId:   genreID,
			CreatedBy: userIDInt64,
			UpdatedBy: userIDInt64,
		}
		genres[i] = &dramaGenre

	}
	if err := s.store.CreateDramaGenres(genres); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	resp := s.mapShortDramaFromEntityToResponse(&drama)
	config.BuildSuccessResponse(ctx, resp)
}

// @Summary Update a drama
// @Description Update a drama
// @Tags dramas
// @Accept json
// @Produce json
// @Param drama body dto.UpdateShortDramaRequest true "Drama Update Request"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=dto.ShortDramaResponse} "Drama updated successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/dramas [put]
func (s *Server) updateDrama(ctx *gin.Context) {
	var req dto.UpdateShortDramaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	userIDInt64, err := ExtractUserID(ctx)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	drama := model.ShortDramaModel{
		Id:        req.Id,
		Thumb:     req.Thumb,
		UpdatedBy: userIDInt64,
	}

	if req.ReleaseDate != "" {
		releaseDate, err := config.ConvertStringToDate(req.ReleaseDate)
		if err != nil {
			config.BuildErrorResponse(ctx, err, nil)
			return
		}
		drama.ReleaseDate = &releaseDate
	}

	if err := s.store.UpdateDrama(&drama); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	if err := s.store.DeleteDramaTranslations(drama.Id); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	translations := make([]*model.DramaTranslationModel, len(req.DramaTranslations))
	for i, translation := range req.DramaTranslations {
		dramaTranslation := model.DramaTranslationModel{
			DramaId:     drama.Id,
			Language:    translation.Language,
			Title:       translation.Title,
			Description: translation.Description,
			CreatedBy:   userIDInt64,
			UpdatedBy:   userIDInt64,
		}
		translations[i] = &dramaTranslation
	}
	if err := s.store.CreateDramaTranslations(translations); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	if err := s.store.DeleteDramaGenres(drama.Id); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	genres := make([]*model.DramaGenreModel, len(req.Genres))
	for i, genreID := range req.Genres {
		dramaGenre := model.DramaGenreModel{
			DramaId:   drama.Id,
			GenreId:   genreID,
			CreatedBy: userIDInt64,
			UpdatedBy: userIDInt64,
		}
		genres[i] = &dramaGenre
	}

	if err := s.store.CreateDramaGenres(genres); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	dramaDTO, err := s.store.GetDrama(drama.Id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}
	dramaDTO.Thumb = s.mapThumb(drama.Thumb)
	config.BuildSuccessResponse(ctx, dramaDTO)
}

// @Summary Delete a drama
// @Description Delete a drama by ID
// @Tags dramas
// @Accept json
// @Param id path int true "Drama ID"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse "Drama deleted successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/dramas/{id} [delete]
func (s *Server) deleteDrama(ctx *gin.Context) {
	dramaID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	if err := s.store.DeleteDramaTranslations(int64(dramaID)); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	if err := s.store.DeleteDramaGenres(int64(dramaID)); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	if err := s.store.DeleteDrama(int64(dramaID)); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, gin.H{"message": "Drama deleted successfully"})
}

// @Summary Active a drama
// @Description Active a drama by ID
// @Tags dramas
// @Accept json
// @Param id path int true "Drama ID"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse "Drama activated/deactivated successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/dramas/active/{id} [put]
func (s *Server) activeDrama(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	userID, err := ExtractUserID(ctx)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	if err := s.store.ActiveDrama(int64(id), userID); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, "Drama activated/deactivated successfully")
}

func (s *Server) mapShortDramaFromEntityToResponse(drama *model.ShortDramaModel) *dto.ShortDramaResponse {
	if drama.Thumb != "" {
		drama.Thumb = s.minio.GetFileUrl(s.config.FileStorage.ShortDramaThumbFolder, drama.Thumb)
	}

	resp := &dto.ShortDramaResponse{
		ShortDrama: dto.ShortDrama{
			Id:          drama.Id,
			ReleaseDate: drama.ReleaseDate.Format(time.RFC3339),
			Thumb:       s.mapThumb(drama.Thumb),
			CreatedAt:   drama.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   drama.UpdatedAt.Format(time.RFC3339),
		},
		Translations: s.mapDramaTranslations(drama.Id),
		Genres:       s.mapDramaGenres(drama.Id),
	}
	if drama.Active != nil {
		resp.Active = *drama.Active
	}
	mapIdUserName, err := s.store.GetUserNamesByIds([]int64{drama.CreatedBy, drama.UpdatedBy})
	if err == nil && len(mapIdUserName) > 0 {
		resp.CreatedByName = mapIdUserName[drama.CreatedBy]
		resp.UpdatedByName = mapIdUserName[drama.UpdatedBy]
	}
	return resp
}

func (s *Server) mapThumb(thump string) string {
	if thump != "" {
		return s.minio.GetFileUrl(s.config.FileStorage.ShortDramaThumbFolder, thump)
	}
	return thump
}

func (s *Server) mapDramaTranslations(id int64) []dto.DramaTranslation {
	translations, _ := s.store.GetDramaTranslations(id)
	return translations
}

func (s *Server) mapDramaGenres(id int64) []dto.GenreForShortDramaResponse {
	genres, _ := s.store.GetDramaGenres(id)
	return genres
}

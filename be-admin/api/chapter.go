package api

import (
	"comics-admin/dto"
	config "comics-admin/util"
	"github.com/gin-gonic/gin"
	"pkg-common/common"
	"pkg-common/model"
	"strconv"
)

func (s *Server) chapterRouter() {
	group := s.router.Group("api/chapters").Use(s.authMiddleware(s.tokenMaker))

	group.POST("", s.createChapter)
	group.GET("/:id", s.getChapter)
	group.GET("", s.getChapters)
	group.PUT("", s.updateChapter)
	group.DELETE("/:id", s.deleteChapter)
	group.PUT("/active/:id", s.activeChapter)
}

// @Summary Create a new chapter
// @Description Create a new chapter
// @Tags chapters
// @Accept json
// @Produce json
// @Param chapter body dto.ChapterRequest true "Chapter Request"
// @Security     BearerAuth
// @Success 200 {object} dto.ChapterResponse
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/chapters [post]
func (s *Server) createChapter(ctx *gin.Context) {
	var req dto.ChapterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	userID, err := ExtractUserID(ctx)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	chapter := model.ChapterModel{
		ComicId:   req.ComicId,
		Name:      req.Name,
		Cover:     req.Cover,
		Number:    req.Number,
		Active:    true,
		CreatedBy: userID,
	}

	if err := s.store.CreateChapter(&chapter); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, chapter)
}

// @Summary Get a chapter
// @Description Get a chapter
// @Tags chapters
// @Accept json
// @Produce json
// @Param id path int true "Chapter ID"
// @Security     BearerAuth
// @Success 200 {object} dto.ChapterResponse
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/chapters/{id} [get]
func (s *Server) getChapter(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	chapter, err := s.store.GetChapter(id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, chapter)
}

// @Summary Get chapters
// @Description Get chapters
// @Tags chapters
// @Accept json
// @Produce json
// @Param comic_id query int false "Comic ID"
// @Param order query string false "ASC or DESC"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Security     BearerAuth
// @Success 200 {object} []dto.ChapterResponse
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/chapters [get]
func (s *Server) getChapters(ctx *gin.Context) {
	var req dto.ChapterListRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	chapters, total, err := s.store.ListChapters(req)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	pagination := common.Pagination{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    int(total),
	}

	config.BuildListResponse(ctx, &pagination, chapters)
}

// @Summary Update a chapter
// @Description Update a chapter
// @Tags chapters
// @Accept json
// @Produce json
// @Param chapter body dto.ChapterUpdateRequest true "Chapter Update Request"
// @Security     BearerAuth
// @Success 200 {object} dto.ChapterResponse
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/chapters [put]
func (s *Server) updateChapter(ctx *gin.Context) {
	var req dto.ChapterUpdateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	userId, err := ExtractUserID(ctx)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	chapter := model.ChapterModel{
		Id:        req.Id,
		ComicId:   req.ComicId,
		Name:      req.Name,
		Cover:     req.Cover,
		Number:    req.Number,
		UpdatedBy: userId,
	}

	if err := s.store.UpdateChapter(&chapter); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, chapter)
}

// @Summary Delete a chapter
// @Description Delete a chapter by ID
// @Tags chapters
// @Accept json
// @Param id path int true "Chapter ID"
// @Security     BearerAuth
// @Success 200 {object} dto.ResponseMessage "Record deleted successfully"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/chapters/{id} [delete]
func (s *Server) deleteChapter(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	if err := s.store.DeleteChapter(id); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, "Record deleted successfully")
}

// @Summary Active a chapter
// @Description Active a chapter by ID
// @Tags chapters
// @Accept json
// @Param id path int true "Chapter ID"
// @Security     BearerAuth
// @Success 200 {object} dto.ResponseMessage "Record updated successfully"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/chapters/active/{id} [put]
func (s *Server) activeChapter(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	userID, err := ExtractUserID(ctx)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	if err := s.store.ActiveChapter(id, userID); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, "Record updated successfully")
}

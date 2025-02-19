package api

import (
	"comics-admin/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

	chapter := model.ChapterModel{
		ComicId:   req.ComicId,
		Name:      req.Name,
		Cover:     req.Cover,
		Number:    req.Number,
		Active:    req.Active,
		CreatedBy: userIDInt64,
	}

	if err := s.store.CreateChapter(&chapter); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, chapter)

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
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	chapter, err := s.store.GetChapter(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, chapter)
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
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	chapters, total, err := s.store.ListChapters(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var response []dto.ChapterResponse
	for _, chapter := range chapters {
		response = append(response, dto.ChapterResponse{
			Id:            chapter.Id,
			ComicId:       chapter.ComicId,
			Name:          chapter.Name,
			Cover:         chapter.Cover,
			Number:        chapter.Number,
			Active:        chapter.Active,
			CreatedAt:     chapter.CreatedAt,
			CreatedBy:     chapter.CreatedBy,
			UpdatedAt:     chapter.UpdatedAt,
			UpdatedBy:     chapter.UpdatedBy,
			CreatedByUser: "",
			UpdatedByUser: "",
		})
	}

	ListResponse(ctx, req.Page, req.PageSize, int(total), response)
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

	chapter := model.ChapterModel{
		Id:        req.Id,
		ComicId:   req.ComicId,
		Name:      req.Name,
		Cover:     req.Cover,
		Number:    req.Number,
		Active:    req.Active,
		UpdatedBy: userIDInt64,
	}

	if err := s.store.UpdateChapter(&chapter); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, chapter)
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
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := s.store.DeleteChapter(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseMessage{
		Status:  "success",
		Message: "Record deleted successfully",
	})
}

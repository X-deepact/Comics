package api

import (
	"comics-admin/dto"
	"pkg-common/common"
	"pkg-common/model"
	"strconv"
	"time"

	config "comics-admin/util"

	"github.com/gin-gonic/gin"
)

// @Summary Create chapter item
// @Description Create a new chapter item (page)
// @Tags chapter-items
// @Accept json
// @Produce json
// @Param item body dto.ChapterItemCreateRequest true "Chapter Item Request"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=dto.ChapterItemResponse} "Chapter item created successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/chapter-items [post]
func (s *Server) createChapterItem(ctx *gin.Context) {
	var req dto.ChapterItemCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	userID := getUserIDFromContext(ctx)

	now := time.Now()
	item := &model.ChapterItemModel{
		ChapterId:  req.ChapterId,
		Page:       req.Page,
		Image:      req.Image,
		Active:     req.Active,
		ActiveFrom: req.ActiveFrom,
		CreatedBy:  userID,
		CreatedAt:  &now,
	}

	if err := s.store.CreateChapterItem(item); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	response, err := s.store.GetChapterItem(item.Id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, gin.H{
		"message": "Chapter item created successfully",
		"data":    response,
	})
}

// @Summary Get chapter item
// @Description Get a chapter item by ID
// @Tags chapter-items
// @Accept json
// @Produce json
// @Param id path int true "Chapter Item ID"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=dto.ChapterItemResponse} "Chapter item retrieved successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/chapter-items/{id} [get]
func (s *Server) getChapterItem(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	item, err := s.store.GetChapterItem(id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, gin.H{
		"message": "Get chapter item successfully",
		"data":    item,
	})
}

// @Summary List chapter items
// @Description Get list of chapter items with pagination
// @Tags chapter-items
// @Accept json
// @Produce json
// @Param chapter_id query int true "Chapter ID"
// @Param page query int true "Page number"
// @Param page_size query int true "Page size"
// @Param sort_by query string false "Sort by field (e.g. page, created_at, updated_at)"
// @Param sort query string false "Sort order (ASC/DESC)"
// @Security BearerAuth
// @Success 200 {object} dto.ListSuccessResponse{data=[]dto.ChapterItemResponse} "List chapter items"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/chapter-items [get]
func (s *Server) listChapterItems(ctx *gin.Context) {
	var req dto.ChapterItemListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	items, total, err := s.store.ListChapterItems(req)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	for i := range items {
		items[i].Image = s.minio.GetFileUrl(s.config.FileStorage.ChapterItemFolder, items[i].Image)
	}

	config.BuildListResponse(ctx, &common.Pagination{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    int(total),
	}, items)
}

// @Summary Update chapter item
// @Description Update an existing chapter item
// @Tags chapter-items
// @Accept json
// @Produce json
// @Param item body dto.ChapterItemUpdateRequest true "Chapter Item Update Request"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=dto.ChapterItemResponse} "Chapter item updated successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/chapter-items [put]
func (s *Server) updateChapterItem(ctx *gin.Context) {
	var req dto.ChapterItemUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	userID := getUserIDFromContext(ctx)

	item := &model.ChapterItemModel{
		Id:         req.Id,
		Page:       req.Page,
		Image:      req.Image,
		Active:     req.Active,
		ActiveFrom: req.ActiveFrom,
		UpdatedBy:  userID,
		UpdatedAt:  &time.Time{},
	}

	if err := s.store.UpdateChapterItem(item); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	response, err := s.store.GetChapterItem(item.Id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, gin.H{
		"message": "Chapter item updated successfully",
		"data":    response,
	})
}

// @Summary Delete chapter item
// @Description Delete a chapter item by ID
// @Tags chapter-items
// @Accept json
// @Produce json
// @Param id path int true "Chapter Item ID"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse "Chapter item deleted successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/chapter-items/{id} [delete]
func (s *Server) deleteChapterItem(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	if err := s.store.DeleteChapterItem(id); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, gin.H{
		"message": "Chapter item deleted successfully",
	})
}

// @Summary Upload chapter image
// @Description Upload an image for a chapter item
// @Tags chapter-items
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Image file"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse "Chapter item image uploaded successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/chapter-items/upload-image [post]
func (s *Server) uploadChapterImage(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	fileName := ""

	if file != nil {
		fileName, err = s.minio.SaveImage(file, s.config.FileStorage.ChapterItemFolder)
		if err != nil {
			config.BuildErrorResponse(ctx, err, nil)
			return
		}
	}

	config.BuildSuccessResponse(ctx, gin.H{
		"message": "Chapter item image uploaded successfully",
		"data":    fileName,
	})
}

// getUserIDFromContext extracts user ID from gin context
func getUserIDFromContext(ctx *gin.Context) int64 {
	userID, exists := ctx.Get("user_id")
	if !exists {
		return 0
	}

	// Convert interface{} to int64
	switch v := userID.(type) {
	case int64:
		return v
	case uint:
		return int64(v)
	case float64:
		return int64(v)
	default:
		return 0
	}
}

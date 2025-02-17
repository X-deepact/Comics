package api

import (
	"comics-admin/dto"
	"fmt"
	"net/http"
	"pkg-common/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Create chapter item
// @Description Create a new chapter item (page)
// @Tags chapter-items
// @Accept json
// @Produce json
// @Param item body dto.ChapterItemCreateRequest true "Chapter Item Request"
// @Security BearerAuth
// @Success 200 {object} dto.ChapterItemResponse
// @Router /api/chapter-items [post]
func (s *Server) createChapterItem(ctx *gin.Context) {
	var req dto.ChapterItemCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response, err := s.store.GetChapterItem(item.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// @Summary Get chapter item
// @Description Get a chapter item by ID
// @Tags chapter-items
// @Accept json
// @Produce json
// @Param id path int true "Chapter Item ID"
// @Security BearerAuth
// @Success 200 {object} dto.ChapterItemResponse
// @Router /api/chapter-items/{id} [get]
func (s *Server) getChapterItem(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	item, err := s.store.GetChapterItem(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, item)
}

// @Summary List chapter items
// @Description Get list of chapter items with pagination
// @Tags chapter-items
// @Accept json
// @Produce json
// @Param chapter_id query int true "Chapter ID"
// @Param page query int true "Page number"
// @Param page_size query int true "Page size"
// @Security BearerAuth
// @Success 200 {object} []dto.ChapterItemResponse
// @Router /api/chapter-items [get]
func (s *Server) listChapterItems(ctx *gin.Context) {
	var req dto.ChapterItemListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	items, total, err := s.store.ListChapterItems(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"pagination": gin.H{
			"page":      req.Page,
			"page_size": req.PageSize,
			"total":     total,
		},
		"data": items,
	})
}

// @Summary Update chapter item
// @Description Update an existing chapter item
// @Tags chapter-items
// @Accept json
// @Produce json
// @Param item body dto.ChapterItemUpdateRequest true "Chapter Item Update Request"
// @Security BearerAuth
// @Success 200 {object} dto.ChapterItemResponse
// @Router /api/chapter-items [put]
func (s *Server) updateChapterItem(ctx *gin.Context) {
	var req dto.ChapterItemUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response, err := s.store.GetChapterItem(item.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// @Summary Delete chapter item
// @Description Delete a chapter item by ID
// @Tags chapter-items
// @Accept json
// @Produce json
// @Param id path int true "Chapter Item ID"
// @Security BearerAuth
// @Success 200 {object} nil
// @Router /api/chapter-items/{id} [delete]
func (s *Server) deleteChapterItem(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := s.store.DeleteChapterItem(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Chapter item deleted successfully"})
}

// @Summary Upload chapter image
// @Description Upload an image for a chapter item
// @Tags chapter-items
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Image file"
// @Security BearerAuth
// @Success 200 {object} dto.UploadImageResponse
// @Router /api/chapter-items/upload-image [post]
func (s *Server) uploadChapterImage(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Generate unique filename
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)

	// Save file to storage
	dst := fmt.Sprintf("%s/%s", s.config.FileStorage.RootFolder, filename)
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return file URL
	fileURL := fmt.Sprintf("%s/%s", s.config.ApiFile.Url, filename)
	ctx.JSON(http.StatusOK, gin.H{
		"url": fileURL,
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

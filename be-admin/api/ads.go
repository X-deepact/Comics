package api

import (
	"comics-admin/dto"
	config "comics-admin/util"
	"mime/multipart"
	"net/http"
	"pkg-common/common"
	"pkg-common/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new advertisement
// @Description Create a new advertisement
// @Tags ads
// @Accept json
// @Produce json
// @Param ad body dto.AdsCreateRequest true "Advertisement Request"
// @Param status query string false "Filter by status (active/inactive)"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=dto.AdsResponse} "Advertisement created successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/ads [post]
func (s *Server) createAds(ctx *gin.Context) {
	var req dto.AdsCreateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	userID, err := ExtractUserID(ctx)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	now := time.Now()

	ads := &model.AdModel{
		Title:      req.Title,
		Image:      req.Image,
		ComicID:    req.ComicID,
		Type:       req.Type,
		Status:     req.Status,
		DirectURL:  req.DirectURL,
		ActiveFrom: req.ActiveFrom,
		ActiveTo:   req.ActiveTo,
		CreatedBy:  userID,
		CreatedAt:  &now,
	}

	if err := s.store.CreateAds(ads); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	response, err := s.store.GetAds(ads.Id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, gin.H{
		"message": "Advertisement created successfully",
		"data":    response,
	})
}

// @Summary Get advertisement list
// @Description Get list of advertisements with pagination
// @Tags ads
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param page_size query int true "Page size"
// @Param title query string false "Filter by title"
// @Param type query string false "Filter by type (internal/external)"
// @Param status query string false "Filter by status (active/inactive)"
// @Param sort_by query string false "Sort by field (e.g. title, created_at, updated_at)"
// @Param sort query string false "Sort order (ASC/DESC)"
// @Security BearerAuth
// @Success 200 {object} dto.ListSuccessResponse{data=[]dto.AdsResponse} "Get advertisement list successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/ads [get]
func (s *Server) getAdsList(ctx *gin.Context) {
	var req dto.AdsListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	ads, total, err := s.store.GetAdsList(req)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	prefix := s.minio.GetFileUrl(s.config.FileStorage.AdsFolder, "")
	for i := 0; i < len(ads); i++ {
		ads[i].Image = prefix + ads[i].Image
	}

	config.BuildListResponse(ctx, &common.Pagination{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    int(total),
	}, ads)
}

// @Summary Update advertisement
// @Description Update an existing advertisement
// @Tags ads
// @Accept json
// @Produce json
// @Param ad body dto.AdsUpdateRequest true "Advertisement Update Request"
// @Param status query string false "Filter by status (active/inactive)"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=dto.AdsResponse} "Advertisement updated successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/ads [put]
func (s *Server) updateAds(ctx *gin.Context) {
	var req dto.AdsUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	userID, err := ExtractUserID(ctx)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	now := time.Now()

	ads := &model.AdModel{
		Id:         req.ID,
		Title:      req.Title,
		Image:      req.Image,
		ComicID:    req.ComicID,
		Type:       req.Type,
		Status:     req.Status,
		DirectURL:  req.DirectURL,
		ActiveFrom: req.ActiveFrom,
		ActiveTo:   req.ActiveTo,
		UpdatedBy:  userID,
		UpdatedAt:  &now,
	}

	if err := s.store.UpdateAds(ads); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	response, err := s.store.GetAds(ads.Id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, gin.H{
		"message": "Advertisement updated successfully",
		"data":    response,
	})
}

// @Summary Delete advertisement
// @Description Delete an advertisement by ID
// @Tags ads
// @Accept json
// @Produce json
// @Param id path int true "Advertisement ID"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse "Advertisement deleted successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/ads/{id} [delete]
func (s *Server) deleteAds(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	if err := s.store.DeleteAds(id); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, gin.H{
		"message": "Advertisement successfully deleted",
	})
}

// @Summary Update advertisement status
// @Description Update the status of an advertisement to active or inactive
// @Tags ads
// @Accept json
// @Produce json
// @Param id path int true "Advertisement ID"
// @Param status body dto.AdsUpdateStatusRequest true "Status Request"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=dto.AdsResponse} "Advertisement status updated successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/ads/{id}/status [patch]
func (s *Server) updateAdsStatus(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	var req dto.AdsUpdateStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	if err := s.store.UpdateAdsStatus(id, req.Status); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	// Get updated ads information
	response, err := s.store.GetAds(id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, gin.H{
		"message": "Advertisement status updated successfully",
		"data":    response,
	})
}

// @Summary Get advertisement by ID
// @Description Get detailed information of an advertisement by ID
// @Tags ads
// @Accept json
// @Produce json
// @Param id path int true "Advertisement ID"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=dto.AdsResponse} "Get advertisement successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/ads/{id} [get]
func (s *Server) getAdsById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	response, err := s.store.GetAds(id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	response.Image = s.minio.GetFileUrl(s.config.FileStorage.AdsFolder, response.Image)

	config.BuildSuccessResponse(ctx, gin.H{
		"message": "Get advertisement successfully",
		"data":    response,
	})
}

// @Summary Upload advertisement image
// @Description Upload an image for advertisement
// @Tags ads
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Image file"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse "Advertisement image uploaded successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/ads/upload-image [post]
func (s *Server) uploadAdsImage(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	fileName := ""

	if file != nil {
		fileName, err = s.minio.SaveImage(file, s.config.FileStorage.AdsFolder)
		if err != nil {
			config.BuildErrorResponse(ctx, err, nil)
			return
		}
	}

	config.BuildSuccessResponse(ctx, gin.H{
		"message": "Advertisement image uploaded successfully",
		"data":    fileName,
	})
}

// isImageFile checks if the file is an image based on its content type
func (s *Server) isImageFile(file *multipart.FileHeader) bool {
	// Get file content type
	openedFile, err := file.Open()
	if err != nil {
		return false
	}
	defer openedFile.Close()

	// Read first 512 bytes to detect content type
	buffer := make([]byte, 512)
	_, err = openedFile.Read(buffer)
	if err != nil {
		return false
	}

	// Check content type
	contentType := http.DetectContentType(buffer)
	validTypes := []string{"image/jpeg", "image/png", "image/gif", "image/webp"}

	for _, t := range validTypes {
		if contentType == t {
			return true
		}
	}
	return false
}

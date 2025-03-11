package api

import (
	"comics-admin/dto"
	config "comics-admin/util"
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
// @Success 200 {object} dto.ResponseMessage{data=dto.AdsResponse} "Advertisement created successfully"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
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
// @Success 200 {object} dto.ResponseMessage{data=[]dto.AdsResponse} "Get advertisement list successfully"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
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
// @Success 200 {object} dto.ResponseMessage{data=dto.AdsResponse} "Advertisement updated successfully"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
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
// @Success 200 {object} dto.ResponseMessage "Advertisement deleted successfully"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
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
// @Success 200 {object} dto.ResponseMessage{data=dto.AdsResponse} "Advertisement status updated successfully"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
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
// @Success 200 {object} dto.ResponseMessage{data=dto.AdsResponse} "Get advertisement successfully"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 404 {object} dto.ResponseMessage "Advertisement not found"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
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

	config.BuildSuccessResponse(ctx, gin.H{
		"message": "Get advertisement successfully",
		"data":    response,
	})
}

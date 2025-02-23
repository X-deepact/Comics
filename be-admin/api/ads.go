package api

import (
	"comics-admin/dto"
	"net/http"
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
// @Success 200 {object} dto.AdsResponse
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/ads [post]
func (s *Server) createAds(ctx *gin.Context) {
	var req dto.AdsCreateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := ExtractUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response, err := s.store.GetAds(ads.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
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
// @Security BearerAuth
// @Success 200 {object} []dto.AdsResponse
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/ads [get]
func (s *Server) getAdsList(ctx *gin.Context) {
	var req dto.AdsListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ads, total, err := s.store.GetAdsList(req)
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
		"data": ads,
	})
}

// @Summary Update advertisement
// @Description Update an existing advertisement
// @Tags ads
// @Accept json
// @Produce json
// @Param ad body dto.AdsUpdateRequest true "Advertisement Update Request"
// @Param status query string false "Filter by status (active/inactive)"
// @Security BearerAuth
// @Success 200 {object} dto.AdsResponse
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/ads [put]
func (s *Server) updateAds(ctx *gin.Context) {
	var req dto.AdsUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := ExtractUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response, err := s.store.GetAds(ads.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
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
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := s.store.DeleteAds(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseMessage{Status: "success", Message: "Advertisement successfully deleted"})
}

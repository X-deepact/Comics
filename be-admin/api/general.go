package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) generalRouter() {
	group := s.router.Group("api/general").Use(s.authMiddleware(s.tokenMaker))

	group.GET("/tiers", s.getTiers)
}

// @Summary Get tiers
// @Description Get tiers
// @Tags general
// @Accept json
// @Produce json
// @Security     BearerAuth
// @Success 200 {object} dto.TierModel
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/general/tiers [get]
func (s *Server) getTiers(ctx *gin.Context) {
	tiers, err := s.store.GetTiers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tiers)
}

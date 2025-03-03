package api

import (
	"comics-admin/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) generalRouter() {
	group := s.router.Group("api/general").Use(s.authMiddleware(s.tokenMaker))

	group.GET("/tiers", s.getGeneralTiers)
	group.GET("/genres", s.getGeneralGenres)
	group.GET("/authors", s.getGeneralAuthors)
}

// @Summary Get tiers
// @Description Get tiers
// @Tags general
// @Accept json
// @Produce json
// @Security     BearerAuth
// @Success 200 {object} []dto.TierModel
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/general/tiers [get]
func (s *Server) getGeneralTiers(ctx *gin.Context) {
	tiers, err := s.store.GetGeneralTiers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tiers)
}

// @Summary Get general genres
// @Description Get general genres
// @Tags general
// @Accept json
// @Produce json
// @Param name query string false "Name"
// @Param language query string false "Language"
// @Security     BearerAuth
// @Success 200 {object} []dto.GeneralGenreResponse
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/general/genres [get]
func (s *Server) getGeneralGenres(ctx *gin.Context) {
	var req dto.GeneralGenreRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	genres, err := s.store.GetGeneralGenres(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, genres)
}

// @Summary Get general authors
// @Description Get general authors
// @Tags general
// @Accept json
// @Produce json
// @Param name query string false "Name"
// @Security     BearerAuth
// @Success 200 {object} []dto.GeneralAuthorResponse
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/general/authors [get]
func (s *Server) getGeneralAuthors(ctx *gin.Context) {
	var req dto.GeneralAuthorRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authors, err := s.store.GetGeneralAuthors(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, authors)
}

package api

import (
	"comics-admin/dto"
	config "comics-admin/util"
	"github.com/gin-gonic/gin"
)

func (s *Server) generalRouter() {
	group := s.router.Group("api/general").Use(s.authMiddleware(s.tokenMaker))

	group.GET("/tiers", s.getGeneralTiers)
	group.GET("/genres", s.getGeneralGenres)
	group.GET("/authors", s.getGeneralAuthors)
	group.GET("/comics", s.getGeneralComics)
	group.GET("/drama-genres", s.getGeneralDramaGenres)
}

// @Summary Get tiers
// @Description Get tiers
// @Tags general
// @Accept json
// @Produce json
// @Security     BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=[]dto.TierModel} "Tiers"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/general/tiers [get]
func (s *Server) getGeneralTiers(ctx *gin.Context) {
	tiers, err := s.store.GetGeneralTiers()
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, tiers)
}

// @Summary Get general genres
// @Description Get general genres
// @Tags general
// @Accept json
// @Produce json
// @Param name query string false "Name"
// @Param language query string false "Language"
// @Security     BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=[]dto.GeneralGenreResponse} "Genres"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/general/genres [get]
func (s *Server) getGeneralGenres(ctx *gin.Context) {
	var req dto.GeneralGenreRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	genres, err := s.store.GetGeneralGenres(req)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, genres)
}

// @Summary Get general authors
// @Description Get general authors
// @Tags general
// @Accept json
// @Produce json
// @Param name query string false "Name"
// @Security     BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=[]dto.GeneralAuthorResponse} "Authors"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/general/authors [get]
func (s *Server) getGeneralAuthors(ctx *gin.Context) {
	var req dto.GeneralAuthorRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	authors, err := s.store.GetGeneralAuthors(req)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, authors)
}

// @Summary Get general comics
// @Description Get general comics
// @Tags general
// @Accept json
// @Produce json
// @Param name query string false "Name"
// @Param author_id query int false "Author ID"
// @Param genre_id query int false "Genre ID"
// @Param lang query string false "Language"
// @Param original_language query string false "Original language"
// @Security     BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=[]dto.GeneralComicResponse} "Comics"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/general/comics [get]
func (s *Server) getGeneralComics(ctx *gin.Context) {
	var req dto.GeneralComicRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	comics, err := s.store.GetGeneralComics(req)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, comics)

}

// @Summary Get general drama genres
// @Description Get general drama genres
// @Tags general
// @Accept json
// @Produce json
// @Param name query string false "Name"
// @Param language query string false "Language"
// @Security     BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=[]dto.GeneralDramaGenreResponse} "Comics"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/general/drama-genres [get]
func (s *Server) getGeneralDramaGenres(ctx *gin.Context) {
	var req dto.GeneralDramaGenreRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	genres, err := s.store.GetGeneralDramaGenres(req)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	config.BuildSuccessResponse(ctx, genres)

}

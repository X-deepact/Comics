package api

import (
	"comics-admin/dto"
	"comics-admin/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pkg-common/model"
	"time"
)

func (s *Server) comicRouter() {
	group := s.router.Group("api/comics").Use(s.authMiddleware(s.tokenMaker))

	group.POST("", s.createComic)
	group.GET("/:id", s.getComic)
	group.GET("", s.listComics)
	group.PUT("", s.updateComic)
	group.DELETE("/:id", s.deleteComic)
	group.POST("/upload-cover", s.saveCover)
}

// @Summary Create a new comic
// @Description Create a new comic with the provided details
// @Tags comics
// @Accept json
// @Produce json
// @Param comic body dto.ComicRequest true "Comic Request"
// @Security     BearerAuth
// @Success 200 {object} dto.ComicResponse
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Failure 401 {object} dto.ErrorResponse "User not authenticated"
// @Router /api/comics [post]
func (s *Server) createComic(ctx *gin.Context) {
	var req dto.ComicRequest

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

	comic := model.ComicModel{
		Name:        req.Name,
		Code:        req.Code,
		Cover:       req.Cover,
		Description: req.Description,
		Active:      req.Active,
		Lang:        req.Language,
		Audience:    req.Audience,
		CreatedBy:   userIDInt64,
	}

	if err := s.store.CreateComic(&comic); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, dto.ComicResponse{
		ID:          comic.Id,
		Name:        comic.Name,
		Code:        comic.Code,
		Cover:       comic.Cover,
		Description: comic.Description,
		Active:      comic.Active,
		CreatedAt:   comic.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   comic.UpdatedAt.Format(time.RFC3339),
		CreatedBy:   comic.CreatedBy,
		UpdatedBy:   comic.UpdatedBy,
	})
}

// @Summary Get a comic
// @Description Get a comic by ID
// @Tags comics
// @Accept json
// @Produce json
// @Param id path int true "Comic ID"
// @Security     BearerAuth
// @Success 200 {object} dto.ComicResponse
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/comics/{id} [get]
func (s *Server) getComic(ctx *gin.Context) {
	// Extract comic ID from URI
	var uri struct {
		ID int64 `uri:"id" binding:"required,min=1"`
	}

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	comic, err := s.store.GetComic(uri.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, dto.ComicResponse{
		ID:          comic.Id,
		Name:        comic.Name,
		Code:        comic.Code,
		Cover:       comic.Cover,
		Description: comic.Description,
		Active:      comic.Active,
		CreatedAt:   comic.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   comic.UpdatedAt.Format(time.RFC3339),
		CreatedBy:   comic.CreatedBy,
		UpdatedBy:   comic.UpdatedBy,
	})

}

// @Summary List comics
// @Description List all comics
// @Tags comics
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param q query string false "Search query"
// @Param sort_by query string false "Sort by"
// @Param sort query string false "Sort order"
// @Param active query bool false "Active"
// @Param language query string false "Language"
// @Param audience query string false "Audience"
// @Security     BearerAuth
// @Success 200 {object} []dto.ComicResponse
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/comics [get]
func (s *Server) listComics(ctx *gin.Context) {
	var req dto.ComicListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	comics, total, err := s.store.ListComics(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var res []dto.ComicResponse
	for _, comic := range comics {
		res = append(res, dto.ComicResponse{
			ID:          comic.ID,
			Name:        comic.Name,
			Code:        comic.Code,
			Cover:       comic.Cover,
			Description: comic.Description,
			Active:      comic.Active,
			CreatedAt:   comic.CreatedAt,
			UpdatedAt:   comic.UpdatedAt,
			CreatedBy:   comic.CreatedBy,
			UpdatedBy:   comic.UpdatedBy,
		})
	}

	ListResponse(ctx, req.Page, req.PageSize, int(total), res)
}

// @Summary Update a comic
// @Description Update a comic with the provided details
// @Tags comics
// @Accept json
// @Produce json
// @Param comic body dto.ComicUpdateRequest true "Comic Update Request"
// @Security     BearerAuth
// @Success 200 {object} dto.ComicResponse
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 401 {object} dto.ErrorResponse "User not authenticated"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/comics [put]
func (s *Server) updateComic(ctx *gin.Context) {
	var req dto.ComicUpdateRequest

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

	userIDUint, ok := userID.(int64)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("invalid user ID type")))
		return
	}

	comic := model.ComicModel{
		Id:          req.ID,
		Name:        req.Name,
		Code:        req.Code,
		Cover:       req.Cover,
		Description: req.Description,
		Active:      req.Active,
		UpdatedBy:   userIDUint,
	}

	if err := s.store.UpdateComic(&comic); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	comicRes, err := s.store.GetComic(comic.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, dto.ComicResponse{
		ID:          comicRes.Id,
		Name:        comicRes.Name,
		Code:        comicRes.Code,
		Cover:       comicRes.Cover,
		Description: comicRes.Description,
		Active:      comicRes.Active,
		CreatedAt:   comicRes.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   comicRes.UpdatedAt.Format(time.RFC3339),
		CreatedBy:   comicRes.CreatedBy,
		UpdatedBy:   comicRes.UpdatedBy,
	})
}

// @Summary Delete a comic
// @Description Delete a comic by ID
// @Tags comics
// @Accept json
// @Produce json
// @Param id path int true "Comic ID"
// @Security     BearerAuth
// @Success 200 {object} nil
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/comics/{id} [delete]
func (s *Server) deleteComic(ctx *gin.Context) {
	var uri struct {
		ID int64 `uri:"id" binding:"required,min=1"`
	}

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := s.store.DeleteComic(uri.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Comic successfully deleted"})
}

// @Summary Upload a comic cover
// @Description Upload a comic cover image
// @Tags comics
// @Accept multipart/form-data
// @Produce json
// @Param cover formData file true "Comic Cover Image"
// @Security     BearerAuth
// @Success 200 {object} dto.UploadImageResponse "Image URL"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/comics/upload-cover [post]
func (s *Server) saveCover(ctx *gin.Context) {
	file, err := ctx.FormFile("cover")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fileLink := ""

	if file != nil {
		fileName, err := config.SaveImage(file, s.config.FileStorage.CoverFolder)
		if err != nil {
			config.BuildErrorResponse(ctx, http.StatusBadRequest, err, nil)
			return
		}

		fileLink = config.GetFileUrl(s.config.ApiFile.Url, s.config.FileStorage.RootFolder, s.config.FileStorage.CoverFolder, fileName)
	}

	ctx.JSON(http.StatusOK, fileLink)
}

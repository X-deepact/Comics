package api

import (
	"comics-admin/dto"
	config "comics-admin/util"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s *Server) comicRouter() {
	group := s.router.Group("api/comics").Use(s.authMiddleware(s.tokenMaker))

	group.POST("", s.createComic)
	group.GET("/:id", s.getComic)
	group.GET("", s.getComics)
	group.PUT("", s.updateComic)
	group.DELETE("/:id", s.deleteComic)
	group.PUT("/:id/active", s.activeComic)
	group.POST("/upload-cover", s.saveCover)
}

// @Summary Create comic
// @Description Create a new comic
// @Tags comics
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Comic name"
// @Param description formData string false "Comic description"
// @Param lang formData string false "Language"
// @Param audience formData string false "Audience"
// @Param genres formData []int false "Genre ID"
// @Param authors formData []int false "Author ID"
// @Param cover formData file false "Comic cover image"
// @Security BearerAuth
// @Success 200 {object} dto.ComicResponse "Comic created successfully"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 401 {object} dto.ResponseMessage "Unauthorized"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/comics [post]
func (s *Server) createComic(ctx *gin.Context) {
	var req dto.ComicRequest

	if err := ctx.ShouldBind(&req); err != nil {
		config.BuildErrorResponse(ctx, http.StatusBadRequest, err, nil)
		return
	}

	file, _ := ctx.FormFile("cover")

	if file != nil {
		fileName, err := config.SaveImage(file, s.config.FileStorage.CoverFolder)
		if err != nil {
			config.BuildErrorResponse(ctx, http.StatusBadRequest, err, nil)
			return
		}
		req.Cover = fileName
	}

	userID, err := ExtractUserID(ctx)
	if err != nil {
		config.BuildErrorResponse(ctx, http.StatusUnauthorized, err, nil)
		return
	}
	req.CreatedBy = userID

	comic, err := s.store.CreateComic(&req)
	if err != nil {
		config.BuildErrorResponse(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	ctx.JSON(http.StatusOK, comic)
}

// @Summary Get a comic
// @Description Get a comic by ID
// @Tags comics
// @Accept json
// @Produce json
// @Param id path int true "Comic ID"
// @Security     BearerAuth
// @Success 200 {object} dto.ComicReturn "Comic found"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/comics/{id} [get]
func (s *Server) getComic(ctx *gin.Context) {
	// Extract comic ID from URI
	var uri struct {
		ID int64 `uri:"id" binding:"required,min=1"`
	}

	if err := ctx.ShouldBindUri(&uri); err != nil {
		config.BuildErrorResponse(ctx, http.StatusBadRequest, err, nil)
		return
	}

	comic, err := s.store.GetComic(uri.ID)
	if err != nil {
		config.BuildErrorResponse(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	userCreate, _ := s.store.GetUser(comic.CreatedBy)
	var userUpdate *dto.UserDetailDto
	if comic.UpdatedBy == comic.CreatedBy {
		userUpdate = userCreate
	} else {
		userUpdate, _ = s.store.GetUser(comic.UpdatedBy)
	}
	if !strings.HasPrefix(comic.Cover, "http") {
		comic.Cover = config.GetFileUrl(s.config.ApiFile.Url, s.config.FileStorage.RootFolder, s.config.FileStorage.CoverFolder, comic.Cover)
	}

	genres, _ := s.store.GetGenresOfAComic(comic.ID)
	authors, _ := s.store.GetAuthorsOfAComic(comic.ID)
	returnComic := dto.ComicReturn{

		ComicResponse: *comic,
		Genres:        genres,
		Authors:       authors,
	}
	if userCreate != nil {
		returnComic.CreatedByUser = *userCreate
	}
	if userUpdate != nil {
		returnComic.UpdatedByUser = *userUpdate
	}

	ctx.JSON(http.StatusOK, returnComic)

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
// @Param sort query string false "Sort order (asc, desc)"
// @Param active query bool false "Active"
// @Param language query string false "Language"
// @Param audience query string false "Audience"
// @Param author query int false "Author ID"
// @Param genre query int false "Genre ID"
// @Security     BearerAuth
// @Success 200 {object} dto.ListResponse{data=dto.ComicReturn} "List of comics"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/comics [get]
func (s *Server) getComics(ctx *gin.Context) {
	var req dto.ComicListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		config.BuildErrorResponse(ctx, http.StatusBadRequest, err, nil)
		return
	}

	if _, ok := ctx.GetQuery("active"); ok {
		req.ActiveValue = true
	}

	comics, total, err := s.store.ListComics(req)
	if err != nil {
		config.BuildErrorResponse(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	comicsReturn := make([]dto.ComicReturn, len(comics))
	user := map[int64]*dto.UserDetailDto{0: nil}

	for i, comic := range comics {
		if !strings.HasPrefix(comic.Cover, "http") {
			comic.Cover = config.GetFileUrl(s.config.ApiFile.Url, s.config.FileStorage.RootFolder, s.config.FileStorage.CoverFolder, comic.Cover)
		}
		comicsReturn[i] = dto.ComicReturn{
			ComicResponse: comic,
		}
		if comic.CreatedBy != 0 {
			if _, ok := user[comic.CreatedBy]; !ok {
				user[comic.CreatedBy], _ = s.store.GetUser(comic.CreatedBy)
			}
			if user[comic.CreatedBy] != nil {

				comicsReturn[i].CreatedByUser = *user[comic.CreatedBy]
			}
		}
		if comic.UpdatedBy != 0 {
			if _, ok := user[comic.UpdatedBy]; !ok {
				user[comic.UpdatedBy], _ = s.store.GetUser(comic.UpdatedBy)
			}
			if user[comic.UpdatedBy] != nil {
				comicsReturn[i].UpdatedByUser = *user[comic.UpdatedBy]
			}
		}
		genres, _ := s.store.GetGenresOfAComic(comic.ID)
		authors, _ := s.store.GetAuthorsOfAComic(comic.ID)
		comicsReturn[i].Genres = genres
		comicsReturn[i].Authors = authors
	}

	ListResponse(ctx, req.Page, req.PageSize, int(total), comicsReturn)
}

// @Summary Update comic
// @Description Update an existing comic
// @Tags comics
// @Accept multipart/form-data
// @Produce json
// @Param id formData int true "Comic ID"
// @Param name formData string true "Comic name"
// @Param description formData string false "Comic description"
// @Param active formData bool false "Active"
// @Param lang formData string false "Language"
// @Param audience formData string false "Audience"
// @Param genres formData []int false "Genre ID"
// @Param authors formData []int false "Author ID"
// @Param cover formData file false "Comic cover image"
// @Security BearerAuth
// @Success 200 {object} dto.ComicResponse "Comic updated successfully"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 401 {object} dto.ResponseMessage "Unauthorized"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/comics [put]
func (s *Server) updateComic(ctx *gin.Context) {
	var req dto.ComicUpdateRequest

	if err := ctx.ShouldBind(&req); err != nil {
		config.BuildErrorResponse(ctx, http.StatusBadRequest, err, nil)
		return
	}

	userID, err := ExtractUserID(ctx)
	if err != nil {
		config.BuildErrorResponse(ctx, http.StatusUnauthorized, err, nil)
		return
	}

	req.UpdatedBy = userID

	file, _ := ctx.FormFile("cover")

	if file != nil {

		fileName, err := config.SaveImage(file, s.config.FileStorage.CoverFolder)
		if err != nil {
			config.BuildErrorResponse(ctx, http.StatusBadRequest, err, nil)
			return
		}
		req.Cover = fileName
	}
	comic, err := s.store.UpdateComic(&req)

	if err != nil {
		config.BuildErrorResponse(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	ctx.JSON(http.StatusOK, &comic)

}

// @Summary Delete a comic
// @Description Delete a comic by ID
// @Tags comics
// @Accept json
// @Produce json
// @Param id path int true "Comic ID"
// @Security     BearerAuth
// @Success 200 {object} dto.ResponseMessage "Comic successfully deleted"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/comics/{id} [delete]
func (s *Server) deleteComic(ctx *gin.Context) {
	var uri struct {
		ID int64 `uri:"id" binding:"required,min=1"`
	}

	if err := ctx.ShouldBindUri(&uri); err != nil {
		config.BuildErrorResponse(ctx, http.StatusBadRequest, err, nil)
		return
	}

	if err := s.store.DeleteComic(uri.ID); err != nil {
		config.BuildErrorResponse(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseMessage{
		Status:  "success",
		Message: "Comic successfully deleted",
	})
}

// @Summary Upload a comic cover
// @Description Upload a comic cover image
// @Tags comics
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Comic cover image"
// @Security     BearerAuth
// @Success 200 {object} dto.ResponseMessage "Comic cover uploaded successfully"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/comics/upload-cover [post]
func (s *Server) saveCover(ctx *gin.Context) {
	file, err := ctx.FormFile("cover")
	if err != nil {
		config.BuildErrorResponse(ctx, http.StatusBadRequest, err, nil)
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

	ctx.JSON(http.StatusOK, dto.ResponseMessage{
		Status:  "success",
		Message: "Comic cover uploaded successfully",
		Data:    fileLink,
	})
}

// @Summary Activate/Deactivate a comic
// @Description Activate/Deactivate a comic by ID
// @Tags comics
// @Accept json
// @Produce json
// @Param id path int true "Comic ID"
// @Security     BearerAuth
// @Success 200 {object} dto.ResponseMessage "Comic activated/deactivated successfully"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 401 {object} dto.ResponseMessage "Unauthorized"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/comics/{id}/active [put]
func (s *Server) activeComic(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		config.BuildErrorResponse(ctx, http.StatusBadRequest, err, nil)
		return
	}

	userID, err := ExtractUserID(ctx)
	if err != nil {
		config.BuildErrorResponse(ctx, http.StatusUnauthorized, err, nil)
		return
	}

	if err := s.store.ActiveComic(int64(id), userID); err != nil {
		config.BuildErrorResponse(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseMessage{Status: "success", Message: "User successfully activated/deactivated"})

}

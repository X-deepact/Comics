package api

import (
	"comics-admin/dto"
	"comics-admin/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	file, _ := ctx.FormFile("cover")
	/*
		fileLink := ""

		if file != nil {
			fileName, err := config.SaveImage(file, s.config.FileStorage.CoverFolder)
			if err != nil {
				config.BuildErrorResponse(ctx, http.StatusBadRequest, err, nil)
				return
			}

			fileLink = config.GetFileUrl(s.config.ApiFile.Url, s.config.FileStorage.RootFolder, s.config.FileStorage.CoverFolder, fileName)
		}

		req.Cover = fileLink

	*/

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
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	req.CreatedBy = userID

	comic, err := s.store.CreateComic(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
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
// @Success 200 {object} dto.ComicReturn
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
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
	returnComic := dto.ComicReturn{

		ComicResponse: *comic,
		CreatedByUser: *userCreate,
		UpdatedByUser: *userUpdate,
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
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if _, ok := ctx.GetQuery("active"); ok {
		req.ActiveValue = true
	}

	comics, total, err := s.store.ListComics(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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
			comicsReturn[i].CreatedByUser = *user[comic.CreatedBy]
		}
		if comic.UpdatedBy != 0 {
			if _, ok := user[comic.UpdatedBy]; !ok {
				user[comic.UpdatedBy], _ = s.store.GetUser(comic.UpdatedBy)
			}
			comicsReturn[i].UpdatedByUser = *user[comic.UpdatedBy]
		}
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
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := ExtractUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	req.UpdatedBy = userID

	file, _ := ctx.FormFile("cover")
	/*
		if file != nil {
			fileLink := ""
			fileName, err := config.SaveImage(file, s.config.FileStorage.CoverFolder)
			if err != nil {
				config.BuildErrorResponse(ctx, http.StatusBadRequest, err, nil)
				return
			}
			fileLink = config.GetFileUrl(s.config.ApiFile.Url, s.config.FileStorage.RootFolder, s.config.FileStorage.CoverFolder, fileName)
			req.Cover = fileLink
		}
	*/
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
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
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := s.store.DeleteComic(uri.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
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
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := ExtractUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if err := s.store.ActiveComic(int64(id), userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseMessage{Status: "success", Message: "User successfully activated/deactivated"})

}

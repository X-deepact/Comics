package api

import (
	config "comics-admin/util"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) uploadDramaRouter() {
	group := s.router.Group("api/upload-short-drama/")

	group.Use(s.authMiddleware(s.tokenMaker))
	group.POST("/video", s.uploadVideoDrama)
	group.POST("/thumbnail", s.uploadThumbnailDrama)
	group.POST("/subtitle", s.uploadSubtitleDrama)
}

// @Summary Upload video for drama
// @Description Upload a video file for a drama
// @Tags upload-short-drama
// @Accept multipart/form-data
// @Produce json
// @Param id formData int true "Drama ID"
// @Param file formData file true "Video file"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=string} "Video uploaded successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/upload-short-drama/video [post]
func (s *Server) uploadVideoDrama(ctx *gin.Context) {
	idStr := ctx.PostForm("id")
	if idStr == "" {
		config.BuildErrorResponse(ctx, errors.New("missing required parameter: id"), nil)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		config.BuildErrorResponse(ctx, fmt.Errorf("invalid id format: %w", err), nil)
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		config.BuildErrorResponse(ctx, fmt.Errorf("failed to retrieve file: %w", err), nil)
		return
	}

	if file != nil {
		fileName, err := s.minio.SaveVideo(file, s.config.FileStorage.ShortDramaFolder, id)
		if err != nil {
			config.BuildErrorResponse(ctx, err, nil)
			return
		}

		if fileName != "" {
			config.BuildSuccessResponse(ctx, fileName)
			return
		}
	}

	config.BuildSuccessResponse(ctx, nil)
}

// @Summary Upload thumbnail for drama
// @Description Upload a thumbnail image for a drama
// @Tags upload-short-drama
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Thumbnail image file"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=string} "Thumbnail uploaded successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/upload-short-drama/thumbnail [post]
func (s *Server) uploadThumbnailDrama(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		config.BuildErrorResponse(ctx, fmt.Errorf("failed to retrieve file: %w", err), nil)
		return
	}

	if file != nil {
		fileName, err := s.minio.SaveImage(file, s.config.FileStorage.ShortDramaThumbFolder)
		if err != nil {
			config.BuildErrorResponse(ctx, err, nil)
			return
		}

		if fileName != "" {
			config.BuildSuccessResponse(ctx, fileName)
			return
		}
	}

	config.BuildSuccessResponse(ctx, nil)
}

// @Summary Upload subtitle for drama
// @Description Upload a subtitle file for a drama
// @Tags upload-short-drama
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Subtitle file"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=string} "Subtitle uploaded successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/upload-short-drama/subtitle [post]
func (s *Server) uploadSubtitleDrama(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		config.BuildErrorResponse(ctx, fmt.Errorf("failed to retrieve file: %w", err), nil)
		return
	}

	if file != nil {
		fileName, err := s.minio.SaveSubtitle(file, s.config.FileStorage.ShortDramaSubFolder)
		if err != nil {
			config.BuildErrorResponse(ctx, err, nil)
			return
		}

		if fileName != "" {
			config.BuildSuccessResponse(ctx, fileName)
			return
		}
	}

	config.BuildSuccessResponse(ctx, nil)
}

package dto

import (
	config "comics-admin/util"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type RecommendCreateRequest struct {
	Title      string `json:"title,omitempty" form:"title,omitempty"`
	Cover      string `json:"cover,omitempty" form:"cover,omitempty"`
	Position   int    `json:"position,omitempty" form:"position,omitempty"`
	ActiveFrom int64  `json:"active_from,omitempty" form:"active_from,omitempty"`
	ActiveTo   int64  `json:"active_to,omitempty" form:"active_to,omitempty"`
}

type RecommendResponse struct {
	Id         int64  `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	Cover      string `json:"cover,omitempty"`
	Position   int    `json:"position,omitempty"`
	ActiveFrom string `json:"active_from,omitempty"`
	ActiveTo   string `json:"active_to,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	CreatedBy  int64  `json:"created_by,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	UpdatedBy  int64  `json:"updated_by,omitempty"`
}

type RecommendUpdateRequest struct {
	Id         int64  `json:"id,omitempty" form:"id,omitempty"`
	Title      string `json:"title,omitempty" form:"title,omitempty"`
	Cover      string `json:"cover,omitempty" form:"cover,omitempty"`
	Position   int    `json:"position,omitempty" form:"position,omitempty"`
	ActiveFrom string `json:"active_from,omitempty" form:"active_from,omitempty"`
	ActiveTo   string `json:"active_to,omitempty" form:"active_to,omitempty"`
}

func (r *RecommendCreateRequest) Bind(ctx *gin.Context, coverFolter string) error {
	if err := ctx.ShouldBindWith(r, binding.Form); err != nil {
		slog.Info(err.Error())
		return err
	}
	file, err := ctx.FormFile("cover")
	if err != nil {
		slog.Error("error bind file", err.Error())
		return err
	}
	if file != nil && file.Size > 0 {
		if err := config.EnsureDir(coverFolter); err != nil {
			slog.Error("Error ensure dir", err.Error())
			return err
		}
		fileName, err := config.SaveImage(file, coverFolter)
		if err != nil {
			slog.Error("Error save image", err.Error())
			return err
		}
		r.Cover = fileName
	}
	return nil
}

func (r *RecommendUpdateRequest) Bind(ctx *gin.Context, coverFolter string) error {
	if err := ctx.ShouldBind(r); err != nil {
		return err

	}
	file, err := ctx.FormFile("cover")
	if err == nil && file != nil && file.Size > 0 {
		fileName, err := config.SaveImage(file, coverFolter)
		if err != nil {
			return err
		}
		r.Cover = fileName
	}
	return nil
}

func (r *RecommendResponse) BindCoverUrl(cf config.Config) {
	if len(r.Cover) <= 0 {
		return
	}
	r.Cover = config.GetFileUrl(cf.ApiFile.Url, cf.FileStorage.RootFolder, cf.FileStorage.RecommendFolder+"/", r.Cover)
}

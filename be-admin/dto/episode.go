package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type EpisodeCreateRequest struct {
	DramaId   int64      `json:"drama_id,omitempty"`
	Number    int64      `json:"number,omitempty"`
	Video     string     `json:"video,omitempty"`
	Active    *bool      `json:"active,omitempty"`
	Subtitles []Subtitle `json:"subtitles,omitempty"`
}

type Subtitle struct {
	Language string `json:"language,omitempty"`
	Url      string `json:"url,omitempty"`
}

type EpisodeResponse struct {
	Id        int64      `json:"id,omitempty"`
	DramaId   int64      `json:"drama_id,omitempty"`
	Number    int64      `json:"number,omitempty"`
	Video     string     `json:"video,omitempty"`
	Active    *bool      `json:"active,omitempty"`
	Subtitles []Subtitle `json:"subtitles,omitempty"`
	CreatedAt string     `json:"created_at,omitempty"`
	UpdatedAt string     `json:"updated_at,omitempty"`
	CreatedBy string     `json:"created_by_name,omitempty"`
	UpdatedBy string     `json:"updated_by_name,omitempty"`
}

type EpisodeListRequest struct {
	ListRequest
	DramaId int64  `form:"drama_id" json:"drama_id"`
	SortBy  string `form:"sort_by" json:"sort_by"`
	Sort    string `form:"sort" json:"sort"`
}

type EpisodeUpdateRequest struct {
	Id        int64      `json:"id,omitempty" binding:"required"`
	DramaId   int64      `json:"drama_id,omitempty"`
	Number    int64      `json:"number,omitempty"`
	Video     string     `json:"video,omitempty"`
	Active    *bool      `json:"active,omitempty"`
	Subtitles []Subtitle `json:"subtitles,omitempty"`
}

func (r *EpisodeUpdateRequest) Bind(ctx *gin.Context) error {
	if err := ctx.ShouldBindWith(r, binding.JSON); err != nil {
		return err
	}
	return nil
}

func (r *EpisodeCreateRequest) Bind(ctx *gin.Context) error {
	if err := ctx.ShouldBindWith(r, binding.JSON); err != nil {
		return err
	}
	return nil
}

func (r *EpisodeListRequest) Bind(ctx *gin.Context) error {
	if err := ctx.ShouldBindWith(r, binding.Form); err != nil {
		return err
	}
	if r.PageSize <= 0 || r.PageSize > 10 {
		r.PageSize = 10
	}
	if r.Page <= 0 {
		r.Page = 1
	}
	if len(r.SortBy) <= 0 {
		r.SortBy = "created_at"
	}
	if len(r.Sort) <= 0 {
		r.Sort = "ASC"
	}
	return nil
}

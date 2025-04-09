package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type EpisodeCreateRequest struct {
	DramaId   int64      `json:"drama_id" binding:"required"`
	Number    int64      `json:"number" binding:"required"`
	Video     string     `json:"video" binding:"required"`
	Active    *bool      `json:"active"`
	Subtitles []Subtitle `json:"subtitles"`
}

type Subtitle struct {
	Language string `json:"language,omitempty" binding:"required"`
	Url      string `json:"url,omitempty" binding:"required"`
}

type EpisodeResponse struct {
	Id        int64      `json:"id,omitempty"`
	DramaId   int64      `json:"drama_id,omitempty"`
	Number    int64      `json:"number,omitempty"`
	Video     string     `json:"video,omitempty"`
	Active    *bool      `json:"active,omitempty"`
	Subtitles []Subtitle `json:"subtitles,omitempty" binding:"required"`
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
	Id        int64      `json:"id"`
	DramaId   int64      `json:"drama_id"`
	Number    int64      `json:"number"`
	Video     string     `json:"video"`
	Active    *bool      `json:"active"`
	Subtitles []Subtitle `json:"subtitles"`
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
	return nil
}

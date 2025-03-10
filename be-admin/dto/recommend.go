package dto

import (
	"errors"

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
	Id            int64                    `json:"id,omitempty"`
	Title         string                   `json:"title,omitempty"`
	Cover         string                   `json:"cover,omitempty"`
	Position      int                      `json:"position,omitempty"`
	ActiveFrom    string                   `json:"active_from,omitempty"`
	ActiveTo      string                   `json:"active_to,omitempty"`
	CreatedAt     string                   `json:"created_at,omitempty"`
	CreatedByName string                   `json:"created_by_name,omitempty"`
	UpdatedAt     string                   `json:"updated_at,omitempty"`
	UpdatedByName string                   `json:"updated_by_name,omitempty"`
	Comic         []ComicRecommendResponse `json:"comics,omitempty"`
}

type ComicRecommendResponse struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type RecommendUpdateRequest struct {
	Id         int64  `json:"id,omitempty" form:"id,omitempty"`
	Title      string `json:"title,omitempty" form:"title,omitempty"`
	Cover      string `json:"cover,omitempty" form:"cover,omitempty"`
	Position   int    `json:"position,omitempty" form:"position,omitempty"`
	ActiveFrom int64  `json:"active_from,omitempty" form:"active_from,omitempty"`
	ActiveTo   int64  `json:"active_to,omitempty" form:"active_to,omitempty"`
}

type RecommendComicRequest struct {
	RecommendId int64 `json:"recommend_id,omitempty" form:"recommend_id,omitempty"`
	ComicId     int64 `json:"comic_id,omitempty" form:"comic_id,omitempty"`
}

type RecommendComicResponse struct {
	Id          int64 `json:"id,omitempty"`
	RecommendId int64 `json:"recommend_id,omitempty"`
	ComicId     int64 `json:"comic_id,omitempty"`
}

func (r *RecommendCreateRequest) Bind(ctx *gin.Context) error {
	if err := ctx.ShouldBindWith(r, binding.Form); err != nil {
		return err
	}
	if len(r.Title) == 0 {
		return errors.New("title is required")
	}

	if r.Position < 0 {
		return errors.New("position is required")
	}

	if r.ActiveFrom <= 0 {
		return errors.New("active_from is required")
	}
	if r.ActiveTo <= 0 {
		return errors.New("active_to is required")
	}

	return nil
}

func (r *RecommendUpdateRequest) Bind(ctx *gin.Context) error {
	if err := ctx.ShouldBindWith(r, binding.Form); err != nil {
		return err
	}
	return nil
}

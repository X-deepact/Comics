package dto

import "github.com/gin-gonic/gin"

type RequestQueryFilter struct {
	Page     int    `form:"page" binding:"required,min=1" json:"page,omitempty"`
	PageSize int    `form:"page_size" binding:"required,min=10,max=50" json:"page_size,omitempty"`
	SortBy   string `form:"sort_by,omitempty" json:"sort_by,omitempty"`
	Sort     string `form:"sort,omitempty" json:"sort,omitempty"`
	Name     string `form:"name,omitempty" json:"name,omitempty"`
	Title    string `form:"title,omitempty" json:"title,omitempty"`
}

func (r *RequestQueryFilter) Bind(ctx *gin.Context) error {
	if err := ctx.ShouldBindQuery(r); err != nil {
		return err
	}
	return nil
}

package api

import (
	"net/http"
	"pkg-common/common"

	"github.com/gin-gonic/gin"
)

func ListResponse(ctx *gin.Context, page, pageSize, total int, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"pagination": gin.H{
			"page":      page,
			"page_size": pageSize,
			"total":     total,
		},
		"data": data,
	})
}

func ListResponseApi(ctx *gin.Context, page, pageSize, total int, data interface{}) {
	pagination := common.Pagination{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}
	ctx.JSON(http.StatusOK, common.ApiResponse(data, &pagination, nil))
}

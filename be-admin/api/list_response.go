package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

package api

import (
	"fmt"
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

func extractUserID(ctx *gin.Context) (int64, error) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		return 0, fmt.Errorf("user not authenticated")
	}

	userIDInt64, ok := userID.(int64)
	if !ok {
		return 0, fmt.Errorf("invalid user ID type")
	}

	return userIDInt64, nil
}

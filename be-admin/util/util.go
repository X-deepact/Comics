package config

import (
	"net/http"
	"pkg-common/common"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	SortBy = "sort_by"
	Sort   = "sort"

	SortAsc      = "asc"
	SortDesc     = "desc"
	MapSortOrder = map[string]bool{
		SortAsc:  true,
		SortDesc: true,
	}

	SortById        = "id"
	SortByUpdatedAt = "updated_at"

	MapSortBy = map[string]bool{
		SortById:        true,
		SortByUpdatedAt: true,
	}
)

/*
	type Response struct {
		Data    interface{} `json:"data,omitempty"`
		Message interface{} `json:"message,omitempty"`
		Code    interface{} `json:"code,omitempty"`
	}
*/
func BuildListResponse(ctx *gin.Context, pagination *common.Pagination, data interface{}) {
	BuildStandardResponse(ctx, common.ApiResponse(data, pagination, nil))
}

func BuildErrorResponse(ctx *gin.Context, err error, body interface{}) {
	BuildStandardResponse(ctx, common.ApiResponse(nil, nil, err))
}

func BuildSuccessResponse(ctx *gin.Context, body interface{}) {
	BuildStandardResponse(ctx, common.ApiResponse(body, nil, nil))
}

func BuildStandardResponse(ctx *gin.Context, resp common.Response) {
	ctx.JSON(http.StatusOK, resp)
}

func GetQuery(ctx *gin.Context, key string) string {
	return ctx.Query(key)
}

func ConvertStringToDate(dateString string) (time.Time, error) {
	return time.Parse("2006-01-02", dateString)
}

func FromLongValueToTime(value int64) *time.Time {
	t := time.UnixMilli(value)
	return &t
}

func GetSortOrder(sortOrder string) string {
	if len(sortOrder) == 0 {
		return Upper(SortAsc)
	}

	if MapSortOrder[Lower(sortOrder)] {
		return Upper(sortOrder)
	}
	return Upper(SortAsc)
}

func GetSortBy(sortBy string) string {
	if len(sortBy) == 0 {
		return SortById
	}
	if MapSortBy[Lower(sortBy)] {
		return sortBy
	}
	return SortById
}

func Upper(str string) string {
	return strings.ToUpper(str)
}

func Lower(str string) string {
	return strings.ToLower(str)
}

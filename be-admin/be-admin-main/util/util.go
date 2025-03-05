package config

import (
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

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Message interface{} `json:"message,omitempty"`
	Code    interface{} `json:"code,omitempty"`
}

func BuildErrorResponse(ctx *gin.Context, status int, err error, body interface{}) {
	BuildStandardResponse(ctx, status, Response{Message: err.Error(), Code: status, Data: body})
}

func BuildSuccessResponse(ctx *gin.Context, status int, body interface{}) {
	BuildStandardResponse(ctx, status, Response{Message: "Successful", Code: status, Data: body})
}

func BuildStandardResponse(ctx *gin.Context, status int, resp Response) {
	ctx.JSON(status, Response{Data: resp.Data, Code: resp.Code, Message: resp.Message})
}

func GetQuery(ctx *gin.Context, key string) string {
	return ctx.Query(key)
}

func ConvertStringToDate(dateString string) (time.Time, error) {
	return time.Parse("2006-01-02", dateString)
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

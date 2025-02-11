package config

import "github.com/gin-gonic/gin"

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

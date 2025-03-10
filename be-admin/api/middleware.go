package api

import (
	"comics-admin/token"
	config "comics-admin/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func (s *Server) authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		authorizationHeader := ctx.GetHeader(s.config.Source.AuthorizationHeaderName)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provide")

			config.BuildErrorResponse(ctx, err, nil)
			ctx.Abort()
			return
		}

		fields := strings.Fields(authorizationHeader)

		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			config.BuildErrorResponse(ctx, err, nil)
			ctx.Abort()
			return
		}

		authorizationType, authorization := strings.ToLower(fields[0]), fields[1]

		if authorizationType != s.config.Source.AuthorizationTypeBearer {
			err := fmt.Errorf("%s authorization type", authorizationType)
			config.BuildErrorResponse(ctx, err, nil)
			ctx.Abort()
			return
		}

		payload, err := tokenMaker.VerifyToken(authorization)

		if err != nil {
			config.BuildErrorResponse(ctx, err, nil)
			ctx.Abort()
			return
		}

		ctx.Set(s.config.Source.AuthorizationPayloadKey, payload)
		ctx.Set("user_id", payload.UserID)
		ctx.Set("username", payload.Username)
		ctx.Next()
	}
}

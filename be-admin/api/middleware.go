package api

import (
	"comics-admin/token"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (s *Server) authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		authorizationHeader := ctx.GetHeader(s.config.Source.AuthorizationHeaderName)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provide")

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)

		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		authorizationType, authorization := strings.ToLower(fields[0]), fields[1]

		if authorizationType != s.config.Source.AuthorizationTypeBearer {
			err := fmt.Errorf("%s authorization type", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		payload, err := tokenMaker.VerifyToken(authorization)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		ctx.Set(s.config.Source.AuthorizationPayloadKey, payload)
		ctx.Set("user_id", payload.UserID)
		ctx.Set("username", payload.Username)
		ctx.Next()
	}
}

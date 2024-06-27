package http

import (
	"strings"
	"toktok-backend/internal/adapter/handler/http/util"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/core/port"

	"github.com/gin-gonic/gin"
)

const (
	// authorizationHeaderKey is the key for authorization header in the request
	authorizationHeaderKey = "authorization"

	// authorizationType is the accepted authorization type
	authorizationType = "bearer"

	// authorizationPayloadKey is the key for authorization payload in the context
	authorizationPayloadKey = "authorization_payload"
)

// authMiddleware is a middleware to check if the user is authenticated
func authMiddleware(t port.TokenService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHader) == 0 {
			util.HandleErr(ctx, domain.ErrEmptyAuthorizationHeader)
			return
		}

		fileds := strings.Fields(authorizationHader)
		if !(len(fileds) == 2) {
			util.HandleErr(ctx, domain.ErrInvalidAuthorizationHeader)
			return
		}

		tokenString := fileds[1]
		payload, err := t.VerifyToken(tokenString)
		if err != nil {
			util.HandleErr(ctx, err)
			return
		}

		if payload.TokenType == domain.RefreshToken {
			util.HandleErr(ctx, err)
			return
		}

		ctx.Set(authorizationHeaderKey, payload)
		ctx.Next()
	}
}

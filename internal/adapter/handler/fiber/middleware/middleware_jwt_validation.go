package middleware

import (
	"strings"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/core/port"
	"toktok-backend/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

const (
	AuthorizationPayloadKey = "authorization_payload"
)

type JWTValidationMiddleware struct {
	jwtService port.JWTService
}

func NewJWTValidationMiddleware(jwtService port.JWTService) *JWTValidationMiddleware {
	return &JWTValidationMiddleware{
		jwtService: jwtService,
	}
}

func (m *JWTValidationMiddleware) ValidateToken(c *fiber.Ctx) error {
	header := c.Get("Authorization")

	// Authorization 헤더의 필드값이 있는가?
	if len(header) == 0 {
		return errors.Wrap(domain.ErrUnauthorized, "err empty authorization header")
	}

	// 잘못된 필드값임을 알림
	fields := strings.Fields(header)
	if len(fields) != 2 {
		return errors.Wrap(domain.ErrUnauthorized, "err empty authorization header")

	}

	// bearer 헤더인지 확인
	authType := strings.ToLower(fields[0])
	if authType != "bearer" {
		return errors.Wrap(domain.ErrUnauthorized, "err invalid authorization header")

	}

	tokenString := fields[1]
	payload, err := m.jwtService.VerifyToken(tokenString)
	if err != nil {
		return err
	}

	if payload.TokenType == domain.RefreshToken {
		return errors.Wrap(domain.ErrUnauthorized, "token type not match")
	}

	c.Locals(AuthorizationPayloadKey, payload)
	return c.Next()
}

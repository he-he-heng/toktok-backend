package middleware

import (
	"strings"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/core/port"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

const (
	AuthorizationPayloadKey = "authorization_payload"
)

type Middleware struct {
	tokenService port.TokenService
}

func New(tokenService port.TokenService) *Middleware {
	return &Middleware{
		tokenService: tokenService,
	}
}

func (m *Middleware) ValidateToken(c *fiber.Ctx) error {

	header := utils.CopyString(c.Get("Authorization"))

	// Authorization 헤더의 필드값이 있는가?
	if len(header) == 0 {
		return domain.ErrEmptyAuthorizationHeader
	}

	// 잘못된 필드값임을 알림
	fields := strings.Fields(header)
	if len(fields) != 2 {
		return domain.ErrInvalidAuthorizationHeader
	}

	// bearer 헤더인지 확인
	authType := strings.ToLower(fields[0])
	if authType != "bearer" {
		return domain.ErrInvalidAuthorizationHeader
	}

	tokenString := fields[1]
	payload, err := m.tokenService.VerifyToken(tokenString)
	if err != nil {
		return err
	}

	if payload.TokenType == domain.RefreshToken {
		return domain.ErrNotAccessToken
	}

	c.Locals(AuthorizationPayloadKey, payload)
	return c.Next()
}

package port

import (
	"context"
	"toktok-backend/internal/core/domain"
)

type TokenServicde interface {
	CreateToken(tokenType domain.TokenType, user *domain.User) (string, error)
	VerifyToken(token string) (*domain.TokenPlayload, error)
}

type AuthService interface {
	Login(ctx context.Context, uid, password string) (accessToken string, refreshToken string, err error)
}

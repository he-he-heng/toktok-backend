package port

import (
	"context"
	"toktok-backend/internal/core/domain"
)

type TokenServicde interface {
	CreateToken(user *domain.User) (string, error)
	VerifyToken(token string) (*domain.TokenPlayload, error)
}

type AuthService interface {
	Login(ctx context.Context, id, password string) (string, error)
}

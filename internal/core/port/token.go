package port

import "toktok-backend/internal/core/domain"

type TokenService interface {
	CreateToken(tokenType domain.TokenType, user *domain.User) (string, error)
	VerifyToken(token string) (*domain.TokenPlayload, error)
}

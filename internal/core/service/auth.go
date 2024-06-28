package service

import (
	"context"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/core/port"
	"toktok-backend/internal/core/util"

	"toktok-backend/pkg/errors"
)

type AuthService struct {
	userRepository port.UserRepository
	tokenService   port.TokenService
}

func NewAuthService(userRepository port.UserRepository, tokenService port.TokenService) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		tokenService:   tokenService,
	}
}

func (s *AuthService) Login(ctx context.Context, uid, password string) (accessToken string, refreshToken string, err error) {
	user, err := s.userRepository.GetUserByUID(ctx, uid)
	if err != nil {
		return "", "", err
	}

	err = util.ComparePassword(password, user.Password)
	if err != nil {
		return "", "", errors.Wrap(domain.ErrInvalidCredentials, err)
	}

	accessToken, err = s.tokenService.CreateToken(domain.AccessToken, user)
	if err != nil {
		return "", "", errors.Wrap(domain.ErrTokenCreation, err)
	}

	refreshToken, err = s.tokenService.CreateToken(domain.RefreshToken, user)
	if err != nil {
		return "", "", errors.Wrap(domain.ErrTokenCreation, err)
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) Refresh(ctx context.Context, token string) (accessToekn string, err error) {
	tokenPayload, err := s.tokenService.VerifyToken(token)
	if err != nil {
		return "", err
	}

	if tokenPayload.TokenType != domain.RefreshToken {
		return "", errors.Wrap(domain.ErrNotRefreshToken, "")
	}

	user := domain.User{
		ID:   tokenPayload.ID,
		Role: tokenPayload.Role,
	}

	accessToken, err := s.tokenService.CreateToken(domain.AccessToken, &user)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

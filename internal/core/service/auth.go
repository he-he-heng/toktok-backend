package service

import (
	"context"
	"errors"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/core/port"
	"toktok-backend/internal/core/util"
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
		if errors.Is(domain.ErrDataNotFound, err) {
			return "", "", err
		}

		return "", "", domain.ErrInternal
	}

	err = util.ComparePassword(password, user.Password)
	if err != nil {
		return "", "", domain.ErrInvalidCredentials
	}

	accessToken, err = s.tokenService.CreateToken(domain.AccessToken, user)
	if err != nil {
		return "", "", domain.ErrTokenCreation
	}

	refreshToken, err = s.tokenService.CreateToken(domain.RefreshToken, user)
	if err != nil {
		return "", "", domain.ErrTokenCreation
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) ReissueToken(ctx context.Context, token string) (accessToekn string, err error) {
	tokenPayload, err := s.tokenService.VerifyToken(token)
	if err != nil {
		return "", err
	}

	if !(tokenPayload.TokenType == domain.RefreshToken) {
		return "", domain.ErrNotRefreshToken
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

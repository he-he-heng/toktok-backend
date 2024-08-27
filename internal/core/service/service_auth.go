package service

import (
	"context"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/core/port"
	"toktok-backend/internal/core/service/utils"
	"toktok-backend/pkg/errors"
)

type AuthService struct {
	userRepository port.UserRepository
	jwtService     port.JWTService
}

func NewAuthService(userRepository port.UserRepository, jwtService port.JWTService) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		jwtService:     jwtService,
	}
}

func (s *AuthService) Login(ctx context.Context, argUser *domain.User) (accessToken string, refreshToken string, err error) {
	user, err := s.userRepository.GetUserByUID(ctx, argUser.UID)
	if err != nil {
		return "", "", err
	}

	err = utils.VerifyPassword(argUser.Password, user.Password)
	if err != nil {
		return "", "", errors.Wrap(domain.ErrInvalidCredentials, err)
	}

	accessToken, err = s.jwtService.CreateToken(domain.AccessToken, user)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = s.jwtService.CreateToken(domain.RefreshToken, user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) Refresh(ctx context.Context, token string) (accessToekn string, err error) {
	tokenPayload, err := s.jwtService.VerifyToken(token)
	if err != nil {
		return "", err
	}

	if tokenPayload.TokenType != domain.RefreshToken {
		return "", errors.Wrap(domain.ErrUnauthorized, "")
	}

	user := domain.User{
		ID:   tokenPayload.Iss,
		Role: tokenPayload.Role,
	}

	accessToken, err := s.jwtService.CreateToken(domain.AccessToken, &user)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

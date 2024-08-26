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

func (s *AuthService) Login(ctx context.Context, uid, password string) (accessToken string, refreshToken string, err error) {
	user, err := s.userRepository.GetUserByUID(ctx, uid)
	if err != nil {
		return "", "", err
	}

	err = utils.VerifyPassword(password, user.Password)
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

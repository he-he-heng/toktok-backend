package dto

import "toktok-backend/internal/core/domain"

type AuthLoginRequest struct {
	Uid      string `json:"uid" validate:"gte=6,lte=18"`
	Password string `json:"password" validate:"gte=6,lte=32"`
}

func (dto *AuthLoginRequest) ToDomainUser() *domain.User {
	ret := domain.User{}

	ret.UID = dto.Uid
	ret.Password = dto.Password

	return &ret
}

type AuthRefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

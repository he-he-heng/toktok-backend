package dto

import (
	"toktok-backend/internal/core/domain"
)

// TODO: using validation
type CreateUserRequest struct {
	Uid      string `json:"uid"`      // require
	Password string `json:"password"` // require
	Role     string `json:"role"`
	Email    string `json:"email"`
	BanState string `json:"banState"`
}

func (dto *CreateUserRequest) ToDomainUser() *domain.User {
	ret := domain.User{}

	ret.UID = dto.Uid
	ret.Password = dto.Password

	if dto.Role != "" {
		ret.Role = domain.UserRoleType(dto.Role)
	}

	if dto.Email != "" {
		ret.Email = &dto.Email
	}

	if dto.BanState != "" {
		ret.BanState = domain.UserBanStateType(dto.BanState)
	}

	return &ret
}

type UpdateUserReqeust struct {
	Uid      string `json:"uid"`
	Password string `json:"password"`
	Email    string `json:"email"`
	BanState string `json:"banState"`
}

func (dto *UpdateUserReqeust) ToDomainUser(id int) *domain.User {
	user := domain.User{
		ID:       id,
		UID:      dto.Uid,
		Password: dto.Password,
		// Email:    &dto.Password,
		BanState: domain.UserBanStateType(dto.BanState),
	}

	if dto.Email != "" {
		user.Email = &dto.Email
	}

	return &user
}

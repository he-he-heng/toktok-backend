package dto

import (
	"toktok-backend/internal/core/domain"
)

// TODO: using validation
type CreateUserRequest struct {
	// validate:"required"
	Uid      string  `json:"uid" validate:"gte=6,lte=18"`
	Password string  `json:"password" validate:"gte=6,lte=32"`
	Email    *string `json:"email" validate:"omitempty,email"`
}

func (dto *CreateUserRequest) ToDomainUser() *domain.User {
	ret := domain.User{}

	ret.UID = dto.Uid
	ret.Password = dto.Password

	if dto.Email != nil {
		ret.Email = dto.Email
	}

	return &ret
}

type UpdateUserReqeust struct {
	Uid      *string `json:"uid" validate:"omitempty,gte=6,lte=18"`
	Password *string `json:"password" validate:"omitempty,gte=6,lte=32"`
	Email    *string `json:"email" validate:"omitempty,email"`
}

func (dto *UpdateUserReqeust) ToDomainUser(id int) *domain.User {
	user := domain.User{
		ID: id,
	}

	if dto.Uid != nil {
		user.UID = *dto.Uid
	}

	if dto.Password != nil {
		user.Password = *dto.Password
	}

	if dto.Email != nil {
		user.Email = dto.Email
	}

	return &user
}

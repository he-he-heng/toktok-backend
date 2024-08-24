package dto

import (
	"toktok-backend/internal/core/domain"
)

type GetUserResponse struct {
	ID       int     `json:"id"`
	UID      string  `json:"uid"`
	Password string  `json:"password"`
	Email    *string `json:"email,omitempty"`
	Role     string  `json:"role"`
	BanState string  `json:"banState"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"deletedAt,omitempty"`
}

func (gur GetUserResponse) Of(domainUser *domain.User) GetUserResponse {
	getUserResponse := GetUserResponse{
		ID:       domainUser.ID,
		UID:      domainUser.Password,
		Password: domainUser.Password,
		Email:    domainUser.Email,
		Role:     string(domainUser.Role),
		BanState: string(domainUser.BanState),
	}

	getUserResponse.CreatedAt = domainUser.CreatedAt.Format("2006-01-02T15:04:05Z")
	getUserResponse.UpdatedAt = domainUser.UpdatedAt.Format("2006-01-02T15:04:05Z")

	if !domainUser.DeletedAt.IsZero() {
		getUserResponse.DeletedAt = domainUser.DeletedAt.Format("2006-01-02T15:04:05Z")
	}

	return getUserResponse
}

type UserListResponse struct {
	Users []GetUserResponse `json:"users"`
}

func (u UserListResponse) Of(domainUsers []*domain.User) (ret UserListResponse) {
	for _, user := range domainUsers {
		ret.Users = append(ret.Users, GetUserResponse{}.Of(user))
	}

	return ret
}

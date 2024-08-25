package dto

import (
	"toktok-backend/internal/core/domain"
)

type User struct {
	Id  int    `json:"id"`
	Uid string `json:"uid"`
	// Password stirng -> empty password
	Email    *string `json:"email,omitempty"`
	Role     string  `json:"role"`
	BanState string  `json:"banState"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"deletedAt,omitempty"`
}

func (u User) Of(domainUser *domain.User) User {
	user := User{
		Id:       domainUser.ID,
		Uid:      domainUser.UID,
		Email:    domainUser.Email,
		Role:     string(domainUser.Role),
		BanState: string(domainUser.BanState),
	}

	user.CreatedAt = domainUser.CreatedAt.Format("2006-01-02T15:04:05Z")
	user.UpdatedAt = domainUser.UpdatedAt.Format("2006-01-02T15:04:05Z")

	if !domainUser.DeletedAt.IsZero() {
		user.DeletedAt = domainUser.DeletedAt.Format("2006-01-02T15:04:05Z")
	}

	return user

}

type GetUserResponse struct {
}

func (GetUserResponse) Of(domainUser *domain.User) User {
	return User{}.Of(domainUser)

}

type UserListResponse struct {
	Users []User `json:"users"`
}

func (u UserListResponse) Of(domainUsers []*domain.User) (ret UserListResponse) {
	for _, user := range domainUsers {
		ret.Users = append(ret.Users, User{}.Of(user))
	}

	return ret
}

type CreateUserResponse struct{}

func (CreateUserRequest) Of(domainUser *domain.User) User {
	return User{}.Of(domainUser)
}

type UpdateUserResponse struct{}

func (UpdateUserResponse) Of(domainUser *domain.User) User {
	return User{}.Of(domainUser)
}

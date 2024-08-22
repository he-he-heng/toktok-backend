package domain

import "time"

type UserRoleType string

const (
	UserRoleAdmin = "admin"
	UserRoleUser  = "user"
)

type UserBanStateType string

const (
	UserBanStateBan   UserBanStateType = "ban"
	UserBanStateUnBan UserBanStateType = "unban"
)

type User struct {
	ID       int
	UID      string
	Password string
	Email    *string
	Role     UserRoleType
	BanState UserBanStateType

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

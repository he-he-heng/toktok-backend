package domain

import "time"

type UserRoleType string

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

type User struct {
	ID       int
	UID      string
	Password string
	Role     UserRoleType
	IsBan    bool

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

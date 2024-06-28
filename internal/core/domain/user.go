package domain

import "time"

type RoleType string

const (
	RoleAdmin RoleType = "admin"
	RoleUser  RoleType = "user"
)

type User struct {
	ID       int
	UID      string
	Password string
	Role     RoleType

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

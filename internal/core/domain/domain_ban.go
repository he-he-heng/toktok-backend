package domain

import "time"

type Ban struct {
	ID int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	IsBan    bool
	BandedAt time.Time

	User *User
}

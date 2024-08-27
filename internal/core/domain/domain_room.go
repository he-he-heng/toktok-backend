package domain

import "time"

type Room struct {
	ID int

	AvatarRelationID int
	FriendRelationID int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

package domain

import "time"

type MessageStateType string

const (
	MessageStateCheck   MessageStateType = "check"
	MessageStateUnCheck MessageStateType = "uncheck"
)

type Message struct {
	ID         int
	AvatarID   int
	RelationID int
	State      MessageStateType
	Content    string
	EnteredAt  time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

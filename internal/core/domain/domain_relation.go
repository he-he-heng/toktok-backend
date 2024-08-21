package domain

import "time"

type RelationStateType string

const (
	RequestFriend RelationStateType = "request-friend"
	Pending       RelationStateType = "pending"
	Friend        RelationStateType = "friend"
	Declined      RelationStateType = "decline"
	Removed       RelationStateType = "remove"
)

type RelationAlertStateType string

const (
	Allow RelationAlertStateType = "allow"
	Deny  RelationAlertStateType = "deny"
)

type Relation struct {
	ID         int
	AvatarID   int
	FriendID   int
	State      RelationStateType
	AlertState RelationAlertStateType

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

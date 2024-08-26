package domain

import "time"

type RelationStateType string

const (
	RelationStateRequestFriend RelationStateType = "request-friend"
	RelationStatePending       RelationStateType = "pending"
	RelationStateFriend        RelationStateType = "friend"
	RelationStateDecline       RelationStateType = "decline"
	RelationStateRemove        RelationStateType = "remove"
)

type RelationAlertStateType string

const (
	RelationAlertStateAllow RelationAlertStateType = "allow"
	RelationAlertStateDeny  RelationAlertStateType = "deny"
)

type Relation struct {
	ID         int
	AvatarID   int
	FriendID   int
	State      RelationStateType
	AlertState RelationAlertStateType

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

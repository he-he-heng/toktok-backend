package domain

import "time"

type RelationStateType string

const (
	Accepted RelationStateType = "accepted" // 친구 요청이 수락된 상태
	Pending  RelationStateType = "pending"  // 친구 요청이 대기 중인 상태
	Declined RelationStateType = "declined" // 친구 요청이 거절된 상태
	Removed  RelationStateType = "removed"  // 친구 관계가 해제된 상태
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

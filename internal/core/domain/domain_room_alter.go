package domain

import "time"

type RoomAlterStateType string

const (
	Allow  RoomAlterStateType = "allow"
	Reject RoomAlterStateType = "reject"
)

type RoomAlter struct {
	ID    int
	State RoomAlterStateType

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	Avatar   *Avatar
	Relation *Relation
}

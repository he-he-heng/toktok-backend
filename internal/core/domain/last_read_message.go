package domain

import "time"

type LastReadMessage struct {
	ID int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	Avatar   *Avatar
	Relation *Relation
	Message  *Message
}

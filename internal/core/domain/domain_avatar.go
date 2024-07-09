package domain

import "time"

type SexType string

const (
	SexMale   SexType = "male"
	SexFemale SexType = "female"
)

type Avatar struct {
	ID        int
	Sex       SexType
	Birthday  time.Time
	Mbti      *string
	Picture   string
	Nickname  string
	Introduce *string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	User *User
}

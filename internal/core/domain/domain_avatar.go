package domain

import "time"

type AvatarSexType string

const (
	SexFemale AvatarSexType = "female"
	SexMale   AvatarSexType = "male"
)

type AvatarStateType string

const (
	Online  AvatarStateType = "online"
	Offline AvatarStateType = "offline"
)

type Avatar struct {
	ID        int
	UserID    int
	Sex       AvatarSexType
	Birthday  string
	Mbti      *string
	Nickname  string
	Picture   string
	Introduce *string
	State     AvatarStateType

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

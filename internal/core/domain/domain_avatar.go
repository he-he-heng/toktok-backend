package domain

import "time"

type AvatarSexType string

const (
	AvatarSexFemale AvatarSexType = "female"
	AvatarSexMale   AvatarSexType = "male"
)

type AvatarStateType string

const (
	AvatarStateOnline  AvatarStateType = "online"
	AvatarStateOffline AvatarStateType = "offline"
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

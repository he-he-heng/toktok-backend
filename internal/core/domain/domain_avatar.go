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

type AvatarPictureType string

const (
	AvatarPictureLightRed    AvatarPictureType = "lightRed"
	AvatarPictureMilkRed     AvatarPictureType = "milkRed"
	AvatarPictureLightOrange AvatarPictureType = "lightOrange"
	AvatarPictureDeepYellow  AvatarPictureType = "deepYellow"
	AvatarPictureDeepPurple  AvatarPictureType = "deepPurple"
	AvatarPictureBluePurple  AvatarPictureType = "bluePurple"
	AvatarPictureLightPink   AvatarPictureType = "lightPink"
	AvatarPictureDeepSky     AvatarPictureType = "deepSky"
)

type Avatar struct {
	ID        int
	UserID    int
	Sex       AvatarSexType
	Birthday  string
	Mbti      *string
	Nickname  string
	Picture   AvatarPictureType
	Introduce *string
	State     AvatarStateType

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

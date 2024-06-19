package model

type SexType string

const (
	// man
	Male SexType = "male"

	// woman
	Female SexType = "woman"
)

type AvatarType string

const (
	RedFace        AvatarType = "red_face"
	OrangeFace     AvatarType = "orange_face"
	OrangeSoftFace AvatarType = "orange_soft_face"
	YellowFace     AvatarType = "yellow_face"
	PurpleFace     AvatarType = "purple_face"
	PinkFace       AvatarType = "pink_face"
	BlueFace       AvatarType = "blue_face"
)

type User struct {
	DeviceIdent string
	PhoneNumber string
	Email       *string
	Sex         SexType
	Birthday    string
	MBTI        *string
	Avatar      AvatarType
	Nickname    string
	Introduce   *string
}

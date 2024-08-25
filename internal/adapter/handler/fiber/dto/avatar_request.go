package dto

type AvatarCreateRequest struct {
	UserId int    `json:"userId"`
	Sex    string `json:"sex" validate:"gte=6,lte=18,oneof=male female"`
}

package dto

import (
	"toktok-backend/internal/core/domain"
)

type CreateAvatarResponse struct {
	Id        int     `json:"id"`
	UserId    int     `json:"userId"`
	Sex       string  `json:"sex"`
	Birthday  string  `json:"birthday"`
	Mbti      *string `json:"mbti,omitempty"`
	Picture   string  `json:"picture"`
	Nickname  string  `json:"nickname"`
	Introduce *string `json:"introduce,omitempty"`
	State     string  `json:"state"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
	DeletedAt string  `json:"deletedAt"`
}

func (dto CreateAvatarResponse) Of(avatar *domain.Avatar) CreateAvatarResponse {
	createAvatarResponse := CreateAvatarResponse{
		Id:       avatar.ID,
		UserId:   avatar.UserID,
		Sex:      string(avatar.Sex),
		Birthday: avatar.Birthday,
		// set Mbti
		Picture:  string(avatar.Picture),
		Nickname: avatar.Nickname,
		// set Nickname
		State: string(avatar.State),
	}

	if avatar.Mbti != nil {
		createAvatarResponse.Mbti = avatar.Mbti
	}

	if avatar.Introduce != nil {
		createAvatarResponse.Introduce = avatar.Introduce
	}

	createAvatarResponse.CreatedAt = avatar.CreatedAt.Format("2006-01-02T15:04:05Z")
	createAvatarResponse.UpdatedAt = avatar.UpdatedAt.Format("2006-01-02T15:04:05Z")

	if !avatar.DeletedAt.IsZero() {
		createAvatarResponse.DeletedAt = avatar.DeletedAt.Format("2006-01-02T15:04:05Z")
	}

	return createAvatarResponse
}

type ListAvatarResponse struct {
	Avatars []CreateAvatarResponse `json:"avatars,omitempty"`
}

func (ListAvatarResponse) Of(avatars []*domain.Avatar) ListAvatarResponse {
	listAvatarResponse := ListAvatarResponse{
		Avatars: make([]CreateAvatarResponse, 0),
	}

	for _, avatar := range avatars {
		listAvatarResponse.Avatars = append(listAvatarResponse.Avatars, CreateAvatarResponse{}.Of(avatar))
	}

	return listAvatarResponse
}

type GetAvatarResponse struct{}

func (GetAvatarResponse) Of(avatar *domain.Avatar) CreateAvatarResponse {
	return CreateAvatarResponse{}.Of(avatar)
}

type UpdateAvatarResponse struct{}

func (UpdateAvatarResponse) Of(avatar *domain.Avatar) CreateAvatarResponse {
	return CreateAvatarResponse{}.Of(avatar)
}

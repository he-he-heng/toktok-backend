package dto

import "toktok-backend/internal/core/domain"

type CreateAvatarRequest struct {
	UserId    int     `json:"userId"`
	Sex       string  `json:"sex" validate:"oneof=male female"`
	Birthday  string  `json:"birthday" validate:"gte=6,lte=4"`
	Mbti      *string `json:"mbti" validate:"gte=4,lte=4,omitempty"`
	Picture   string  `json:"picture" validate:"oneof=lightRed milkRed lightOrange deepYellow deepPurple bluePurple lightPink deepSky"`
	Nickname  string  `json:"nickname" validate:"gte=4,lte=18"`
	Introduce *string `json:"introduce" validate:"gte=1,lte=300,omitempty"`
}

func (dto CreateAvatarRequest) ToDomainAvatar(userId int) *domain.Avatar {
	avatar := domain.Avatar{
		UserID:   userId,
		Sex:      domain.AvatarSexType(dto.Sex),
		Birthday: dto.Birthday,
		// set mbti
		Picture:  domain.AvatarPictureType(dto.Picture),
		Nickname: dto.Nickname,
		// set introduce
	}

	if dto.Mbti != nil {
		avatar.Mbti = dto.Mbti
	}

	if dto.Introduce != nil {
		avatar.Introduce = dto.Introduce
	}

	return &avatar
}

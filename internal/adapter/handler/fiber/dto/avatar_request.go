package dto

import "toktok-backend/internal/core/domain"

type CreateAvatarRequest struct {
	UserId    int     `json:"userId" validate:"gte=1"`
	Sex       string  `json:"sex" validate:"oneof=male female"`
	Birthday  string  `json:"birthday" validate:"gte=8,lte=8"`
	Mbti      *string `json:"mbti" validate:"omitempty,gte=4,lte=4"`
	Picture   string  `json:"picture" validate:"oneof=lightRed milkRed lightOrange deepYellow deepPurple bluePurple lightPink deepSky"`
	Nickname  string  `json:"nickname" validate:"gte=4,lte=18"`
	Introduce *string `json:"introduce" validate:"omitempty,gte=1,lte=300"`
}

func (dto CreateAvatarRequest) ToDomainAvatar() *domain.Avatar {
	avatar := domain.Avatar{
		UserID:   dto.UserId,
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

type UpdateAvatarRequest struct {
	Sex       *string `json:"sex" validate:"omitempty,oneof=male female"`
	Birthday  *string `json:"birthday" validate:"omitempty,gte=8,lte=8"`
	Mbti      *string `json:"mbti" validate:"omitempty,gte=4,lte=4"`
	Picture   *string `json:"picture" validate:"omitempty,oneof=lightRed milkRed lightOrange deepYellow deepPurple bluePurple lightPink deepSky"`
	Nickname  *string `json:"nickname" validate:"omitempty,gte=4,lte=18"`
	Introduce *string `json:"introduce" validate:"omitempty,gte=1,lte=300"`
}

func (dto UpdateAvatarRequest) ToDomainAvatar(id int) *domain.Avatar {
	avatar := domain.Avatar{
		ID: id,
	}

	if dto.Sex != nil {
		avatar.Sex = domain.AvatarSexType(*dto.Sex)
	}

	if dto.Birthday != nil {
		avatar.Birthday = *dto.Birthday
	}

	if dto.Mbti != nil {
		avatar.Mbti = dto.Mbti
	}

	if dto.Picture != nil {
		avatar.Picture = domain.AvatarPictureType(*dto.Picture)
	}

	if dto.Nickname != nil {
		avatar.Nickname = *dto.Nickname
	}

	if dto.Introduce != nil {
		avatar.Introduce = dto.Introduce
	}

	return &avatar
}

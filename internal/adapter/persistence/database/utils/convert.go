package utils

import (
	"toktok-backend/internal/adapter/persistence/database/ent"
	"toktok-backend/internal/core/domain"
)

func ToDomainUser(user *ent.User) *domain.User {
	if user == nil {
		return nil
	}

	retUser := domain.User{
		ID:       user.ID,
		UID:      user.UID,
		Password: user.Password,
		Role:     domain.UserRoleType(user.Role),
		IsBan:    user.IsBan,

		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	return &retUser
}

func ToDomainUsers(users []*ent.User) []*domain.User {
	retUsers := []*domain.User{}
	for _, user := range users {
		retUsers = append(retUsers, ToDomainUser(user))
	}

	return retUsers
}

func ToDomainAvatar(avatar *ent.Avatar) *domain.Avatar {
	if avatar == nil {
		return nil
	}

	retAvatar := domain.Avatar{
		ID:        avatar.ID,
		UserID:    avatar.ID,
		Sex:       domain.AvatarSexType(avatar.Sex),
		Birthday:  avatar.Birthday,
		Mbti:      avatar.Mbti,
		Nickname:  avatar.Nickname,
		Picture:   avatar.Picture,
		Introduce: avatar.Introduce,
		State:     domain.AvatarStateType(avatar.State),

		CreatedAt: avatar.CreatedAt,
		UpdatedAt: avatar.UpdatedAt,
		DeletedAt: avatar.DeletedAt,
	}

	return &retAvatar
}

func ToDomainAvatars(avatars []*ent.Avatar) []*domain.Avatar {
	retAvatars := []*domain.Avatar{}
	for _, avatar := range avatars {
		retAvatars = append(retAvatars, ToDomainAvatar(avatar))
	}

	return retAvatars
}

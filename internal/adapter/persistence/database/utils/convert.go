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
		Email:    user.Email,
		Role:     domain.UserRoleType(user.Role),
		BanState: domain.UserBanStateType(user.BanState),

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

func ToDomainRelation(relation *ent.Relation) *domain.Relation {
	if relation == nil {
		return nil
	}

	retAvatar := domain.Relation{
		ID:         relation.ID,
		AvatarID:   relation.Edges.Avatar.ID,
		FriendID:   relation.Edges.Friend.ID,
		State:      domain.RelationStateType(relation.AlertState),
		AlertState: domain.RelationAlertStateType(relation.AlertState),

		CreatedAt: relation.CreatedAt,
		UpdatedAt: relation.UpdatedAt,
		DeletedAt: relation.DeletedAt,
	}

	return &retAvatar
}

func ToDomainRelations(relations []*ent.Relation) []*domain.Relation {
	retRelations := []*domain.Relation{}
	for _, relation := range relations {
		retRelations = append(retRelations, ToDomainRelation(relation))
	}

	return retRelations
}

func ToDomainMessage(message *ent.Message) *domain.Message {
	retMessage := domain.Message{
		ID:         message.ID,
		AvatarID:   message.Edges.Avatar.ID,
		RelationID: message.Edges.Relation.ID,
		State:      domain.MessageStateType(message.State),
		Content:    message.Content,
		EnteredAt:  message.EnteredAt,

		CreatedAt: message.CreatedAt,
		UpdatedAt: message.UpdatedAt,
		DeletedAt: message.DeletedAt,
	}

	return &retMessage
}

func ToDomainMesssages(messages []*ent.Message) []*domain.Message {
	retMessages := []*domain.Message{}
	for _, message := range messages {
		retMessages = append(retMessages, ToDomainMessage(message))
	}

	return retMessages
}

package repository

import (
	"context"
	"toktok-backend/internal/adapter/persistence/database"
	entavatar "toktok-backend/internal/adapter/persistence/database/ent/avatar"
	"toktok-backend/internal/adapter/persistence/database/utils"
	"toktok-backend/internal/core/domain"
)

type AvatarRepository struct {
	client *database.Client
}

func (r *AvatarRepository) CreateAvatar(ctx context.Context, avatar *domain.Avatar) (*domain.Avatar, error) {
	builder := r.client.Avatar.Create().
		SetUserID(avatar.UserID).
		SetSex(entavatar.Sex(avatar.Sex)).
		SetBirthday(avatar.Birthday).
		SetNickname(avatar.Nickname).
		SetPicture(avatar.Picture).
		SetState(entavatar.State(avatar.State))

	if avatar.Mbti != nil {
		builder.SetBirthday(avatar.Birthday)
	}

	if avatar.Introduce != nil {
		builder.SetIntroduce(*avatar.Introduce)
	}

	createdAvatar, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ToDomainAvatar(createdAvatar), nil
}

func (r *AvatarRepository) GetAvatar(ctx context.Context, id int) (*domain.Avatar, error) {
	queriedAvatar, err := r.client.Avatar.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return utils.ToDomainAvatar(queriedAvatar), nil
}

func (r *AvatarRepository) ListAvatar(ctx context.Context, skip, limit int, filter, order string) ([]*domain.Avatar, error) {
	builder := r.client.Avatar.Query()

	err := utils.HandleOrderAndFilter(builder, filter, order, utils.FieldMap{
		DefaultField: "id",
		Candidates: utils.Candidates{
			"id":      "id",
			"user_id": "user_id",
		},
	})
	if err != nil {
		return nil, err
	}

	avatars, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ToDomainAvatars(avatars), nil
}

func (r *AvatarRepository) UpdateUser(ctx context.Context, avatar *domain.Avatar) (*domain.Avatar, error) {
	builder := r.client.Avatar.UpdateOneID(avatar.ID)

	if utils.Changeable(avatar.Sex) {
		builder.SetSex(entavatar.Sex(avatar.Sex))
	}

	if utils.Changeable(avatar.Birthday) {
		builder.SetBirthday(avatar.Birthday)
	}

	if avatar.Mbti != nil {
		builder.SetMbti(avatar.Birthday)
	}

	if utils.Changeable(avatar.Nickname) {
		builder.SetNickname(avatar.Nickname)
	}

	if utils.Changeable(avatar.Picture) {
		builder.SetPicture(avatar.Picture)
	}

	if avatar.Introduce != nil {
		builder.SetIntroduce(*avatar.Introduce)
	}

	if utils.Changeable(avatar.State) {
		builder.SetPicture(avatar.Picture)
	}

	updatedAvatar, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ToDomainAvatar(updatedAvatar), nil
}

func (r *AvatarRepository) DeleteUser(ctx context.Context, id int) error {
	err := r.client.Avatar.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

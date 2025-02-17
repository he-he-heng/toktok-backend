package repository

import (
	"context"
	"toktok-backend/internal/adapter/persistence/mysql"
	entavatar "toktok-backend/internal/adapter/persistence/mysql/ent/avatar"
	"toktok-backend/internal/adapter/persistence/mysql/utils"

	"toktok-backend/internal/core/domain"

	"entgo.io/ent/dialect/sql"
)

type AvatarRepository struct {
	client *mysql.Client
}

func NewAvatarRepository(client *mysql.Client) *AvatarRepository {
	avatarRepository := AvatarRepository{
		client: client,
	}
	return &avatarRepository
}

func (r *AvatarRepository) CreateAvatar(ctx context.Context, avatar *domain.Avatar) (*domain.Avatar, error) {
	builder := r.client.Avatar.Create().
		SetUserID(avatar.UserID).
		SetSex(entavatar.Sex(avatar.Sex)).
		SetBirthday(avatar.Birthday).
		SetNickname(avatar.Nickname).
		SetPicture(entavatar.Picture(avatar.Picture))

	if avatar.State != "" {
		builder.SetState(entavatar.State(avatar.State))
	}

	if avatar.Mbti != nil {
		builder.SetMbti(*avatar.Mbti)
	}

	if avatar.Introduce != nil {
		builder.SetIntroduce(*avatar.Introduce)
	}

	createdAvatar, err := builder.Save(ctx)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	return utils.ToDomainAvatar(createdAvatar), nil
}

func (r *AvatarRepository) GetAvatar(ctx context.Context, id int) (*domain.Avatar, error) {
	queriedAvatar, err := r.client.Avatar.Get(ctx, id)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	return utils.ToDomainAvatar(queriedAvatar), nil
}

func (r *AvatarRepository) ListAvatar(ctx context.Context, skip, limit int, order, criterion string) ([]*domain.Avatar, error) {
	builder := r.client.Avatar.Query().
		Limit(limit).
		Offset((skip - 1) * limit)

	orderTermOption := sql.OrderAsc()
	if order == "desc" {
		orderTermOption = sql.OrderDesc()
	}

	switch criterion {

	default:
		builder.Order(
			entavatar.ByID(
				orderTermOption,
			),
		)
	}

	avatars, err := builder.All(ctx)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	return utils.ToDomainAvatars(avatars), nil
}

func (r *AvatarRepository) UpdateAvatar(ctx context.Context, avatar *domain.Avatar) (*domain.Avatar, error) {
	builder := r.client.Avatar.UpdateOneID(avatar.ID)

	if avatar.Sex != "" {
		builder.SetSex(entavatar.Sex(avatar.Sex))
	}

	if avatar.Birthday != "" {
		builder.SetBirthday(avatar.Birthday)
	}

	if avatar.Mbti != nil {
		builder.SetMbti(*avatar.Mbti)
	}

	if avatar.Nickname != "" {
		builder.SetNickname(avatar.Nickname)
	}

	if avatar.Picture != "" {
		builder.SetPicture(entavatar.Picture(avatar.Picture))
	}

	if avatar.Introduce != nil {
		builder.SetIntroduce(*avatar.Introduce)
	}

	if avatar.State != "" {
		builder.SetState(entavatar.State(avatar.State))
	}

	updatedAvatar, err := builder.Save(ctx)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	return utils.ToDomainAvatar(updatedAvatar), nil
}

func (r *AvatarRepository) DeleteAvatar(ctx context.Context, id int) error {
	err := r.client.Avatar.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return utils.ErrWrap(err)
	}

	return nil
}

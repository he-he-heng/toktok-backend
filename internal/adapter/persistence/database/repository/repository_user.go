package repository

import (
	"context"
	"toktok-backend/internal/adapter/persistence/database"
	entuser "toktok-backend/internal/adapter/persistence/database/ent/user"
	"toktok-backend/internal/adapter/persistence/database/utils"
	"toktok-backend/internal/core/domain"
)

type UserRepository struct {
	client *database.Client
}

func NewUserRepository(client *database.Client) *UserRepository {
	userRepository := UserRepository{
		client: client,
	}

	return &userRepository
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	builder := r.client.User.Create().
		SetUID(user.UID).
		SetPassword(user.Password).
		SetBanState(entuser.BanState(user.BanState)).
		SetRole(entuser.Role(user.Role))

	if user.Email != nil {
		builder.SetEmail(*user.Email)
	}

	createdUser, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ToDomainUser(createdUser), nil
}

func (r *UserRepository) GetUser(ctx context.Context, id int) (*domain.User, error) {
	queriedUser, err := r.client.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return utils.ToDomainUser(queriedUser), err
}

func (r *UserRepository) ListUser(ctx context.Context, skip, limit int, filter, order string) ([]*domain.User, error) {
	builder := r.client.User.Query().
		Limit(limit).
		Offset((skip - 1) * limit)

	err := utils.HandleOrderAndFilter(builder, filter, order, utils.FieldMap{
		DefaultField: "id",
		Candidates: utils.Candidates{
			"id":  "id",
			"uid": "uid",
		},
	})
	if err != nil {
		return nil, err
	}

	queriedUsers, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ToDomainUsers(queriedUsers), nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	builder := r.client.User.UpdateOneID(user.ID)

	if utils.Changeable(user.UID) {
		builder.SetUID(user.UID)
	}

	if utils.Changeable(user.Password) {
		builder.SetUID(user.Password)
	}

	if utils.Changeable(user.Role) {
		builder.SetUID(user.Password)
	}

	if user.Email != nil {
		builder.SetEmail(*user.Email)
	}

	if utils.Changeable(user.BanState) {
		builder.SetBanState(entuser.BanState(user.BanState))
	}

	updatedUser, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ToDomainUser(updatedUser), nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	err := r.client.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}

	return err
}

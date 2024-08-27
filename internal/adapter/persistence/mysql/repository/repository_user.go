package repository

import (
	"context"
	"toktok-backend/internal/adapter/persistence/mysql"
	entuser "toktok-backend/internal/adapter/persistence/mysql/ent/user"
	"toktok-backend/internal/adapter/persistence/mysql/utils"
	"toktok-backend/internal/core/domain"

	"entgo.io/ent/dialect/sql"
)

type UserRepository struct {
	client *mysql.Client
}

func NewUserRepository(client *mysql.Client) *UserRepository {
	userRepository := UserRepository{
		client: client,
	}

	return &userRepository
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	builder := r.client.User.Create().
		SetUID(user.UID).
		SetPassword(user.Password)

	if user.BanState != "" {
		builder.SetBanState(entuser.BanState(user.BanState))
	}

	if user.Role != "" {
		builder.SetRole(entuser.Role(user.Role))
	}

	if user.Email != nil {
		builder.SetEmail(*user.Email)
	}

	createdUser, err := builder.Save(ctx)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	return utils.ToDomainUser(createdUser), nil
}

func (r *UserRepository) GetUser(ctx context.Context, id int) (*domain.User, error) {
	queriedUser, err := r.client.User.Get(ctx, id)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	return utils.ToDomainUser(queriedUser), err
}

func (r *UserRepository) GetUserByUID(ctx context.Context, uid string) (*domain.User, error) {
	queriedUser, err := r.client.User.Query().
		Where(entuser.UIDEQ(uid)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ToDomainUser(queriedUser), nil
}

func (r *UserRepository) ListUser(ctx context.Context, skip, limit int, order, critertion string) ([]*domain.User, error) {
	builder := r.client.User.Query().
		Limit(limit).
		Offset((skip - 1) * limit)

	orderTermOption := sql.OrderAsc()
	if order == "desc" {
		orderTermOption = sql.OrderDesc()
	}

	switch critertion {
	case "uid":
		builder.Order(
			entuser.ByUID(
				orderTermOption,
			),
		)
	default:
		builder.Order(
			entuser.ByID(
				orderTermOption,
			),
		)
	}

	queriedUsers, err := builder.All(ctx)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	return utils.ToDomainUsers(queriedUsers), nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	builder := r.client.User.UpdateOneID(user.ID)

	if user.UID != "" {
		builder.SetUID(user.UID)
	}

	if user.Password != "" {
		builder.SetPassword(user.Password)
	}

	if user.Role != "" {
		builder.SetRole(entuser.Role(user.Role))
	}

	if user.Email != nil {
		builder.SetEmail(*user.Email)
	}

	if user.BanState != "" {
		builder.SetBanState(entuser.BanState(user.BanState))
	}

	updatedUser, err := builder.Save(ctx)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	return utils.ToDomainUser(updatedUser), nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	err := r.client.User.DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return utils.ErrWrap(err)
	}

	return nil
}

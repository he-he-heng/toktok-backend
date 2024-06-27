package repository

import (
	"context"
	"toktok-backend/internal/adapter/persistence/mysql"
	"toktok-backend/internal/adapter/persistence/mysql/ent"
	entuser "toktok-backend/internal/adapter/persistence/mysql/ent/user"
	"toktok-backend/internal/core/domain"
)

type UserRepository struct {
	db *mysql.Database
}

func NewUserRepository(db *mysql.Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// CreateUser inserts a new user into the database
func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	// fmt.Printf("user : %+v\n", user)
	// fmt.Println("userRole: ", user.Role)
	// fmt.Println("casted userRole: ", entuser.Role(user.Role))
	_, err := r.db.User.
		Create().
		SetUID(user.UID).
		SetPassword(user.Password).
		SetRole(entuser.Role(user.Role)).
		Save(ctx)

	// fmt.Println("err: ", err, "error type: ", reflect.TypeOf(err))

	if ent.IsValidationError(err) {
		return nil, domain.ErrValidation
	}

	if ent.IsConstraintError(err) {
		return nil, domain.ErrConstraint
	}

	return user, nil
}

// GetUserByID selects a user by id
func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	user, err := r.db.User.
		Get(ctx, id)

	if ent.IsNotFound(err) {
		return nil, domain.ErrDataNotFound
	}

	return &domain.User{
		ID:       user.ID,
		UID:      user.UID,
		Password: user.Password,
		Role:     domain.RoleType(user.Role),

		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: &user.DeletedAt,
	}, nil
}

// GetUserByUID selects a user by uid
func (r *UserRepository) GetUserByUID(ctx context.Context, uid string) (*domain.User, error) {

	user, err := r.db.User.
		Query().
		Where(entuser.UID(uid)).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, domain.ErrDataNotFound
	}

	if ent.IsValidationError(err) {
		return nil, domain.ErrValidation
	}

	if ent.IsConstraintError(err) {
		return nil, domain.ErrConstraint
	}

	return &domain.User{
		ID:       user.ID,
		UID:      user.UID,
		Password: user.Password,
		Role:     domain.RoleType(user.Role),

		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: &user.DeletedAt,
	}, nil
}

// ListUsers selects a list of users with pagination
func (r *UserRepository) ListUsers(ctx context.Context, skip, limit int) ([]domain.User, error) {
	users, err := r.db.User.
		Query().
		Order(ent.Desc(entuser.FieldID)).
		Limit(limit).
		Offset((skip - 1) * limit).
		All(ctx)

	if ent.IsNotFound(err) {
		return nil, domain.ErrDataNotFound
	}

	if ent.IsValidationError(err) {
		return nil, domain.ErrValidation
	}

	if ent.IsConstraintError(err) {
		return nil, domain.ErrConstraint
	}

	var domainUsers = make([]domain.User, 0)
	for _, u := range users {
		domainUsers = append(domainUsers, domain.User{
			ID:       u.ID,
			UID:      u.UID,
			Password: u.Password,
			Role:     domain.RoleType(u.Role),

			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			DeletedAt: &u.DeletedAt,
		})
	}

	return domainUsers, nil
}

// UpdateUser updates a user
func (r *UserRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	_, err := r.db.User.
		Update().
		Where(entuser.IDEQ(user.ID)).
		SetUID(user.UID).
		SetPassword(user.Password).
		SetRole(entuser.Role(user.Role)).
		Save(ctx)

	if ent.IsValidationError(err) {
		return nil, domain.ErrValidation
	}

	if ent.IsConstraintError(err) {
		return nil, domain.ErrConstraint
	}

	return user, err
}

// DeleteUser deletes a user
func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	err := r.db.User.
		DeleteOneID(id).
		Exec(ctx)

	if ent.IsValidationError(err) {
		return domain.ErrValidation
	}

	if ent.IsConstraintError(err) {
		return domain.ErrConstraint
	}

	return nil
}

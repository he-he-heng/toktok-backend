package port

import (
	"context"
	"toktok-backend/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUser(ctx context.Context, id int) (*domain.User, error)
	GetUserByUID(ctx context.Context, uid string) (*domain.User, error)
	ListUser(ctx context.Context, skip, limit int, order, criterion string) ([]*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type UserService interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUser(ctx context.Context, id int) (*domain.User, error)
	ListUser(ctx context.Context, skip, limit int, order, criterion string) ([]*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id int) error
}

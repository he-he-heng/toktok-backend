package port

import (
	"context"
	"toktok-backend/internal/core/domain"
)

type AvatarRepository interface {
	CreateAvatar(ctx context.Context, avatar *domain.Avatar) (*domain.Avatar, error)
	GetAvatar(ctx context.Context, id int) (*domain.Avatar, error)
	ListAvatar(ctx context.Context, skip, limit int, filter, order string) ([]*domain.Avatar, error)
	UpdateUser(ctx context.Context, avatar *domain.Avatar) (*domain.Avatar, error)
	DeleteUser(ctx context.Context, id int) error
}

type AvatarService interface{}

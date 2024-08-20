package port

import (
	"context"
	"toktok-backend/internal/core/domain"
)

type AvatarRepository interface {
	CreateAvatar(ctx context.Context, avatar *domain.Avatar) (*domain.User, error)
	GetAvatar(ctx context.Context, id int) (*domain.Avatar, error)
	ListUser(ctx context.Context, skip, limit int) ([]*domain.Avatar, error)
	UpdateUser(ctx context.Context, user *domain.Avatar) (*domain.Avatar, error)
	DeleteUser(ctx context.Context, id int) error
}

type AvatarService interface{}

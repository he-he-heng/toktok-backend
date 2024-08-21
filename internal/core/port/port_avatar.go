package port

import (
	"context"
	"toktok-backend/internal/core/domain"
)

type AvatarRepository interface {
	CreateAvatar(ctx context.Context, avatar *domain.Avatar) (*domain.Avatar, error)
	GetAvatar(ctx context.Context, id int) (*domain.Avatar, error)
	ListAvatar(ctx context.Context, skip, limit int, order, criterion string) ([]*domain.Avatar, error)
	UpdateAvatar(ctx context.Context, avatar *domain.Avatar) (*domain.Avatar, error)
	DeleteAvatar(ctx context.Context, id int) error
}

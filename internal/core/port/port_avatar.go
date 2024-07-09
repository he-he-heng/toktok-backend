package port

import (
	"context"
	"toktok-backend/internal/core/domain"
)

// [OUTBOUND]
type AvatarRepository interface {
	CreateAvatar(ctx context.Context, avatar *domain.Avatar) (*domain.Avatar, error)
	GetAvatarByID(ctx context.Context, id int) (*domain.Avatar, error)
	UpdateAvatar(ctx context.Context, avatar *domain.Avatar) (*domain.Avatar, error)
	DeleteAvatar(ctx context.Context, id int) error
}

type AvatarService interface {
}

package service

import (
	"context"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/core/port"
)

type AvatarService struct {
	avatarRepository port.AvatarRepository
}

func NewAvatarService(avatarRepository port.AvatarRepository) *AvatarService {
	userService := AvatarService{
		avatarRepository: avatarRepository,
	}

	return &userService
}

func (s *AvatarService) CreateAvatar(ctx context.Context, avatar *domain.Avatar) (*domain.Avatar, error) {
	return s.avatarRepository.CreateAvatar(ctx, avatar)
}

func (s *AvatarService) GetAvatar(ctx context.Context, id int) (*domain.Avatar, error) {
	return s.avatarRepository.GetAvatar(ctx, id)
}

func (s *AvatarService) ListAvatar(ctx context.Context, skip, limit int, order, criterion string) ([]*domain.Avatar, error) {
	return s.avatarRepository.ListAvatar(ctx, skip, limit, order, criterion)
}

func (s *AvatarService) UpdateAvatar(ctx context.Context, avatar *domain.Avatar) (*domain.Avatar, error) {
	return s.avatarRepository.UpdateAvatar(ctx, avatar)
}

func (s *AvatarService) DeleteAvatar(ctx context.Context, id int) error {
	return s.avatarRepository.DeleteAvatar(ctx, id)
}

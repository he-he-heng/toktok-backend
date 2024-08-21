package port

import (
	"context"
	"toktok-backend/internal/core/domain"
)

type MessageRepository interface {
	CreateMessage(ctx context.Context, message *domain.Message) (*domain.Message, error)
	GetMessage(ctx context.Context, id int) (*domain.Message, error)
	ListMessaeeByRealtionID(ctx context.Context, relationId, skip, limit int, order, critertion string)
	// UpdateMessage(ctx context.Context, message *domain.Message) (*domain.Message, error)
	// DeleteMessage(ctx context.Context, id int) error
}

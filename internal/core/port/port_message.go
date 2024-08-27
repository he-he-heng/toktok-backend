package port

import (
	"context"
	"toktok-backend/internal/core/domain"
)

type MessageRepository interface {
	CreateMessage(ctx context.Context, message *domain.Message) (*domain.Message, error)
	GetMessage(ctx context.Context, id int) (*domain.Message, error)
	ListMessaeeByRoomID(ctx context.Context, roomID, skip, limit int, order, critertion string)
	UpdateMessage(ctx context.Context, message *domain.Message) (*domain.Message, error)
	DeleteMessage(ctx context.Context, id int) error
}

type MessageService interface {
	ListMessage(ctx context.Context, roomID, skip, limit int, order, criterion string)
}

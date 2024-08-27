package port

import (
	"context"
	"toktok-backend/internal/core/domain"
)

type RoomRepository interface {
	CreateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error)
	GetRoom(ctx context.Context, id int) (*domain.Room, error)
	ListRoom(ctx context.Context, skip, limit int, order, critertion string) ([]*domain.Room, error)
	// UpdateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error)
	DeleteRoom(ctx context.Context, id int) error

	// using by relation service
	GetRoomByRelationID(ctx context.Context, relationID int) (*domain.Room, error)
}

type RoomService interface {
	CreateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error)
	GetRoom(ctx context.Context, id int) (*domain.Room, error)
	ListRoom(ctx context.Context, skip, limit int, order, critertion string) ([]*domain.Room, error)
	DeleteRoom(ctx context.Context, id int) error
}

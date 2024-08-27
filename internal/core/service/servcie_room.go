package service

import (
	"context"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/core/port"
)

type RoomService struct {
	roomRepository port.RoomRepository
}

func NewRoomService(roomRepository port.RoomRepository) *RoomService {
	roomService := RoomService{
		roomRepository: roomRepository,
	}

	return &roomService
}

func (s *RoomService) CreateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error) {
	return s.roomRepository.CreateRoom(ctx, room)
}

func (s *RoomService) GetRoom(ctx context.Context, id int) (*domain.Room, error) {
	return s.roomRepository.GetRoom(ctx, id)
}

func (s *RoomService) ListRoom(ctx context.Context, skip, limit int, order, critertion string) ([]*domain.Room, error) {
	return s.roomRepository.ListRoom(ctx, skip, limit, order, critertion)
}

func (s *RoomService) DeleteRoom(ctx context.Context, id int) error {
	return s.roomRepository.DeleteRoom(ctx, id)
}

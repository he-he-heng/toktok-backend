package repository

import "toktok-backend/internal/adapter/persistence/mysql"

type RoomRepository struct {
	client *mysql.Client
}

func NewRoomRepoistory(client *mysql.Client) *RoomRepository {
	roomRepository := RoomRepository{
		client: client,
	}

	return &roomRepository
}

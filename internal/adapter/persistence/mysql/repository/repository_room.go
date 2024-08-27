package repository

import (
	"context"
	"toktok-backend/internal/adapter/persistence/mysql"
	entrelation "toktok-backend/internal/adapter/persistence/mysql/ent/relation"
	entroom "toktok-backend/internal/adapter/persistence/mysql/ent/room"
	"toktok-backend/internal/adapter/persistence/mysql/utils"
	"toktok-backend/internal/core/domain"

	"entgo.io/ent/dialect/sql"
)

type RoomRepository struct {
	client *mysql.Client
}

func NewRoomRepoistory(client *mysql.Client) *RoomRepository {
	roomRepository := RoomRepository{
		client: client,
	}

	return &roomRepository
}

func (r *RoomRepository) CreateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error) {
	createdRoom, err := r.client.Room.Create().
		SetAvatarID(room.AvatarRelationID).
		SetFriendID(room.FriendRelationID).
		Save(ctx)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	loadedRoom, err := r.client.Room.Query().
		Where(entroom.ID(createdRoom.ID)).
		WithAvatar().
		WithFriend().
		Only(ctx)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	return utils.ToDomainRoom(loadedRoom), nil
}

func (r *RoomRepository) GetRoom(ctx context.Context, id int) (*domain.Room, error) {

	queriedRoom, err := r.client.Room.Query().
		Where(entroom.ID(id)).
		WithAvatar().
		WithFriend().
		Only(ctx)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	return utils.ToDomainRoom(queriedRoom), nil
}

func (r *RoomRepository) GetRoomByRelationID(ctx context.Context, avatarRelationID int) (*domain.Room, error) {
	queriedRoom, err := r.client.Room.Query().
		Where(entroom.HasAvatarWith(
			entrelation.ID(avatarRelationID),
		)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ToDomainRoom(queriedRoom), nil
}

func (r *RoomRepository) ListRoom(ctx context.Context, skip, limit int, order, critertion string) ([]*domain.Room, error) {
	builder := r.client.Room.Query().
		Limit(limit).
		Offset((skip - 1) * limit)

	orderTermOption := sql.OrderAsc()
	if order == "desc" {
		orderTermOption = sql.OrderDesc()
	}

	switch critertion {
	case "uid":

	default:
		builder.Order(
			entroom.ByID(
				orderTermOption,
			),
		)
	}

	queriedUsers, err := builder.All(ctx)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	return utils.ToDomainRooms(queriedUsers), nil
}

// func (r *RoomRepository) UpdateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error) {

// }

func (r *RoomRepository) DeleteRoom(ctx context.Context, id int) error {
	err := r.client.Room.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return utils.ErrWrap(err)
	}

	return nil
}

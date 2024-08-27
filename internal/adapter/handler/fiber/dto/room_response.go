package dto

import "toktok-backend/internal/core/domain"

type GetRoomResponse struct {
	ID               int `json:"id"`
	AvatarRelationID int `json:"avatarRelationID"`
	FriendRelationID int `json:"friendRelationID"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"deletedAt"`
}

func (GetRoomResponse) Of(room *domain.Room) GetRoomResponse {
	ret := GetRoomResponse{
		ID:               room.ID,
		AvatarRelationID: room.AvatarRelationID,
		FriendRelationID: room.FriendRelationID,
	}

	ret.CreatedAt = room.CreatedAt.Format("2006-01-02T15:04:05Z")
	ret.UpdatedAt = room.UpdatedAt.Format("2006-01-02T15:04:05Z")

	if !room.DeletedAt.IsZero() {
		ret.DeletedAt = room.DeletedAt.Format("2006-01-02T15:04:05Z")
	}

	return ret
}

package dto

import (
	"toktok-backend/internal/core/domain"
)

type relation struct {
	ID            int    `json:"id"`
	AvatarID      int    `json:"avatarId"`
	FriendID      int    `json:"friendId"`
	State         string `json:"state"`
	AlertState    string `json:"alertState,omitempty"`
	HasNewMessage bool   `json:"hasNewMessage,omitempty"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updaetdAt"`
	DeletedAt string `json:"deletedAt,omitempty"`
}

func (relation) of(r *domain.Relation, hasNewMessages ...bool) relation {
	retR := relation{
		ID:         r.ID,
		AvatarID:   r.AvatarID,
		FriendID:   r.FriendID,
		State:      string(r.State),
		AlertState: string(r.AlertState),
		// set hasNewMessage
	}

	if len(hasNewMessages) > 0 {
		retR.HasNewMessage = hasNewMessages[0]
	}

	retR.CreatedAt = r.CreatedAt.Format("2006-01-02T15:04:05Z")
	retR.UpdatedAt = r.UpdatedAt.Format("2006-01-02T15:04:05Z")

	if !r.DeletedAt.IsZero() {
		retR.DeletedAt = r.DeletedAt.Format("2006-01-02T15:04:05Z")
	}

	return retR
}

type CreateRelationResponse struct {
	Relations [2]relation `json:"relations,omitempty"`
}

func (CreateRelationResponse) Of(args *[2]domain.Relation) CreateRelationResponse {
	ret := CreateRelationResponse{
		Relations: [2]relation{},
	}

	for i := 0; i < 2; i++ {
		ret.Relations[i] = relation{}.of(&args[i])
	}

	return ret
}

type GetRelationResponse struct {
}

func (GetRelationResponse) Of(r *domain.Relation) relation {
	return relation{}.of(r)
}

type ListRelationResponse struct {
	Relations []relation `json:"relations,omitempty"`
}

func (ListRelationResponse) Of(rs []*domain.Relation) ListRelationResponse {
	listRelation := ListRelationResponse{
		Relations: make([]relation, 0),
	}

	for _, r := range rs {
		listRelation.Relations = append(listRelation.Relations, relation{}.of(r))
	}

	return listRelation
}

type UpdateRelationResponse struct {
	Relations [2]relation `json:"relations,omitempty"`
}

func (UpdateRelationResponse) Of(rs *[2]domain.Relation) UpdateRelationResponse {
	returnValue := UpdateRelationResponse{
		Relations: [2]relation{},
	}

	for i, r := range rs {
		returnValue.Relations[i] = relation{}.of(&r)
	}

	return returnValue
}

type GetRoomByRelationByIDResponse struct {
	ID               int `json:"id"`
	AvatarRelationID int `json:"avatarRelationID"`
	FriendRelationID int `json:"friendRelationID"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"deletedAt"`
}

func (GetRoomByRelationByIDResponse) Of(room *domain.Room) GetRoomByRelationByIDResponse {
	ret := GetRoomByRelationByIDResponse{
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

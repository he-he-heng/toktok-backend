package dto

import (
	"toktok-backend/internal/core/domain"
)

type CreateRelationRequest struct {
	AvatarID     int `json:"avatarId" validate:""`
	WantFriendID int `json:"wantFriendId" validate:""`
}

func (c CreateRelationRequest) ToDomainRelation() *domain.Relation {
	r := domain.Relation{
		AvatarID: c.AvatarID,
		FriendID: c.WantFriendID,
	}

	return &r
}

type UpdateRelationRequest struct {
	State    string `json:"state" validate:"omitempty,oneof=request-friend pending decline friend remove"`
	AvatarID int    `json:"avatarId" validate:""`
}

func (dto UpdateRelationRequest) ToDomainRelation(id int) *domain.Relation {
	r := domain.Relation{
		ID:       id,
		AvatarID: dto.AvatarID,
		State:    domain.RelationStateType(dto.State),
	}

	return &r
}

type DeleteRelationRequest struct {
	AvatarID int `json:"avatarId" validate:""`
}

func (dto DeleteRelationRequest) ToDomainRelation() *domain.Relation {
	r := domain.Relation{
		AvatarID: dto.AvatarID,
	}

	return &r
}

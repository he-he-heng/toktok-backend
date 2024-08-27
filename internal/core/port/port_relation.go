package port

import (
	"context"
	"toktok-backend/internal/core/domain"
)

type RelationRepository interface {
	CreateRelation(ctx context.Context, relation *domain.Relation) (*domain.Relation, error)

	GetRelation(ctx context.Context, id int) (*domain.Relation, error)
	GetRelationByAvatarIDAndRelationIDWithState(ctx context.Context, argRelation *domain.Relation) (*domain.Relation, error)

	ListRelationByAvatarIDAndFriendID(ctx context.Context, avatarID int, friendID int) ([]*domain.Relation, error)
	ListRelation(ctx context.Context, skip, limit int, order, criterion string, filter domain.RelationStateType) ([]*domain.Relation, error)
	ListRelationByAvatarID(ctx context.Context, skip, limit int, order, criterion string, filter domain.RelationStateType, avatarID int) ([]*domain.Relation, error)

	UpdateRelation(ctx context.Context, relation *domain.Relation) (*domain.Relation, error)
	DeleteRelation(ctx context.Context, id int) error
}

type RelationService interface {
	CreateRelation(ctx context.Context, relation *domain.Relation) (*[2]domain.Relation, error)
	GetRelation(ctx context.Context, id int) (*domain.Relation, error)
	ListRelation(ctx context.Context, skip, limit int, order, criterion string, filter domain.RelationStateType) ([]*domain.Relation, error)
	ListRelationByAvatarID(ctx context.Context, skip, limit int, order, criterion string, filter domain.RelationStateType, avatarID int) ([]*domain.Relation, error)
	UpdateRelation(ctx context.Context, relation *domain.Relation) (*[2]domain.Relation, error)
	DeleteRelation(ctx context.Context, id int) error

	// using room repository
	GetRoomByRelationID(ctx context.Context, relationID int) (*domain.Room, error)
}

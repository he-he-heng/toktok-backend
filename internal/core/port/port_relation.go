package port

import (
	"context"
	"toktok-backend/internal/core/domain"
)

type RelationRepository interface {
	CreateRelation(ctx context.Context, relation *domain.Relation) (*domain.Relation, error)
	GetRelation(ctx context.Context, id int) (*domain.Relation, error)
	ListRelation(ctx context.Context, skip, limit int, order, criterion, filter string) ([]*domain.Relation, error)
	UpdateRelation(ctx context.Context, relation *domain.Relation) (*domain.Relation, error)
	DeleteRelation(ctx context.Context, id int) error
}

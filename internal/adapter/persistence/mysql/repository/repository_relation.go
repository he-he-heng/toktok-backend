package repository

import (
	"context"
	"slices"
	"toktok-backend/internal/adapter/persistence/mysql"
	entavatar "toktok-backend/internal/adapter/persistence/mysql/ent/avatar"
	entrelation "toktok-backend/internal/adapter/persistence/mysql/ent/relation"

	"toktok-backend/internal/adapter/persistence/mysql/utils"
	"toktok-backend/internal/core/domain"

	"entgo.io/ent/dialect/sql"
)

type RelationRepository struct {
	client *mysql.Client
}

func NewRelationRepository(client *mysql.Client) *RelationRepository {
	relationRepository := RelationRepository{
		client: client,
	}

	return &relationRepository
}

func (r *RelationRepository) CreateRelation(ctx context.Context, relation *domain.Relation) (*domain.Relation, error) {
	createdRelation, err := r.client.Relation.Create().
		SetAvatarID(relation.AvatarID).
		SetFriendID(relation.FriendID).
		SetState(entrelation.State(relation.State)).
		SetAlertState(entrelation.AlertState(relation.State)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ToDomainRelation(createdRelation), nil
}

func (r *RelationRepository) GetRelation(ctx context.Context, id int) (*domain.Relation, error) {
	queriedRelation, err := r.client.Relation.Query().
		Where(entrelation.IDEQ(id)).
		WithAvatar().
		WithFriend().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ToDomainRelation(queriedRelation), nil
}

func (r *RelationRepository) ListRelation(ctx context.Context, skip, limit int, order, criterion string, filter domain.RelationStateType) ([]*domain.Relation, error) {
	builder := r.client.Relation.Query().
		Limit(limit).
		Offset((skip - 1) * limit).
		WithAvatar().
		WithFriend()

	validStates := []domain.RelationStateType{
		domain.RelationStateFriend,
		domain.RealtionStatePending,
		domain.RelationStateRequestFriend,
		domain.RelationStateDeclined,
		domain.RelationStateRemoved,
	}
	if slices.Contains(validStates, filter) {
		builder.Where(entrelation.StateEQ(entrelation.State(filter)))
	}

	orderTermOption := sql.OrderAsc()
	if order == "desc" {
		orderTermOption = sql.OrderDesc()
	}

	switch criterion {
	case "avatar_id":
		builder.Order(
			entrelation.ByAvatarField(
				entavatar.FieldID,
				orderTermOption,
			),
		)

	case "friend_id":
		builder.Order(
			entrelation.ByFriendField(
				entavatar.FieldID,
				orderTermOption,
			),
		)

	case "id":
		builder.Order(
			entrelation.ByID(
				orderTermOption,
			),
		)
	}

	queriedRelations, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ToDomainRelations(queriedRelations), nil
}

func (r *RelationRepository) UpdateRelation(ctx context.Context, relation *domain.Relation) (*domain.Relation, error) {
	builder := r.client.Relation.UpdateOneID(relation.ID)

	if utils.Changeable(relation.State) {
		builder.SetState(entrelation.State(relation.State))
	}

	if utils.Changeable(relation.AlertState) {
		builder.SetAlertState(entrelation.AlertState(relation.AlertState))
	}

	updatedRelation, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ToDomainRelation(updatedRelation), nil
}

func (r *RelationRepository) DeleteRelation(ctx context.Context, id int) error {
	err := r.client.Client.Relation.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

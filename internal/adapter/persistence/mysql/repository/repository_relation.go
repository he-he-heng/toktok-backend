package repository

import (
	"context"
	"slices"
	"toktok-backend/internal/adapter/persistence/mysql"
	entavatar "toktok-backend/internal/adapter/persistence/mysql/ent/avatar"
	"toktok-backend/internal/adapter/persistence/mysql/ent/relation"
	entrelation "toktok-backend/internal/adapter/persistence/mysql/ent/relation"
	"toktok-backend/internal/adapter/persistence/mysql/utils"
	"toktok-backend/internal/core/domain"
	"toktok-backend/pkg/errors"

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

	builder := r.client.Relation.Create().
		SetAvatarID(relation.AvatarID).
		SetFriendID(relation.FriendID)

	if relation.State != "" {
		builder.SetState(entrelation.State(relation.State))
	}

	if relation.AlertState != "" {
		builder.SetAlertState(entrelation.AlertState(relation.AlertState))
	}

	createdRelation, err := builder.Save(ctx)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	loadedRelation, err := r.client.Relation.Query().
		Where(entrelation.IDEQ(createdRelation.ID)).
		WithAvatar().
		WithFriend().
		Only(ctx)

	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	return utils.ToDomainRelation(loadedRelation), nil

}

func (r *RelationRepository) GetRelation(ctx context.Context, id int) (*domain.Relation, error) {
	queriedRelation, err := r.client.Relation.Query().
		Where(entrelation.IDEQ(id)).
		WithAvatar().
		WithFriend().
		Only(ctx)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	return utils.ToDomainRelation(queriedRelation), nil

}

func (r *RelationRepository) GetRelationByAvatarIDAndRelationIDWithState(ctx context.Context, argRelation *domain.Relation) (*domain.Relation, error) {
	// 오직 이 함수는 state의 값이 domain.RelationStatePending과 RelationReqeustFriend만 가능합니다.
	//fmt.Printf("argRelation: %+v\n", argRelation)
	//fmt.Printf("bool %v\n", !(argRelation.State == domain.RelationStatePending || argRelation.State == domain.RelationStateRequestFriend))

	if !(argRelation.State == domain.RelationStatePending || argRelation.State == domain.RelationStateFriend) {
		return nil, errors.Wrap(domain.ErrBadParam, "invalid filter")
	}

	queriedRelation, err := r.client.Relation.Query().Where(
		relation.And(
			relation.HasAvatarWith(entavatar.ID(argRelation.AvatarID)),
			relation.HasFriendWith(entavatar.ID(argRelation.FriendID)),
			relation.StateEQ(entrelation.State(argRelation.State)),
		),
	).WithAvatar().WithFriend().Only(ctx)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	return utils.ToDomainRelation(queriedRelation), nil
}

func (r *RelationRepository) ListRelationByAvatarIDAndFriendID(ctx context.Context, avatarID int, friendID int) ([]*domain.Relation, error) {
	builder := r.client.Relation.Query().
		Where(
			relation.And(
				relation.HasAvatarWith(entavatar.ID(avatarID)),
				relation.HasFriendWith(entavatar.ID(friendID)),
			),
		).
		WithAvatar().
		WithFriend()

	rels, err := builder.All(ctx)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	return utils.ToDomainRelations(rels), nil
}

func (r *RelationRepository) ListRelation(ctx context.Context, skip, limit int, order, criterion string, filter domain.RelationStateType) ([]*domain.Relation, error) {
	builder := r.client.Relation.Query().
		Limit(limit).
		Offset((skip - 1) * limit).
		WithAvatar().
		WithFriend()

	validStates := []domain.RelationStateType{
		domain.RelationStateFriend,
		domain.RelationStatePending,
		domain.RelationStateRequestFriend,
		domain.RelationStateDecline,
		domain.RelationStateRemove,
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
		return nil, utils.ErrWrap(err)
	}

	return utils.ToDomainRelations(queriedRelations), nil
}

func (r *RelationRepository) ListRelationByAvatarID(ctx context.Context, skip, limit int, order, criterion string, filter domain.RelationStateType, avatarID int) ([]*domain.Relation, error) {
	builder := r.client.Relation.Query().
		Limit(limit).
		Offset((skip - 1) * limit).
		Where(entrelation.HasFriendWith(entavatar.ID(avatarID))).
		WithAvatar().
		WithFriend()

	validStates := []domain.RelationStateType{
		domain.RelationStateFriend,
		domain.RelationStatePending,
		domain.RelationStateRequestFriend,
		domain.RelationStateDecline,
		domain.RelationStateRemove,
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
		return nil, utils.ErrWrap(err)
	}

	return utils.ToDomainRelations(queriedRelations), nil
}

func (r *RelationRepository) UpdateRelation(ctx context.Context, relation *domain.Relation) (*domain.Relation, error) {
	builder := r.client.Relation.UpdateOneID(relation.ID)

	if relation.State != "" {
		builder.SetState(entrelation.State(relation.State))
	}

	if relation.AlertState != "" {
		builder.SetAlertState(entrelation.AlertState(relation.AlertState))
	}

	updatedRelation, err := builder.Save(ctx)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	loadedRelation, err := r.client.Relation.Query().
		Where(entrelation.ID(updatedRelation.ID)).
		WithAvatar().
		WithFriend().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ToDomainRelation(loadedRelation), nil
}

func (r *RelationRepository) DeleteRelation(ctx context.Context, id int) error {
	err := r.client.Client.Relation.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return utils.ErrWrap(err)
	}

	return nil
}

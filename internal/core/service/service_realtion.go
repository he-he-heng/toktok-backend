package service

import (
	"context"
	"fmt"
	"toktok-backend/internal/adapter/persistence/mysql/utils"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/core/port"
	"toktok-backend/pkg/errors"
)

type RelationService struct {
	relationRepository port.RelationRepository
	roomRepository     port.RoomRepository
}

func NewRelationService(relationRepository port.RelationRepository, roomRepository port.RoomRepository) *RelationService {
	relationService := RelationService{
		relationRepository: relationRepository,
		roomRepository:     roomRepository,
	}

	return &relationService
}

func (s *RelationService) CreateRelation(ctx context.Context, relation *domain.Relation) (*[2]domain.Relation, error) {

	// 친구를 추가할 수 없는 경우
	/*
		A, B가 있을 때


		A(Friend) < --- > B(Friend) 이면 친추를 할 수 없음

		A(Pending) < --- > B(Request-Fried) 이면 친추를 할 수 없음

		A(Request-Friend) < --- > B(Pending) 이면 친추를 할 수 없음

		즉, 서로 친구이거나, 친구요청 후 응답대기 상태라면 친추가 불가능하다.
	*/

	queriedAvatarRelations, err := s.relationRepository.ListRelationByAvatarIDAndFriendID(ctx, relation.AvatarID, relation.FriendID)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	queriedFriendRelations, err := s.relationRepository.ListRelationByAvatarIDAndFriendID(ctx, relation.FriendID, relation.AvatarID)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	// avatar id와 friend id를 넣었을 때, 이 필드 두 개와 일치하는 어트리뷰트를 찾고, 그 어트리뷰트의 state필드가 reuqest-friend 이거나, pending 이라면 그 어트리뷰티를 반환하는 함수명

	for i := 0; i < len(queriedAvatarRelations); i++ {
		avatarState := queriedAvatarRelations[i].State
		friendState := queriedFriendRelations[i].State

		// avatar의 state가 firned, pending, reqeust-friend라면
		if avatarState == domain.RelationStateFriend || avatarState == domain.RelationStateRequestFriend || avatarState == domain.RelationStatePending {
			return nil, errors.Wrap(domain.ErrConstraint, "avatarState is not (domainRelationStateFriend | domainStatePending)")
		}

		// friend의 state가 firned, pending, reqeust-friend라면
		if friendState == domain.RelationStateFriend || friendState == domain.RelationStateRequestFriend || friendState == domain.RelationStatePending {
			return nil, errors.Wrap(domain.ErrConstraint, "friendState is not (domainRelationStateFriend | domainStatePending)")
		}
	}

	sender := domain.Relation{
		AvatarID:   relation.AvatarID,
		FriendID:   relation.FriendID,
		State:      domain.RelationStateRequestFriend,
		AlertState: domain.RelationAlertStateAllow,
	}

	receiver := domain.Relation{
		AvatarID:   relation.FriendID,
		FriendID:   relation.AvatarID,
		State:      domain.RelationStatePending,
		AlertState: domain.RelationAlertStateAllow,
	}

	createAvatarRelation, err := s.relationRepository.CreateRelation(ctx, &sender)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	createFriendRelation, err := s.relationRepository.CreateRelation(ctx, &receiver)
	if err != nil {
		return nil, utils.ErrWrap(err)
	}

	return &[2]domain.Relation{*createAvatarRelation, *createFriendRelation}, nil
}

func (s *RelationService) GetRelation(ctx context.Context, id int) (*domain.Relation, error) {
	return s.relationRepository.GetRelation(ctx, id)
}

func (s *RelationService) ListRelation(ctx context.Context, skip, limit int, order, criterion string, filter domain.RelationStateType) ([]*domain.Relation, error) {
	return s.relationRepository.ListRelation(ctx, skip, limit, order, criterion, domain.RelationStateType(filter))
}

func (s *RelationService) ListRelationByAvatarID(ctx context.Context, skip, limit int, order, criterion string, filter domain.RelationStateType, avatarID int) ([]*domain.Relation, error) {
	return s.relationRepository.ListRelationByAvatarID(ctx, skip, limit, order, criterion, filter, avatarID)
}

func (s *RelationService) UpdateRelation(ctx context.Context, relation *domain.Relation) (*[2]domain.Relation, error) {
	// Input : ID, State(only decline, friend, remove)

	isFriend := false

	gotRelation, err := s.relationRepository.GetRelation(ctx, relation.ID)
	if err != nil {
		return nil, err
	}

	/*
		변경할 수 있는 조건
		A(friend) - B(friend) 일 때  오직 remove state 로만 변경할 수 있다.

		A(friend-request) - B(pending) 이라면 오직 decline, friend 관계으로만 변경할 수 있다.
	*/

	updateSenderRelation := &domain.Relation{}
	updateReceiverRelation := &domain.Relation{}

	if relation.State == domain.RelationStateDecline && gotRelation.State == domain.RelationStatePending {

		friendRelation, err := s.relationRepository.GetRelationByAvatarIDAndRelationIDWithState(ctx, &domain.Relation{
			AvatarID: gotRelation.FriendID,
			FriendID: gotRelation.AvatarID,
			State:    domain.RelationStateRequestFriend,
		})
		if err != nil {
			return nil, err
		}

		updateSenderRelation = &domain.Relation{
			ID:    gotRelation.ID,
			State: domain.RelationStateDecline,
		}

		updateReceiverRelation = &domain.Relation{
			ID:    friendRelation.ID,
			State: domain.RelationStateDecline,
		}

	} else if relation.State == domain.RelationStateFriend && gotRelation.State == domain.RelationStatePending {
		friendRelation, err := s.relationRepository.GetRelationByAvatarIDAndRelationIDWithState(ctx, &domain.Relation{
			AvatarID: gotRelation.FriendID,
			FriendID: gotRelation.AvatarID,
			State:    domain.RelationStateRequestFriend,
		})
		if err != nil {
			fmt.Println("reutnr")
			return nil, err
		}

		updateSenderRelation = &domain.Relation{
			ID:    gotRelation.ID,
			State: domain.RelationStateFriend,
		}

		updateReceiverRelation = &domain.Relation{
			ID:    friendRelation.ID,
			State: domain.RelationStateFriend,
		}

		isFriend = true

	} else if relation.State == domain.RelationStateRemove && gotRelation.State == domain.RelationStateFriend {
		friendRelation, err := s.relationRepository.GetRelationByAvatarIDAndRelationIDWithState(ctx, &domain.Relation{
			AvatarID: gotRelation.FriendID,
			FriendID: gotRelation.AvatarID,
			State:    domain.RelationStateFriend,
		})
		if err != nil {
			return nil, err
		}

		updateSenderRelation = &domain.Relation{
			ID:    gotRelation.ID,
			State: domain.RelationStateRemove,
		}

		updateReceiverRelation = &domain.Relation{
			ID:    friendRelation.ID,
			State: domain.RelationStateRemove,
		}
	} else {
		return nil, errors.Wrap(domain.ErrConstraint, errors.New("you do not have the right to update this request"))
	}

	updatedSenderRelation, err := s.relationRepository.UpdateRelation(ctx, updateSenderRelation)
	if err != nil {
		return nil, err
	}
	updatedReceiverRelation, err := s.relationRepository.UpdateRelation(ctx, updateReceiverRelation)
	if err != nil {
		return nil, err
	}

	if isFriend {
		_, err := s.roomRepository.CreateRoom(ctx, &domain.Room{
			AvatarRelationID: updateSenderRelation.ID,
			FriendRelationID: updateReceiverRelation.ID,
		})

		if err != nil {
			return nil, err
		}
	}

	return &[2]domain.Relation{
		*updatedSenderRelation,
		*updatedReceiverRelation,
	}, nil
}

func (s *RelationService) DeleteRelation(ctx context.Context, id int) error {
	return s.relationRepository.DeleteRelation(ctx, id)
}

/////

func (s *RelationService) GetRoomByRelationID(ctx context.Context, relationID int) (*domain.Room, error) {
	room, err := s.roomRepository.GetRoomByRelationID(ctx, relationID)
	if err != nil {
		return nil, err
	}

	return room, nil
}

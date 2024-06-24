package service

import (
	"context"
	"errors"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/core/port"
	"toktok-backend/internal/core/util"
)

type UserService struct {
	repository port.UserRepository
}

func NewUserService(userRepository port.UserRepository) *UserService {
	return &UserService{
		repository: userRepository,
	}
}

// Register registers a new user
func (s *UserService) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, domain.ErrInternal
	}

	user.Password = hashedPassword

	user, err = s.repository.CreateUser(ctx, user)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return user, nil
}

// GetUser returns a user by id
func (s *UserService) GetUser(ctx context.Context, id int) (*domain.User, error) {
	user, err := s.repository.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(domain.ErrDataNotFound, err) {
			return nil, err
		}

		return nil, domain.ErrInternal
	}

	return user, nil
}

// ListUsers returns a list of users with pagination
func (s *UserService) ListUsers(ctx context.Context, skip, limit int) ([]domain.User, error) {
	users, err := s.repository.ListUsers(ctx, skip, limit)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return users, err
}

// UpdateUser updates a user
func (s *UserService) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	existUser, err := s.repository.GetUserByID(ctx, user.ID)
	if err != nil {
		if errors.Is(domain.ErrDataNotFound, err) {
			return nil, err
		}

		return nil, domain.ErrInternal
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, domain.ErrInternal
	}

	existUser.Password = hashedPassword

	user, err = s.repository.UpdateUser(ctx, existUser)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return user, nil
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	_, err := s.repository.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(domain.ErrDataNotFound, err) {
			return err
		}

		return domain.ErrInternal
	}

	return s.repository.DeleteUser(ctx, id)
}

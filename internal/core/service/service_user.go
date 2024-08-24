package service

import (
	"context"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/core/port"
	"toktok-backend/internal/core/service/utils"
	"toktok-backend/pkg/errors"
)

type UserService struct {
	UserRepository port.UserRepository
}

func NewUserService(userRepository port.UserRepository) *UserService {
	userService := UserService{
		UserRepository: userRepository,
	}

	return &userService
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {

	hashedPssword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, errors.Wrap(domain.ErrInternalServerError, hashedPssword)
	}

	return s.UserRepository.CreateUser(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, id int) (*domain.User, error) {
	return s.UserRepository.GetUser(ctx, id)
}

func (s *UserService) ListUser(ctx context.Context, skip, limit int, order, criterion string) ([]*domain.User, error) {
	return s.UserRepository.ListUser(ctx, skip, limit, order, criterion)
}

func (s *UserService) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	if user.Password != "" {
		hashedPssword, err := utils.HashPassword(user.Password)
		if err != nil {
			return nil, errors.Wrap(domain.ErrInternalServerError, err)
		}

		user.Password = hashedPssword
	}

	return s.UserRepository.UpdateUser(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.UserRepository.DeleteUser(ctx, id)
}

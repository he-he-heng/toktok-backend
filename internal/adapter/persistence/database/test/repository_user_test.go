package test

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"toktok-backend/internal/adapter/persistence/database/repository"
	"toktok-backend/internal/adapter/persistence/database/test/provider"
	"toktok-backend/internal/core/domain"
)

var userRepo *repository.UserRepository

func init() {
	userRepo = repository.NewUserRepository(provider.Get().Client)
}

func TestCreateUser(t *testing.T) {
	defer provider.Get().TerminateContainer(context.Background())

	testCases := []struct {
		user   *domain.User
		expect *domain.User
	}{
		{
			user: &domain.User{
				UID:      "socketdokkdj3",
				Password: "fjfji3lski3",
			},

			expect: &domain.User{
				ID:       1,
				UID:      "socketdokkdj3",
				Password: "fjfji3lski3",
				Email:    nil,
				Role:     domain.UserRoleUser,
				BanState: domain.UserBanStateUnBan,
			},
		},
	}

	for i, tc := range testCases {
		name := fmt.Sprintf("#%d CreateUser", i)
		t.Run(name, func(tsub *testing.T) {
			got, err := userRepo.CreateUser(context.Background(), tc.user)
			if err != nil {
				tsub.Errorf("failed to create user: %v", err)
			}

			tc.expect.CreatedAt = got.CreatedAt
			tc.expect.UpdatedAt = got.UpdatedAt
			tc.expect.DeletedAt = got.DeletedAt

			if !reflect.DeepEqual(tc.expect, got) {
				tsub.Errorf("tc.expect and got don't have each other. got=%+v\nexpect=%+v\n", got, tc.expect)
			}
		})
	}
}

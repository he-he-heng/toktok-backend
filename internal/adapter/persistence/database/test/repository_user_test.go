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
	defer provider.Get().TerminateContainer(provider.MyContext)

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

		{
			user: &domain.User{
				UID:      "himynameisback",
				Password: "wangood",
				Email:    strPtr("hijak"),
				Role:     domain.UserRoleAdmin,
				BanState: domain.UserBanStateUnBan,
			},

			expect: &domain.User{
				ID:       2,
				UID:      "himynameisback",
				Password: "wangood",
				// Email:    strPtr("hijak"),
				Role:     domain.UserRoleAdmin,
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

			tc.expect.Email = got.Email
			tc.expect.CreatedAt = got.CreatedAt
			tc.expect.UpdatedAt = got.UpdatedAt
			tc.expect.DeletedAt = got.DeletedAt

			if !reflect.DeepEqual(tc.expect, got) {
				tsub.Errorf("\ntc.expect and got don't have each other.\ngot=%+v\nexpect=%+v\n", got, tc.expect)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	defer provider.Get().TerminateContainer(provider.MyContext)

	ctx := context.Background()
	// database seedling
	userRepo.CreateUser(ctx, &domain.User{
		UID:      "socketdokkdj3",
		Password: "fjfji3lski3",
	})

	testCases := []struct {
		selectId int
		expect   *domain.User
	}{
		{
			selectId: 1,
			expect: &domain.User{
				ID:       1,
				UID:      "socketdokkdj3",
				Password: "fjfji3lski3",
				Role:     domain.UserRoleUser,
				BanState: domain.UserBanStateUnBan,
			},
		},
	}

	for i, tc := range testCases {
		name := fmt.Sprintf("#%d GetUser", i)
		t.Run(name, func(tsub *testing.T) {
			got, err := userRepo.GetUser(ctx, tc.selectId)
			if err != nil {
				tsub.Errorf("failed to get user: %v", err)
			}

			tc.expect.Email = got.Email
			tc.expect.CreatedAt = got.CreatedAt
			tc.expect.UpdatedAt = got.UpdatedAt
			tc.expect.DeletedAt = got.DeletedAt

			if !reflect.DeepEqual(tc.expect, got) {
				tsub.Errorf("\ntc.expect and got don't have each other.\ngot=%+v\nexpect=%+v\n", got, tc.expect)
			}
		})
	}

}

func TestListUser(t *testing.T) {
	defer provider.Get().TerminateContainer(provider.MyContext)

	ctx := context.Background()
	// database seedling
	err := func() error {
		// field 1
		if _, err := userRepo.CreateUser(ctx, &domain.User{
			UID:      "atesttest",
			Password: "43299d9s",
		}); err != nil {
			return err
		}

		// field 2
		if _, err := userRepo.CreateUser(ctx, &domain.User{
			UID:      "btesttest",
			Password: "sldkkj39990",
		}); err != nil {
			return err
		}

		// field 3
		if _, err := userRepo.CreateUser(ctx, &domain.User{
			UID:      "ctesttest",
			Password: "d0993jd",
			Role:     domain.UserRoleAdmin,
		}); err != nil {
			return err
		}

		// field 4
		if _, err := userRepo.CreateUser(ctx, &domain.User{
			UID:      "dtestest",
			Password: "109djj223",
			BanState: domain.UserBanStateBan,
		}); err != nil {
			return err
		}

		return nil

	}()
	if err != nil {
		t.Errorf("failed to create user: %v", err)
	}

	testCases := []struct {
		skip, limit       int
		order, critertion string
		expect            []*domain.User
	}{
		// case 1 , 페이징 기법과 기본값들이 잘 동작하는 확인
		{
			skip:  2,
			limit: 2,
			// order: "asc", //-->  default
			// critertion: "id" --> default
			expect: []*domain.User{
				&domain.User{
					ID:       3,
					UID:      "ctesttest",
					Password: "d0993jd",
					Email:    nil,
					Role:     domain.UserRoleAdmin,
					BanState: domain.UserBanStateUnBan,
				},

				&domain.User{
					ID:       4,
					UID:      "dtestest",
					Password: "109djj223",
					Email:    nil,
					Role:     domain.UserRoleUser,
					BanState: domain.UserBanStateBan,
				},
			},
		},

		// 기대값 Filed3, 모두 옵션 주고것 까지
		{
			order:      "desc",
			critertion: "uid",

			skip:  2,
			limit: 1,

			expect: []*domain.User{
				&domain.User{
					ID:       3,
					UID:      "ctesttest",
					Email:    nil,
					Password: "d0993jd",
					Role:     domain.UserRoleAdmin,
					BanState: domain.UserBanStateUnBan,
				},
			},
		},
	}

	for i, tc := range testCases {
		name := fmt.Sprintf("#%d ListUser", i)
		t.Run(name, func(tsub *testing.T) {

			got, err := userRepo.ListUser(ctx, tc.skip, tc.limit, tc.order, tc.critertion)
			if err != nil {
				t.Errorf("failed to list user : %v", err)
			}

			for i := 0; i < len(got); i++ {

				tc.expect[i].Email = got[i].Email
				tc.expect[i].CreatedAt = got[i].CreatedAt
				tc.expect[i].UpdatedAt = got[i].UpdatedAt
				tc.expect[i].DeletedAt = got[i].DeletedAt

				if !reflect.DeepEqual(got[i], tc.expect[i]) {
					tsub.Errorf("\ntc.expect[i] and got[i] don't have each other.\ngot[i]=%+v\nexpect[i]=%+v\n", got[i], tc.expect[i])

				}
			}

		})
	}

}

func TestUpdateUser(t *testing.T) {
	defer provider.Get().TerminateContainer(provider.MyContext)

	ctx := context.Background()

	// database seedling
	err := func() error {
		// field 1
		if _, err := userRepo.CreateUser(ctx, &domain.User{
			UID:      "updateuserid",
			Password: "updateuserpassword",
			Email:    strPtr("2hanjunbum6@gmail.com"),
			Role:     domain.UserRoleAdmin,
			BanState: domain.UserBanStateBan,
		}); err != nil {
			return err
		}

		return nil

	}()
	if err != nil {
		t.Errorf("failed to create user: %v", err)
	}

	testCases := []struct {
		user   *domain.User
		expect *domain.User
	}{
		{
			user: &domain.User{
				ID:       1,
				UID:      "hahahahaid",
				Password: "hahahahapassword",
				Email:    strPtr("gogogo@gmail.com"),
				Role:     domain.UserRoleUser,
				BanState: domain.UserBanStateUnBan,
			},

			expect: &domain.User{
				ID:       1,
				UID:      "hahahahaid",
				Password: "hahahahapassword",
				Email:    strPtr("gogogo@gmail.com"),
				Role:     domain.UserRoleUser,
				BanState: domain.UserBanStateUnBan,
			},
		},
	}

	for i, tc := range testCases {
		name := fmt.Sprintf("#%d UpdateUser", i)
		t.Run(name, func(tsub *testing.T) {

			got, err := userRepo.UpdateUser(ctx, tc.user)
			if err != nil {
				tsub.Errorf("failed to update user: %v", err)
			}

			if *tc.expect.Email != *got.Email {
				tsub.Errorf("\ntc.expect'email and got'email don't have each other.\ngot=%+v\nexpect=%+v\n", got.Email, tc.expect.Email)
			}

			tc.expect.Email = got.Email

			tc.expect.CreatedAt = got.CreatedAt
			tc.expect.UpdatedAt = got.UpdatedAt
			tc.expect.DeletedAt = got.DeletedAt

			if !reflect.DeepEqual(tc.expect, got) {
				tsub.Errorf("\ntc.expect and got don't have each other.\ngot=%+v\nexpect=%+v\n", got, tc.expect)
			}
		})
	}
}

package test

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"toktok-backend/internal/adapter/persistence/mysql/repository"
	"toktok-backend/internal/adapter/persistence/mysql/test/provider"
	"toktok-backend/internal/core/domain"
)

var avatarRepo *repository.AvatarRepository

func init() {
	avatarRepo = repository.NewAvatarRepository(provider.Get().Client)
}

func TestCreateAvatar(t *testing.T) {
	defer provider.Get().TerminateContainer(provider.MyContext)

	//setting
	ctx := context.Background()

	// seed
	_, err := userRepo.CreateUser(ctx, &domain.User{
		UID:      "2hanjunbum6",
		Password: "12345678",
	})
	if err != nil {
		t.Errorf("[Seed ERROR] failed to create user: %v", err)
	}

	// TestCases
	testCases := []struct {
		avatar *domain.Avatar
		expect *domain.Avatar
	}{
		{
			avatar: &domain.Avatar{
				UserID:    1,
				Sex:       domain.AvatarSexMale,
				Birthday:  "20240607",
				Mbti:      strPtr("infp"),
				Nickname:  "fullgukbap",
				Picture:   "http://test.com",
				Introduce: strPtr("안녕하세요 ㅎㅎ 저는 fullgukbap!"),
				State:     domain.AvatarStateOffline,
			},

			expect: &domain.Avatar{
				ID:        1,
				UserID:    1,
				Sex:       domain.AvatarSexMale,
				Birthday:  "20240607",
				Mbti:      strPtr("infp"),
				Nickname:  "fullgukbap",
				Picture:   "http://test.com",
				Introduce: strPtr("안녕하세요 ㅎㅎ 저는 fullgukbap!"),
				State:     domain.AvatarStateOffline,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%d Create Avatar", i), func(t *testing.T) {
			got, err := avatarRepo.CreateAvatar(ctx, tc.avatar)
			if err != nil {
				t.Errorf("failed to create avatar: %v\n", err)
			}

			fmt.Printf("got: %+v", got)

			if *tc.expect.Mbti != *got.Mbti {
				t.Errorf("\nexpect, got not match [mbti] \nexpect: %+v\ngot: %+v\n", tc.expect, got)
			}
			tc.expect.Mbti = got.Mbti

			if *tc.expect.Introduce != *got.Introduce {
				t.Errorf("\nexpect, got not match, [introduce] \nexpect: %+v\ngot: %+v\n", tc.expect, got)
			}
			tc.expect.Introduce = got.Introduce

			tc.expect.CreatedAt = got.CreatedAt
			tc.expect.UpdatedAt = got.UpdatedAt
			tc.expect.DeletedAt = got.DeletedAt

			if !reflect.DeepEqual(got, tc.expect) {
				t.Errorf("\nexpect, got not match\nexpect: %+v\ngot: %+v\n", tc.expect, got)
			}
		})
	}
}

func TestGetAvatar(t *testing.T) {
	defer provider.Get().TerminateContainer(provider.MyContext)

	//setting
	ctx := context.Background()

	// seed
	_, err := userRepo.CreateUser(ctx, &domain.User{
		UID:      "2hanjunbum6",
		Password: "12345678",
	})
	if err != nil {
		t.Errorf("[Seed ERROR] failed to create user: %v", err)
	}

	createdAvatar, err := avatarRepo.CreateAvatar(ctx, &domain.Avatar{
		UserID:   1,
		Sex:      domain.AvatarSexFemale,
		Birthday: "20240606",
		Nickname: "fullgukbap",
		Picture:  "http://test.com",
		State:    domain.AvatarStateOnline,
	})
	if err != nil {
		t.Errorf("[Seed ERROR] failed to create avatar")
	}

	// testCases
	testCases := []struct {
		id     int
		expect *domain.Avatar
	}{
		// avatar.id가 1인 아바타 조회
		{
			id: 1,
			expect: &domain.Avatar{
				ID:        1,
				UserID:    1,
				Sex:       domain.AvatarSexFemale,
				Mbti:      nil,
				Birthday:  "20240606",
				Nickname:  "fullgukbap",
				Picture:   "http://test.com",
				Introduce: nil,
				State:     domain.AvatarStateOnline,

				CreatedAt: createdAvatar.CreatedAt.UTC(),
				UpdatedAt: createdAvatar.DeletedAt.UTC(),
				DeletedAt: createdAvatar.DeletedAt.UTC(),
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%d Create Avatar", i), func(t *testing.T) {
			got, err := avatarRepo.GetAvatar(ctx, tc.id)
			if err != nil {
				t.Errorf("failed to create avatar: %v\n", err)
			}

			fmt.Printf("got: %+v", got)

			if tc.expect.Mbti != got.Mbti {
				t.Errorf("\nexpect, got not match [mbti] \nexpect: %+v\ngot: %+v\n", tc.expect, got)
			}
			tc.expect.Mbti = got.Mbti

			if tc.expect.Introduce != got.Introduce {
				t.Errorf("\nexpect, got not match, [introduce] \nexpect: %+v\ngot: %+v\n", tc.expect, got)
			}
			tc.expect.Introduce = got.Introduce

			tc.expect.CreatedAt = got.CreatedAt
			tc.expect.UpdatedAt = got.UpdatedAt
			tc.expect.DeletedAt = got.DeletedAt

			if !reflect.DeepEqual(got, tc.expect) {
				t.Errorf("\nexpect, got not match\nexpect: %+v\ngot: %+v\n", tc.expect, got)
			}
		})
	}

}

func TestListAvatar(t *testing.T) {
	defer provider.Get().TerminateContainer(provider.MyContext)

	//setting
	ctx := context.Background()

	err := func() error {

		_, err := userRepo.CreateUser(ctx, &domain.User{
			UID:      "a123456",
			Password: "qwe123qwe123",
		})

		_, err = userRepo.CreateUser(ctx, &domain.User{
			UID:      "b123456",
			Password: "qwe123qwe123",
		})

		_, err = userRepo.CreateUser(ctx, &domain.User{
			UID:      "c123456",
			Password: "qwe123qwe123",
		})

		_, err = userRepo.CreateUser(ctx, &domain.User{
			UID:      "d123456",
			Password: "qwe123qwe123",
		})

		_, err = avatarRepo.CreateAvatar(ctx, &domain.Avatar{
			UserID:   1,
			Sex:      domain.AvatarSexFemale,
			Birthday: "20240606",
			Nickname: "a123456",
			Picture:  "http://test.com",
			State:    domain.AvatarStateOnline,
		})

		_, err = avatarRepo.CreateAvatar(ctx, &domain.Avatar{
			UserID:   2,
			Sex:      domain.AvatarSexFemale,
			Birthday: "20240606",
			Nickname: "b123456",
			Picture:  "http://test.com",
			State:    domain.AvatarStateOnline,
		})

		_, err = avatarRepo.CreateAvatar(ctx, &domain.Avatar{
			UserID:   3,
			Sex:      domain.AvatarSexFemale,
			Birthday: "20240606",
			Nickname: "c123456",
			Picture:  "http://test.com",
			State:    domain.AvatarStateOnline,
		})

		_, err = avatarRepo.CreateAvatar(ctx, &domain.Avatar{
			UserID:   4,
			Sex:      domain.AvatarSexFemale,
			Birthday: "20240606",
			Nickname: "d123456",
			Picture:  "http://test.com",
			State:    domain.AvatarStateOnline,
		})

		return err
	}()
	if err != nil {
		t.Errorf("failed to create avatar, user: %s", err)
	}

	// test cases
	testCases := []struct {
		skip, limit       int
		order, critertion string
		expect            []*domain.Avatar
	}{
		{
			skip:  2,
			limit: 2,
			// order: "asc", //-->  default
			// critertion: "id" --> default
			expect: []*domain.Avatar{
				&domain.Avatar{
					ID:        3,
					UserID:    3,
					Sex:       domain.AvatarSexFemale,
					Birthday:  "20240606",
					Mbti:      nil,
					Nickname:  "c123456",
					Picture:   "http://test.com",
					Introduce: nil,
					State:     domain.AvatarStateOnline,
				},

				&domain.Avatar{
					ID:        4,
					UserID:    4,
					Sex:       domain.AvatarSexFemale,
					Birthday:  "20240606",
					Nickname:  "d123456",
					Picture:   "http://test.com",
					State:     domain.AvatarStateOnline,
					Mbti:      nil,
					Introduce: nil,
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%d Create Avatar", i), func(tsub *testing.T) {
			gots, err := avatarRepo.ListAvatar(ctx, tc.skip, tc.limit, tc.order, tc.critertion)
			if err != nil {
				t.Errorf("failed to list avatar: %s", err)
			}

			for i := 0; i < len(gots); i++ {
				tc.expect[i].CreatedAt = gots[i].CreatedAt
				tc.expect[i].UpdatedAt = gots[i].UpdatedAt
				tc.expect[i].DeletedAt = gots[i].DeletedAt

				if !reflect.DeepEqual(gots[i], tc.expect[i]) {
					tsub.Errorf("\ntc.expect[i] and got[i] don't have each other.\ngot[i]=%+v\nexpect[i]=%+v\n", gots[i], tc.expect[i])
				}
			}

		})
	}

}

func TestUpdateUpdateAvatar(t *testing.T) {
	defer provider.Get().TerminateContainer(provider.MyContext)

	// setting
	ctx := context.Background()

	// seedling
	err := func() error {
		_, err := userRepo.CreateUser(ctx, &domain.User{
			UID:      "a123456",
			Password: "qwe123qwe123",
		})
		if err != nil {
			return err
		}

		_, err = avatarRepo.CreateAvatar(ctx, &domain.Avatar{
			UserID:   1,
			Sex:      domain.AvatarSexFemale,
			Birthday: "20240606",
			Nickname: "c123456",
			Picture:  "http://test.com",
			State:    domain.AvatarStateOnline,
		})

		return err
	}()
	if err != nil {
		t.Errorf("failed to seedling: %v", err)
	}

	// testCases
	testCases := []struct {
		avatar *domain.Avatar
		expect *domain.Avatar
	}{
		{
			avatar: &domain.Avatar{
				ID:       1,
				Sex:      domain.AvatarSexMale,
				Birthday: "20248888",
				Mbti:     strPtr("intj"),
				Nickname: "fullgukbap",
				Picture:  "http://test-retry.com",
			},

			expect: &domain.Avatar{
				ID:        1,
				UserID:    1,
				Sex:       domain.AvatarSexMale,
				Birthday:  "20248888",
				Mbti:      strPtr("intj"),
				Nickname:  "fullgukbap",
				Picture:   "http://test-retry.com",
				Introduce: nil,
				State:     domain.AvatarStateOnline,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%d Update Avatar", i), func(tsub *testing.T) {

			got, err := avatarRepo.UpdateAvatar(ctx, tc.avatar)
			if err != nil {
				t.Errorf("failed to update avatar: %v", err)
			}

			if *got.Mbti != *tc.expect.Mbti {
				t.Errorf("got.Mbti, tc.expect not macth: %v", err)
			}
			got.Mbti = tc.expect.Mbti

			tc.expect.CreatedAt = got.CreatedAt
			tc.expect.UpdatedAt = got.UpdatedAt
			tc.expect.DeletedAt = got.DeletedAt

			if !reflect.DeepEqual(got, tc.expect) {
				tsub.Errorf("\ntc.expect and got don't have each other.\ngot=%+v\nexpect=%+v\n", got, tc.expect)
			}
		})
	}

}

func TestDeleteAvatar(t *testing.T) {
	defer provider.Get().TerminateContainer(provider.MyContext)

	// setting
	ctx := context.Background()

	// seedling

	_, err := userRepo.CreateUser(ctx, &domain.User{
		UID:      "a123456",
		Password: "qwe123qwe123",
	})
	if err != nil {
		t.Errorf("failed to seedling: %v", err)
	}

	_, err = avatarRepo.CreateAvatar(ctx, &domain.Avatar{
		UserID:   1,
		Sex:      domain.AvatarSexFemale,
		Birthday: "20240606",
		Nickname: "c123456",
		Picture:  "http://test.com",
		State:    domain.AvatarStateOnline,
	})

	if err != nil {
		t.Errorf("failed to seedling: %v", err)
	}

	err = avatarRepo.DeleteAvatar(ctx, 1)
	if err != nil {
		t.Errorf("delete error: %v", err)
	}

}

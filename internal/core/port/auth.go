package port

import (
	"context"
)

type AuthService interface {
	Login(ctx context.Context, uid, password string) (accessToken string, refreshToken string, err error)
}

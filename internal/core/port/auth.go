package port

import (
	"context"
)

type AuthService interface {
	Login(ctx context.Context, uid, password string) (accessToken string, refreshToken string, err error)
	Refresh(ctx context.Context, token string) (accessToekn string, err error)
}

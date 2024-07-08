package port

import (
	"context"
)

// [INBOUND] AuthService 인터페이스는 인증(Authenication)과 상호 작용하기 위한 주체입니다.
type AuthService interface {
	// Login 함수는 uid, password를 인자로 받아 accessToken과 refreshToken을 반환합니다.
	Login(ctx context.Context, uid, password string) (accessToken string, refreshToken string, err error)

	// Refresh 함수는 token을 인자로 받아 accessToken을 반환합니다.
	Refresh(ctx context.Context, token string) (accessToekn string, err error)
}

package port

import "toktok-backend/internal/core/domain"

// [OUTBOUND] TokenService 인터페이스는 토큰과 관련된(인증에 가까운) 행위를 명세합니다.
type TokenService interface {
	// CreateToken 함수는 tokenPayload와 user값을 기반으로 새로운 토큰을 생성합니다.
	CreateToken(tokenType domain.TokenType, user *domain.User) (string, error)

	// VerifyToken 함수는 token을 검증하여 tokenPayload로 변환해줍니다.
	VerifyToken(token string) (*domain.TokenPlayload, error)
}

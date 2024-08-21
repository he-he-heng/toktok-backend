package domain

import "time"

type TokenPayload struct {
	// 사용자나 엔터티를 교유하게 식별하는 ID
	Sub int

	// 토큰을 발급한 사람의 userID
	Iss int

	// 토큰 만료 시간 (unix timestamp로 표현)
	Exp time.Time

	// 토큰이 발급된 시간 (unix timestamp로 표현)
	Iat time.Time

	Alg string

	Role UserRoleType
}

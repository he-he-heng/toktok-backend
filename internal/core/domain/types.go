package domain

import "errors"

var (
	// ErrInternal은 내부 서비스가 요청을 처리하지 못할 때 발생하는 오류입니다.
	ErrInternal = errors.New("내부 오류")

	// ErrInvalidCredentials은 잘못된 자격 증명일 때 발생하는 오류입니다.
	ErrInvalidCredentials = errors.New("잘못된 자격증명입니다")

	// ErrTokenCreation은 토큰 생성 실패 시 발생하는 오류입니다.
	ErrTokenCreation = errors.New("토큰 생성 오류")

	// ErrExpiredToken은 액세스 토큰이 만료된 경우 발생하는 오류입니다.
	ErrExpiredToken = errors.New("토큰이 만료되었습니다")

	// ErrInvalidToken은 액세스 토큰이 잘못된 경우 발생하는 오류입니다.
	ErrInvalidToken = errors.New("잘못된 토큰입니다")

	// ErrTokenDuration은 토큰 기간 형식이 잘못된 경우 발생하는 오류입니다.
	ErrTokenDuration = errors.New("잘못된 토큰 기간 형식입니다")

	// ErrEmptyAuthorizationHeader는 인증 헤더가 비어있는 경우 발생하는 오류입니다.
	ErrEmptyAuthorizationHeader = errors.New("인증 헤더가 제공되지 않았습니다")

	// ErrInvalidAuthorizationHeader는 인증 헤더 형식이 잘못된 경우 발생하는 오류입니다.
	ErrInvalidAuthorizationHeader = errors.New("인증 헤더 형식이 잘못되었습니다")

	// ErrNotAccessToken은 액세스 토큰이 아닐떄 사용하는 오류입니다.
	ErrNotAccessToken = errors.New("요청된 토큰은 액세스 토큰이 아닙니다")

	// ErrNotRefreshToken은 토큰이 리프레시 토큰이 아닐 때 사하는 오류입니다.
	ErrNotRefreshToken = errors.New("요청된 토큰은 리프레시 토큰이 아닙니다")

	ErrInvalidTokenSignMethod = errors.New("요청한 토큰의 sign 메소드가 일치하지 않습니다")

	ErrInvalidBodyParsing = errors.New("요청된 값을 파싱하는데 실패했습니다")

	// --- 데이터베이스 ---

	// ErrDataNotFound는 요청한 데이터를 찾을 수 없을 때 발생하는 오류입니다.
	ErrNotFound = errors.New("데이터를 찾을 수 없습니다")

	// ErrValidadtion은 데이터베이스에 삽입할 요소들이 유효하지 않을 때 발생하는 에러 입니다.
	ErrValidation = errors.New("데이터가 유효하지 않습니다")

	// ErrConstraint은 데이터베이스의 제약조건이 실패할 때 발생하는 에러 입니다.
	ErrConstraint = errors.New("제약 조건이 실패했습니다")
)

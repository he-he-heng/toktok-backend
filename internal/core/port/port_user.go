package port

import (
	"context"
	"toktok-backend/internal/core/domain"
)

// [OUTBOUND] UserRepository 인터페이스는 사용자 관련 데이터와 상호 작용하기 위한 주체입니다.
type UserRepository interface {
	// CreateUser 함수는 데이터베이스에 새 사용자를 생성합니다.
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// GetUserByID 함수는 아이디로 사용자를 조회합니다.
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	// GetUserByUID 함수는 uid로 사용자를 조회합니다.
	GetUserByUID(ctx context.Context, uid string) (*domain.User, error)
	// ListUsers 함수는 페이지 매김이 있는 사용자 목록을 선택합니다.
	ListUsers(ctx context.Context, skip, limit int) ([]domain.User, error)
	// UpdateUser 함수는 사용자를 업데이트 합니다.
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// DeleteUser 함수는 id로 사용자를 제거 합니다.
	DeleteUser(ctx context.Context, id int) error
}

// [INBOUND] UserService 인터페이스는 사용자 관련 비지니스 로직과 상호 작용하기 위한 주체 입니다.
type UserService interface {
	// Register 함수는 새로운 사용자를 등록합니다.
	Register(ctx context.Context, user *domain.User) (*domain.User, error)
	// GetUser 함수는 id로 사용자를 조회합니다.
	GetUser(ctx context.Context, id int) (*domain.User, error)
	// ListUsers 함수는 페이지 매김이 있는 사용자 목록을 선택합니다.
	ListUsers(ctx context.Context, skip, limit int) ([]domain.User, error)
	// UpdateUser 함수는 사용자를 업데이트 합니다.
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// DeleteUser 함수는 id로 사용자를 제거 합니다.
	DeleteUser(ctx context.Context, id int) error
}

package domain

type TokenType string

const (
	RefreshToken = "refreshToken"
	AccessToken  = "accessToken"
)

type TokenPlayload struct {
	UID  int
	Role RoleType

	// Expiration Time (토큰 만료 시간)
	Exp int64

	// Issued At (토큰 발급 시간)
	Ita int64

	TokenType TokenType
}

package domain

type TokenPayload struct {
	ID     int
	UserID int
	Role   UserRoleType
}

package dto

type RegisterReqeust struct {
	UID      string `json:"uid" binding:"required,regexp=^[A-Za-z0-9]+$"`
	Password string `json:"password" binding:"required,regexp=^^[A-Za-z0-9!@#$%^&*]+$"`
}

package dto

type RegisterRequest struct {
	UID      string `json:"uid" validate:"required"`
	Password string `json:"password" validate:"required"`
}

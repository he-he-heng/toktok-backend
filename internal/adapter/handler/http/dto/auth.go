package dto

type LoginRequest struct {
	UID      string `json:"uid" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type RefreshReponse struct {
	AccessToken string `json:"accessToken"`
}

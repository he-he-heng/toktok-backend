package dto

type LoginRequest struct {
	UID      string `json:"uid" binding:"required,regexp=^[A-Za-z0-9]+$"`
	Password string `json:"password" binding:"required,regexp=^^[A-Za-z0-9!@#$%^&*]+$"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refershToken"`
}

type ReissueTokenReqeust struct {
	RefreshToken string `json:"accessToken" binding:"required"`
}

type ReissueTokenResponsea struct {
	AccessTokens string `json:"accessToken"`
}

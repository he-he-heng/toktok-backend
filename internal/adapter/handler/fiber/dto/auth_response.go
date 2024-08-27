package dto

import "toktok-backend/internal/core/domain"

type AuthLoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (AuthLoginResponse) Of(accessToken, refreshToken string) AuthLoginResponse {
	authLoginResponse := AuthLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return authLoginResponse
}

type AuthRefreshResponse struct {
	AccessToken string `json:"acessToken"`
}

func (AuthRefreshResponse) Of(accessToken string) AuthRefreshResponse {
	authRefreshResponse := AuthRefreshResponse{
		AccessToken: accessToken,
	}

	return authRefreshResponse
}

type AuthValidationResponse struct {
	Iss       int    `json:"iss"`
	Exp       int64  `json:"exp"`
	Ita       int64  `json:"ita"`
	TokenType string `json:"tokenType"`
	Role      string `json:"role"`
}

func (AuthValidationResponse) Of(tokenPayload *domain.TokenPayload) AuthValidationResponse {
	authValidationResponse := AuthValidationResponse{
		Iss:       tokenPayload.Iss,
		Exp:       tokenPayload.Exp,
		Ita:       tokenPayload.Ita,
		TokenType: string(tokenPayload.TokenType),
		Role:      string(tokenPayload.Role),
	}

	return authValidationResponse
}

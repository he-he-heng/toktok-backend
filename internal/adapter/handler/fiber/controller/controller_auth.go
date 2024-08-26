package controller

import "toktok-backend/internal/core/port"

type AuthController struct {
	authService port.AuthService
}

func NewAuthController(authService port.AuthService) *AuthController {
	authController := AuthController{
		authService: authService,
	}

	return &authController
}

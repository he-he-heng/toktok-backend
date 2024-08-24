package controller

import "toktok-backend/internal/core/port"

type AvatarController struct {
	avatarService port.UserService
}

func NewAvatarController(avatarService port.UserService) *AvatarController {
	userController := AvatarController{
		avatarService: avatarService,
	}

	return &userController
}

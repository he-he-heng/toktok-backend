package http

import "toktok-backend/internal/core/port"

type UserHandler struct {
	userService port.UserService
}

func NewUserHandler(userSerivce port.UserService) *UserHandler {
	return &UserHandler{
		userService: userSerivce,
	}
}

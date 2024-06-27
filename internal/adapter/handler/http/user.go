package http

import (
	"toktok-backend/internal/adapter/handler/http/dto"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService port.UserService
}

func NewUserHandler(userService port.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	// TODO validation

	user := domain.User{
		UID:      req.UID,
		Password: req.Password,
	}

	h.userService.Register(c.Context(), &user)
}

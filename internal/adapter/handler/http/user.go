package http

import (
	"toktok-backend/internal/adapter/handler/http/dto"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/core/port"
	"toktok-backend/pkg/errors"

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
		return errors.Wrap(domain.ErrInvalidBodyParsing, err)
	}

	// TODO validation

	user := domain.User{
		UID:      req.UID,
		Password: req.Password,
	}

	_, err := h.userService.Register(c.Context(), &user)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&dto.General{
		Status:  fiber.StatusCreated,
		Message: "성공적으로 회원을 생성했습니다.",
	})
}

package http

import (
	"toktok-backend/internal/adapter/handler/http/dto"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/core/port"
	"toktok-backend/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService port.AuthService
}

func NewAuthHandler(authService port.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return errors.Wrap(domain.ErrInvalidBodyParsing, err)
	}

	accessToken, refreshToken, err := h.authService.Login(c.Context(), req.UID, req.Password)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(&dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
	// return c.Status(fiber.StatusCreated).JSON(&dto.General{
	// 	// Status:  fiber.StatusCreated,
	// 	// // Message: "토큰을 성공적으로 생성했습니다.",
	// 	Data: &dto.LoginResponse{
	// 		AccessToken:  accessToken,
	// 		RefreshToken: refreshToken,
	// 	},
	// })

}

func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	var req dto.RefreshRequest
	if err := c.BodyParser(&req); err != nil {
		return errors.Wrap(domain.ErrInvalidBodyParsing, err)
	}

	accssToken, err := h.authService.Refresh(c.Context(), req.RefreshToken)
	if err != nil {
		return err
	}

	// return c.Status(fiber.StatusOK).JSON(&dto.General{
	// 	Status:  fiber.StatusOK,
	// 	Message: "성공적으로 재발급을 완료했습니다.",
	// 	Data: &dto.RefreshReponse{
	// 		AccessToken: accssToken,
	// 	},
	// })

	return c.Status(fiber.StatusOK).JSON(&dto.RefreshReponse{
		AccessToken: accssToken,
	})
}

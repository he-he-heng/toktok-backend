package controller

import (
	"fmt"
	"toktok-backend/internal/adapter/handler/fiber/dto"
	"toktok-backend/internal/adapter/handler/fiber/middleware"
	"toktok-backend/internal/adapter/handler/fiber/validator"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/core/port"
	"toktok-backend/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService port.AuthService
}

func NewAuthController(authService port.AuthService) *AuthController {
	authController := AuthController{
		authService: authService,
	}

	return &authController
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	body := dto.AuthLoginRequest{}
	if err := ctx.BodyParser(&body); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	if err := validator.Get().Verify(&body); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	accessToken, refreshToken, err := c.authService.Login(ctx.Context(), body.ToDomainUser())
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(dto.AuthLoginResponse{}.Of(accessToken, refreshToken))
}

func (c *AuthController) Refresh(ctx *fiber.Ctx) error {
	body := dto.AuthRefreshRequest{}
	if err := ctx.BodyParser(&body); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	if err := validator.Get().Verify(&body); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	accessToken, err := c.authService.Refresh(ctx.Context(), body.RefreshToken)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.AuthRefreshResponse{}.Of(accessToken))
}

func (c *AuthController) Validation(ctx *fiber.Ctx) error {
	warpData := ctx.Locals(middleware.AuthorizationPayloadKey)
	fmt.Println(warpData)
	data, ok := warpData.(*domain.TokenPayload)
	if !ok {
		return errors.Wrap(domain.ErrInternalServerError, "warpData not convert to *domain.TokenPayload")
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.AuthValidationResponse{}.Of(data))
}

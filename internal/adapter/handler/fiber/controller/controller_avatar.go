package controller

import (
	"toktok-backend/internal/adapter/handler/fiber/dto"
	"toktok-backend/internal/adapter/handler/fiber/validator"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/core/port"
	"toktok-backend/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

type AvatarController struct {
	avatarService port.AvatarService
}

func NewAvatarController(avatarService port.AvatarService) *AvatarController {
	userController := AvatarController{
		avatarService: avatarService,
	}

	return &userController
}

func (c *AvatarController) CreateAvatar(ctx *fiber.Ctx) error {
	body := dto.CreateAvatarRequest{}
	if err := ctx.BodyParser(&body); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	if err := validator.Get().Verify(&body); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	avatar, err := c.avatarService.CreateAvatar(ctx.Context(), body.ToDomainAvatar())
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateAvatarResponse{}.Of(avatar))
}

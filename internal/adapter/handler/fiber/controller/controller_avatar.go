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

	return ctx.Status(fiber.StatusCreated).JSON(dto.CreateAvatarResponse{}.Of(avatar))
}

func (c *AvatarController) ListAvatar(ctx *fiber.Ctx) error {
	skip := ctx.QueryInt("skip", 1)
	limit := ctx.QueryInt("limit", 5)
	order := ctx.Query("order", "asc", "asc")
	criterion := ctx.Query("criterion", "id")

	users, err := c.avatarService.ListAvatar(ctx.Context(), skip, limit, order, criterion)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.ListAvatarResponse{}.Of(users))
}

func (c *AvatarController) GetAvatar(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 1)
	if err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	user, err := c.avatarService.GetAvatar(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.GetAvatarResponse{}.Of(user))
}

func (c *AvatarController) UpdateAvatar(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 1)
	if err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	body := dto.UpdateAvatarRequest{}
	if err := ctx.BodyParser(&body); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	if err := validator.Get().Verify(&body); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	avatar, err := c.avatarService.UpdateAvatar(ctx.Context(), body.ToDomainAvatar(id))
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.UpdateAvatarResponse{}.Of(avatar))

}

func (c *AvatarController) DeleteAvatar(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 1)
	if err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	err = c.avatarService.DeleteAvatar(ctx.Context(), id)
	if err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

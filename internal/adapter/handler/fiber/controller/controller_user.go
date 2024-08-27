package controller

import (
	"toktok-backend/internal/adapter/handler/fiber/dto"
	"toktok-backend/internal/adapter/handler/fiber/validator"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/core/port"
	"toktok-backend/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService port.UserService
}

func NewUserController(userService port.UserService) *UserController {
	userController := UserController{
		userService: userService,
	}

	return &userController
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	body := dto.CreateUserRequest{}
	if err := ctx.BodyParser(&body); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}
	if err := validator.Get().Verify(&body); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	user, err := c.userService.CreateUser(ctx.Context(), body.ToDomainUser())
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateUserRequest{}.Of(user))
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 1)
	if err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	user, err := c.userService.GetUser(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.GetUserResponse{}.Of(user))

}

func (c *UserController) ListUser(ctx *fiber.Ctx) error {
	skip := ctx.QueryInt("skip", 1)
	limit := ctx.QueryInt("limit", 5)
	order := ctx.Query("order", "asc", "asc")
	criterion := ctx.Query("criterion", "id")

	if skip > 30 {
		return errors.Wrap(domain.ErrBadParam, "value of skip is greater than 30")
	}

	users, err := c.userService.ListUser(ctx.Context(), skip, limit, order, criterion)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.UserListResponse{}.Of(users))
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil || id < 1 {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	body := dto.UpdateUserReqeust{}
	if err := ctx.BodyParser(&body); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	user, err := c.userService.UpdateUser(ctx.Context(), body.ToDomainUser(id))
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.UpdateUserResponse{}.Of(user))
}

func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil || id < 1 {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	err = c.userService.DeleteUser(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

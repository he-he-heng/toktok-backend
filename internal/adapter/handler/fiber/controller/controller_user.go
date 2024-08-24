package controller

import (
	"toktok-backend/internal/adapter/handler/fiber/dto"
	"toktok-backend/internal/core/port"

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
		return err
	}

	_, err := c.userService.CreateUser(ctx.Context(), body.ToDomainUser())
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 1)
	if err != nil {
		return err
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
	order := ctx.Query("order", "asc")
	criterion := ctx.Query("criterion", "id")

	if skip > 30 {
		// TODO: error 반환
	}

	users, err := c.userService.ListUser(ctx.Context(), skip, limit, order, criterion)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.UserListResponse{}.Of(users))
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.QueryInt("id", 1)
	body := dto.UpdateUserReqeust{}
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	_, err := c.userService.UpdateUser(ctx.Context(), body.ToDomainUser(id))
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.QueryInt("id", -1)
	if id == -1 {
		// TODO: error 반환
	}

	err := c.userService.DeleteUser(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)

}

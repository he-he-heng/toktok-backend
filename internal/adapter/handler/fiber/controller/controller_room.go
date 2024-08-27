package controller

import (
	"toktok-backend/internal/adapter/handler/fiber/dto"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/core/port"
	"toktok-backend/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

type RoomController struct {
	roomService port.RoomService
}

func NewRoomController(roomService port.RoomService) *RoomController {
	roomController := RoomController{
		roomService: roomService,
	}

	return &roomController
}

// func (c *RoomController) CreateRoom(ctx *fiber.Ctx) error {

// }

func (c *RoomController) GetRoom(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	if id < 0 {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	queriedRoom, err := c.roomService.GetRoom(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.GetRoomResponse{}.Of(queriedRoom))
}

// func (c *RoomController) ListRoom(ctx *fiber.Ctx) error {

// }

// func (c *RoomController) DeleteRoom(ctx *fiber.Ctx) error {

// }

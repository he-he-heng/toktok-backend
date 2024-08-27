package controller

import (
	"fmt"
	"toktok-backend/internal/adapter/handler/fiber/dto"
	"toktok-backend/internal/adapter/handler/fiber/validator"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/core/port"
	"toktok-backend/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

type RelationController struct {
	relationService port.RelationService
}

func NewRelationController(relationService port.RelationService) *RelationController {
	relationController := RelationController{
		relationService: relationService,
	}

	return &relationController
}

func (c *RelationController) CreateRelation(ctx *fiber.Ctx) error {
	body := dto.CreateRelationRequest{}
	if err := ctx.BodyParser(&body); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	fmt.Printf("parsing body: %+v\n", body)

	if err := validator.Get().Verify(&body); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	relations, err := c.relationService.CreateRelation(ctx.Context(), body.ToDomainRelation())
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(dto.CreateRelationResponse{}.Of(relations))
}

func (c *RelationController) GetRelation(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 1)
	if err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	relation, err := c.relationService.GetRelation(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.GetRelationResponse{}.Of(relation))

}

func (c *RelationController) ListRelation(ctx *fiber.Ctx) error {
	skip := ctx.QueryInt("skip", 1)
	limit := ctx.QueryInt("limit", 5)
	order := ctx.Query("order", "asc", "asc")
	criterion := ctx.Query("criterion", "id")
	filter := ctx.Query("filter", "")
	avatarID := ctx.QueryInt("avatar_id", 0)

	if avatarID > 0 {
		relations, err := c.relationService.ListRelationByAvatarID(ctx.Context(), skip, limit, order, criterion, domain.RelationStateType(filter), avatarID)
		if err != nil {
			return err
		}

		return ctx.Status(fiber.StatusOK).JSON(dto.ListRelationResponse{}.Of(relations))
	}

	relations, err := c.relationService.ListRelation(ctx.Context(), skip, limit, order, criterion, domain.RelationStateType(filter))
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.ListRelationResponse{}.Of(relations))
}

func (c *RelationController) UpdateRelation(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id", 0)
	if err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	if id < 1 {
		return errors.Wrap(domain.ErrBadParam, "id is less than 1")
	}

	body := dto.UpdateRelationRequest{}
	if err := ctx.BodyParser(&body); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	if err := validator.Get().Verify(&body); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	relation, err := c.relationService.UpdateRelation(ctx.Context(), body.ToDomainRelation(id))
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.UpdateRelationResponse{}.Of(relation))
}

func (c *RelationController) DeleteRelation(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}
	if id < 1 {
		return errors.Wrap(domain.ErrBadParam, "id is less than 1")
	}

	body := dto.DeleteRelationRequest{}
	if err := ctx.BodyParser(&body); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	if err := validator.Get().Verify(&body); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	if err = c.relationService.DeleteRelation(ctx.Context(), id); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)

}

func (c *RelationController) GetRoomByRealtionID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	if id < 0 {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	room, err := c.relationService.GetRoomByRelationID(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.GetRoomByRelationByIDResponse{}.Of(room))
}

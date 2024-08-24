package catcher

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Default error handler
var CustomErrorHandler = func(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	msg := err.Error()
	status := fiber.StatusInternalServerError

	gotStatus, originErrMsg, err := errSetGet().Get(err)
	if err == nil {
		status = gotStatus
		msg = fmt.Sprintf("[%s] %s", originErrMsg, msg)
	}

	return ctx.Status(status).JSON(&errResponse{
		Status: status,
		Msg:    msg,
	})
}

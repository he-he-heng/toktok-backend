package myerror

import (
	"toktok-backend/internal/adapter/handler/http/dto"
	"toktok-backend/internal/core/domain"
	"toktok-backend/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

type pair struct {
	statusCode int
	err        error
}

type errorMapper struct {
	pairs []pair
}

func newErrorMapper() *errorMapper {
	return &errorMapper{
		pairs: make([]pair, 0),
	}
}

func (e *errorMapper) register(pairs ...pair) {
	e.pairs = append(e.pairs, pairs...)
}

func (e *errorMapper) findStatus(err error) (int, error) {
	for _, pair := range e.pairs {
		if errors.Equal(err, pair.err) {
			return pair.statusCode, nil
		}
	}

	return 0, errors.New("not found status")
}

func (e *errorMapper) autoSetting() *errorMapper {
	pairs := []pair{
		{statusCode: fiber.StatusInternalServerError, err: domain.ErrInternal},
		{statusCode: fiber.StatusUnauthorized, err: domain.ErrInvalidCredentials},
		{statusCode: fiber.StatusInternalServerError, err: domain.ErrTokenCreation},
		{statusCode: fiber.StatusUnauthorized, err: domain.ErrExpiredToken},
		{statusCode: fiber.StatusUnauthorized, err: domain.ErrInvalidToken},
		{statusCode: fiber.StatusBadRequest, err: domain.ErrTokenDuration},
		{statusCode: fiber.StatusUnauthorized, err: domain.ErrEmptyAuthorizationHeader},
		{statusCode: fiber.StatusUnauthorized, err: domain.ErrInvalidAuthorizationHeader},
		{statusCode: fiber.StatusUnauthorized, err: domain.ErrNotAccessToken},
		{statusCode: fiber.StatusUnauthorized, err: domain.ErrNotRefreshToken},
		{statusCode: fiber.StatusUnauthorized, err: domain.ErrInvalidTokenSignMethod},
		{statusCode: fiber.StatusBadRequest, err: domain.ErrInvalidBodyParsing},
		{statusCode: fiber.StatusNotFound, err: domain.ErrNotFound},
		{statusCode: fiber.StatusBadRequest, err: domain.ErrValidation},
		{statusCode: fiber.StatusConflict, err: domain.ErrConstraint},
	}

	e.register(pairs...)
	return e
}

var mapper = newErrorMapper().autoSetting()

func Handler(c *fiber.Ctx, argErr error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(argErr, &e) {
		code = e.Code
	}

	findCode, err := mapper.findStatus(argErr)
	if err == nil {
		code = findCode
	}

	// Set Content-Type: text/plain; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	// Return status code with error message
	return c.Status(code).JSON(&dto.General{
		Status:  code,
		Message: argErr.Error(),
	})
}

package catcher

import (
	"sync"
	"toktok-backend/internal/core/domain"
	"toktok-backend/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

var es *errorSet
var once sync.Once

type statusAndErr struct {
	status int
	err    error
}

func newStatusAndErr(status int, err error) statusAndErr {
	return statusAndErr{
		status: status,
		err:    err,
	}
}

type errorSet struct {
	statusAndErrs []statusAndErr
}

func (es *errorSet) add(statusCode int, errs ...error) {
	for _, err := range errs {
		es.statusAndErrs = append(es.statusAndErrs, newStatusAndErr(statusCode, err))
	}
}

func (es *errorSet) Get(argErr error) (int, string, error) {
	for _, statusAndErr := range es.statusAndErrs {
		if errors.DeepEqual(argErr, statusAndErr.err) {
			return statusAndErr.status, statusAndErr.err.Error(), nil
		}
	}

	return 0, "", errors.New("err not add in errorSet")
}

func errSetGet() *errorSet {
	if es == nil {
		once.Do(func() {
			es = &errorSet{
				statusAndErrs: make([]statusAndErr, 10),
			}
		})

		// StatusBadRequest 400
		es.add(fiber.StatusBadRequest, domain.ErrNotLoaded, domain.ErrValidation, domain.ErrNotSingular, domain.ErrBadParam)

		// StatusUnauthorized
		es.add(fiber.StatusUnauthorized, domain.ErrInvalidCredentials, domain.ErrUnauthorized)

		// StatusNotFound 404
		es.add(fiber.StatusNotFound, domain.ErrNotFound)

		// StatusConflict 409
		es.add(fiber.StatusConflict, domain.ErrConstraint)

	}

	return es
}

// 400 -> domain.NotFound, domain.conflict

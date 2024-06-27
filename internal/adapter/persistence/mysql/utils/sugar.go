package utils

import (
	"toktok-backend/internal/adapter/persistence/mysql/ent"
	"toktok-backend/internal/core/domain"

	"github.com/pkg/errors"
)

func Wrap(err error, args ...string) (r error) {
	r = domain.ErrInternal

	if ent.IsNotFound(err) {
		r = errors.Wrap(domain.ErrNotFound, err.Error())
	}

	if ent.IsValidationError(err) {
		r = errors.Wrap(domain.ErrValidation, err.Error())
	}

	if ent.IsConstraintError(err) {
		r = errors.Wrap(domain.ErrConstraint, err.Error())
	}

	for _, arg := range args {
		r = errors.Wrap(r, arg)
	}

	return r
}

package utils

import (
	"toktok-backend/internal/adapter/persistence/mysql/ent"
	"toktok-backend/internal/core/domain"
	"toktok-backend/pkg/errors"
)

func ErrWrap(err error, args ...string) (r error) {

	if ent.IsNotFound(err) {
		r = errors.Wrap(domain.ErrNotFound, err)
	} else if ent.IsValidationError(err) {
		r = errors.Wrap(domain.ErrValidation, err)
	} else if ent.IsConstraintError(err) {
		r = errors.Wrap(domain.ErrConstraint, err)
	} else if ent.IsNotLoaded(err) {
		r = errors.Wrap(domain.ErrNotLoaded, err)
	} else if ent.IsNotSingular(err) {
		r = errors.Wrap(domain.ErrNotSingular, err)
	} else {
		r = errors.Wrap(domain.ErrInternalServerError, err)
	}

	for _, arg := range args {
		r = errors.Wrap(r, arg)
	}

	return r
}

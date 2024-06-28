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
	} else {
		r = errors.Wrap(domain.ErrInternal, err)
	}

	for _, arg := range args {
		r = errors.Wrap(r, arg)
	}

	return r
}

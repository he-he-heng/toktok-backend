package http

import (
	"toktok-backend/internal/core/domain"

	"github.com/go-playground/validator/v10"
)

// userRoleValidator is a custom validator for validating user roles
var userRoleValidator validator.Func = func(fl validator.FieldLevel) bool {
	userRole := fl.Field().Interface().(domain.RoleType)

	switch userRole {
	case "admin", "user":
		return true
	default:
		return false
	}
}

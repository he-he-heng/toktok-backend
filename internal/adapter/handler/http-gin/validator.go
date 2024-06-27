package http

import (
	"regexp"
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

var regexpValidator validator.Func = func(fl validator.FieldLevel) bool {
	// 태그에서 정규식을 가져옴
	pattern := fl.Param()
	value := fl.Field().String()

	// 정규식 컴파일
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}

	// 값이 정규식을 만족하는지 확인
	return re.MatchString(value)
}

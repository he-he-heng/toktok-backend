package domain

import "toktok-backend/pkg/errors"

// 공통으로 사용하는 에러
var (
	ErrBadParam            = errors.New("given param is not valid")
	ErrInternalServerError = errors.New("internal server error")
)

// handler에서 사용되는 에러
var ()

// service 에서 사용되는 에러
var ()

// auth 에서 사용되는 에러
var (
	ErrInvalidCredentials = errors.New("err invalid credentials")
	ErrUnauthorized       = errors.New("unauthorized")
)

// repository 에서 사용되는 에러
var (
	ErrNotFound    = errors.New("not found")      // 404
	ErrConstraint  = errors.New("constraint")     // 409
	ErrNotLoaded   = errors.New("not loaded")     // 400
	ErrNotSingular = errors.New("not singular")   // 400
	ErrValidation  = errors.New("validation err") // 400
)

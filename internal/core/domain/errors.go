package domain

import "errors"

// 공통으로 사용하는 에러
var (
	ErrBadParamInput       = errors.New("given param is not valid")
	ErrInternalServerError = errors.New("internal Server Error")
)

// handler에서 사용되는 에러
var ()

// service 에서 사용되는 에러
var ()

// repository 에서 사용되는 에러
var (
	ErrNotFound = errors.New("your requsted Item is not found")
	ErrConflict = errors.New("your item already exist")
)

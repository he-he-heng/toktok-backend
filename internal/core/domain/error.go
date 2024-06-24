package domain

import "errors"

var (
	// ErrDataNotFound is an error for when requested data is not found
	ErrDataNotFound = errors.New("data not found")

	// ErrInternal is an error for when an internal service fails to process the request
	ErrInternal = errors.New("internal error")

	// ErrInvalidCredentials is an error for when the credentials are invalid
	ErrInvalidCredentials = errors.New("invalid id or password")

	// ErrTokenCreation is an error for when the token creation fails
	ErrTokenCreation = errors.New("error creating token")
)

package domain

import "errors"

var (

	// ErrInternal is an error for when an internal service fails to process the request
	ErrInternal = errors.New("internal error")

	// ErrInvalidCredentials is an error for when the credentials are invalid
	ErrInvalidCredentials = errors.New("invalid id or password")

	// ErrTokenCreation is an error for when the token creation fails
	ErrTokenCreation = errors.New("error creating token")

	// ErrExpiredToken is an error for when the access token is expired
	ErrExpiredToken = errors.New("token has expired")

	// ErrInvalidToken is an error for when the access token is invalid
	ErrInvalidToken = errors.New("token is invalid")

	// ErrTokenDuration is an error for when the token duration format is invalid
	ErrTokenDuration = errors.New("invalid token duration format")

	// ErrEmptyAuthorizationHeader is an error for when the authorization header is empty
	ErrEmptyAuthorizationHeader = errors.New("authorization header is not provided")

	// ErrInvalidAuthorizationHeader is an error for when the authorization header is invalid
	ErrInvalidAuthorizationHeader = errors.New("authorization header format is invalid")

	ErrNotAccessToken = errors.New("requested token is not an accedss token")

	ErrNotRefreshToken = errors.New("requested token is not an refersh token")

	// --- database ---

	// ErrDataNotFound is an error for when requested data is not found
	ErrDataNotFound = errors.New("data not found")

	ErrValidation = errors.New("data not valid form")

	ErrConstraint = errors.New("constraint failed")
)

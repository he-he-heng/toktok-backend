package util

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"toktok-backend/internal/core/domain"

	"github.com/go-playground/validator/v10"
)

var errStatusMap = map[error]int{
	domain.ErrInternal:     http.StatusInternalServerError,
	domain.ErrDataNotFound: http.StatusNotFound,
	// domain.ErrConflictingData:            http.StatusConflict,
	domain.ErrInvalidCredentials: http.StatusUnauthorized,
	// domain.ErrUnauthorized:               http.StatusUnauthorized,
	domain.ErrEmptyAuthorizationHeader:   http.StatusUnauthorized,
	domain.ErrInvalidAuthorizationHeader: http.StatusUnauthorized,
	domain.ErrNotAccessToken:             http.StatusUnauthorized,
	// domain.ErrInvalidAuthorizationType:   http.StatusUnauthorized,
	domain.ErrInvalidToken: http.StatusUnauthorized,
	domain.ErrExpiredToken: http.StatusUnauthorized,
	// domain.ErrForbidden:                  http.StatusForbidden,
	// domain.ErrNoUpdatedData:              http.StatusBadRequest,
	// domain.ErrInsufficientStock:          http.StatusBadRequest,
	// domain.ErrInsufficientPayment:        http.StatusBadRequest,
}

func status(err error) int {
	statusCode, ok := errStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	return statusCode
}

// parseError parses error messages from the error object and returns a slice of error messages
func parseErr(err error) string {
	var errMsgs []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.Error())
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}

	var sb strings.Builder

	for _, msg := range errMsgs {
		sb.WriteString(fmt.Sprintf("%s ", msg))
	}

	return sb.String()
}

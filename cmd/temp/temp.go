package main

import (
	"fmt"
	"toktok-backend/pkg/errors"
)

var NotFound = errors.New("hihi")

func main() {
	err := errors.Wrap(NotFound, "not found 입니다만")

	newErr := errors.Wrap(err, "새로운 에러 입니다만")

	fmt.Println(errors.Is(newErr, NotFound))
}

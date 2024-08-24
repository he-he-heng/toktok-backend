package validator

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

type verifier struct {
	instance *validator.Validate
}

func newVerifier() *verifier {
	verifier := verifier{
		instance: validator.New(),
	}

	return &verifier
}

func (v verifier) Verify(s any) error {
	return v.instance.Struct(s)
}

var once sync.Once
var instance *verifier

func Get() *verifier {
	once.Do(func() {
		//init
		instance = newVerifier()

		// register custom tag

	})

	return instance
}

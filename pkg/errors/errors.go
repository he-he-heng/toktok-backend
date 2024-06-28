package errors

import stderrors "errors"

func New(text string) error {
	return stderrors.New(text)
}

func Is(err, target error) bool             { return stderrors.Is(err, target) }
func As(err error, target interface{}) bool { return stderrors.As(err, target) }

type withMessage struct {
	cause error
	msg   string
}

func Wrap(err error, help any) error {

	var msg string
	switch v := help.(type) {
	case string:
		msg = v
	case error:
		msg = v.Error()
	default:
		panic("only string and error")
	}

	return &withMessage{
		cause: err,
		msg:   msg,
	}
}

func (w *withMessage) Error() string { return w.msg + ": " + w.cause.Error() }
func (w *withMessage) Cause() error  { return w.cause }
func (w *withMessage) Unwrap() error { return w.cause }

func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}

// Eample
//
//	if Equal(err, domain.InvalidTokenErr) {
//	    fmt.Println("It's Equal!")
//	}
func Equal(err error, target error) bool {
	return Is(Cause(err), target)
}

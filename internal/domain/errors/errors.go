package errors

import "errors"

var (
	ErrNotFound           = errors.New("not found")
	ErrUnexpectedError    = errors.New("unexpected error")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

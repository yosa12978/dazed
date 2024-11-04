package errs

import "errors"

var (
	ErrNotFound        error = errors.New("not found")
	ErrValidationError       = errors.New("validation error")
	ErrInternalError         = errors.New("internal error")
)

func NewNotFound(err error) error {
	return errors.Join(ErrNotFound, err)
}

func NewValidationError(err error) error {
	return errors.Join(ErrValidationError, err)
}

func NewInternalError(err error) error {
	return errors.Join(ErrInternalError, err)
}

package category

import "errors"

var (
	ErrInvalidID       = errors.New("invalid id provided")
	ErrInvalidName     = errors.New("invalid name provided")
	ErrInvalidParentId = errors.New("invalid parent_id provided")

	ErrNameIsTooLong = errors.New("name must me less than 100 characters long")
	ErrNameIsEmpty   = errors.New("no name provided")
)

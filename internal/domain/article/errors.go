package article

import "errors"

var (
	ErrInvalidID          error = errors.New("invalid id provided")
	ErrInvalidTitle             = errors.New("invalid title provided")
	ErrInvalidBody              = errors.New("invalid body provided")
	ErrInvalidDescription       = errors.New("invalid description provided")

	ErrTitleIsTooLong  = errors.New("title length must be less than 100 characters")
	ErrNoTitleProvided = errors.New("no title provided")
	ErrNoBodyProvided  = errors.New("no body provoded")
	ErrDescIsTooLong   = errors.New("description length must be less than 257 characters")

	ErrChangingTitle    = errors.New("couldn't change title")
	ErrChangingDesc     = errors.New("couldn't change description")
	ErrChangingBody     = errors.New("couldn't change body")
	ErrChangingCategory = errors.New("couldn't change category")
)

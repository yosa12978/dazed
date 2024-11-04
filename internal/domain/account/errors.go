package account

import "errors"

var (
	ErrInvalidID       error = errors.New("invalid id provided")
	ErrInvalidPassword       = errors.New("invalid password provided")
	ErrInvalidUsername       = errors.New("invalid username provided")
	ErrInvalidSalt           = errors.New("invalid salt provided")
	ErrInvalidRole           = errors.New("invalid role provided")

	ErrUsernameIsTooLong              = errors.New("username length must be shorter than 32 characters")
	ErrUsernameIsTooShort             = errors.New("username length must be longer than 2 characters")
	ErrUsernameContainsForbiddenChars = errors.New("username can't contain following chars: " + forbiddenList)
	ErrPasswordIsTooShort             = errors.New("password length must be greater than 7")
	ErrSaltIsEmpty                    = errors.New("salt is an empty string")
	ErrRoleNotFound                   = errors.New("role doesn't exist")
)

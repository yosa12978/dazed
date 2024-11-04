package account

import (
	"errors"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/yosa12978/dazed/pkg/util"
)

var forbiddenList string = "!@#$%^&*()-=[]{}`\\|;:\"'<>,.?/"

type ID string
type Username string
type CreatedAt time.Time
type UpdatedAt time.Time
type Salt string
type PasswordHash string
type Password string
type Role string

func NewId(id string) (ID, error) {
	if _, err := uuid.Parse(id); err != nil {
		return "", errors.Join()
	}
	return ID(id), nil
}

func (id ID) Value() string {
	return string(id)
}

func NewPasswordHash(password string) (PasswordHash, error) {
	password = strings.TrimSpace(password)
	if len(password) < 8 {
		return "", errors.Join(ErrInvalidPassword, ErrPasswordIsTooShort)
	}
	hash, err := util.HashPassword(password)
	if err != nil {
		return "", errors.Join(ErrInvalidPassword, err)
	}
	return PasswordHash(hash), nil
}

func (p PasswordHash) Value() string {
	return string(p)
}

func NewSalt(salt string) (Salt, error) {
	salt = strings.TrimSpace(salt)
	if len(salt) == 0 {
		return "", errors.Join(ErrInvalidSalt, ErrSaltIsEmpty)
	}
	return Salt(salt), nil
}

func (s Salt) Value() string {
	return string(s)
}

func NewUsername(username string) (Username, error) {
	username = strings.TrimSpace(username)
	if len(username) < 3 {
		return "", errors.Join(ErrInvalidUsername, ErrUsernameIsTooShort)
	}
	if len(username) > 32 {
		return "", errors.Join(ErrInvalidUsername, ErrUsernameIsTooLong)
	}
	if strings.ContainsAny(username, forbiddenList) {
		return "", errors.Join(ErrInvalidUsername, ErrUsernameContainsForbiddenChars)
	}
	return Username(username), nil
}

func (u Username) Value() string {
	return string(u)
}

const (
	USER  Role = "USER"
	ADMIN Role = "ADMIN"
)

var roles = []Role{USER, ADMIN}

func NewRole(role string) (Role, error) {
	if !slices.Contains(roles, Role(role)) {
		return "", errors.Join(ErrInvalidRole, ErrRoleNotFound)
	}
	return Role(role), nil
}

func (r Role) Value() string {
	return string(r)
}

func NewCreatedAt() CreatedAt {
	return CreatedAt(time.Now().UTC())
}

func (c CreatedAt) Value() time.Time {
	return time.Time(c)
}

func NewUpdatedAt() UpdatedAt {
	return UpdatedAt(time.Now().UTC())
}

func (u UpdatedAt) Value() time.Time {
	return time.Time(u)
}

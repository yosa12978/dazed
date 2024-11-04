package account

import (
	"errors"

	"github.com/google/uuid"
	"github.com/yosa12978/dazed/pkg/util"
)

type Account struct {
	Id           ID
	Username     Username
	PasswordHash PasswordHash
	Salt         Salt
	CreatedAt    CreatedAt
	UpdatedAt    UpdatedAt
	Role         Role
}

func New(username Username, passwordHash PasswordHash, role Role) Account {
	salt, _ := NewSalt(uuid.NewString())
	id, _ := NewId(uuid.NewString())
	return Account{
		Id:           id,
		Username:     username,
		PasswordHash: passwordHash,
		Role:         role,
		Salt:         salt,
		CreatedAt:    NewCreatedAt(),
		UpdatedAt:    NewUpdatedAt(),
	}
}

func (a *Account) ChangePassword(oldPassword, newPassword string) error {
	if !util.CheckPasswordHash(oldPassword, a.PasswordHash.Value()) {
		return errors.Join(ErrInvalidPassword, errors.New("passwords doesn't match"))
	}
	salt, err := NewSalt(uuid.NewString())
	if err != nil {
		return err
	}
	a.Salt = salt
	newPasswordHash, err := NewPasswordHash(newPassword + salt.Value())
	if err != nil {
		return err
	}
	a.PasswordHash = newPasswordHash
	a.UpdatedAt = NewUpdatedAt()
	return nil
}

func (a *Account) ChangeUsername(username Username) {
	a.Username = username
	a.UpdatedAt = NewUpdatedAt()
}

func (a *Account) ChangeRole(role Role) {
	a.Role = role
	a.UpdatedAt = NewUpdatedAt()
}

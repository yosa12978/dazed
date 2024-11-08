package category

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type ID string
type Name string
type ParentID ID
type CreatedAt time.Time
type UpdatedAt time.Time

func (id ID) Value() string {
	return string(id)
}

func (name Name) Value() string {
	return string(name)
}

func (id ParentID) Value() string {
	return string(ID(id))
}

func (c CreatedAt) Value() time.Time {
	return time.Time(c)
}

func (u UpdatedAt) Value() time.Time {
	return time.Time(u)
}

func NewID(id string) (ID, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return "", errors.Join(ErrInvalidID, err)
	}
	return ID(id), nil
}

func NewName(name string) (Name, error) {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return "", errors.Join(ErrInvalidName, ErrNameIsEmpty)
	}
	if len(name) >= 100 {
		return "", errors.Join(ErrInvalidName, ErrNameIsTooLong)
	}
	return Name(name), nil
}

func NewParentID(pid string) (ParentID, error) {
	if len(pid) == 0 {
		return "", nil
	}
	id, err := NewID(pid)
	return ParentID(id), err
}

func NewCreatedAt() CreatedAt {
	return CreatedAt(time.Now().UTC())
}

func NewUpdatedAt() UpdatedAt {
	return UpdatedAt(time.Now().UTC())
}

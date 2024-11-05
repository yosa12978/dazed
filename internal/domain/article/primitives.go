package article

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type ID string
type Title string
type Description string
type Body string
type CreatedAt time.Time
type UpdatedAt time.Time

func (id ID) Value() string {
	return string(id)
}

func (t Title) Value() string {
	return string(t)
}

func (d Description) Value() string {
	return string(d)
}

func (b Body) Value() string {
	return string(b)
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

func NewTitle(title string) (Title, error) {
	title = strings.TrimSpace(title)
	if len(title) == 0 {
		return "", errors.Join(ErrInvalidTitle, ErrNoTitleProvided)
	}
	if len(title) > 100 {
		return "", errors.Join(ErrInvalidTitle, ErrTitleIsTooLong)
	}
	return Title(title), nil
}

func NewDescription(description string) (Description, error) {
	description = strings.TrimSpace(description)
	if len(description) > 256 {
		return "", errors.Join(ErrInvalidDescription, ErrDescIsTooLong)
	}
	return Description(description), nil
}

func NewBody(body string) (Body, error) {
	body = strings.TrimSpace(body)
	if len(body) == 0 {
		return "", errors.Join(ErrInvalidBody, ErrNoBodyProvided)
	}
	return Body(body), nil
}

func NewCreatedAt() CreatedAt {
	return CreatedAt(time.Now().UTC())
}
func NewUpdatedAt() UpdatedAt {
	return UpdatedAt(time.Now().UTC())
}

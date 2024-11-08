package article

import (
	"errors"

	"github.com/google/uuid"
	"github.com/yosa12978/dazed/internal/domain/account"
	"github.com/yosa12978/dazed/internal/domain/category"
)

// Fully encapsulate this (and other) structure fields
type Article struct {
	Id          ID
	Title       Title
	Description Description
	Body        Body
	CreatedAt   CreatedAt
	UpdatedAt   UpdatedAt
	Category    category.Category
	Author      account.Account
}

func New(title Title, desc Description, body Body) Article {
	id, _ := NewID(uuid.NewString())
	return Article{
		Id:          id,
		Title:       title,
		Body:        body,
		Description: desc,
		CreatedAt:   NewCreatedAt(),
		UpdatedAt:   NewUpdatedAt(),
	}
}

func (a *Article) ChangeTitle(title Title, initiator account.Account) error {
	if ok, err := a.Author.HasPermission(initiator); !ok {
		return errors.Join(ErrChangingTitle, err)
	}
	a.Title = title
	a.UpdatedAt = NewUpdatedAt()
	return nil
}

func (a *Article) ChangeDesc(desc Description, initiator account.Account) error {
	if ok, err := a.Author.HasPermission(initiator); !ok {
		return errors.Join(ErrChangingDesc, err)
	}
	a.Description = desc
	a.UpdatedAt = NewUpdatedAt()
	return nil
}

func (a *Article) ChangeBody(body Body, initiator account.Account) error {
	if ok, err := a.Author.HasPermission(initiator); !ok {
		return errors.Join(ErrChangingBody, err)
	}
	a.Body = body
	a.UpdatedAt = NewUpdatedAt()
	return nil
}

func (a *Article) ChangeCategory(category category.Category, initiator account.Account) error {
	if ok, err := a.Author.HasPermission(initiator); !ok {
		return errors.Join(ErrChangingCategory, err)
	}
	a.Category = category
	a.UpdatedAt = NewUpdatedAt()
	return nil
}

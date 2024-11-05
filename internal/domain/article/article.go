package article

import "github.com/google/uuid"

// Fully encapsulate this (and other) structure fields
type Article struct {
	Id          ID
	Title       Title
	Description Description
	Body        Body
	CreatedAt   CreatedAt
	UpdatedAt   UpdatedAt
	// Also an author and a category
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

func (a *Article) ChangeTitle(title Title) {
	a.Title = title
	a.UpdatedAt = NewUpdatedAt()
}

func (a *Article) ChangeDescription(desc Description) {
	a.Description = desc
	a.UpdatedAt = NewUpdatedAt()
}

func (a *Article) ChangeBody(body Body) {
	a.Body = body
	a.UpdatedAt = NewUpdatedAt()
}

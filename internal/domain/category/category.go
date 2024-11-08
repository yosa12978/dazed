package category

import "github.com/google/uuid"

type Category struct {
	Id        ID
	Name      Name
	CreatedAt CreatedAt
	UpdatedAt UpdatedAt
	ParentId  ParentID
}

func New(name Name, parentId ParentID) Category {
	id, _ := NewID(uuid.NewString())
	return Category{
		Id:        id,
		Name:      name,
		ParentId:  parentId,
		CreatedAt: NewCreatedAt(),
		UpdatedAt: NewUpdatedAt(),
	}
}

func (c *Category) ChangeName(name Name) {
	c.Name = name
	c.UpdatedAt = NewUpdatedAt()
}

func (c *Category) ChangeParentId(parentId ParentID) {
	c.ParentId = parentId
	c.UpdatedAt = NewUpdatedAt()
}

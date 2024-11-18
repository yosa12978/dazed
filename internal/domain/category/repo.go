package category

import "context"

type Repo interface {
	GetAll(ctx context.Context) ([]Category, error)
	GetRoots(ctx context.Context) ([]Category, error)
	GetById(ctx context.Context, id ID) (*Category, error)
	GetSubcategories(ctx context.Context, parentID ParentID) ([]Category, error)
	Create(ctx context.Context, category Category) error
	Update(ctx context.Context, category Category) error
	Delete(ctx context.Context, id ID) error
}

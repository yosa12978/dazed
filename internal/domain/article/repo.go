package article

import (
	"context"
)

type Repo interface {
	GetAll(ctx context.Context, page, size int) ([]Article, error)
	GetById(ctx context.Context, id ID) (*Article, error)
	Create(ctx context.Context, article Article) error
	Update(ctx context.Context, article Article) error
	Delete(ctx context.Context, id ID) error
}

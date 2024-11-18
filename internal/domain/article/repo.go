package article

import (
	"context"

	"github.com/yosa12978/dazed/internal/domain/category"
)

type Repo interface {
	GetAll(ctx context.Context, page, size int) ([]Article, error)
	GetById(ctx context.Context, id ID) (*Article, error)
	GetByCategory(ctx context.Context, categoryId category.ID) ([]Article, error)
	GetUncategorized(ctx context.Context) ([]Article, error)
	Create(ctx context.Context, article Article) error
	Update(ctx context.Context, article Article) error
	Delete(ctx context.Context, id ID) error
}

package account

import "context"

type Repo interface {
	GetAll(ctx context.Context, page, size int) ([]Account, error)
	GetById(ctx context.Context, id ID) (*Account, error)
	GetByUsername(ctx context.Context, username Username) (*Account, error)
	Create(ctx context.Context, account Account) error
	Update(ctx context.Context, account Account) error
	Delete(ctx context.Context, id ID) error
}

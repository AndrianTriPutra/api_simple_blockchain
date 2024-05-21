package transaction

import (
	"atp/payment/pkg/adapter/badGer"
	"context"
)

type repository struct {
	provider badGer.DatabaseI
}

func NewRepository(provider badGer.DatabaseI) RepositoryI {
	return repository{
		provider,
	}
}

type RepositoryI interface {
	SetTemp(ctx context.Context, value string) error
	Set(ctx context.Context, key, value string) error

	Get(ctx context.Context, key string) (string, error)
	GetALL(ctx context.Context) (map[string]string, error)
}

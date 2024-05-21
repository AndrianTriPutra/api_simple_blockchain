package badGer

import (
	"context"
	"fmt"

	"github.com/dgraph-io/badger/v3"
)

type badgerDB struct {
	db *badger.DB
}

type DatabaseI interface {
	Db(ctx context.Context) interface{}
}

func NewConnection(pathToDb string) (DatabaseI, error) {
	opts := badger.DefaultOptions(pathToDb)

	opts.Logger = nil
	db, err := badger.Open(opts)
	if err != nil {
		return nil, fmt.Errorf("opening kv: %w", err)
	}

	return &badgerDB{db: db}, nil
}

func (k *badgerDB) Db(ctx context.Context) interface{} {
	tx := ctx.Value("txContext")
	if tx == nil {
		return k.db
	}
	return tx.(*badger.DB)
}

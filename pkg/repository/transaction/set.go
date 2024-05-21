package transaction

import (
	"context"

	"github.com/dgraph-io/badger/v3"
)

func (r repository) Set(ctx context.Context, key, value string) error {
	db := r.provider.Db(ctx).(*badger.DB)

	return db.Update(
		func(txn *badger.Txn) error {
			return txn.Set([]byte(key), []byte(value))
		})
}

func (r repository) SetTemp(ctx context.Context, value string) error {
	db := r.provider.Db(ctx).(*badger.DB)

	return db.Update(
		func(txn *badger.Txn) error {
			return txn.Set([]byte("temp"), []byte(value))
		})
}

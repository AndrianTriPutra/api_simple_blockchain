package transaction

import (
	"context"
	"fmt"

	"github.com/dgraph-io/badger/v3"
)

func (r repository) Get(ctx context.Context, key string) (string, error) {
	db := r.provider.Db(ctx).(*badger.DB)

	var value string
	return value, db.View(
		func(tx *badger.Txn) error {
			item, err := tx.Get([]byte(key))
			if err != nil {
				return fmt.Errorf("getting value: %w", err)
			}
			valCopy, err := item.ValueCopy(nil)
			if err != nil {
				return fmt.Errorf("copying value: %w", err)
			}
			value = string(valCopy)
			return nil
		})
}

func (r repository) GetALL(ctx context.Context) (map[string]string, error) {
	var data = map[string]string{}
	db := r.provider.Db(ctx).(*badger.DB)
	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				//fmt.Printf("key=%s, value=%s\n", k, v)
				data[string(k)] = string(v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	return data, err
}

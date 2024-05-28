package transaction

import (
	"atp/payment/pkg/adapter/levelDB"
	"context"
	"errors"

	"github.com/syndtr/goleveldb/leveldb"
)

type repository struct {
	provider levelDB.DatabaseI
}

func NewRepository(provider levelDB.DatabaseI) RepositoryI {
	return repository{
		provider,
	}
}

type RepositoryI interface {
	PUT(ctx context.Context, key, value string) error
	GET(ctx context.Context, key string) (string, error)
	GetALL(ctx context.Context) (map[string]string, error)
}

func (r repository) PUT(ctx context.Context, key, value string) error {
	db := r.provider.Db(ctx).(*leveldb.DB)
	err := db.Put([]byte(key), []byte(value), nil)
	if err != nil {
		errN := errors.New("failed when PUT key [" + key + "]->" + err.Error())
		return errN
	}
	return nil
}

func (r repository) GET(ctx context.Context, key string) (string, error) {
	var value string
	db := r.provider.Db(ctx).(*leveldb.DB)
	data, err := db.Get([]byte(key), nil)
	if err != nil {
		errN := errors.New("failed when GET key [" + key + "]->" + err.Error())
		return value, errN
	}
	value = string(data)
	return value, nil
}

func (r repository) GetALL(ctx context.Context) (map[string]string, error) {
	var data = map[string]string{}

	db := r.provider.Db(ctx).(*leveldb.DB)
	iter := db.NewIterator(nil, nil)
	for ok := iter.First(); ok; ok = iter.Next() {
		key := iter.Key()
		value := iter.Value()
		data[string(key)] = string(value)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		errN := errors.New("failed when GetALL ->" + err.Error())
		return data, errN
	}
	return data, nil
}

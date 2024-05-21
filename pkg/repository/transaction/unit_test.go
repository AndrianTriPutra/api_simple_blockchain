package transaction_test

import (
	"atp/payment/pkg/adapter/badGer"
	"atp/payment/pkg/repository/transaction"
	"context"
	"log"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func Test_SetTemp(t *testing.T) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	log.Printf("basepath:%s", basepath)
	base := basepath[0:strings.Index(basepath, "pkg")]
	path := base + "database/"
	log.Printf("path:%s", path)

	db, err := badGer.NewConnection(path)
	if err != nil {
		log.Fatalf("failed connect to database:" + err.Error())
	}
	log.Printf("succes connect to database:%s", path)

	repo := transaction.NewRepository(db)
	ctx := context.Background()

	err = repo.SetTemp(ctx, "init")
	if err != nil {
		log.Fatalf("failed SetTemp:" + err.Error())
	}
}

func Test_Set(t *testing.T) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	log.Printf("basepath:%s", basepath)
	base := basepath[0:strings.Index(basepath, "pkg")]
	path := base + "database/"
	log.Printf("path:%s", path)

	db, err := badGer.NewConnection(path)
	if err != nil {
		log.Fatalf("failed connect to database:" + err.Error())
	}
	log.Printf("succes connect to database:%s", path)

	repo := transaction.NewRepository(db)
	ctx := context.Background()

	err = repo.Set(ctx, "1", "a")
	if err != nil {
		log.Fatalf("failed Set:" + err.Error())
	}
}

func Test_Get(t *testing.T) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	log.Printf("basepath:%s", basepath)
	base := basepath[0:strings.Index(basepath, "pkg")]
	path := base + "database/"
	log.Printf("path:%s", path)

	db, err := badGer.NewConnection(path)
	if err != nil {
		log.Fatalf("failed connect to database:" + err.Error())
	}
	log.Printf("succes connect to database:%s", path)

	repo := transaction.NewRepository(db)
	ctx := context.Background()

	key := "temp"
	val, err := repo.Get(ctx, key)
	if err != nil {
		log.Fatalf("failed Get:" + err.Error())
	}
	log.Printf("[key:%s || value:%s]", key, val)

}

func Test_GetALL(t *testing.T) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	log.Printf("basepath:%s", basepath)
	base := basepath[0:strings.Index(basepath, "pkg")]
	path := base + "database/"
	log.Printf("path:%s", path)

	db, err := badGer.NewConnection(path)
	if err != nil {
		log.Fatalf("failed connect to database:" + err.Error())
	}
	log.Printf("succes connect to database:%s", path)

	repo := transaction.NewRepository(db)
	ctx := context.Background()

	data, err := repo.GetALL(ctx)
	if err != nil {
		log.Fatalf("failed GetALL:" + err.Error())
	}

	for key, val := range data {
		log.Printf("[key:%s || value:%s]", key, val)
	}
}

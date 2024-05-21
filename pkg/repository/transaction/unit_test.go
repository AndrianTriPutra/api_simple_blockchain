package transaction_test

import (
	"atp/payment/pkg/adapter/levelDB"
	"atp/payment/pkg/repository/transaction"
	"context"
	"log"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func Test_LastKey(t *testing.T) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	log.Printf("basepath:%s", basepath)
	base := basepath[0:strings.Index(basepath, "pkg")]
	path := base + "database/"
	log.Printf("path:%s", path)

	db, err := levelDB.NewConnection(path)
	if err != nil {
		log.Fatalf("failed connect to database:" + err.Error())
	}

	repo := transaction.NewRepository(db, path+"key.db")

	ctx := context.Background()
	last, err := repo.GET(ctx, "temp")
	if err != nil {
		log.Fatalf("failed LastKey:" + err.Error())
	}

	log.Printf("last:%s", last)

}

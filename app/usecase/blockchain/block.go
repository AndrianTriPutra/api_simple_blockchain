package blockchain

import (
	"atp/payment/pkg/utils/domain"
	"context"
	"encoding/json"
	"log"
	"time"
)

func (bc *blockchain) NewBlock(ctx context.Context, prevHash string, data []*domain.Data) *domain.Block {
	b := new(domain.Block)
	b.Data = data

	b.Header = &domain.Header{
		PrevHash: prevHash,
		Time:     time.Now().UnixNano(),
	}
	nowHash := b.Hash()

	prev, err := bc.transRepo.GET(ctx, "temp")
	if err != nil {
		log.Fatalf("FAILED NewBlock GET:" + err.Error())
	} else {
		log.Printf("key temp [%s] found", prev)
		b.Header.PrevHash = prev
	}

	js, _ := json.Marshal(b)
	err = bc.transRepo.PUT(ctx, nowHash, string(js))
	if err != nil {
		log.Fatalf("FAILED NewBlock PUT-data:" + err.Error())
	}

	err = bc.transRepo.PUT(ctx, "temp", nowHash)
	if err != nil {
		log.Fatalf("FAILED NewBlock PUT-temp:" + err.Error())
	}

	return b
}

func (bc *blockchain) CreateBlock(ctx context.Context, a *domain.Blockchain, prevHash string) *domain.Block {
	b := bc.NewBlock(ctx, prevHash, a.Pool)
	a.Chain = append(a.Chain, b)
	a.Pool = []*domain.Data{}
	return b
}

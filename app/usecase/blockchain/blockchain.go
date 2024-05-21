package blockchain

import (
	"atp/payment/pkg/utils/domain"
	"context"
	"encoding/json"
	"log"
	"time"
)

func (bc blockchain) CreateBlockchain(ctx context.Context, prevHash string) *domain.Blockchain {
	nbc := new(domain.Blockchain)
	bc.CreateGenesis(ctx, nbc, prevHash)
	return nbc
}

func (bc *blockchain) CreateGenesis(ctx context.Context, a *domain.Blockchain, prevHash string) *domain.Block {
	b := bc.NewGenesis(ctx, prevHash, a.Pool)
	a.Chain = append(a.Chain, b)
	a.Pool = []*domain.Data{}
	return b
}

func (bc *blockchain) NewGenesis(ctx context.Context, prevHash string, data []*domain.Data) *domain.Block {
	b := new(domain.Block)
	b.Data = data

	b.Header = &domain.Header{
		PrevHash: prevHash,
		Time:     time.Now().UnixNano(),
	}
	nowHash := b.Hash()

	exist, err := bc.transRepo.GetALL(ctx)
	if err != nil {
		log.Fatalf("FAILED NewGenesis GetALL:" + err.Error())
	}

	if len(exist) <= 0 {
		log.Println(" ==== Create GENESIS ====")
		js, _ := json.Marshal(b)
		err = bc.transRepo.Set(ctx, nowHash, string(js))
		if err != nil {
			log.Fatalf("FAILED NewGenesis SET-data:" + err.Error())
		}

		err = bc.transRepo.Set(ctx, "temp", nowHash)
		if err != nil {
			log.Fatalf("FAILED NewGenesis SET-temp:" + err.Error())
		}
	}

	return b
}

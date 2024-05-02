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

	ctrans := false
	prev, err := bc.transRepo.GetALL(ctx)
	if err != nil {
		log.Fatalf("FAILED GetALLX:" + err.Error())
	}

	if len(prev) > 0 {
		for _, val := range prev {
			block := domain.Block{}
			err = json.Unmarshal([]byte(val), &block)
			if err != nil {
				log.Fatalf("FAILED Unmarshalx:" + err.Error())
			}

			if block.Header.PrevHash == prevHash {
				log.Println(" prevHash is EXIST [" + prevHash + "]")
			} else {
				ctrans = true
			}
		}
	} else {
		ctrans = true
	}

	if ctrans {
		js, _ := json.Marshal(b)
		err := bc.transRepo.PUT(ctx, nowHash, string(js))
		if err != nil {
			log.Fatalf("FAILED NewBlock PUT:" + err.Error())
		}

		err = bc.transRepo.Save(ctx, nowHash)
		if err != nil {
			log.Fatalf("FAILED NewBlock Save key:" + err.Error())
		}
	}
	return b
}

func (bc *blockchain) CreateBlock(ctx context.Context, a *domain.Blockchain, prevHash string) *domain.Block {
	b := bc.NewBlock(ctx, prevHash, a.Pool)
	a.Chain = append(a.Chain, b)
	a.Pool = []*domain.Data{}
	return b
}

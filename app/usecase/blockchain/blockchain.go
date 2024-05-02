package blockchain

import (
	"atp/payment/pkg/utils/domain"
	"context"
)

func (bc blockchain) CreateBlockchain(ctx context.Context, prevHash string) *domain.Blockchain {
	nbc := new(domain.Blockchain)
	bc.CreateBlock(ctx, nbc, prevHash)
	return nbc
}

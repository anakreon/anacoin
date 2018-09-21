package blockchain

import (
	"github.com/anakreon/anacoin/pkg/hasher"
)

type Block struct {
	PreviousBlock *Block
	NextBlock     *Block
	Index         int64
	PreviousHash  string
	Hash          string
	Timestamp     int64
	Transactions  []Transaction
	Nonce         string
	Target        int
}

func (block Block) CalculateHash() string {
	hashData := string(block.Timestamp) + block.PreviousHash + block.Nonce + string(block.Target)
	return hasher.GetSha256Hash(hashData)
}

package blockchain

import (
	"time"

	"github.com/anakreon/anacoin/internal/hasher"
)

type Block struct {
	previousHash string
	hash         string
	timestamp    int64
	transactions []Transaction
	nonce        string
	target       int
}

func NewBlock(previousHash string, transactions []Transaction) *Block {
	return &Block{
		timestamp:    time.Now().Unix(),
		previousHash: previousHash,
		target:       5,
		transactions: transactions,
	}
}

func (block Block) CalculateHash() string {
	hashData := string(block.timestamp) + block.previousHash + block.nonce + string(block.target)
	return hasher.GetSha256Hash(hashData)
}

func (block *Block) CalculateAndSetHash() {
	block.hash = block.CalculateHash()
}

func (block Block) GetHash() string {
	return block.hash
}

func (block Block) GetTarget() int {
	return block.target
}

func (block Block) GetTransactions() []Transaction {
	return block.transactions
}

func (block *Block) SetNonce(nonce string) {
	block.nonce = nonce
}

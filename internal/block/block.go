package block

import (
	"time"

	"github.com/anakreon/anacoin/internal/hasher"
)

type Block struct {
	previousHash string
	timestamp    int64
	merkleRoot   string
	transactions []Transaction
	nonce        string
	target       int
}

func NewBlock(previousHash string, transactions []Transaction) *Block {
	return &Block{
		timestamp:    time.Now().Unix(),
		previousHash: previousHash,
		target:       4,
		transactions: transactions,
		merkleRoot:   calculateMerkleRoot(transactions),
	}
}

func (block Block) CalculateHash() string {
	hashData := string(block.timestamp) + block.previousHash + block.nonce + string(block.target) + block.merkleRoot
	return hasher.GetSha256Hash(hashData)
}

func (block Block) GetTarget() int {
	return block.target
}

func (block Block) GetPreviousHash() string {
	return block.previousHash
}

func (block Block) GetTransactions() []Transaction {
	return block.transactions
}

func (block *Block) SetNonce(nonce string) {
	block.nonce = nonce
}

func calculateMerkleRoot(transactions []Transaction) string {
	if len(transactions)%2 != 0 {
		transactions = append(transactions, transactions[len(transactions)-1])
	}
	return recursivelyCalculateMerkleRoot(transactions)
}

func recursivelyCalculateMerkleRoot(transactions []Transaction) string {
	if len(transactions) == 1 {
		return transactions[0].CalculateHash()
	} else {
		halfLength := len(transactions) / 2
		return recursivelyCalculateMerkleRoot(transactions[:halfLength]) + recursivelyCalculateMerkleRoot(transactions[halfLength:])
	}
}

package blockchain

import (
	"regexp"
	"strings"
)

type Block struct {
	Index        int64
	PreviousHash string
	Hash         string
	Timestamp    int64
	Transactions []Transaction
	Nonce        string
	Target       int
}

func (block Block) CalculateHash() string {
	hashData := string(block.Timestamp) + block.PreviousHash + block.Nonce
	return getSha256Hash(hashData)
}

func (block Block) Validate(previousBlock Block) bool {
	hasValidPreviousHash := previousBlock.Hash == block.PreviousHash
	hasValidCalculatedHash := block.CalculateHash() == block.Hash
	hasValidHashAsPerTarget := block.IsValidTargetHash()
	return hasValidPreviousHash && hasValidCalculatedHash && hasValidHashAsPerTarget && block.hasCoinbaseTransaction() && block.areTransactionsValid()
}

func (block Block) IsValidTargetHash() bool {
	prefix := strings.Repeat("8", block.Target)
	return strings.HasPrefix(block.Hash, prefix) && matchEndTargetCharacter(block)
}

func (block Block) hasCoinbaseTransaction() bool {
	return len(block.Transactions[0].In) == 1 &&
		block.Transactions[0].In[0].ScriptSig == "COINBASE" &&
		len(block.Transactions[0].Out) == 1 &&
		block.Transactions[0].Out[0].Value == 1
}

func (block Block) areTransactionsValid() bool {
	areValid := true
	for i := 1; i < len(block.Transactions); i++ {
		if !isTransactionValid() {
			areValid = false
			break
		}
	}
	return areValid
}

func isTransactionValid() bool {
	return true
}

func matchEndTargetCharacter(block Block) bool {
	character := string([]rune(block.Hash)[block.Target])
	match, _ := regexp.MatchString("[5-9]", character)
	return match
}

package validator

import (
	"regexp"
	"strings"

	"github.com/anakreon/anacoin/pkg/blockchain"
)

func IsValidBlock(block *blockchain.Block) bool {
	hasValidCalculatedHash := block.CalculateHash() == block.Hash
	hasValidHashAsPerTarget := IsValidHashAsPerTarget(block.Hash, block.Target)
	return hasValidCalculatedHash &&
		hasValidHashAsPerTarget &&
		hasCoinbaseTransaction(block) &&
		areTransactionsValid(block)
}

func hasCoinbaseTransaction(block *blockchain.Block) bool {
	return len(block.Transactions) > 0 &&
		len(block.Transactions[0].In) == 1 &&
		block.Transactions[0].In[0].ScriptSig == "COINBASE" &&
		len(block.Transactions[0].Out) == 1 &&
		block.Transactions[0].Out[0].Value == 1
}

func areTransactionsValid(block *blockchain.Block) bool {
	areValid := true
	for i := 1; i < len(block.Transactions); i++ {
		if !IsTransactionValid() {
			areValid = false
			break
		}
	}
	return areValid
}

func IsTransactionValid() bool {
	return true
}

func IsValidHashAsPerTarget(hash string, target int) bool {
	prefix := strings.Repeat("8", target)
	return strings.HasPrefix(hash, prefix) && restrictEndTargetCharacter(hash, target)
}

func restrictEndTargetCharacter(hash string, target int) bool {
	if len(hash) <= target {
		return false
	}
	character := string([]rune(hash)[target])
	match, _ := regexp.MatchString("[5-9]", character)
	return match
}

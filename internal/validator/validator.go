package validator

import (
	"regexp"
	"strings"

	"github.com/anakreon/anacoin/internal/blockchain"
)

func IsValidBlock(block *blockchain.Block) bool {
	hash := block.GetHash()
	hasValidCalculatedHash := block.CalculateHash() == hash
	hasValidHashAsPerTarget := IsValidHashAsPerTarget(hash, block.GetTarget())
	return hasValidCalculatedHash &&
		hasValidHashAsPerTarget &&
		hasCoinbaseTransaction(block) &&
		areTransactionsValid(block)
}

func hasCoinbaseTransaction(block *blockchain.Block) bool {
	transactions := block.GetTransactions()
	return len(transactions) > 0 &&
		len(transactions[0].In) == 1 &&
		transactions[0].In[0].ScriptSig == "COINBASE" &&
		len(transactions[0].Out) == 1 &&
		transactions[0].Out[0].Value == 1
}

func areTransactionsValid(block *blockchain.Block) bool {
	transactions := block.GetTransactions()
	areValid := true
	for i := 1; i < len(transactions); i++ {
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

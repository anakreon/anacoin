package validator

import (
	"regexp"
	"strings"

	"github.com/anakreon/anacoin/internal/block"
)

func IsValidBlock(block *block.Block) bool {
	hasValidHashAsPerTarget := IsValidHashAsPerTarget(block.CalculateHash(), block.GetTarget())
	return hasValidHashAsPerTarget &&
		hasCoinbaseTransaction(block) &&
		areTransactionsValid(block)
}

func hasCoinbaseTransaction(block *block.Block) bool {
	transactions := block.GetTransactions()
	return len(transactions) > 0 &&
		transactions[0].GetInputLength() == 1 &&
		transactions[0].GetInputs()[0].GetSignature() == "COINBASE" &&
		transactions[0].GetOutputLength() == 1 &&
		transactions[0].GetOutputs()[0].GetValue() == 1
}

func areTransactionsValid(block *block.Block) bool {
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

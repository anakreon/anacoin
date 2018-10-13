package blockchain

import (
	"strconv"

	"github.com/anakreon/anacoin/internal/hasher"
)

type TransactionOutput struct {
	value        uint64
	scriptPubKey string
}

func NewTransactionOutput(value uint64, scriptPubKey string) TransactionOutput {
	return TransactionOutput{value, scriptPubKey}
}

func (output TransactionOutput) CalculateHash() string {
	outValue := output.scriptPubKey + strconv.FormatUint(output.value, 16)
	return hasher.GetDoubleHashBase64(outValue)
}

func (output TransactionOutput) GetValue() uint64 {
	return output.value
}

func (output TransactionOutput) GetPubKey() string {
	return output.scriptPubKey
}

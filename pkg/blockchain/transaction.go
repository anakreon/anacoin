package blockchain

import (
	"strconv"

	"github.com/anakreon/anacoin/pkg/hasher"
)

type Transaction struct {
	In  []TransactionInput
	Out []TransactionOutput
}

type TransactionInput struct {
	TransactionID    string
	TransactionIndex uint8
	ScriptSig        string
}

type TransactionOutput struct {
	Value        float64
	ScriptPubKey string
}

func (transaction Transaction) CalculateHash() string {
	inHash, outHash := "", ""
	for _, input := range transaction.In {
		inValue := input.TransactionID + string(input.TransactionIndex) + inHash
		inHash = hasher.GetDoubleHashBase64(inValue)
	}

	for _, output := range transaction.Out {
		outValue := output.ScriptPubKey + strconv.FormatFloat(output.Value, 'f', 6, 64) + outHash
		outHash = hasher.GetDoubleHashBase64(outValue)
	}
	return hasher.GetDoubleHashBase64(inHash + outHash)
}

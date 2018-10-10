package blockchain

import (
	"strconv"

	"github.com/anakreon/anacoin/internal/hasher"
)

type Transaction struct {
	in  []TransactionInput
	out []TransactionOutput
}

func NewTransaction(inputs []TransactionInput, outputs []TransactionOutput) *Transaction {
	return &Transaction{
		in:  inputs,
		out: outputs,
	}
}

/*func (transaction Transaction) buildTransactionInputFromOutput(transactionOutputIndex uint8, scriptSig string) TransactionInput {
	transactionID := transaction.CalculateHash()
	return newTransactionInput(transactionID, transactionOutputIndex, scriptSig)
}*/

func (transaction Transaction) CalculateHash() string {
	inHash, outHash := "", ""
	for _, input := range transaction.in {
		inValue := input.transactionID + string(input.transactionIndex) + inHash
		inHash = hasher.GetDoubleHashBase64(inValue)
	}

	for _, output := range transaction.out {
		outValue := output.scriptPubKey + strconv.FormatUint(output.value, 64) + outHash
		outHash = hasher.GetDoubleHashBase64(outValue)
	}
	return hasher.GetDoubleHashBase64(inHash + outHash)
}

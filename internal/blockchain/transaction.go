package blockchain

import (
	"github.com/anakreon/anacoin/internal/hasher"
)

type Transaction struct {
	in  []TransactionInput
	out []TransactionOutput
}

type hashable interface {
	CalculateHash() string
}

func NewTransaction(in []TransactionInput, out []TransactionOutput) Transaction {
	return Transaction{in, out}
}

func (transaction Transaction) CalculateHash() string {
	inHash, outHash := "", ""
	for _, value := range transaction.in {
		inHash = calculateHashableValueHash(value, inHash)
	}
	for _, value := range transaction.out {
		outHash = calculateHashableValueHash(value, inHash)
	}
	return hasher.GetDoubleHashBase64(inHash + outHash)
}

func calculateHashableValueHash(value hashable, hash string) string {
	loopHash := value.CalculateHash() + hash
	return hasher.GetDoubleHashBase64(loopHash)
}

func (transaction *Transaction) GetInputLength() int {
	return len(transaction.in)
}

func (transaction *Transaction) GetOutputLength() int {
	return len(transaction.out)
}

func (transaction *Transaction) GetInputs() []TransactionInput {
	return transaction.in
}

func (transaction *Transaction) GetOutputs() []TransactionOutput {
	return transaction.out
}

func (transaction *Transaction) SetSignatureForInputs(signature string) {
	for _, input := range transaction.in {
		input.SetSignature(signature)
	}
}

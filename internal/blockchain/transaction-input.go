package blockchain

import "github.com/anakreon/anacoin/internal/hasher"

type TransactionInput struct {
	transactionID    string
	transactionIndex int
	scriptSig        string
}

func NewTransactionInput(transactionID string, transactionIndex int, scriptSig string) TransactionInput {
	return TransactionInput{transactionID, transactionIndex, scriptSig}
}

func (input TransactionInput) CalculateHash() string {
	inValue := input.transactionID + string(input.transactionIndex)
	return hasher.GetDoubleHashBase64(inValue)
}

func (input TransactionInput) GetSignature() string {
	return input.scriptSig
}

func (input *TransactionInput) SetSignature(signature string) {
	input.scriptSig = signature
}

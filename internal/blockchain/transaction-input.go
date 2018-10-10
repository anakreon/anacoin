package blockchain

type TransactionInput struct {
	transactionID    string
	transactionIndex uint8
	scriptSig        string
}

func newTransactionInput(transactionID string, transactionIndex uint8, scriptSig string) TransactionInput {
	return TransactionInput{
		transactionID:    transactionID,
		transactionIndex: transactionIndex,
		scriptSig:        scriptSig,
	}
}

package blockchain

type TransactionOutput struct {
	value        uint64
	scriptPubKey string
}

func newTransactionOutput(value uint64, scriptPubKey string) TransactionOutput {
	return TransactionOutput{
		value:        value,
		scriptPubKey: scriptPubKey,
	}
}

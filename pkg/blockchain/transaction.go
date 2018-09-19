package blockchain

type Transaction struct {
	In  TransactionInput
	Out []TransactionOutput
}

type TransactionInput struct {
	TxID      string
	TxIndex   int8
	ScriptSig string
}

type TransactionOutput struct {
	Value        float64
	ScriptPubKey string
}

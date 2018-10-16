package block

type transactionID string
type transactionIndexOutputMap map[int]TransactionOutput
type UnspentTransactionOutputs map[transactionID]transactionIndexOutputMap

type correctTransactionOutputCallback func(transactionOutput TransactionOutput) bool

func (unspentTransactionOutputs UnspentTransactionOutputs) FilterUnspentTransactionOutputs(isCorrectTransactionOutput correctTransactionOutputCallback) UnspentTransactionOutputs {
	filteredUnspentTransactionOutputs := make(UnspentTransactionOutputs)
	for transactionID, indexOutputMap := range unspentTransactionOutputs {
		for index, output := range indexOutputMap {
			if isCorrectTransactionOutput(output) {
				if _, filteredKeyExists := filteredUnspentTransactionOutputs[transactionID]; !filteredKeyExists {
					filteredUnspentTransactionOutputs[transactionID] = make(transactionIndexOutputMap)
				}
				filteredUnspentTransactionOutputs[transactionID][index] = output
			}
		}
	}
	return filteredUnspentTransactionOutputs
}

func (unspentTransactionOutputs UnspentTransactionOutputs) UpdateFromTransactions(transactions []Transaction) {
	for _, transaction := range transactions {
		transactionID := transactionID(transaction.CalculateHash())
		unspentTransactionOutputs.addTransactionOutputs(transactionID, transaction.out)
		unspentTransactionOutputs.removeTransactionInputs(transaction.in)
	}
}

func (unspentTransactionOutputs UnspentTransactionOutputs) addTransactionOutputs(transactionID transactionID, transactionOutputs []TransactionOutput) {
	unspentTransactionOutputs[transactionID] = make(transactionIndexOutputMap)
	for index, output := range transactionOutputs {
		unspentTransactionOutputs[transactionID][index] = output
	}
}

func (unspentTransactionOutputs UnspentTransactionOutputs) removeTransactionInputs(transactionInputs []TransactionInput) {
	for _, input := range transactionInputs {
		transactionID := transactionID(input.transactionID)
		delete(unspentTransactionOutputs[transactionID], input.transactionIndex)
		if len(unspentTransactionOutputs[transactionID]) == 0 {
			delete(unspentTransactionOutputs, transactionID)
		}
	}
}

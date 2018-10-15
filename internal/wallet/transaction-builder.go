package wallet

import "github.com/anakreon/anacoin/internal/blockchain"

type NotEnoughValueError struct{}

func (error NotEnoughValueError) Error() string {
	return "not enough value"
}

func (wallet *Wallet) buildTransactionInputsForValue(value uint64) ([]blockchain.TransactionInput, uint64, error) {
	result := []blockchain.TransactionInput{}
	valueToAdd := value
	myUnspentTransactionOutputs := wallet.getMyUnspentTransactionOutputs()
	for transactionID, indexOutputMap := range myUnspentTransactionOutputs {
		for index, output := range indexOutputMap {
			transactionInput := blockchain.NewTransactionInput(string(transactionID), index)
			result = append(result, transactionInput)
			valueToAdd -= output.GetValue()
			if valueToAdd <= 0 {
				break
			}
		}
	}
	if valueToAdd <= 0 {
		return result, -valueToAdd, nil
	} else {
		return nil, 0, NotEnoughValueError{}
	}
}

func (wallet *Wallet) buildTransactionOutputs(value uint64, remainderValue uint64, targetAddress string) []blockchain.TransactionOutput {
	result := []blockchain.TransactionOutput{
		blockchain.NewTransactionOutput(value, targetAddress),
	}
	if remainderValue != 0 {
		remainderOutput := blockchain.NewTransactionOutput(remainderValue, wallet.GetPublicAddress())
		result = append(result, remainderOutput)
	}
	return result
}

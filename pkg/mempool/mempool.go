package mempool

import "github.com/anakreon/anacoin/pkg/blockchain"

var unconfirmedTransactions = []blockchain.Transaction{}

func AddTransaction(transaction blockchain.Transaction) {
	unconfirmedTransactions = append(unconfirmedTransactions, transaction)
}

func GetAllTransactions() []blockchain.Transaction {
	return unconfirmedTransactions
}

package mempool

import "github.com/anakreon/anacoin/internal/blockchain"

type UnconfirmedTransactions []blockchain.Transaction

func (transactions *UnconfirmedTransactions) AddTransaction(transaction blockchain.Transaction) {
	*transactions = append(*transactions, transaction)
}

func (transactions *UnconfirmedTransactions) Clear() {
	*transactions = []blockchain.Transaction{}
}

func NewUnconfirmedTransactions() UnconfirmedTransactions {
	return []blockchain.Transaction{}
}

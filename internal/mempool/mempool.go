package mempool

import (
	"sync"

	"github.com/anakreon/anacoin/internal/blockchain"
)

type UnconfirmedTransactions []blockchain.Transaction

var mutex *sync.Mutex

func (transactions *UnconfirmedTransactions) AddTransaction(transaction blockchain.Transaction) {
	mutex.Lock()
	*transactions = append(*transactions, transaction)
	mutex.Unlock()
}

func (transactions *UnconfirmedTransactions) Clear() {
	mutex.Lock()
	*transactions = []blockchain.Transaction{}
	mutex.Unlock()
}

func NewUnconfirmedTransactions() UnconfirmedTransactions {
	mutex = &sync.Mutex{}
	return []blockchain.Transaction{}
}

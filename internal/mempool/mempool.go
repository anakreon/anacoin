package mempool

import (
	"sync"

	"github.com/anakreon/anacoin/internal/block"
)

type UnconfirmedTransactions []block.Transaction

var mutex *sync.Mutex

func (transactions *UnconfirmedTransactions) AddTransaction(transaction block.Transaction) {
	mutex.Lock()
	*transactions = append(*transactions, transaction)
	mutex.Unlock()
}

func (transactions *UnconfirmedTransactions) Clear() {
	mutex.Lock()
	*transactions = []block.Transaction{}
	mutex.Unlock()
}

func NewUnconfirmedTransactions() UnconfirmedTransactions {
	mutex = &sync.Mutex{}
	return []block.Transaction{}
}

package mempool

import (
	"testing"

	"github.com/anakreon/anacoin/internal/blockchain"
)

func TestNewUnconfirmedTransactions(t *testing.T) {
	unconfirmedTransactions := NewUnconfirmedTransactions()
	if len(unconfirmedTransactions) != 0 {
		t.Fail()
	}
}
func TestAddTransaction(t *testing.T) {
	newTransactionOne := blockchain.Transaction{}
	newTransactionTwo := blockchain.Transaction{}
	unconfirmedTransactions := NewUnconfirmedTransactions()
	unconfirmedTransactions.AddTransaction(newTransactionOne)
	unconfirmedTransactions.AddTransaction(newTransactionTwo)
	if len(unconfirmedTransactions) != 2 {
		t.Error("not 2 transactions")
	}
	if &unconfirmedTransactions[0] != &newTransactionOne {
		t.Error("transaction 1 not in slot 0")
	}
	if &unconfirmedTransactions[1] != &newTransactionTwo {
		t.Error("transaction 2 not in slot 1")
	}
}

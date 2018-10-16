package mempool

import (
	"testing"

	"github.com/anakreon/anacoin/internal/block"
)

func TestNewUnconfirmedTransactions(t *testing.T) {
	unconfirmedTransactions := NewUnconfirmedTransactions()
	if len(unconfirmedTransactions) != 0 {
		t.Fail()
	}
}
func TestAddTransaction(t *testing.T) {
	newTransactionOne := buildTransactionWithId("one")
	newTransactionTwo := buildTransactionWithId("two")
	unconfirmedTransactions := NewUnconfirmedTransactions()
	unconfirmedTransactions.AddTransaction(newTransactionOne)
	unconfirmedTransactions.AddTransaction(newTransactionTwo)
	if len(unconfirmedTransactions) != 2 {
		t.Error("not 2 transactions")
	}
	if len(unconfirmedTransactions[0].In) == 1 && unconfirmedTransactions[0].In[0].TransactionID != "one" {
		t.Error("transaction 1 not in slot 0")
	}
	if len(unconfirmedTransactions[1].In) == 1 && unconfirmedTransactions[1].In[0].TransactionID != "two" {
		t.Error("transaction 2 not in slot 1")
	}
}

func TestClear(t *testing.T) {
	newTransactionOne := buildTransactionWithId("one")
	newTransactionTwo := buildTransactionWithId("two")
	unconfirmedTransactions := NewUnconfirmedTransactions()
	unconfirmedTransactions.AddTransaction(newTransactionOne)
	unconfirmedTransactions.AddTransaction(newTransactionTwo)
	unconfirmedTransactions.Clear()
	if len(unconfirmedTransactions) != 0 {
		t.Error("not 0 transactions")
	}
}

func buildTransactionWithId(id string) block.Transaction {
	return block.Transaction{
		In: []block.TransactionInput{
			block.TransactionInput{TransactionID: id},
		},
	}
}

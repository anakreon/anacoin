package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"

	"github.com/anakreon/anacoin/internal/blockchain"
)

func (wallet *Wallet) createTransaction(value uint64) blockchain.Transaction {
	lastBlock := wallet.storage.GetLastBlock()
	sourceTransaction := lastBlock.Transactions[0]
	return wallet.buildTransaction(sourceTransaction.CalculateHash())
}

func (wallet *Wallet) buildTransaction(txid string) blockchain.Transaction {
	transaction := blockchain.Transaction{
		In: []blockchain.TransactionInput{
			blockchain.TransactionInput{
				TransactionID:    txid,
				TransactionIndex: 0,
			},
		},
		Out: []blockchain.TransactionOutput{
			blockchain.TransactionOutput{
				Value:        1000,
				ScriptPubKey: wallet.GetPublicAddress(),
			},
		},
	}
	transaction.In[0].ScriptSig = wallet.buildSignature(transaction.CalculateHash())
	return transaction
}

func (wallet *Wallet) buildSignature(transactionHash string) string {
	signedHash := wallet.signTransactionHash(transactionHash)
	publicKeyString := wallet.getPublicKeyString()
	return signedHash + " " + publicKeyString
}

func (wallet *Wallet) signTransactionHash(transactionHash string) string {
	r, s, _ := ecdsa.Sign(rand.Reader, wallet.privateKey, []byte(transactionHash))
	return r.Text(16) + "," + s.Text(16)
}

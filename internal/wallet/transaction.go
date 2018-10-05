package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"

	"github.com/anakreon/anacoin/internal/blockchain"
)

func (wallet *Wallet) createTransaction(targetAddress string, value uint64) blockchain.Transaction {
	//lastBlock := wallet.storage.GetLastBlock()
	//sourceTransaction := lastBlock.Transactions[0], sourceTransaction.CalculateHash()
	return wallet.buildTransaction(targetAddress, value, "some hash")
}

func (wallet *Wallet) buildTransaction(targetAddress string, value uint64, txid string) blockchain.Transaction {
	transaction := blockchain.Transaction{
		In: []blockchain.TransactionInput{
			blockchain.TransactionInput{
				TransactionID:    txid,
				TransactionIndex: 0,
			},
		},
		Out: []blockchain.TransactionOutput{
			blockchain.TransactionOutput{
				Value:        value,
				ScriptPubKey: targetAddress,
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

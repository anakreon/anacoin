package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/anakreon/anacoin/pkg/blockchain"
	"github.com/anakreon/anacoin/pkg/connector"
	"github.com/anakreon/anacoin/pkg/hasher"
	"github.com/anakreon/anacoin/pkg/mempool"
)

var privateKey *ecdsa.PrivateKey
var storage *blockchain.Blockchain

func Initialize(storageInstance *blockchain.Blockchain) {
	storage = storageInstance
	generateKeys()
}

func generateKeys() {
	curve := elliptic.P521()
	privateKeyPointer, _ := ecdsa.GenerateKey(curve, rand.Reader)
	privateKey = privateKeyPointer
}

func GetPublicAddress() string {
	publicKey := getPublicKeyString()
	base58Address := hasher.GetDoubleHashBase64(publicKey)
	return base58Address
}

func getPublicKeyString() string {
	return privateKey.PublicKey.X.Text(16) + "," + privateKey.PublicKey.Y.Text(16)
}

func TestSign() {
	value := "hello world"
	valueHash := sha256.Sum256([]byte(value))
	fmt.Println("hash", valueHash)
	r, s, _ := ecdsa.Sign(rand.Reader, privateKey, valueHash[:])
	fmt.Printf("signature: (0x%x, 0x%x)\n", r, s)
	valid := ecdsa.Verify(&privateKey.PublicKey, valueHash[:], r, s)
	fmt.Println("signature verified:", valid)
}

func AddTransactionX() {
	block := storage.GetLastBlock()
	sourceTransaction := block.Transactions[0]
	transaction := buildTransaction(sourceTransaction.CalculateHash())
	mempool.AddTransaction(transaction)
	connector.BroadcastNewTransaction(transaction)
}

func buildTransaction(txid string) blockchain.Transaction {
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
				ScriptPubKey: GetPublicAddress(),
			},
		},
	}
	transaction.In[0].ScriptSig = buildSignature(transaction.CalculateHash())
	return transaction
}

func buildSignature(transactionHash string) string {
	signedHash := signTransactionHash(transactionHash)
	publicKeyString := getPublicKeyString()
	return signedHash + " " + publicKeyString
}

func signTransactionHash(transactionHash string) string {
	r, s, _ := ecdsa.Sign(rand.Reader, privateKey, []byte(transactionHash))
	return r.Text(16) + "," + s.Text(16)
}

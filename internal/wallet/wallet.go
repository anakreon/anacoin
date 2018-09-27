package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/anakreon/anacoin/internal/blockchain"
	"github.com/anakreon/anacoin/internal/connector"
	"github.com/anakreon/anacoin/internal/hasher"
	"github.com/anakreon/anacoin/internal/mempool"
)

type Wallet struct {
	storage                 *blockchain.Blockchain
	unconfirmedTransactions *mempool.UnconfirmedTransactions
	connector               *connector.Connector
	privateKey              *ecdsa.PrivateKey
}

func NewWallet(storage *blockchain.Blockchain, unconfirmedTransactions *mempool.UnconfirmedTransactions, connector *connector.Connector) Wallet {
	return Wallet{
		privateKey:              generatePrivateKey(),
		storage:                 storage,
		unconfirmedTransactions: unconfirmedTransactions,
		connector:               connector,
	}
}

func generatePrivateKey() *ecdsa.PrivateKey {
	curve := elliptic.P521()
	privateKey, _ := ecdsa.GenerateKey(curve, rand.Reader)
	return privateKey
}

func (wallet *Wallet) GetPublicAddress() string {
	publicKey := wallet.getPublicKeyString()
	base58Address := hasher.GetDoubleHashBase64(publicKey)
	return base58Address
}

func (wallet *Wallet) getPublicKeyString() string {
	return wallet.privateKey.PublicKey.X.Text(16) + "," + wallet.privateKey.PublicKey.Y.Text(16)
}

func (wallet *Wallet) AddNewTransaction(value uint64) {
	transaction := wallet.createTransaction(value)
	wallet.unconfirmedTransactions.AddTransaction(transaction)
	wallet.connector.BroadcastNewTransaction(transaction)
}

/*func TestSign() {
	value := "hello world"
	valueHash := sha256.Sum256([]byte(value))
	fmt.Println("hash", valueHash)
	r, s, _ := ecdsa.Sign(rand.Reader, privateKey, valueHash[:])
	fmt.Printf("signature: (0x%x, 0x%x)\n", r, s)
	valid := ecdsa.Verify(&privateKey.PublicKey, valueHash[:], r, s)
	fmt.Println("signature verified:", valid)
}*/

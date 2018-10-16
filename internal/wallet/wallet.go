package wallet

import (
	"crypto/ecdsa"

	"github.com/anakreon/anacoin/internal/block"
	"github.com/anakreon/anacoin/internal/blockchain"
	"github.com/anakreon/anacoin/internal/hasher"
)

type connector interface {
	BroadcastNewTransaction(transaction block.Transaction)
}

type unconfirmedTransactions interface {
	AddTransaction(transaction block.Transaction)
}

type Wallet struct {
	storage                 *blockchain.Blockchain
	unconfirmedTransactions unconfirmedTransactions
	connector               connector
	privateKey              *ecdsa.PrivateKey
}

func NewWallet(storage *blockchain.Blockchain, unconfirmedTransactions unconfirmedTransactions, connector connector) Wallet {
	return Wallet{
		privateKey:              generatePrivateKey(),
		storage:                 storage,
		unconfirmedTransactions: unconfirmedTransactions,
		connector:               connector,
	}
}

func (wallet *Wallet) GetPublicAddress() string {
	publicKey := wallet.getPublicKeyString()
	base58Address := hasher.GetDoubleHashBase64(publicKey)
	return base58Address
}

func (wallet *Wallet) getPublicKeyString() string {
	return convertRSBigIntToText(wallet.privateKey.PublicKey.X, wallet.privateKey.PublicKey.Y)
}

func (wallet *Wallet) AddTransaction(targetAddress string, value uint64) {
	transaction := wallet.createTransaction(targetAddress, value)
	wallet.unconfirmedTransactions.AddTransaction(transaction)
	wallet.connector.BroadcastNewTransaction(transaction)
}

func (wallet *Wallet) GetBalance() uint64 {
	return wallet.getBalance()
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

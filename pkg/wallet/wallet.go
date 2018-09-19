package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/anakreon/anacoin/pkg/blockchain"
	"github.com/anakreon/anacoin/pkg/mempool"
	"golang.org/x/crypto/ripemd160"
)

var privateKey *ecdsa.PrivateKey

func Initialize() {
	generateKeys()
	addTransactionX()
}

func generateKeys() {
	curve := elliptic.P521()
	privateKeyPointer, _ := ecdsa.GenerateKey(curve, rand.Reader)
	privateKey = privateKeyPointer
}

func GetPublicAddress() string {
	publicKey := getPublicKeyString()
	base58Address := getDoubleHashBase64Key(publicKey)
	return base58Address
}

func getDoubleHashBase64Key(key string) string {
	sha256Hash := getSha256Hash(key)
	ripemd160Hash := getRipemd160Hash(sha256Hash)
	return getBase64String(ripemd160Hash)
}

func getSha256Hash(inputString string) string {
	hasher := sha256.New()
	hasher.Write([]byte(inputString))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes[:])
}

func getRipemd160Hash(inputString string) string {
	hasher := ripemd160.New()
	hasher.Write([]byte(inputString))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes[:])
}

func getBase64String(inputString string) string {
	return base64.StdEncoding.EncodeToString([]byte(inputString))
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

func addTransactionX() {
	transaction := buildTransaction()
	mempool.AddTransaction(transaction)
}

func buildTransaction() blockchain.Transaction {
	out := []blockchain.TransactionOutput{
		blockchain.TransactionOutput{
			Value:        1000,
			ScriptPubKey: GetPublicAddress(),
		},
	}
	return blockchain.Transaction{
		In: blockchain.TransactionInput{
			ScriptSig: buildSignature(out),
		},
		Out: out,
	}
}

func buildSignature(out []blockchain.TransactionOutput) string {
	signedHash := signTransactionHash(out)
	publicKeyString := getPublicKeyString()
	return signedHash + " " + publicKeyString
}

func signTransactionHash(out []blockchain.TransactionOutput) string {
	hasher := sha256.New()
	for _, transaction := range out {
		value := strconv.FormatFloat(transaction.Value, 'f', 6, 64) + transaction.ScriptPubKey
		hasher.Write([]byte(value))
	}
	hashBytes := hasher.Sum(nil)
	r, s, _ := ecdsa.Sign(rand.Reader, privateKey, hashBytes[:])
	return r.Text(16) + "," + s.Text(16)
}

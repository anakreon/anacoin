package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
	"strings"

	"github.com/anakreon/anacoin/internal/blockchain"
)

func (wallet *Wallet) getMyUnspentTransactionOutputs() blockchain.UnspentTransactionOutputs {
	unspentTransactionOutputs := wallet.storage.FindUnspentTransactionOutputs()
	return unspentTransactionOutputs.FilterUnspentTransactionOutputs(func(output blockchain.TransactionOutput) bool {
		return output.GetPubKey() == wallet.GetPublicAddress()
	})
}

func (wallet *Wallet) getBalance() uint64 {
	var balance uint64 = 0
	myUnspentTransactionOutputs := wallet.getMyUnspentTransactionOutputs()
	for _, indexOutputMap := range myUnspentTransactionOutputs {
		for _, output := range indexOutputMap {
			balance += output.GetValue()
		}
	}
	return balance
}

func (wallet *Wallet) createTransaction(targetAddress string, value uint64) blockchain.Transaction {
	return wallet.buildTransaction(targetAddress, value)
}

func (wallet *Wallet) buildTransaction(targetAddress string, value uint64) blockchain.Transaction {
	inputs, remainderValue, _ := wallet.buildTransactionInputsForValue(value)
	outputs := wallet.buildTransactionOutputs(value, remainderValue, targetAddress)
	transaction := blockchain.NewTransaction(inputs, outputs)
	signature := wallet.buildSignature(transaction.CalculateHash())
	transaction.SetSignatureForInputs(signature)
	return transaction
}

func (wallet *Wallet) buildSignature(transactionHash string) string {
	signedHash := wallet.signTransactionHash(transactionHash)
	publicKeyString := wallet.getPublicKeyString()
	return signedHash + " " + publicKeyString
}

func (wallet *Wallet) signTransactionHash(transactionHash string) string {
	r, s, _ := ecdsa.Sign(rand.Reader, wallet.privateKey, []byte(transactionHash))
	return convertRSBigIntToText(r, s)
}

func signTransactionHash(privateKey *ecdsa.PrivateKey, transactionHash string) string {
	r, s, _ := ecdsa.Sign(rand.Reader, privateKey, []byte(transactionHash))
	return convertRSBigIntToText(r, s)
}

func convertRSBigIntToText(r, s *big.Int) string {
	rText, _ := r.MarshalText()
	sText, _ := s.MarshalText()
	return string(rText) + "," + string(sText)
}

func isValidSignature(scriptSig string, scriptPubKey string, transactionHash string) bool {
	split := strings.Split(scriptSig, " ")
	signedHash := split[0]
	publicKeyString := split[1]

	r, s := convertRSTextToBigInt(signedHash)
	publicKeyX, publicKeyY := convertRSTextToBigInt(publicKeyString)
	publicKey := buildPublicKey(publicKeyX, publicKeyY)
	isValid := ecdsa.Verify(&publicKey, []byte(transactionHash), r, s)
	return isValid
}

func convertRSTextToBigInt(rs string) (r, s *big.Int) {
	split := strings.Split(rs, ",")
	r.UnmarshalText([]byte(split[0]))
	s.UnmarshalText([]byte(split[1]))
	return
}

func buildPublicKey(x, y *big.Int) ecdsa.PublicKey {
	curve := elliptic.P521()
	return ecdsa.PublicKey{curve, x, y}
}

func generatePrivateKey() *ecdsa.PrivateKey {
	curve := elliptic.P521()
	privateKey, _ := ecdsa.GenerateKey(curve, rand.Reader)
	return privateKey
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

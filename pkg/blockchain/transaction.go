package blockchain

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strconv"

	"golang.org/x/crypto/ripemd160"
)

type Transaction struct {
	In  []TransactionInput
	Out []TransactionOutput
}

type TransactionInput struct {
	TxID      string
	TxIndex   int8
	ScriptSig string
}

type TransactionOutput struct {
	Value        float64
	ScriptPubKey string
}

func (transaction Transaction) CalculateHash() string {
	inHash, outHash := "", ""
	for _, input := range transaction.In {
		inValue := input.TxID + string(input.TxIndex) + inHash
		inHash = getDoubleHashBase64Key(inValue)
	}

	for _, output := range transaction.Out {
		outValue := output.ScriptPubKey + strconv.FormatFloat(output.Value, 'f', 6, 64) + outHash
		outHash = getDoubleHashBase64Key(outValue)
	}
	return getDoubleHashBase64Key(inHash + outHash)
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

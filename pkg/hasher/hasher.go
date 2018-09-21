package hasher

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	"golang.org/x/crypto/ripemd160"
)

func GetDoubleHashBase64(input string) string {
	sha256Hash := GetSha256Hash(input)
	ripemd160Hash := GetRipemd160Hash(sha256Hash)
	return GetBase64String(ripemd160Hash)
}

func GetSha256Hash(inputString string) string {
	hasher := sha256.New()
	hasher.Write([]byte(inputString))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes[:])
}

func GetRipemd160Hash(inputString string) string {
	hasher := ripemd160.New()
	hasher.Write([]byte(inputString))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes[:])
}

func GetBase64String(inputString string) string {
	return base64.StdEncoding.EncodeToString([]byte(inputString))
}

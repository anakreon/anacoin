package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"regexp"
	"strings"
)

type Block struct {
	PreviousHash string
	Hash         string
	Timestamp    int64
	Transactions []Transaction
	Nonce        string
	Target       int
}

func (block Block) CalculateHash() string {
	hashData := string(block.Timestamp) + block.PreviousHash + block.Nonce
	hash := sha256.New()
	hash.Write([]byte(hashData))
	byteHash := hash.Sum(nil)
	return hex.EncodeToString(byteHash)
}

func (block Block) Validate(previousBlock Block) bool {
	hasValidPreviousHash := previousBlock.Hash == block.PreviousHash
	hasValidCalculatedHash := block.CalculateHash() == block.Hash
	hasValidHasAsPerTarget := block.IsValidTargetHash()
	return hasValidPreviousHash && hasValidCalculatedHash && hasValidHasAsPerTarget
}

func (block Block) IsValidTargetHash() bool {
	prefix := strings.Repeat("8", block.Target)
	return strings.HasPrefix(block.Hash, prefix) && matchEndTargetCharacter(block)
}

func matchEndTargetCharacter(block Block) bool {
	character := string([]rune(block.Hash)[block.Target])
	match, _ := regexp.MatchString("[5-9]", character)
	return match
}

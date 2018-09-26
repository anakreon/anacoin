package validator

import (
	"testing"
)

func TestIsValidBlock(t *testing.T) {
	validBlock := buildValidBlock()
	t.Run("is valid", testIsValidBlock(&validBlock, true))

	invalidNonce := buildBlockInvalidNonce()
	t.Run("invalid nonce", testIsValidBlock(&invalidNonce, false))

	invalidHash := buildBlockInvalidHash()
	t.Run("invalid hash", testIsValidBlock(&invalidHash, false))

	noCoinbaseTransaction := buildBlockNoCoinbaseTransaction()
	t.Run("no coinbase transaction", testIsValidBlock(&noCoinbaseTransaction, false))
}

func buildBlockInvalidNonce() blockchain.Block {
	block := buildValidBlock()
	block.Nonce = "not a valid nonce"
	return block
}

func buildBlockInvalidHash() blockchain.Block {
	block := buildValidBlock()
	block.Hash = "not a valid hash"
	return block
}

func buildBlockNoCoinbaseTransaction() blockchain.Block {
	block := buildValidBlock()
	block.Transactions = []blockchain.Transaction{}
	return block
}

func buildValidBlock() blockchain.Block {
	return blockchain.Block{
		PreviousHash: "PreviousHash",
		Timestamp:    1234,
		Target:       3,
		Nonce:        "2bbc74fc8b213512",
		Hash:         "8888fbf1517a4b6a55a6972da919d1bd8d8a3ee3b5e99f3387d61cac95988389",
		Transactions: []blockchain.Transaction{
			blockchain.Transaction{
				In: []blockchain.TransactionInput{
					blockchain.TransactionInput{
						ScriptSig: "COINBASE",
					},
				},
				Out: []blockchain.TransactionOutput{
					blockchain.TransactionOutput{
						Value:        1,
						ScriptPubKey: "asdf",
					},
				},
			},
		},
	}
}

func testIsValidBlock(block *blockchain.Block, expected bool) func(*testing.T) {
	return func(t *testing.T) {
		if IsValidBlock(block) != expected {
			t.Fail()
		}
	}
}

func TestIsTransactionValid(t *testing.T) {
	if !IsTransactionValid() {
		t.Fail()
	}
}

func TestIsValidHashAsPerTarget(t *testing.T) {
	t.Run("is valid, valid prefix & endChar", testIsValidHashAsPerTarget("88885", 4, true))
	t.Run("is valid, no target, only endChar", testIsValidHashAsPerTarget("5", 0, true))
	t.Run("is valid, target 1 & endChar", testIsValidHashAsPerTarget("86", 1, true))
	t.Run("is not valid, invalid prefix", testIsValidHashAsPerTarget("888aa", 4, false))
	t.Run("is not valid, invalid endChar", testIsValidHashAsPerTarget("8888a", 4, false))
	t.Run("is not valid", testIsValidHashAsPerTarget("asdfadsf", 4, false))
}

func testIsValidHashAsPerTarget(hash string, target int, expected bool) func(*testing.T) {
	return func(t *testing.T) {
		if IsValidHashAsPerTarget(hash, target) != expected {
			t.Fail()
		}
	}
}

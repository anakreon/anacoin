package validator

import (
	"testing"

	"github.com/anakreon/anacoin/internal/block"
)

/*
func TestIsValidBlock(t *testing.T) {
	validBlock := buildValidBlock()
	invalidNonce := buildBlockInvalidNonce()
	invalidHash := buildBlockInvalidHash()
	noCoinbaseTransaction := buildBlockNoCoinbaseTransaction()

	var isValidTests = []struct {
		label    string
		block    *block.Block
		expected bool
	}{
		{"is valid", &validBlock, true},
		{"invalid once", &invalidNonce, false},
		{"invalid hash", &invalidHash, false},
		{"no coinbase transaction", &noCoinbaseTransaction, false},
	}

	for _, test := range isValidTests {
		t.Run(test.label, testIsValidBlock(test.block, test.expected))
	}
}

func buildBlockInvalidNonce() block.Block {
	block := buildValidBlock()
	block.SetNonce("not a valid nonce")
	return block
}


func buildBlockInvalidHash() block.Block {
	block := buildValidBlock()
	block.hash = "not a valid hash"
	return block
}

func buildBlockNoCoinbaseTransaction() block.Block {
	block := buildValidBlock()
	block.transactions = []block.Transaction{}
	return block
}

func buildValidBlock() block.Block {
	return block.Block{
		previousHash: "PreviousHash",
		timestamp:    1234,
		target:       3,
		nonce:        "2bbc74fc8b213512",
		hash:         "8888fbf1517a4b6a55a6972da919d1bd8d8a3ee3b5e99f3387d61cac95988389",
		transactions: []block.Transaction{
			block.Transaction{
				In: []block.TransactionInput{
					block.TransactionInput{
						ScriptSig: "COINBASE",
					},
				},
				Out: []block.TransactionOutput{
					block.TransactionOutput{
						Value:        1,
						ScriptPubKey: "asdf",
					},
				},
			},
		},
	}
}*/

func testIsValidBlock(block *block.Block, expected bool) func(*testing.T) {
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

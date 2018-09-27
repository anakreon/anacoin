package blockchain

import "testing"

func TestNewBlockchain(t *testing.T) {
	blockchain := NewBlockchain()
	if blockchain.tail != blockchain.head {
		t.Error("tail should point to head")
	}
	if len(blockchain.forkTails) != 1 {
		t.Error("there should be one forktail")
	}
	if blockchain.head != blockchain.forkTails[0] {
		t.Error("head should be the only forktail")
	}
	if blockchain.forkTails[0].Index != 0 {
		t.Error("genesis block index should be 0")
	}
	if blockchain.forkTails[0].Timestamp != 0 {
		t.Error("genesis block Timestamp should be 0")
	}
	if blockchain.forkTails[0].Nonce != "imGenesis" {
		t.Error("genesis block Nonce should be imGenesis")
	}
	if blockchain.forkTails[0].Hash != blockchain.forkTails[0].CalculateHash() {
		t.Error("genesis block Hash is invalid")
	}
}

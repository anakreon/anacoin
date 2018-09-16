package miner

import (
	"time"

	"github.com/anakreon/anacoin/pkg/blockchain"
)

var shouldMine = false

func StartMining() {
	shouldMine = true
	mine()
}

func StopMining() {
	shouldMine = false
}

func mine() {
	for shouldMine {
		block := buildBlock()
		minedBlock := mineBlock(block)
		blockchain.AddToChain(minedBlock)
		blockchain.PrintChain()
	}
}

func buildBlock() blockchain.Block {
	block := blockchain.Block{
		Timestamp:    time.Now().Unix(),
		PreviousHash: blockchain.GetLastBlockHash(),
		Target:       4,
	}
	return block
}

func mineBlock(block blockchain.Block) blockchain.Block {
	for nonce := int64(0); shouldContinueMining(block); nonce++ {
		block.Nonce = string(nonce)
		block.Hash = block.CalculateHash()
	}
	return block
}

func shouldContinueMining(block blockchain.Block) bool {
	return shouldMine && !block.IsValidTargetHash()
}

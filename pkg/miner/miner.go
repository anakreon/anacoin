package miner

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/anakreon/anacoin/pkg/blockchain"
	"github.com/anakreon/anacoin/pkg/mempool"
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
		Transactions: buildTransactions(),
	}
	return block
}

func mineBlock(block blockchain.Block) blockchain.Block {
	for shouldContinueMining(block) {
		block.Nonce = generateRandomHex()
		block.Hash = block.CalculateHash()
	}
	return block
}

func shouldContinueMining(block blockchain.Block) bool {
	return shouldMine && !block.IsValidTargetHash()
}

func generateRandomHex() string {
	randomInt := rand.Int63()
	return strconv.FormatInt(randomInt, 16)
}

func buildTransactions() []blockchain.Transaction {
	return mempool.GetAllTransactions()
}

package miner

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/anakreon/anacoin/pkg/blockchain"
	"github.com/anakreon/anacoin/pkg/connector"
	"github.com/anakreon/anacoin/pkg/mempool"
	"github.com/anakreon/anacoin/pkg/validator"
)

var shouldMine = false
var coinbaseValue float64 = 1

var storage *blockchain.Blockchain

func Initialize(storageInstance *blockchain.Blockchain) {
	storage = storageInstance
}

func StartMining(pubKey string) {
	shouldMine = true
	mine(pubKey)
}

func StopMining() {
	shouldMine = false
}

func mine(pubKey string) {
	for shouldMine {
		candidateBlock := buildCandidateBlock(pubKey)
		minedBlock := mineBlock(candidateBlock)
		storage.AddBlock(minedBlock)
		connector.BroadcastNewBlock(minedBlock)
		mempool.Clear()
	}
}

func buildCandidateBlock(pubKey string) blockchain.Block {
	previousBlock := storage.GetLastBlock()
	candidateBlock := blockchain.Block{
		Index:        previousBlock.Index + 1,
		Timestamp:    time.Now().Unix(),
		PreviousHash: previousBlock.Hash,
		Target:       5,
		Transactions: buildTransactions(pubKey),
	}
	return candidateBlock
}

func mineBlock(block blockchain.Block) blockchain.Block {
	for shouldContinueMining(block) {
		block.Nonce = generateRandomHex()
		block.Hash = block.CalculateHash()
	}
	return block
}

func shouldContinueMining(block blockchain.Block) bool {
	return shouldMine && !validator.IsValidHashAsPerTarget(block.Hash, block.Target)
}

func generateRandomHex() string {
	randomInt := rand.Int63()
	return strconv.FormatInt(randomInt, 16)
}

func buildTransactions(pubKey string) []blockchain.Transaction {
	mempoolTransactions := mempool.GetAllTransactions()
	coinbaseTransactions := []blockchain.Transaction{
		buildCoinbaseTransaction(pubKey),
	}
	return append(coinbaseTransactions, mempoolTransactions...)
}

func buildCoinbaseTransaction(pubKey string) blockchain.Transaction {
	return blockchain.Transaction{
		In: []blockchain.TransactionInput{
			blockchain.TransactionInput{
				ScriptSig: "COINBASE",
			},
		},
		Out: []blockchain.TransactionOutput{
			blockchain.TransactionOutput{
				Value:        coinbaseValue,
				ScriptPubKey: pubKey,
			},
		},
	}
}

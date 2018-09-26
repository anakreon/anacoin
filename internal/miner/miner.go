package miner

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/anakreon/anacoin/internal/blockchain"
	"github.com/anakreon/anacoin/internal/connector"
	"github.com/anakreon/anacoin/internal/mempool"
	"github.com/anakreon/anacoin/internal/validator"
)

type storage interface {
	GetLastBlock() blockchain.Block
	AddBlock(block blockchain.Block)
}

type Miner struct {
	shouldMine bool
}

func NewMiner() Miner {
	return Miner{
		shouldMine: false,
	}
}

func (miner *Miner) Mine(pubKey string, storage storage, unconfirmedTransactions *mempool.UnconfirmedTransactions) {
	miner.shouldMine = true
	for miner.shouldMine {
		candidateBlock := buildCandidateBlock(pubKey, storage.GetLastBlock(), unconfirmedTransactions)
		minedBlock := miner.mineBlock(candidateBlock)
		storage.AddBlock(minedBlock)
		connector.BroadcastNewBlock(minedBlock)
		unconfirmedTransactions.Clear()
	}
}

func (miner *Miner) Stop() {
	miner.shouldMine = false
}

func buildCandidateBlock(pubKey string, lastBlock blockchain.Block, unconfirmedTransactions *mempool.UnconfirmedTransactions) blockchain.Block {
	candidateBlock := blockchain.Block{
		Index:        lastBlock.Index + 1,
		Timestamp:    time.Now().Unix(),
		PreviousHash: lastBlock.Hash,
		Target:       5,
		Transactions: buildTransactions(pubKey, unconfirmedTransactions),
	}
	return candidateBlock
}

func (miner *Miner) mineBlock(block blockchain.Block) blockchain.Block {
	for miner.shouldContinueMining(block) {
		block.Nonce = generateRandomHex()
		block.Hash = block.CalculateHash()
	}
	return block
}

func (miner *Miner) shouldContinueMining(block blockchain.Block) bool {
	return miner.shouldMine && !validator.IsValidHashAsPerTarget(block.Hash, block.Target)
}

func generateRandomHex() string {
	randomInt := rand.Int63()
	return strconv.FormatInt(randomInt, 16)
}

func buildTransactions(pubKey string, unconfirmedTransactions *mempool.UnconfirmedTransactions) []blockchain.Transaction {
	mempoolTransactions := unconfirmedTransactions.GetAllTransactions()
	coinbaseTransactions := []blockchain.Transaction{
		buildCoinbaseTransaction(pubKey),
	}
	return append(coinbaseTransactions, *mempoolTransactions...)
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
				Value:        1,
				ScriptPubKey: pubKey,
			},
		},
	}
}

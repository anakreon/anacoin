package miner

import (
	"math/rand"
	"strconv"

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
	storage                 storage
	unconfirmedTransactions *mempool.UnconfirmedTransactions
	connector               *connector.Connector
	shouldMine              bool
}

func NewMiner(storage *blockchain.Blockchain, unconfirmedTransactions *mempool.UnconfirmedTransactions, connector *connector.Connector) Miner {
	return Miner{
		storage:                 storage,
		unconfirmedTransactions: unconfirmedTransactions,
		connector:               connector,
		shouldMine:              false,
	}
}

func (miner *Miner) Mine(pubKey string) {
	miner.shouldMine = true
	for miner.shouldMine {
		candidateBlock := miner.buildCandidateBlock(pubKey, miner.storage.GetLastBlock())
		minedBlock := miner.mineBlock(candidateBlock)
		miner.storage.AddBlock(minedBlock)
		miner.connector.BroadcastNewBlock(minedBlock)
		miner.unconfirmedTransactions.Clear()
	}
}

func (miner *Miner) Stop() {
	miner.shouldMine = false
}

func (miner *Miner) buildCandidateBlock(pubKey string, lastBlock blockchain.Block) blockchain.Block {
	transactions := miner.buildTransactions(pubKey)
	candidateBlock := blockchain.NewBlock(lastBlock.GetHash(), transactions)
	return *candidateBlock
}

func (miner *Miner) mineBlock(block blockchain.Block) blockchain.Block {
	for miner.shouldContinueMining(block) {
		block.SetNonce(generateRandomHex())
		block.CalculateAndSetHash()
	}
	return block
}

func (miner *Miner) shouldContinueMining(block blockchain.Block) bool {
	return miner.shouldMine && !validator.IsValidHashAsPerTarget(block.GetHash(), block.GetTarget())
}

func generateRandomHex() string {
	randomInt := rand.Int63()
	return strconv.FormatInt(randomInt, 16)
}

func (miner *Miner) buildTransactions(pubKey string) []blockchain.Transaction {
	coinbaseTransactions := []blockchain.Transaction{
		buildCoinbaseTransaction(pubKey),
	}
	return append(coinbaseTransactions, *miner.unconfirmedTransactions...)
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

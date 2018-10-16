package miner

import (
	"math/rand"
	"strconv"

	"github.com/anakreon/anacoin/internal/block"
	"github.com/anakreon/anacoin/internal/mempool"
	"github.com/anakreon/anacoin/internal/validator"
)

type storage interface {
	GetLastBlock() block.Block
	AddBlock(block block.Block)
}

type connector interface {
	BroadcastNewBlock(block block.Block)
}

type Miner struct {
	storage                 storage
	unconfirmedTransactions *mempool.UnconfirmedTransactions
	connector               connector
	shouldMine              bool
}

func NewMiner(storage storage, unconfirmedTransactions *mempool.UnconfirmedTransactions, connector connector) Miner {
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

func (miner *Miner) buildCandidateBlock(pubKey string, lastBlock block.Block) block.Block {
	transactions := miner.buildTransactions(pubKey)
	candidateBlock := block.NewBlock(lastBlock.CalculateHash(), transactions)
	return *candidateBlock
}

func (miner *Miner) mineBlock(block block.Block) block.Block {
	for miner.shouldContinueMining(block) {
		block.SetNonce(generateRandomHex())
	}
	return block
}

func (miner *Miner) shouldContinueMining(block block.Block) bool {
	return miner.shouldMine && !validator.IsValidHashAsPerTarget(block.CalculateHash(), block.GetTarget())
}

func generateRandomHex() string {
	randomInt := rand.Int63()
	return strconv.FormatInt(randomInt, 16)
}

func (miner *Miner) buildTransactions(pubKey string) []block.Transaction {
	coinbaseTransaction := buildCoinbaseTransaction(pubKey)
	return append([]block.Transaction{coinbaseTransaction}, *miner.unconfirmedTransactions...)
}

func buildCoinbaseTransaction(pubKey string) block.Transaction {
	coinbaseTransactionInput := block.NewTransactionInput("", 0)
	coinbaseTransactionInput.SetSignature("COINBASE")
	inputs := []block.TransactionInput{
		coinbaseTransactionInput,
	}
	outputs := []block.TransactionOutput{
		block.NewTransactionOutput(1, pubKey),
	}
	return block.NewTransaction(inputs, outputs)
}

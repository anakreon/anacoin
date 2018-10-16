package blockchain

import (
	"fmt"
	"sync"

	"github.com/anakreon/anacoin/internal/block"
	linkedlist "github.com/anakreon/anacoin/internal/linkedlist-gen"
)

type Blockchain struct {
	list  *linkedlist.List
	mutex *sync.Mutex
}

func NewBlockchain() *Blockchain {
	genesisBlock := createGenesisBlock()
	return &Blockchain{
		list:  linkedlist.NewList(genesisBlock),
		mutex: &sync.Mutex{},
	}
}

func createGenesisBlock() block.Block {
	block := block.Block{}
	block.SetNonce("imGenesis")
	return block
}

func (blockchain *Blockchain) AddBlock(newBlock block.Block) {
	fmt.Println(newBlock)
	blockchain.mutex.Lock()
	previousBlock := blockchain.findBlockByHash(newBlock.GetPreviousHash())
	if previousBlock != nil {
		blockchain.list.AddNode(previousBlock, newBlock)
	}
	blockchain.mutex.Unlock()
}

func (blockchain *Blockchain) findBlockByHash(hash string) (resultBlock *block.Block) {
	iterator := blockchain.list.AllTailsIterator()
	for iterator.HasNext() {
		currentBlock := iterator.Next()
		if currentBlock.CalculateHash() == hash {
			resultBlock = currentBlock
			break
		}
	}
	return
}

func (blockchain *Blockchain) GetLastBlock() block.Block {
	return *blockchain.list.GetMainTailData()
}

type CorrectTransactionCallback func(transaction block.Transaction) bool

func (blockchain *Blockchain) FindTransactions(isCorrectTransaction CorrectTransactionCallback) []block.Transaction {
	matchingTransactions := []block.Transaction{}
	iterator := blockchain.list.Iterator()
	for iterator.HasNext() {
		block := iterator.Next()
		transactions := block.GetTransactions()
		for _, transaction := range transactions {
			if isCorrectTransaction(transaction) {
				matchingTransactions = append(matchingTransactions, transaction)
			}
		}
	}
	return matchingTransactions
}

func (blockchain *Blockchain) FindUnspentTransactionOutputs() block.UnspentTransactionOutputs {
	unspentTransactionOutputs := make(block.UnspentTransactionOutputs)
	iterator := blockchain.list.Iterator()
	for iterator.HasNext() {
		block := iterator.Next()
		transactions := block.GetTransactions()
		unspentTransactionOutputs.UpdateFromTransactions(transactions)
	}
	return unspentTransactionOutputs
}

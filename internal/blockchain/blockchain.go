package blockchain

import (
	"fmt"
	"sync"

	"github.com/anakreon/anacoin/internal/linkedlist"
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

func createGenesisBlock() *Block {
	genesisBlock := Block{
		timestamp: 0,
		nonce:     "imGenesis",
	}
	return &genesisBlock
}

func (blockchain *Blockchain) AddBlock(newBlock Block) {
	blockchain.mutex.Lock()
	previousBlock := blockchain.findBlockByHash(newBlock.previousHash)
	if previousBlock != nil {
		blockchain.list.AddNode(previousBlock, newBlock)
	}
	blockchain.mutex.Unlock()
}

func (blockchain *Blockchain) findBlockByHash(hash string) (resultBlock *Block) {
	iterator := blockchain.list.AllTailsIterator()
	for iterator.HasNext() {
		currentBlockX := iterator.Next()
		fmt.Println(currentBlockX)
		//
		currentBlock := currentBlockX.(*Block)
		if currentBlock.CalculateHash() == hash {
			resultBlock = currentBlock
			break
		}
	}
	return
}

func (blockchain *Blockchain) GetLastBlock() Block {
	block, _ := blockchain.list.GetMainTailData().(Block)
	return block
}

type CorrectTransactionCallback func(transaction Transaction) bool

func (blockchain *Blockchain) FindTransactions(isCorrectTransaction CorrectTransactionCallback) []Transaction {
	matchingTransactions := []Transaction{}
	iterator := blockchain.list.Iterator()
	for iterator.HasNext() {
		block := iterator.Next().(Block)
		transactions := block.GetTransactions()
		for _, transaction := range transactions {
			if isCorrectTransaction(transaction) {
				matchingTransactions = append(matchingTransactions, transaction)
			}
		}
	}
	return matchingTransactions
}

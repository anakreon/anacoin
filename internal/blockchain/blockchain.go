package blockchain

import (
	"log"
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
	genesisBlock.CalculateAndSetHash()
	return &genesisBlock
}

func (blockchain *Blockchain) AddBlock(newBlock Block) {
	log.Println(newBlock)
	blockchain.mutex.Lock()
	previousNode := blockchain.findNodeByBlockHash(newBlock.previousHash)
	if previousNode != nil {
		blockchain.list.CreateNewNodeAndLinkWithPreviousNode(previousNode, newBlock)
	}
	blockchain.mutex.Unlock()
}

func (blockchain *Blockchain) findNodeByBlockHash(hash string) *linkedlist.Node {
	return blockchain.list.FindNodeInList(func(node *linkedlist.Node) bool {
		nodeData := node.GetData()
		return (*nodeData).GetHash() == hash
	})
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

package blockchain

import (
	"fmt"
	"sync"
)

type Blockchain struct {
	head      *Block
	tail      *Block
	forkTails []*Block
	mutex     *sync.Mutex
}

func NewBlockchain() *Blockchain {
	head := createGenesisBlock()
	return &Blockchain{
		head:      head,
		tail:      head,
		forkTails: []*Block{head},
		mutex:     &sync.Mutex{},
	}
}

func createGenesisBlock() *Block {
	genesisBlock := Block{
		Index:     0,
		Timestamp: 0,
		Nonce:     "imGenesis",
	}
	genesisBlock.Hash = genesisBlock.CalculateHash()
	return &genesisBlock
}

func (blockchain *Blockchain) AddBlock(newBlock Block) {
	fmt.Println(newBlock)
	blockchain.mutex.Lock()
	previousBlock := blockchain.findBlockByHash(newBlock.PreviousHash)
	if previousBlock != nil {
		newBlock.PreviousBlock = previousBlock
		blockchain.linkNewBlock(previousBlock, &newBlock)
	}
	blockchain.mutex.Unlock()
}

func (blockchain *Blockchain) linkNewBlock(block *Block, newBlock *Block) {
	if block.HasNextBlock() {
		blockchain.addNewForkTail(newBlock)
	} else {
		block.NextBlock = newBlock
		blockchain.switchToMainChainIfLonger(newBlock)
	}
}

func (blockchain *Blockchain) switchToMainChainIfLonger(block *Block) {
	if block.Index > blockchain.tail.Index {
		blockchain.switchToMainChain(block)
	}
}

func (blockchain *Blockchain) switchToMainChain(newTail *Block) {
	blockchain.tail = newTail
	for iterator := blockchain.tail; iterator != blockchain.head; iterator = iterator.PreviousBlock {
		iterator.PreviousBlock.NextBlock = iterator
	}
}

func (blockchain *Blockchain) addNewForkTail(block *Block) {
	blockchain.forkTails = append(blockchain.forkTails, block)
}

func (blockchain *Blockchain) findBlockByHash(hash string) (block *Block) {
Mainloop:
	for _, forkTail := range blockchain.forkTails {
		for iterator := forkTail; iterator.IsLinkedWithPreviousBlock(); iterator = iterator.PreviousBlock {
			if iterator.Hash == hash {
				block = iterator
				break Mainloop
			}
		}
	}
	return
}

func (blockchain *Blockchain) GetLastBlock() Block {
	return *blockchain.tail
}

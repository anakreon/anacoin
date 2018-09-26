package blockchain

import (
	"fmt"
)

type Blockchain struct {
	head      *Block
	tail      *Block
	forkTails []*Block
}

func NewBlockchain() *Blockchain {
	head := generateGenesisBlock()
	return &Blockchain{
		head:      head,
		tail:      head,
		forkTails: []*Block{head},
	}
}

func generateGenesisBlock() *Block {
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
	previousBlock := blockchain.findBlockByHash(newBlock.PreviousHash)
	if previousBlock != nil {
		newBlock.PreviousBlock = previousBlock
		if previousBlock.HasNextBlock() {
			blockchain.addNewForkTail(&newBlock)
		} else {
			previousBlock.NextBlock = &newBlock
			blockchain.switchToMainChainIfLonger(&newBlock)
		}
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

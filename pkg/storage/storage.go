package storage

import (
	"github.com/anakreon/anacoin/pkg/blockchain"
)

var head = generateGenesisBlock()
var tail = head
var forkTails = []*blockchain.Block{tail}

func generateGenesisBlock() *blockchain.Block {
	genesisBlock := blockchain.Block{
		Index:     0,
		Timestamp: 0,
		Nonce:     "imGenesis",
	}
	genesisBlock.Hash = genesisBlock.CalculateHash()
	return &genesisBlock
}

func AddBlock(newBlock blockchain.Block) {
	previousBlock := findBlockByHash(newBlock.PreviousHash)
	if previousBlock != nil {
		newBlock.PreviousBlock = previousBlock
		if isLastElementOfChain(previousBlock) {
			previousBlock.NextBlock = &newBlock
			switchToMainChainIfLonger(&newBlock)
		} else {
			addNewForkTail(&newBlock)
		}
	}
}

func isLastElementOfChain(block *blockchain.Block) bool {
	return block.NextBlock == nil
}

func switchToMainChainIfLonger(block *blockchain.Block) {
	if block.Index > tail.Index {
		switchToMainChain(block)
	}
}

func switchToMainChain(newTail *blockchain.Block) {
	tail = newTail
	for iterator := tail; iterator != head; iterator = iterator.PreviousBlock {
		iterator.PreviousBlock.NextBlock = iterator
	}
}

func addNewForkTail(block *blockchain.Block) {
	forkTails = append(forkTails, block)
}

func findBlockByHash(hash string) (block *blockchain.Block) {
Mainloop:
	for _, forkTail := range forkTails {
		for iterator := forkTail; isPreviousBlockLinkedToMe(iterator); iterator = iterator.PreviousBlock {
			if iterator.Hash == hash {
				block = iterator
				break Mainloop
			}
		}
	}
	return
}

func isPreviousBlockLinkedToMe(block *blockchain.Block) bool {
	return block.PreviousBlock != nil && block.PreviousBlock.NextBlock == block
}

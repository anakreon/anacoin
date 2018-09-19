package blockchain

import (
	"fmt"
)

var chain = []Block{generateGenesisBlock()}

func generateGenesisBlock() Block {
	genesisBlock := Block{
		Index:        0,
		Timestamp:    0,
		PreviousHash: "",
		Nonce:        "imGenesis",
	}
	genesisBlock.Hash = genesisBlock.CalculateHash()
	return genesisBlock
}

func AddToChain(block Block) {
	if block.Validate(getLastBlock()) {
		fmt.Println(block)
		chain = append(chain, block)
	}
}

func PrintChain() {
	fmt.Println(chain)
}

func getLastBlock() Block {
	return chain[len(chain)-1]
}

func GetLastBlock() Block {
	return getLastBlock()
}

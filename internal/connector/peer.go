package connector

import "github.com/anakreon/anacoin/internal/blockchain"

type Peer interface {
	SendBlock(block blockchain.Block)
	SendTransaction(transaction blockchain.Transaction)
}

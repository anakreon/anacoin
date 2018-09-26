package connector

import "github.com/anakreon/anacoin/pkg/blockchain"

type Peer interface {
	SendBlock(block blockchain.Block)
	SendTransaction(transaction blockchain.Transaction)
}

type Receiver struct{}

func (receiver Receiver) SendBlock(block blockchain.Block) {
	ReceiveBlock(block)
	//else remove peer
}

func (receiver Receiver) SendTransaction(transaction blockchain.Transaction) {
	ReceiveTransaction(transaction)
	//else remove peer
}

package connector

import "github.com/anakreon/anacoin/internal/blockchain"

type Peer interface {
	SendBlock(block blockchain.Block)
	SendTransaction(transaction blockchain.Transaction)
}

type Receiver struct {
	connector *Connector
}

func (receiver Receiver) SendBlock(block blockchain.Block) {
	receiver.connector.ReceiveBlock(block)
	//else remove peer
}

func (receiver Receiver) SendTransaction(transaction blockchain.Transaction) {
	receiver.connector.ReceiveTransaction(transaction)
	//else remove peer
}

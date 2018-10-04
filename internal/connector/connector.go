package connector

import (
	"log"

	"github.com/anakreon/anacoin/internal/blockchain"
	"github.com/anakreon/anacoin/internal/mempool"
	"github.com/anakreon/anacoin/internal/validator"
)

type Connector struct {
	storage                 *blockchain.Blockchain
	unconfirmedTransactions *mempool.UnconfirmedTransactions
	peers                   Peers
}

func NewConnector(storage *blockchain.Blockchain, unconfirmedTransactions *mempool.UnconfirmedTransactions) *Connector {
	connector := Connector{
		storage:                 storage,
		unconfirmedTransactions: unconfirmedTransactions,
		peers:                   Peers{},
	}
	return &connector
}

func (connector *Connector) AddPeer(peer Peer) {
	connector.peers = append(connector.peers, peer)
}

func (connector *Connector) ReceiveBlock(block blockchain.Block) {
	if validator.IsValidBlock(&block) {
		log.Println("receiving valid block")
		connector.storage.AddBlock(block)
	}
}

func (connector *Connector) ReceiveTransaction(transaction blockchain.Transaction) {
	if validator.IsTransactionValid() {
		log.Println("receiving valid transaction")
		connector.unconfirmedTransactions.AddTransaction(transaction)
	}
}

func (connector *Connector) BroadcastNewBlock(block blockchain.Block) {
	for _, peer := range connector.peers {
		log.Println("sending new block")
		peer.SendBlock(block)
	}
}

func (connector *Connector) BroadcastNewTransaction(transaction blockchain.Transaction) {
	for _, peer := range connector.peers {
		log.Println("sending new transaction")
		peer.SendTransaction(transaction)
	}
}

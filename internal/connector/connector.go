package connector

import (
	"log"

	"github.com/anakreon/anacoin/internal/block"
	"github.com/anakreon/anacoin/internal/validator"
)

type storage interface {
	AddBlock(block block.Block)
}

type unconfirmedTransactions interface {
	AddTransaction(transaction block.Transaction)
}

type Connector struct {
	storage                 storage
	unconfirmedTransactions unconfirmedTransactions
	peers                   Peers
}

func NewConnector(storage storage, unconfirmedTransactions unconfirmedTransactions) *Connector {
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

func (connector *Connector) ReceiveBlock(block block.Block) {
	if validator.IsValidBlock(&block) {
		log.Println("receiving valid block")
		connector.storage.AddBlock(block)
	}
}

func (connector *Connector) ReceiveTransaction(transaction block.Transaction) {
	if validator.IsTransactionValid() {
		log.Println("receiving valid transaction")
		connector.unconfirmedTransactions.AddTransaction(transaction)
	}
}

func (connector *Connector) BroadcastNewBlock(block block.Block) {
	for _, peer := range connector.peers {
		log.Println("sending new block")
		peer.SendBlock(block)
	}
}

func (connector *Connector) BroadcastNewTransaction(transaction block.Transaction) {
	for _, peer := range connector.peers {
		log.Println("sending new transaction")
		peer.SendTransaction(transaction)
	}
}

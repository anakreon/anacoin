package connector

import (
	"github.com/anakreon/anacoin/internal/blockchain"
	"github.com/anakreon/anacoin/internal/mempool"
	"github.com/anakreon/anacoin/internal/validator"
)

type Connection interface {
	GetConnector() *Connector
}

type Connector struct {
	storage                 *blockchain.Blockchain
	unconfirmedTransactions *mempool.UnconfirmedTransactions
	peerReceivers           PeerReceivers
}

func NewConnector(storage *blockchain.Blockchain, unconfirmedTransactions *mempool.UnconfirmedTransactions) *Connector {
	connector := Connector{
		storage:                 storage,
		unconfirmedTransactions: unconfirmedTransactions,
	}
	connector.peerReceivers = initiateConnectionWithPeers(&connector)
	return &connector
}

func initiateConnectionWithPeers(connector *Connector) PeerReceivers {
	peerReceivers := PeerReceivers{}
	connections := []Connection{}
	for _, connection := range connections {
		receiver := Receiver{connector}
		peer := connection.GetConnector().Connect(receiver)
		peerReceivers.AddPeer(peer, receiver)
	}
	return peerReceivers
}

func (connector *Connector) Connect(peer Peer) Peer {
	receiver := Receiver{connector}
	connector.peerReceivers.AddPeer(peer, receiver)
	return receiver
}

func (connector *Connector) ReceiveBlock(block blockchain.Block) {
	if validator.IsValidBlock(&block) {
		connector.storage.AddBlock(block)
	}
}

func (connector *Connector) ReceiveTransaction(transaction blockchain.Transaction) {
	if validator.IsTransactionValid() {
		connector.unconfirmedTransactions.AddTransaction(transaction)
	}
}

func (connector *Connector) BroadcastNewBlock(block blockchain.Block) {
	for _, peerReceiver := range connector.peerReceivers {
		peerReceiver.Peer.SendBlock(block)
	}
}

func (connector *Connector) BroadcastNewTransaction(transaction blockchain.Transaction) {
	for _, peerReceiver := range connector.peerReceivers {
		peerReceiver.Peer.SendTransaction(transaction)
	}
}

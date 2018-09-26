package connector

import (
	"github.com/anakreon/anacoin/pkg/blockchain"
	"github.com/anakreon/anacoin/pkg/mempool"
	"github.com/anakreon/anacoin/pkg/validator"
)

var storage *blockchain.Blockchain

func Initialize(storageInstance *blockchain.Blockchain) {
	storage = storageInstance
	initiateConnectionWithPeers()
}

func initiateConnectionWithPeers() {
	/* for each connection
	receiver := Receiver{}
	peer := CONNECTION.connector.Connect(receiver)
	AddPeer(peer, receiver)
	*/
}

func Connect(peer Peer) Peer {
	receiver := Receiver{}
	AddPeer(peer, receiver)
	return receiver
}

func ReceiveBlock(block blockchain.Block) {
	if validator.IsValidBlock(&block) {
		storage.AddBlock(block)
	}
}

func ReceiveTransaction(transaction blockchain.Transaction) {
	if validator.IsTransactionValid() {
		mempool.AddTransaction(transaction)
	}
}

func BroadcastNewBlock(block blockchain.Block) {
	for _, peer := range peers {
		peer.Peer.SendBlock(block)
	}
}

func BroadcastNewTransaction(transaction blockchain.Transaction) {
	for _, peer := range peers {
		peer.Peer.SendTransaction(transaction)
	}
}

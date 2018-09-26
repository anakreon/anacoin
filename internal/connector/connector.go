package connector

import (
	"github.com/anakreon/anacoin/internal/blockchain"
	"github.com/anakreon/anacoin/internal/mempool"
	"github.com/anakreon/anacoin/internal/validator"
)

var storage *blockchain.Blockchain
var unconfirmedTransactions *mempool.UnconfirmedTransactions

func Initialize(storageInstance *blockchain.Blockchain, unconfirmedTransactionsInstance *mempool.UnconfirmedTransactions) {
	storage = storageInstance
	unconfirmedTransactions = unconfirmedTransactionsInstance
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
		unconfirmedTransactions.AddTransaction(transaction)
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

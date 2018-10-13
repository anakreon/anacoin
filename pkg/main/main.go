package main

import (
	"net/rpc"
	"time"

	"github.com/anakreon/anacoin/internal/arpc"
	"github.com/anakreon/anacoin/internal/blockchain"
	"github.com/anakreon/anacoin/internal/connector"
	"github.com/anakreon/anacoin/internal/mempool"
	"github.com/anakreon/anacoin/internal/miner"
	"github.com/anakreon/anacoin/internal/wallet"
)

func main() {
	channel := make(chan string, 0)
	run()
	<-channel
}

type RpcPeer struct {
	rpcClient *rpc.Client
}

func (peer RpcPeer) SendBlock(block blockchain.Block) {
	peer.rpcClient.Call("ConnectorRpc.ReceiveBlock", arpc.ReceiveBlockArgs{block}, nil)
}

func (peer RpcPeer) SendTransaction(transaction blockchain.Transaction) {
	peer.rpcClient.Call("ConnectorRpc.ReceiveTransaction", arpc.ReceiveTransactionArgs{transaction}, nil)
}

func run() {
	storage := blockchain.NewBlockchain()
	unconfirmedTransactions := mempool.NewUnconfirmedTransactions()
	connectorInstance := connector.NewConnector(storage, &unconfirmedTransactions)
	wallet := wallet.NewWallet(storage, &unconfirmedTransactions, connectorInstance)
	publicAddress := wallet.GetPublicAddress()

	miner := miner.NewMiner(storage, &unconfirmedTransactions, connectorInstance)
	go miner.Mine(publicAddress)

	/*anacoinRPC := arpc.NewAnacoinRpc(&miner, &wallet)
	go arpc.Serve(&anacoinRPC, "", "2222")

	client := arpc.Connect("localhost", "2222")
	client.Call("AnacoinRpc.Mine", arpc.MineArgs{publicAddress}, nil)*/

	connectorRPC := arpc.NewConnectorRpc(connectorInstance)
	go arpc.Serve(&connectorRPC, "", "1111")

	time.Sleep(30 * time.Second)
	clientOne := arpc.Connect("localhost", "2222")
	peerOne := RpcPeer{clientOne}
	connectorInstance.AddPeer(peerOne)

	//miner.Mine(publicAddress)
	/*peerReceivers := PeerReceivers{}
	for _, connection := range connections {
		receiver := Receiver{connector}
		peer := connection.GetConnector().Connect(receiver)
		peerReceivers.AddPeer(peer, receiver)
	}
	return peerReceivers*/
}

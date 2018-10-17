package main

import (
	"encoding/json"
	"net/rpc"
	"syscall/js"

	"github.com/anakreon/anacoin/internal/arpc"
	"github.com/anakreon/anacoin/internal/block"
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

func (peer RpcPeer) SendBlock(block block.Block) {
	peer.rpcClient.Call("ConnectorRpc.ReceiveBlock", arpc.ReceiveBlockArgs{block}, nil)
}

func (peer RpcPeer) SendTransaction(transaction block.Transaction) {
	peer.rpcClient.Call("ConnectorRpc.ReceiveTransaction", arpc.ReceiveTransactionArgs{transaction}, nil)
}

type WasmPeer struct {
	connector *connector.Connector
}

func (peer WasmPeer) SendBlock(block block.Block) {
	blockJson, _ := json.Marshal(block)
	wasmAPI := js.Global().Get("wasmSendBlock")
	wasmAPI.Invoke(string(blockJson))
}

func (peer WasmPeer) SendTransaction(transaction block.Transaction) {
	transactionJson, _ := json.Marshal(transaction)
	wasmAPI := js.Global().Get("wasmSendTransaction")
	wasmAPI.Invoke(string(transactionJson))
}

func run() {
	storage := blockchain.NewBlockchain()
	unconfirmedTransactions := mempool.NewUnconfirmedTransactions()
	connectorInstance := connector.NewConnector(storage, &unconfirmedTransactions)
	wallet := wallet.NewWallet(storage, &unconfirmedTransactions, connectorInstance)
	publicAddress := wallet.GetPublicAddress()

	miner := miner.NewMiner(storage, &unconfirmedTransactions, connectorInstance)
	go miner.Mine(publicAddress)

	wasmPeer := WasmPeer{connectorInstance}
	connectorInstance.AddPeer(wasmPeer)

	/*anacoinRPC := arpc.NewAnacoinRpc(&miner, &wallet)
	go arpc.Serve(&anacoinRPC, "", "2222")

	client := arpc.Connect("localhost", "2222")
	client.Call("AnacoinRpc.Mine", arpc.MineArgs{publicAddress}, nil)*/
	/*
		connectorRPC := arpc.NewConnectorRpc(connectorInstance)
		go arpc.Serve(&connectorRPC, "", "1111")

		time.Sleep(30 * time.Second)
		clientOne := arpc.Connect("localhost", "2222")
		peerOne := RpcPeer{clientOne}
		connectorInstance.AddPeer(peerOne)*/

	//miner.Mine(publicAddress)
	/*peerReceivers := PeerReceivers{}
	for _, connection := range connections {
		receiver := Receiver{connector}
		peer := connection.GetConnector().Connect(receiver)
		peerReceivers.AddPeer(peer, receiver)
	}
	return peerReceivers*/
}

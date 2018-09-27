package main

import (
	"github.com/anakreon/anacoin/internal/blockchain"
	"github.com/anakreon/anacoin/internal/connector"
	"github.com/anakreon/anacoin/internal/mempool"
	"github.com/anakreon/anacoin/internal/miner"
	"github.com/anakreon/anacoin/internal/wallet"
)

func main() {
	storage := blockchain.NewBlockchain()
	unconfirmedTransactions := mempool.NewUnconfirmedTransactions()
	connector := connector.NewConnector(storage, &unconfirmedTransactions)
	wallet := wallet.NewWallet(storage, &unconfirmedTransactions, connector)
	publicAddress := wallet.GetPublicAddress()
	miner := miner.NewMiner(storage, &unconfirmedTransactions, connector)
	miner.Mine(publicAddress)
}

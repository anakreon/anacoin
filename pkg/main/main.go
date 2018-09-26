package main

import (
	"github.com/anakreon/anacoin/internal/blockchain"
	"github.com/anakreon/anacoin/internal/mempool"
	"github.com/anakreon/anacoin/internal/miner"
	"github.com/anakreon/anacoin/internal/wallet"
)

func main() {
	unconfirmedTransactions := mempool.NewUnconfirmedTransactions()
	storage := blockchain.NewBlockchain()
	wallet := wallet.NewWallet(storage, &unconfirmedTransactions)
	publicAddress := wallet.GetPublicAddress()
	miner := miner.NewMiner()
	miner.Mine(publicAddress, storage, &unconfirmedTransactions)
}

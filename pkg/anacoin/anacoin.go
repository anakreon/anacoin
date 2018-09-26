package main

import (
	"github.com/anakreon/anacoin/pkg/blockchain"
	"github.com/anakreon/anacoin/pkg/miner"
	"github.com/anakreon/anacoin/pkg/wallet"
)

func main() {
	storage := blockchain.NewBlockchain()
	wallet.Initialize(storage)
	publicAddress := wallet.GetPublicAddress()
	miner.Initialize(storage)
	miner.StartMining(publicAddress)
}

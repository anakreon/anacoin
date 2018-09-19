package main

import (
	"github.com/anakreon/anacoin/pkg/miner"
	"github.com/anakreon/anacoin/pkg/wallet"
)

func main() {
	wallet.Initialize()
	publicAddress := wallet.GetPublicAddress()
	miner.StartMining(publicAddress)
}

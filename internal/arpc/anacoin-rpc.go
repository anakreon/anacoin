package arpc

import (
	"github.com/anakreon/anacoin/internal/miner"
	"github.com/anakreon/anacoin/internal/wallet"
)

type AnacoinRpc struct {
	miner  *miner.Miner
	wallet *wallet.Wallet
}

type MineArgs struct {
	PublicAddress string
}

type AddTransactionArgs struct {
	TargetAddress string
	Value         uint64
}

func NewAnacoinRpc(miner *miner.Miner, wallet *wallet.Wallet) AnacoinRpc {
	return AnacoinRpc{
		miner:  miner,
		wallet: wallet,
	}
}

func (rpc *AnacoinRpc) Mine(args *MineArgs, _ *bool) error {
	rpc.miner.Mine(args.PublicAddress)
	return nil
}

func (rpc *AnacoinRpc) StopMining(_ bool, _ *bool) error {
	rpc.miner.Stop()
	return nil
}

func (rpc *AnacoinRpc) GetPublicAddress(_ bool, publicAddress *string) error {
	*publicAddress = rpc.wallet.GetPublicAddress()
	return nil
}

func (rpc *AnacoinRpc) AddTransaction(args *AddTransactionArgs, _ *bool) error {
	rpc.wallet.AddTransaction(args.TargetAddress, args.Value)
	return nil
}

func (rpc *AnacoinRpc) GetUnspentValue(_ bool, _ *bool) error {
	rpc.wallet.GetUnspentValue()
	return nil
}

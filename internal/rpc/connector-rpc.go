package rpc

import (
	"github.com/anakreon/anacoin/internal/blockchain"
	"github.com/anakreon/anacoin/internal/connector"
)

type ConnectorRpc struct {
	connector *connector.Connector
}

type ReceiveBlockArgs struct {
	Block blockchain.Block
}

type ReceiveTransactionArgs struct {
	Transaction blockchain.Transaction
}

func (rpc *ConnectorRpc) ReceiveBlock(args *ReceiveBlockArgs, _ *bool) error {
	rpc.connector.ReceiveBlock(args.Block)
	return nil
}

func (rpc *ConnectorRpc) ReceiveTransaction(args *ReceiveTransactionArgs, _ *bool) error {
	rpc.connector.ReceiveTransaction(args.Transaction)
	return nil
}

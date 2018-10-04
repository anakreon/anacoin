package arpc

import (
	"github.com/anakreon/anacoin/internal/blockchain"
	"github.com/anakreon/anacoin/internal/connector"
)

type ConnectorRpc struct {
	connector *connector.Connector
}

type ConnectArgs struct {
	Peer connector.Peer
}

type ReceiveBlockArgs struct {
	Block blockchain.Block
}

type ReceiveTransactionArgs struct {
	Transaction blockchain.Transaction
}

func NewConnectorRpc(connector *connector.Connector) ConnectorRpc {
	return ConnectorRpc{
		connector: connector,
	}
}

func (rpc *ConnectorRpc) ReceiveBlock(args *ReceiveBlockArgs, _ *bool) error {
	rpc.connector.ReceiveBlock(args.Block)
	return nil
}

func (rpc *ConnectorRpc) ReceiveTransaction(args *ReceiveTransactionArgs, _ *bool) error {
	rpc.connector.ReceiveTransaction(args.Transaction)
	return nil
}

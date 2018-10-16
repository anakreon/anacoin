package connector

import (
	"github.com/anakreon/anacoin/internal/block"
)

type Peer interface {
	SendBlock(block block.Block)
	SendTransaction(transaction block.Transaction)
}

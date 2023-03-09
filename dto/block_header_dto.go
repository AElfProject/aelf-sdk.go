package dto

import (
	"time"
)

//BlockHeaderDto BlockHeaderDto
type BlockHeaderDto struct {
	PreviousBlockHash                string
	MerkleTreeRootOfTransactions     string
	MerkleTreeRootOfWorldState       string
	MerkleTreeRootOfTransactionState string
	Extra                            string
	Height                           int64
	Time                             time.Time
	ChainId                          string
	Bloom                            string
	SignerPubkey                     string
}

package dto

//BlockHeaderDto BlockHeaderDto
type BlockHeaderDto struct {
	PreviousBlockHash            string
	MerkleTreeRootOfTransactions string
	MerkleTreeRootOfWorldState   string
	Extra                        string
	Height                       int64
	Time                         string
	ChainId                      string
	Bloom                        string
	SignerPubkey                 string
}

package dto

//ChainStatusDto ChainStatusDto
type ChainStatusDto struct {
	ChainId                     string
	Branches                    map[string]interface{}
	NotLinkedBlocks             map[string]interface{}
	LongestChainHeight          int64
	LongestChainHash            string
	GenesisBlockHash            string
	GenesisContractAddress      string
	LastIrreversibleBlockHash   string
	LastIrreversibleBlockHeight int64
	BestChainHash               string
	BestChainHeight             int64
}

//TaskQueueInfoDto TaskQueueInfoDto
type TaskQueueInfoDto struct {
	Name string
	Size int
}

//MerklePathDto MerklePathDto
type MerklePathDto struct {
	MerklePathNodes []MerklePathNodeDto
}

//MerklePathNodeDto MerklePathNodeDto
type MerklePathNodeDto struct {
	Hash            string
	IsLeftChildNode bool
}

//RoundDto RoundDto
type RoundDto struct {
	RoundNumber                           int64
	TermNumber                            int64
	RoundId                               int64
	RealTimeMinerInformation              []byte
	ExtraBlockProducerOfPreviousRound     string
	ConfirmedIrreversibleBlockRoundNumber int64
	ConfirmedIrreversibleBlockHeight      int64
	IsMinerListJustChanged                bool
}

//MinerInRoundDto MinerInRoundDto
type MinerInRoundDto struct {
	Order                          int
	ProducedTinyBlocks             int
	ExpectedMiningTime             string
	ActualMiningTimes              []string
	InValue                        string
	PreviousInValue                string
	OutValue                       string
	ProducedBlocks                 int32
	MissedBlocks                   int64
	ImpliedIrreversibleBlockHeight int64
}

//BlockBodyDto BlockBodyDto
type BlockBodyDto struct {
	TransactionsCount int
	Transactions      []string
}

//BlockDto BlockDto
type BlockDto struct {
	BlockHash string
	Header    BlockHeaderDto
	Body      BlockBodyDto
}

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

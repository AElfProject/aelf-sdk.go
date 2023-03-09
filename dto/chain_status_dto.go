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

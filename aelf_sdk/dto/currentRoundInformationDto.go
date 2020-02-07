package dto

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

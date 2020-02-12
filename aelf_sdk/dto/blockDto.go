package dto

//BlockDto BlockDto
type BlockDto struct {
	BlockHash string
	Header    BlockHeaderDto
	Body      BlockBodyDto
}

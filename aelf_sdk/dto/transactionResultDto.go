package dto

//TransactionResultDto TransactionResultDto
type TransactionResultDto struct {
	TransactionId       string
	Status              string
	Logs                []LogEventDto
	Bloom               string
	BlockNumber         int64
	BlockHash           string
	Transaction         TransactionDto
	ReturnValue         string
	ReadableReturnValue string
	Error               string
}

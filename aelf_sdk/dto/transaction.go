package dto

//CreateRawTransactionInput CreateRawTransactionInput
type CreateRawTransactionInput struct {
	From           string
	MethodName     string
	Params         string
	RefBlockHash   string
	RefBlockNumber int64
	To             string
}

//ExecuteRawTransactionDto ExecuteRawTransactionDto
type ExecuteRawTransactionDto struct {
	RawTransaction string
	Signature      string
}

//SendRawTransactionInput SendRawTransactionInput
type SendRawTransactionInput struct {
	Transaction       string
	Signature         string
	ReturnTransaction bool
}

//SendRawTransactionOutput SendRawTransactionOutput
type SendRawTransactionOutput struct {
	TransactionId string
	Transaction   TransactionDto
}

//TransactionDto TransactionDto
type TransactionDto struct {
	From           string
	To             string
	RefBlockNumber int64
	RefBlockPrefix string
	MethodName     string
	Params         string
	Signature      string
}

//TransactionPoolStatus TransactionPoolStatus
type TransactionPoolStatus struct {
	Queued    int
	Validated int
}

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

//LogEventDto LogEventDto
type LogEventDto struct {
	Address    string
	Name       string
	Indexed    []string
	NonIndexed string
}

//SendTransactionOutput SendTransactionOutput
type SendTransactionOutput struct {
	TransactionID string
}

//CreateRawTransactionOutput Create RawTransactionOutput
type CreateRawTransactionOutput struct {
	RawTransaction string
}

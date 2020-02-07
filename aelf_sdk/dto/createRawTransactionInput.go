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

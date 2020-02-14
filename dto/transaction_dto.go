package dto

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

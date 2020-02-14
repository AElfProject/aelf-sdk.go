package dto

//SendRawTransactionInput SendRawTransactionInput
type SendRawTransactionInput struct {
	Transaction       string
	Signature         string
	ReturnTransaction bool
}
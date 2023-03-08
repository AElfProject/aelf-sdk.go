package dto

type CalculateTransactionFeeOutput struct {
	Success        bool
	TransactionFee map[string]interface{}
	ResourceFee    map[string]interface{}
}

package dto

type TransactionFeeResultOutput struct {
	Success        bool
	TransactionFee map[string]interface{}
	ResourceFee    map[string]interface{}
}

package client

import (
	"encoding/json"

	"aelf_sdk.go/aelf_sdk/dto"
	util "aelf_sdk.go/aelf_sdk/utils"
)

//GetTransactionPoolStatus Get information about the current transaction pool.
func (a *AElfClient) GetTransactionPoolStatus() (*dto.TransactionPoolStatus, error) {
	url := a.Host + TRANSACTIONPOOLSTATUS
	transactionPoolBytes, err := util.GetRequest("GET", url, a.Version, nil)
	if err != nil {
		return nil, err
	}
	var transactionPool = new(dto.TransactionPoolStatus)
	json.Unmarshal(transactionPoolBytes, &transactionPool)
	return transactionPool, nil
}

//GetTransactionResult Gets the result of transaction execution by the given transactionId.
func (a *AElfClient) GetTransactionResult(transactionID string) (*dto.TransactionResultDto, error) {
	url := a.Host + TRANSACTIONRESULT
	params := map[string]interface{}{"transactionId": transactionID}
	transactionBytes, err := util.GetRequest("GET", url, a.Version, params)
	if err != nil {
		return nil, err
	}
	var transaction = new(dto.TransactionResultDto)
	json.Unmarshal(transactionBytes, &transaction)
	return transaction, nil
}

//GetTransactionResults Get results of multiple transactions by specified blockHash
func (a *AElfClient) GetTransactionResults(blockHash string, offset, limit int) ([]*dto.TransactionResultDto, error) {
	url := a.Host + TRANSACTIONRESULTS
	params := map[string]interface{}{
		"blockHash": blockHash,
		"offset":    offset,
		"limit":     limit,
	}
	transactionsBytes, err := util.GetRequest("GET", url, a.Version, params)
	if err != nil {
		return nil, err
	}
	var datas interface{}
	json.Unmarshal(transactionsBytes, &datas)
	var transactions []*dto.TransactionResultDto
	for _, d := range datas.([]interface{}) {
		var transaction = new(dto.TransactionResultDto)
		Bytes, err := json.Marshal(d)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(Bytes, &transaction)
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

//GetMerklePathByTransactionID Get merkle path of a transaction.
func (a *AElfClient) GetMerklePathByTransactionID(transactionID string) (*dto.MerklePathDto, error) {
	url := a.Host + MBYTRANSACTIONID
	params := map[string]interface{}{"transactionId": transactionID}
	merkleBytes, err := util.GetRequest("GET", url, a.Version, params)
	if err != nil {
		return nil, err
	}
	var merkle = new(dto.MerklePathDto)
	json.Unmarshal(merkleBytes, &merkle)
	return merkle, nil
}

//ExecuteTransaction  Call a read-only method of a contract.
func (a *AElfClient) ExecuteTransaction(rawTransaction string) (string, error) {
	url := a.Host + EXECUTETRANSACTION
	params := map[string]interface{}{"RawTransaction": rawTransaction}
	transactionBytes, err := util.PostRequest(url, a.Version, params)
	if err != nil {
		return "", err
	}
	return util.BytesToString(transactionBytes), nil
}

//ExecuteRawTransaction Call a method of a contract by given serialized strings.
func (a *AElfClient) ExecuteRawTransaction(input *dto.ExecuteRawTransactionDto) (map[string]interface{}, error) {
	url := a.Host + EXECUTERAWTRANSACTION
	params := map[string]interface{}{
		"RawTransaction": input.RawTransaction,
		"Signature":      input.Signature,
	}
	transactionBytes, err := util.PostRequest(url, a.Version, params)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	json.Unmarshal(transactionBytes, &data)
	return data, nil
}

//SendTransaction Broadcast a transaction.
func (a *AElfClient) SendTransaction(transaction string) (*dto.SendTransactionOutput, error) {
	url := a.Host + SENDTRANSACTION
	params := map[string]interface{}{"RawTransaction": transaction}
	transactionBytes, err := util.PostRequest(url, a.Version, params)
	if err != nil {
		return nil, err
	}
	var output = new(dto.SendTransactionOutput)
	json.Unmarshal(transactionBytes, &output)
	return output, nil
}

//CreateRawTransaction Creates an unsigned serialized transaction.
func (a *AElfClient) CreateRawTransaction(input *dto.CreateRawTransactionInput) (*dto.CreateRawTransactionOutput, error) {
	url := a.Host + RAWTRANSACTION
	params := map[string]interface{}{
		"From":           input.From,
		"MethodName":     input.MethodName,
		"Params":         input.Params,
		"RefBlockHash":   input.RefBlockHash,
		"RefBlockNumber": input.RefBlockNumber,
		"To":             input.To,
	}
	transactionBytes, err := util.PostRequest(url, a.Version, params)
	if err != nil {
		return nil, err
	}
	var output = new(dto.CreateRawTransactionOutput)
	json.Unmarshal(transactionBytes, &output)
	return output, nil
}

//SendRawTransaction Broadcast a serialized transaction.
func (a *AElfClient) SendRawTransaction(input *dto.SendRawTransactionInput) (*dto.SendRawTransactionOutput, error) {
	url := a.Host + SENDRAWTRANSACTION
	params := map[string]interface{}{
		"Transaction":       input.Transaction,
		"Signature":         input.Signature,
		"ReturnTransaction": input.ReturnTransaction,
	}
	rawTransactionBytes, err := util.PostRequest(url, a.Version, params)
	if err != nil {
		return nil, err
	}
	var rawTransaction = new(dto.SendRawTransactionOutput)
	json.Unmarshal(rawTransactionBytes, &rawTransaction)
	return rawTransaction, nil
}

//SendTransactions Broadcast volume transactions.
func (a *AElfClient) SendTransactions(rawTransactions string) (map[string]interface{}, error) {
	url := a.Host + SENDTRANSACTIONS
	params := map[string]interface{}{
		"RawTransactions": rawTransactions,
	}
	transactionsBytes, err := util.PostRequest(url, a.Version, params)
	if err != nil {
		return nil, err
	}
	var datas map[string]interface{}
	json.Unmarshal(transactionsBytes, &datas)
	return datas, nil
}

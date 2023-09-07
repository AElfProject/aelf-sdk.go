package client

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/AElfProject/aelf-sdk.go/dto"
	util "github.com/AElfProject/aelf-sdk.go/utils"
	"github.com/davecgh/go-spew/spew"
)

// GetTransactionPoolStatus Get information about the current transaction pool.
func (client *AElfClient) GetTransactionPoolStatus() (*dto.TransactionPoolStatusOutput, error) {
	url := client.Host + TRANSACTIONPOOLSTATUS
	transactionPoolBytes, err := util.GetRequest("GET", url, client.Version, nil)
	if err != nil {
		return nil, errors.New("Get Transaction Pool Status error:" + err.Error())
	}
	var transactionPool = new(dto.TransactionPoolStatusOutput)
	json.Unmarshal(transactionPoolBytes, &transactionPool)
	return transactionPool, nil
}

// GetTransactionResult Gets the result of transaction execution by the given transactionId.
func (client *AElfClient) GetTransactionResult(transactionID string) (*dto.TransactionResultDto, error) {
	url := client.Host + TRANSACTIONRESULT
	_, err := hex.DecodeString(transactionID)
	if err != nil {
		return nil, errors.New("transactionID hex to []byte error:" + err.Error())
	}
	params := map[string]interface{}{"transactionId": transactionID}
	transactionBytes, err := util.GetRequest("GET", url, client.Version, params)
	if err != nil {
		return nil, errors.New("Get Transaction Result error:" + err.Error())
	}
	var transaction = new(dto.TransactionResultDto)
	json.Unmarshal(transactionBytes, &transaction)
	return transaction, nil
}

// GetTransactionResults Get results of multiple transactions by specified blockHash.
func (client *AElfClient) GetTransactionResults(blockHash string, offset, limit int) ([]*dto.TransactionResultDto, error) {
	url := client.Host + TRANSACTIONRESULTS
	_, err := hex.DecodeString(blockHash)
	if err != nil {
		return nil, errors.New("blockHash hex to []byte error:" + err.Error())
	}
	params := map[string]interface{}{
		"blockHash": blockHash,
		"offset":    offset,
		"limit":     limit,
	}
	transactionsBytes, err := util.GetRequest("GET", url, client.Version, params)
	if err != nil {
		return nil, errors.New("Get Transaction Results error:" + err.Error())
	}
	var datas interface{}
	json.Unmarshal(transactionsBytes, &datas)
	var transactions []*dto.TransactionResultDto
	for _, d := range datas.([]interface{}) {
		var transaction = new(dto.TransactionResultDto)
		Bytes, _ := json.Marshal(d)
		json.Unmarshal(Bytes, &transaction)
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

// GetMerklePathByTransactionID Get merkle path of a transaction.
func (client *AElfClient) GetMerklePathByTransactionID(transactionID string) (*dto.MerklePathDto, error) {
	url := client.Host + MBYTRANSACTIONID
	_, err := hex.DecodeString(transactionID)
	if err != nil {
		return nil, errors.New("transactionID hex to []byte error:" + err.Error())
	}
	params := map[string]interface{}{"transactionId": transactionID}
	merkleBytes, err := util.GetRequest("GET", url, client.Version, params)
	if err != nil {
		return nil, errors.New("Get MerklePath By TransactionID error:" + err.Error())
	}
	var merkle = new(dto.MerklePathDto)
	json.Unmarshal(merkleBytes, &merkle)
	return merkle, nil
}

// ExecuteTransaction  Call a read-only method of a contract.
func (client *AElfClient) ExecuteTransaction(rawTransaction string) (string, error) {
	url := client.Host + EXECUTETRANSACTION
	params := map[string]interface{}{"RawTransaction": rawTransaction}
	transactionBytes, err := util.PostRequest(url, client.Version, params)
	if err != nil {
		return "", errors.New("Execute Transaction error:" + err.Error())
	}
	return util.BytesToString(transactionBytes), nil
}

// ExecuteRawTransaction Call a method of a contract by given serialized strings.
func (client *AElfClient) ExecuteRawTransaction(input *dto.ExecuteRawTransactionDto) (string, error) {
	url := client.Host + EXECUTERAWTRANSACTION
	params := map[string]interface{}{
		"RawTransaction": input.RawTransaction,
		"Signature":      input.Signature,
	}
	transactionBytes, err := util.PostRequest(url, client.Version, params)
	if err != nil {
		return "", errors.New("Execute RawTransaction error:" + err.Error())
	}
	//var data interface{}
	//json.Unmarshal(transactionBytes, &data)
	return util.BytesToString(transactionBytes), nil
}

// SendTransaction Broadcast a transaction.
func (client *AElfClient) SendTransaction(transaction string) (*dto.SendTransactionOutput, error) {
	url := client.Host + SENDTRANSACTION
	params := map[string]interface{}{"RawTransaction": transaction}
	transactionBytes, err := util.PostRequest(url, client.Version, params)
	if err != nil {
		return nil, errors.New("Send Transaction error:" + err.Error())
	}
	var output = new(dto.SendTransactionOutput)
	json.Unmarshal(transactionBytes, &output)
	return output, nil
}

// CreateRawTransaction Creates an unsigned serialized transaction.
func (client *AElfClient) CreateRawTransaction(input *dto.CreateRawTransactionInput) (*dto.CreateRawTransactionOutput, error) {
	url := client.Host + RAWTRANSACTION
	params := map[string]interface{}{
		"From":           input.From,
		"MethodName":     input.MethodName,
		"Params":         input.Params,
		"RefBlockHash":   input.RefBlockHash,
		"RefBlockNumber": input.RefBlockNumber,
		"To":             input.To,
	}
	transactionBytes, err := util.PostRequest(url, client.Version, params)
	if err != nil {
		return nil, errors.New("Create RawTransaction error:" + err.Error())
	}
	var output = new(dto.CreateRawTransactionOutput)
	json.Unmarshal(transactionBytes, &output)
	return output, nil
}

// SendRawTransaction Broadcast a serialized transaction.
func (client *AElfClient) SendRawTransaction(input *dto.SendRawTransactionInput) (*dto.SendRawTransactionOutput, error) {
	url := client.Host + SENDRAWTRANSACTION
	params := map[string]interface{}{
		"Transaction":       input.Transaction,
		"Signature":         input.Signature,
		"ReturnTransaction": input.ReturnTransaction,
	}
	rawTransactionBytes, err := util.PostRequest(url, client.Version, params)
	if err != nil {
		return nil, errors.New("Send RawTransaction error:" + err.Error())
	}
	var rawTransaction = new(dto.SendRawTransactionOutput)
	json.Unmarshal(rawTransactionBytes, &rawTransaction)
	return rawTransaction, nil
}

// SendTransactions Broadcast volume transactions.
func (client *AElfClient) SendTransactions(rawTransactions string) ([]interface{}, error) {
	url := client.Host + SENDTRANSACTIONS
	params := map[string]interface{}{
		"RawTransactions": rawTransactions,
	}
	transactionsBytes, err := util.PostRequest(url, client.Version, params)
	if err != nil {
		return nil, errors.New("Send Transaction error:" + err.Error())
	}
	var data interface{}
	json.Unmarshal(transactionsBytes, &data)
	var transactions []interface{}
	for _, d := range data.([]interface{}) {
		transactions = append(transactions, d)
	}
	return transactions, nil
}

func (client *AElfClient) CalculateTransactionFee(input *dto.CalculateTransactionFeeInput) (*dto.CalculateTransactionFeeOutput, error) {
	url := client.Host + CALCULATETRANSACTIONFEE
	params := map[string]interface{}{
		"RawTransaction": input.RawTransaction,
	}
	transactionFeeResult, err := util.PostRequest(url, client.Version, params)
	if err != nil {
		return nil, errors.New("CalculateTransactionFee error:" + err.Error())
	}
	var feeResult = new(dto.CalculateTransactionFeeOutput)
	json.Unmarshal(transactionFeeResult, &feeResult)
	spew.Dump("CalculateTransactionFee : ", feeResult.Success)
	return feeResult, nil

}

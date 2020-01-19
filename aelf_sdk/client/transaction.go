package client

import (
	"encoding/json"

	"aelf_sdk.go/aelf_sdk/dto"
)

//GetTransactionPoolStatus Get Transaction Pool Status  //已完成
func (a *AElfClient) GetTransactionPoolStatus() (*dto.TransactionPoolStatus, error) {
	url := a.Host + TRANSACTIONPOOLSTATUS
	transactionPoolBytes, err := GetRequest("GET", url, a.Version, nil)
	if err != nil {
		return nil, err
	}
	var transactionPool = new(dto.TransactionPoolStatus)
	json.Unmarshal(transactionPoolBytes, &transactionPool)
	return transactionPool, nil
}

//ExecuteTransaction Execute Transaction   //已完成
func (a *AElfClient) ExecuteTransaction(params map[string]interface{}) (string, error) {
	url := a.Host + EXECUTETRANSACTION
	transactionBytes, err := PostRequest(url, a.Version, params)
	if err != nil {
		return "", err
	}
	var data interface{}
	json.Unmarshal(transactionBytes, &data)
	return data.(string), nil
}

//ExecuteRawTransaction Execute Raw Transaction //已完成
func (a *AElfClient) ExecuteRawTransaction(params map[string]interface{}) (string, error) {
	url := a.Host + EXECUTERAWTRANSACTION
	transactionBytes, err := PostRequest(url, a.Version, params)
	if err != nil {
		return "", err
	}
	var data interface{}
	json.Unmarshal(transactionBytes, &data)
	return data.(string), nil
}

//CreateRawTransaction Create Raw Transaction  //已完成
func (a *AElfClient) CreateRawTransaction(params map[string]interface{}) (string, error) {
	url := a.Host + RAWTRANSACTION
	transactionBytes, err := PostRequest(url, a.Version, params)
	if err != nil {
		return "", err
	}
	var data interface{}
	json.Unmarshal(transactionBytes, &data)
	return data.(string), nil
}

//SendTransaction Send Transaction   //已完成
// func (a *AElfClient) SendTransaction(transaction string) (string, error) {
// 	url := a.Host + SENDTRANSACTION
// 	params := map[string]interface{}{"RawTransaction": transaction} //hex string
// 	data, err := PostRequest(url, a.Version, params)
// 	if err != nil {
// 		return "", err
// 	}
// 	return data.(string), nil
// }

//SendRawTransaction Send  raw Transaction  //已完成
func (a *AElfClient) SendRawTransaction(params map[string]interface{}) (*dto.SendRawTransactionOutput, error) {
	url := a.Host + SENDRAWTRANSACTION
	rawTransactionBytes, err := PostRequest(url, a.Version, params)
	if err != nil {
		return nil, err
	}
	var rawTransaction = new(dto.SendRawTransactionOutput)
	json.Unmarshal(rawTransactionBytes, &rawTransaction)
	return rawTransaction, nil
}

//GetTransactionResult Get Transaction Result
func (a *AElfClient) GetTransactionResult(transactionID string) (*dto.TransactionResultDto, error) {
	url := a.Host + TRANSACTIONRESULT
	params := map[string]interface{}{"transactionId": transactionID}
	transactionBytes, err := GetRequest("GET", url, a.Version, params)
	if err != nil {
		return nil, err
	}
	var transaction = new(dto.TransactionResultDto)
	json.Unmarshal(transactionBytes, &transaction)
	return transaction, nil
}

//GetTransactionResults Get Transaction Results  //已完成
func (a *AElfClient) GetTransactionResults(blockHash string) ([]*dto.TransactionResultDto, error) {
	url := a.Host + TRANSACTIONRESULTS
	params := map[string]interface{}{"blockHash": blockHash}
	transactionsBytes, err := GetRequest("GET", url, a.Version, params)
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

//GetMerklePathByTransactionID Get MerklePath By TransactionId  //已完成
func (a *AElfClient) GetMerklePathByTransactionID(transactionID string) (*dto.MerklePathDto, error) {
	url := a.Host + MBYTRANSACTIONID
	params := map[string]interface{}{"transactionId": transactionID}
	merkleBytes, err := GetRequest("GET", url, a.Version, params)
	if err != nil {
		return nil, err
	}
	var merkle = new(dto.MerklePathDto)
	json.Unmarshal(merkleBytes, &merkle)
	return merkle, nil
}

// type SendTransactionsInput struct {
// 	RawTransactions string   post参数
// }

//SendTransactions SendTransactions  //已完成
// func (a *AElfClient) SendTransactions(params map[string]interface{}) ([]string, error) {
// 	url := a.Host + SENDTRANSACTIONS
// 	datas, err := PostRequest(url, a.Version, params)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var transactions []string
// 	for _, dat := range datas.([]interface{}) {
// 		transactions = append(transactions, dat.(string))
// 	}
// 	return transactions, nil
// }

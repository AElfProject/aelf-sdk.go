package client

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/AElfProject/aelf-sdk.go/dto"
	"github.com/AElfProject/aelf-sdk.go/model/consts"
	client "github.com/AElfProject/aelf-sdk.go/protobuf/generated"
	util "github.com/AElfProject/aelf-sdk.go/utils"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/protobuf/proto"
)

// GetTransactionPoolStatus Get information about the current transaction pool.
func (a *AElfClient) GetTransactionPoolStatus() (*dto.TransactionPoolStatusOutput, error) {
	url := a.Host + TRANSACTIONPOOLSTATUS
	transactionPoolBytes, err := util.GetRequest("GET", url, a.Version, nil)
	if err != nil {
		return nil, errors.New("Get Transaction Pool Status error:" + err.Error())
	}
	var transactionPool = new(dto.TransactionPoolStatusOutput)
	json.Unmarshal(transactionPoolBytes, &transactionPool)
	return transactionPool, nil
}

// GetTransactionResult Gets the result of transaction execution by the given transactionId.
func (a *AElfClient) GetTransactionResult(transactionID string) (*dto.TransactionResultDto, error) {
	url := a.Host + TRANSACTIONRESULT
	_, err := hex.DecodeString(transactionID)
	if err != nil {
		return nil, errors.New("transactionID hex to []byte error:" + err.Error())
	}
	params := map[string]interface{}{"transactionId": transactionID}
	transactionBytes, err := util.GetRequest("GET", url, a.Version, params)
	if err != nil {
		return nil, errors.New("Get Transaction Result error:" + err.Error())
	}
	var transaction = new(dto.TransactionResultDto)
	json.Unmarshal(transactionBytes, &transaction)
	return transaction, nil
}

// GetTransactionResults Get results of multiple transactions by specified blockHash.
func (a *AElfClient) GetTransactionResults(blockHash string, offset, limit int) ([]*dto.TransactionResultDto, error) {
	url := a.Host + TRANSACTIONRESULTS
	_, err := hex.DecodeString(blockHash)
	if err != nil {
		return nil, errors.New("blockHash hex to []byte error:" + err.Error())
	}
	params := map[string]interface{}{
		"blockHash": blockHash,
		"offset":    offset,
		"limit":     limit,
	}
	transactionsBytes, err := util.GetRequest("GET", url, a.Version, params)
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
func (a *AElfClient) GetMerklePathByTransactionID(transactionID string) (*dto.MerklePathDto, error) {
	url := a.Host + MBYTRANSACTIONID
	_, err := hex.DecodeString(transactionID)
	if err != nil {
		return nil, errors.New("transactionID hex to []byte error:" + err.Error())
	}
	params := map[string]interface{}{"transactionId": transactionID}
	merkleBytes, err := util.GetRequest("GET", url, a.Version, params)
	if err != nil {
		return nil, errors.New("Get MerklePath By TransactionID error:" + err.Error())
	}
	var merkle = new(dto.MerklePathDto)
	json.Unmarshal(merkleBytes, &merkle)
	return merkle, nil
}

// ExecuteTransaction  Call a read-only method of a contract.
func (a *AElfClient) ExecuteTransaction(rawTransaction string) (string, error) {
	url := a.Host + EXECUTETRANSACTION
	params := map[string]interface{}{"RawTransaction": rawTransaction}
	transactionBytes, err := util.PostRequest(url, a.Version, params)
	if err != nil {
		return "", errors.New("Execute Transaction error:" + err.Error())
	}
	return util.BytesToString(transactionBytes), nil
}

// ExecuteRawTransaction Call a method of a contract by given serialized strings.
func (a *AElfClient) ExecuteRawTransaction(input *dto.ExecuteRawTransactionDto) (string, error) {
	url := a.Host + EXECUTERAWTRANSACTION
	params := map[string]interface{}{
		"RawTransaction": input.RawTransaction,
		"Signature":      input.Signature,
	}
	transactionBytes, err := util.PostRequest(url, a.Version, params)
	if err != nil {
		return "", errors.New("Execute RawTransaction error:" + err.Error())
	}
	//var data interface{}
	//json.Unmarshal(transactionBytes, &data)
	return util.BytesToString(transactionBytes), nil
}

// SendTransaction Broadcast a transaction.
func (a *AElfClient) SendTransaction(transaction string) (*dto.SendTransactionOutput, error) {
	url := a.Host + SENDTRANSACTION
	params := map[string]interface{}{"RawTransaction": transaction}
	transactionBytes, err := util.PostRequest(url, a.Version, params)
	if err != nil {
		return nil, errors.New("Send Transaction error:" + err.Error())
	}
	var output = new(dto.SendTransactionOutput)
	json.Unmarshal(transactionBytes, &output)
	return output, nil
}

// CreateRawTransaction Creates an unsigned serialized transaction.
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
		return nil, errors.New("Create RawTransaction error:" + err.Error())
	}
	var output = new(dto.CreateRawTransactionOutput)
	json.Unmarshal(transactionBytes, &output)
	return output, nil
}

// SendRawTransaction Broadcast a serialized transaction.
func (a *AElfClient) SendRawTransaction(input *dto.SendRawTransactionInput) (*dto.SendRawTransactionOutput, error) {
	url := a.Host + SENDRAWTRANSACTION
	params := map[string]interface{}{
		"Transaction":       input.Transaction,
		"Signature":         input.Signature,
		"ReturnTransaction": input.ReturnTransaction,
	}
	rawTransactionBytes, err := util.PostRequest(url, a.Version, params)
	if err != nil {
		return nil, errors.New("Send RawTransaction error:" + err.Error())
	}
	var rawTransaction = new(dto.SendRawTransactionOutput)
	json.Unmarshal(rawTransactionBytes, &rawTransaction)
	return rawTransaction, nil
}

// SendTransactions Broadcast volume transactions.
func (a *AElfClient) SendTransactions(rawTransactions string) ([]interface{}, error) {
	url := a.Host + SENDTRANSACTIONS
	params := map[string]interface{}{
		"RawTransactions": rawTransactions,
	}
	transactionsBytes, err := util.PostRequest(url, a.Version, params)
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

func (a *AElfClient) CalculateTransactionFee(input *dto.CalculateTransactionFeeInput) (*dto.CalculateTransactionFeeOutput, error) {
	url := a.Host + CALCULATETRANSACTIONFEE
	params := map[string]interface{}{
		"RawTransaction": input.RawTransaction,
	}
	transactionFeeResult, err := util.PostRequest(url, a.Version, params)
	if err != nil {
		return nil, errors.New("CalculateTransactionFee error:" + err.Error())
	}
	var feeResult = new(dto.CalculateTransactionFeeOutput)
	json.Unmarshal(transactionFeeResult, &feeResult)
	spew.Dump("CalculateTransactionFee : ", feeResult.Success)
	return feeResult, nil

}

func (a *AElfClient) GetTransferred(txId string) []*client.Transferred {
	transffereds := make([]*client.Transferred, 0)
	result, err := a.GetTransactionResult(txId)
	if err != nil || len(result.Logs) == 0 {
		return transffereds
	}

	contractAddr, _ := a.GetContractAddressByName(consts.TokenContractSystemName)

	for _, log := range result.Logs {
		if log.Name == "Transferred" && log.Address == contractAddr {
			transferred := new(client.Transferred)
			if nonIndexedBytes, err := util.Base64DecodeBytes(log.NonIndexed); err == nil {
				proto.Unmarshal(nonIndexedBytes, transferred)
			}
			if fromBytes, err := util.Base64DecodeBytes(log.Indexed[0]); err == nil {
				temp := new(client.Transferred)
				proto.Unmarshal(fromBytes, temp)
				transferred.From = temp.From
			}
			if toBytes, err := util.Base64DecodeBytes(log.Indexed[1]); err == nil {
				temp := new(client.Transferred)
				proto.Unmarshal(toBytes, temp)
				transferred.To = temp.To
			}
			if symbolBytes, err := util.Base64DecodeBytes(log.Indexed[2]); err == nil {
				temp := new(client.Transferred)
				proto.Unmarshal(symbolBytes, temp)
				transferred.Symbol = temp.Symbol
			}
			transffereds = append(transffereds, transferred)
		}
	}

	return transffereds
}

func (a *AElfClient) GetCrossChainTransferred(txId string) []*client.CrossChainTransferred {
	crossChainTransferreds := make([]*client.CrossChainTransferred, 0)
	result, err := a.GetTransactionResult(txId)
	if err != nil || len(result.Logs) == 0 {
		return crossChainTransferreds
	}

	contractAddr, _ := a.GetContractAddressByName(consts.TokenContractSystemName)

	for _, log := range result.Logs {
		if log.Name == "CrossChainTransferred" && log.Address == contractAddr {
			crossChainTransferred := new(client.CrossChainTransferred)
			if nonIndexedBytes, err := util.Base64DecodeBytes(log.NonIndexed); err == nil {
				proto.Unmarshal(nonIndexedBytes, crossChainTransferred)
			}
			crossChainTransferreds = append(crossChainTransferreds, crossChainTransferred)
		}
	}

	return crossChainTransferreds
}

func (a *AElfClient) GetCrossChainReceived(txId string) []*client.CrossChainReceived {
	crossChainReceiveds := make([]*client.CrossChainReceived, 0)
	result, err := a.GetTransactionResult(txId)
	if err != nil || len(result.Logs) == 0 {
		return crossChainReceiveds
	}

	contractAddr, _ := a.GetContractAddressByName(consts.TokenContractSystemName)

	for _, log := range result.Logs {
		if log.Name == "CrossChainReceived" && log.Address == contractAddr {
			crossChainReceived := new(client.CrossChainReceived)
			if nonIndexedBytes, err := util.Base64DecodeBytes(log.NonIndexed); err == nil {
				proto.Unmarshal(nonIndexedBytes, crossChainReceived)
			}
			crossChainReceiveds = append(crossChainReceiveds, crossChainReceived)
		}
	}

	return crossChainReceiveds
}

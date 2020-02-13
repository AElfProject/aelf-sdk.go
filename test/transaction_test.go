package test

import (
	"encoding/hex"
	"testing"

	"aelf_sdk.go/client"
	"aelf_sdk.go/dto"
	util "aelf_sdk.go/utils"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

var ContractMethodName = "GetContractAddressByName"
var ContractAddress, _ = aelf.GetGenesisContractAddress()

func TestGetTransactionResult(t *testing.T) {
	var isTransactions = true
	height, err := aelf.GetBlockHeight()
	block, err := aelf.GetBlockByHeight(int(height), isTransactions)
	assert.NoError(t, err)
	transactionID := block.Body.Transactions[0]
	transactionResult, err := aelf.GetTransactionResult(transactionID)
	assert.NoError(t, err)
	spew.Dump("Get Transaction Result", transactionResult)
}

func TestGetTransactionResults(t *testing.T) {
	var isTransactions = true
	height, err := aelf.GetBlockHeight()
	block, err := aelf.GetBlockByHeight(int(height), isTransactions)
	assert.NoError(t, err)
	blockHash := block.BlockHash
	transactionResults, err := aelf.GetTransactionResults(blockHash, 0, 10)
	assert.NoError(t, err)
	spew.Dump("Get Transaction Results", transactionResults)
}

func TestGetMerklePathByTransactionID(t *testing.T) {
	var isTransactions = true
	height, err := aelf.GetBlockHeight()
	block, err := aelf.GetBlockByHeight(int(height), isTransactions)
	assert.NoError(t, err)
	transactionID := block.Body.Transactions[0]
	merklePath, err := aelf.GetMerklePathByTransactionID(transactionID)
	assert.NoError(t, err)
	spew.Dump("Get Merkle Path By TransactionID Result", merklePath)
}

func TestGetTransactionPoolStatus(t *testing.T) {
	poolStatus, err := aelf.GetTransactionPoolStatus()
	assert.NoError(t, err)
	spew.Dump("Get TransactionPool Status Result", poolStatus)
}

func TestCreateRawTransaction(t *testing.T) {
	chainStatus, err := aelf.GetChainStatus()
	assert.NoError(t, err)
	var input = &dto.CreateRawTransactionInput{
		From:           _address,
		To:             ContractAddress,
		MethodName:     ContractMethodName,
		Params:         util.ParamsToString("AElf.ContractNames.Token"),
		RefBlockHash:   chainStatus.BestChainHash,
		RefBlockNumber: chainStatus.BestChainHeight,
	}
	result, err := aelf.CreateRawTransaction(input)
	assert.NoError(t, err)
	spew.Dump("Create RawTransaction result", result)
}

func TestSendRawTransaction(t *testing.T) {
	chainStatus, err := aelf.GetChainStatus()
	assert.NoError(t, err)
	var input = &dto.CreateRawTransactionInput{
		From:           _address,
		To:             ContractAddress,
		MethodName:     ContractMethodName,
		RefBlockNumber: chainStatus.BestChainHeight,
		RefBlockHash:   chainStatus.BestChainHash,
		Params:         util.ParamsToString("AElf.ContractNames.Token"),
	}
	createRaw, err := aelf.CreateRawTransaction(input)
	assert.NoError(t, err)
	spew.Dump("Create Raw Transaction result", createRaw)
	rawTransactionBytes, err := hex.DecodeString(createRaw.RawTransaction)
	signature, _ := client.GetSignatureWithPrivateKey(aelf.PrivateKey, rawTransactionBytes)
	var sendRawinput = &dto.SendRawTransactionInput{
		Transaction:       createRaw.RawTransaction,
		Signature:         signature,
		ReturnTransaction: true,
	}
	executeRawresult, err := aelf.SendRawTransaction(sendRawinput)
	assert.NoError(t, err)
	spew.Dump("Send Raw Transaction result", executeRawresult)
}

func TestExecuteRawTransaction(t *testing.T) {
	chainStatus, err := aelf.GetChainStatus()
	assert.NoError(t, err)
	var input = &dto.CreateRawTransactionInput{
		From:           _address,
		To:             ContractAddress,
		MethodName:     ContractMethodName,
		RefBlockNumber: chainStatus.BestChainHeight,
		RefBlockHash:   chainStatus.BestChainHash,
		Params:         util.ParamsToString("AElf.ContractNames.Consensus"),
	}
	createRaw, err := aelf.CreateRawTransaction(input)
	assert.NoError(t, err)
	spew.Dump("Create Raw Transaction result", createRaw)
	rawTransactionBytes, err := hex.DecodeString(createRaw.RawTransaction)
	signature, _ := client.GetSignatureWithPrivateKey(aelf.PrivateKey, rawTransactionBytes)
	var executeRawinput = &dto.ExecuteRawTransactionDto{
		RawTransaction: createRaw.RawTransaction,
		Signature:      signature,
	}
	executeRawresult, err := aelf.ExecuteRawTransaction(executeRawinput)
	assert.NoError(t, err)
	spew.Dump("Execute RawTransaction result", executeRawresult)
}

func TestSendTransaction(t *testing.T) {
	fromAddress := _address
	toAddress := ContractAddress
	methodName := ContractMethodName
	params := util.GetBytesSha256("AElf.ContractNames.Vote")
	transaction, err := aelf.CreateTransaction(fromAddress, toAddress, methodName, params)
	assert.NoError(t, err)
	signature, err := aelf.SignTransaction(aelf.PrivateKey, transaction)
	transaction.Signature = signature
	assert.NoError(t, err)
	transactionByets, _ := proto.Marshal(transaction)
	sendResult, err := aelf.SendTransaction(hex.EncodeToString(transactionByets))
	assert.NoError(t, err)
	spew.Dump("Send Transaction result", sendResult)
}

func TestExecuteTransaction(t *testing.T) {
	fromAddress := _address
	toAddress := ContractAddress
	methodName := ContractMethodName
	params := util.GetBytesSha256("AElf.ContractNames.TokenConverter")
	transaction, err := aelf.CreateTransaction(fromAddress, toAddress, methodName, params)
	assert.NoError(t, err)
	signature, err := aelf.SignTransaction(aelf.PrivateKey, transaction)
	transaction.Signature = signature
	transactionByets, _ := proto.Marshal(transaction)
	executeresult, err := aelf.ExecuteTransaction(hex.EncodeToString(transactionByets))
	assert.NoError(t, err)
	spew.Dump("Execute Transaction result", executeresult)
}

func TestSendTransctions(t *testing.T) {
	fromAddress := _address
	toAddress := ContractAddress
	methodName := ContractMethodName
	param1 := util.GetBytesSha256("AElf.ContractNames.Vote")
	param2 := util.GetBytesSha256("AElf.ContractNames.Token")
	var parameters [][]byte
	parameters = append(parameters, param1)
	parameters = append(parameters, param2)
	for _, param := range parameters {
		transaction, err := aelf.CreateTransaction(fromAddress, toAddress, methodName, param)
		assert.NoError(t, err)
		signature, err := aelf.SignTransaction(aelf.PrivateKey, transaction)
		transaction.Signature = signature
		assert.NoError(t, err)
		transactionByets, _ := proto.Marshal(transaction)
		results, err := aelf.SendTransactions(hex.EncodeToString(transactionByets))
		assert.NoError(t, err)
		spew.Dump("Send Transactions result", results)
	}
}

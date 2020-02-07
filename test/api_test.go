package test

import (
	"encoding/hex"
	"testing"

	"aelf_sdk.go/aelf_sdk/client"
	"aelf_sdk.go/aelf_sdk/dto"
	util "aelf_sdk.go/aelf_sdk/utils"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/proto"
	secp256 "github.com/skycoin/skycoin/src/cipher/secp256k1-go"
	"github.com/stretchr/testify/assert"
)

var aelf = client.AElfClient{
	Host:       "http://127.0.0.1:8000",
	Version:    "1.0",
	PrivateKey: "680afd630d82ae5c97942c4141d60b8a9fedfa5b2864fca84072c17ee1f72d9d",
}

var ContractMethodName = "GetContractAddressByName"
var TestAddress = "127.0.0.1:6801"
var ContractAddress, _ = aelf.GetGenesisContractAddress()
var _address = aelf.GetAddressFromPrivateKey(aelf.PrivateKey, false)

func TestGetAddressFromPubKey(t *testing.T) {
	privateKeyBytes, _ := hex.DecodeString(aelf.PrivateKey)
	pubkeyBytes := secp256.UncompressedPubkeyFromSeckey(privateKeyBytes)
	pubKeyAddress := aelf.GetAddressFromPubKey(hex.EncodeToString(pubkeyBytes))
	assert.Equal(t, _address, pubKeyAddress)
	spew.Dump("Get Address From Public Key", pubKeyAddress)
}

func TestGetChainStatus(t *testing.T) {
	chainStatus, err := aelf.GetChainStatus()
	assert.NoError(t, err)
	spew.Dump("Get Chain Status Result", chainStatus)
}

func TestGetChainID(t *testing.T) {
	chainID, err := aelf.GetChainID()
	assert.NoError(t, err)
	spew.Dump("Get ChainID Result", chainID)
}

func TestGetBlockHeight(t *testing.T) {
	height, err := aelf.GetBlockHeight()
	assert.NoError(t, err)
	assert.True(t, height > 0)
}

func TestGetBlockByHeightAndByHash(t *testing.T) {
	//Test GetBlockByHeight
	var includeTransactions = true
	height, err := aelf.GetBlockHeight()
	spew.Dump(">>>>>>>height", height)
	// s := strconv.Itoa(int(height))
	// s64, _ := strconv.ParseInt(s, 10, 64)
	spew.Dump(">>>>>>>heights64s64", int64(height))

	HeightBlock, err := aelf.GetBlockByHeight(int(height), includeTransactions)
	assert.NoError(t, err)
	spew.Dump("Get Block ByHeight Result", HeightBlock)

	//Test GetBlockByHash
	// blockHash := HeightBlock.BlockHash
	// byHashBlock, err := aelf.GetBlockByHash(blockHash, includeTransactions)
	// assert.NoError(t, err)
	// spew.Dump("Get Block ByHash Result", byHashBlock)
}

func TestTransactionResult(t *testing.T) {
	//Test GetTransactionResult
	var isTransactions = true
	height, err := aelf.GetBlockHeight()
	block, err := aelf.GetBlockByHeight(int(height), isTransactions)
	assert.NoError(t, err)

	transactionID := block.Body.Transactions[0]
	transactionResult, err := aelf.GetTransactionResult(transactionID)
	assert.NoError(t, err)
	spew.Dump("Get Transaction Result", transactionResult)

	// Test GetTransactionResults
	blockHash := block.BlockHash
	transactionResults, err := aelf.GetTransactionResults(blockHash, 0, 10)
	assert.NoError(t, err)
	spew.Dump("Get Transaction Results", transactionResults)

	//Test GetMerklePathByTransactionID
	merklePath, err := aelf.GetMerklePathByTransactionID(transactionID)
	assert.NoError(t, err)
	spew.Dump("Get Merkle Path By TransactionID Result", merklePath)
}

func TestNetworkApi(t *testing.T) {
	//Test GetNetworkInfo
	netWorkInfo, err := aelf.GetNetworkInfo()
	assert.NoError(t, err)
	spew.Dump("Get Network Info Result", netWorkInfo)

	//Test AddPeer
	addPeer, err := aelf.AddPeer(TestAddress)
	assert.NoError(t, err)
	assert.True(t, addPeer == true)

	//Test RemovePeer
	removePeer, err := aelf.RemovePeer(TestAddress)
	assert.NoError(t, err)
	assert.True(t, removePeer == true)

	//Test GetPeers
	peers, err := aelf.GetPeers(true)
	assert.NoError(t, err)
	spew.Dump("Get Peers Result", peers)
}

func TestGetTransactionPoolStatus(t *testing.T) {
	poolStatus, err := aelf.GetTransactionPoolStatus()
	assert.NoError(t, err)
	spew.Dump("Get TransactionPool Status Result", poolStatus)
}

func TestGetTaskQueueStatus(t *testing.T) {
	taskQueueStatus, err := aelf.GetTaskQueueStatus()
	assert.NoError(t, err)
	spew.Dump("Get Task Queue Status Result", taskQueueStatus)
}

func TestClient(t *testing.T) {
	//Test IsConnected
	isConnected := aelf.IsConnected()
	assert.True(t, isConnected == true)

	//Test GetGenesisContractAddress
	ContractAddress, err := aelf.GetGenesisContractAddress()
	assert.NoError(t, err)
	spew.Dump("Get Genesis Contract Address Result", ContractAddress)
}

func TestGetContractFileDescriptorSet(t *testing.T) {
	ContractAddress, err := aelf.GetGenesisContractAddress()
	assert.NoError(t, err)
	contractFile, err := aelf.GetContractFileDescriptorSet(ContractAddress)
	assert.NoError(t, err)
	spew.Dump("Get Contract File Descriptor Set Result", contractFile)
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

func TestGetContractAddressByName(t *testing.T) {
	contractAddress, err := aelf.GetContractAddressByName("AElf.ContractNames.Token")
	assert.NoError(t, err)
	spew.Dump("Get ContractAddress By Name Result", contractAddress)
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

func TestGetFormattedAddress(t *testing.T) {
	formattedAddress, err := aelf.GetFormattedAddress(_address)
	assert.NoError(t, err)
	spew.Dump("Get Formatted Address result", formattedAddress, err)

	privateKeyaddress := aelf.GetAddressFromPrivateKey(aelf.PrivateKey, false)
	assert.Equal(t, formattedAddress, ("ELF_" + privateKeyaddress + "_AELF"))
}

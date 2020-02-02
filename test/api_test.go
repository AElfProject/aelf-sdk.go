package test

import (
	"encoding/hex"
	"testing"

	"aelf_sdk.go/aelf_sdk/client"

	"aelf_sdk.go/aelf_sdk/dto"
	pt "aelf_sdk.go/aelf_sdk/proto"

	util "aelf_sdk.go/aelf_sdk/utils"
	secp256 "github.com/skycoin/skycoin/src/cipher/secp256k1-go"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

var aelf = client.AElfClient{
	Host:       "http://127.0.0.1:8000",
	Version:    "1.0",
	PrivateKey: "680afd630d82ae5c97942c4141d60b8a9fedfa5b2864fca84072c17ee1f72d9d",
	PublicKey:  "",
}

var ContractMethodName = "GetContractAddressByName"
var TestAddress = "127.0.0.1:6801"
var ContractAddress, _ = aelf.GetGenesisContractAddress()
var _address = aelf.GetAddressFromPrivateKey(aelf.PrivateKey, false)

func TestChainApi(t *testing.T) {
	chainStatus, err := aelf.GetChainStatus()
	if err != nil {
		t.Error("GetChainStatus error", err)
	}
	spew.Dump("get chain status", chainStatus)
	chainID, err := aelf.GetChainID()
	if err != nil {
		t.Error("GetChainID error", err)
	}
	spew.Dump("get chain ID", chainID)
}

func TestGetBlockApi(t *testing.T) {

	//Test GetBlockHeight
	height, err := aelf.GetBlockHeight()
	if err != nil || height == 0 {
		t.Error("GetBlockHeight err", err)
	}
	spew.Dump("# get_block_height", height)

	//Test GetBlockByHeight
	var isTransactions = true
	var blockHeight = 1
	byHeightBlock, err := aelf.GetBlockByHeight(blockHeight, isTransactions)
	if err != nil {
		t.Error("GetBlockByHeight err", err)
	}
	spew.Dump("get_block_by_height", byHeightBlock)

	//Test GetBlockByHash
	blockHash := byHeightBlock.BlockHash
	byHashBlock, err := aelf.GetBlockByHash(blockHash, isTransactions)
	if err != nil {
		t.Error("Get Block By Height err", err)
	}
	spew.Dump("# get_block_by_hash", byHashBlock)
}

func TestTransactionResultApi(t *testing.T) {
	//Test GetTransactionResult
	var isTransactions = true
	var blockHeight = 1
	block, err := aelf.GetBlockByHeight(blockHeight, isTransactions)
	if err != nil || block == nil {
		t.Error("Get Block By Height error", err)
	}

	transactionID := block.Body.Transactions[0]
	transactionResult, err := aelf.GetTransactionResult(transactionID)
	if err != nil {
		t.Error("Get Transaction Result error", err)
	}
	spew.Dump("# Get Transaction Result", transactionResult)

	// Test GetTransactionResults
	blockHash := block.BlockHash
	transactionResults, err := aelf.GetTransactionResults(blockHash, 0, 10)
	if err != nil {
		t.Error("Get Transaction Results error", err)
	}
	spew.Dump("# Get Transaction Results", transactionResults)

	//Test GetMerklePathByTransactionID
	merklePath, err := aelf.GetMerklePathByTransactionID(transactionID)
	if err != nil {
		t.Error("Get merkle Path  error", err)
	}
	spew.Dump("Get Merkle Path By TransactionID", merklePath)
}

func TestNetworkApi(t *testing.T) {

	//Test GetNetworkInfo
	netWorkInfo, err := aelf.GetNetworkInfo()
	if err != nil {
		t.Error("get net work info error", err)
	}
	spew.Dump("# get_network_info", netWorkInfo)

	//Test AddPeer
	addPeer, err := aelf.AddPeer(TestAddress)
	if err != nil {
		t.Error("add peer error", err)
	}
	spew.Dump("# add_peer", addPeer)

	//Test RemovePeer
	removePeer, err := aelf.RemovePeer(TestAddress)
	if err != nil {
		t.Error("remove peer error", err)
	}
	spew.Dump("# remove_peer", removePeer)

	//Test GetPeers
	peers, err := aelf.GetPeers(true)
	if err != nil {
		t.Error("get peers error", err)
	}
	spew.Dump("# get peers", peers)
}

func TestTransactionPoolApi(t *testing.T) {
	poolStatus, err := aelf.GetTransactionPoolStatus()
	if err != nil {
		t.Error("Get Transaction Pool Status error", err)
	}
	spew.Dump("", poolStatus)
}

func TestTaskQueueApi(t *testing.T) {
	taskQueueStatus, err := aelf.GetTaskQueueStatus()
	if err != nil || len(taskQueueStatus) == 0 {
		t.Error("get task Queue Status errro", err)
	}
	spew.Dump("# get task Queue  status", taskQueueStatus)
}

func TestCurrentRoundInformation(t *testing.T) {
	roundInfo, err := aelf.GetCurrentRoundInformation()
	if err != nil {
		t.Error("get Current Round Information error", err)
	}
	spew.Dump("get Current Round Information", roundInfo)
}

func TestClient(t *testing.T) {
	//Test IsConnected
	isConnected := aelf.IsConnected()
	if !isConnected {
		t.Error("connect faild")
	}
	assert.True(t, true, isConnected == true)

	//Test GetGenesisContractAddress
	if ContractAddress != "" {
		spew.Dump("Get Genesis Contract Address", ContractAddress)
	}
}

func TestGetContractFileDescriptorSet(t *testing.T) {
	if ContractAddress == "" {
		t.Error("Get Genesis ContractAddress error")
	}
	contractFile, err := aelf.GetContractFileDescriptorSet(ContractAddress)
	if err != nil {
		t.Error("Get Contract File Descriptor Set error", err)
	}
	spew.Dump("Get Contract File Descriptor Set", contractFile)
}

func TestCreateRawTransaction(t *testing.T) {
	chainStatus, err := aelf.GetChainStatus()
	if err != nil {
		t.Error("Get Chain Status error", err)
	}
	var input = &dto.CreateRawTransactionInput{
		From:           _address,
		To:             ContractAddress,
		MethodName:     ContractMethodName,
		Params:         util.ParamsToString("AElf.ContractNames.Token"),
		RefBlockHash:   chainStatus.BestChainHash,
		RefBlockNumber: chainStatus.BestChainHeight,
	}
	result, err := aelf.CreateRawTransaction(input)
	if err != nil || result.RawTransaction != "" {
		t.Error("Create RawTransaction error", err)
	}
	spew.Dump("Create RawTransaction result", result)
}

func TestGetAddressFromPubKey(t *testing.T) {
	var pubkeyBytes []byte
	pubkeyBytes = secp256.UncompressedPubkeyFromSeckey(util.StringTo32Bytes(aelf.PrivateKey))
	pubKeyAddress := aelf.GetAddressFromPubKey(string(pubkeyBytes))
	spew.Dump("Get Address From Public Key", pubKeyAddress)
}

func TestSendTransctions(t *testing.T) {
	fromAddress := _address
	toAddress := ContractAddress
	methodName := ContractMethodName
	param1 := util.GetBytesSha256("AElf.ContractNames.Vote")
	param2 := util.GetBytesSha256("AElf.ContractNames.Token")
	var parameters [][]byte
	var transaction = new(pt.Transaction)
	parameters = append(parameters, param1)
	parameters = append(parameters, param2)
	for _, param := range parameters {
		transaction, _ = aelf.CreateTransaction(fromAddress, toAddress, methodName, param)
		transactionBytes := util.SerializeToBytes(transaction)
		results, err := aelf.SendTransactions(hex.EncodeToString(transactionBytes))
		if err != nil {
			t.Error("Send Transactions  error", err)
		}
		spew.Dump("Send Transactions result", results)
	}
}

func TestSendTransaction(t *testing.T) {
	fromAddress := _address
	toAddress := ContractAddress
	methodName := ContractMethodName
	params := util.GetBytesSha256("AElf.ContractNames.Vote")
	transaction, err := aelf.CreateTransaction(fromAddress, toAddress, methodName, params)
	if err != nil {
		t.Error("Create Transaction error", err)
	}
	signature, err := aelf.SignTransaction(aelf.PrivateKey, transaction)
	if err != nil {
		t.Error("Sign Transaction with private key error", err)
	}
	transaction.Signature = signature

	transactionByets := util.SerializeToBytes(transaction)
	rawTransaction := hex.EncodeToString(transactionByets)
	sendResult, err := aelf.SendTransaction(rawTransaction)
	if err != nil {
		t.Error("Send Transaction error", err)
	}
	spew.Dump("Send Transaction result", sendResult.TransactionID, err)
}

func TestExecuteTransaction(t *testing.T) {
	fromAddress := _address
	toAddress := ContractAddress
	methodName := ContractMethodName
	params := util.GetBytesSha256("AElf.ContractNames.TokenConverter")
	transaction, err := aelf.CreateTransaction(fromAddress, toAddress, methodName, params)
	if err != nil {
		t.Error("create Transaction error", err)
	}
	spew.Dump("create Transaction result", transaction)
	signature, err := aelf.SignTransaction(aelf.PrivateKey, transaction)
	signatureBytes, _ := hex.DecodeString(string(signature))
	transaction.Signature = signatureBytes
	RawTransaction := hex.EncodeToString(util.SerializeToBytes(transaction))
	executeresult, err := aelf.ExecuteTransaction(RawTransaction)
	if err != nil {
		t.Error("Execute Transaction error", err)
	}
	spew.Dump("Execute Transaction result", executeresult)
}

func TestSendRawTransaction(t *testing.T) {
	chainStatus, err := aelf.GetChainStatus()
	if err != nil {
		t.Error("Get Chain Status error", err)
	}
	var input = &dto.CreateRawTransactionInput{
		From:           _address,
		To:             ContractAddress,
		MethodName:     ContractMethodName,
		RefBlockNumber: chainStatus.BestChainHeight,
		RefBlockHash:   chainStatus.BestChainHash,
		Params:         util.ParamsToString("AElf.ContractNames.Token"),
	}
	createRaw, err := aelf.CreateRawTransaction(input)
	var sendRawinput = &dto.SendRawTransactionInput{
		Transaction:       createRaw.RawTransaction,
		Signature:         "6935fca171452304204121397ec5917aca44aad0e7e478fed2a4a02b8f39facb0de7c752148c4fbc6366d63a32be5b824a89830bd5a3336b692be41029dabde501",
		ReturnTransaction: true,
	}
	executeRawresult, _ := aelf.SendRawTransaction(sendRawinput)
	spew.Dump("Send Raw Transaction result", executeRawresult)
}

func TestExecuteRawTransaction(t *testing.T) {
	chainStatus, err := aelf.GetChainStatus()
	if err != nil {
		t.Error("Get Chain Status error", err)
	}
	var input = &dto.CreateRawTransactionInput{
		From:           _address,
		To:             ContractAddress,
		MethodName:     ContractMethodName,
		RefBlockNumber: chainStatus.BestChainHeight,
		RefBlockHash:   chainStatus.BestChainHash,
		Params:         util.ParamsToString("AElf.ContractNames.Consensus"),
	}
	createRaw, err := aelf.CreateRawTransaction(input)
	spew.Dump("Create Raw Transaction result", createRaw, err)
	rawTransactionBytes, err := hex.DecodeString(createRaw.RawTransaction)
	signature := client.GetSignatureWithPrivateKey(aelf.PrivateKey, rawTransactionBytes)

	var executeRawinput = &dto.ExecuteRawTransactionDto{
		RawTransaction: createRaw.RawTransaction,
		Signature:      signature,
	}
	executeRawresult, err := aelf.ExecuteRawTransaction(executeRawinput)
	spew.Dump("Execute RawTransaction result", executeRawresult)
}

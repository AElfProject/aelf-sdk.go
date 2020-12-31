package test

import (
	"encoding/base64"
	"encoding/hex"
	"strings"
	"testing"
	"time"

	"github.com/AElfProject/aelf-sdk.go/client"
	"github.com/AElfProject/aelf-sdk.go/dto"
	pb "github.com/AElfProject/aelf-sdk.go/protobuf/generated"
	"github.com/AElfProject/aelf-sdk.go/utils"
	util "github.com/AElfProject/aelf-sdk.go/utils"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

var ContractMethodName = "GetContractAddressByName"
var ContractAddress, _ = aelf.GetGenesisContractAddress()

func TestGetTransactionResult(t *testing.T) {
	var isTransactions = true
	height, err := aelf.GetBlockHeight()
	block, err := aelf.GetBlockByHeight(height, isTransactions)
	assert.NoError(t, err)
	transactionID := block.Body.Transactions[0]
	transactionResult, err := aelf.GetTransactionResult(transactionID)
	assert.NoError(t, err)
	assert.Equal(t, transactionID, transactionResult.TransactionId)
	assert.Equal(t, "MINED", transactionResult.Status)
	assert.Equal(t, block.Header.Height, transactionResult.BlockNumber)
	assert.Equal(t, block.BlockHash, transactionResult.BlockHash)
	assert.NotEmpty(t, transactionResult.Bloom)
	assert.NotEmpty(t, transactionResult.Transaction)
	//spew.Dump("Get Transaction Result", transactionResult)
}

func TestGetTransactionResults(t *testing.T) {
	var isTransactions = true
	block, err := aelf.GetBlockByHeight(1, isTransactions)
	assert.NoError(t, err)
	transactionResults, err := aelf.GetTransactionResults(block.BlockHash, 0, 10)
	assert.NoError(t, err)
	assert.Equal(t, 10, len(transactionResults))
	for _, txResult := range transactionResults {
		assert.Equal(t, "MINED", txResult.Status)
		assert.Equal(t, block.Header.Height, txResult.BlockNumber)
		assert.Equal(t, block.BlockHash, txResult.BlockHash)
		assert.NotEmpty(t, txResult.Bloom)
		assert.NotEmpty(t, txResult.Transaction)
	}
	//spew.Dump("Get Transaction Results", transactionResults)
}

func TestGetMerklePathByTransactionID(t *testing.T) {
	var isTransactions = true
	height, err := aelf.GetBlockHeight()
	block, err := aelf.GetBlockByHeight(height, isTransactions)
	assert.NoError(t, err)
	transactionID := block.Body.Transactions[0]
	merklePath, err := aelf.GetMerklePathByTransactionID(transactionID)
	assert.NoError(t, err)
	assert.NotEmpty(t, merklePath)
	//spew.Dump("Get Merkle Path By TransactionID Result", merklePath)
}

func TestGetTransactionPoolStatus(t *testing.T) {
	poolStatus, err := aelf.GetTransactionPoolStatus()
	assert.NoError(t, err)
	spew.Dump("Get TransactionPool Status Result", poolStatus)
}

func TestCreateRawTransaction(t *testing.T) {
	chainStatus, err := aelf.GetChainStatus()
	assert.NoError(t, err)
	params := &pb.Hash{
		Value: utils.GetBytesSha256("AElf.ContractNames.Token"),
	}
	paramsByte, _ := protojson.Marshal(params)
	var input = &dto.CreateRawTransactionInput{
		From:           _address,
		To:             ContractAddress,
		MethodName:     ContractMethodName,
		Params:         string(paramsByte),
		RefBlockHash:   chainStatus.BestChainHash,
		RefBlockNumber: chainStatus.BestChainHeight,
	}
	result, err := aelf.CreateRawTransaction(input)
	assert.NoError(t, err)
	spew.Dump("Create RawTransaction result", result)
}

func TestSendRaßwTransaction(t *testing.T) {
	chainStatus, err := aelf.GetChainStatus()
	tokenContractAddress, _ := aelf.GetContractAddressByName("AElf.ContractNames.Token")
	userKeyPairInfo := aelf.GenerateKeyPairInfo()
	toAddress, _ := util.Base58StringToAddress(userKeyPairInfo.Address)
	params := &pb.TransferInput{
		To:     toAddress,
		Symbol: "ELF",
		Amount: 1000000000,
		Memo:   "transfer in test",
	}

	paramsByte, _ := protojson.Marshal(params)
	var input = &dto.CreateRawTransactionInput{
		From:           _address,
		To:             tokenContractAddress,
		MethodName:     "Transfer",
		RefBlockNumber: chainStatus.BestChainHeight,
		RefBlockHash:   chainStatus.BestChainHash,
		Params:         string(paramsByte),
	}
	createRaw, err := aelf.CreateRawTransaction(input)
	assert.NoError(t, err)
	//spew.Dump("Create Raw Transaction result", createRaw)
	rawTransactionBytes, err := hex.DecodeString(createRaw.RawTransaction)
	signature, _ := client.GetSignatureWithPrivateKey(aelf.PrivateKey, rawTransactionBytes)
	var sendRawinput = &dto.SendRawTransactionInput{
		Transaction:       createRaw.RawTransaction,
		Signature:         signature,
		ReturnTransaction: true,
	}
	executeRawResult, err := aelf.SendRawTransaction(sendRawinput)
	assert.NoError(t, err)
	//spew.Dump("Send Raw Transaction result", executeRawResult)
	assert.NotEmpty(t, executeRawResult.TransactionId)
	assert.Equal(t, _address, executeRawResult.Transaction.From)
	assert.Equal(t, tokenContractAddress, executeRawResult.Transaction.To)
	assert.Equal(t, chainStatus.BestChainHeight, executeRawResult.Transaction.RefBlockNumber)
	prefixBytes, _ := hex.DecodeString(chainStatus.BestChainHash)
	assert.Equal(t, base64.StdEncoding.EncodeToString(prefixBytes[:4]), executeRawResult.Transaction.RefBlockPrefix)
	assert.Equal(t, "Transfer", executeRawResult.Transaction.MethodName)
	assert.Equal(t, "{ \"to\": \""+userKeyPairInfo.Address+"\", \"symbol\": \"ELF\", \"amount\": \"1000000000\", \"memo\": \"transfer in test\" }", executeRawResult.Transaction.Params)
	signatureBytes, _ := hex.DecodeString(signature)
	assert.Equal(t, base64.StdEncoding.EncodeToString(signatureBytes), executeRawResult.Transaction.Signature)

	time.Sleep(time.Duration(4) * time.Second)

	balance := getBanlance(toAddress)
	assert.Equal(t, "ELF", balance.Symbol)
	assert.Equal(t, toAddress.Value, balance.Owner.Value)
	assert.Equal(t, int64(1000000000), balance.Balance)
}

func TestSendRawTransactionWithoutReturnTransaction(t *testing.T) {
	chainStatus, err := aelf.GetChainStatus()
	tokenContractAddress, _ := aelf.GetContractAddressByName("AElf.ContractNames.Token")
	userKeyPairInfo := aelf.GenerateKeyPairInfo()
	toAddress, _ := util.Base58StringToAddress(userKeyPairInfo.Address)
	params := &pb.TransferInput{
		To:     toAddress,
		Symbol: "ELF",
		Amount: 1000000000,
		Memo:   "transfer in test",
	}

	paramsByte, _ := protojson.Marshal(params)
	var input = &dto.CreateRawTransactionInput{
		From:           _address,
		To:             tokenContractAddress,
		MethodName:     "Transfer",
		RefBlockNumber: chainStatus.BestChainHeight,
		RefBlockHash:   chainStatus.BestChainHash,
		Params:         string(paramsByte),
	}
	createRaw, err := aelf.CreateRawTransaction(input)
	assert.NoError(t, err)
	//spew.Dump("Create Raw Transaction result", createRaw)
	rawTransactionBytes, err := hex.DecodeString(createRaw.RawTransaction)
	signature, _ := client.GetSignatureWithPrivateKey(aelf.PrivateKey, rawTransactionBytes)
	var sendRawinput = &dto.SendRawTransactionInput{
		Transaction:       createRaw.RawTransaction,
		Signature:         signature,
		ReturnTransaction: false,
	}
	executeRawResult, err := aelf.SendRawTransaction(sendRawinput)
	assert.NoError(t, err)
	//spew.Dump("Send Raw Transaction result", executeRawResult)
	assert.NotEmpty(t, executeRawResult.TransactionId)
	assert.Empty(t, executeRawResult.Transaction.From)
	assert.Empty(t, executeRawResult.Transaction.To)
	assert.Equal(t, int64(0), executeRawResult.Transaction.RefBlockNumber)
	assert.Empty(t, executeRawResult.Transaction.RefBlockPrefix)
	assert.Empty(t, executeRawResult.Transaction.MethodName)
	assert.Empty(t, executeRawResult.Transaction.Params)
	assert.Empty(t, executeRawResult.Transaction.Signature)

	time.Sleep(time.Duration(4) * time.Second)

	balance := getBanlance(toAddress)
	assert.Equal(t, "ELF", balance.Symbol)
	assert.Equal(t, toAddress.Value, balance.Owner.Value)
	assert.Equal(t, int64(1000000000), balance.Balance)
}

func TestExecuteRawTransaction(t *testing.T) {
	chainStatus, err := aelf.GetChainStatus()
	tokenContractAddress, _ := aelf.GetContractAddressByName("AElf.ContractNames.Token")
	userKeyPairInfo := aelf.GenerateKeyPairInfo()
	toAddress, _ := util.Base58StringToAddress(userKeyPairInfo.Address)
	transaction := createTransferTransaction(toAddress)
	transactionByets, _ := proto.Marshal(transaction)
	sendResult, err := aelf.SendTransaction(hex.EncodeToString(transactionByets))
	assert.NoError(t, err)
	assert.NotEmpty(t, sendResult.TransactionID)

	time.Sleep(time.Duration(4) * time.Second)

	getBalanceInput := &pb.GetBalanceInput{
		Symbol: "ELF",
		Owner:  toAddress,
	}
	paramsByte, _ := protojson.Marshal(getBalanceInput)
	//spew.Dump(paramsByte)
	var input = &dto.CreateRawTransactionInput{
		From:           _address,
		To:             tokenContractAddress,
		MethodName:     "GetBalance",
		RefBlockNumber: chainStatus.BestChainHeight,
		RefBlockHash:   chainStatus.BestChainHash,
		Params:         string(paramsByte),
	}
	createRaw, err := aelf.CreateRawTransaction(input)
	assert.NoError(t, err)
	//spew.Dump("Create Raw Transaction result", createRaw)
	rawTransactionBytes, err := hex.DecodeString(createRaw.RawTransaction)
	signature, _ := client.GetSignatureWithPrivateKey(aelf.PrivateKey, rawTransactionBytes)
	var executeRawinput = &dto.ExecuteRawTransactionDto{
		RawTransaction: createRaw.RawTransaction,
		Signature:      signature,
	}
	executeRawresult, err := aelf.ExecuteRawTransaction(executeRawinput)
	assert.NoError(t, err)

	assert.Equal(t, "{ \"symbol\": \"ELF\", \"owner\": \""+userKeyPairInfo.Address+"\", \"balance\": \"1000000000\" }", executeRawresult)
}

func TestSendTransaction(t *testing.T) {
	// Get token contract address.
	tokenContractAddress, _ := aelf.GetContractAddressByName("AElf.ContractNames.Token")
	fromAddress := aelf.GetAddressFromPrivateKey(aelf.PrivateKey)
	userKeyPairInfo := aelf.GenerateKeyPairInfo()
	toAddress, _ := util.Base58StringToAddress(userKeyPairInfo.Address)
	transaction := createTransferTransaction(toAddress)

	// Send the transfer transaction to AElf chain node.
	transactionByets, _ := proto.Marshal(transaction)
	sendResult, err := aelf.SendTransaction(hex.EncodeToString(transactionByets))
	assert.NoError(t, err)
	assert.NotEmpty(t, sendResult.TransactionID)

	time.Sleep(time.Duration(4) * time.Second)

	transactionResult, err := aelf.GetTransactionResult(sendResult.TransactionID)
	//spew.Dump("Create Raw Transaction result", transactionResult)
	assert.NoError(t, err)
	assert.Equal(t, sendResult.TransactionID, transactionResult.TransactionId)
	assert.Equal(t, "MINED", transactionResult.Status)
	assert.Empty(t, transactionResult.Error)

	assert.Equal(t, 2, len(transactionResult.Logs))

	assert.Equal(t, tokenContractAddress, transactionResult.Logs[0].Address)
	assert.Equal(t, "TransactionFeeCharged", transactionResult.Logs[0].Name)
	var feeCharged = new(pb.TransactionFeeCharged)
	nonIndexedBytes, _ := util.Base64DecodeBytes(transactionResult.Logs[0].NonIndexed)
	proto.Unmarshal(nonIndexedBytes, feeCharged)
	assert.Equal(t, "ELF", feeCharged.Symbol)
	assert.True(t, feeCharged.Amount > 0)

	assert.Equal(t, tokenContractAddress, transactionResult.Logs[1].Address)
	assert.Equal(t, "Transferred", transactionResult.Logs[1].Name)
	var transferred = new(pb.Transferred)
	indexedBytes, _ := util.Base64DecodeBytes(transactionResult.Logs[1].Indexed[0])
	proto.Unmarshal(indexedBytes, transferred)
	assert.Equal(t, fromAddress, utils.AddressToBase58String(transferred.From))

	transferred = new(pb.Transferred)
	indexedBytes, _ = util.Base64DecodeBytes(transactionResult.Logs[1].Indexed[1])
	proto.Unmarshal(indexedBytes, transferred)
	assert.Equal(t, userKeyPairInfo.Address, utils.AddressToBase58String(transferred.To))

	transferred = new(pb.Transferred)
	indexedBytes, _ = util.Base64DecodeBytes(transactionResult.Logs[1].Indexed[2])
	proto.Unmarshal(indexedBytes, transferred)
	assert.Equal(t, "ELF", transferred.Symbol)

	transferred = new(pb.Transferred)
	nonIndexedBytes, _ = util.Base64DecodeBytes(transactionResult.Logs[1].NonIndexed)
	proto.Unmarshal(nonIndexedBytes, transferred)
	assert.Equal(t, int64(1000000000), transferred.Amount)
	assert.Equal(t, "transfer in test", transferred.Memo)
}

func TestSendFailedTransaction(t *testing.T) {
	tokenContractAddress, _ := aelf.GetContractAddressByName("AElf.ContractNames.Token")
	toAddress, _ := util.Base58StringToAddress(aelf.GetAddressFromPrivateKey(aelf.PrivateKey))
	userKeyPairInfo := aelf.GenerateKeyPairInfo()
	methodName := "Transfer"

	params := &pb.TransferInput{
		To:     toAddress,
		Symbol: "ELF",
		Amount: 1000000000,
		Memo:   "transfer in test",
	}
	paramsByte, _ := proto.Marshal(params)

	transaction, err := aelf.CreateTransaction(userKeyPairInfo.Address, tokenContractAddress, methodName, paramsByte)

	signature, _ := aelf.SignTransaction(userKeyPairInfo.PrivateKey, transaction)
	transaction.Signature = signature

	transactionByets, _ := proto.Marshal(transaction)
	sendResult, err := aelf.SendTransaction(hex.EncodeToString(transactionByets))

	assert.NoError(t, err)
	assert.NotEmpty(t, sendResult.TransactionID)

	time.Sleep(time.Duration(4) * time.Second)

	transactionResult, err := aelf.GetTransactionResult(sendResult.TransactionID)
	assert.NoError(t, err)
	assert.Equal(t, "NODEVALIDATIONFAILED", transactionResult.Status)
	assert.Equal(t, "Pre-Error: Transaction fee not enough.", transactionResult.Error)
}

func TestExecuteTransaction(t *testing.T) {
	userKeyPairInfo := aelf.GenerateKeyPairInfo()
	toAddress, _ := util.Base58StringToAddress(userKeyPairInfo.Address)
	transaction := createTransferTransaction(toAddress)

	transactionByets, _ := proto.Marshal(transaction)
	sendResult, err := aelf.SendTransaction(hex.EncodeToString(transactionByets))
	assert.NoError(t, err)
	assert.NotEmpty(t, sendResult.TransactionID)

	time.Sleep(time.Duration(4) * time.Second)

	balance := getBanlance(toAddress)

	assert.Equal(t, "ELF", balance.Symbol)
	assert.Equal(t, toAddress.Value, balance.Owner.Value)
	assert.Equal(t, int64(1000000000), balance.Balance)
}

func TestSendTransctions(t *testing.T) {
	var transactions []string
	for i := 0; i < 2; i++ {
		user := aelf.GenerateKeyPairInfo()
		userAddress, _ := util.Base58StringToAddress(user.Address)
		transaction := createTransferTransaction(userAddress)
		transactionByets, _ := proto.Marshal(transaction)
		transactions = append(transactions, hex.EncodeToString(transactionByets))
	}
	txs := strings.Join(transactions, ",")
	results, err := aelf.SendTransactions(txs)
	assert.NoError(t, err)
	assert.True(t, len(results) == 2)
	assert.NotEmpty(t, results[0])
	assert.NotEmpty(t, results[1])

	time.Sleep(time.Duration(4) * time.Second)

	for i := 0; i < 2; i++ {
		transactionResult, _ := aelf.GetTransactionResult(results[i].(string))
		assert.Equal(t, "MINED", transactionResult.Status)
	}
}

func createTransferTransaction(toAddress *pb.Address) *pb.Transaction {
	tokenContractAddress, _ := aelf.GetContractAddressByName("AElf.ContractNames.Token")
	fromAddress := aelf.GetAddressFromPrivateKey(aelf.PrivateKey)
	methodName := "Transfer"

	params := &pb.TransferInput{
		To:     toAddress,
		Symbol: "ELF",
		Amount: 1000000000,
		Memo:   "transfer in test",
	}
	paramsByte, _ := proto.Marshal(params)

	transaction, _ := aelf.CreateTransaction(fromAddress, tokenContractAddress, methodName, paramsByte)
	signature, _ := aelf.SignTransaction(aelf.PrivateKey, transaction)
	transaction.Signature = signature

	return transaction
}

func getBanlance(owner *pb.Address) *pb.GetBalanceOutput {
	tokenContractAddress, _ := aelf.GetContractAddressByName("AElf.ContractNames.Token")
	fromAddress := aelf.GetAddressFromPrivateKey(aelf.PrivateKey)
	getBalanceInput := &pb.GetBalanceInput{
		Symbol: "ELF",
		Owner:  owner,
	}
	getBalanceInputByte, _ := proto.Marshal(getBalanceInput)

	getBalanceTransaction, _ := aelf.CreateTransaction(fromAddress, tokenContractAddress, "GetBalance", getBalanceInputByte)
	getBalanceSignature, _ := aelf.SignTransaction(aelf.PrivateKey, getBalanceTransaction)
	getBalanceTransaction.Signature = getBalanceSignature

	getBalanceTransactionByets, _ := proto.Marshal(getBalanceTransaction)
	getBalanceResult, _ := aelf.ExecuteTransaction(hex.EncodeToString(getBalanceTransactionByets))
	balance := &pb.GetBalanceOutput{}
	getBalanceResultBytes, _ := hex.DecodeString(getBalanceResult)
	proto.Unmarshal(getBalanceResultBytes, balance)

	return balance
}

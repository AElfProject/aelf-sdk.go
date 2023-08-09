package test

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"github.com/AElfProject/aelf-sdk.go/model"
	"github.com/AElfProject/aelf-sdk.go/model/consts"
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

var (
	ContractAddress, _               = aelf.GetGenesisContractAddress()
	defaultTestHolder                = getDefaultTestHolder(true)
	defaultSideChainTestHolder       = getDefaultTestHolder(false)
	defaultTestCrossChainFromChainId = utils.ConvertBase58ToChainId(DefaultMainChain)
	defaultTestCrossChainToChainId   = utils.ConvertBase58ToChainId(DefaultTestSideChain)
	testSideChainClient              = client.AElfClient{
		Host:       "http://192.168.66.106:8000",
		Version:    "1.0",
		PrivateKey: "36bc3f264aa340d44aada5759a5a86aac6d734f19932397e551d9e69edffe0d2",
	}
)

const (
	ContractMethodName = "GetContractAddressByName"

	DefaultTestSymbol            = "ELF"
	DefaultTransferTestAmount    = 1000000000
	DefaultTransferTestMemo      = "transfer in test"
	DefaultTransferTestWaitTime  = 8 * time.Second
	DefaultIndexingTestWaitTime  = 2 * time.Minute
	DefaultMainChain             = "AELF"
	DefaultTestSideChain         = "tDVW"
	DefaultTestTokenTotalSupply  = int64(100000000000000000)
	DefaultTestTokenDecimals     = int32(8)
	DefaultTestTokenIsBurnable   = true
	DefaultTestTokenIssueChainId = int32(9992731)
)

type TestHolder struct {
	KeyPair *model.KeyPairInfo
	Address *pb.Address
}

func getDefaultTestHolder(isMainChain bool) *TestHolder {
	userKeyPairInfo := aelf.GenerateKeyPairInfo()
	if isMainChain == false {
		userKeyPairInfo = testSideChainClient.GenerateKeyPairInfo()
	}
	toAddress, _ := util.Base58StringToAddress(userKeyPairInfo.Address)
	return &TestHolder{
		KeyPair: userKeyPairInfo,
		Address: toAddress,
	}
}

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
	block, err := aelf.GetBlockByHeight(1, isTransactions)
	assert.NoError(t, err)
	transactionID := block.Body.Transactions[0]
	merklePath, err := aelf.GetMerklePathByTransactionID(transactionID)
	assert.NoError(t, err)
	assert.True(t, len(merklePath.MerklePathNodes) == 4)
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
		Value: utils.GetBytesSha256(consts.TokenContractSystemName),
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

func TestSendRawTransaction(t *testing.T) {
	chainStatus, err := aelf.GetChainStatus()
	tokenContractAddress, _ := aelf.GetContractAddressByName(consts.TokenContractSystemName)
	userKeyPairInfo := aelf.GenerateKeyPairInfo()
	toAddress, _ := util.Base58StringToAddress(userKeyPairInfo.Address)
	params := &pb.TransferInput{
		To:     toAddress,
		Symbol: DefaultTestSymbol,
		Amount: DefaultTransferTestAmount,
		Memo:   DefaultTransferTestMemo,
	}

	paramsByte, _ := protojson.Marshal(params)
	var input = &dto.CreateRawTransactionInput{
		From:           _address,
		To:             tokenContractAddress,
		MethodName:     consts.TokenContractTransfer,
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
	assert.Equal(t, consts.TokenContractTransfer, executeRawResult.Transaction.MethodName)
	assert.Equal(t, "{ \"to\": \""+userKeyPairInfo.Address+"\", \"symbol\": \"ELF\", \"amount\": \"1000000000\", \"memo\": \"transfer in test\" }", executeRawResult.Transaction.Params)
	signatureBytes, _ := hex.DecodeString(signature)
	assert.Equal(t, base64.StdEncoding.EncodeToString(signatureBytes), executeRawResult.Transaction.Signature)

	time.Sleep(DefaultTransferTestWaitTime)

	balance, _ := aelf.GetTokenBalance(DefaultTestSymbol, userKeyPairInfo.Address)
	assert.Equal(t, "ELF", balance.Symbol)
	assert.Equal(t, toAddress.Value, balance.Owner.Value)
	assert.Equal(t, int64(DefaultTransferTestAmount), balance.Balance)
}

func TestSendRawTransactionWithoutReturnTransaction(t *testing.T) {
	chainStatus, err := aelf.GetChainStatus()
	tokenContractAddress, _ := aelf.GetContractAddressByName(consts.TokenContractSystemName)
	userKeyPairInfo := aelf.GenerateKeyPairInfo()
	toAddress, _ := util.Base58StringToAddress(userKeyPairInfo.Address)
	params := &pb.TransferInput{
		To:     toAddress,
		Symbol: DefaultTestSymbol,
		Amount: DefaultTransferTestAmount,
		Memo:   DefaultTransferTestMemo,
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

	time.Sleep(DefaultTransferTestWaitTime)

	balance, _ := aelf.GetTokenBalance(DefaultTestSymbol, userKeyPairInfo.Address)
	assert.Equal(t, DefaultTestSymbol, balance.Symbol)
	assert.Equal(t, toAddress.Value, balance.Owner.Value)
	assert.Equal(t, int64(1000000000), balance.Balance)
}

func TestExecuteRawTransaction(t *testing.T) {
	chainStatus, err := aelf.GetChainStatus()
	tokenContractAddress, _ := aelf.GetContractAddressByName(consts.TokenContractSystemName)
	userKeyPairInfo := aelf.GenerateKeyPairInfo()
	toAddress, _ := util.Base58StringToAddress(userKeyPairInfo.Address)
	transaction := createTransferTransaction(toAddress)
	transactionByets, _ := proto.Marshal(transaction)
	sendResult, err := aelf.SendTransaction(hex.EncodeToString(transactionByets))
	assert.NoError(t, err)
	assert.NotEmpty(t, sendResult.TransactionID)

	time.Sleep(DefaultTransferTestWaitTime)

	getBalanceInput := &pb.GetBalanceInput{
		Symbol: DefaultTestSymbol,
		Owner:  toAddress,
	}
	paramsByte, _ := protojson.Marshal(getBalanceInput)
	//spew.Dump(paramsByte)
	var input = &dto.CreateRawTransactionInput{
		From:           _address,
		To:             tokenContractAddress,
		MethodName:     consts.TokenContractGetBalance,
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
	tokenContractAddress, _ := aelf.GetContractAddressByName(consts.TokenContractSystemName)
	fromAddress := aelf.GetAddressFromPrivateKey(aelf.PrivateKey)
	userKeyPairInfo := aelf.GenerateKeyPairInfo()
	toAddress, _ := util.Base58StringToAddress(userKeyPairInfo.Address)
	transaction := createTransferTransaction(toAddress)

	// Send the transfer transaction to AElf chain node.
	transactionByets, _ := proto.Marshal(transaction)
	sendResult, err := aelf.SendTransaction(hex.EncodeToString(transactionByets))
	assert.NoError(t, err)
	assert.NotEmpty(t, sendResult.TransactionID)

	time.Sleep(DefaultTransferTestWaitTime)

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
	assert.Equal(t, DefaultTestSymbol, feeCharged.Symbol)
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
	assert.Equal(t, DefaultTestSymbol, transferred.Symbol)

	transferred = new(pb.Transferred)
	nonIndexedBytes, _ = util.Base64DecodeBytes(transactionResult.Logs[1].NonIndexed)
	proto.Unmarshal(nonIndexedBytes, transferred)
	assert.Equal(t, int64(1000000000), transferred.Amount)
	assert.Equal(t, "transfer in test", transferred.Memo)
}

func TestSendFailedTransaction(t *testing.T) {
	tokenContractAddress, _ := aelf.GetContractAddressByName(consts.TokenContractSystemName)
	toAddress, _ := util.Base58StringToAddress(aelf.GetAddressFromPrivateKey(aelf.PrivateKey))
	userKeyPairInfo := aelf.GenerateKeyPairInfo()
	methodName := consts.TokenContractTransfer

	params := &pb.TransferInput{
		To:     toAddress,
		Symbol: DefaultTestSymbol,
		Amount: DefaultTransferTestAmount,
		Memo:   DefaultTransferTestMemo,
	}
	paramsByte, _ := proto.Marshal(params)

	transaction, err := aelf.CreateTransaction(userKeyPairInfo.Address, tokenContractAddress, methodName, paramsByte)

	signature, _ := aelf.SignTransaction(userKeyPairInfo.PrivateKey, transaction)
	transaction.Signature = signature

	transactionByets, _ := proto.Marshal(transaction)
	sendResult, err := aelf.SendTransaction(hex.EncodeToString(transactionByets))

	assert.NoError(t, err)
	assert.NotEmpty(t, sendResult.TransactionID)

	time.Sleep(DefaultTransferTestWaitTime)

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

	time.Sleep(DefaultTransferTestWaitTime)

	balance, _ := aelf.GetTokenBalance(DefaultTestSymbol, userKeyPairInfo.Address)
	assert.Equal(t, DefaultTestSymbol, balance.Symbol)
	assert.Equal(t, toAddress.Value, balance.Owner.Value)
	assert.Equal(t, int64(DefaultTransferTestAmount), balance.Balance)
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

	time.Sleep(DefaultTransferTestWaitTime)

	for i := 0; i < 2; i++ {
		transactionResult, _ := aelf.GetTransactionResult(results[i].(string))
		assert.Equal(t, "MINED", transactionResult.Status)
	}
}

func TestCalculateTransactionFee(t *testing.T) {
	chainStatus, err := aelf.GetChainStatus()
	assert.NoError(t, err)
	params := &pb.Hash{
		Value: utils.GetBytesSha256(consts.TokenContractSystemName),
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
	var rawTransaction = result.RawTransaction
	var transactionFeeInput = &dto.CalculateTransactionFeeInput{
		RawTransaction: rawTransaction,
	}
	feeResult, err := aelf.CalculateTransactionFee(transactionFeeInput)
	assert.NoError(t, err)
	jsonStr, err := json.Marshal(feeResult.TransactionFee)
	assert.True(t, feeResult.Success)
	assert.NotEmpty(t, feeResult.TransactionFee[DefaultTestSymbol])
	assert.Greater(t, feeResult.TransactionFee[DefaultTestSymbol], float64(1.7e+05))
	assert.Less(t, feeResult.TransactionFee[DefaultTestSymbol], float64(1.9e+07))
	spew.Dump("CalculateTransactionFeeResult : ", jsonStr)

}

func TestGetTransferred(t *testing.T) {
	transaction := createTransferTransaction(defaultTestHolder.Address)

	transactionByets, _ := proto.Marshal(transaction)
	sendResult, err := aelf.SendTransaction(hex.EncodeToString(transactionByets))
	assert.NoError(t, err)
	assert.NotEmpty(t, sendResult.TransactionID)

	time.Sleep(DefaultTransferTestWaitTime)

	transferreds := aelf.GetTransferred(sendResult.TransactionID)
	assert.Len(t, transferreds, 1)
	assert.Equal(t, DefaultTestSymbol, transferreds[0].GetSymbol())
	assert.Equal(t, int64(DefaultTransferTestAmount), transferreds[0].GetAmount())
	assert.Equal(t, DefaultTransferTestMemo, transferreds[0].GetMemo())
}

//func TestGetCrossChainTransferred(t *testing.T) {
//	result, _ := createCrossChainTransferTx(defaultSideChainTestHolder.Address)
//
//	time.Sleep(DefaultTransferTestWaitTime)
//	crossChainTransferred := aelf.GetCrossChainTransferred(result.TransactionID)
//	assert.Len(t, crossChainTransferred, 1)
//	assert.Equal(t, DefaultTestSymbol, crossChainTransferred[0].GetSymbol())
//	assert.Equal(t, int64(DefaultTransferTestAmount), crossChainTransferred[0].GetAmount())
//	assert.Equal(t, DefaultTransferTestMemo, crossChainTransferred[0].GetMemo())
//}
//
//func TestGetCrossChainReceived(t *testing.T) {
//	crossChainTxOutput, crossChainTxBytes := createCrossChainTransferTx(defaultSideChainTestHolder.Address)
//
//	time.Sleep(DefaultTransferTestWaitTime)
//	txResult, _ := aelf.GetTransactionResult(crossChainTxOutput.TransactionID)
//
//	time.Sleep(DefaultIndexingTestWaitTime)
//	mp, _ := aelf.GetMerklePathByTransactionID(crossChainTxOutput.TransactionID)
//
//	params := &pb.CrossChainReceiveTokenInput{
//		FromChainId:              defaultTestCrossChainFromChainId,
//		ParentChainHeight:        txResult.BlockNumber,
//		TransferTransactionBytes: crossChainTxBytes,
//		MerklePath:               getTxMerklePath(mp),
//	}
//
//	paramsByte, _ := proto.Marshal(params)
//
//	tokenContractAddress, _ := testSideChainClient.GetContractAddressByName(consts.TokenContractSystemName)
//	transaction, _ := testSideChainClient.CreateTransaction(testSideChainClient.GetAddressFromPrivateKey(testSideChainClient.PrivateKey),
//		tokenContractAddress, consts.CrossChainContractCrossChainReceiveToken, paramsByte)
//	signature, _ := testSideChainClient.SignTransaction(testSideChainClient.PrivateKey, transaction)
//	transaction.Signature = signature
//
//	txBytes, _ := proto.Marshal(transaction)
//	result, err := testSideChainClient.SendTransaction(hex.EncodeToString(txBytes))
//	assert.NoError(t, err)
//	assert.NotEmpty(t, result.TransactionID)
//
//	time.Sleep(DefaultTransferTestWaitTime)
//	crossChainReceiveds := testSideChainClient.GetCrossChainReceived(result.TransactionID)
//
//	assert.Len(t, crossChainReceiveds, 1)
//	assert.Equal(t, defaultSideChainTestHolder.KeyPair.Address, utils.AddressToBase58String(crossChainReceiveds[0].GetTo()))
//	assert.Equal(t, aelf.GetAddressFromPrivateKey(aelf.PrivateKey), utils.AddressToBase58String(crossChainReceiveds[0].GetFrom()))
//	assert.Equal(t, DefaultTestSymbol, crossChainReceiveds[0].GetSymbol())
//	assert.Equal(t, int64(DefaultTransferTestAmount), crossChainReceiveds[0].GetAmount())
//	assert.Equal(t, DefaultTransferTestMemo, crossChainReceiveds[0].GetMemo())
//	assert.Equal(t, defaultTestCrossChainFromChainId, crossChainReceiveds[0].GetFromChainId())
//	assert.Equal(t, txResult.BlockNumber, crossChainReceiveds[0].GetParentChainHeight())
//	assert.Equal(t, defaultTestCrossChainFromChainId, crossChainReceiveds[0].GetIssueChainId())
//}

func createTransferTransaction(toAddress *pb.Address) *pb.Transaction {
	tokenContractAddress, _ := aelf.GetContractAddressByName(consts.TokenContractSystemName)
	methodName := consts.TokenContractTransfer

	params := &pb.TransferInput{
		To:     toAddress,
		Symbol: DefaultTestSymbol,
		Amount: DefaultTransferTestAmount,
		Memo:   DefaultTransferTestMemo,
	}

	paramsByte, _ := proto.Marshal(params)

	transaction, _ := aelf.CreateTransaction(aelf.GetAddressFromPrivateKey(aelf.PrivateKey), tokenContractAddress, methodName, paramsByte)
	signature, _ := aelf.SignTransaction(aelf.PrivateKey, transaction)
	transaction.Signature = signature

	return transaction
}

func createCrossChainTransferTx(toAddress *pb.Address) (*dto.SendTransactionOutput, []byte) {
	crossChainContractAddress, _ := aelf.GetContractAddressByName(consts.TokenContractSystemName)
	fromAddress := aelf.GetAddressFromPrivateKey(aelf.PrivateKey)
	methodName := consts.TokenContractCrossChainTransfer

	params := &pb.CrossChainTransferInput{
		To:           toAddress,
		Symbol:       DefaultTestSymbol,
		Amount:       DefaultTransferTestAmount,
		Memo:         DefaultTransferTestMemo,
		ToChainId:    defaultTestCrossChainToChainId,
		IssueChainId: defaultTestCrossChainFromChainId,
	}

	paramsByte, _ := proto.Marshal(params)

	result, txByets := sendTx(fromAddress, crossChainContractAddress, methodName, paramsByte)
	return result, txByets
}

func getTxMerklePath(merklePath *dto.MerklePathDto) *pb.MerklePath {
	mpn := make([]*pb.MerklePathNode, len(merklePath.MerklePathNodes))
	for i, node := range merklePath.MerklePathNodes {
		hashByte, _ := hex.DecodeString(node.Hash)
		mpn[i] = &pb.MerklePathNode{
			Hash:            &pb.Hash{Value: hashByte},
			IsLeftChildNode: node.IsLeftChildNode,
		}
	}
	return &pb.MerklePath{MerklePathNodes: mpn}
}

func exTx(fromAddr, contractAddr, methodName string, inputBytes []byte) string {
	tx, _ := aelf.CreateTransaction(fromAddr, contractAddr, methodName, inputBytes)
	sign, _ := aelf.SignTransaction(aelf.PrivateKey, tx)
	tx.Signature = sign

	byets, _ := proto.Marshal(tx)
	exResult, _ := aelf.ExecuteTransaction(hex.EncodeToString(byets))
	return exResult
}

func sendTx(fromAddr, contractAddr, methodName string, inputBytes []byte) (*dto.SendTransactionOutput, []byte) {
	tx, _ := aelf.CreateTransaction(fromAddr, contractAddr, methodName, inputBytes)
	sign, _ := aelf.SignTransaction(aelf.PrivateKey, tx)
	tx.Signature = sign

	byets, _ := proto.Marshal(tx)
	exResult, _ := aelf.SendTransaction(hex.EncodeToString(byets))
	return exResult, byets
}

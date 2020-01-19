package test

import (
	"testing"

	"aelf_sdk.go/aelf_sdk/client"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

var aelf = client.AElfClient{
	Host:       "http://127.0.0.1:8000",
	Version:    "",
	PrivateKey: "680afd630d82ae5c97942c4141d60b8a9fedfa5b2864fca84072c17ee1f72d9d",
	PublicKey:  "",
}

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
	height, err := aelf.GetBlockHeight()
	if err != nil || height == 0 {
		t.Error("GetBlockHeight err", err)
	}
	spew.Dump("# get_block_height", height)

	var isTransactions = true
	var blockHeight = 1
	byHeightBlock, err := aelf.GetBlockByHeight(blockHeight, isTransactions)
	if err != nil {
		t.Error("GetBlockByHeight err", err)
	}
	spew.Dump("get_block_by_height>>>>>>", byHeightBlock)

	blockHash := byHeightBlock.BlockHash
	byHashBlock, err := aelf.GetBlockByHash(blockHash, isTransactions)
	if err != nil {
		t.Error("Get Block By Height err", err)
	}
	spew.Dump("# get_block_by_hash", byHashBlock)

}

func TestTransactionResultApi(t *testing.T) {
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

	blockHash := block.BlockHash
	transactionResults, err := aelf.GetTransactionResults(blockHash)
	if err != nil {
		t.Error("Get Transaction Results error", err)
	}
	spew.Dump("# Get Transaction Results", transactionResults)

	merklePath, err := aelf.GetMerklePathByTransactionID(transactionID)
	if err != nil {
		t.Error("Get merkle Path  error", err)
	}
	spew.Dump("Get Merkle Path By TransactionID", merklePath)
}

func TestNetworkApi(t *testing.T) {
	netWorkInfo, err := aelf.GetNetworkInfo()
	if err != nil {
		t.Error("get net work info error", err)
	}
	spew.Dump("# get_network_info", netWorkInfo)

	removePeer, err := aelf.RemovePeer("127.0.0.1:6801")
	if err != nil {
		t.Error("remove peer error", err)
	}
	assert.True(t, true, removePeer)

	addPeer, err := aelf.AddPeer("127.0.0.1:6801")
	if err != nil {
		t.Error("add peer error", err)
	}
	assert.True(t, true, addPeer)

	peers, err := aelf.GetPeers()
	if err != nil {
		t.Error("get peers error", err)
	}
	spew.Dump("# get peers", peers)
}

func TestTransactionPoolApi(t *testing.T) {
	transactionPoolStatus, err := aelf.GetTransactionPoolStatus()
	if err != nil {
		t.Error("Get Transaction Pool Status error", err)
	}
	assert.True(t, true, transactionPoolStatus.Queued >= 0)

}

func TestTaskQueueApi(t *testing.T) {
	taskQueueStatus, err := aelf.GetTaskQueueStatus()
	if err != nil || len(taskQueueStatus) == 0 {
		t.Error("get task Queue Status errro", err)
	}
	spew.Dump("# get task Queue  status", taskQueueStatus)
}

//以上已完成测试

func TestRawTransactionApi(t *testing.T) {
	// reqParams := map[string]interface{}{
	// 	"From":           from,
	// 	"To":             to,
	// 	"RefBlockNumber": strconv.Itoa(blockNumber),
	// 	"RefBlockHash":   blockHash,
	// 	"MethodName":     method,
	// 	"Params":         params,
	// }
}

// // func SendRawTransaction() {
// // 	params := map[string]interface{}{
// // 		"Transaction":       transaction,
// // 		"Signature":         signature,
// // 		"ReturnTransaction": returnTransaction,
// // 	}
// // }

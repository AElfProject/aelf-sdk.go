package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetBlockHeight(t *testing.T) {
	height, err := aelf.GetBlockHeight()
	assert.NoError(t, err)
	assert.True(t, height > 0)
}

func TestGetBlock(t *testing.T) {
	height, err := aelf.GetBlockHeight()

	blockByHeight, err := aelf.GetBlockByHeight(height, false)
	assert.NoError(t, err)

	blockByHash, err := aelf.GetBlockByHash(blockByHeight.BlockHash, false)
	assert.NoError(t, err)

	assert.Equal(t, blockByHeight, blockByHash)
	assert.NotEmpty(t, blockByHeight.BlockHash)

	assert.Equal(t, blockByHeight.Header.Height, height)
	assert.NotEmpty(t, blockByHeight.Header.PreviousBlockHash)
	assert.NotEmpty(t, blockByHeight.Header.MerkleTreeRootOfTransactions)
	assert.NotEmpty(t, blockByHeight.Header.MerkleTreeRootOfWorldState)
	assert.NotEmpty(t, blockByHeight.Header.Extra)
	assert.Equal(t, "AELF", blockByHeight.Header.ChainId)
	assert.NotEmpty(t, blockByHeight.Header.Bloom)
	assert.NotEmpty(t, blockByHeight.Header.SignerPubkey)
	assert.True(t, blockByHeight.Header.Time.After(time.Time{}))

	assert.True(t, blockByHeight.Body.TransactionsCount > 0)
	assert.Equal(t, 0, len(blockByHeight.Body.Transactions))

	previousBlock, err := aelf.GetBlockByHash(blockByHeight.Header.PreviousBlockHash, false)
	assert.NoError(t, err)
	assert.Equal(t, previousBlock.BlockHash, blockByHeight.Header.PreviousBlockHash)
	assert.Equal(t, previousBlock.Header.Height, blockByHeight.Header.Height-1)
}

func TestGetBlockWithTransaction(t *testing.T) {
	height, err := aelf.GetBlockHeight()

	blockByHeight, err := aelf.GetBlockByHeight(height, true)
	assert.NoError(t, err)

	blockByHash, err := aelf.GetBlockByHash(blockByHeight.BlockHash, true)
	assert.NoError(t, err)

	assert.Equal(t, blockByHeight, blockByHash)
	assert.NotEmpty(t, blockByHeight.BlockHash)

	assert.Equal(t, blockByHeight.Header.Height, height)
	assert.NotEmpty(t, blockByHeight.Header.PreviousBlockHash)
	assert.NotEmpty(t, blockByHeight.Header.MerkleTreeRootOfTransactions)
	assert.NotEmpty(t, blockByHeight.Header.MerkleTreeRootOfWorldState)
	assert.NotEmpty(t, blockByHeight.Header.Extra)
	assert.Equal(t, "AELF", blockByHeight.Header.ChainId)
	assert.NotEmpty(t, blockByHeight.Header.Bloom)
	assert.NotEmpty(t, blockByHeight.Header.SignerPubkey)

	assert.True(t, blockByHeight.Body.TransactionsCount > 0)
	assert.Equal(t, blockByHeight.Body.TransactionsCount, len(blockByHeight.Body.Transactions))
	for _, tx := range blockByHeight.Body.Transactions {
		assert.NotEmpty(t, tx)
	}

	previousBlock, err := aelf.GetBlockByHash(blockByHeight.Header.PreviousBlockHash, true)
	assert.NoError(t, err)
	assert.Equal(t, previousBlock.BlockHash, blockByHeight.Header.PreviousBlockHash)
	assert.Equal(t, previousBlock.Header.Height, blockByHeight.Header.Height-1)
}

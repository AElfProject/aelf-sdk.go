package test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestGetChainStatus(t *testing.T) {
	chainStatus, err := aelf.GetChainStatus()

	assert.NoError(t, err)
	assert.Equal(t, "AELF", chainStatus.ChainId)
	assert.NotEmpty(t, chainStatus.Branches)
	assert.True(t, chainStatus.LongestChainHeight > 0)
	assert.NotEmpty(t, chainStatus.LongestChainHash)
	assert.NotEmpty(t, chainStatus.GenesisContractAddress)
	assert.NotEmpty(t, chainStatus.GenesisBlockHash)
	assert.True(t, chainStatus.LastIrreversibleBlockHeight > 0)
	assert.NotEmpty(t, chainStatus.LastIrreversibleBlockHash)
	assert.True(t, chainStatus.BestChainHeight > 0)
	assert.NotEmpty(t, chainStatus.BestChainHash)

	longestChainBlock, err := aelf.GetBlockByHash(chainStatus.LongestChainHash, false)
	assert.Equal(t, longestChainBlock.Header.Height, chainStatus.LongestChainHeight)

	genesisBlock, err := aelf.GetBlockByHash(chainStatus.GenesisBlockHash, false)
	assert.Equal(t, int64(1), genesisBlock.Header.Height)

	lastIrreversibleBlock, err := aelf.GetBlockByHash(chainStatus.LastIrreversibleBlockHash, false)
	assert.Equal(t, lastIrreversibleBlock.Header.Height, chainStatus.LastIrreversibleBlockHeight)

	bestChainBlock, err := aelf.GetBlockByHash(chainStatus.BestChainHash, false)
	assert.Equal(t, bestChainBlock.Header.Height, chainStatus.BestChainHeight)

	genesisContractAddress, err := aelf.GetGenesisContractAddress()
	assert.Equal(t, genesisContractAddress, chainStatus.GenesisContractAddress)
}

func TestGetChainID(t *testing.T) {
	chainID, err := aelf.GetChainID()
	assert.NoError(t, err)
	assert.Equal(t, 9992731, chainID)
}

func TestGetTaskQueueStatus(t *testing.T) {
	taskQueueStatus, err := aelf.GetTaskQueueStatus()
	assert.NoError(t, err)
	spew.Dump("Get Task Queue Status Result", taskQueueStatus)
}

func TestGetContractFileDescriptorSet(t *testing.T) {
	contractAddr, err := aelf.GetGenesisContractAddress()
	assert.NoError(t, err)
	contractFile, err := aelf.GetContractFileDescriptorSet(contractAddr)
	assert.NoError(t, err)
	assert.NotEmpty(t, contractFile)
	spew.Dump("Get Contract File Descriptor Set Result", contractFile)
}

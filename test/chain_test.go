package test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

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
	spew.Dump("Get Contract File Descriptor Set Result", contractFile)
}

package test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestGetBlockHeight(t *testing.T) {
	height, err := aelf.GetBlockHeight()
	assert.NoError(t, err)
	assert.True(t, height > 0)
}

func TestGetBlockByHeight(t *testing.T) {
	var includeTransactions = true
	height, err := aelf.GetBlockHeight()
	spew.Dump("Get Block height Result", height)

	HeightBlock, err := aelf.GetBlockByHeight(int(height), includeTransactions)
	assert.NoError(t, err)
	spew.Dump("Get Block ByHeight Result", HeightBlock)

}

func TestGetBlockByHash(t *testing.T) {
	var includeTransactions = true
	height, _ := aelf.GetBlockHeight()
	HeightBlock, _ := aelf.GetBlockByHeight(int(height), includeTransactions)
	byHashBlock, err := aelf.GetBlockByHash(HeightBlock.BlockHash, includeTransactions)
	assert.NoError(t, err)
	spew.Dump("Get Block ByHash Result", byHashBlock)
}

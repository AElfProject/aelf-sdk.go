package test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

// var TestAddress = "127.0.0.1:6801"
const version = "1.3.0.0"

func TestNetworkApi(t *testing.T) {
	// Test GetNetworkInfo
	netWorkInfo, err := aelf.GetNetworkInfo()
	assert.NoError(t, err)
	assert.Equal(t, version, netWorkInfo.Version)
	spew.Dump("Get Network Info Result", netWorkInfo)

	// Test AddPeer
	// addPeer, err := aelf.AddPeer(TestAddress)
	// assert.NoError(t, err)
	// assert.True(t, addPeer == true)

	// Test RemovePeer
	// removePeer, err := aelf.RemovePeer(TestAddress)
	// assert.NoError(t, err)
	// assert.True(t, removePeer == true)

	//Test GetPeers
	peers, err := aelf.GetPeers(true)
	assert.NoError(t, err)
	spew.Dump("Get Peers Result", peers)
}

package client

import (
	"encoding/hex"
	"encoding/json"
	"errors"

	"github.com/AElfProject/aelf-sdk.go/dto"
	util "github.com/AElfProject/aelf-sdk.go/utils"
)

// GetBlockHeight Get height of the current chain.
func (client *AElfClient) GetBlockHeight() (int64, error) {
	url := client.Host + BLOCKHEIGHT
	heightBytes, err := util.GetRequest("GET", url, client.Version, nil)
	if err != nil {
		return 0, errors.New("Get BlockHeight error:" + err.Error())
	}
	var data interface{}
	json.Unmarshal(heightBytes, &data)
	return int64(data.(float64)), nil
}

// GetBlockByHash Get information of a block by given block hash. Optional whether to include transaction information.
func (client *AElfClient) GetBlockByHash(blockHash string, includeTransactions bool) (*dto.BlockDto, error) {
	_, err := hex.DecodeString(blockHash)
	if err != nil {
		return nil, errors.New("transactionID hex to []byte error:" + err.Error())
	}
	params := map[string]interface{}{
		"blockHash":           blockHash,
		"includeTransactions": includeTransactions,
	}
	url := client.Host + BLOCKBYHASH
	blockBytes, err := util.GetRequest("GET", url, client.Version, params)
	if err != nil {
		return nil, errors.New("Get Block ByHash error:" + err.Error())
	}
	var block = new(dto.BlockDto)
	json.Unmarshal(blockBytes, &block)
	return block, nil
}

// GetBlockByHeight Get information of a block by specified height. Optional whether to include transaction information.
func (client *AElfClient) GetBlockByHeight(blockHeight int64, includeTransactions bool) (*dto.BlockDto, error) {
	params := map[string]interface{}{
		"blockHeight":         blockHeight,
		"includeTransactions": includeTransactions,
	}
	url := client.Host + BLOCKBYHEIGHT
	blockBytes, err := util.GetRequest("GET", url, client.Version, params)
	if err != nil {
		return nil, errors.New("Get Block ByHeight error:" + err.Error())
	}
	var block = new(dto.BlockDto)
	json.Unmarshal(blockBytes, &block)
	return block, nil
}

package client

import (
	"encoding/json"

	"aelf_sdk.go/aelf_sdk/dto"
	util "aelf_sdk.go/aelf_sdk/utils"
)

//GetBlockHeight Get height of the current chain
func (a *AElfClient) GetBlockHeight() (float64, error) {
	url := a.Host + BLOCKHEIGHT
	heightBytes, err := util.GetRequest("GET", url, a.Version, nil)
	if err != nil {
		return 0, err
	}
	var data interface{}
	json.Unmarshal(heightBytes, &data)
	return data.(float64), nil
}

// GetBlockByHash Get information of a block by given block hash. Optional whether to include transaction information.
func (a *AElfClient) GetBlockByHash(blockHash string, isTransactions bool) (*dto.BlockDto, error) {
	params := map[string]interface{}{
		"blockHash":           blockHash,
		"includeTransactions": isTransactions,
	}
	url := a.Host + BLOCKBYHASH
	blockBytes, err := util.GetRequest("GET", url, a.Version, params)
	if err != nil {
		return nil, err
	}
	var block = new(dto.BlockDto)
	json.Unmarshal(blockBytes, &block)
	return block, nil
}

//GetBlockByHeight Get information of a block by specified height. Optional whether to include transaction information.
func (a *AElfClient) GetBlockByHeight(blockHeight int, isTransactions bool) (*dto.BlockDto, error) {
	params := map[string]interface{}{
		"blockHeight":         blockHeight,
		"includeTransactions": isTransactions,
	}
	url := a.Host + BLOCKBYHEIGHT
	blockBytes, err := util.GetRequest("GET", url, a.Version, params)
	if err != nil {
		return nil, err
	}
	var block = new(dto.BlockDto)
	json.Unmarshal(blockBytes, &block)
	return block, nil
}

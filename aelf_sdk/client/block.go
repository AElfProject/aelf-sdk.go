package client

import (
	"encoding/json"

	"aelf_sdk.go/aelf_sdk/dto"
)

//GetBlockHeight Get Block Height
func (a *AElfClient) GetBlockHeight() (float64, error) {
	url := a.Host + BLOCKHEIGHT
	heightBytes, err := GetRequest("GET", url, a.Version, nil)
	if err != nil {
		return 0, err
	}
	var data interface{}
	json.Unmarshal(heightBytes, &data)
	return data.(float64), err
}

// GetBlockByHash Get Block By Hash
func (a *AElfClient) GetBlockByHash(blockHash string, isTransactions bool) (*dto.BlockDto, error) {
	params := map[string]interface{}{
		"blockHash":           blockHash,
		"includeTransactions": isTransactions,
	}
	url := a.Host + BLOCKBYHASH
	blockBytes, err := GetRequest("GET", url, a.Version, params)
	if err != nil {
		return nil, err
	}
	var block = new(dto.BlockDto)
	json.Unmarshal(blockBytes, &block)
	return block, nil
}

//GetBlockByHeight Get Block By Height
func (a *AElfClient) GetBlockByHeight(blockHeight int, isTransactions bool) (*dto.BlockDto, error) {
	params := map[string]interface{}{
		"blockHeight":         blockHeight,
		"includeTransactions": isTransactions,
	}
	url := a.Host + BLOCKBYHEIGHT
	blockBytes, err := GetRequest("GET", url, a.Version, params)
	if err != nil {
		return nil, err
	}
	var block = new(dto.BlockDto)
	json.Unmarshal(blockBytes, &block)
	return block, nil
}

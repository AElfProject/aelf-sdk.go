package client

import (
	"encoding/binary"
	"encoding/json"

	"aelf_sdk.go/aelf_sdk/dto"
	"github.com/btcsuite/btcutil/base58"
)

//GetChainStatus Get Chain Status
func (a *AElfClient) GetChainStatus() (*dto.ChainStatusDto, error) {
	url := a.Host + CHAINSTATUS
	chainBytes, err := GetRequest("GET", url, a.Version, nil)
	if err != nil {
		return nil, err
	}
	var chain = new(dto.ChainStatusDto)
	json.Unmarshal(chainBytes, &chain)
	return chain, nil
}

//GetContractFileDescriptorSet Get Contract File Descriptor Set
func (a *AElfClient) GetContractFileDescriptorSet(address string) ([]byte, error) {
	url := a.Host + FILEDESCRIPTOR
	params := map[string]interface{}{"address": address}
	data, err := GetRequest("GET", url, a.Version, params)
	if err != nil {
		return nil, err
	}
	return data, err
}

//GetCurrentRoundInformation Get Current Round Information   //已完成
func (a *AElfClient) GetCurrentRoundInformation() (*dto.RoundDto, error) {
	url := a.Host + ROUNDINFORMATION
	roundBytes, err := GetRequest("GET", url, a.Version, nil)
	if err != nil {
		return nil, err
	}
	var round = new(dto.RoundDto)
	json.Unmarshal(roundBytes, &round)
	return round, nil
}

//GetChainID Get Chain ID
func (a *AElfClient) GetChainID() (uint16, error) {
	chainStatus, err := a.GetChainStatus()
	if err != nil {
		return 0, err
	}
	chainIDBytes := base58.Decode(chainStatus.ChainId)
	return binary.BigEndian.Uint16(chainIDBytes), nil
}

//GetTaskQueueStatus Get Task Queue Status  //已完成
func (a *AElfClient) GetTaskQueueStatus() ([]*dto.TaskQueueInfoDto, error) {
	url := a.Host + TASKQUEUESTATUS
	queues, err := GetRequest("GET", url, a.Version, nil)
	if err != nil {
		return nil, err
	}
	var datas interface{}
	json.Unmarshal(queues, &datas)
	var queueInfos []*dto.TaskQueueInfoDto
	for _, data := range datas.([]interface{}) {
		var queue = new(dto.TaskQueueInfoDto)
		queueBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(queueBytes, &queue)
		queueInfos = append(queueInfos, queue)
	}

	return queueInfos, nil
}

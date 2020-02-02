package client

import (
	"encoding/json"

	"aelf_sdk.go/aelf_sdk/dto"
	util "aelf_sdk.go/aelf_sdk/utils"
	"github.com/btcsuite/btcutil/base58"
)

//GetChainStatus Get the current status of the block chain
func (a *AElfClient) GetChainStatus() (*dto.ChainStatusDto, error) {
	url := a.Host + CHAINSTATUS
	chainBytes, err := util.GetRequest("GET", url, a.Version, nil)
	if err != nil {
		return nil, err
	}
	var chain = new(dto.ChainStatusDto)
	json.Unmarshal(chainBytes, &chain)
	return chain, nil
}

//GetContractFileDescriptorSet Get the definitions of proto-buff related to a contract
func (a *AElfClient) GetContractFileDescriptorSet(address string) ([]byte, error) {
	url := a.Host + FILEDESCRIPTOR
	params := map[string]interface{}{"address": address}
	data, err := util.GetRequest("GET", url, a.Version, params)
	if err != nil {
		return nil, err
	}
	return data, err
}

//GetCurrentRoundInformation Get the latest round of consensus information from data on the last blockHeader of best-chain.
func (a *AElfClient) GetCurrentRoundInformation() (*dto.RoundDto, error) {
	url := a.Host + ROUNDINFORMATION
	roundBytes, err := util.GetRequest("GET", url, a.Version, nil)
	if err != nil {
		return nil, err
	}
	var round = new(dto.RoundDto)
	json.Unmarshal(roundBytes, &round)
	return round, nil
}

//GetChainID Get id of the chain
func (a *AElfClient) GetChainID() (int, error) {
	chainStatus, err := a.GetChainStatus()
	if err != nil {
		return 0, err
	}
	chainIDBytes := base58.Decode(chainStatus.ChainId)

	if len(chainIDBytes) < 4 {
		var bs [4]byte
		for i := 0; i < 4; i++ {
			bs[i] = 0
			if len(chainIDBytes) > i {
				bs[i] = chainIDBytes[i]
			}
		}
		chainIDBytes = bs[:]
	}
	return util.BytesToInt(chainIDBytes), nil
}

//GetTaskQueueStatus Get the status information of the task queue
func (a *AElfClient) GetTaskQueueStatus() ([]*dto.TaskQueueInfoDto, error) {
	url := a.Host + TASKQUEUESTATUS
	queues, err := util.GetRequest("GET", url, a.Version, nil)
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

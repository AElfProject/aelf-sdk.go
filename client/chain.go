package client

import (
	"encoding/json"
	"errors"

	"github.com/AElfProject/aelf-sdk.go/dto"
	util "github.com/AElfProject/aelf-sdk.go/utils"

	"github.com/btcsuite/btcutil/base58"
)

//GetChainStatus Get the current status of the block chain.
func (a *AElfClient) GetChainStatus() (*dto.ChainStatusDto, error) {
	url := a.Host + CHAINSTATUS
	chainBytes, err := util.GetRequest("GET", url, a.Version, nil)
	if err != nil {
		return nil, errors.New("Get ChainStatus error:" + err.Error())
	}
	var chain = new(dto.ChainStatusDto)
	json.Unmarshal(chainBytes, &chain)
	return chain, nil
}

//GetContractFileDescriptorSet Get the definitions of proto-buff related to a contract.
func (a *AElfClient) GetContractFileDescriptorSet(address string) ([]byte, error) {
	url := a.Host + FILEDESCRIPTOR
	params := map[string]interface{}{"address": address}
	data, err := util.GetRequest("GET", url, a.Version, params)
	if err != nil {
		return nil, errors.New("Get ContractFile Descriptor Set error:" + err.Error())
	}
	return data, err
}

//GetChainID Get id of the chain.
func (a *AElfClient) GetChainID() (int, error) {
	chainStatus, err := a.GetChainStatus()
	if err != nil {
		return 0, errors.New("Get Chain Status error:" + err.Error())
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

//GetTaskQueueStatus Get the status information of the task queue.
func (a *AElfClient) GetTaskQueueStatus() ([]*dto.TaskQueueInfoDto, error) {
	url := a.Host + TASKQUEUESTATUS
	queues, err := util.GetRequest("GET", url, a.Version, nil)
	if err != nil {
		return nil, errors.New("Get Task Queue Status error:" + err.Error())
	}
	var datas interface{}
	json.Unmarshal(queues, &datas)
	var queueInfos []*dto.TaskQueueInfoDto
	for _, data := range datas.([]interface{}) {
		var queue = new(dto.TaskQueueInfoDto)
		queueBytes, err := json.Marshal(data)
		if err != nil {
			return nil, errors.New("json Marshal data error:" + err.Error())
		}
		json.Unmarshal(queueBytes, &queue)
		queueInfos = append(queueInfos, queue)
	}
	return queueInfos, nil
}

package client

import (
	"encoding/json"

	"aelf_sdk.go/aelf_sdk/dto"
	util "aelf_sdk.go/aelf_sdk/utils"
)

//GetNetworkInfo Get the node's network information.
func (a *AElfClient) GetNetworkInfo() (*dto.NetworkInfo, error) {
	url := a.Host + NETWORKINFO
	networkBytes, err := util.GetRequest("GET", url, a.Version, nil)
	if err != nil {
		return nil, err
	}
	var network = new(dto.NetworkInfo)
	json.Unmarshal(networkBytes, &network)
	return network, nil
}

//RemovePeer Attempt to remove a node from the connected network nodes by given the ipAddress.
func (a *AElfClient) RemovePeer(ipAddress string) (bool, error) {
	url := a.Host + REMOVEPEER
	params := map[string]interface{}{"address": ipAddress}
	peerBytes, err := util.GetRequest("DELETE", url, a.Version, params)
	if err != nil {
		return false, err
	}
	var data interface{}
	json.Unmarshal(peerBytes, &data)
	return data.(bool), nil
}

//AddPeer Attempt to add a node to the connected network nodes.Input parameter contains the ipAddress of the node.
func (a *AElfClient) AddPeer(ipAddress string) (bool, error) {
	url := a.Host + ADDPEER
	params := map[string]interface{}{"Address": ipAddress}
	peerBytes, err := util.PostRequest(url, a.Version, params)
	if err != nil {
		return false, err
	}
	var data interface{}
	json.Unmarshal(peerBytes, &data)
	return data.(bool), nil
}

//GetPeers Gets information about the peer nodes of the current node.Optional whether to include metrics.
func (a *AElfClient) GetPeers(withMetrics bool) ([]*dto.PeerDto, error) {
	url := a.Host + PEERS
	params := map[string]interface{}{"withMetrics": withMetrics}
	peerBytes, err := util.GetRequest("GET", url, a.Version, params)
	if err != nil {
		return nil, err
	}
	var datas interface{}
	var peers []*dto.PeerDto
	json.Unmarshal(peerBytes, &datas)
	for _, data := range datas.([]interface{}) {
		var peer = new(dto.PeerDto)
		peerBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(peerBytes, &peer)
		peers = append(peers, peer)
	}
	return peers, nil
}

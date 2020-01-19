package client

import (
	"encoding/json"

	"aelf_sdk.go/aelf_sdk/dto"
)

//GetNetworkInfo Get Network Info
func (a *AElfClient) GetNetworkInfo() (*dto.NetworkInfo, error) {
	url := a.Host + NETWORKINFO
	networkBytes, err := GetRequest("GET", url, a.Version, nil)
	if err != nil {
		return nil, err
	}
	var network = new(dto.NetworkInfo)
	json.Unmarshal(networkBytes, &network)
	return network, nil
}

//RemovePeer Remove Peer
func (a *AElfClient) RemovePeer(peerAddress string) (bool, error) {
	url := a.Host + REMOVEPEER
	params := map[string]interface{}{"address": peerAddress}
	peerBytes, err := GetRequest("DELETE", url, a.Version, params)
	if err != nil {
		return false, err
	}
	var data interface{}
	json.Unmarshal(peerBytes, &data)
	return data.(bool), nil
}

//AddPeer Add Peer
func (a *AElfClient) AddPeer(peerAddress string) (interface{}, error) {
	url := a.Host + ADDPEER
	params := map[string]interface{}{"Address": peerAddress}
	peerBytes, err := PostRequest(url, a.Version, params)
	if err != nil {
		return false, err
	}
	var data map[string]interface{}
	json.Unmarshal(peerBytes, &data)
	return data, nil
}

//GetPeers GetPeers
func (a *AElfClient) GetPeers() ([]*dto.PeerDto, error) {
	url := a.Host + PEERS
	peerBytes, err := GetRequest("GET", url, a.Version, nil)
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

package client

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/AElfProject/aelf-sdk.go/dto"
	util "github.com/AElfProject/aelf-sdk.go/utils"
)

//GetNetworkInfo Get the node's network information.
func (a *AElfClient) GetNetworkInfo() (*dto.NetworkInfo, error) {
	url := a.Host + NETWORKINFO
	networkBytes, err := util.GetRequest("GET", url, a.Version, nil)
	if err != nil {
		return nil, errors.New("Get Network Info error:" + err.Error())
	}
	var network = new(dto.NetworkInfo)
	json.Unmarshal(networkBytes, &network)
	return network, nil
}

//RemovePeer Attempt to remove a node from the connected network nodes by given the ipAddress.
func (a *AElfClient) RemovePeer(ipAddress string) (bool, error) {
	url := a.Host + REMOVEPEER
	combine := a.UserName + ":" + a.Password
	combineToBase64 := "Basic " + base64.StdEncoding.EncodeToString([]byte(combine))
	params := map[string]interface{}{"address": ipAddress}
	peerBytes, err := util.GetRequestWithAuth("DELETE", url, a.Version, params, combineToBase64)
	if err != nil {
		return false, errors.New("Remove Peer error:" + err.Error())
	}
	var data interface{}
	json.Unmarshal(peerBytes, &data)
	return data.(bool), nil
}

//AddPeer Attempt to add a node to the connected network nodes.Input parameter contains the ipAddress of the node.
func (a *AElfClient) AddPeer(ipAddress string) (bool, error) {
	url := a.Host + ADDPEER
	combine := a.UserName + ":" + a.Password
	combineToBase64 := "Basic " + base64.StdEncoding.EncodeToString([]byte(combine))
	params := map[string]interface{}{"Address": ipAddress}
	peerBytes, err := util.PostRequestWithAuth(url, a.Version, params, combineToBase64)
	if err != nil {
		return false, errors.New("Add Peer error:" + err.Error())
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
		return nil, errors.New("Get Peers error:" + err.Error())
	}
	var datas interface{}
	var peers []*dto.PeerDto
	json.Unmarshal(peerBytes, &datas)
	for _, data := range datas.([]interface{}) {
		var peer = new(dto.PeerDto)
		peerBytes, _ := json.Marshal(data)
		json.Unmarshal(peerBytes, &peer)
		peers = append(peers, peer)
	}
	return peers, nil
}

package client

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/AElfProject/aelf-sdk.go/dto"
	util "github.com/AElfProject/aelf-sdk.go/utils"
)

// GetNetworkInfo Get the node's network information.
func (client *AElfClient) GetNetworkInfo() (*dto.NetworkInfo, error) {
	url := client.Host + NETWORKINFO
	networkBytes, err := util.GetRequest("GET", url, client.Version, nil)
	if err != nil {
		return nil, errors.New("Get Network Info error:" + err.Error())
	}
	var network = new(dto.NetworkInfo)
	json.Unmarshal(networkBytes, &network)
	return network, nil
}

// RemovePeer Attempt to remove a node from the connected network nodes by given the ipAddress.
func (client *AElfClient) RemovePeer(ipAddress string) (bool, error) {
	url := client.Host + REMOVEPEER
	combine := client.UserName + ":" + client.Password
	combineToBase64 := "Basic " + base64.StdEncoding.EncodeToString([]byte(combine))
	params := map[string]interface{}{"address": ipAddress}
	peerBytes, err := util.GetRequestWithAuth("DELETE", url, client.Version, params, combineToBase64)
	if err != nil {
		return false, errors.New("Remove Peer error:" + err.Error())
	}
	var data interface{}
	json.Unmarshal(peerBytes, &data)
	return data.(bool), nil
}

// AddPeer Attempt to add a node to the connected network nodes.Input parameter contains the ipAddress of the node.
func (client *AElfClient) AddPeer(ipAddress string) (bool, error) {
	url := client.Host + ADDPEER
	combine := client.UserName + ":" + client.Password
	combineToBase64 := "Basic " + base64.StdEncoding.EncodeToString([]byte(combine))
	params := map[string]interface{}{"Address": ipAddress}
	peerBytes, err := util.PostRequestWithAuth(url, client.Version, params, combineToBase64)
	if err != nil {
		return false, errors.New("Add Peer error:" + err.Error())
	}
	var data interface{}
	json.Unmarshal(peerBytes, &data)
	return data.(bool), nil
}

// GetPeers Gets information about the peer nodes of the current node.Optional whether to include metrics.
func (client *AElfClient) GetPeers(withMetrics bool) ([]*dto.PeerDto, error) {
	url := client.Host + PEERS
	params := map[string]interface{}{"withMetrics": withMetrics}
	peerBytes, err := util.GetRequest("GET", url, client.Version, params)
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

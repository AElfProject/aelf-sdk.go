package test

import (
	"encoding/base64"
	"encoding/hex"
	"testing"

	"github.com/AElfProject/aelf-sdk.go/client"
	"github.com/AElfProject/aelf-sdk.go/dto"
	"github.com/AElfProject/aelf-sdk.go/extension"

	pb "github.com/AElfProject/aelf-sdk.go/protobuf/generated"
	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/proto"
	secp256 "github.com/haltingstate/secp256k1-go"

	"github.com/stretchr/testify/assert"
)

var aelf = client.AElfClient{
	Host:       "http://127.0.0.1:8000",
	Version:    "1.0",
	PrivateKey: "4980f789235c786462b24ef23849e90bb1b6d590402a130313ce88d79a6ca3da",
}

var _address = aelf.GetAddressFromPrivateKey(aelf.PrivateKey)

func TestGetAddressFromPubKey(t *testing.T) {
	privateKeyBytes, _ := hex.DecodeString(aelf.PrivateKey)
	pubkeyBytes := secp256.UncompressedPubkeyFromSeckey(privateKeyBytes)
	pubKeyAddress := aelf.GetAddressFromPubKey(hex.EncodeToString(pubkeyBytes))
	assert.Equal(t, _address, pubKeyAddress)
	spew.Dump("Get Address From Public Key", pubKeyAddress)
}

func TestGetFormattedAddress(t *testing.T) {
	formattedAddress, err := aelf.GetFormattedAddress(_address)
	assert.NoError(t, err)
	spew.Dump("Get Formatted Address result", formattedAddress, err)

	privateKeyaddress := aelf.GetAddressFromPrivateKey(aelf.PrivateKey)
	assert.Equal(t, formattedAddress, ("ELF_" + privateKeyaddress + "_AELF"))
}

func TestGenerateKeyPairInfo(t *testing.T) {
	keyPair := aelf.GenerateKeyPairInfo()
	spew.Dump("Generate KeyPair Info Result", keyPair)
}

func TestGetContractAddressByName(t *testing.T) {
	contractAddress, err := aelf.GetContractAddressByName("AElf.ContractNames.Token")
	assert.NoError(t, err)
	spew.Dump("Get ContractAddress By Name Result", contractAddress)
}

func TestClient(t *testing.T) {
	//Test IsConnected
	isConnected := aelf.IsConnected()
	assert.True(t, isConnected)

	wrongClient := client.AElfClient{
		Host:       "http://127.0.0.1:8008",
		Version:    "1.0",
		PrivateKey: "cd86ab6347d8e52bbbe8532141fc59ce596268143a308d1d40fedf385528b458",
	}

	isConnected = wrongClient.IsConnected()
	assert.False(t, isConnected)
}

func TestGetTransactionFee(t *testing.T) {
	var result dto.TransactionResultDto
	var logEventDto dto.LogEventDto
	logEventDto.Name = "TransactionFeeCharged"
	var param = &pb.TransactionFeeCharged{Symbol: "ELF", Amount: 1000}
	paramBytes, _ := proto.Marshal(param)
	logEventDto.NonIndexed = base64.StdEncoding.EncodeToString(paramBytes)
	result.Logs = append(result.Logs, logEventDto)

	logEventDto.Name = "ResourceTokenCharged"
	var params = &pb.ResourceTokenCharged{Symbol: "READ", Amount: 800}
	paramsBytes, _ := proto.Marshal(params)
	logEventDto.NonIndexed = base64.StdEncoding.EncodeToString(paramsBytes)
	result.Logs = append(result.Logs, logEventDto)

	logEventDto.Name = "ResourceTokenCharged"
	params = &pb.ResourceTokenCharged{Symbol: "WRITE", Amount: 600}
	paramsBytes, _ = proto.Marshal(params)
	logEventDto.NonIndexed = base64.StdEncoding.EncodeToString(paramsBytes)
	result.Logs = append(result.Logs, logEventDto)

	logEventDto.Name = "ResourceTokenCharged"
	params = &pb.ResourceTokenCharged{Symbol: "READ", Amount: 200}
	paramsBytes, _ = proto.Marshal(params)
	logEventDto.NonIndexed = base64.StdEncoding.EncodeToString(paramsBytes)
	result.Logs = append(result.Logs, logEventDto)

	res, _ := extension.GetTransactionFees(result)
	assert.Equal(t, int64(1000), res["TransactionFeeCharged"][0]["ELF"])
	assert.Equal(t, int64(800), res["ResourceTokenCharged"][0]["READ"])
	assert.Equal(t, int64(600), res["ResourceTokenCharged"][1]["WRITE"])
	assert.Equal(t, int64(200), res["ResourceTokenCharged"][2]["READ"])
}

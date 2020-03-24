package test

import (
	"encoding/base64"
	"encoding/hex"
	"testing"

	"aelf-sdk.go/client"
	"aelf-sdk.go/dto"
	"aelf-sdk.go/extension"

	pb "aelf-sdk.go/protobuf/generated"
	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/proto"
	secp256 "github.com/haltingstate/secp256k1-go"

	"github.com/stretchr/testify/assert"
)

var aelf = client.AElfClient{
	Host:       "http://127.0.0.1:8000",
	Version:    "1.0",
	PrivateKey: "680afd630d82ae5c97942c4141d60b8a9fedfa5b2864fca84072c17ee1f72d9d",
}

var _address = aelf.GetAddressFromPrivateKey(aelf.PrivateKey, false)

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

	privateKeyaddress := aelf.GetAddressFromPrivateKey(aelf.PrivateKey, false)
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
	assert.True(t, isConnected == true)

	//Test GetGenesisContractAddress
	contractAddr, err := aelf.GetGenesisContractAddress()
	assert.NoError(t, err)
	spew.Dump("Get Genesis Contract Address Result", contractAddr)
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

package test

import (
	"encoding/hex"
	"testing"

	"aelf-sdk.go/client"
	"aelf-sdk.go/extension"
	pb "aelf-sdk.go/protobuf/generated"
	util "aelf-sdk.go/utils"

	"github.com/davecgh/go-spew/spew"
	proto "github.com/golang/protobuf/proto"
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
	toAccount := "2DyzHMD1DqurK9hhiPa91mTBEtcPNrPvY5Uh7tnqRMXGnB381R"
	toAddress, _ := aelf.GetContractAddressByName("AElf.ContractNames.Token")
	methodName := "TransferFrom"
	fromAddressBytes, _ := util.Base58StringToAddress(_address)
	toAddressBytes, _ := util.Base58StringToAddress(toAccount)
	var param = &pb.TransferFromInput{
		From:   fromAddressBytes,
		To:     toAddressBytes,
		Symbol: "ELF",
		Amount: 10000,
	}
	paramBytes, _ := proto.Marshal(param)
	transaction, _ := aelf.CreateTransaction(_address, toAddress, methodName, paramBytes)
	signature, err := aelf.SignTransaction(aelf.PrivateKey, transaction)
	transaction.Signature = signature
	assert.NoError(t, err)
	transactionByets, _ := proto.Marshal(transaction)
	result, err := aelf.SendTransaction(hex.EncodeToString(transactionByets))
	assert.NoError(t, err)

	transactionResult, _ := aelf.GetTransactionResult(result.TransactionID)
	res := extension.GetTransactionFees(transactionResult)
	spew.Dump("Get Transaction Fee Result", res)
}

package example

import (
	"encoding/hex"
	"errors"
	pb "github.com/AElfProject/aelf-sdk.go/protobuf/generated"

	"github.com/AElfProject/aelf-sdk.go/client"
	"github.com/AElfProject/aelf-sdk.go/dto"
	"github.com/AElfProject/aelf-sdk.go/utils"
	"google.golang.org/protobuf/encoding/protojson"

	secp256 "github.com/skycoin/skycoin/src/cipher/secp256k1-go"
)

var aelf = client.AElfClient{
	Host:       "http://127.0.0.1:8000", //your host
	Version:    "1.0",
	PrivateKey: "680afd630d82ae5c97942c4141d60b8a9fedfa5b2864fca84072c17ee1f72d9d", //your private key
}

var privatekeyAddress = aelf.GetAddressFromPrivateKey(aelf.PrivateKey)
var contractMethodName = "GetContractAddressByName"
var contractAddress, _ = aelf.GetGenesisContractAddress()

// DemoGetBlockByHash Get Block ByHash demo.
func DemoGetBlockByHash() (*dto.BlockDto, error) {
	var includeTransactions = true
	height, _ := aelf.GetBlockHeight()
	HeightBlock, _ := aelf.GetBlockByHeight(height, includeTransactions)
	byHashBlock, err := aelf.GetBlockByHash(HeightBlock.BlockHash, includeTransactions)
	if err != nil {
		return nil, errors.New("Get Block By Hash error: " + err.Error())
	}
	return byHashBlock, nil
}

// DemoGetAddressFromPubKey Get Address From Public Key demo.
func DemoGetAddressFromPubKey() string {
	privateKeyBytes, _ := hex.DecodeString(aelf.PrivateKey)
	pubkeyBytes := secp256.UncompressedPubkeyFromSeckey(privateKeyBytes)
	pubKeyAddress := aelf.GetAddressFromPubKey(hex.EncodeToString(pubkeyBytes))
	return pubKeyAddress
}

// DemoExecuteRawTransaction ExecuteRawTransaction demo.
func DemoExecuteRawTransaction() (string, error) {
	chainStatus, err := aelf.GetChainStatus()
	params := &pb.Hash{
		Value: utils.GetBytesSha256("AElf.ContractNames.Token"),
	}
	paramsByte, _ := protojson.Marshal(params)
	var input = &dto.CreateRawTransactionInput{
		From:           privatekeyAddress,
		To:             contractAddress,
		MethodName:     contractMethodName,
		RefBlockNumber: chainStatus.BestChainHeight,
		RefBlockHash:   chainStatus.BestChainHash,
		Params:         string(paramsByte),
	}
	createRaw, err := aelf.CreateRawTransaction(input)
	if err != nil {
		return "", errors.New("create rawTransaction error: " + err.Error())
	}
	rawTransactionBytes, err := hex.DecodeString(createRaw.RawTransaction)
	signature, _ := client.GetSignatureWithPrivateKey(aelf.PrivateKey, rawTransactionBytes)
	var executeRawinput = &dto.ExecuteRawTransactionDto{
		RawTransaction: createRaw.RawTransaction,
		Signature:      signature,
	}
	executeRawresult, err := aelf.ExecuteRawTransaction(executeRawinput)
	if err != nil {
		return "", errors.New("execute rawTransaction error: " + err.Error())
	}
	return executeRawresult, nil
}

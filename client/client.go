package client

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/AElfProject/aelf-sdk.go/dto"
	"github.com/AElfProject/aelf-sdk.go/model"
	pb "github.com/AElfProject/aelf-sdk.go/protobuf/generated"
	util "github.com/AElfProject/aelf-sdk.go/utils"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	secp256 "github.com/haltingstate/secp256k1-go"
	wrap "google.golang.org/protobuf/types/known/wrapperspb"
)

// AElfClient AElf Client.
type AElfClient struct {
	Host       string
	Version    string
	PrivateKey string
	UserName   string
	Password   string
}

// const const.
const (
	CHAINSTATUS             = "/api/blockChain/chainStatus"
	BLOCKHEIGHT             = "/api/blockChain/blockHeight"
	BLOCKBYHASH             = "/api/blockChain/block"
	BLOCKBYHEIGHT           = "/api/blockChain/blockByHeight"
	TRANSACTIONPOOLSTATUS   = "/api/blockChain/transactionPoolStatus"
	RAWTRANSACTION          = "/api/blockChain/rawTransaction"
	SENDTRANSACTION         = "/api/blockChain/sendTransaction"
	SENDRAWTRANSACTION      = "/api/blockChain/sendRawTransaction"
	TASKQUEUESTATUS         = "/api/blockChain/taskQueueStatus"
	TRANSACTIONRESULT       = "/api/blockChain/transactionResult"
	TRANSACTIONRESULTS      = "/api/blockChain/transactionResults"
	MBYTRANSACTIONID        = "/api/blockChain/merklePathByTransactionId"
	ADDPEER                 = "/api/net/peer"
	REMOVEPEER              = "/api/net/peer"
	PEERS                   = "/api/net/peers"
	NETWORKINFO             = "/api/net/networkInfo"
	SENDTRANSACTIONS        = "/api/blockChain/sendTransactions"
	EXECUTETRANSACTION      = "/api/blockChain/executeTransaction"
	EXECUTERAWTRANSACTION   = "/api/blockChain/executeRawTransaction"
	FILEDESCRIPTOR          = "/api/blockChain/contractFileDescriptorSet"
	CALCULATETRANSACTIONFEE = "/api/blockChain/calculateTransactionFee"

	ExamplePrivateKey = "680afd630d82ae5c97942c4141d60b8a9fedfa5b2864fca84072c17ee1f72d9d"
)

// GetAddressFromPubKey Get the account address through the public key.
func (client *AElfClient) GetAddressFromPubKey(pubkey string) string {
	bytes, _ := hex.DecodeString(pubkey)
	return util.GetAddressByBytes(bytes)
}

// GetAddressFromPrivateKey Get the account address through the private key.
func (client *AElfClient) GetAddressFromPrivateKey(privateKey string) string {
	bytes, _ := hex.DecodeString(privateKey)
	pubkeyBytes := secp256.UncompressedPubkeyFromSeckey(bytes)
	return util.GetAddressByBytes(pubkeyBytes)
}

// GetFormattedAddress Convert the Address to the displayed string:symbol_base58-string_base58-string-chain-id.
func (client *AElfClient) GetFormattedAddress(address string) (string, error) {
	chain, _ := client.GetChainStatus()
	methodName := "GetPrimaryTokenSymbol"
	fromAddress := client.GetAddressFromPrivateKey(ExamplePrivateKey)
	contractAddress, _ := client.GetContractAddressByName("AElf.ContractNames.Token")
	transaction, _ := client.CreateTransaction(fromAddress, contractAddress, methodName, nil)
	signature, _ := client.SignTransaction(ExamplePrivateKey, transaction)
	transaction.Signature = signature
	transactionBytes, err := proto.Marshal(transaction)
	if err != nil {
		return "", errors.New("proto marshasl transaction error" + err.Error())
	}
	executeResult, _ := client.ExecuteTransaction(hex.EncodeToString(transactionBytes))
	var symbol = new(wrap.StringValue)
	executeBytes, err := hex.DecodeString(executeResult)
	proto.Unmarshal(executeBytes, symbol)
	return symbol.Value + "_" + address + "_" + chain.ChainId, nil
}

// GetContractAddressByName Get  contract address by contract name.
func (client *AElfClient) GetContractAddressByName(contractName string) (string, error) {
	fromAddress := client.GetAddressFromPrivateKey(ExamplePrivateKey)
	toAddress, err := client.GetGenesisContractAddress()
	if err != nil {
		return "", errors.New("Get Genesis Contract Address error")
	}
	contractNameBytes := util.GetBytesSha256(contractName)
	var hash = new(pb.Hash)
	hash.Value = contractNameBytes
	hashBytes, _ := proto.Marshal(hash)

	transaction, _ := client.CreateTransaction(fromAddress, toAddress, "GetContractAddressByName", hashBytes)
	signature, _ := client.SignTransaction(ExamplePrivateKey, transaction)
	transaction.Signature = signature
	transactionBytes, err := proto.Marshal(transaction)
	if err != nil {
		return "", errors.New("proto marshasl transaction error" + err.Error())
	}
	result, _ := client.ExecuteTransaction(hex.EncodeToString(transactionBytes))
	var address = new(pb.Address)
	resultBytes, err := hex.DecodeString(result)
	proto.Unmarshal(resultBytes, address)
	return util.EncodeCheck(address.Value), nil
}

// SignTransaction Sign a transaction using private key.
func (client *AElfClient) SignTransaction(privateKey string, transaction *pb.Transaction) ([]byte, error) {
	transactionBytes, _ := proto.Marshal(transaction)
	txDataBytes := sha256.Sum256(transactionBytes)
	privateKeyBytes, _ := hex.DecodeString(privateKey)
	signatureBytes := secp256.Sign(txDataBytes[:], privateKeyBytes)
	return signatureBytes, nil
}

// CreateTransaction create a transaction from the input parameters.
func (client *AElfClient) CreateTransaction(from, to, method string, params []byte) (*pb.Transaction, error) {
	chainStatus, err := client.GetChainStatus()
	if err != nil {
		return nil, errors.New("Get Chain Status error ")
	}
	prefixBytes, _ := hex.DecodeString(chainStatus.BestChainHash)
	fromAddressBytes, _ := util.Base58StringToAddress(from)
	toAddressBytes, _ := util.Base58StringToAddress(to)
	var transaction = &pb.Transaction{
		From:           fromAddressBytes,
		To:             toAddressBytes,
		MethodName:     method,
		RefBlockNumber: chainStatus.BestChainHeight,
		RefBlockPrefix: prefixBytes[:4],
		Params:         params,
	}
	return transaction, nil
}

// GetGenesisContractAddress Get the address of genesis contract.
func (client *AElfClient) GetGenesisContractAddress() (string, error) {
	chainStatus, err := client.GetChainStatus()
	if err != nil {
		return "", errors.New("Get Genesis Contract Address error:" + err.Error())
	}
	address := chainStatus.GenesisContractAddress
	return address, nil
}

// IsConnected Verify whether this sdk successfully connects the chain.
func (client *AElfClient) IsConnected() bool {
	data, err := client.GetChainStatus()
	if err != nil || data == nil {
		return false
	}
	return true
}

// GenerateKeyPairInfo Generate KeyPair Info.
func (client *AElfClient) GenerateKeyPairInfo() *model.KeyPairInfo {
	publicKeyBytes, privateKeyBytes := secp256.GenerateKeyPair()
	publicKey := hex.EncodeToString(publicKeyBytes)
	privateKey := hex.EncodeToString(privateKeyBytes)
	privateKeyAddress := client.GetAddressFromPrivateKey(privateKey)
	var keyPair = &model.KeyPairInfo{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Address:    privateKeyAddress,
	}
	return keyPair
}

// GetSignatureWithPrivateKey Get Signature With PrivateKey.
func GetSignatureWithPrivateKey(privateKey string, txData []byte) (string, error) {
	privateKeyBytes, _ := hex.DecodeString(privateKey)
	txDataBytes := sha256.Sum256(txData)
	signatureBytes := secp256.Sign(txDataBytes[:], privateKeyBytes)
	return hex.EncodeToString(signatureBytes), nil
}

// callWriteContract 发送交易并轮询获取交易结果
func (client *AElfClient) CallWriteContract(contractAddress string, methodName string, params interface{}) (*dto.TransactionResultDto, error) {

	paramsByte, err := proto.Marshal(params.(proto.Message))
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}

	transaction, _ := client.CreateTransaction(client.GetAddressFromPrivateKey(client.PrivateKey), contractAddress, methodName, paramsByte)
	signature, _ := client.SignTransaction(client.PrivateKey, transaction)
	transaction.Signature = signature

	// Send the transfer transaction to AElf chain node.
	transactionByets, _ := proto.Marshal(transaction)
	sendResult, err := client.SendTransaction(hex.EncodeToString(transactionByets))
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}

	var times = 0
	for {
		if times > 5 {
			return nil, fmt.Errorf("transaction not mined within time limit")
		}
		transactionResult, err := client.GetTransactionResult(sendResult.TransactionID)
		if err != nil {
			return nil, fmt.Errorf("failed to get transaction result: %w", err)
		}

		if transactionResult.Status == "MINED" {
			return transactionResult, nil
		}

		time.Sleep(1 * time.Second)
		times++
	}
}

func (client *AElfClient) CallViewContract(contractAddress, methodName string, params interface{}) (json.RawMessage, error) {
	// Step 1: Get the chain status and contract address
	chainStatus, err := client.GetChainStatus()
	if err != nil {
		return nil, err
	}

	// Step 2: Prepare the parameters for the contract call
	var paramsBytes []byte
	if params != nil {
		if msg, ok := params.(proto.Message); ok {
			paramsBytes, err = protojson.Marshal(msg)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("params is not of type proto.Message")
		}
	}

	// Step 3: Create a raw transaction input
	input := &dto.CreateRawTransactionInput{
		From:           client.GetAddressFromPrivateKey(client.PrivateKey),
		To:             contractAddress,
		MethodName:     methodName,
		RefBlockNumber: chainStatus.BestChainHeight,
		RefBlockHash:   chainStatus.BestChainHash,
		Params:         string(paramsBytes),
	}

	createRaw, err := client.CreateRawTransaction(input)
	if err != nil {
		return nil, err
	}

	// Step 4: Sign the raw transaction
	rawTransactionBytes, err := hex.DecodeString(createRaw.RawTransaction)
	if err != nil {
		return nil, err
	}

	signature, err := GetSignatureWithPrivateKey(client.PrivateKey, rawTransactionBytes)
	if err != nil {
		return nil, err
	}

	// Step 5: Execute the raw transaction
	executeRawInput := &dto.ExecuteRawTransactionDto{
		RawTransaction: createRaw.RawTransaction,
		Signature:      signature,
	}

	executeRawResult, err := client.ExecuteRawTransaction(executeRawInput)
	if err != nil {
		return nil, err
	}

	// Log the result for debugging
	fmt.Printf("Transaction result: %s\n", executeRawResult)

	return json.RawMessage(executeRawResult), nil
}

package client

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/AElfProject/aelf-sdk.go/model"
	pb "github.com/AElfProject/aelf-sdk.go/protobuf/generated"
	util "github.com/AElfProject/aelf-sdk.go/utils"

	proto "github.com/golang/protobuf/proto"
	wrap "github.com/golang/protobuf/ptypes/wrappers"
	secp256 "github.com/haltingstate/secp256k1-go"
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
	CHAINSTATUS           = "/api/blockChain/chainStatus"
	BLOCKHEIGHT           = "/api/blockChain/blockHeight"
	BLOCKBYHASH           = "/api/blockChain/block"
	BLOCKBYHEIGHT         = "/api/blockChain/blockByHeight"
	TRANSACTIONPOOLSTATUS = "/api/blockChain/transactionPoolStatus"
	RAWTRANSACTION        = "/api/blockChain/rawTransaction"
	SENDTRANSACTION       = "/api/blockChain/sendTransaction"
	SENDRAWTRANSACTION    = "/api/blockChain/sendRawTransaction"
	TASKQUEUESTATUS       = "/api/blockChain/taskQueueStatus"
	TRANSACTIONRESULT     = "/api/blockChain/transactionResult"
	TRANSACTIONRESULTS    = "/api/blockChain/transactionResults"
	MBYTRANSACTIONID      = "/api/blockChain/merklePathByTransactionId"
	ADDPEER               = "/api/net/peer"
	REMOVEPEER            = "/api/net/peer"
	PEERS                 = "/api/net/peers"
	NETWORKINFO           = "/api/net/networkInfo"
	SENDTRANSACTIONS      = "/api/blockChain/sendTransactions"
	EXECUTETRANSACTION    = "/api/blockChain/executeTransaction"
	EXECUTERAWTRANSACTION = "/api/blockChain/executeRawTransaction"
	FILEDESCRIPTOR        = "/api/blockChain/contractFileDescriptorSet"
	TRANSACTIONFEERESULT  = "/api/blockChain/calculateTransactionFee"

	ExamplePrivateKey = "680afd630d82ae5c97942c4141d60b8a9fedfa5b2864fca84072c17ee1f72d9d"
)

// GetAddressFromPubKey Get the account address through the public key.
func (a *AElfClient) GetAddressFromPubKey(pubkey string) string {
	bytes, _ := hex.DecodeString(pubkey)
	return util.GetAddressByBytes(bytes)
}

// GetAddressFromPrivateKey Get the account address through the private key.
func (a *AElfClient) GetAddressFromPrivateKey(privateKey string) string {
	bytes, _ := hex.DecodeString(privateKey)
	pubkeyBytes := secp256.UncompressedPubkeyFromSeckey(bytes)
	return util.GetAddressByBytes(pubkeyBytes)
}

// GetFormattedAddress Convert the Address to the displayed string:symbol_base58-string_base58-string-chain-id.
func (a *AElfClient) GetFormattedAddress(address string) (string, error) {
	chain, _ := a.GetChainStatus()
	methodName := "GetPrimaryTokenSymbol"
	fromAddress := a.GetAddressFromPrivateKey(ExamplePrivateKey)
	contractAddress, _ := a.GetContractAddressByName("AElf.ContractNames.Token")
	transaction, _ := a.CreateTransaction(fromAddress, contractAddress, methodName, nil)
	signature, _ := a.SignTransaction(ExamplePrivateKey, transaction)
	transaction.Signature = signature
	transactionBytes, err := proto.Marshal(transaction)
	if err != nil {
		return "", errors.New("proto marshasl transaction error" + err.Error())
	}
	executeResult, _ := a.ExecuteTransaction(hex.EncodeToString(transactionBytes))
	var symbol = new(wrap.StringValue)
	executeBytes, err := hex.DecodeString(executeResult)
	proto.Unmarshal(executeBytes, symbol)
	return symbol.Value + "_" + address + "_" + chain.ChainId, nil
}

// GetContractAddressByName Get  contract address by contract name.
func (a *AElfClient) GetContractAddressByName(contractName string) (string, error) {
	fromAddress := a.GetAddressFromPrivateKey(ExamplePrivateKey)
	toAddress, err := a.GetGenesisContractAddress()
	if err != nil {
		return "", errors.New("Get Genesis Contract Address error")
	}
	contractNameBytes := util.GetBytesSha256(contractName)
	var hash = new(pb.Hash)
	hash.Value = contractNameBytes
	hashBytes, _ := proto.Marshal(hash)

	transaction, _ := a.CreateTransaction(fromAddress, toAddress, "GetContractAddressByName", hashBytes)
	signature, _ := a.SignTransaction(ExamplePrivateKey, transaction)
	transaction.Signature = signature
	transactionBytes, err := proto.Marshal(transaction)
	if err != nil {
		return "", errors.New("proto marshasl transaction error" + err.Error())
	}
	result, _ := a.ExecuteTransaction(hex.EncodeToString(transactionBytes))
	var address = new(pb.Address)
	resultBytes, err := hex.DecodeString(result)
	proto.Unmarshal(resultBytes, address)
	return util.EncodeCheck(address.Value), nil
}

// SignTransaction Sign a transaction using private key.
func (a *AElfClient) SignTransaction(privateKey string, transaction *pb.Transaction) ([]byte, error) {
	transactionBytes, _ := proto.Marshal(transaction)
	txDataBytes := sha256.Sum256(transactionBytes)
	privateKeyBytes, _ := hex.DecodeString(privateKey)
	signatureBytes := secp256.Sign(txDataBytes[:], privateKeyBytes)
	return signatureBytes, nil
}

// CreateTransaction create a transaction from the input parameters.
func (a *AElfClient) CreateTransaction(from, to, method string, params []byte) (*pb.Transaction, error) {
	chainStatus, err := a.GetChainStatus()
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
func (a *AElfClient) GetGenesisContractAddress() (string, error) {
	chainStatus, err := a.GetChainStatus()
	if err != nil {
		return "", errors.New("Get Genesis Contract Address error:" + err.Error())
	}
	address := chainStatus.GenesisContractAddress
	return address, nil
}

// IsConnected Verify whether this sdk successfully connects the chain.
func (a *AElfClient) IsConnected() bool {
	data, err := a.GetChainStatus()
	if err != nil || data == nil {
		return false
	}
	return true
}

// GenerateKeyPairInfo Generate KeyPair Info.
func (a *AElfClient) GenerateKeyPairInfo() *model.KeyPairInfo {
	publicKeyBytes, privateKeyBytes := secp256.GenerateKeyPair()
	publicKey := hex.EncodeToString(publicKeyBytes)
	privateKey := hex.EncodeToString(privateKeyBytes)
	privateKeyAddress := a.GetAddressFromPrivateKey(privateKey)
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

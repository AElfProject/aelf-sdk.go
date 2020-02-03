package client

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	pt "aelf_sdk.go/aelf_sdk/proto"
	util "aelf_sdk.go/aelf_sdk/utils"
	proto "github.com/golang/protobuf/proto"

	"github.com/davecgh/go-spew/spew"
	secp256 "github.com/skycoin/skycoin/src/cipher/secp256k1-go"

	secp256k1 "github.com/ethereum/go-ethereum/crypto/secp256k1"
)

//AElfClient AElf Client
type AElfClient struct {
	Host       string
	Version    string
	PrivateKey string
	PublicKey  string
}

//const const
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
	ROUNDINFORMATION      = "/api/blockChain/currentRoundInformation"
)

//GetAddressFromPubKey Get the account address through the public key
func (a *AElfClient) GetAddressFromPubKey(pubkey string) string {
	bytes, _ := hex.DecodeString(pubkey)
	pubkeyBytes1 := sha256.Sum256(bytes)
	pubkeyBytes2 := sha256.Sum256(pubkeyBytes1[:])
	address := util.EncodeCheck(pubkeyBytes2[:])
	return address
}

//GetAddressFromPrivateKey Get the account address through the private key
func (a *AElfClient) GetAddressFromPrivateKey(privateKey string, compress bool) string {
	var pubkeyBytes []byte
	bytes, _ := hex.DecodeString(privateKey)
	if compress {
		pubkeyBytes = secp256.PubkeyFromSeckey(bytes)
	} else {
		pubkeyBytes = secp256.UncompressedPubkeyFromSeckey(bytes)
	}
	pubkeyBytes1 := sha256.Sum256(pubkeyBytes)
	pubkeyBytes2 := sha256.Sum256(pubkeyBytes1[:])
	address := util.EncodeCheck(pubkeyBytes2[:])
	return address
}

// GetFormattedAddress Convert the Address to the displayed string：symbol_base58-string_base58-string-chain-id
func (a *AElfClient) GetFormattedAddress(privateKey, address string) (string, error) {
	chain, err := a.GetChainStatus()
	if err != nil {
		return "", errors.New("get chain Status error" + err.Error())
	}
	methodName := "GetPrimaryTokenSymbol"
	fromAddress := a.GetAddressFromPrivateKey(privateKey, false)
	toAddress, err := a.GetContractAddressByName(privateKey, util.GetBytesSha256("AElf.ContractNames.Token"))
	if err != nil {
		return "", errors.New("Get Contract Address By Name error" + err.Error())
	}
	transaction, err := a.CreateTransaction(fromAddress, toAddress, methodName, nil)
	if err != nil {
		return "", errors.New("Create transaction error")
	}
	signature, err := a.SignTransaction(privateKey, transaction)
	if err != nil {
		return "", errors.New("sign transaction error" + err.Error())
	}
	transaction.Signature = signature
	transactionBytes, err := proto.Marshal(transaction)
	if err != nil {
		return "", errors.New("proto marshasl transaction error" + err.Error())
	}

	executeResult, err := a.ExecuteTransaction(hex.EncodeToString(transactionBytes))
	spew.Dump(">>>>>>>>>>>>result", executeResult, err)
	if err != nil {
		return "", errors.New("Execute Transaction error" + err.Error())
	}
	// var symbol = new(wrap.StringValue)
	// json.Unmarshal(util.StringToBytes(executeResult), &symbol)
	// return symbol.GetValue() + address + "_" + chain.ChainId, nil
	return chain.ChainId, nil
}

//GetContractAddressByName Get  contract address by contract name
func (a *AElfClient) GetContractAddressByName(privateKey string, contractName []byte) (string, error) {
	fromAddress := a.GetAddressFromPrivateKey(privateKey, false)
	toAddress, err := a.GetGenesisContractAddress()
	if err != nil {
		return "", errors.New("Get Genesis Contract Address error")
	}
	methodName := "GetContractAddressByName"
	transaction, err := a.CreateTransaction(fromAddress, toAddress, methodName, contractName)
	if err != nil {
		return "", errors.New("Create transaction error")
	}
	signature, err := a.SignTransaction(privateKey, transaction)
	if err != nil {
		return "", errors.New("sign transaction error" + err.Error())
	}
	transaction.Signature = signature
	transactionBytes, err := proto.Marshal(transaction)
	if err != nil {
		return "", errors.New("proto marshasl transaction error" + err.Error())
	}
	result, err := a.ExecuteTransaction(hex.EncodeToString(transactionBytes))
	if err != nil {
		return "", errors.New("Execute Transaction error" + err.Error())
	}
	return result, nil
}

//SignTransaction Sign a transaction using private key
func (a *AElfClient) SignTransaction(privateKey string, transaction *pt.Transaction) ([]byte, error) {
	transactionBytes, _ := proto.Marshal(transaction)
	txDataBytes := sha256.Sum256(transactionBytes)
	privateKeyBytes, _ := hex.DecodeString(privateKey)
	signatureBytes, err := secp256k1.Sign(txDataBytes[:], privateKeyBytes)
	if err != nil {
		return nil, errors.New("Get Signature With PrivateKey error: " + err.Error())
	}
	return signatureBytes, nil
}

//GetSignatureWithPrivateKey Get Signature With PrivateKey
func GetSignatureWithPrivateKey(privateKey string, txData []byte) (string, error) {
	privateKeyBytes, _ := hex.DecodeString(privateKey)
	txDataBytes := sha256.Sum256(txData)
	signatureBytes, err := secp256k1.Sign(txDataBytes[:], privateKeyBytes)
	if err != nil {
		return "", errors.New("Get Signature With PrivateKey error: " + err.Error())
	}
	return hex.EncodeToString(signatureBytes), nil
}

//CreateTransaction create a transaction from the input parameters
func (a *AElfClient) CreateTransaction(from, to, method string, params []byte) (*pt.Transaction, error) {
	chainStatus, err := a.GetChainStatus()
	if err != nil {
		return nil, errors.New("Get Chain Status error ")
	}
	prefixBytes, _ := hex.DecodeString(chainStatus.BestChainHash)
	fromAddressBytes, _ := util.Base58StringToAddress(from)
	toAddressBytes, _ := util.Base58StringToAddress(to)
	var transaction = &pt.Transaction{
		From:           fromAddressBytes,
		To:             toAddressBytes,
		MethodName:     method,
		RefBlockNumber: chainStatus.BestChainHeight,
		RefBlockPrefix: prefixBytes[:4],
	}
	var hash = new(pt.Hash)
	hash.Value = params
	hashBytes, _ := proto.Marshal(hash)
	if params != nil {
		transaction.Params = hashBytes
	}
	return transaction, nil
}

//GetGenesisContractAddress Get the address of genesis contract
func (a *AElfClient) GetGenesisContractAddress() (string, error) {
	chainStatus, err := a.GetChainStatus()
	if err != nil {
		return "", err
	}
	address := chainStatus.GenesisContractAddress
	return address, nil
}

//IsConnected Verify whether this sdk successfully connects the chain
func (a *AElfClient) IsConnected() bool {
	data, err := a.GetChainStatus()
	if err != nil || data == nil {
		return false
	}
	return true
}

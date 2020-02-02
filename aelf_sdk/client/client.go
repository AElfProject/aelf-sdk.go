package client

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"

	pt "aelf_sdk.go/aelf_sdk/proto"
	util "aelf_sdk.go/aelf_sdk/utils"
	proto "github.com/golang/protobuf/proto"
	wrap "github.com/golang/protobuf/ptypes/wrappers"

	secp256 "github.com/skycoin/skycoin/src/cipher/secp256k1-go"
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
		return "", errors.New("json marshasl transaction error" + err.Error())
	}

	result, err := a.ExecuteTransaction(hex.EncodeToString(transactionBytes))
	if err != nil {
		return "", errors.New("Execute Transaction error" + err.Error())
	}
	var symbol = new(wrap.StringValue)
	json.Unmarshal(util.StringToBytes(result), &symbol)
	return symbol.GetValue() + address + "_" + chain.ChainId, nil
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
	transactionBytes, err := json.Marshal(&transaction)
	if err != nil {
		return "", errors.New("json marshasl transaction error" + err.Error())
	}
	result, err := a.ExecuteTransaction(hex.EncodeToString(transactionBytes))
	if err != nil {
		return "", errors.New("Execute Transaction error" + err.Error())
	}
	return result, nil
}

//SignTransaction Sign a transaction using private key
func (a *AElfClient) SignTransaction(privateKey string, transaction *pt.Transaction) ([]byte, error) {
	transactionBytes, err := json.Marshal(transaction)
	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, errors.New("marshasl transaction error: " + err.Error())
	}
	signatureBytes := secp256.Sign(transactionBytes, privateKeyBytes)
	return signatureBytes, nil
}

//GetSignatureWithPrivateKey Get Signature With PrivateKey
func GetSignatureWithPrivateKey(privateKey string, txData []byte) string {
	privateKeyBytes, _ := hex.DecodeString(privateKey)
	signatureBytes := secp256.Sign(txData, privateKeyBytes)
	return hex.EncodeToString(signatureBytes)
}

//CreateTransaction create a transaction from the input parameters
func (a *AElfClient) CreateTransaction(from, to, method string, params []byte) (*pt.Transaction, error) {
	chainStatus, err := a.GetChainStatus()
	if err != nil {
		return nil, errors.New("Get Chain Status error ")
	}
	var transaction = new(pt.Transaction)
	transaction.MethodName = method
	if params != nil {
		transaction.Params = params
	}
	transaction.ProtoMessage()
	transaction.To, _ = util.Base58StringToAddress(to)
	transaction.From, _ = util.Base58StringToAddress(from)
	transaction.RefBlockNumber = chainStatus.BestChainHeight
	bytes, _ := hex.DecodeString(chainStatus.BestChainHash)
	transaction.RefBlockPrefix = bytes[0:4]
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

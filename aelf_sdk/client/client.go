package client

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

//GetAddressFromPubKey Get Address From Public Key
// func (a *AElfClient) GetAddressFromPubKey(pubkey string) (*proto.Address, error) {
// 	var address = new(proto.Address)
// 	h := sha256.New()
// 	h.Write([]byte(pubkey))
// 	h.Write(h.Sum(nil))
// 	address.value = h.Sum(nil)
// 	return address, nil
// }

// //GetFormattedAddress Get Formatted Address
// func (a *AElfClient) GetFormattedAddress(address *proto.Address) {

// }

//GetSystemContractAddress Get system contract address by contract name
// func (a *AElfClient) GetSystemContractAddress(contractName string) (*proto.Address, error) {
// 	toAddress, err := a.GetGenesisContractAddress()
// 	if err != nil {
// 		return "", errors.New("Get Genesis Contract Address error")
// 	}

// 	Sha256.hashTwice(publicKey)

// }

//SignTransaction Sign Transaction
// func (a *AElfClient) SignTransaction(privateKey string, transaction *proto.Transaction) {

// }

// //CreateTransaction Create Transaction
// func (a *AElfClient) CreateTransaction(from, to, method string, params []byte) (*proto.Transaction, error) {
// 	chainStatus, err := a.GetChainStatus()
// 	if err != nil {
// 		return nil, errors.New("Get Chain Status error ")
// 	}
// 	chainHash := chainStatus.BestChainHash
// 	chainHeight := chainStatus.BestChainHeight

// 	var transaction = new(proto.Transaction)
// 	transaction.refBlockNumber = chainHeight
// 	transaction.refBlockPrefix = []byte(chainHash)

// 	// Address from = 1;
// 	// Address to = 2;
// 	transaction.methodName = method
// 	transaction.params = params
// 	// bytes signature = 10000;
// 	return transaction, nil
// }

//GetGenesisContractAddress Get Genesis Contract Address
// func (a *AElfClient) GetGenesisContractAddress() (string, error) { //已完成
// 	chainStatus, err := a.GetChainStatus()
// 	if err != nil {
// 		return "", err
// 	}
// 	address := chainStatus.GenesisContractAddress
// 	return address, nil
// }

//IsConnected IsConnected
func (a *AElfClient) IsConnected() bool { //已完成
	data, err := a.GetChainStatus()
	if err != nil || data == nil {
		return false
	}
	return true
}

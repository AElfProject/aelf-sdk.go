

# AElf-Client

## Introduction

This is a Golang client library, used to communicate with the [AElf](https://github.com/AElfProject/AElf)  API.

### Getting Started

You should build the "AElf_SDK" project 


### Basic usage

``` Golang
private const string BaseUrl = "Http://127.0.0.1:8000";

// get client instance
AElfClient aelf = client.AElfClient{Host:BaseUrl}
var height =  aelf.GetBlockHeight();
```


#### BlockService

```Golang
 func (a *AElfClient) GetBlockHeight() (float64, error);

 func (a *AElfClient) GetBlockByHash(blockHash string, isTransactions bool) (BlockDto, error);

 func (a *AElfClient) GetBlockByHeight(blockHeight int, isTransactions bool) (BlockDto, error);
```

#### ChainService

```Golang
 func (a *AElfClient) GetChainStatus() (ChainStatusDto, error);

 func (a *AElfClient) GetContractFileDescriptorSet(address string) ([]byte, error);

 func (a *AElfClient) GetChainID() (int, error);

 func (a *AElfClient) GetTaskQueueStatus() ([]TaskQueueInfoDto, error);
```

#### NetworkService

```Golang
 func (a *AElfClient) AddPeer(ipAddress string) (bool), error);

 func (a *AElfClient) RemovePeer(address string) (bool, error);

 func (a *AElfClient) GetPeers(withMetrics bool) ([]PeerDto, error);

 func (a *AElfClient) GetNetworkInfo() (NetworkInfo, error);

```

#### TransactionService

```Golang

func (a *AElfClient) GetTransactionPoolStatus() (TransactionPoolStatus, error);

func (a *AElfClient) GetTransactionResult(transactionID string) (TransactionResultDto, error);

func (a *AElfClient) GetTransactionResults(blockHash string, offset = 0, limit = 10) ([]TransactionResultDto, error);

func (a *AElfClient) GetMerklePathByTransactionID(transactionID string) (*dto.MerklePathDto, error);

func (a *AElfClient) ExecuteTransaction(rawTransaction string) (string, error);

func (a *AElfClient) ExecuteRawTransaction(input ExecuteRawTransactionDto) (map[string]interface{}, error);

func (a *AElfClient) SendTransaction(transaction string) (SendTransactionOutput, error);

func (a *AElfClient) CreateRawTransaction(input CreateRawTransactionInput) (CreateRawTransactionOutput, error);

func (a *AElfClient) SendRawTransaction(input SendRawTransactionInput) (SendRawTransactionOutput, error);

func (a *AElfClient) SendTransactions(rawTransactions string) ([]string, error);

```

#### ClientService

```Golang
 func (a *AElfClient) IsConnected() bool;

 func (a *AElfClient) GetFormattedAddress(address string) (string, error);

 func (a *AElfClient) GetAddressFromPubKey(pubkey string) string;

 func (a *AElfClient) GetAddressFromPrivateKey(privateKey string, compress = false bool) string;

 func (a *AElfClient) GetGenesisContractAddress() (string, error);

 func (a *AElfClient) GetContractAddressByName(contractName []byte) (string, error);

```

### Test

This module contains tests for all services provided by AElfClient. You can see how to properly use services provided by AElfClient here.

You need to firstly set necessary parameters to make sure tests can run successfully.

1. Set baseUrl to your target url.

   ```Golang
   string BaseUrl = "Http://127.0.0.1:8000";
   ```

2. Give a valid privateKey of a node.

   ```Golang
   string PrivateKey = "680afd630d82ae5c97942c4141d60b8a9fedfa5b2864fca84072c17ee1f72d9d";
   ```

### ProtoBuff Build
Compile protobuff before secondary development.
You need to go to the resource directory.

Linxu Or Mac ./protobuff.sh

### Note

You need to run a local or remote AElf node to run the unit test successfully. If you're not familiar with how to run a node or multiple nodes, please see [Running a node](https://docs.aelf.io/v/dev/main/main/run-node) / [Running multiple nodes](https://docs.aelf.io/v/dev/main/main/multi-nodes) for more information.

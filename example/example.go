package example

import (
	"encoding/hex"
	"fmt"
	"github.com/AElfProject/aelf-sdk.go/client"
	"github.com/AElfProject/aelf-sdk.go/model/consts"
	pb "github.com/AElfProject/aelf-sdk.go/protobuf/generated"
	"github.com/AElfProject/aelf-sdk.go/utils"
	"google.golang.org/protobuf/proto"
)

const (
	MainChainUrl = "https://aelf-public-node.aelf.io"
	SideChainUrl = "https://tdvv-public-node.aelf.io"
)

var mainChainClient = client.AElfClient{
	Host:    MainChainUrl,
	Version: "1.0",
}

var sideChainClient = client.AElfClient{
	Host:    SideChainUrl,
	Version: "1.0",
}

func getTransferred(txId string) []*pb.Transferred {
	transffereds := make([]*pb.Transferred, 0)
	result, err := mainChainClient.GetTransactionResult(txId)
	if err != nil || len(result.Logs) == 0 {
		return transffereds
	}

	contractAddr, _ := mainChainClient.GetContractAddressByName(consts.TokenContractSystemName)

	for _, log := range result.Logs {
		if log.Name == consts.TransferredLogEventName && log.Address == contractAddr {
			transferred := new(pb.Transferred)
			if nonIndexedBytes, err := utils.Base64DecodeBytes(log.NonIndexed); err == nil {
				proto.Unmarshal(nonIndexedBytes, transferred)
			}
			if fromBytes, err := utils.Base64DecodeBytes(log.Indexed[0]); err == nil {
				temp := new(pb.Transferred)
				proto.Unmarshal(fromBytes, temp)
				transferred.From = temp.From
			}
			if toBytes, err := utils.Base64DecodeBytes(log.Indexed[1]); err == nil {
				temp := new(pb.Transferred)
				proto.Unmarshal(toBytes, temp)
				transferred.To = temp.To
			}
			if symbolBytes, err := utils.Base64DecodeBytes(log.Indexed[2]); err == nil {
				temp := new(pb.Transferred)
				proto.Unmarshal(symbolBytes, temp)
				transferred.Symbol = temp.Symbol
			}
			transffereds = append(transffereds, transferred)
		}
	}

	return transffereds
}

func getCrossChainTransferred(txId string) []*pb.CrossChainTransferred {
	crossChainTransferreds := make([]*pb.CrossChainTransferred, 0)
	result, err := sideChainClient.GetTransactionResult(txId)
	if err != nil || len(result.Logs) == 0 {
		return crossChainTransferreds
	}

	contractAddr, _ := sideChainClient.GetContractAddressByName(consts.TokenContractSystemName)

	for _, log := range result.Logs {
		if log.Name == consts.CrossChainTransferredLogEventName && log.Address == contractAddr {
			crossChainTransferred := new(pb.CrossChainTransferred)
			if nonIndexedBytes, err := utils.Base64DecodeBytes(log.NonIndexed); err == nil {
				proto.Unmarshal(nonIndexedBytes, crossChainTransferred)
			}
			crossChainTransferreds = append(crossChainTransferreds, crossChainTransferred)
		}
	}

	return crossChainTransferreds
}

func getCrossChainReceived(txId string) []*pb.CrossChainReceived {
	crossChainReceiveds := make([]*pb.CrossChainReceived, 0)
	result, err := mainChainClient.GetTransactionResult(txId)
	if err != nil || len(result.Logs) == 0 {
		return crossChainReceiveds
	}

	contractAddr, _ := mainChainClient.GetContractAddressByName(consts.TokenContractSystemName)

	for _, log := range result.Logs {
		if log.Name == consts.CrossChainReceivedLogEventName && log.Address == contractAddr {
			crossChainReceived := new(pb.CrossChainReceived)
			if nonIndexedBytes, err := utils.Base64DecodeBytes(log.NonIndexed); err == nil {
				proto.Unmarshal(nonIndexedBytes, crossChainReceived)
			}
			crossChainReceiveds = append(crossChainReceiveds, crossChainReceived)
		}
	}

	return crossChainReceiveds
}

func getTokenBalance(symbol, owner string) (*pb.GetBalanceOutput, error) {
	tokenContractAddr, _ := mainChainClient.GetContractAddressByName(consts.TokenContractSystemName)
	addr := mainChainClient.GetAddressFromPrivateKey(mainChainClient.PrivateKey)
	ownerAddr, err := utils.Base58StringToAddress(owner)
	if err != nil {
		return &pb.GetBalanceOutput{}, err
	}
	inputByte, _ := proto.Marshal(&pb.GetBalanceInput{
		Symbol: symbol,
		Owner:  ownerAddr,
	})

	tx, _ := mainChainClient.CreateTransaction(addr, tokenContractAddr, consts.TokenContractGetBalance, inputByte)
	tx.Signature, _ = mainChainClient.SignTransaction(mainChainClient.PrivateKey, tx)

	txByets, _ := proto.Marshal(tx)
	re, _ := mainChainClient.ExecuteTransaction(hex.EncodeToString(txByets))

	balance := &pb.GetBalanceOutput{}
	bytes, _ := hex.DecodeString(re)
	proto.Unmarshal(bytes, balance)

	return balance, nil
}

func getTokenInfo(symbol string) (*pb.TokenInfo, error) {
	tokenContractAddr, _ := mainChainClient.GetContractAddressByName(consts.TokenContractSystemName)
	addr := mainChainClient.GetAddressFromPrivateKey(mainChainClient.PrivateKey)
	inputByte, _ := proto.Marshal(&pb.TokenInfo{
		Symbol: symbol,
	})

	tx, _ := mainChainClient.CreateTransaction(addr, tokenContractAddr, consts.TokenContractGetTokenInfo, inputByte)
	tx.Signature, _ = mainChainClient.SignTransaction(mainChainClient.PrivateKey, tx)

	txBytes, _ := proto.Marshal(tx)
	re, _ := mainChainClient.ExecuteTransaction(hex.EncodeToString(txBytes))

	tokenInfo := &pb.TokenInfo{}
	bytes, _ := hex.DecodeString(re)
	proto.Unmarshal(bytes, tokenInfo)

	return tokenInfo, nil
}

func main() {
	// handle transferred log event
	transferTxId := "5c2b267f436b7b50f53acb7f6ebc8221f4167405c042862155734c414c63c501"
	transferreds := getTransferred(transferTxId)

	for _, t := range transferreds {
		fmt.Printf("Transferred logevent params: from:%s, to:%s, symbol:%s, amount:%d, memo:%s\n",
			utils.AddressToBase58String(t.From), utils.AddressToBase58String(t.To), t.Symbol, t.Amount, t.Memo)
	}

	// handle crossChainTransfer log event
	crossChainTransferTxId := "0cc31ff44f14d4f155c8bc09b6e2fed4dcbe923c049df4e27c2a14831d5af031"
	crossChainTransferred := getCrossChainTransferred(crossChainTransferTxId)

	for _, t := range crossChainTransferred {
		fmt.Printf("CrossChainTransferred logevent params: from:%s, to:%s, symbol:%s, amount:%d, memo:%s, "+
			"toChainId:%d, issueChainId:%d\n", utils.AddressToBase58String(t.From), utils.AddressToBase58String(t.To),
			t.Symbol, t.Amount, t.Memo, t.ToChainId, t.IssueChainId)
	}

	// handle crossChainReceived log event
	crossChainReceivedTxId := "df731ace1caec3d2d047c5dd03997a2ad1b6e8fc032b40fc073339623031c036"
	crossChainReceived := getCrossChainReceived(crossChainReceivedTxId)

	for _, t := range crossChainReceived {
		fmt.Printf("CrossChainTransferred logevent params: from:%s, to:%s, symbol:%s, amount:%d, memo:%s, "+
			"toChainId:%d, issueChainId:%d, parentChainHeight:%d,transferTransactionId: %s\n",
			utils.AddressToBase58String(t.From), utils.AddressToBase58String(t.To), t.Symbol, t.Amount, t.Memo,
			t.FromChainId, t.IssueChainId, t.ParentChainHeight, hex.EncodeToString(t.TransferTransactionId.GetValue()))
	}

	// get token info
	elf, _ := getTokenInfo("ELF")
	fmt.Printf("ELF symbol:%s decimals:%d tokenName:%s,supply:%d totalSupply:%d issuer:%s issueChainId:%d",
		elf.Symbol, elf.Decimals, elf.TokenName, elf.Supply, elf.TotalSupply, utils.AddressToBase58String(elf.Issuer), elf.IssueChainId)

	// get owner token balance
	userKeyPairInfo := mainChainClient.GenerateKeyPairInfo()
	balance, _ := getTokenBalance("ELF", userKeyPairInfo.Address)
	fmt.Printf("%s ELF Amount: %d.", userKeyPairInfo.Address, balance.Balance)
}

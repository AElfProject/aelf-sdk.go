package extension

import (
	"errors"

	"aelf-sdk.go/dto"
	pb "aelf-sdk.go/protobuf/generated"
	util "aelf-sdk.go/utils"
	proto "github.com/golang/protobuf/proto"
)

//GetTransactionFees Get Transaction Fees
func GetTransactionFees(transactionResultDto dto.TransactionResultDto) (map[string][]map[string]interface{}, error) {
	var feeDicts = map[string][]map[string]interface{}{}
	eventLogs := transactionResultDto.Logs
	if len(eventLogs) == 0 {
		return nil, errors.New("transaction Result Dto not found  Logs error")
	}
	for _, log := range eventLogs {
		nonIndexedBytes, _ := util.Base64DecodeBytes(log.NonIndexed)
		if log.Name == "TransactionFeeCharged" {
			var feeCharged = new(pb.TransactionFeeCharged)
			proto.Unmarshal(nonIndexedBytes, feeCharged)
			var feeMap = map[string]interface{}{feeCharged.Symbol: feeCharged.Amount}
			feeDicts["TransactionFeeCharged"] = append(feeDicts["TransactionFeeCharged"], feeMap)
		}

		if log.Name == "ResourceTokenCharged" {
			var tokenCharged = new(pb.ResourceTokenCharged)
			proto.Unmarshal(nonIndexedBytes, tokenCharged)
			var feeMap = map[string]interface{}{tokenCharged.Symbol: tokenCharged.Amount}
			feeDicts["ResourceTokenCharged"] = append(feeDicts["ResourceTokenCharged"], feeMap)
		}
	}
	return feeDicts, nil
}

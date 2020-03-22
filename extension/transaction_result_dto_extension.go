package extension

import (
	"encoding/json"

	"aelf-sdk.go/dto"
	pb "aelf-sdk.go/protobuf/generated"
	proto "github.com/golang/protobuf/proto"
)

//GetTransactionFees Get Transaction Fees
func GetTransactionFees(transactionResultDto *dto.TransactionResultDto) []map[string]interface{} {
	var feeDicts []map[string]interface{}
	eventLogs := transactionResultDto.Logs
	if len(eventLogs) == 0 {
		return nil
	}
	for _, log := range eventLogs {
		if log.Name == "ResourceTokenCharged" || log.Name == "TransactionFeeCharged" {
			var feeCharged = new(pb.TransactionFeeCharged)
			nonIndexedBytes, _ := json.Marshal(log.NonIndexed)
			proto.Unmarshal(nonIndexedBytes, feeCharged)
			var feeDict = map[string]interface{}{feeCharged.Amount: feeCharged.Symbol}
			feeDicts = append(feeDicts, feeDict)
		}
	}
	return feeDicts
}

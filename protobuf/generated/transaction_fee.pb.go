// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.11.4
// source: transaction_fee.proto

package client

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransactionSizeFeeSymbols struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionSizeFeeSymbolList []*TransactionSizeFeeSymbol `protobuf:"bytes,1,rep,name=transaction_size_fee_symbol_list,json=transactionSizeFeeSymbolList,proto3" json:"transaction_size_fee_symbol_list,omitempty"`
}

func (x *TransactionSizeFeeSymbols) Reset() {
	*x = TransactionSizeFeeSymbols{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_fee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionSizeFeeSymbols) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionSizeFeeSymbols) ProtoMessage() {}

func (x *TransactionSizeFeeSymbols) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_fee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionSizeFeeSymbols.ProtoReflect.Descriptor instead.
func (*TransactionSizeFeeSymbols) Descriptor() ([]byte, []int) {
	return file_transaction_fee_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionSizeFeeSymbols) GetTransactionSizeFeeSymbolList() []*TransactionSizeFeeSymbol {
	if x != nil {
		return x.TransactionSizeFeeSymbolList
	}
	return nil
}

type TransactionSizeFeeSymbol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenSymbol      string `protobuf:"bytes,1,opt,name=token_symbol,json=tokenSymbol,proto3" json:"token_symbol,omitempty"`
	BaseTokenWeight  int32  `protobuf:"varint,2,opt,name=base_token_weight,json=baseTokenWeight,proto3" json:"base_token_weight,omitempty"`
	AddedTokenWeight int32  `protobuf:"varint,3,opt,name=added_token_weight,json=addedTokenWeight,proto3" json:"added_token_weight,omitempty"`
}

func (x *TransactionSizeFeeSymbol) Reset() {
	*x = TransactionSizeFeeSymbol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_fee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionSizeFeeSymbol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionSizeFeeSymbol) ProtoMessage() {}

func (x *TransactionSizeFeeSymbol) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_fee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionSizeFeeSymbol.ProtoReflect.Descriptor instead.
func (*TransactionSizeFeeSymbol) Descriptor() ([]byte, []int) {
	return file_transaction_fee_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionSizeFeeSymbol) GetTokenSymbol() string {
	if x != nil {
		return x.TokenSymbol
	}
	return ""
}

func (x *TransactionSizeFeeSymbol) GetBaseTokenWeight() int32 {
	if x != nil {
		return x.BaseTokenWeight
	}
	return 0
}

func (x *TransactionSizeFeeSymbol) GetAddedTokenWeight() int32 {
	if x != nil {
		return x.AddedTokenWeight
	}
	return 0
}

type TransactionFeeCharged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Amount int64  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransactionFeeCharged) Reset() {
	*x = TransactionFeeCharged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_fee_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionFeeCharged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionFeeCharged) ProtoMessage() {}

func (x *TransactionFeeCharged) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_fee_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionFeeCharged.ProtoReflect.Descriptor instead.
func (*TransactionFeeCharged) Descriptor() ([]byte, []int) {
	return file_transaction_fee_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionFeeCharged) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *TransactionFeeCharged) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type ResourceTokenCharged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol          string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Amount          int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	ContractAddress *Address `protobuf:"bytes,3,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
}

func (x *ResourceTokenCharged) Reset() {
	*x = ResourceTokenCharged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_fee_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceTokenCharged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceTokenCharged) ProtoMessage() {}

func (x *ResourceTokenCharged) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_fee_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceTokenCharged.ProtoReflect.Descriptor instead.
func (*ResourceTokenCharged) Descriptor() ([]byte, []int) {
	return file_transaction_fee_proto_rawDescGZIP(), []int{3}
}

func (x *ResourceTokenCharged) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *ResourceTokenCharged) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ResourceTokenCharged) GetContractAddress() *Address {
	if x != nil {
		return x.ContractAddress
	}
	return nil
}

type ResourceTokenOwned struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol          string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Amount          int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	ContractAddress *Address `protobuf:"bytes,3,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
}

func (x *ResourceTokenOwned) Reset() {
	*x = ResourceTokenOwned{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_fee_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceTokenOwned) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceTokenOwned) ProtoMessage() {}

func (x *ResourceTokenOwned) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_fee_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceTokenOwned.ProtoReflect.Descriptor instead.
func (*ResourceTokenOwned) Descriptor() ([]byte, []int) {
	return file_transaction_fee_proto_rawDescGZIP(), []int{4}
}

func (x *ResourceTokenOwned) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *ResourceTokenOwned) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ResourceTokenOwned) GetContractAddress() *Address {
	if x != nil {
		return x.ContractAddress
	}
	return nil
}

var File_transaction_fee_proto protoreflect.FileDescriptor

var file_transaction_fee_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x65,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a,
	0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a,
	0x19, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x7a, 0x65,
	0x46, 0x65, 0x65, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x12, 0x68, 0x0a, 0x20, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x66,
	0x65, 0x65, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x7a, 0x65, 0x46, 0x65, 0x65,
	0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x1c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x69, 0x7a, 0x65, 0x46, 0x65, 0x65, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x97, 0x01, 0x0a, 0x18, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x7a, 0x65, 0x46, 0x65, 0x65, 0x53, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x2a, 0x0a, 0x11, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x62, 0x61, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x2c, 0x0a, 0x12, 0x61, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x61, 0x64,
	0x64, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x4d,
	0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65,
	0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x04, 0xa0, 0xbb, 0x18, 0x01, 0x22, 0x88, 0x01,
	0x0a, 0x14, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43,
	0x68, 0x61, 0x72, 0x67, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x3a, 0x04, 0xa0, 0xbb, 0x18, 0x01, 0x22, 0x86, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x3a, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3a, 0x04, 0xa0, 0xbb, 0x18,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_fee_proto_rawDescOnce sync.Once
	file_transaction_fee_proto_rawDescData = file_transaction_fee_proto_rawDesc
)

func file_transaction_fee_proto_rawDescGZIP() []byte {
	file_transaction_fee_proto_rawDescOnce.Do(func() {
		file_transaction_fee_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_fee_proto_rawDescData)
	})
	return file_transaction_fee_proto_rawDescData
}

var file_transaction_fee_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_transaction_fee_proto_goTypes = []interface{}{
	(*TransactionSizeFeeSymbols)(nil), // 0: client.TransactionSizeFeeSymbols
	(*TransactionSizeFeeSymbol)(nil),  // 1: client.TransactionSizeFeeSymbol
	(*TransactionFeeCharged)(nil),     // 2: client.TransactionFeeCharged
	(*ResourceTokenCharged)(nil),      // 3: client.ResourceTokenCharged
	(*ResourceTokenOwned)(nil),        // 4: client.ResourceTokenOwned
	(*Address)(nil),                   // 5: client.Address
}
var file_transaction_fee_proto_depIdxs = []int32{
	1, // 0: client.TransactionSizeFeeSymbols.transaction_size_fee_symbol_list:type_name -> client.TransactionSizeFeeSymbol
	5, // 1: client.ResourceTokenCharged.contract_address:type_name -> client.Address
	5, // 2: client.ResourceTokenOwned.contract_address:type_name -> client.Address
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_transaction_fee_proto_init() }
func file_transaction_fee_proto_init() {
	if File_transaction_fee_proto != nil {
		return
	}
	file_options_proto_init()
	file_client_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_transaction_fee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionSizeFeeSymbols); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transaction_fee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionSizeFeeSymbol); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transaction_fee_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionFeeCharged); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transaction_fee_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceTokenCharged); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transaction_fee_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceTokenOwned); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transaction_fee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transaction_fee_proto_goTypes,
		DependencyIndexes: file_transaction_fee_proto_depIdxs,
		MessageInfos:      file_transaction_fee_proto_msgTypes,
	}.Build()
	File_transaction_fee_proto = out.File
	file_transaction_fee_proto_rawDesc = nil
	file_transaction_fee_proto_goTypes = nil
	file_transaction_fee_proto_depIdxs = nil
}

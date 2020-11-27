// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.11.4
// source: aedpos_contract.proto

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

//aedpos_contract
type MinerList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The miners public key list.
	Pubkeys [][]byte `protobuf:"bytes,1,rep,name=pubkeys,proto3" json:"pubkeys,omitempty"`
}

func (x *MinerList) Reset() {
	*x = MinerList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aedpos_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MinerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MinerList) ProtoMessage() {}

func (x *MinerList) ProtoReflect() protoreflect.Message {
	mi := &file_aedpos_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MinerList.ProtoReflect.Descriptor instead.
func (*MinerList) Descriptor() ([]byte, []int) {
	return file_aedpos_contract_proto_rawDescGZIP(), []int{0}
}

func (x *MinerList) GetPubkeys() [][]byte {
	if x != nil {
		return x.Pubkeys
	}
	return nil
}

type PubkeyList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Candidates’ public keys
	Value [][]byte `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *PubkeyList) Reset() {
	*x = PubkeyList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aedpos_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubkeyList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubkeyList) ProtoMessage() {}

func (x *PubkeyList) ProtoReflect() protoreflect.Message {
	mi := &file_aedpos_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubkeyList.ProtoReflect.Descriptor instead.
func (*PubkeyList) Descriptor() ([]byte, []int) {
	return file_aedpos_contract_proto_rawDescGZIP(), []int{1}
}

func (x *PubkeyList) GetValue() [][]byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type MinerListWithRoundNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of miners.
	MinerList *MinerList `protobuf:"bytes,1,opt,name=miner_list,json=minerList,proto3" json:"miner_list,omitempty"`
	// The round number.
	RoundNumber int64 `protobuf:"varint,2,opt,name=round_number,json=roundNumber,proto3" json:"round_number,omitempty"`
}

func (x *MinerListWithRoundNumber) Reset() {
	*x = MinerListWithRoundNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aedpos_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MinerListWithRoundNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MinerListWithRoundNumber) ProtoMessage() {}

func (x *MinerListWithRoundNumber) ProtoReflect() protoreflect.Message {
	mi := &file_aedpos_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MinerListWithRoundNumber.ProtoReflect.Descriptor instead.
func (*MinerListWithRoundNumber) Descriptor() ([]byte, []int) {
	return file_aedpos_contract_proto_rawDescGZIP(), []int{2}
}

func (x *MinerListWithRoundNumber) GetMinerList() *MinerList {
	if x != nil {
		return x.MinerList
	}
	return nil
}

func (x *MinerListWithRoundNumber) GetRoundNumber() int64 {
	if x != nil {
		return x.RoundNumber
	}
	return 0
}

type GetMinerListInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The term number.
	TermNumber int64 `protobuf:"varint,1,opt,name=term_number,json=termNumber,proto3" json:"term_number,omitempty"`
}

func (x *GetMinerListInput) Reset() {
	*x = GetMinerListInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aedpos_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMinerListInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMinerListInput) ProtoMessage() {}

func (x *GetMinerListInput) ProtoReflect() protoreflect.Message {
	mi := &file_aedpos_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMinerListInput.ProtoReflect.Descriptor instead.
func (*GetMinerListInput) Descriptor() ([]byte, []int) {
	return file_aedpos_contract_proto_rawDescGZIP(), []int{3}
}

func (x *GetMinerListInput) GetTermNumber() int64 {
	if x != nil {
		return x.TermNumber
	}
	return 0
}

var File_aedpos_contract_proto protoreflect.FileDescriptor

var file_aedpos_contract_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x65, 0x64, 0x70, 0x6f, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22,
	0x25, 0x0a, 0x09, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x07, 0x70,
	0x75, 0x62, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x22, 0x0a, 0x0a, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x6f, 0x0a, 0x18, 0x4d, 0x69,
	0x6e, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x52, 0x6f, 0x75, 0x6e, 0x64,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x0a, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x09, 0x6d,
	0x69, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x34, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x72, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x65, 0x72, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aedpos_contract_proto_rawDescOnce sync.Once
	file_aedpos_contract_proto_rawDescData = file_aedpos_contract_proto_rawDesc
)

func file_aedpos_contract_proto_rawDescGZIP() []byte {
	file_aedpos_contract_proto_rawDescOnce.Do(func() {
		file_aedpos_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_aedpos_contract_proto_rawDescData)
	})
	return file_aedpos_contract_proto_rawDescData
}

var file_aedpos_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_aedpos_contract_proto_goTypes = []interface{}{
	(*MinerList)(nil),                // 0: client.MinerList
	(*PubkeyList)(nil),               // 1: client.PubkeyList
	(*MinerListWithRoundNumber)(nil), // 2: client.MinerListWithRoundNumber
	(*GetMinerListInput)(nil),        // 3: client.GetMinerListInput
}
var file_aedpos_contract_proto_depIdxs = []int32{
	0, // 0: client.MinerListWithRoundNumber.miner_list:type_name -> client.MinerList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_aedpos_contract_proto_init() }
func file_aedpos_contract_proto_init() {
	if File_aedpos_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aedpos_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MinerList); i {
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
		file_aedpos_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubkeyList); i {
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
		file_aedpos_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MinerListWithRoundNumber); i {
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
		file_aedpos_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMinerListInput); i {
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
			RawDescriptor: file_aedpos_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aedpos_contract_proto_goTypes,
		DependencyIndexes: file_aedpos_contract_proto_depIdxs,
		MessageInfos:      file_aedpos_contract_proto_msgTypes,
	}.Build()
	File_aedpos_contract_proto = out.File
	file_aedpos_contract_proto_rawDesc = nil
	file_aedpos_contract_proto_goTypes = nil
	file_aedpos_contract_proto_depIdxs = nil
}

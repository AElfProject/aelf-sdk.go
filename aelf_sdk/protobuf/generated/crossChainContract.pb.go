// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crossChainContract.proto

package client

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//cross_chain_contract
type VerifyTransactionInput struct {
	TransactionId        *Hash       `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id" form:"transaction_id"`
	Path                 *MerklePath `protobuf:"bytes,2,opt,name=path,proto3" json:"path" form:"path"`
	ParentChainHeight    int64       `protobuf:"zigzag64,3,opt,name=parent_chain_height,json=parentChainHeight,proto3" json:"parent_chain_height" form:"parent_chain_height"`
	VerifiedChainId      int32       `protobuf:"varint,4,opt,name=verified_chain_id,json=verifiedChainId,proto3" json:"verified_chain_id" form:"verified_chain_id"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *VerifyTransactionInput) Reset()         { *m = VerifyTransactionInput{} }
func (m *VerifyTransactionInput) String() string { return proto.CompactTextString(m) }
func (*VerifyTransactionInput) ProtoMessage()    {}
func (*VerifyTransactionInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_e668cd6f553fe010, []int{0}
}

func (m *VerifyTransactionInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyTransactionInput.Unmarshal(m, b)
}
func (m *VerifyTransactionInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyTransactionInput.Marshal(b, m, deterministic)
}
func (m *VerifyTransactionInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyTransactionInput.Merge(m, src)
}
func (m *VerifyTransactionInput) XXX_Size() int {
	return xxx_messageInfo_VerifyTransactionInput.Size(m)
}
func (m *VerifyTransactionInput) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyTransactionInput.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyTransactionInput proto.InternalMessageInfo

func (m *VerifyTransactionInput) GetTransactionId() *Hash {
	if m != nil {
		return m.TransactionId
	}
	return nil
}

func (m *VerifyTransactionInput) GetPath() *MerklePath {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *VerifyTransactionInput) GetParentChainHeight() int64 {
	if m != nil {
		return m.ParentChainHeight
	}
	return 0
}

func (m *VerifyTransactionInput) GetVerifiedChainId() int32 {
	if m != nil {
		return m.VerifiedChainId
	}
	return 0
}

func init() {
	proto.RegisterType((*VerifyTransactionInput)(nil), "client.VerifyTransactionInput")
}

func init() { proto.RegisterFile("crossChainContract.proto", fileDescriptor_e668cd6f553fe010) }

var fileDescriptor_e668cd6f553fe010 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x14, 0x45, 0x89, 0x8e, 0xb3, 0x88, 0xa3, 0x32, 0x11, 0x24, 0xb8, 0x2a, 0x2e, 0xa4, 0xb8, 0xe8,
	0xc2, 0xf9, 0x84, 0x6e, 0xda, 0x85, 0x20, 0x41, 0xdc, 0x96, 0x67, 0x12, 0xcd, 0xc3, 0x92, 0x84,
	0xe4, 0x29, 0xf8, 0x99, 0xfe, 0x91, 0x34, 0x6d, 0x71, 0x56, 0x81, 0x7b, 0x4e, 0x1e, 0xf7, 0x72,
	0xa9, 0x53, 0xc8, 0xb9, 0x75, 0x80, 0xbe, 0x0d, 0x9e, 0x12, 0x68, 0x6a, 0x62, 0x0a, 0x14, 0xc4,
	0x56, 0x8f, 0x68, 0x3d, 0xdd, 0xee, 0xe6, 0x77, 0x4e, 0xef, 0x7e, 0x19, 0xbf, 0x79, 0xb5, 0x09,
	0xdf, 0x7f, 0x5e, 0x12, 0xf8, 0x0c, 0x9a, 0x30, 0xf8, 0xde, 0xc7, 0x2f, 0x12, 0x07, 0x7e, 0x49,
	0xff, 0xd9, 0x80, 0x46, 0xb2, 0x8a, 0xd5, 0xe7, 0x8f, 0xbb, 0x66, 0xb9, 0xd0, 0x41, 0x76, 0xea,
	0xe2, 0xc8, 0xe9, 0x8d, 0xb8, 0xe7, 0x9b, 0x08, 0xe4, 0xe4, 0x49, 0x51, 0xc5, 0xaa, 0x3e, 0xd9,
	0xf4, 0x39, 0xda, 0x67, 0x20, 0xa7, 0x0a, 0x17, 0x0d, 0xbf, 0x8e, 0x90, 0xac, 0xa7, 0x41, 0x4f,
	0x5d, 0x07, 0x67, 0xf1, 0xc3, 0x91, 0x3c, 0xad, 0x58, 0x2d, 0xd4, 0x7e, 0x46, 0x65, 0x45, 0x57,
	0x80, 0x78, 0xe0, 0xfb, 0xef, 0xa9, 0x26, 0x5a, 0xb3, 0xfc, 0x40, 0x23, 0x37, 0x15, 0xab, 0xcf,
	0xd4, 0xd5, 0x0a, 0x8a, 0xdf, 0x9b, 0xb7, 0x6d, 0x99, 0x76, 0xf8, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0x43, 0x48, 0x6b, 0xaa, 0x0c, 0x01, 0x00, 0x00,
}

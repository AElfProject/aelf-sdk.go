// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.11.4
// source: profit_contract.proto

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

//profit_contract
type CreatedSchemeIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The scheme ids.
	SchemeIds []*Hash `protobuf:"bytes,1,rep,name=scheme_ids,json=schemeIds,proto3" json:"scheme_ids,omitempty"`
}

func (x *CreatedSchemeIds) Reset() {
	*x = CreatedSchemeIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profit_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatedSchemeIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatedSchemeIds) ProtoMessage() {}

func (x *CreatedSchemeIds) ProtoReflect() protoreflect.Message {
	mi := &file_profit_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatedSchemeIds.ProtoReflect.Descriptor instead.
func (*CreatedSchemeIds) Descriptor() ([]byte, []int) {
	return file_profit_contract_proto_rawDescGZIP(), []int{0}
}

func (x *CreatedSchemeIds) GetSchemeIds() []*Hash {
	if x != nil {
		return x.SchemeIds
	}
	return nil
}

type GetManagingSchemeIdsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The manager address.
	Manager *Address `protobuf:"bytes,1,opt,name=manager,proto3" json:"manager,omitempty"`
}

func (x *GetManagingSchemeIdsInput) Reset() {
	*x = GetManagingSchemeIdsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profit_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetManagingSchemeIdsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetManagingSchemeIdsInput) ProtoMessage() {}

func (x *GetManagingSchemeIdsInput) ProtoReflect() protoreflect.Message {
	mi := &file_profit_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetManagingSchemeIdsInput.ProtoReflect.Descriptor instead.
func (*GetManagingSchemeIdsInput) Descriptor() ([]byte, []int) {
	return file_profit_contract_proto_rawDescGZIP(), []int{1}
}

func (x *GetManagingSchemeIdsInput) GetManager() *Address {
	if x != nil {
		return x.Manager
	}
	return nil
}

type SchemeBeneficiaryShare struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the sub scheme.
	SchemeId *Hash `protobuf:"bytes,1,opt,name=scheme_id,json=schemeId,proto3" json:"scheme_id,omitempty"`
	// The weight of the sub scheme.
	Shares int64 `protobuf:"varint,2,opt,name=shares,proto3" json:"shares,omitempty"`
}

func (x *SchemeBeneficiaryShare) Reset() {
	*x = SchemeBeneficiaryShare{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profit_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchemeBeneficiaryShare) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchemeBeneficiaryShare) ProtoMessage() {}

func (x *SchemeBeneficiaryShare) ProtoReflect() protoreflect.Message {
	mi := &file_profit_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchemeBeneficiaryShare.ProtoReflect.Descriptor instead.
func (*SchemeBeneficiaryShare) Descriptor() ([]byte, []int) {
	return file_profit_contract_proto_rawDescGZIP(), []int{2}
}

func (x *SchemeBeneficiaryShare) GetSchemeId() *Hash {
	if x != nil {
		return x.SchemeId
	}
	return nil
}

func (x *SchemeBeneficiaryShare) GetShares() int64 {
	if x != nil {
		return x.Shares
	}
	return 0
}

type Scheme struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The virtual address of the scheme.
	VirtualAddress *Address `protobuf:"bytes,1,opt,name=virtual_address,json=virtualAddress,proto3" json:"virtual_address,omitempty"`
	// The total weight of the scheme.
	TotalShares int64 `protobuf:"varint,2,opt,name=total_shares,json=totalShares,proto3" json:"total_shares,omitempty"`
	// The manager of the scheme.
	Manager *Address `protobuf:"bytes,3,opt,name=manager,proto3" json:"manager,omitempty"`
	// The current period.
	CurrentPeriod int64 `protobuf:"varint,4,opt,name=current_period,json=currentPeriod,proto3" json:"current_period,omitempty"`
	// Sub schemes information.
	SubSchemes []*SchemeBeneficiaryShare `protobuf:"bytes,5,rep,name=sub_schemes,json=subSchemes,proto3" json:"sub_schemes,omitempty"`
	// Whether you can directly remove the beneficiary.
	CanRemoveBeneficiaryDirectly bool `protobuf:"varint,6,opt,name=can_remove_beneficiary_directly,json=canRemoveBeneficiaryDirectly,proto3" json:"can_remove_beneficiary_directly,omitempty"`
	// Period of profit distribution.
	ProfitReceivingDuePeriodCount int64 `protobuf:"varint,7,opt,name=profit_receiving_due_period_count,json=profitReceivingDuePeriodCount,proto3" json:"profit_receiving_due_period_count,omitempty"`
	// Whether all the schemes balance will be distributed during distribution each period.
	IsReleaseAllBalanceEveryTimeByDefault bool `protobuf:"varint,8,opt,name=is_release_all_balance_every_time_by_default,json=isReleaseAllBalanceEveryTimeByDefault,proto3" json:"is_release_all_balance_every_time_by_default,omitempty"`
	// The is of the scheme.
	SchemeId *Hash `protobuf:"bytes,9,opt,name=scheme_id,json=schemeId,proto3" json:"scheme_id,omitempty"`
	// Delay distribute period.
	DelayDistributePeriodCount int32 `protobuf:"varint,10,opt,name=delay_distribute_period_count,json=delayDistributePeriodCount,proto3" json:"delay_distribute_period_count,omitempty"`
	// Record the scheme's current total share for deferred distribution of benefits, period -> total shares.
	CachedDelayTotalShares map[int64]int64 `protobuf:"bytes,11,rep,name=cached_delay_total_shares,json=cachedDelayTotalShares,proto3" json:"cached_delay_total_shares,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// The received token symbols.
	ReceivedTokenSymbols []string `protobuf:"bytes,12,rep,name=received_token_symbols,json=receivedTokenSymbols,proto3" json:"received_token_symbols,omitempty"`
}

func (x *Scheme) Reset() {
	*x = Scheme{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profit_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scheme) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scheme) ProtoMessage() {}

func (x *Scheme) ProtoReflect() protoreflect.Message {
	mi := &file_profit_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scheme.ProtoReflect.Descriptor instead.
func (*Scheme) Descriptor() ([]byte, []int) {
	return file_profit_contract_proto_rawDescGZIP(), []int{3}
}

func (x *Scheme) GetVirtualAddress() *Address {
	if x != nil {
		return x.VirtualAddress
	}
	return nil
}

func (x *Scheme) GetTotalShares() int64 {
	if x != nil {
		return x.TotalShares
	}
	return 0
}

func (x *Scheme) GetManager() *Address {
	if x != nil {
		return x.Manager
	}
	return nil
}

func (x *Scheme) GetCurrentPeriod() int64 {
	if x != nil {
		return x.CurrentPeriod
	}
	return 0
}

func (x *Scheme) GetSubSchemes() []*SchemeBeneficiaryShare {
	if x != nil {
		return x.SubSchemes
	}
	return nil
}

func (x *Scheme) GetCanRemoveBeneficiaryDirectly() bool {
	if x != nil {
		return x.CanRemoveBeneficiaryDirectly
	}
	return false
}

func (x *Scheme) GetProfitReceivingDuePeriodCount() int64 {
	if x != nil {
		return x.ProfitReceivingDuePeriodCount
	}
	return 0
}

func (x *Scheme) GetIsReleaseAllBalanceEveryTimeByDefault() bool {
	if x != nil {
		return x.IsReleaseAllBalanceEveryTimeByDefault
	}
	return false
}

func (x *Scheme) GetSchemeId() *Hash {
	if x != nil {
		return x.SchemeId
	}
	return nil
}

func (x *Scheme) GetDelayDistributePeriodCount() int32 {
	if x != nil {
		return x.DelayDistributePeriodCount
	}
	return 0
}

func (x *Scheme) GetCachedDelayTotalShares() map[int64]int64 {
	if x != nil {
		return x.CachedDelayTotalShares
	}
	return nil
}

func (x *Scheme) GetReceivedTokenSymbols() []string {
	if x != nil {
		return x.ReceivedTokenSymbols
	}
	return nil
}

type SchemePeriod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The scheme id.
	SchemeId *Hash `protobuf:"bytes,1,opt,name=scheme_id,json=schemeId,proto3" json:"scheme_id,omitempty"`
	// The period number.
	Period int64 `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
}

func (x *SchemePeriod) Reset() {
	*x = SchemePeriod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profit_contract_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchemePeriod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchemePeriod) ProtoMessage() {}

func (x *SchemePeriod) ProtoReflect() protoreflect.Message {
	mi := &file_profit_contract_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchemePeriod.ProtoReflect.Descriptor instead.
func (*SchemePeriod) Descriptor() ([]byte, []int) {
	return file_profit_contract_proto_rawDescGZIP(), []int{4}
}

func (x *SchemePeriod) GetSchemeId() *Hash {
	if x != nil {
		return x.SchemeId
	}
	return nil
}

func (x *SchemePeriod) GetPeriod() int64 {
	if x != nil {
		return x.Period
	}
	return 0
}

type DistributedProfitsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The total amount of shares in this scheme at the current period.
	TotalShares int64 `protobuf:"varint,1,opt,name=total_shares,json=totalShares,proto3" json:"total_shares,omitempty"`
	// The contributed amount in this scheme at the current period.
	AmountsMap map[string]int64 `protobuf:"bytes,2,rep,name=amounts_map,json=amountsMap,proto3" json:"amounts_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"zigzag64,2,opt,name=value,proto3"`
	// Whether released.
	IsReleased bool `protobuf:"varint,3,opt,name=is_released,json=isReleased,proto3" json:"is_released,omitempty"`
}

func (x *DistributedProfitsInfo) Reset() {
	*x = DistributedProfitsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profit_contract_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributedProfitsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributedProfitsInfo) ProtoMessage() {}

func (x *DistributedProfitsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_profit_contract_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributedProfitsInfo.ProtoReflect.Descriptor instead.
func (*DistributedProfitsInfo) Descriptor() ([]byte, []int) {
	return file_profit_contract_proto_rawDescGZIP(), []int{5}
}

func (x *DistributedProfitsInfo) GetTotalShares() int64 {
	if x != nil {
		return x.TotalShares
	}
	return 0
}

func (x *DistributedProfitsInfo) GetAmountsMap() map[string]int64 {
	if x != nil {
		return x.AmountsMap
	}
	return nil
}

func (x *DistributedProfitsInfo) GetIsReleased() bool {
	if x != nil {
		return x.IsReleased
	}
	return false
}

type GetProfitDetailsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The scheme id.
	SchemeId *Hash `protobuf:"bytes,1,opt,name=scheme_id,json=schemeId,proto3" json:"scheme_id,omitempty"`
	// The address of beneficiary.
	Beneficiary *Address `protobuf:"bytes,2,opt,name=beneficiary,proto3" json:"beneficiary,omitempty"`
}

func (x *GetProfitDetailsInput) Reset() {
	*x = GetProfitDetailsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profit_contract_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfitDetailsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfitDetailsInput) ProtoMessage() {}

func (x *GetProfitDetailsInput) ProtoReflect() protoreflect.Message {
	mi := &file_profit_contract_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfitDetailsInput.ProtoReflect.Descriptor instead.
func (*GetProfitDetailsInput) Descriptor() ([]byte, []int) {
	return file_profit_contract_proto_rawDescGZIP(), []int{6}
}

func (x *GetProfitDetailsInput) GetSchemeId() *Hash {
	if x != nil {
		return x.SchemeId
	}
	return nil
}

func (x *GetProfitDetailsInput) GetBeneficiary() *Address {
	if x != nil {
		return x.Beneficiary
	}
	return nil
}

type ProfitDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The profit information.
	Details []*ProfitDetail `protobuf:"bytes,1,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *ProfitDetails) Reset() {
	*x = ProfitDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profit_contract_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfitDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfitDetails) ProtoMessage() {}

func (x *ProfitDetails) ProtoReflect() protoreflect.Message {
	mi := &file_profit_contract_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfitDetails.ProtoReflect.Descriptor instead.
func (*ProfitDetails) Descriptor() ([]byte, []int) {
	return file_profit_contract_proto_rawDescGZIP(), []int{7}
}

func (x *ProfitDetails) GetDetails() []*ProfitDetail {
	if x != nil {
		return x.Details
	}
	return nil
}

type ProfitDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The start period number.
	StartPeriod int64 `protobuf:"varint,1,opt,name=start_period,json=startPeriod,proto3" json:"start_period,omitempty"`
	// The end period number.
	EndPeriod int64 `protobuf:"varint,2,opt,name=end_period,json=endPeriod,proto3" json:"end_period,omitempty"`
	// The weight of the proceeds on the current period of the scheme.
	Shares int64 `protobuf:"varint,3,opt,name=shares,proto3" json:"shares,omitempty"`
	// The last period number that the beneficiary received the profit.
	LastProfitPeriod int64 `protobuf:"varint,4,opt,name=last_profit_period,json=lastProfitPeriod,proto3" json:"last_profit_period,omitempty"`
	// Whether the weight  has been removed.
	IsWeightRemoved bool `protobuf:"varint,5,opt,name=is_weight_removed,json=isWeightRemoved,proto3" json:"is_weight_removed,omitempty"`
}

func (x *ProfitDetail) Reset() {
	*x = ProfitDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profit_contract_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfitDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfitDetail) ProtoMessage() {}

func (x *ProfitDetail) ProtoReflect() protoreflect.Message {
	mi := &file_profit_contract_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfitDetail.ProtoReflect.Descriptor instead.
func (*ProfitDetail) Descriptor() ([]byte, []int) {
	return file_profit_contract_proto_rawDescGZIP(), []int{8}
}

func (x *ProfitDetail) GetStartPeriod() int64 {
	if x != nil {
		return x.StartPeriod
	}
	return 0
}

func (x *ProfitDetail) GetEndPeriod() int64 {
	if x != nil {
		return x.EndPeriod
	}
	return 0
}

func (x *ProfitDetail) GetShares() int64 {
	if x != nil {
		return x.Shares
	}
	return 0
}

func (x *ProfitDetail) GetLastProfitPeriod() int64 {
	if x != nil {
		return x.LastProfitPeriod
	}
	return 0
}

func (x *ProfitDetail) GetIsWeightRemoved() bool {
	if x != nil {
		return x.IsWeightRemoved
	}
	return false
}

type ClaimProfitsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The scheme id.
	SchemeId *Hash `protobuf:"bytes,1,opt,name=scheme_id,json=schemeId,proto3" json:"scheme_id,omitempty"`
	// The address of beneficiary.
	Beneficiary *Address `protobuf:"bytes,2,opt,name=beneficiary,proto3" json:"beneficiary,omitempty"`
}

func (x *ClaimProfitsInput) Reset() {
	*x = ClaimProfitsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profit_contract_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimProfitsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimProfitsInput) ProtoMessage() {}

func (x *ClaimProfitsInput) ProtoReflect() protoreflect.Message {
	mi := &file_profit_contract_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimProfitsInput.ProtoReflect.Descriptor instead.
func (*ClaimProfitsInput) Descriptor() ([]byte, []int) {
	return file_profit_contract_proto_rawDescGZIP(), []int{9}
}

func (x *ClaimProfitsInput) GetSchemeId() *Hash {
	if x != nil {
		return x.SchemeId
	}
	return nil
}

func (x *ClaimProfitsInput) GetBeneficiary() *Address {
	if x != nil {
		return x.Beneficiary
	}
	return nil
}

var File_profit_contract_proto protoreflect.FileDescriptor

var file_profit_contract_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a,
	0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x49, 0x64,
	0x73, 0x12, 0x2b, 0x0a, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x48,
	0x61, 0x73, 0x68, 0x52, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x49, 0x64, 0x73, 0x22, 0x46,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x49, 0x64, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0x5b, 0x0a, 0x16, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65,
	0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x12, 0x29, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x61, 0x73,
	0x68, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x73, 0x22, 0xbc, 0x06, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x38,
	0x0a, 0x0f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0e, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61,
	0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x07, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x3f, 0x0a,
	0x0b, 0x73, 0x75, 0x62, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x65, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x12, 0x45,
	0x0a, 0x1f, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x62, 0x65, 0x6e,
	0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6c,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1c, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6c, 0x79, 0x12, 0x48, 0x0a, 0x21, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x5f,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x75, 0x65, 0x5f, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x1d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x69, 0x6e,
	0x67, 0x44, 0x75, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x5b, 0x0a, 0x2c, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x6c,
	0x6c, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x72, 0x79, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x62, 0x79, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x25, 0x69, 0x73, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x41, 0x6c, 0x6c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x45, 0x76, 0x65, 0x72, 0x79, 0x54,
	0x69, 0x6d, 0x65, 0x42, 0x79, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x29, 0x0a, 0x09,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x08, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x1d, 0x64, 0x65, 0x6c, 0x61, 0x79,
	0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1a,
	0x64, 0x65, 0x6c, 0x61, 0x79, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x65, 0x0a, 0x19, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x64, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x16, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x64, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x73, 0x12, 0x34, 0x0a, 0x16, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x14, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x1a, 0x49, 0x0a, 0x1b, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x64, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x51, 0x0a, 0x0c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x12, 0x29, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x48,
	0x61, 0x73, 0x68, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0xec, 0x01, 0x0a, 0x16, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x0b, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x5f, 0x6d,
	0x61, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x4d, 0x61, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x64, 0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x12, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x75, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x29, 0x0a,
	0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x08,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x0b, 0x62, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0b,
	0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x22, 0x3f, 0x0a, 0x0d, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2e, 0x0a, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xc2, 0x01, 0x0a,
	0x0c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x69, 0x73, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x64, 0x22, 0x71, 0x0a, 0x11, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74,
	0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x29, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x49,
	0x64, 0x12, 0x31, 0x0a, 0x0b, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0b, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63,
	0x69, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_profit_contract_proto_rawDescOnce sync.Once
	file_profit_contract_proto_rawDescData = file_profit_contract_proto_rawDesc
)

func file_profit_contract_proto_rawDescGZIP() []byte {
	file_profit_contract_proto_rawDescOnce.Do(func() {
		file_profit_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_profit_contract_proto_rawDescData)
	})
	return file_profit_contract_proto_rawDescData
}

var file_profit_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_profit_contract_proto_goTypes = []interface{}{
	(*CreatedSchemeIds)(nil),          // 0: client.CreatedSchemeIds
	(*GetManagingSchemeIdsInput)(nil), // 1: client.GetManagingSchemeIdsInput
	(*SchemeBeneficiaryShare)(nil),    // 2: client.SchemeBeneficiaryShare
	(*Scheme)(nil),                    // 3: client.Scheme
	(*SchemePeriod)(nil),              // 4: client.SchemePeriod
	(*DistributedProfitsInfo)(nil),    // 5: client.DistributedProfitsInfo
	(*GetProfitDetailsInput)(nil),     // 6: client.GetProfitDetailsInput
	(*ProfitDetails)(nil),             // 7: client.ProfitDetails
	(*ProfitDetail)(nil),              // 8: client.ProfitDetail
	(*ClaimProfitsInput)(nil),         // 9: client.ClaimProfitsInput
	nil,                               // 10: client.Scheme.CachedDelayTotalSharesEntry
	nil,                               // 11: client.DistributedProfitsInfo.AmountsMapEntry
	(*Hash)(nil),                      // 12: client.Hash
	(*Address)(nil),                   // 13: client.Address
}
var file_profit_contract_proto_depIdxs = []int32{
	12, // 0: client.CreatedSchemeIds.scheme_ids:type_name -> client.Hash
	13, // 1: client.GetManagingSchemeIdsInput.manager:type_name -> client.Address
	12, // 2: client.SchemeBeneficiaryShare.scheme_id:type_name -> client.Hash
	13, // 3: client.Scheme.virtual_address:type_name -> client.Address
	13, // 4: client.Scheme.manager:type_name -> client.Address
	2,  // 5: client.Scheme.sub_schemes:type_name -> client.SchemeBeneficiaryShare
	12, // 6: client.Scheme.scheme_id:type_name -> client.Hash
	10, // 7: client.Scheme.cached_delay_total_shares:type_name -> client.Scheme.CachedDelayTotalSharesEntry
	12, // 8: client.SchemePeriod.scheme_id:type_name -> client.Hash
	11, // 9: client.DistributedProfitsInfo.amounts_map:type_name -> client.DistributedProfitsInfo.AmountsMapEntry
	12, // 10: client.GetProfitDetailsInput.scheme_id:type_name -> client.Hash
	13, // 11: client.GetProfitDetailsInput.beneficiary:type_name -> client.Address
	8,  // 12: client.ProfitDetails.details:type_name -> client.ProfitDetail
	12, // 13: client.ClaimProfitsInput.scheme_id:type_name -> client.Hash
	13, // 14: client.ClaimProfitsInput.beneficiary:type_name -> client.Address
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_profit_contract_proto_init() }
func file_profit_contract_proto_init() {
	if File_profit_contract_proto != nil {
		return
	}
	file_client_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_profit_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatedSchemeIds); i {
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
		file_profit_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetManagingSchemeIdsInput); i {
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
		file_profit_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchemeBeneficiaryShare); i {
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
		file_profit_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scheme); i {
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
		file_profit_contract_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchemePeriod); i {
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
		file_profit_contract_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributedProfitsInfo); i {
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
		file_profit_contract_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfitDetailsInput); i {
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
		file_profit_contract_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfitDetails); i {
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
		file_profit_contract_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfitDetail); i {
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
		file_profit_contract_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimProfitsInput); i {
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
			RawDescriptor: file_profit_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_profit_contract_proto_goTypes,
		DependencyIndexes: file_profit_contract_proto_depIdxs,
		MessageInfos:      file_profit_contract_proto_msgTypes,
	}.Build()
	File_profit_contract_proto = out.File
	file_profit_contract_proto_rawDesc = nil
	file_profit_contract_proto_goTypes = nil
	file_profit_contract_proto_depIdxs = nil
}

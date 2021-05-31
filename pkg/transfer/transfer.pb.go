// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: transfer.proto

package transfer

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

type ACTION int32

const (
	ACTION_START_AGENT ACTION = 0
	ACTION_STOP_AGENT  ACTION = 1
	ACTION_REPORT_DATA ACTION = 2
	ACTION_FINISH_PLAN ACTION = 3
)

// Enum value maps for ACTION.
var (
	ACTION_name = map[int32]string{
		0: "START_AGENT",
		1: "STOP_AGENT",
		2: "REPORT_DATA",
		3: "FINISH_PLAN",
	}
	ACTION_value = map[string]int32{
		"START_AGENT": 0,
		"STOP_AGENT":  1,
		"REPORT_DATA": 2,
		"FINISH_PLAN": 3,
	}
)

func (x ACTION) Enum() *ACTION {
	p := new(ACTION)
	*p = x
	return p
}

func (x ACTION) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ACTION) Descriptor() protoreflect.EnumDescriptor {
	return file_transfer_proto_enumTypes[0].Descriptor()
}

func (ACTION) Type() protoreflect.EnumType {
	return &file_transfer_proto_enumTypes[0]
}

func (x ACTION) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ACTION.Descriptor instead.
func (ACTION) EnumDescriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{0}
}

type STATUS int32

const (
	STATUS_WAITING STATUS = 0
	STATUS_SUCCESS STATUS = 1
	STATUS_FAILURE STATUS = 2
	STATUS_RUNNING STATUS = 3
	STATUS_ERROR   STATUS = 4
)

// Enum value maps for STATUS.
var (
	STATUS_name = map[int32]string{
		0: "WAITING",
		1: "SUCCESS",
		2: "FAILURE",
		3: "RUNNING",
		4: "ERROR",
	}
	STATUS_value = map[string]int32{
		"WAITING": 0,
		"SUCCESS": 1,
		"FAILURE": 2,
		"RUNNING": 3,
		"ERROR":   4,
	}
)

func (x STATUS) Enum() *STATUS {
	p := new(STATUS)
	*p = x
	return p
}

func (x STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_transfer_proto_enumTypes[1].Descriptor()
}

func (STATUS) Type() protoreflect.EnumType {
	return &file_transfer_proto_enumTypes[1]
}

func (x STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use STATUS.Descriptor instead.
func (STATUS) EnumDescriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{1}
}

type Mail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action  ACTION `protobuf:"varint,1,opt,name=action,proto3,enum=ACTION" json:"action,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Mail) Reset() {
	*x = Mail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mail) ProtoMessage() {}

func (x *Mail) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mail.ProtoReflect.Descriptor instead.
func (*Mail) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *Mail) GetAction() ACTION {
	if x != nil {
		return x.Action
	}
	return ACTION_START_AGENT
}

func (x *Mail) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type Plan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TreeName string `protobuf:"bytes,1,opt,name=treeName,proto3" json:"treeName,omitempty"`
	RobotNum int32  `protobuf:"varint,2,opt,name=robotNum,proto3" json:"robotNum,omitempty"`
	Parallel int32  `protobuf:"varint,3,opt,name=parallel,proto3" json:"parallel,omitempty"`
	Interval int32  `protobuf:"varint,4,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *Plan) Reset() {
	*x = Plan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plan) ProtoMessage() {}

func (x *Plan) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plan.ProtoReflect.Descriptor instead.
func (*Plan) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *Plan) GetTreeName() string {
	if x != nil {
		return x.TreeName
	}
	return ""
}

func (x *Plan) GetRobotNum() int32 {
	if x != nil {
		return x.RobotNum
	}
	return 0
}

func (x *Plan) GetParallel() int32 {
	if x != nil {
		return x.Parallel
	}
	return 0
}

func (x *Plan) GetInterval() int32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image    string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	AgentNum int32  `protobuf:"varint,2,opt,name=agentNum,proto3" json:"agentNum,omitempty"`
	RawTree  []byte `protobuf:"bytes,3,opt,name=rawTree,proto3" json:"rawTree,omitempty"`
	RawConf  []byte `protobuf:"bytes,4,opt,name=rawConf,proto3" json:"rawConf,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Event) GetAgentNum() int32 {
	if x != nil {
		return x.AgentNum
	}
	return 0
}

func (x *Event) GetRawTree() []byte {
	if x != nil {
		return x.RawTree
	}
	return nil
}

func (x *Event) GetRawConf() []byte {
	if x != nil {
		return x.RawConf
	}
	return nil
}

type EventReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventID string       `protobuf:"bytes,1,opt,name=eventID,proto3" json:"eventID,omitempty"`
	Status  string       `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Replies []*PlanReply `protobuf:"bytes,3,rep,name=replies,proto3" json:"replies,omitempty"`
}

func (x *EventReply) Reset() {
	*x = EventReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventReply) ProtoMessage() {}

func (x *EventReply) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventReply.ProtoReflect.Descriptor instead.
func (*EventReply) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{3}
}

func (x *EventReply) GetEventID() string {
	if x != nil {
		return x.EventID
	}
	return ""
}

func (x *EventReply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *EventReply) GetReplies() []*PlanReply {
	if x != nil {
		return x.Replies
	}
	return nil
}

type Quantity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string           `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	TotalNum int64            `protobuf:"varint,2,opt,name=totalNum,proto3" json:"totalNum,omitempty"`
	ErrorNum int64            `protobuf:"varint,3,opt,name=errorNum,proto3" json:"errorNum,omitempty"`
	MaxTime  int64            `protobuf:"varint,4,opt,name=maxTime,proto3" json:"maxTime,omitempty"`
	MinTime  int64            `protobuf:"varint,5,opt,name=minTime,proto3" json:"minTime,omitempty"`
	AvgTime  int64            `protobuf:"varint,6,opt,name=avgTime,proto3" json:"avgTime,omitempty"`
	ErrorMap map[string]int64 `protobuf:"bytes,7,rep,name=errorMap,proto3" json:"errorMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *Quantity) Reset() {
	*x = Quantity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quantity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quantity) ProtoMessage() {}

func (x *Quantity) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quantity.ProtoReflect.Descriptor instead.
func (*Quantity) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{4}
}

func (x *Quantity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Quantity) GetTotalNum() int64 {
	if x != nil {
		return x.TotalNum
	}
	return 0
}

func (x *Quantity) GetErrorNum() int64 {
	if x != nil {
		return x.ErrorNum
	}
	return 0
}

func (x *Quantity) GetMaxTime() int64 {
	if x != nil {
		return x.MaxTime
	}
	return 0
}

func (x *Quantity) GetMinTime() int64 {
	if x != nil {
		return x.MinTime
	}
	return 0
}

func (x *Quantity) GetAvgTime() int64 {
	if x != nil {
		return x.AvgTime
	}
	return 0
}

func (x *Quantity) GetErrorMap() map[string]int64 {
	if x != nil {
		return x.ErrorMap
	}
	return nil
}

type PlanReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BeginTime  int64       `protobuf:"varint,2,opt,name=beginTime,proto3" json:"beginTime,omitempty"`
	EndTime    int64       `protobuf:"varint,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Quantities []*Quantity `protobuf:"bytes,4,rep,name=quantities,proto3" json:"quantities,omitempty"`
}

func (x *PlanReply) Reset() {
	*x = PlanReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanReply) ProtoMessage() {}

func (x *PlanReply) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanReply.ProtoReflect.Descriptor instead.
func (*PlanReply) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{5}
}

func (x *PlanReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlanReply) GetBeginTime() int64 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *PlanReply) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *PlanReply) GetQuantities() []*Quantity {
	if x != nil {
		return x.Quantities
	}
	return nil
}

type Outcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status  STATUS `protobuf:"varint,2,opt,name=status,proto3,enum=STATUS" json:"status,omitempty"`
	Consume int64  `protobuf:"varint,3,opt,name=consume,proto3" json:"consume,omitempty"`
	Err     string `protobuf:"bytes,4,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *Outcome) Reset() {
	*x = Outcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outcome) ProtoMessage() {}

func (x *Outcome) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outcome.ProtoReflect.Descriptor instead.
func (*Outcome) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{6}
}

func (x *Outcome) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Outcome) GetStatus() STATUS {
	if x != nil {
		return x.Status
	}
	return STATUS_WAITING
}

func (x *Outcome) GetConsume() int64 {
	if x != nil {
		return x.Consume
	}
	return 0
}

func (x *Outcome) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlanID   string     `protobuf:"bytes,1,opt,name=planID,proto3" json:"planID,omitempty"`
	Outcomes []*Outcome `protobuf:"bytes,2,rep,name=outcomes,proto3" json:"outcomes,omitempty"`
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{7}
}

func (x *Report) GetPlanID() string {
	if x != nil {
		return x.PlanID
	}
	return ""
}

func (x *Report) GetOutcomes() []*Outcome {
	if x != nil {
		return x.Outcomes
	}
	return nil
}

var File_transfer_proto protoreflect.FileDescriptor

var file_transfer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x41, 0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x41, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x76, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x72, 0x65, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x72, 0x65, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x6d, 0x0a, 0x05, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x61, 0x77, 0x54, 0x72, 0x65,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x61, 0x77, 0x54, 0x72, 0x65, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x61, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x72, 0x61, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x22, 0x64, 0x0a, 0x0a, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x07, 0x72, 0x65,
	0x70, 0x6c, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73,
	0x22, 0x96, 0x02, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x0a,
	0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x76, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x61, 0x76, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x61, 0x70, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x61, 0x70, 0x1a, 0x3b, 0x0a, 0x0d,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x82, 0x01, 0x0a, 0x09, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62,
	0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x0a, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x52, 0x0a, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x6a,
	0x0a, 0x07, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x46, 0x0a, 0x06, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x08,
	0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x08, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d,
	0x65, 0x73, 0x2a, 0x4b, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x12, 0x0f, 0x0a, 0x0b,
	0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x02, 0x12, 0x0f,
	0x0a, 0x0b, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x10, 0x03, 0x2a,
	0x47, 0x0a, 0x06, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x49,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x02,
	0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x32, 0x2c, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x4d, 0x61,
	0x69, 0x6c, 0x12, 0x05, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x1a, 0x05, 0x2e, 0x4d, 0x61, 0x69, 0x6c,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transfer_proto_rawDescOnce sync.Once
	file_transfer_proto_rawDescData = file_transfer_proto_rawDesc
)

func file_transfer_proto_rawDescGZIP() []byte {
	file_transfer_proto_rawDescOnce.Do(func() {
		file_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_transfer_proto_rawDescData)
	})
	return file_transfer_proto_rawDescData
}

var file_transfer_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_transfer_proto_goTypes = []interface{}{
	(ACTION)(0),        // 0: ACTION
	(STATUS)(0),        // 1: STATUS
	(*Mail)(nil),       // 2: Mail
	(*Plan)(nil),       // 3: Plan
	(*Event)(nil),      // 4: Event
	(*EventReply)(nil), // 5: EventReply
	(*Quantity)(nil),   // 6: Quantity
	(*PlanReply)(nil),  // 7: PlanReply
	(*Outcome)(nil),    // 8: Outcome
	(*Report)(nil),     // 9: Report
	nil,                // 10: Quantity.ErrorMapEntry
}
var file_transfer_proto_depIdxs = []int32{
	0,  // 0: Mail.action:type_name -> ACTION
	7,  // 1: EventReply.replies:type_name -> PlanReply
	10, // 2: Quantity.errorMap:type_name -> Quantity.ErrorMapEntry
	6,  // 3: PlanReply.quantities:type_name -> Quantity
	1,  // 4: Outcome.status:type_name -> STATUS
	8,  // 5: Report.outcomes:type_name -> Outcome
	2,  // 6: Courier.DeliverMail:input_type -> Mail
	2,  // 7: Courier.DeliverMail:output_type -> Mail
	7,  // [7:8] is the sub-list for method output_type
	6,  // [6:7] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_transfer_proto_init() }
func file_transfer_proto_init() {
	if File_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mail); i {
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
		file_transfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plan); i {
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
		file_transfer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_transfer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventReply); i {
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
		file_transfer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quantity); i {
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
		file_transfer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanReply); i {
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
		file_transfer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outcome); i {
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
		file_transfer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report); i {
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
			RawDescriptor: file_transfer_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transfer_proto_goTypes,
		DependencyIndexes: file_transfer_proto_depIdxs,
		EnumInfos:         file_transfer_proto_enumTypes,
		MessageInfos:      file_transfer_proto_msgTypes,
	}.Build()
	File_transfer_proto = out.File
	file_transfer_proto_rawDesc = nil
	file_transfer_proto_goTypes = nil
	file_transfer_proto_depIdxs = nil
}
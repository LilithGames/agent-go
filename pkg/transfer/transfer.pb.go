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
	ACTION_START_AGENT     ACTION = 0
	ACTION_STOP_AGENT      ACTION = 1
	ACTION_REPORT_DATA     ACTION = 2
	ACTION_FINISH_PLAN     ACTION = 3
	ACTION_START_CIRCLE    ACTION = 4
	ACTION_REPORT_PARALLEL ACTION = 5
)

// Enum value maps for ACTION.
var (
	ACTION_name = map[int32]string{
		0: "START_AGENT",
		1: "STOP_AGENT",
		2: "REPORT_DATA",
		3: "FINISH_PLAN",
		4: "START_CIRCLE",
		5: "REPORT_PARALLEL",
	}
	ACTION_value = map[string]int32{
		"START_AGENT":     0,
		"STOP_AGENT":      1,
		"REPORT_DATA":     2,
		"FINISH_PLAN":     3,
		"START_CIRCLE":    4,
		"REPORT_PARALLEL": 5,
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

type CLASS int32

const (
	CLASS_HANDLER CLASS = 0
	CLASS_EVENT   CLASS = 1
)

// Enum value maps for CLASS.
var (
	CLASS_name = map[int32]string{
		0: "HANDLER",
		1: "EVENT",
	}
	CLASS_value = map[string]int32{
		"HANDLER": 0,
		"EVENT":   1,
	}
)

func (x CLASS) Enum() *CLASS {
	p := new(CLASS)
	*p = x
	return p
}

func (x CLASS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CLASS) Descriptor() protoreflect.EnumDescriptor {
	return file_transfer_proto_enumTypes[2].Descriptor()
}

func (CLASS) Type() protoreflect.EnumType {
	return &file_transfer_proto_enumTypes[2]
}

func (x CLASS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CLASS.Descriptor instead.
func (CLASS) EnumDescriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{2}
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

	Image        string            `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	AgentNum     int32             `protobuf:"varint,2,opt,name=agentNum,proto3" json:"agentNum,omitempty"`
	Parallel     int32             `protobuf:"varint,3,opt,name=parallel,proto3" json:"parallel,omitempty"`
	Environments map[string]string `protobuf:"bytes,4,rep,name=environments,proto3" json:"environments,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
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

func (x *Event) GetParallel() int32 {
	if x != nil {
		return x.Parallel
	}
	return 0
}

func (x *Event) GetEnvironments() map[string]string {
	if x != nil {
		return x.Environments
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

	Name     string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TotalNum int64            `protobuf:"varint,2,opt,name=totalNum,proto3" json:"totalNum,omitempty"`
	ErrorNum int64            `protobuf:"varint,3,opt,name=errorNum,proto3" json:"errorNum,omitempty"`
	MaxTime  int64            `protobuf:"varint,4,opt,name=maxTime,proto3" json:"maxTime,omitempty"`
	MinTime  int64            `protobuf:"varint,5,opt,name=minTime,proto3" json:"minTime,omitempty"`
	AvgTime  int64            `protobuf:"varint,6,opt,name=avgTime,proto3" json:"avgTime,omitempty"`
	Le50Ms   int64            `protobuf:"varint,7,opt,name=le50ms,proto3" json:"le50ms,omitempty"`
	Le100Ms  int64            `protobuf:"varint,8,opt,name=le100ms,proto3" json:"le100ms,omitempty"`
	Le200Ms  int64            `protobuf:"varint,9,opt,name=le200ms,proto3" json:"le200ms,omitempty"`
	Le500Ms  int64            `protobuf:"varint,10,opt,name=le500ms,proto3" json:"le500ms,omitempty"`
	Le1S     int64            `protobuf:"varint,11,opt,name=le1s,proto3" json:"le1s,omitempty"`
	Le2S     int64            `protobuf:"varint,12,opt,name=le2s,proto3" json:"le2s,omitempty"`
	Le5S     int64            `protobuf:"varint,13,opt,name=le5s,proto3" json:"le5s,omitempty"`
	Le10S    int64            `protobuf:"varint,14,opt,name=le10s,proto3" json:"le10s,omitempty"`
	ErrorMap map[string]int64 `protobuf:"bytes,15,rep,name=errorMap,proto3" json:"errorMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Class    CLASS            `protobuf:"varint,16,opt,name=class,proto3,enum=CLASS" json:"class,omitempty"`
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

func (x *Quantity) GetLe50Ms() int64 {
	if x != nil {
		return x.Le50Ms
	}
	return 0
}

func (x *Quantity) GetLe100Ms() int64 {
	if x != nil {
		return x.Le100Ms
	}
	return 0
}

func (x *Quantity) GetLe200Ms() int64 {
	if x != nil {
		return x.Le200Ms
	}
	return 0
}

func (x *Quantity) GetLe500Ms() int64 {
	if x != nil {
		return x.Le500Ms
	}
	return 0
}

func (x *Quantity) GetLe1S() int64 {
	if x != nil {
		return x.Le1S
	}
	return 0
}

func (x *Quantity) GetLe2S() int64 {
	if x != nil {
		return x.Le2S
	}
	return 0
}

func (x *Quantity) GetLe5S() int64 {
	if x != nil {
		return x.Le5S
	}
	return 0
}

func (x *Quantity) GetLe10S() int64 {
	if x != nil {
		return x.Le10S
	}
	return 0
}

func (x *Quantity) GetErrorMap() map[string]int64 {
	if x != nil {
		return x.ErrorMap
	}
	return nil
}

func (x *Quantity) GetClass() CLASS {
	if x != nil {
		return x.Class
	}
	return CLASS_HANDLER
}

type Quantities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Handler map[string]*Quantity `protobuf:"bytes,1,rep,name=handler,proto3" json:"handler,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Event   map[string]*Quantity `protobuf:"bytes,2,rep,name=event,proto3" json:"event,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Quantities) Reset() {
	*x = Quantities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quantities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quantities) ProtoMessage() {}

func (x *Quantities) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Quantities.ProtoReflect.Descriptor instead.
func (*Quantities) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{5}
}

func (x *Quantities) GetHandler() map[string]*Quantity {
	if x != nil {
		return x.Handler
	}
	return nil
}

func (x *Quantities) GetEvent() map[string]*Quantity {
	if x != nil {
		return x.Event
	}
	return nil
}

type PlanReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BeginTime   int64       `protobuf:"varint,2,opt,name=beginTime,proto3" json:"beginTime,omitempty"`
	EndTime     int64       `protobuf:"varint,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
	HQuantities []*Quantity `protobuf:"bytes,4,rep,name=hQuantities,proto3" json:"hQuantities,omitempty"`
	EQuantities []*Quantity `protobuf:"bytes,5,rep,name=eQuantities,proto3" json:"eQuantities,omitempty"`
}

func (x *PlanReply) Reset() {
	*x = PlanReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanReply) ProtoMessage() {}

func (x *PlanReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PlanReply.ProtoReflect.Descriptor instead.
func (*PlanReply) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{6}
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

func (x *PlanReply) GetHQuantities() []*Quantity {
	if x != nil {
		return x.HQuantities
	}
	return nil
}

func (x *PlanReply) GetEQuantities() []*Quantity {
	if x != nil {
		return x.EQuantities
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
	Class   CLASS  `protobuf:"varint,5,opt,name=class,proto3,enum=CLASS" json:"class,omitempty"`
}

func (x *Outcome) Reset() {
	*x = Outcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outcome) ProtoMessage() {}

func (x *Outcome) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Outcome.ProtoReflect.Descriptor instead.
func (*Outcome) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{7}
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

func (x *Outcome) GetClass() CLASS {
	if x != nil {
		return x.Class
	}
	return CLASS_HANDLER
}

type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlanID   string     `protobuf:"bytes,1,opt,name=planID,proto3" json:"planID,omitempty"`
	EventNum int64      `protobuf:"varint,2,opt,name=eventNum,proto3" json:"eventNum,omitempty"`
	Outcomes []*Outcome `protobuf:"bytes,3,rep,name=outcomes,proto3" json:"outcomes,omitempty"`
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[8]
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
	return file_transfer_proto_rawDescGZIP(), []int{8}
}

func (x *Report) GetPlanID() string {
	if x != nil {
		return x.PlanID
	}
	return ""
}

func (x *Report) GetEventNum() int64 {
	if x != nil {
		return x.EventNum
	}
	return 0
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
	0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0xd4, 0x01, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x61, 0x6c,
	0x6c, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x72, 0x61, 0x6c,
	0x6c, 0x65, 0x6c, 0x12, 0x3c, 0x0a, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x1a, 0x3f, 0x0a, 0x11, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x64, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x24, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52,
	0x07, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x22, 0xec, 0x03, 0x0a, 0x08, 0x51, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4e, 0x75,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4e, 0x75,
	0x6d, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x69,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x76, 0x67, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x76, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6c, 0x65, 0x35, 0x30, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6c, 0x65, 0x35, 0x30, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x65, 0x31, 0x30, 0x30,
	0x6d, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x65, 0x31, 0x30, 0x30, 0x6d,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x65, 0x32, 0x30, 0x30, 0x6d, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x6c, 0x65, 0x32, 0x30, 0x30, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6c,
	0x65, 0x35, 0x30, 0x30, 0x6d, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x65,
	0x35, 0x30, 0x30, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x65, 0x31, 0x73, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x6c, 0x65, 0x31, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x65, 0x32,
	0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6c, 0x65, 0x32, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x65, 0x35, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6c, 0x65, 0x35,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x31, 0x30, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6c, 0x65, 0x31, 0x30, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x61, 0x70, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x61, 0x70, 0x12, 0x1c, 0x0a, 0x05,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x06, 0x2e, 0x43, 0x4c,
	0x41, 0x53, 0x53, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xfa, 0x01, 0x0a, 0x0a, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x45, 0x0a, 0x0c, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x43, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x1f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xb1, 0x01, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2b,
	0x0a, 0x0b, 0x68, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0b,
	0x68, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x0b, 0x65,
	0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0b, 0x65, 0x51, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x88, 0x01, 0x0a, 0x07, 0x4f, 0x75, 0x74,
	0x63, 0x6f, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x1c, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x06, 0x2e, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x52, 0x05, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x22, 0x62, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x6c, 0x61, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x75,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x75,
	0x6d, 0x12, 0x24, 0x0a, 0x08, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x08, 0x6f,
	0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x73, 0x2a, 0x72, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54,
	0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54,
	0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x44, 0x41, 0x54,
	0x41, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x5f, 0x50, 0x4c,
	0x41, 0x4e, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x43, 0x49,
	0x52, 0x43, 0x4c, 0x45, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54,
	0x5f, 0x50, 0x41, 0x52, 0x41, 0x4c, 0x4c, 0x45, 0x4c, 0x10, 0x05, 0x2a, 0x47, 0x0a, 0x06, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07,
	0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x04, 0x2a, 0x1f, 0x0a, 0x05, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x12, 0x0b, 0x0a,
	0x07, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x45, 0x52, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x10, 0x01, 0x32, 0x2c, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x12, 0x21, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x4d, 0x61, 0x69, 0x6c, 0x12,
	0x05, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x1a, 0x05, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_transfer_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_transfer_proto_goTypes = []interface{}{
	(ACTION)(0),        // 0: ACTION
	(STATUS)(0),        // 1: STATUS
	(CLASS)(0),         // 2: CLASS
	(*Mail)(nil),       // 3: Mail
	(*Plan)(nil),       // 4: Plan
	(*Event)(nil),      // 5: Event
	(*EventReply)(nil), // 6: EventReply
	(*Quantity)(nil),   // 7: Quantity
	(*Quantities)(nil), // 8: Quantities
	(*PlanReply)(nil),  // 9: PlanReply
	(*Outcome)(nil),    // 10: Outcome
	(*Report)(nil),     // 11: Report
	nil,                // 12: Event.EnvironmentsEntry
	nil,                // 13: Quantity.ErrorMapEntry
	nil,                // 14: Quantities.HandlerEntry
	nil,                // 15: Quantities.EventEntry
}
var file_transfer_proto_depIdxs = []int32{
	0,  // 0: Mail.action:type_name -> ACTION
	12, // 1: Event.environments:type_name -> Event.EnvironmentsEntry
	9,  // 2: EventReply.replies:type_name -> PlanReply
	13, // 3: Quantity.errorMap:type_name -> Quantity.ErrorMapEntry
	2,  // 4: Quantity.class:type_name -> CLASS
	14, // 5: Quantities.handler:type_name -> Quantities.HandlerEntry
	15, // 6: Quantities.event:type_name -> Quantities.EventEntry
	7,  // 7: PlanReply.hQuantities:type_name -> Quantity
	7,  // 8: PlanReply.eQuantities:type_name -> Quantity
	1,  // 9: Outcome.status:type_name -> STATUS
	2,  // 10: Outcome.class:type_name -> CLASS
	10, // 11: Report.outcomes:type_name -> Outcome
	7,  // 12: Quantities.HandlerEntry.value:type_name -> Quantity
	7,  // 13: Quantities.EventEntry.value:type_name -> Quantity
	3,  // 14: Courier.DeliverMail:input_type -> Mail
	3,  // 15: Courier.DeliverMail:output_type -> Mail
	15, // [15:16] is the sub-list for method output_type
	14, // [14:15] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
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
			switch v := v.(*Quantities); i {
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
		file_transfer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_transfer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
			NumEnums:      3,
			NumMessages:   13,
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

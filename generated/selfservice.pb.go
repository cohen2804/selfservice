//generate protobuf:
// protoc --go_out=./generated --go_opt=paths=source_relative \
// --go-grpc_out=./generated --go-grpc_opt=paths=source_relative \
// grpcmain.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: selfservice.proto

package generated

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

type Status int32

const (
	Status_IN_PROCESS Status = 0
	Status_FINISHED   Status = 1
	Status_FAILED     Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "IN_PROCESS",
		1: "FINISHED",
		2: "FAILED",
	}
	Status_value = map[string]int32{
		"IN_PROCESS": 0,
		"FINISHED":   1,
		"FAILED":     2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_selfservice_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_selfservice_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_selfservice_proto_rawDescGZIP(), []int{0}
}

type ActionType int32

const (
	ActionType_EXPORT ActionType = 0
	ActionType_REMOVE ActionType = 1
)

// Enum value maps for ActionType.
var (
	ActionType_name = map[int32]string{
		0: "EXPORT",
		1: "REMOVE",
	}
	ActionType_value = map[string]int32{
		"EXPORT": 0,
		"REMOVE": 1,
	}
)

func (x ActionType) Enum() *ActionType {
	p := new(ActionType)
	*p = x
	return p
}

func (x ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_selfservice_proto_enumTypes[1].Descriptor()
}

func (ActionType) Type() protoreflect.EnumType {
	return &file_selfservice_proto_enumTypes[1]
}

func (x ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionType.Descriptor instead.
func (ActionType) EnumDescriptor() ([]byte, []int) {
	return file_selfservice_proto_rawDescGZIP(), []int{1}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,4,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Subject  string `protobuf:"bytes,5,opt,name=Subject,proto3" json:"Subject,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_selfservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_selfservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_selfservice_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *User) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

type ActionRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantID     string     `protobuf:"bytes,1,opt,name=TenantID,proto3" json:"TenantID,omitempty"`
	SessionID    string     `protobuf:"bytes,2,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	Action       ActionType `protobuf:"varint,3,opt,name=Action,proto3,enum=selfservice.ActionType" json:"Action,omitempty"`
	Status       Status     `protobuf:"varint,4,opt,name=Status,proto3,enum=selfservice.Status" json:"Status,omitempty"`
	User         *User      `protobuf:"bytes,5,opt,name=User,proto3" json:"User,omitempty"`
	StartedDate  string     `protobuf:"bytes,6,opt,name=StartedDate,proto3" json:"StartedDate,omitempty"`
	FinishedDate string     `protobuf:"bytes,7,opt,name=FinishedDate,proto3" json:"FinishedDate,omitempty"`
}

func (x *ActionRecord) Reset() {
	*x = ActionRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_selfservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionRecord) ProtoMessage() {}

func (x *ActionRecord) ProtoReflect() protoreflect.Message {
	mi := &file_selfservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionRecord.ProtoReflect.Descriptor instead.
func (*ActionRecord) Descriptor() ([]byte, []int) {
	return file_selfservice_proto_rawDescGZIP(), []int{1}
}

func (x *ActionRecord) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *ActionRecord) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *ActionRecord) GetAction() ActionType {
	if x != nil {
		return x.Action
	}
	return ActionType_EXPORT
}

func (x *ActionRecord) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_IN_PROCESS
}

func (x *ActionRecord) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ActionRecord) GetStartedDate() string {
	if x != nil {
		return x.StartedDate
	}
	return ""
}

func (x *ActionRecord) GetFinishedDate() string {
	if x != nil {
		return x.FinishedDate
	}
	return ""
}

type FilterOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"` // Field to filter by (e.g., "action")
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"` // Value to filter for (e.g., "export")
}

func (x *FilterOption) Reset() {
	*x = FilterOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_selfservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterOption) ProtoMessage() {}

func (x *FilterOption) ProtoReflect() protoreflect.Message {
	mi := &file_selfservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterOption.ProtoReflect.Descriptor instead.
func (*FilterOption) Descriptor() ([]byte, []int) {
	return file_selfservice_proto_rawDescGZIP(), []int{2}
}

func (x *FilterOption) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *FilterOption) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageToken int32         `protobuf:"varint,1,opt,name=PageToken,proto3" json:"PageToken,omitempty"` // Token to retrieve the next page (optional)
	PageSize  int32         `protobuf:"varint,2,opt,name=PageSize,proto3" json:"PageSize,omitempty"`   // Number of results per page (optional)
	Filter    *FilterOption `protobuf:"bytes,3,opt,name=Filter,proto3" json:"Filter,omitempty"`        // Filter by specific field and value (optional)
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_selfservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_selfservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_selfservice_proto_rawDescGZIP(), []int{3}
}

func (x *Query) GetPageToken() int32 {
	if x != nil {
		return x.PageToken
	}
	return 0
}

func (x *Query) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *Query) GetFilter() *FilterOption {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListActionsRq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantID string `protobuf:"bytes,1,opt,name=TenantID,proto3" json:"TenantID,omitempty"`
	Query    *Query `protobuf:"bytes,2,opt,name=Query,proto3" json:"Query,omitempty"`
}

func (x *ListActionsRq) Reset() {
	*x = ListActionsRq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_selfservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListActionsRq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListActionsRq) ProtoMessage() {}

func (x *ListActionsRq) ProtoReflect() protoreflect.Message {
	mi := &file_selfservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListActionsRq.ProtoReflect.Descriptor instead.
func (*ListActionsRq) Descriptor() ([]byte, []int) {
	return file_selfservice_proto_rawDescGZIP(), []int{4}
}

func (x *ListActionsRq) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *ListActionsRq) GetQuery() *Query {
	if x != nil {
		return x.Query
	}
	return nil
}

type ListActionsRs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*ActionRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *ListActionsRs) Reset() {
	*x = ListActionsRs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_selfservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListActionsRs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListActionsRs) ProtoMessage() {}

func (x *ListActionsRs) ProtoReflect() protoreflect.Message {
	mi := &file_selfservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListActionsRs.ProtoReflect.Descriptor instead.
func (*ListActionsRs) Descriptor() ([]byte, []int) {
	return file_selfservice_proto_rawDescGZIP(), []int{5}
}

func (x *ListActionsRs) GetRecords() []*ActionRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

type LogItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp string `protobuf:"bytes,1,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *LogItem) Reset() {
	*x = LogItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_selfservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogItem) ProtoMessage() {}

func (x *LogItem) ProtoReflect() protoreflect.Message {
	mi := &file_selfservice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogItem.ProtoReflect.Descriptor instead.
func (*LogItem) Descriptor() ([]byte, []int) {
	return file_selfservice_proto_rawDescGZIP(), []int{6}
}

func (x *LogItem) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *LogItem) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetActionLogRq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantID  string `protobuf:"bytes,1,opt,name=TenantID,proto3" json:"TenantID,omitempty"`
	SessionID string `protobuf:"bytes,2,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
}

func (x *GetActionLogRq) Reset() {
	*x = GetActionLogRq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_selfservice_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActionLogRq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActionLogRq) ProtoMessage() {}

func (x *GetActionLogRq) ProtoReflect() protoreflect.Message {
	mi := &file_selfservice_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActionLogRq.ProtoReflect.Descriptor instead.
func (*GetActionLogRq) Descriptor() ([]byte, []int) {
	return file_selfservice_proto_rawDescGZIP(), []int{7}
}

func (x *GetActionLogRq) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *GetActionLogRq) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

type GetActionLogRs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogsList []*LogItem `protobuf:"bytes,1,rep,name=LogsList,proto3" json:"LogsList,omitempty"`
}

func (x *GetActionLogRs) Reset() {
	*x = GetActionLogRs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_selfservice_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActionLogRs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActionLogRs) ProtoMessage() {}

func (x *GetActionLogRs) ProtoReflect() protoreflect.Message {
	mi := &file_selfservice_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActionLogRs.ProtoReflect.Descriptor instead.
func (*GetActionLogRs) Descriptor() ([]byte, []int) {
	return file_selfservice_proto_rawDescGZIP(), []int{8}
}

func (x *GetActionLogRs) GetLogsList() []*LogItem {
	if x != nil {
		return x.LogsList
	}
	return nil
}

var File_selfservice_proto protoreflect.FileDescriptor

var file_selfservice_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x6c, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x65, 0x6c, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x3c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x93,
	0x02, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x2f, 0x0a, 0x06, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x6c, 0x66,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x6c,
	0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x6c, 0x66, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20,
	0x0a, 0x0b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x22, 0x3a, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x74, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x6c, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x55, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x6c, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x22, 0x44, 0x0a,
	0x0d, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x73, 0x12, 0x33,
	0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x73, 0x65, 0x6c, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x22, 0x41, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c,
	0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x52, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x22, 0x42, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x6f, 0x67, 0x52, 0x73, 0x12, 0x30, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x6c, 0x66, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x4c, 0x6f,
	0x67, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x2a, 0x32, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x24, 0x0a, 0x0a, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x58, 0x50, 0x4f,
	0x52, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x01,
	0x32, 0xa8, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x6c, 0x66,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x71, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x6c, 0x66, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x73, 0x12, 0x48, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f,
	0x67, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x6c, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x52, 0x71, 0x1a, 0x1b,
	0x2e, 0x73, 0x65, 0x6c, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x52, 0x73, 0x42, 0x0d, 0x5a, 0x0b, 0x2e,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_selfservice_proto_rawDescOnce sync.Once
	file_selfservice_proto_rawDescData = file_selfservice_proto_rawDesc
)

func file_selfservice_proto_rawDescGZIP() []byte {
	file_selfservice_proto_rawDescOnce.Do(func() {
		file_selfservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_selfservice_proto_rawDescData)
	})
	return file_selfservice_proto_rawDescData
}

var file_selfservice_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_selfservice_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_selfservice_proto_goTypes = []interface{}{
	(Status)(0),            // 0: selfservice.Status
	(ActionType)(0),        // 1: selfservice.ActionType
	(*User)(nil),           // 2: selfservice.User
	(*ActionRecord)(nil),   // 3: selfservice.ActionRecord
	(*FilterOption)(nil),   // 4: selfservice.FilterOption
	(*Query)(nil),          // 5: selfservice.Query
	(*ListActionsRq)(nil),  // 6: selfservice.ListActionsRq
	(*ListActionsRs)(nil),  // 7: selfservice.ListActionsRs
	(*LogItem)(nil),        // 8: selfservice.LogItem
	(*GetActionLogRq)(nil), // 9: selfservice.GetActionLogRq
	(*GetActionLogRs)(nil), // 10: selfservice.GetActionLogRs
}
var file_selfservice_proto_depIdxs = []int32{
	1,  // 0: selfservice.ActionRecord.Action:type_name -> selfservice.ActionType
	0,  // 1: selfservice.ActionRecord.Status:type_name -> selfservice.Status
	2,  // 2: selfservice.ActionRecord.User:type_name -> selfservice.User
	4,  // 3: selfservice.Query.Filter:type_name -> selfservice.FilterOption
	5,  // 4: selfservice.ListActionsRq.Query:type_name -> selfservice.Query
	3,  // 5: selfservice.ListActionsRs.records:type_name -> selfservice.ActionRecord
	8,  // 6: selfservice.GetActionLogRs.LogsList:type_name -> selfservice.LogItem
	6,  // 7: selfservice.SessionRecordsService.ListActions:input_type -> selfservice.ListActionsRq
	9,  // 8: selfservice.SessionRecordsService.GetActionLog:input_type -> selfservice.GetActionLogRq
	7,  // 9: selfservice.SessionRecordsService.ListActions:output_type -> selfservice.ListActionsRs
	10, // 10: selfservice.SessionRecordsService.GetActionLog:output_type -> selfservice.GetActionLogRs
	9,  // [9:11] is the sub-list for method output_type
	7,  // [7:9] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_selfservice_proto_init() }
func file_selfservice_proto_init() {
	if File_selfservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_selfservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_selfservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionRecord); i {
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
		file_selfservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterOption); i {
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
		file_selfservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
		file_selfservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListActionsRq); i {
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
		file_selfservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListActionsRs); i {
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
		file_selfservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogItem); i {
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
		file_selfservice_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetActionLogRq); i {
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
		file_selfservice_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetActionLogRs); i {
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
			RawDescriptor: file_selfservice_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_selfservice_proto_goTypes,
		DependencyIndexes: file_selfservice_proto_depIdxs,
		EnumInfos:         file_selfservice_proto_enumTypes,
		MessageInfos:      file_selfservice_proto_msgTypes,
	}.Build()
	File_selfservice_proto = out.File
	file_selfservice_proto_rawDesc = nil
	file_selfservice_proto_goTypes = nil
	file_selfservice_proto_depIdxs = nil
}
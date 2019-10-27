// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/server/server.proto

package server

import (
	context "context"
	fmt "fmt"
	event "github.com/egor1344/otus_calendar/calendar/proto/event"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetEventListRequest_Type int32

const (
	GetEventListRequest_week  GetEventListRequest_Type = 0
	GetEventListRequest_month GetEventListRequest_Type = 1
	GetEventListRequest_year  GetEventListRequest_Type = 2
)

var GetEventListRequest_Type_name = map[int32]string{
	0: "week",
	1: "month",
	2: "year",
}

var GetEventListRequest_Type_value = map[string]int32{
	"week":  0,
	"month": 1,
	"year":  2,
}

func (x GetEventListRequest_Type) String() string {
	return proto.EnumName(GetEventListRequest_Type_name, int32(x))
}

func (GetEventListRequest_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aedf378765df4bae, []int{8, 0}
}

type AddEventRequest struct {
	Event                *event.Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AddEventRequest) Reset()         { *m = AddEventRequest{} }
func (m *AddEventRequest) String() string { return proto.CompactTextString(m) }
func (*AddEventRequest) ProtoMessage()    {}
func (*AddEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aedf378765df4bae, []int{0}
}

func (m *AddEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddEventRequest.Unmarshal(m, b)
}
func (m *AddEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddEventRequest.Marshal(b, m, deterministic)
}
func (m *AddEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddEventRequest.Merge(m, src)
}
func (m *AddEventRequest) XXX_Size() int {
	return xxx_messageInfo_AddEventRequest.Size(m)
}
func (m *AddEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddEventRequest proto.InternalMessageInfo

func (m *AddEventRequest) GetEvent() *event.Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type AddEventResponse struct {
	// Types that are valid to be assigned to Result:
	//	*AddEventResponse_Event
	//	*AddEventResponse_Error
	Result               isAddEventResponse_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *AddEventResponse) Reset()         { *m = AddEventResponse{} }
func (m *AddEventResponse) String() string { return proto.CompactTextString(m) }
func (*AddEventResponse) ProtoMessage()    {}
func (*AddEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aedf378765df4bae, []int{1}
}

func (m *AddEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddEventResponse.Unmarshal(m, b)
}
func (m *AddEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddEventResponse.Marshal(b, m, deterministic)
}
func (m *AddEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddEventResponse.Merge(m, src)
}
func (m *AddEventResponse) XXX_Size() int {
	return xxx_messageInfo_AddEventResponse.Size(m)
}
func (m *AddEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddEventResponse proto.InternalMessageInfo

type isAddEventResponse_Result interface {
	isAddEventResponse_Result()
}

type AddEventResponse_Event struct {
	Event *event.Event `protobuf:"bytes,1,opt,name=event,proto3,oneof"`
}

type AddEventResponse_Error struct {
	Error string `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*AddEventResponse_Event) isAddEventResponse_Result() {}

func (*AddEventResponse_Error) isAddEventResponse_Result() {}

func (m *AddEventResponse) GetResult() isAddEventResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *AddEventResponse) GetEvent() *event.Event {
	if x, ok := m.GetResult().(*AddEventResponse_Event); ok {
		return x.Event
	}
	return nil
}

func (m *AddEventResponse) GetError() string {
	if x, ok := m.GetResult().(*AddEventResponse_Error); ok {
		return x.Error
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AddEventResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AddEventResponse_Event)(nil),
		(*AddEventResponse_Error)(nil),
	}
}

type GetEventRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventRequest) Reset()         { *m = GetEventRequest{} }
func (m *GetEventRequest) String() string { return proto.CompactTextString(m) }
func (*GetEventRequest) ProtoMessage()    {}
func (*GetEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aedf378765df4bae, []int{2}
}

func (m *GetEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventRequest.Unmarshal(m, b)
}
func (m *GetEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventRequest.Marshal(b, m, deterministic)
}
func (m *GetEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventRequest.Merge(m, src)
}
func (m *GetEventRequest) XXX_Size() int {
	return xxx_messageInfo_GetEventRequest.Size(m)
}
func (m *GetEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventRequest proto.InternalMessageInfo

func (m *GetEventRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetEventResponse struct {
	// Types that are valid to be assigned to Result:
	//	*GetEventResponse_Event
	//	*GetEventResponse_Error
	Result               isGetEventResponse_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GetEventResponse) Reset()         { *m = GetEventResponse{} }
func (m *GetEventResponse) String() string { return proto.CompactTextString(m) }
func (*GetEventResponse) ProtoMessage()    {}
func (*GetEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aedf378765df4bae, []int{3}
}

func (m *GetEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventResponse.Unmarshal(m, b)
}
func (m *GetEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventResponse.Marshal(b, m, deterministic)
}
func (m *GetEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventResponse.Merge(m, src)
}
func (m *GetEventResponse) XXX_Size() int {
	return xxx_messageInfo_GetEventResponse.Size(m)
}
func (m *GetEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventResponse proto.InternalMessageInfo

type isGetEventResponse_Result interface {
	isGetEventResponse_Result()
}

type GetEventResponse_Event struct {
	Event *event.Event `protobuf:"bytes,1,opt,name=event,proto3,oneof"`
}

type GetEventResponse_Error struct {
	Error string `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*GetEventResponse_Event) isGetEventResponse_Result() {}

func (*GetEventResponse_Error) isGetEventResponse_Result() {}

func (m *GetEventResponse) GetResult() isGetEventResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GetEventResponse) GetEvent() *event.Event {
	if x, ok := m.GetResult().(*GetEventResponse_Event); ok {
		return x.Event
	}
	return nil
}

func (m *GetEventResponse) GetError() string {
	if x, ok := m.GetResult().(*GetEventResponse_Error); ok {
		return x.Error
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GetEventResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GetEventResponse_Event)(nil),
		(*GetEventResponse_Error)(nil),
	}
}

type UpdateEventRequest struct {
	Event                *event.Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateEventRequest) Reset()         { *m = UpdateEventRequest{} }
func (m *UpdateEventRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateEventRequest) ProtoMessage()    {}
func (*UpdateEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aedf378765df4bae, []int{4}
}

func (m *UpdateEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEventRequest.Unmarshal(m, b)
}
func (m *UpdateEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEventRequest.Marshal(b, m, deterministic)
}
func (m *UpdateEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEventRequest.Merge(m, src)
}
func (m *UpdateEventRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateEventRequest.Size(m)
}
func (m *UpdateEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEventRequest proto.InternalMessageInfo

func (m *UpdateEventRequest) GetEvent() *event.Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type UpdateEventResponse struct {
	// Types that are valid to be assigned to Result:
	//	*UpdateEventResponse_Event
	//	*UpdateEventResponse_Error
	Result               isUpdateEventResponse_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *UpdateEventResponse) Reset()         { *m = UpdateEventResponse{} }
func (m *UpdateEventResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateEventResponse) ProtoMessage()    {}
func (*UpdateEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aedf378765df4bae, []int{5}
}

func (m *UpdateEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEventResponse.Unmarshal(m, b)
}
func (m *UpdateEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEventResponse.Marshal(b, m, deterministic)
}
func (m *UpdateEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEventResponse.Merge(m, src)
}
func (m *UpdateEventResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateEventResponse.Size(m)
}
func (m *UpdateEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEventResponse proto.InternalMessageInfo

type isUpdateEventResponse_Result interface {
	isUpdateEventResponse_Result()
}

type UpdateEventResponse_Event struct {
	Event *event.Event `protobuf:"bytes,1,opt,name=event,proto3,oneof"`
}

type UpdateEventResponse_Error struct {
	Error string `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*UpdateEventResponse_Event) isUpdateEventResponse_Result() {}

func (*UpdateEventResponse_Error) isUpdateEventResponse_Result() {}

func (m *UpdateEventResponse) GetResult() isUpdateEventResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *UpdateEventResponse) GetEvent() *event.Event {
	if x, ok := m.GetResult().(*UpdateEventResponse_Event); ok {
		return x.Event
	}
	return nil
}

func (m *UpdateEventResponse) GetError() string {
	if x, ok := m.GetResult().(*UpdateEventResponse_Error); ok {
		return x.Error
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UpdateEventResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UpdateEventResponse_Event)(nil),
		(*UpdateEventResponse_Error)(nil),
	}
}

type DeleteEventRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEventRequest) Reset()         { *m = DeleteEventRequest{} }
func (m *DeleteEventRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteEventRequest) ProtoMessage()    {}
func (*DeleteEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aedf378765df4bae, []int{6}
}

func (m *DeleteEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEventRequest.Unmarshal(m, b)
}
func (m *DeleteEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEventRequest.Marshal(b, m, deterministic)
}
func (m *DeleteEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEventRequest.Merge(m, src)
}
func (m *DeleteEventRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteEventRequest.Size(m)
}
func (m *DeleteEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEventRequest proto.InternalMessageInfo

func (m *DeleteEventRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteEventResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEventResponse) Reset()         { *m = DeleteEventResponse{} }
func (m *DeleteEventResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteEventResponse) ProtoMessage()    {}
func (*DeleteEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aedf378765df4bae, []int{7}
}

func (m *DeleteEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEventResponse.Unmarshal(m, b)
}
func (m *DeleteEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEventResponse.Marshal(b, m, deterministic)
}
func (m *DeleteEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEventResponse.Merge(m, src)
}
func (m *DeleteEventResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteEventResponse.Size(m)
}
func (m *DeleteEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEventResponse proto.InternalMessageInfo

func (m *DeleteEventResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetEventListRequest struct {
	Type                 GetEventListRequest_Type `protobuf:"varint,1,opt,name=type,proto3,enum=calendar.GetEventListRequest_Type" json:"type,omitempty"`
	UserId               string                   `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *GetEventListRequest) Reset()         { *m = GetEventListRequest{} }
func (m *GetEventListRequest) String() string { return proto.CompactTextString(m) }
func (*GetEventListRequest) ProtoMessage()    {}
func (*GetEventListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aedf378765df4bae, []int{8}
}

func (m *GetEventListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventListRequest.Unmarshal(m, b)
}
func (m *GetEventListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventListRequest.Marshal(b, m, deterministic)
}
func (m *GetEventListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventListRequest.Merge(m, src)
}
func (m *GetEventListRequest) XXX_Size() int {
	return xxx_messageInfo_GetEventListRequest.Size(m)
}
func (m *GetEventListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventListRequest proto.InternalMessageInfo

func (m *GetEventListRequest) GetType() GetEventListRequest_Type {
	if m != nil {
		return m.Type
	}
	return GetEventListRequest_week
}

func (m *GetEventListRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetEventListResponse struct {
	Event                []*event.Event `protobuf:"bytes,1,rep,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetEventListResponse) Reset()         { *m = GetEventListResponse{} }
func (m *GetEventListResponse) String() string { return proto.CompactTextString(m) }
func (*GetEventListResponse) ProtoMessage()    {}
func (*GetEventListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aedf378765df4bae, []int{9}
}

func (m *GetEventListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventListResponse.Unmarshal(m, b)
}
func (m *GetEventListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventListResponse.Marshal(b, m, deterministic)
}
func (m *GetEventListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventListResponse.Merge(m, src)
}
func (m *GetEventListResponse) XXX_Size() int {
	return xxx_messageInfo_GetEventListResponse.Size(m)
}
func (m *GetEventListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventListResponse proto.InternalMessageInfo

func (m *GetEventListResponse) GetEvent() []*event.Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func init() {
	proto.RegisterEnum("calendar.GetEventListRequest_Type", GetEventListRequest_Type_name, GetEventListRequest_Type_value)
	proto.RegisterType((*AddEventRequest)(nil), "calendar.AddEventRequest")
	proto.RegisterType((*AddEventResponse)(nil), "calendar.AddEventResponse")
	proto.RegisterType((*GetEventRequest)(nil), "calendar.GetEventRequest")
	proto.RegisterType((*GetEventResponse)(nil), "calendar.GetEventResponse")
	proto.RegisterType((*UpdateEventRequest)(nil), "calendar.UpdateEventRequest")
	proto.RegisterType((*UpdateEventResponse)(nil), "calendar.UpdateEventResponse")
	proto.RegisterType((*DeleteEventRequest)(nil), "calendar.DeleteEventRequest")
	proto.RegisterType((*DeleteEventResponse)(nil), "calendar.DeleteEventResponse")
	proto.RegisterType((*GetEventListRequest)(nil), "calendar.GetEventListRequest")
	proto.RegisterType((*GetEventListResponse)(nil), "calendar.GetEventListResponse")
}

func init() { proto.RegisterFile("proto/server/server.proto", fileDescriptor_aedf378765df4bae) }

var fileDescriptor_aedf378765df4bae = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcf, 0xcb, 0xd3, 0x40,
	0x10, 0xfd, 0x12, 0xdb, 0x90, 0xce, 0x57, 0xdb, 0xb0, 0x95, 0xda, 0x06, 0x2b, 0x75, 0xa9, 0xd0,
	0x8b, 0x11, 0x2a, 0x8a, 0x78, 0x6b, 0x55, 0xfc, 0x81, 0x5e, 0x82, 0x5e, 0x0a, 0x1e, 0xa2, 0x19,
	0xb0, 0x58, 0x93, 0xb8, 0xbb, 0xa9, 0xf4, 0x7f, 0xf0, 0x4f, 0xf6, 0x20, 0xd9, 0xdd, 0x98, 0x4d,
	0xfa, 0x03, 0x3e, 0xe8, 0x25, 0x61, 0x66, 0xde, 0xbc, 0x79, 0x3b, 0xfb, 0x58, 0x18, 0x67, 0x2c,
	0x15, 0xe9, 0x63, 0x8e, 0x6c, 0x87, 0x4c, 0xff, 0x02, 0x99, 0x23, 0xee, 0xb7, 0x68, 0x8b, 0x49,
	0x1c, 0x31, 0xff, 0xae, 0x02, 0xe1, 0x0e, 0x13, 0xa1, 0xbe, 0x0a, 0x42, 0x9f, 0x42, 0x7f, 0x19,
	0xc7, 0xaf, 0x8b, 0x4c, 0x88, 0xbf, 0x72, 0xe4, 0x82, 0x50, 0x68, 0x4b, 0xc4, 0xc8, 0x9a, 0x5a,
	0xf3, 0xeb, 0x45, 0x37, 0x50, 0x78, 0x85, 0x51, 0x25, 0xba, 0x06, 0xaf, 0x6a, 0xe3, 0x59, 0x9a,
	0x70, 0x24, 0xb3, 0x33, 0x7d, 0x6f, 0xaf, 0x74, 0x27, 0x19, 0x42, 0x1b, 0x19, 0x4b, 0xd9, 0xc8,
	0x9e, 0x5a, 0xf3, 0x8e, 0xcc, 0x17, 0xe1, 0xca, 0x05, 0x87, 0x21, 0xcf, 0xb7, 0x82, 0x3e, 0x80,
	0xfe, 0x1b, 0x14, 0x35, 0x49, 0x3d, 0xb0, 0x37, 0xb1, 0xe4, 0xed, 0x84, 0xf6, 0x26, 0x2e, 0xc6,
	0x57, 0x90, 0x0b, 0x8f, 0x7f, 0x0e, 0xe4, 0x73, 0x16, 0x47, 0x02, 0x6f, 0xbc, 0x94, 0x2f, 0x30,
	0xa8, 0x75, 0x5e, 0x58, 0xd8, 0x0c, 0xc8, 0x2b, 0xdc, 0x62, 0x43, 0x58, 0x73, 0x35, 0x8f, 0x60,
	0x50, 0x43, 0x69, 0x11, 0x43, 0x70, 0xb8, 0x88, 0x44, 0xce, 0x35, 0x54, 0x47, 0xf4, 0x8f, 0x05,
	0x83, 0x72, 0x95, 0x1f, 0x36, 0xfc, 0x3f, 0xed, 0x33, 0x68, 0x89, 0x7d, 0x86, 0x12, 0xdd, 0x5b,
	0xd0, 0xa0, 0x74, 0x52, 0x70, 0x04, 0x1c, 0x7c, 0xda, 0x67, 0x18, 0x4a, 0x7c, 0x31, 0x27, 0xe7,
	0xc8, 0xde, 0xc5, 0xea, 0x1c, 0xa1, 0x8e, 0xe8, 0x43, 0x68, 0x15, 0x28, 0xe2, 0x42, 0xeb, 0x37,
	0xe2, 0x0f, 0xef, 0x8a, 0x74, 0xa0, 0xfd, 0x33, 0x4d, 0xc4, 0x77, 0xcf, 0x2a, 0x92, 0x7b, 0x8c,
	0x98, 0x67, 0xd3, 0x17, 0x70, 0xa7, 0x3e, 0x40, 0xcb, 0x37, 0xd6, 0x7f, 0xeb, 0xc4, 0xfa, 0x17,
	0x7f, 0x6d, 0xb8, 0xfd, 0x52, 0xcb, 0x94, 0x05, 0xb2, 0x04, 0xb7, 0x74, 0x29, 0x19, 0x57, 0x47,
	0x68, 0x18, 0xde, 0xf7, 0x8f, 0x95, 0xf4, 0xe0, 0x25, 0xb8, 0xa5, 0x20, 0x93, 0xa2, 0x61, 0x50,
	0x93, 0xe2, 0xc0, 0x98, 0xef, 0xe1, 0xda, 0xb0, 0x05, 0xb9, 0x57, 0x41, 0x0f, 0x7d, 0xe6, 0x4f,
	0x4e, 0x54, 0x2b, 0x2e, 0xe3, 0x76, 0x4d, 0xae, 0x43, 0x6b, 0x98, 0x5c, 0xc7, 0x2c, 0xf1, 0x11,
	0xba, 0xe6, 0xae, 0xc9, 0xe4, 0xec, 0x25, 0xfb, 0xf7, 0x4f, 0x95, 0x15, 0xdd, 0xca, 0x5d, 0x3b,
	0xea, 0xf1, 0xf9, 0xea, 0xc8, 0xa7, 0xe5, 0xc9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0x13,
	0x6f, 0xf3, 0x9a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalendarEventClient is the client API for CalendarEvent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalendarEventClient interface {
	AddEvent(ctx context.Context, in *AddEventRequest, opts ...grpc.CallOption) (*AddEventResponse, error)
	GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*GetEventResponse, error)
	UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*UpdateEventResponse, error)
	DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*DeleteEventResponse, error)
	GetEventList(ctx context.Context, in *GetEventListRequest, opts ...grpc.CallOption) (*GetEventListResponse, error)
}

type calendarEventClient struct {
	cc *grpc.ClientConn
}

func NewCalendarEventClient(cc *grpc.ClientConn) CalendarEventClient {
	return &calendarEventClient{cc}
}

func (c *calendarEventClient) AddEvent(ctx context.Context, in *AddEventRequest, opts ...grpc.CallOption) (*AddEventResponse, error) {
	out := new(AddEventResponse)
	err := c.cc.Invoke(ctx, "/calendar.CalendarEvent/AddEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarEventClient) GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*GetEventResponse, error) {
	out := new(GetEventResponse)
	err := c.cc.Invoke(ctx, "/calendar.CalendarEvent/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarEventClient) UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*UpdateEventResponse, error) {
	out := new(UpdateEventResponse)
	err := c.cc.Invoke(ctx, "/calendar.CalendarEvent/UpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarEventClient) DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*DeleteEventResponse, error) {
	out := new(DeleteEventResponse)
	err := c.cc.Invoke(ctx, "/calendar.CalendarEvent/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarEventClient) GetEventList(ctx context.Context, in *GetEventListRequest, opts ...grpc.CallOption) (*GetEventListResponse, error) {
	out := new(GetEventListResponse)
	err := c.cc.Invoke(ctx, "/calendar.CalendarEvent/GetEventList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarEventServer is the server API for CalendarEvent service.
type CalendarEventServer interface {
	AddEvent(context.Context, *AddEventRequest) (*AddEventResponse, error)
	GetEvent(context.Context, *GetEventRequest) (*GetEventResponse, error)
	UpdateEvent(context.Context, *UpdateEventRequest) (*UpdateEventResponse, error)
	DeleteEvent(context.Context, *DeleteEventRequest) (*DeleteEventResponse, error)
	GetEventList(context.Context, *GetEventListRequest) (*GetEventListResponse, error)
}

// UnimplementedCalendarEventServer can be embedded to have forward compatible implementations.
type UnimplementedCalendarEventServer struct {
}

func (*UnimplementedCalendarEventServer) AddEvent(ctx context.Context, req *AddEventRequest) (*AddEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEvent not implemented")
}
func (*UnimplementedCalendarEventServer) GetEvent(ctx context.Context, req *GetEventRequest) (*GetEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (*UnimplementedCalendarEventServer) UpdateEvent(ctx context.Context, req *UpdateEventRequest) (*UpdateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (*UnimplementedCalendarEventServer) DeleteEvent(ctx context.Context, req *DeleteEventRequest) (*DeleteEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (*UnimplementedCalendarEventServer) GetEventList(ctx context.Context, req *GetEventListRequest) (*GetEventListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventList not implemented")
}

func RegisterCalendarEventServer(s *grpc.Server, srv CalendarEventServer) {
	s.RegisterService(&_CalendarEvent_serviceDesc, srv)
}

func _CalendarEvent_AddEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarEventServer).AddEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.CalendarEvent/AddEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarEventServer).AddEvent(ctx, req.(*AddEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarEvent_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarEventServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.CalendarEvent/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarEventServer).GetEvent(ctx, req.(*GetEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarEvent_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarEventServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.CalendarEvent/UpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarEventServer).UpdateEvent(ctx, req.(*UpdateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarEvent_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarEventServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.CalendarEvent/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarEventServer).DeleteEvent(ctx, req.(*DeleteEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarEvent_GetEventList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarEventServer).GetEventList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.CalendarEvent/GetEventList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarEventServer).GetEventList(ctx, req.(*GetEventListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalendarEvent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calendar.CalendarEvent",
	HandlerType: (*CalendarEventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEvent",
			Handler:    _CalendarEvent_AddEvent_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _CalendarEvent_GetEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _CalendarEvent_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _CalendarEvent_DeleteEvent_Handler,
		},
		{
			MethodName: "GetEventList",
			Handler:    _CalendarEvent_GetEventList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/server/server.proto",
}

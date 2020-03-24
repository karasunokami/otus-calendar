// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/api/proto/api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Event struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea8a93f38f8eab, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Event) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Event) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type CreateEventRequest struct {
	Title                string               `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateEventRequest) Reset()         { *m = CreateEventRequest{} }
func (m *CreateEventRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEventRequest) ProtoMessage()    {}
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea8a93f38f8eab, []int{1}
}

func (m *CreateEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventRequest.Unmarshal(m, b)
}
func (m *CreateEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventRequest.Marshal(b, m, deterministic)
}
func (m *CreateEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventRequest.Merge(m, src)
}
func (m *CreateEventRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEventRequest.Size(m)
}
func (m *CreateEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventRequest proto.InternalMessageInfo

func (m *CreateEventRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateEventRequest) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *CreateEventRequest) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type CreateEventResponse struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEventResponse) Reset()         { *m = CreateEventResponse{} }
func (m *CreateEventResponse) String() string { return proto.CompactTextString(m) }
func (*CreateEventResponse) ProtoMessage()    {}
func (*CreateEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea8a93f38f8eab, []int{2}
}

func (m *CreateEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventResponse.Unmarshal(m, b)
}
func (m *CreateEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventResponse.Marshal(b, m, deterministic)
}
func (m *CreateEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventResponse.Merge(m, src)
}
func (m *CreateEventResponse) XXX_Size() int {
	return xxx_messageInfo_CreateEventResponse.Size(m)
}
func (m *CreateEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventResponse proto.InternalMessageInfo

func (m *CreateEventResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type UpdateEventRequest struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text                 string               `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UpdateEventRequest) Reset()         { *m = UpdateEventRequest{} }
func (m *UpdateEventRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateEventRequest) ProtoMessage()    {}
func (*UpdateEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea8a93f38f8eab, []int{3}
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

func (m *UpdateEventRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateEventRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdateEventRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *UpdateEventRequest) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *UpdateEventRequest) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type UpdateEventResponse struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateEventResponse) Reset()         { *m = UpdateEventResponse{} }
func (m *UpdateEventResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateEventResponse) ProtoMessage()    {}
func (*UpdateEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea8a93f38f8eab, []int{4}
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

func (m *UpdateEventResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
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
	return fileDescriptor_aaea8a93f38f8eab, []int{5}
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
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEventResponse) Reset()         { *m = DeleteEventResponse{} }
func (m *DeleteEventResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteEventResponse) ProtoMessage()    {}
func (*DeleteEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea8a93f38f8eab, []int{6}
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
	return fileDescriptor_aaea8a93f38f8eab, []int{7}
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
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventResponse) Reset()         { *m = GetEventResponse{} }
func (m *GetEventResponse) String() string { return proto.CompactTextString(m) }
func (*GetEventResponse) ProtoMessage()    {}
func (*GetEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea8a93f38f8eab, []int{8}
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

func (m *GetEventResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "Event")
	proto.RegisterType((*CreateEventRequest)(nil), "CreateEventRequest")
	proto.RegisterType((*CreateEventResponse)(nil), "CreateEventResponse")
	proto.RegisterType((*UpdateEventRequest)(nil), "UpdateEventRequest")
	proto.RegisterType((*UpdateEventResponse)(nil), "UpdateEventResponse")
	proto.RegisterType((*DeleteEventRequest)(nil), "DeleteEventRequest")
	proto.RegisterType((*DeleteEventResponse)(nil), "DeleteEventResponse")
	proto.RegisterType((*GetEventRequest)(nil), "GetEventRequest")
	proto.RegisterType((*GetEventResponse)(nil), "GetEventResponse")
}

func init() {
	proto.RegisterFile("internal/api/proto/api.proto", fileDescriptor_aaea8a93f38f8eab)
}

var fileDescriptor_aaea8a93f38f8eab = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0x59, 0xa4, 0x08, 0x43, 0x22, 0x38, 0xad, 0x09, 0x69, 0x48, 0xc4, 0xc6, 0x03, 0xa7,
	0x45, 0x21, 0x1e, 0xf4, 0x8a, 0xc6, 0x7b, 0xd5, 0xb3, 0x29, 0x76, 0x24, 0x4d, 0xca, 0xb6, 0xb6,
	0x0b, 0xf1, 0x61, 0x4c, 0x7c, 0x15, 0x1f, 0xcb, 0xa3, 0xe9, 0x6e, 0x08, 0x50, 0x10, 0x2b, 0xb7,
	0x99, 0xe9, 0xcc, 0xce, 0xff, 0xcd, 0x5f, 0xe8, 0x04, 0x42, 0x52, 0x22, 0xbc, 0xb0, 0xef, 0xc5,
	0x41, 0x3f, 0x4e, 0x22, 0x19, 0x65, 0x11, 0x57, 0x91, 0x7d, 0x3a, 0x89, 0xa2, 0x49, 0x48, 0xba,
	0x3e, 0x9e, 0xbd, 0xf6, 0x65, 0x30, 0xa5, 0x54, 0x7a, 0xd3, 0x58, 0x37, 0x38, 0x9f, 0x0c, 0x8c,
	0xbb, 0x39, 0x09, 0x89, 0x47, 0x50, 0x0e, 0xfc, 0x36, 0xeb, 0xb2, 0x5e, 0xdd, 0x2d, 0x07, 0x3e,
	0x5a, 0x60, 0xc8, 0x40, 0x86, 0xd4, 0x2e, 0xab, 0x92, 0x4e, 0xf0, 0x1a, 0x20, 0x95, 0x5e, 0x22,
	0x9f, 0xb3, 0x87, 0xda, 0x46, 0x97, 0xf5, 0x1a, 0x03, 0x9b, 0xeb, 0x2d, 0x7c, 0xb1, 0x85, 0x3f,
	0x2e, 0xb6, 0xb8, 0x75, 0xd5, 0x9d, 0xe5, 0x78, 0x05, 0x35, 0x12, 0xbe, 0x1e, 0xac, 0xfe, 0x39,
	0x78, 0x48, 0xc2, 0xcf, 0x32, 0xe7, 0x83, 0x01, 0x8e, 0x12, 0xf2, 0x24, 0x29, 0x9d, 0x2e, 0xbd,
	0xcd, 0x28, 0x95, 0x4b, 0x79, 0xec, 0x77, 0x79, 0x07, 0xfb, 0xca, 0xab, 0x14, 0x97, 0x37, 0x04,
	0x73, 0x4d, 0x5d, 0x1a, 0x47, 0x22, 0x25, 0xec, 0x80, 0x41, 0x59, 0x41, 0xc9, 0x6b, 0x0c, 0xaa,
	0x5c, 0x7f, 0xd6, 0x45, 0xe7, 0x8b, 0x01, 0x3e, 0xc5, 0x7e, 0x9e, 0xa9, 0x98, 0x05, 0x08, 0x15,
	0x49, 0xef, 0x52, 0xd1, 0xd5, 0x5d, 0x15, 0xe7, 0xb8, 0x2b, 0xfb, 0x72, 0x1b, 0xff, 0xe2, 0x5e,
	0x23, 0x28, 0xc4, 0x7d, 0x0e, 0x78, 0x4b, 0x21, 0xed, 0xc6, 0x76, 0x4e, 0xc0, 0x5c, 0xeb, 0xd2,
	0x4f, 0x3b, 0x67, 0xd0, 0xbc, 0x27, 0xb9, 0x73, 0xf2, 0x02, 0x5a, 0xcb, 0x96, 0x22, 0x8a, 0x06,
	0xdf, 0x0c, 0x9a, 0x23, 0x2f, 0x24, 0xe1, 0x7b, 0xc9, 0x03, 0x25, 0xf3, 0xe0, 0x85, 0xf0, 0x06,
	0x1a, 0x2b, 0x96, 0xa2, 0xc9, 0x37, 0x7f, 0x3f, 0xdb, 0xe2, 0x5b, 0x5c, 0x77, 0x4a, 0xd9, 0xec,
	0xca, 0x59, 0xd0, 0xe4, 0x9b, 0x36, 0xdb, 0x16, 0xdf, 0x72, 0x39, 0xa7, 0x84, 0x97, 0x50, 0x5b,
	0xa8, 0xc7, 0x16, 0xcf, 0xb1, 0xda, 0xc7, 0x3c, 0x8f, 0xa6, 0xd7, 0xad, 0x9c, 0x0a, 0x4d, 0xbe,
	0x79, 0x5e, 0xdb, 0xe2, 0xdb, 0xae, 0x59, 0x1a, 0x57, 0x95, 0xbd, 0xc3, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x9e, 0xc7, 0xd8, 0xd0, 0x42, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalendarServiceClient is the client API for CalendarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalendarServiceClient interface {
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error)
	UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*UpdateEventResponse, error)
	GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*GetEventResponse, error)
	DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*DeleteEventResponse, error)
}

type calendarServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarServiceClient(cc grpc.ClientConnInterface) CalendarServiceClient {
	return &calendarServiceClient{cc}
}

func (c *calendarServiceClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error) {
	out := new(CreateEventResponse)
	err := c.cc.Invoke(ctx, "/CalendarService/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*UpdateEventResponse, error) {
	out := new(UpdateEventResponse)
	err := c.cc.Invoke(ctx, "/CalendarService/UpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*GetEventResponse, error) {
	out := new(GetEventResponse)
	err := c.cc.Invoke(ctx, "/CalendarService/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*DeleteEventResponse, error) {
	out := new(DeleteEventResponse)
	err := c.cc.Invoke(ctx, "/CalendarService/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServiceServer is the server API for CalendarService service.
type CalendarServiceServer interface {
	CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error)
	UpdateEvent(context.Context, *UpdateEventRequest) (*UpdateEventResponse, error)
	GetEvent(context.Context, *GetEventRequest) (*GetEventResponse, error)
	DeleteEvent(context.Context, *DeleteEventRequest) (*DeleteEventResponse, error)
}

// UnimplementedCalendarServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalendarServiceServer struct {
}

func (*UnimplementedCalendarServiceServer) CreateEvent(ctx context.Context, req *CreateEventRequest) (*CreateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (*UnimplementedCalendarServiceServer) UpdateEvent(ctx context.Context, req *UpdateEventRequest) (*UpdateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (*UnimplementedCalendarServiceServer) GetEvent(ctx context.Context, req *GetEventRequest) (*GetEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (*UnimplementedCalendarServiceServer) DeleteEvent(ctx context.Context, req *DeleteEventRequest) (*DeleteEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}

func RegisterCalendarServiceServer(s *grpc.Server, srv CalendarServiceServer) {
	s.RegisterService(&_CalendarService_serviceDesc, srv)
}

func _CalendarService_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CalendarService/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).CreateEvent(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CalendarService/UpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).UpdateEvent(ctx, req.(*UpdateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CalendarService/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).GetEvent(ctx, req.(*GetEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CalendarService/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).DeleteEvent(ctx, req.(*DeleteEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalendarService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CalendarService",
	HandlerType: (*CalendarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _CalendarService_CreateEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _CalendarService_UpdateEvent_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _CalendarService_GetEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _CalendarService_DeleteEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/api/proto/api.proto",
}
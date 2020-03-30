// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/calendar.proto

package calendar

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
	return fileDescriptor_46317338ada08f2e, []int{0}
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
	return fileDescriptor_46317338ada08f2e, []int{1}
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
	return fileDescriptor_46317338ada08f2e, []int{2}
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
	return fileDescriptor_46317338ada08f2e, []int{3}
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
	return fileDescriptor_46317338ada08f2e, []int{4}
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
	return fileDescriptor_46317338ada08f2e, []int{5}
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
	return fileDescriptor_46317338ada08f2e, []int{6}
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
	return fileDescriptor_46317338ada08f2e, []int{7}
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
	return fileDescriptor_46317338ada08f2e, []int{8}
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
	proto.RegisterFile("api/calendar.proto", fileDescriptor_46317338ada08f2e)
}

var fileDescriptor_46317338ada08f2e = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x4e, 0xfa, 0x40,
	0x10, 0xc6, 0x59, 0xfe, 0x94, 0x3f, 0x0c, 0x89, 0xe0, 0xb4, 0x26, 0xa4, 0x31, 0x11, 0x37, 0x1e,
	0x38, 0x2d, 0x0a, 0xf1, 0xa0, 0x57, 0x34, 0xde, 0xab, 0x9e, 0x4d, 0xa1, 0x23, 0x69, 0x52, 0xda,
	0xda, 0x2e, 0xc4, 0x87, 0x31, 0xf1, 0x55, 0x7c, 0x2c, 0x8f, 0xa6, 0xbb, 0x12, 0xa0, 0x20, 0x56,
	0x6e, 0x9d, 0xed, 0xcc, 0xce, 0xf7, 0xfb, 0xbe, 0x16, 0xd0, 0x8d, 0xfd, 0xde, 0xd8, 0x0d, 0x28,
	0xf4, 0xdc, 0x44, 0xc4, 0x49, 0x24, 0x23, 0xfb, 0x64, 0x12, 0x45, 0x93, 0x80, 0x7a, 0xaa, 0x1a,
	0xcd, 0x9e, 0x7b, 0xd2, 0x9f, 0x52, 0x2a, 0xdd, 0x69, 0xac, 0x1b, 0xf8, 0x3b, 0x03, 0xe3, 0x76,
	0x4e, 0xa1, 0xc4, 0x03, 0x28, 0xfb, 0x5e, 0x9b, 0x75, 0x58, 0xb7, 0xee, 0x94, 0x7d, 0x0f, 0x2d,
	0x30, 0xa4, 0x2f, 0x03, 0x6a, 0x97, 0xd5, 0x91, 0x2e, 0xf0, 0x0a, 0x20, 0x95, 0x6e, 0x22, 0x9f,
	0xb2, 0x8b, 0xda, 0x46, 0x87, 0x75, 0x1b, 0x7d, 0x5b, 0xe8, 0x2d, 0x62, 0xb1, 0x45, 0x3c, 0x2c,
	0xb6, 0x38, 0x75, 0xd5, 0x9d, 0xd5, 0x78, 0x09, 0x35, 0x0a, 0x3d, 0x3d, 0x58, 0xfd, 0x75, 0xf0,
	0x3f, 0x85, 0x5e, 0x56, 0xf1, 0x37, 0x06, 0x38, 0x4c, 0xc8, 0x95, 0xa4, 0x74, 0x3a, 0xf4, 0x32,
	0xa3, 0x54, 0x2e, 0xe5, 0xb1, 0x9f, 0xe5, 0xfd, 0xdb, 0x57, 0x5e, 0xa5, 0xb8, 0xbc, 0x01, 0x98,
	0x6b, 0xea, 0xd2, 0x38, 0x0a, 0x53, 0xc2, 0x63, 0x30, 0x28, 0x3b, 0x50, 0xf2, 0x1a, 0xfd, 0xaa,
	0xd0, 0xaf, 0xf5, 0x21, 0xff, 0x60, 0x80, 0x8f, 0xb1, 0x97, 0x67, 0x2a, 0x16, 0x01, 0x42, 0x45,
	0xd2, 0xab, 0x54, 0x74, 0x75, 0x47, 0x3d, 0xe7, 0xb8, 0x2b, 0xfb, 0x72, 0x1b, 0x7f, 0xe2, 0x5e,
	0x23, 0x28, 0xc4, 0x7d, 0x06, 0x78, 0x43, 0x01, 0xed, 0xc6, 0xe6, 0x47, 0x60, 0xae, 0x75, 0xe9,
	0xab, 0xf9, 0x29, 0x34, 0xef, 0x48, 0xee, 0x9c, 0x3c, 0x87, 0xd6, 0xb2, 0xa5, 0x88, 0xa2, 0xfe,
	0x27, 0x83, 0xe6, 0xf0, 0xfb, 0x9f, 0xb9, 0xa7, 0x64, 0xee, 0x8f, 0x09, 0xaf, 0xa1, 0xb1, 0x12,
	0x29, 0x9a, 0x62, 0xf3, 0xf3, 0xb3, 0x2d, 0xb1, 0x25, 0x75, 0x5e, 0xca, 0x66, 0x57, 0x6c, 0x41,
	0x53, 0x6c, 0xc6, 0x6c, 0x5b, 0x62, 0x8b, 0x73, 0xbc, 0x84, 0x17, 0x50, 0x5b, 0xa8, 0xc7, 0x96,
	0xc8, 0xb1, 0xda, 0x87, 0x22, 0x8f, 0xa6, 0xd7, 0xad, 0x58, 0x85, 0xa6, 0xd8, 0xb4, 0xd7, 0xb6,
	0xc4, 0x36, 0x37, 0x4b, 0xa3, 0xaa, 0x8a, 0x77, 0xf0, 0x15, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x3b,
	0x95, 0x0b, 0x38, 0x04, 0x00, 0x00,
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
	Metadata: "api/calendar.proto",
}
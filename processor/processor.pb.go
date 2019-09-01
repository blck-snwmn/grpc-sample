// Code generated by protoc-gen-go. DO NOT EDIT.
// source: processor.proto

package processer

import (
	context "context"
	fmt "fmt"
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

type NotificationRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotificationRequest) Reset()         { *m = NotificationRequest{} }
func (m *NotificationRequest) String() string { return proto.CompactTextString(m) }
func (*NotificationRequest) ProtoMessage()    {}
func (*NotificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6783724e039e1aa6, []int{0}
}

func (m *NotificationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationRequest.Unmarshal(m, b)
}
func (m *NotificationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationRequest.Marshal(b, m, deterministic)
}
func (m *NotificationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationRequest.Merge(m, src)
}
func (m *NotificationRequest) XXX_Size() int {
	return xxx_messageInfo_NotificationRequest.Size(m)
}
func (m *NotificationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationRequest proto.InternalMessageInfo

type Notification struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_6783724e039e1aa6, []int{1}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Process struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Process) Reset()         { *m = Process{} }
func (m *Process) String() string { return proto.CompactTextString(m) }
func (*Process) ProtoMessage()    {}
func (*Process) Descriptor() ([]byte, []int) {
	return fileDescriptor_6783724e039e1aa6, []int{2}
}

func (m *Process) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process.Unmarshal(m, b)
}
func (m *Process) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process.Marshal(b, m, deterministic)
}
func (m *Process) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process.Merge(m, src)
}
func (m *Process) XXX_Size() int {
	return xxx_messageInfo_Process.Size(m)
}
func (m *Process) XXX_DiscardUnknown() {
	xxx_messageInfo_Process.DiscardUnknown(m)
}

var xxx_messageInfo_Process proto.InternalMessageInfo

type RegisteredMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisteredMessage) Reset()         { *m = RegisteredMessage{} }
func (m *RegisteredMessage) String() string { return proto.CompactTextString(m) }
func (*RegisteredMessage) ProtoMessage()    {}
func (*RegisteredMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_6783724e039e1aa6, []int{3}
}

func (m *RegisteredMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisteredMessage.Unmarshal(m, b)
}
func (m *RegisteredMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisteredMessage.Marshal(b, m, deterministic)
}
func (m *RegisteredMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisteredMessage.Merge(m, src)
}
func (m *RegisteredMessage) XXX_Size() int {
	return xxx_messageInfo_RegisteredMessage.Size(m)
}
func (m *RegisteredMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisteredMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RegisteredMessage proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NotificationRequest)(nil), "processer.NotificationRequest")
	proto.RegisterType((*Notification)(nil), "processer.Notification")
	proto.RegisterType((*Process)(nil), "processer.Process")
	proto.RegisterType((*RegisteredMessage)(nil), "processer.RegisteredMessage")
}

func init() { proto.RegisterFile("processor.proto", fileDescriptor_6783724e039e1aa6) }

var fileDescriptor_6783724e039e1aa6 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0xca, 0x4f,
	0x4e, 0x2d, 0x2e, 0xce, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0a, 0xa4,
	0x16, 0x29, 0x89, 0x72, 0x09, 0xfb, 0xe5, 0x97, 0x64, 0xa6, 0x65, 0x26, 0x27, 0x96, 0x64, 0xe6,
	0xe7, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x28, 0x69, 0x70, 0xf1, 0x20, 0x0b, 0x0b, 0x49,
	0x70, 0xb1, 0xe7, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0xc1, 0xb8, 0x4a, 0x9c, 0x5c, 0xec, 0x01, 0x10, 0xd3, 0x94, 0x84, 0xb9, 0x04, 0x83, 0x52, 0xd3,
	0x33, 0x8b, 0x4b, 0x52, 0x8b, 0x52, 0x53, 0x7c, 0x21, 0xf2, 0x46, 0x8b, 0x18, 0xb9, 0x38, 0x03,
	0x60, 0xf6, 0x0b, 0x05, 0x70, 0x09, 0x43, 0xad, 0x40, 0x31, 0x5e, 0x4e, 0x0f, 0xee, 0x22, 0x3d,
	0x2c, 0xce, 0x91, 0x12, 0xc7, 0x21, 0x6f, 0xc0, 0x28, 0xe4, 0xcc, 0xc5, 0x0f, 0xb3, 0x14, 0x6a,
	0x8d, 0x90, 0x10, 0x92, 0x6a, 0xa8, 0x98, 0x94, 0x0c, 0x92, 0x18, 0x86, 0x23, 0x93, 0xd8, 0xc0,
	0xe1, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xbd, 0x79, 0x9b, 0x2a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProcessorClient is the client API for Processor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProcessorClient interface {
	RequestNotification(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (Processor_RequestNotificationClient, error)
	RegisterProcess(ctx context.Context, in *Process, opts ...grpc.CallOption) (*RegisteredMessage, error)
}

type processorClient struct {
	cc *grpc.ClientConn
}

func NewProcessorClient(cc *grpc.ClientConn) ProcessorClient {
	return &processorClient{cc}
}

func (c *processorClient) RequestNotification(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (Processor_RequestNotificationClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Processor_serviceDesc.Streams[0], "/processer.Processor/RequestNotification", opts...)
	if err != nil {
		return nil, err
	}
	x := &processorRequestNotificationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Processor_RequestNotificationClient interface {
	Recv() (*Notification, error)
	grpc.ClientStream
}

type processorRequestNotificationClient struct {
	grpc.ClientStream
}

func (x *processorRequestNotificationClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *processorClient) RegisterProcess(ctx context.Context, in *Process, opts ...grpc.CallOption) (*RegisteredMessage, error) {
	out := new(RegisteredMessage)
	err := c.cc.Invoke(ctx, "/processer.Processor/RegisterProcess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessorServer is the server API for Processor service.
type ProcessorServer interface {
	RequestNotification(*NotificationRequest, Processor_RequestNotificationServer) error
	RegisterProcess(context.Context, *Process) (*RegisteredMessage, error)
}

// UnimplementedProcessorServer can be embedded to have forward compatible implementations.
type UnimplementedProcessorServer struct {
}

func (*UnimplementedProcessorServer) RequestNotification(req *NotificationRequest, srv Processor_RequestNotificationServer) error {
	return status.Errorf(codes.Unimplemented, "method RequestNotification not implemented")
}
func (*UnimplementedProcessorServer) RegisterProcess(ctx context.Context, req *Process) (*RegisteredMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterProcess not implemented")
}

func RegisterProcessorServer(s *grpc.Server, srv ProcessorServer) {
	s.RegisterService(&_Processor_serviceDesc, srv)
}

func _Processor_RequestNotification_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NotificationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProcessorServer).RequestNotification(m, &processorRequestNotificationServer{stream})
}

type Processor_RequestNotificationServer interface {
	Send(*Notification) error
	grpc.ServerStream
}

type processorRequestNotificationServer struct {
	grpc.ServerStream
}

func (x *processorRequestNotificationServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

func _Processor_RegisterProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Process)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessorServer).RegisterProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/processer.Processor/RegisterProcess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessorServer).RegisterProcess(ctx, req.(*Process))
	}
	return interceptor(ctx, in, info, handler)
}

var _Processor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "processer.Processor",
	HandlerType: (*ProcessorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterProcess",
			Handler:    _Processor_RegisterProcess_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RequestNotification",
			Handler:       _Processor_RequestNotification_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "processor.proto",
}

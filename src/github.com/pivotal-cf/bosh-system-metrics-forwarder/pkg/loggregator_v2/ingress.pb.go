// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ingress.proto

package loggregator_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IngressResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IngressResponse) Reset()         { *m = IngressResponse{} }
func (m *IngressResponse) String() string { return proto.CompactTextString(m) }
func (*IngressResponse) ProtoMessage()    {}
func (*IngressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ingress_1df30cb6e71681fb, []int{0}
}
func (m *IngressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IngressResponse.Unmarshal(m, b)
}
func (m *IngressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IngressResponse.Marshal(b, m, deterministic)
}
func (dst *IngressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngressResponse.Merge(dst, src)
}
func (m *IngressResponse) XXX_Size() int {
	return xxx_messageInfo_IngressResponse.Size(m)
}
func (m *IngressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IngressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IngressResponse proto.InternalMessageInfo

type BatchSenderResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchSenderResponse) Reset()         { *m = BatchSenderResponse{} }
func (m *BatchSenderResponse) String() string { return proto.CompactTextString(m) }
func (*BatchSenderResponse) ProtoMessage()    {}
func (*BatchSenderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ingress_1df30cb6e71681fb, []int{1}
}
func (m *BatchSenderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchSenderResponse.Unmarshal(m, b)
}
func (m *BatchSenderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchSenderResponse.Marshal(b, m, deterministic)
}
func (dst *BatchSenderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchSenderResponse.Merge(dst, src)
}
func (m *BatchSenderResponse) XXX_Size() int {
	return xxx_messageInfo_BatchSenderResponse.Size(m)
}
func (m *BatchSenderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchSenderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchSenderResponse proto.InternalMessageInfo

type SendResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendResponse) Reset()         { *m = SendResponse{} }
func (m *SendResponse) String() string { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()    {}
func (*SendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ingress_1df30cb6e71681fb, []int{2}
}
func (m *SendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResponse.Unmarshal(m, b)
}
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResponse.Marshal(b, m, deterministic)
}
func (dst *SendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResponse.Merge(dst, src)
}
func (m *SendResponse) XXX_Size() int {
	return xxx_messageInfo_SendResponse.Size(m)
}
func (m *SendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*IngressResponse)(nil), "loggregator.v2.IngressResponse")
	proto.RegisterType((*BatchSenderResponse)(nil), "loggregator.v2.BatchSenderResponse")
	proto.RegisterType((*SendResponse)(nil), "loggregator.v2.SendResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IngressClient is the client API for Ingress service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IngressClient interface {
	Sender(ctx context.Context, opts ...grpc.CallOption) (Ingress_SenderClient, error)
	BatchSender(ctx context.Context, opts ...grpc.CallOption) (Ingress_BatchSenderClient, error)
	Send(ctx context.Context, in *EnvelopeBatch, opts ...grpc.CallOption) (*SendResponse, error)
}

type ingressClient struct {
	cc *grpc.ClientConn
}

func NewIngressClient(cc *grpc.ClientConn) IngressClient {
	return &ingressClient{cc}
}

func (c *ingressClient) Sender(ctx context.Context, opts ...grpc.CallOption) (Ingress_SenderClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Ingress_serviceDesc.Streams[0], "/loggregator.v2.Ingress/Sender", opts...)
	if err != nil {
		return nil, err
	}
	x := &ingressSenderClient{stream}
	return x, nil
}

type Ingress_SenderClient interface {
	Send(*Envelope) error
	CloseAndRecv() (*IngressResponse, error)
	grpc.ClientStream
}

type ingressSenderClient struct {
	grpc.ClientStream
}

func (x *ingressSenderClient) Send(m *Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ingressSenderClient) CloseAndRecv() (*IngressResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(IngressResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ingressClient) BatchSender(ctx context.Context, opts ...grpc.CallOption) (Ingress_BatchSenderClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Ingress_serviceDesc.Streams[1], "/loggregator.v2.Ingress/BatchSender", opts...)
	if err != nil {
		return nil, err
	}
	x := &ingressBatchSenderClient{stream}
	return x, nil
}

type Ingress_BatchSenderClient interface {
	Send(*EnvelopeBatch) error
	CloseAndRecv() (*BatchSenderResponse, error)
	grpc.ClientStream
}

type ingressBatchSenderClient struct {
	grpc.ClientStream
}

func (x *ingressBatchSenderClient) Send(m *EnvelopeBatch) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ingressBatchSenderClient) CloseAndRecv() (*BatchSenderResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BatchSenderResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ingressClient) Send(ctx context.Context, in *EnvelopeBatch, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/loggregator.v2.Ingress/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IngressServer is the server API for Ingress service.
type IngressServer interface {
	Sender(Ingress_SenderServer) error
	BatchSender(Ingress_BatchSenderServer) error
	Send(context.Context, *EnvelopeBatch) (*SendResponse, error)
}

func RegisterIngressServer(s *grpc.Server, srv IngressServer) {
	s.RegisterService(&_Ingress_serviceDesc, srv)
}

func _Ingress_Sender_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IngressServer).Sender(&ingressSenderServer{stream})
}

type Ingress_SenderServer interface {
	SendAndClose(*IngressResponse) error
	Recv() (*Envelope, error)
	grpc.ServerStream
}

type ingressSenderServer struct {
	grpc.ServerStream
}

func (x *ingressSenderServer) SendAndClose(m *IngressResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ingressSenderServer) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Ingress_BatchSender_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IngressServer).BatchSender(&ingressBatchSenderServer{stream})
}

type Ingress_BatchSenderServer interface {
	SendAndClose(*BatchSenderResponse) error
	Recv() (*EnvelopeBatch, error)
	grpc.ServerStream
}

type ingressBatchSenderServer struct {
	grpc.ServerStream
}

func (x *ingressBatchSenderServer) SendAndClose(m *BatchSenderResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ingressBatchSenderServer) Recv() (*EnvelopeBatch, error) {
	m := new(EnvelopeBatch)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Ingress_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnvelopeBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngressServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loggregator.v2.Ingress/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngressServer).Send(ctx, req.(*EnvelopeBatch))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ingress_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loggregator.v2.Ingress",
	HandlerType: (*IngressServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Ingress_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sender",
			Handler:       _Ingress_Sender_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BatchSender",
			Handler:       _Ingress_BatchSender_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "ingress.proto",
}

func init() { proto.RegisterFile("ingress.proto", fileDescriptor_ingress_1df30cb6e71681fb) }

var fileDescriptor_ingress_1df30cb6e71681fb = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0xcc, 0x4b, 0x2f,
	0x4a, 0x2d, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcb, 0xc9, 0x4f, 0x4f, 0x2f,
	0x4a, 0x4d, 0x4f, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x33, 0x92, 0xe2, 0x4b, 0xcd, 0x2b, 0x4b, 0xcd,
	0xc9, 0x2f, 0x48, 0x85, 0xc8, 0x2b, 0x09, 0x72, 0xf1, 0x7b, 0x42, 0x34, 0x04, 0xa5, 0x16, 0x17,
	0xe4, 0xe7, 0x15, 0xa7, 0x2a, 0x89, 0x72, 0x09, 0x3b, 0x25, 0x96, 0x24, 0x67, 0x04, 0xa7, 0xe6,
	0xa5, 0xa4, 0x16, 0xc1, 0x85, 0xf9, 0xb8, 0x78, 0x40, 0x22, 0x30, 0xbe, 0xd1, 0x07, 0x46, 0x2e,
	0x76, 0xa8, 0x56, 0x21, 0x77, 0x2e, 0x36, 0x88, 0x6a, 0x21, 0x09, 0x3d, 0x54, 0x0b, 0xf5, 0x5c,
	0xa1, 0xf6, 0x49, 0xc9, 0xa3, 0xcb, 0xa0, 0xdb, 0xcb, 0xa0, 0xc1, 0x28, 0x14, 0xca, 0xc5, 0x8d,
	0x64, 0xb7, 0x90, 0x2c, 0x2e, 0xd3, 0xc0, 0x8a, 0xa4, 0x94, 0xd1, 0xa5, 0xb1, 0xb9, 0x1b, 0x64,
	0xac, 0x2b, 0x17, 0x0b, 0x48, 0x94, 0x90, 0x79, 0x32, 0xe8, 0xd2, 0xc8, 0x1e, 0x56, 0x62, 0x70,
	0x32, 0xe5, 0x92, 0xcf, 0x2f, 0x4a, 0xd7, 0x4b, 0xce, 0xc9, 0x2f, 0x4d, 0x49, 0xcb, 0x2f, 0xcd,
	0x4b, 0x29, 0xaa, 0x44, 0xd3, 0xe1, 0x24, 0xe4, 0x83, 0xe0, 0x43, 0x3d, 0x98, 0xc4, 0x06, 0x0e,
	0x6a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0x97, 0x11, 0x79, 0x9b, 0x01, 0x00, 0x00,
}
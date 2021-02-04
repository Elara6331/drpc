// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service.proto

package grpccompat

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	drpc "storj.io/drpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type In struct {
	In                   int64    `protobuf:"varint,1,opt,name=in,proto3" json:"in,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *In) Reset()         { *m = In{} }
func (m *In) String() string { return proto.CompactTextString(m) }
func (*In) ProtoMessage()    {}
func (*In) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}
func (m *In) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_In.Unmarshal(m, b)
}
func (m *In) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_In.Marshal(b, m, deterministic)
}
func (m *In) XXX_Merge(src proto.Message) {
	xxx_messageInfo_In.Merge(m, src)
}
func (m *In) XXX_Size() int {
	return xxx_messageInfo_In.Size(m)
}
func (m *In) XXX_DiscardUnknown() {
	xxx_messageInfo_In.DiscardUnknown(m)
}

var xxx_messageInfo_In proto.InternalMessageInfo

func (m *In) GetIn() int64 {
	if m != nil {
		return m.In
	}
	return 0
}

type Out struct {
	Out                  int64    `protobuf:"varint,1,opt,name=out,proto3" json:"out,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Out) Reset()         { *m = Out{} }
func (m *Out) String() string { return proto.CompactTextString(m) }
func (*Out) ProtoMessage()    {}
func (*Out) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}
func (m *Out) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Out.Unmarshal(m, b)
}
func (m *Out) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Out.Marshal(b, m, deterministic)
}
func (m *Out) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Out.Merge(m, src)
}
func (m *Out) XXX_Size() int {
	return xxx_messageInfo_Out.Size(m)
}
func (m *Out) XXX_DiscardUnknown() {
	xxx_messageInfo_Out.DiscardUnknown(m)
}

var xxx_messageInfo_Out proto.InternalMessageInfo

func (m *Out) GetOut() int64 {
	if m != nil {
		return m.Out
	}
	return 0
}

func init() {
	proto.RegisterType((*In)(nil), "service.In")
	proto.RegisterType((*Out)(nil), "service.Out")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x44, 0xb8,
	0x98, 0x3c, 0xf3, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0xf3, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83,
	0x98, 0x32, 0xf3, 0x94, 0xc4, 0xb9, 0x98, 0xfd, 0x4b, 0x4b, 0x84, 0x04, 0xb8, 0x98, 0xf3, 0x4b,
	0x4b, 0xa0, 0xe2, 0x20, 0xa6, 0xd1, 0x4a, 0x46, 0x2e, 0xf6, 0x60, 0x88, 0x56, 0x21, 0x15, 0x2e,
	0x76, 0xdf, 0xd4, 0x92, 0x8c, 0xfc, 0x14, 0x43, 0x21, 0x6e, 0x3d, 0x98, 0xf1, 0x9e, 0x79, 0x52,
	0x3c, 0x70, 0x0e, 0xc8, 0x0c, 0x35, 0x98, 0x2a, 0x23, 0x3c, 0xaa, 0x34, 0x18, 0x11, 0xea, 0x8c,
	0xf1, 0xa8, 0x33, 0x60, 0x14, 0xd2, 0x80, 0xa9, 0x33, 0xc1, 0x6b, 0x9e, 0x01, 0xa3, 0x13, 0x4f,
	0x14, 0x57, 0x7a, 0x51, 0x41, 0x72, 0x72, 0x7e, 0x6e, 0x41, 0x62, 0x49, 0x12, 0x1b, 0xd8, 0xe3,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x3b, 0x1a, 0xb7, 0x09, 0x01, 0x00, 0x00,
}

// --- DRPC BEGIN ---

type DRPCServiceClient interface {
	DRPCConn() drpc.Conn

	Method1(ctx context.Context, in *In) (*Out, error)
	Method2(ctx context.Context) (DRPCService_Method2Client, error)
	Method3(ctx context.Context, in *In) (DRPCService_Method3Client, error)
	Method4(ctx context.Context) (DRPCService_Method4Client, error)
}

type drpcServiceClient struct {
	cc drpc.Conn
}

func NewDRPCServiceClient(cc drpc.Conn) DRPCServiceClient {
	return &drpcServiceClient{cc}
}

func (c *drpcServiceClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcServiceClient) Method1(ctx context.Context, in *In) (*Out, error) {
	out := new(Out)
	err := c.cc.Invoke(ctx, "/service.Service/Method1", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcServiceClient) Method2(ctx context.Context) (DRPCService_Method2Client, error) {
	stream, err := c.cc.NewStream(ctx, "/service.Service/Method2")
	if err != nil {
		return nil, err
	}
	x := &drpcServiceMethod2Client{stream}
	return x, nil
}

type DRPCService_Method2Client interface {
	drpc.Stream
	Send(*In) error
	CloseAndRecv() (*Out, error)
}

type drpcServiceMethod2Client struct {
	drpc.Stream
}

func (x *drpcServiceMethod2Client) Send(m *In) error {
	return x.MsgSend(m)
}

func (x *drpcServiceMethod2Client) CloseAndRecv() (*Out, error) {
	if err := x.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Out)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *drpcServiceClient) Method3(ctx context.Context, in *In) (DRPCService_Method3Client, error) {
	stream, err := c.cc.NewStream(ctx, "/service.Service/Method3")
	if err != nil {
		return nil, err
	}
	x := &drpcServiceMethod3Client{stream}
	if err := x.MsgSend(in); err != nil {
		return nil, err
	}
	if err := x.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DRPCService_Method3Client interface {
	drpc.Stream
	Recv() (*Out, error)
}

type drpcServiceMethod3Client struct {
	drpc.Stream
}

func (x *drpcServiceMethod3Client) Recv() (*Out, error) {
	m := new(Out)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *drpcServiceClient) Method4(ctx context.Context) (DRPCService_Method4Client, error) {
	stream, err := c.cc.NewStream(ctx, "/service.Service/Method4")
	if err != nil {
		return nil, err
	}
	x := &drpcServiceMethod4Client{stream}
	return x, nil
}

type DRPCService_Method4Client interface {
	drpc.Stream
	Send(*In) error
	Recv() (*Out, error)
}

type drpcServiceMethod4Client struct {
	drpc.Stream
}

func (x *drpcServiceMethod4Client) Send(m *In) error {
	return x.MsgSend(m)
}

func (x *drpcServiceMethod4Client) Recv() (*Out, error) {
	m := new(Out)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

type DRPCServiceServer interface {
	Method1(context.Context, *In) (*Out, error)
	Method2(DRPCService_Method2Stream) error
	Method3(*In, DRPCService_Method3Stream) error
	Method4(DRPCService_Method4Stream) error
}

type DRPCServiceDescription struct{}

func (DRPCServiceDescription) NumMethods() int { return 4 }

func (DRPCServiceDescription) Method(n int) (string, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/service.Service/Method1",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCServiceServer).
					Method1(
						ctx,
						in1.(*In),
					)
			}, DRPCServiceServer.Method1, true
	case 1:
		return "/service.Service/Method2",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCServiceServer).
					Method2(
						&drpcServiceMethod2Stream{in1.(drpc.Stream)},
					)
			}, DRPCServiceServer.Method2, true
	case 2:
		return "/service.Service/Method3",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCServiceServer).
					Method3(
						in1.(*In),
						&drpcServiceMethod3Stream{in2.(drpc.Stream)},
					)
			}, DRPCServiceServer.Method3, true
	case 3:
		return "/service.Service/Method4",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCServiceServer).
					Method4(
						&drpcServiceMethod4Stream{in1.(drpc.Stream)},
					)
			}, DRPCServiceServer.Method4, true
	default:
		return "", nil, nil, false
	}
}

func DRPCRegisterService(mux drpc.Mux, impl DRPCServiceServer) error {
	return mux.Register(impl, DRPCServiceDescription{})
}

type DRPCService_Method1Stream interface {
	drpc.Stream
	SendAndClose(*Out) error
}

type drpcServiceMethod1Stream struct {
	drpc.Stream
}

func (x *drpcServiceMethod1Stream) SendAndClose(m *Out) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCService_Method2Stream interface {
	drpc.Stream
	SendAndClose(*Out) error
	Recv() (*In, error)
}

type drpcServiceMethod2Stream struct {
	drpc.Stream
}

func (x *drpcServiceMethod2Stream) SendAndClose(m *Out) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

func (x *drpcServiceMethod2Stream) Recv() (*In, error) {
	m := new(In)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

type DRPCService_Method3Stream interface {
	drpc.Stream
	Send(*Out) error
}

type drpcServiceMethod3Stream struct {
	drpc.Stream
}

func (x *drpcServiceMethod3Stream) Send(m *Out) error {
	return x.MsgSend(m)
}

type DRPCService_Method4Stream interface {
	drpc.Stream
	Send(*Out) error
	Recv() (*In, error)
}

type drpcServiceMethod4Stream struct {
	drpc.Stream
}

func (x *drpcServiceMethod4Stream) Send(m *Out) error {
	return x.MsgSend(m)
}

func (x *drpcServiceMethod4Stream) Recv() (*In, error) {
	m := new(In)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

// --- DRPC END ---

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	Method1(ctx context.Context, in *In, opts ...grpc.CallOption) (*Out, error)
	Method2(ctx context.Context, opts ...grpc.CallOption) (Service_Method2Client, error)
	Method3(ctx context.Context, in *In, opts ...grpc.CallOption) (Service_Method3Client, error)
	Method4(ctx context.Context, opts ...grpc.CallOption) (Service_Method4Client, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Method1(ctx context.Context, in *In, opts ...grpc.CallOption) (*Out, error) {
	out := new(Out)
	err := c.cc.Invoke(ctx, "/service.Service/Method1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Method2(ctx context.Context, opts ...grpc.CallOption) (Service_Method2Client, error) {
	stream, err := c.cc.NewStream(ctx, &_Service_serviceDesc.Streams[0], "/service.Service/Method2", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceMethod2Client{stream}
	return x, nil
}

type Service_Method2Client interface {
	Send(*In) error
	CloseAndRecv() (*Out, error)
	grpc.ClientStream
}

type serviceMethod2Client struct {
	grpc.ClientStream
}

func (x *serviceMethod2Client) Send(m *In) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceMethod2Client) CloseAndRecv() (*Out, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Out)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) Method3(ctx context.Context, in *In, opts ...grpc.CallOption) (Service_Method3Client, error) {
	stream, err := c.cc.NewStream(ctx, &_Service_serviceDesc.Streams[1], "/service.Service/Method3", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceMethod3Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_Method3Client interface {
	Recv() (*Out, error)
	grpc.ClientStream
}

type serviceMethod3Client struct {
	grpc.ClientStream
}

func (x *serviceMethod3Client) Recv() (*Out, error) {
	m := new(Out)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) Method4(ctx context.Context, opts ...grpc.CallOption) (Service_Method4Client, error) {
	stream, err := c.cc.NewStream(ctx, &_Service_serviceDesc.Streams[2], "/service.Service/Method4", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceMethod4Client{stream}
	return x, nil
}

type Service_Method4Client interface {
	Send(*In) error
	Recv() (*Out, error)
	grpc.ClientStream
}

type serviceMethod4Client struct {
	grpc.ClientStream
}

func (x *serviceMethod4Client) Send(m *In) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceMethod4Client) Recv() (*Out, error) {
	m := new(Out)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Method1(context.Context, *In) (*Out, error)
	Method2(Service_Method2Server) error
	Method3(*In, Service_Method3Server) error
	Method4(Service_Method4Server) error
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Method1(ctx context.Context, req *In) (*Out, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Method1 not implemented")
}
func (*UnimplementedServiceServer) Method2(srv Service_Method2Server) error {
	return status.Errorf(codes.Unimplemented, "method Method2 not implemented")
}
func (*UnimplementedServiceServer) Method3(req *In, srv Service_Method3Server) error {
	return status.Errorf(codes.Unimplemented, "method Method3 not implemented")
}
func (*UnimplementedServiceServer) Method4(srv Service_Method4Server) error {
	return status.Errorf(codes.Unimplemented, "method Method4 not implemented")
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Method1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(In)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Method1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/Method1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Method1(ctx, req.(*In))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Method2_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).Method2(&serviceMethod2Server{stream})
}

type Service_Method2Server interface {
	SendAndClose(*Out) error
	Recv() (*In, error)
	grpc.ServerStream
}

type serviceMethod2Server struct {
	grpc.ServerStream
}

func (x *serviceMethod2Server) SendAndClose(m *Out) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceMethod2Server) Recv() (*In, error) {
	m := new(In)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Service_Method3_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(In)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).Method3(m, &serviceMethod3Server{stream})
}

type Service_Method3Server interface {
	Send(*Out) error
	grpc.ServerStream
}

type serviceMethod3Server struct {
	grpc.ServerStream
}

func (x *serviceMethod3Server) Send(m *Out) error {
	return x.ServerStream.SendMsg(m)
}

func _Service_Method4_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).Method4(&serviceMethod4Server{stream})
}

type Service_Method4Server interface {
	Send(*Out) error
	Recv() (*In, error)
	grpc.ServerStream
}

type serviceMethod4Server struct {
	grpc.ServerStream
}

func (x *serviceMethod4Server) Send(m *Out) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceMethod4Server) Recv() (*In, error) {
	m := new(In)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Method1",
			Handler:    _Service_Method1_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Method2",
			Handler:       _Service_Method2_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Method3",
			Handler:       _Service_Method3_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Method4",
			Handler:       _Service_Method4_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}

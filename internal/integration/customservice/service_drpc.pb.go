// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: (devel)
// source: service.proto

package service

import (
	context "context"
	errors "errors"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
	customencoding "storj.io/drpc/internal/integration/customencoding"
)

type drpcEncoding_File_service_proto struct{}

func (drpcEncoding_File_service_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return customencoding.Marshal(msg)
}

func (drpcEncoding_File_service_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return customencoding.Unmarshal(buf, msg)
}

func (drpcEncoding_File_service_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	return customencoding.JSONMarshal(msg)
}

func (drpcEncoding_File_service_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return customencoding.JSONUnmarshal(buf, msg)
}

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
	err := c.cc.Invoke(ctx, "/service.Service/Method1", drpcEncoding_File_service_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcServiceClient) Method2(ctx context.Context) (DRPCService_Method2Client, error) {
	stream, err := c.cc.NewStream(ctx, "/service.Service/Method2", drpcEncoding_File_service_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcService_Method2Client{stream}
	return x, nil
}

type DRPCService_Method2Client interface {
	drpc.Stream
	Send(*In) error
	CloseAndRecv() (*Out, error)
}

type drpcService_Method2Client struct {
	drpc.Stream
}

func (x *drpcService_Method2Client) Send(m *In) error {
	return x.MsgSend(m, drpcEncoding_File_service_proto{})
}

func (x *drpcService_Method2Client) CloseAndRecv() (*Out, error) {
	if err := x.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Out)
	if err := x.MsgRecv(m, drpcEncoding_File_service_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *drpcServiceClient) Method3(ctx context.Context, in *In) (DRPCService_Method3Client, error) {
	stream, err := c.cc.NewStream(ctx, "/service.Service/Method3", drpcEncoding_File_service_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcService_Method3Client{stream}
	if err := x.MsgSend(in, drpcEncoding_File_service_proto{}); err != nil {
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

type drpcService_Method3Client struct {
	drpc.Stream
}

func (x *drpcService_Method3Client) Recv() (*Out, error) {
	m := new(Out)
	if err := x.MsgRecv(m, drpcEncoding_File_service_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *drpcServiceClient) Method4(ctx context.Context) (DRPCService_Method4Client, error) {
	stream, err := c.cc.NewStream(ctx, "/service.Service/Method4", drpcEncoding_File_service_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcService_Method4Client{stream}
	return x, nil
}

type DRPCService_Method4Client interface {
	drpc.Stream
	Send(*In) error
	Recv() (*Out, error)
}

type drpcService_Method4Client struct {
	drpc.Stream
}

func (x *drpcService_Method4Client) Send(m *In) error {
	return x.MsgSend(m, drpcEncoding_File_service_proto{})
}

func (x *drpcService_Method4Client) Recv() (*Out, error) {
	m := new(Out)
	if err := x.MsgRecv(m, drpcEncoding_File_service_proto{}); err != nil {
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

type DRPCServiceUnimplementedServer struct{}

func (s *DRPCServiceUnimplementedServer) Method1(context.Context, *In) (*Out, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), 12)
}

func (s *DRPCServiceUnimplementedServer) Method2(DRPCService_Method2Stream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), 12)
}

func (s *DRPCServiceUnimplementedServer) Method3(*In, DRPCService_Method3Stream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), 12)
}

func (s *DRPCServiceUnimplementedServer) Method4(DRPCService_Method4Stream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), 12)
}

type DRPCServiceDescription struct{}

func (DRPCServiceDescription) NumMethods() int { return 4 }

func (DRPCServiceDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/service.Service/Method1", drpcEncoding_File_service_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCServiceServer).
					Method1(
						ctx,
						in1.(*In),
					)
			}, DRPCServiceServer.Method1, true
	case 1:
		return "/service.Service/Method2", drpcEncoding_File_service_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCServiceServer).
					Method2(
						&drpcService_Method2Stream{in1.(drpc.Stream)},
					)
			}, DRPCServiceServer.Method2, true
	case 2:
		return "/service.Service/Method3", drpcEncoding_File_service_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCServiceServer).
					Method3(
						in1.(*In),
						&drpcService_Method3Stream{in2.(drpc.Stream)},
					)
			}, DRPCServiceServer.Method3, true
	case 3:
		return "/service.Service/Method4", drpcEncoding_File_service_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCServiceServer).
					Method4(
						&drpcService_Method4Stream{in1.(drpc.Stream)},
					)
			}, DRPCServiceServer.Method4, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterService(mux drpc.Mux, impl DRPCServiceServer) error {
	return mux.Register(impl, DRPCServiceDescription{})
}

type DRPCService_Method1Stream interface {
	drpc.Stream
	SendAndClose(*Out) error
}

type drpcService_Method1Stream struct {
	drpc.Stream
}

func (x *drpcService_Method1Stream) SendAndClose(m *Out) error {
	if err := x.MsgSend(m, drpcEncoding_File_service_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCService_Method2Stream interface {
	drpc.Stream
	SendAndClose(*Out) error
	Recv() (*In, error)
}

type drpcService_Method2Stream struct {
	drpc.Stream
}

func (x *drpcService_Method2Stream) SendAndClose(m *Out) error {
	if err := x.MsgSend(m, drpcEncoding_File_service_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

func (x *drpcService_Method2Stream) Recv() (*In, error) {
	m := new(In)
	if err := x.MsgRecv(m, drpcEncoding_File_service_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

type DRPCService_Method3Stream interface {
	drpc.Stream
	Send(*Out) error
}

type drpcService_Method3Stream struct {
	drpc.Stream
}

func (x *drpcService_Method3Stream) Send(m *Out) error {
	return x.MsgSend(m, drpcEncoding_File_service_proto{})
}

type DRPCService_Method4Stream interface {
	drpc.Stream
	Send(*Out) error
	Recv() (*In, error)
}

type drpcService_Method4Stream struct {
	drpc.Stream
}

func (x *drpcService_Method4Stream) Send(m *Out) error {
	return x.MsgSend(m, drpcEncoding_File_service_proto{})
}

func (x *drpcService_Method4Stream) Recv() (*In, error) {
	m := new(In)
	if err := x.MsgRecv(m, drpcEncoding_File_service_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

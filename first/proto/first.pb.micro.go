// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/first.proto

package first

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for First service

func NewFirstEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for First service

type FirstService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
	ClientStream(ctx context.Context, opts ...client.CallOption) (First_ClientStreamService, error)
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (First_ServerStreamService, error)
	BidiStream(ctx context.Context, opts ...client.CallOption) (First_BidiStreamService, error)
}

type firstService struct {
	c    client.Client
	name string
}

func NewFirstService(name string, c client.Client) FirstService {
	return &firstService{
		c:    c,
		name: name,
	}
}

func (c *firstService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "First.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firstService) ClientStream(ctx context.Context, opts ...client.CallOption) (First_ClientStreamService, error) {
	req := c.c.NewRequest(c.name, "First.ClientStream", &ClientStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &firstServiceClientStream{stream}, nil
}

type First_ClientStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*ClientStreamRequest) error
}

type firstServiceClientStream struct {
	stream client.Stream
}

func (x *firstServiceClientStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *firstServiceClientStream) Close() error {
	return x.stream.Close()
}

func (x *firstServiceClientStream) Context() context.Context {
	return x.stream.Context()
}

func (x *firstServiceClientStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *firstServiceClientStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *firstServiceClientStream) Send(m *ClientStreamRequest) error {
	return x.stream.Send(m)
}

func (c *firstService) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (First_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "First.ServerStream", &ServerStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &firstServiceServerStream{stream}, nil
}

type First_ServerStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Recv() (*ServerStreamResponse, error)
}

type firstServiceServerStream struct {
	stream client.Stream
}

func (x *firstServiceServerStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *firstServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *firstServiceServerStream) Context() context.Context {
	return x.stream.Context()
}

func (x *firstServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *firstServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *firstServiceServerStream) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *firstService) BidiStream(ctx context.Context, opts ...client.CallOption) (First_BidiStreamService, error) {
	req := c.c.NewRequest(c.name, "First.BidiStream", &BidiStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &firstServiceBidiStream{stream}, nil
}

type First_BidiStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*BidiStreamRequest) error
	Recv() (*BidiStreamResponse, error)
}

type firstServiceBidiStream struct {
	stream client.Stream
}

func (x *firstServiceBidiStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *firstServiceBidiStream) Close() error {
	return x.stream.Close()
}

func (x *firstServiceBidiStream) Context() context.Context {
	return x.stream.Context()
}

func (x *firstServiceBidiStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *firstServiceBidiStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *firstServiceBidiStream) Send(m *BidiStreamRequest) error {
	return x.stream.Send(m)
}

func (x *firstServiceBidiStream) Recv() (*BidiStreamResponse, error) {
	m := new(BidiStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for First service

type FirstHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
	ClientStream(context.Context, First_ClientStreamStream) error
	ServerStream(context.Context, *ServerStreamRequest, First_ServerStreamStream) error
	BidiStream(context.Context, First_BidiStreamStream) error
}

func RegisterFirstHandler(s server.Server, hdlr FirstHandler, opts ...server.HandlerOption) error {
	type first interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
		ClientStream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
		BidiStream(ctx context.Context, stream server.Stream) error
	}
	type First struct {
		first
	}
	h := &firstHandler{hdlr}
	return s.Handle(s.NewHandler(&First{h}, opts...))
}

type firstHandler struct {
	FirstHandler
}

func (h *firstHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.FirstHandler.Call(ctx, in, out)
}

func (h *firstHandler) ClientStream(ctx context.Context, stream server.Stream) error {
	return h.FirstHandler.ClientStream(ctx, &firstClientStreamStream{stream})
}

type First_ClientStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ClientStreamRequest, error)
}

type firstClientStreamStream struct {
	stream server.Stream
}

func (x *firstClientStreamStream) Close() error {
	return x.stream.Close()
}

func (x *firstClientStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *firstClientStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *firstClientStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *firstClientStreamStream) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *firstHandler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(ServerStreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.FirstHandler.ServerStream(ctx, m, &firstServerStreamStream{stream})
}

type First_ServerStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ServerStreamResponse) error
}

type firstServerStreamStream struct {
	stream server.Stream
}

func (x *firstServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *firstServerStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *firstServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *firstServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *firstServerStreamStream) Send(m *ServerStreamResponse) error {
	return x.stream.Send(m)
}

func (h *firstHandler) BidiStream(ctx context.Context, stream server.Stream) error {
	return h.FirstHandler.BidiStream(ctx, &firstBidiStreamStream{stream})
}

type First_BidiStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamResponse) error
	Recv() (*BidiStreamRequest, error)
}

type firstBidiStreamStream struct {
	stream server.Stream
}

func (x *firstBidiStreamStream) Close() error {
	return x.stream.Close()
}

func (x *firstBidiStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *firstBidiStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *firstBidiStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *firstBidiStreamStream) Send(m *BidiStreamResponse) error {
	return x.stream.Send(m)
}

func (x *firstBidiStreamStream) Recv() (*BidiStreamRequest, error) {
	m := new(BidiStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

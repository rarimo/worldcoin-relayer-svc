// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: service.proto

package types

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	StateRelay(ctx context.Context, in *MsgStateRelayRequest, opts ...grpc.CallOption) (*MsgRelayResponse, error)
	GistRelay(ctx context.Context, in *MsgGISTRelayRequest, opts ...grpc.CallOption) (*MsgRelayResponse, error)
	StateRelays(ctx context.Context, in *MsgRelaysRequest, opts ...grpc.CallOption) (*MsgRelaysResponse, error)
	GISTRelays(ctx context.Context, in *MsgRelaysRequest, opts ...grpc.CallOption) (*MsgRelaysResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) StateRelay(ctx context.Context, in *MsgStateRelayRequest, opts ...grpc.CallOption) (*MsgRelayResponse, error) {
	out := new(MsgRelayResponse)
	err := c.cc.Invoke(ctx, "/Service/StateRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GistRelay(ctx context.Context, in *MsgGISTRelayRequest, opts ...grpc.CallOption) (*MsgRelayResponse, error) {
	out := new(MsgRelayResponse)
	err := c.cc.Invoke(ctx, "/Service/GistRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) StateRelays(ctx context.Context, in *MsgRelaysRequest, opts ...grpc.CallOption) (*MsgRelaysResponse, error) {
	out := new(MsgRelaysResponse)
	err := c.cc.Invoke(ctx, "/Service/StateRelays", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GISTRelays(ctx context.Context, in *MsgRelaysRequest, opts ...grpc.CallOption) (*MsgRelaysResponse, error) {
	out := new(MsgRelaysResponse)
	err := c.cc.Invoke(ctx, "/Service/GISTRelays", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations should embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	StateRelay(context.Context, *MsgStateRelayRequest) (*MsgRelayResponse, error)
	GistRelay(context.Context, *MsgGISTRelayRequest) (*MsgRelayResponse, error)
	StateRelays(context.Context, *MsgRelaysRequest) (*MsgRelaysResponse, error)
	GISTRelays(context.Context, *MsgRelaysRequest) (*MsgRelaysResponse, error)
}

// UnimplementedServiceServer should be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) StateRelay(context.Context, *MsgStateRelayRequest) (*MsgRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StateRelay not implemented")
}
func (UnimplementedServiceServer) GistRelay(context.Context, *MsgGISTRelayRequest) (*MsgRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GistRelay not implemented")
}
func (UnimplementedServiceServer) StateRelays(context.Context, *MsgRelaysRequest) (*MsgRelaysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StateRelays not implemented")
}
func (UnimplementedServiceServer) GISTRelays(context.Context, *MsgRelaysRequest) (*MsgRelaysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GISTRelays not implemented")
}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_StateRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStateRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).StateRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Service/StateRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).StateRelay(ctx, req.(*MsgStateRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GistRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGISTRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GistRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Service/GistRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GistRelay(ctx, req.(*MsgGISTRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_StateRelays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRelaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).StateRelays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Service/StateRelays",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).StateRelays(ctx, req.(*MsgRelaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GISTRelays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRelaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GISTRelays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Service/GISTRelays",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GISTRelays(ctx, req.(*MsgRelaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StateRelay",
			Handler:    _Service_StateRelay_Handler,
		},
		{
			MethodName: "GistRelay",
			Handler:    _Service_GistRelay_Handler,
		},
		{
			MethodName: "StateRelays",
			Handler:    _Service_StateRelays_Handler,
		},
		{
			MethodName: "GISTRelays",
			Handler:    _Service_GISTRelays_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
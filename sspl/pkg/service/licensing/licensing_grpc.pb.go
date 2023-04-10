// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: service/licensing/licensing.proto

package licensing

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

// LicensingClient is the client API for Licensing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LicensingClient interface {
	// Activate activates a Mutagen Pro license.
	Activate(ctx context.Context, in *ActivateRequest, opts ...grpc.CallOption) (*ActivateResponse, error)
	// Status returns Mutagen Pro license status information.
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	// Deactivate deactivates a Mutagen Pro license.
	Deactivate(ctx context.Context, in *DeactivateRequest, opts ...grpc.CallOption) (*DeactivateResponse, error)
}

type licensingClient struct {
	cc grpc.ClientConnInterface
}

func NewLicensingClient(cc grpc.ClientConnInterface) LicensingClient {
	return &licensingClient{cc}
}

func (c *licensingClient) Activate(ctx context.Context, in *ActivateRequest, opts ...grpc.CallOption) (*ActivateResponse, error) {
	out := new(ActivateResponse)
	err := c.cc.Invoke(ctx, "/licensing.Licensing/Activate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *licensingClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/licensing.Licensing/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *licensingClient) Deactivate(ctx context.Context, in *DeactivateRequest, opts ...grpc.CallOption) (*DeactivateResponse, error) {
	out := new(DeactivateResponse)
	err := c.cc.Invoke(ctx, "/licensing.Licensing/Deactivate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LicensingServer is the server API for Licensing service.
// All implementations must embed UnimplementedLicensingServer
// for forward compatibility
type LicensingServer interface {
	// Activate activates a Mutagen Pro license.
	Activate(context.Context, *ActivateRequest) (*ActivateResponse, error)
	// Status returns Mutagen Pro license status information.
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	// Deactivate deactivates a Mutagen Pro license.
	Deactivate(context.Context, *DeactivateRequest) (*DeactivateResponse, error)
	mustEmbedUnimplementedLicensingServer()
}

// UnimplementedLicensingServer must be embedded to have forward compatible implementations.
type UnimplementedLicensingServer struct {
}

func (UnimplementedLicensingServer) Activate(context.Context, *ActivateRequest) (*ActivateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}
func (UnimplementedLicensingServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedLicensingServer) Deactivate(context.Context, *DeactivateRequest) (*DeactivateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deactivate not implemented")
}
func (UnimplementedLicensingServer) mustEmbedUnimplementedLicensingServer() {}

// UnsafeLicensingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LicensingServer will
// result in compilation errors.
type UnsafeLicensingServer interface {
	mustEmbedUnimplementedLicensingServer()
}

func RegisterLicensingServer(s grpc.ServiceRegistrar, srv LicensingServer) {
	s.RegisterService(&Licensing_ServiceDesc, srv)
}

func _Licensing_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicensingServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/licensing.Licensing/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicensingServer).Activate(ctx, req.(*ActivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Licensing_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicensingServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/licensing.Licensing/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicensingServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Licensing_Deactivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicensingServer).Deactivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/licensing.Licensing/Deactivate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicensingServer).Deactivate(ctx, req.(*DeactivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Licensing_ServiceDesc is the grpc.ServiceDesc for Licensing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Licensing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "licensing.Licensing",
	HandlerType: (*LicensingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Activate",
			Handler:    _Licensing_Activate_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Licensing_Status_Handler,
		},
		{
			MethodName: "Deactivate",
			Handler:    _Licensing_Deactivate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/licensing/licensing.proto",
}
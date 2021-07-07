// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PaymantServiceClient is the client API for PaymantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymantServiceClient interface {
	Paymant(ctx context.Context, in *PaymantRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type paymantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymantServiceClient(cc grpc.ClientConnInterface) PaymantServiceClient {
	return &paymantServiceClient{cc}
}

func (c *paymantServiceClient) Paymant(ctx context.Context, in *PaymantRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/paymant.PaymantService/Paymant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymantServiceServer is the server API for PaymantService service.
// All implementations must embed UnimplementedPaymantServiceServer
// for forward compatibility
type PaymantServiceServer interface {
	Paymant(context.Context, *PaymantRequest) (*empty.Empty, error)
	mustEmbedUnimplementedPaymantServiceServer()
}

// UnimplementedPaymantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaymantServiceServer struct {
}

func (UnimplementedPaymantServiceServer) Paymant(context.Context, *PaymantRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Paymant not implemented")
}
func (UnimplementedPaymantServiceServer) mustEmbedUnimplementedPaymantServiceServer() {}

// UnsafePaymantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymantServiceServer will
// result in compilation errors.
type UnsafePaymantServiceServer interface {
	mustEmbedUnimplementedPaymantServiceServer()
}

func RegisterPaymantServiceServer(s grpc.ServiceRegistrar, srv PaymantServiceServer) {
	s.RegisterService(&PaymantService_ServiceDesc, srv)
}

func _PaymantService_Paymant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymantServiceServer).Paymant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paymant.PaymantService/Paymant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymantServiceServer).Paymant(ctx, req.(*PaymantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymantService_ServiceDesc is the grpc.ServiceDesc for PaymantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "paymant.PaymantService",
	HandlerType: (*PaymantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Paymant",
			Handler:    _PaymantService_Paymant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protofile/paymant.proto",
}

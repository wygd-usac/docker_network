// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package calculadora

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

// CalculadoraServiceClient is the client API for CalculadoraService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculadoraServiceClient interface {
	Operar(ctx context.Context, in *Operacion, opts ...grpc.CallOption) (*Resultado, error)
}

type calculadoraServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculadoraServiceClient(cc grpc.ClientConnInterface) CalculadoraServiceClient {
	return &calculadoraServiceClient{cc}
}

func (c *calculadoraServiceClient) Operar(ctx context.Context, in *Operacion, opts ...grpc.CallOption) (*Resultado, error) {
	out := new(Resultado)
	err := c.cc.Invoke(ctx, "/calculadora.CalculadoraService/Operar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculadoraServiceServer is the server API for CalculadoraService service.
// All implementations must embed UnimplementedCalculadoraServiceServer
// for forward compatibility
type CalculadoraServiceServer interface {
	Operar(context.Context, *Operacion) (*Resultado, error)
	mustEmbedUnimplementedCalculadoraServiceServer()
}

// UnimplementedCalculadoraServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculadoraServiceServer struct {
}

func (UnimplementedCalculadoraServiceServer) Operar(context.Context, *Operacion) (*Resultado, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Operar not implemented")
}
func (UnimplementedCalculadoraServiceServer) mustEmbedUnimplementedCalculadoraServiceServer() {}

// UnsafeCalculadoraServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculadoraServiceServer will
// result in compilation errors.
type UnsafeCalculadoraServiceServer interface {
	mustEmbedUnimplementedCalculadoraServiceServer()
}

func RegisterCalculadoraServiceServer(s grpc.ServiceRegistrar, srv CalculadoraServiceServer) {
	s.RegisterService(&CalculadoraService_ServiceDesc, srv)
}

func _CalculadoraService_Operar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Operacion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculadoraServiceServer).Operar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculadora.CalculadoraService/Operar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculadoraServiceServer).Operar(ctx, req.(*Operacion))
	}
	return interceptor(ctx, in, info, handler)
}

// CalculadoraService_ServiceDesc is the grpc.ServiceDesc for CalculadoraService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculadoraService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculadora.CalculadoraService",
	HandlerType: (*CalculadoraServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Operar",
			Handler:    _CalculadoraService_Operar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculadora.proto",
}

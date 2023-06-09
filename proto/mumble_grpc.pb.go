// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/mumble.proto

package proto

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

const (
	MumbleService_Murmer_FullMethodName = "/proto.MumbleService/Murmer"
)

// MumbleServiceClient is the client API for MumbleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MumbleServiceClient interface {
	Murmer(ctx context.Context, in *MurmerRequest, opts ...grpc.CallOption) (*MurmerResponse, error)
}

type mumbleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMumbleServiceClient(cc grpc.ClientConnInterface) MumbleServiceClient {
	return &mumbleServiceClient{cc}
}

func (c *mumbleServiceClient) Murmer(ctx context.Context, in *MurmerRequest, opts ...grpc.CallOption) (*MurmerResponse, error) {
	out := new(MurmerResponse)
	err := c.cc.Invoke(ctx, MumbleService_Murmer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MumbleServiceServer is the server API for MumbleService service.
// All implementations must embed UnimplementedMumbleServiceServer
// for forward compatibility
type MumbleServiceServer interface {
	Murmer(context.Context, *MurmerRequest) (*MurmerResponse, error)
	mustEmbedUnimplementedMumbleServiceServer()
}

// UnimplementedMumbleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMumbleServiceServer struct {
}

func (UnimplementedMumbleServiceServer) Murmer(context.Context, *MurmerRequest) (*MurmerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Murmer not implemented")
}
func (UnimplementedMumbleServiceServer) mustEmbedUnimplementedMumbleServiceServer() {}

// UnsafeMumbleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MumbleServiceServer will
// result in compilation errors.
type UnsafeMumbleServiceServer interface {
	mustEmbedUnimplementedMumbleServiceServer()
}

func RegisterMumbleServiceServer(s grpc.ServiceRegistrar, srv MumbleServiceServer) {
	s.RegisterService(&MumbleService_ServiceDesc, srv)
}

func _MumbleService_Murmer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MurmerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MumbleServiceServer).Murmer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MumbleService_Murmer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MumbleServiceServer).Murmer(ctx, req.(*MurmerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MumbleService_ServiceDesc is the grpc.ServiceDesc for MumbleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MumbleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MumbleService",
	HandlerType: (*MumbleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Murmer",
			Handler:    _MumbleService_Murmer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/mumble.proto",
}

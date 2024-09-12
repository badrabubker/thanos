// Copyright (c) The Thanos Authors.
// Licensed under the Apache License 2.0.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.0
// source: rules/rulespb/rpc.proto

package rulespb

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Rules_Rules_FullMethodName = "/thanos.Rules/Rules"
)

// RulesClient is the client API for Rules service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// / Rules represents API that is responsible for gathering rules and their statuses.
type RulesClient interface {
	// / Rules has info for all rules.
	// / Returned rules are expected to include external labels.
	Rules(ctx context.Context, in *RulesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[RulesResponse], error)
}

type rulesClient struct {
	cc grpc.ClientConnInterface
}

func NewRulesClient(cc grpc.ClientConnInterface) RulesClient {
	return &rulesClient{cc}
}

func (c *rulesClient) Rules(ctx context.Context, in *RulesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[RulesResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Rules_ServiceDesc.Streams[0], Rules_Rules_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[RulesRequest, RulesResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Rules_RulesClient = grpc.ServerStreamingClient[RulesResponse]

// RulesServer is the server API for Rules service.
// All implementations must embed UnimplementedRulesServer
// for forward compatibility.
//
// / Rules represents API that is responsible for gathering rules and their statuses.
type RulesServer interface {
	// / Rules has info for all rules.
	// / Returned rules are expected to include external labels.
	Rules(*RulesRequest, grpc.ServerStreamingServer[RulesResponse]) error
	mustEmbedUnimplementedRulesServer()
}

// UnimplementedRulesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRulesServer struct{}

func (UnimplementedRulesServer) Rules(*RulesRequest, grpc.ServerStreamingServer[RulesResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Rules not implemented")
}
func (UnimplementedRulesServer) mustEmbedUnimplementedRulesServer() {}
func (UnimplementedRulesServer) testEmbeddedByValue()               {}

// UnsafeRulesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RulesServer will
// result in compilation errors.
type UnsafeRulesServer interface {
	mustEmbedUnimplementedRulesServer()
}

func RegisterRulesServer(s grpc.ServiceRegistrar, srv RulesServer) {
	// If the following call pancis, it indicates UnimplementedRulesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Rules_ServiceDesc, srv)
}

func _Rules_Rules_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RulesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RulesServer).Rules(m, &grpc.GenericServerStream[RulesRequest, RulesResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Rules_RulesServer = grpc.ServerStreamingServer[RulesResponse]

// Rules_ServiceDesc is the grpc.ServiceDesc for Rules service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rules_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thanos.Rules",
	HandlerType: (*RulesServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Rules",
			Handler:       _Rules_Rules_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rules/rulespb/rpc.proto",
}

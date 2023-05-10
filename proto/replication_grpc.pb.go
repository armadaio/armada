// Copyright JAMF Software, LLC

//
// Regatta replication protobuffer specification
//

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: replication.proto

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
	Metadata_Get_FullMethodName = "/replication.v1.Metadata/Get"
)

// MetadataClient is the client API for Metadata service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetadataClient interface {
	Get(ctx context.Context, in *MetadataRequest, opts ...grpc.CallOption) (*MetadataResponse, error)
}

type metadataClient struct {
	cc grpc.ClientConnInterface
}

func NewMetadataClient(cc grpc.ClientConnInterface) MetadataClient {
	return &metadataClient{cc}
}

func (c *metadataClient) Get(ctx context.Context, in *MetadataRequest, opts ...grpc.CallOption) (*MetadataResponse, error) {
	out := new(MetadataResponse)
	err := c.cc.Invoke(ctx, Metadata_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataServer is the server API for Metadata service.
// All implementations must embed UnimplementedMetadataServer
// for forward compatibility
type MetadataServer interface {
	Get(context.Context, *MetadataRequest) (*MetadataResponse, error)
	mustEmbedUnimplementedMetadataServer()
}

// UnimplementedMetadataServer must be embedded to have forward compatible implementations.
type UnimplementedMetadataServer struct {
}

func (UnimplementedMetadataServer) Get(context.Context, *MetadataRequest) (*MetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMetadataServer) mustEmbedUnimplementedMetadataServer() {}

// UnsafeMetadataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetadataServer will
// result in compilation errors.
type UnsafeMetadataServer interface {
	mustEmbedUnimplementedMetadataServer()
}

func RegisterMetadataServer(s grpc.ServiceRegistrar, srv MetadataServer) {
	s.RegisterService(&Metadata_ServiceDesc, srv)
}

func _Metadata_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Metadata_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).Get(ctx, req.(*MetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Metadata_ServiceDesc is the grpc.ServiceDesc for Metadata service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Metadata_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "replication.v1.Metadata",
	HandlerType: (*MetadataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Metadata_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "replication.proto",
}

const (
	Snapshot_Stream_FullMethodName = "/replication.v1.Snapshot/Stream"
)

// SnapshotClient is the client API for Snapshot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SnapshotClient interface {
	Stream(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (Snapshot_StreamClient, error)
}

type snapshotClient struct {
	cc grpc.ClientConnInterface
}

func NewSnapshotClient(cc grpc.ClientConnInterface) SnapshotClient {
	return &snapshotClient{cc}
}

func (c *snapshotClient) Stream(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (Snapshot_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Snapshot_ServiceDesc.Streams[0], Snapshot_Stream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &snapshotStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Snapshot_StreamClient interface {
	Recv() (*SnapshotChunk, error)
	grpc.ClientStream
}

type snapshotStreamClient struct {
	grpc.ClientStream
}

func (x *snapshotStreamClient) Recv() (*SnapshotChunk, error) {
	m := new(SnapshotChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SnapshotServer is the server API for Snapshot service.
// All implementations must embed UnimplementedSnapshotServer
// for forward compatibility
type SnapshotServer interface {
	Stream(*SnapshotRequest, Snapshot_StreamServer) error
	mustEmbedUnimplementedSnapshotServer()
}

// UnimplementedSnapshotServer must be embedded to have forward compatible implementations.
type UnimplementedSnapshotServer struct {
}

func (UnimplementedSnapshotServer) Stream(*SnapshotRequest, Snapshot_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedSnapshotServer) mustEmbedUnimplementedSnapshotServer() {}

// UnsafeSnapshotServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SnapshotServer will
// result in compilation errors.
type UnsafeSnapshotServer interface {
	mustEmbedUnimplementedSnapshotServer()
}

func RegisterSnapshotServer(s grpc.ServiceRegistrar, srv SnapshotServer) {
	s.RegisterService(&Snapshot_ServiceDesc, srv)
}

func _Snapshot_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SnapshotRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SnapshotServer).Stream(m, &snapshotStreamServer{stream})
}

type Snapshot_StreamServer interface {
	Send(*SnapshotChunk) error
	grpc.ServerStream
}

type snapshotStreamServer struct {
	grpc.ServerStream
}

func (x *snapshotStreamServer) Send(m *SnapshotChunk) error {
	return x.ServerStream.SendMsg(m)
}

// Snapshot_ServiceDesc is the grpc.ServiceDesc for Snapshot service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Snapshot_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "replication.v1.Snapshot",
	HandlerType: (*SnapshotServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Snapshot_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "replication.proto",
}

const (
	Log_Replicate_FullMethodName = "/replication.v1.Log/Replicate"
)

// LogClient is the client API for Log service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogClient interface {
	// Replicate is method to ask for data of specified table from the specified index.
	Replicate(ctx context.Context, in *ReplicateRequest, opts ...grpc.CallOption) (Log_ReplicateClient, error)
}

type logClient struct {
	cc grpc.ClientConnInterface
}

func NewLogClient(cc grpc.ClientConnInterface) LogClient {
	return &logClient{cc}
}

func (c *logClient) Replicate(ctx context.Context, in *ReplicateRequest, opts ...grpc.CallOption) (Log_ReplicateClient, error) {
	stream, err := c.cc.NewStream(ctx, &Log_ServiceDesc.Streams[0], Log_Replicate_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &logReplicateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Log_ReplicateClient interface {
	Recv() (*ReplicateResponse, error)
	grpc.ClientStream
}

type logReplicateClient struct {
	grpc.ClientStream
}

func (x *logReplicateClient) Recv() (*ReplicateResponse, error) {
	m := new(ReplicateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogServer is the server API for Log service.
// All implementations must embed UnimplementedLogServer
// for forward compatibility
type LogServer interface {
	// Replicate is method to ask for data of specified table from the specified index.
	Replicate(*ReplicateRequest, Log_ReplicateServer) error
	mustEmbedUnimplementedLogServer()
}

// UnimplementedLogServer must be embedded to have forward compatible implementations.
type UnimplementedLogServer struct {
}

func (UnimplementedLogServer) Replicate(*ReplicateRequest, Log_ReplicateServer) error {
	return status.Errorf(codes.Unimplemented, "method Replicate not implemented")
}
func (UnimplementedLogServer) mustEmbedUnimplementedLogServer() {}

// UnsafeLogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServer will
// result in compilation errors.
type UnsafeLogServer interface {
	mustEmbedUnimplementedLogServer()
}

func RegisterLogServer(s grpc.ServiceRegistrar, srv LogServer) {
	s.RegisterService(&Log_ServiceDesc, srv)
}

func _Log_Replicate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReplicateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogServer).Replicate(m, &logReplicateServer{stream})
}

type Log_ReplicateServer interface {
	Send(*ReplicateResponse) error
	grpc.ServerStream
}

type logReplicateServer struct {
	grpc.ServerStream
}

func (x *logReplicateServer) Send(m *ReplicateResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Log_ServiceDesc is the grpc.ServiceDesc for Log service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Log_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "replication.v1.Log",
	HandlerType: (*LogServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Replicate",
			Handler:       _Log_Replicate_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "replication.proto",
}

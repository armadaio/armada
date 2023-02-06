// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: maintenance.proto

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

// MaintenanceClient is the client API for Maintenance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MaintenanceClient interface {
	Backup(ctx context.Context, in *BackupRequest, opts ...grpc.CallOption) (Maintenance_BackupClient, error)
	Restore(ctx context.Context, opts ...grpc.CallOption) (Maintenance_RestoreClient, error)
	Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*ResetResponse, error)
}

type maintenanceClient struct {
	cc grpc.ClientConnInterface
}

func NewMaintenanceClient(cc grpc.ClientConnInterface) MaintenanceClient {
	return &maintenanceClient{cc}
}

func (c *maintenanceClient) Backup(ctx context.Context, in *BackupRequest, opts ...grpc.CallOption) (Maintenance_BackupClient, error) {
	stream, err := c.cc.NewStream(ctx, &Maintenance_ServiceDesc.Streams[0], "/maintenance.v1.Maintenance/Backup", opts...)
	if err != nil {
		return nil, err
	}
	x := &maintenanceBackupClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Maintenance_BackupClient interface {
	Recv() (*SnapshotChunk, error)
	grpc.ClientStream
}

type maintenanceBackupClient struct {
	grpc.ClientStream
}

func (x *maintenanceBackupClient) Recv() (*SnapshotChunk, error) {
	m := new(SnapshotChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *maintenanceClient) Restore(ctx context.Context, opts ...grpc.CallOption) (Maintenance_RestoreClient, error) {
	stream, err := c.cc.NewStream(ctx, &Maintenance_ServiceDesc.Streams[1], "/maintenance.v1.Maintenance/Restore", opts...)
	if err != nil {
		return nil, err
	}
	x := &maintenanceRestoreClient{stream}
	return x, nil
}

type Maintenance_RestoreClient interface {
	Send(*RestoreMessage) error
	CloseAndRecv() (*RestoreResponse, error)
	grpc.ClientStream
}

type maintenanceRestoreClient struct {
	grpc.ClientStream
}

func (x *maintenanceRestoreClient) Send(m *RestoreMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *maintenanceRestoreClient) CloseAndRecv() (*RestoreResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RestoreResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *maintenanceClient) Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*ResetResponse, error) {
	out := new(ResetResponse)
	err := c.cc.Invoke(ctx, "/maintenance.v1.Maintenance/Reset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MaintenanceServer is the server API for Maintenance service.
// All implementations must embed UnimplementedMaintenanceServer
// for forward compatibility
type MaintenanceServer interface {
	Backup(*BackupRequest, Maintenance_BackupServer) error
	Restore(Maintenance_RestoreServer) error
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)
	mustEmbedUnimplementedMaintenanceServer()
}

// UnimplementedMaintenanceServer must be embedded to have forward compatible implementations.
type UnimplementedMaintenanceServer struct {
}

func (UnimplementedMaintenanceServer) Backup(*BackupRequest, Maintenance_BackupServer) error {
	return status.Errorf(codes.Unimplemented, "method Backup not implemented")
}
func (UnimplementedMaintenanceServer) Restore(Maintenance_RestoreServer) error {
	return status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedMaintenanceServer) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reset not implemented")
}
func (UnimplementedMaintenanceServer) mustEmbedUnimplementedMaintenanceServer() {}

// UnsafeMaintenanceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MaintenanceServer will
// result in compilation errors.
type UnsafeMaintenanceServer interface {
	mustEmbedUnimplementedMaintenanceServer()
}

func RegisterMaintenanceServer(s grpc.ServiceRegistrar, srv MaintenanceServer) {
	s.RegisterService(&Maintenance_ServiceDesc, srv)
}

func _Maintenance_Backup_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BackupRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MaintenanceServer).Backup(m, &maintenanceBackupServer{stream})
}

type Maintenance_BackupServer interface {
	Send(*SnapshotChunk) error
	grpc.ServerStream
}

type maintenanceBackupServer struct {
	grpc.ServerStream
}

func (x *maintenanceBackupServer) Send(m *SnapshotChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _Maintenance_Restore_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MaintenanceServer).Restore(&maintenanceRestoreServer{stream})
}

type Maintenance_RestoreServer interface {
	SendAndClose(*RestoreResponse) error
	Recv() (*RestoreMessage, error)
	grpc.ServerStream
}

type maintenanceRestoreServer struct {
	grpc.ServerStream
}

func (x *maintenanceRestoreServer) SendAndClose(m *RestoreResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *maintenanceRestoreServer) Recv() (*RestoreMessage, error) {
	m := new(RestoreMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Maintenance_Reset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenanceServer).Reset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/maintenance.v1.Maintenance/Reset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenanceServer).Reset(ctx, req.(*ResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Maintenance_ServiceDesc is the grpc.ServiceDesc for Maintenance service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Maintenance_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "maintenance.v1.Maintenance",
	HandlerType: (*MaintenanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reset",
			Handler:    _Maintenance_Reset_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Backup",
			Handler:       _Maintenance_Backup_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Restore",
			Handler:       _Maintenance_Restore_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "maintenance.proto",
}

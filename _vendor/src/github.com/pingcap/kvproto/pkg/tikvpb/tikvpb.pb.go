// Code generated by protoc-gen-gogo.
// source: tikvpb.proto
// DO NOT EDIT!

/*
	Package tikvpb is a generated protocol buffer package.

	It is generated from these files:
		tikvpb.proto

	It has these top-level messages:
*/
package tikvpb

import (
	"fmt"
	"math"

	proto "github.com/golang/protobuf/proto"

	coprocessor "github.com/pingcap/kvproto/pkg/coprocessor"

	kvrpcpb "github.com/pingcap/kvproto/pkg/kvrpcpb"

	raft_serverpb "github.com/pingcap/kvproto/pkg/raft_serverpb"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Tikv service

type TikvClient interface {
	// KV commands with mvcc/txn supported.
	KvGet(ctx context.Context, in *kvrpcpb.GetRequest, opts ...grpc.CallOption) (*kvrpcpb.GetResponse, error)
	KvScan(ctx context.Context, in *kvrpcpb.ScanRequest, opts ...grpc.CallOption) (*kvrpcpb.ScanResponse, error)
	KvPrewrite(ctx context.Context, in *kvrpcpb.PrewriteRequest, opts ...grpc.CallOption) (*kvrpcpb.PrewriteResponse, error)
	KvCommit(ctx context.Context, in *kvrpcpb.CommitRequest, opts ...grpc.CallOption) (*kvrpcpb.CommitResponse, error)
	KvImport(ctx context.Context, in *kvrpcpb.ImportRequest, opts ...grpc.CallOption) (*kvrpcpb.ImportResponse, error)
	KvCleanup(ctx context.Context, in *kvrpcpb.CleanupRequest, opts ...grpc.CallOption) (*kvrpcpb.CleanupResponse, error)
	KvBatchGet(ctx context.Context, in *kvrpcpb.BatchGetRequest, opts ...grpc.CallOption) (*kvrpcpb.BatchGetResponse, error)
	KvBatchRollback(ctx context.Context, in *kvrpcpb.BatchRollbackRequest, opts ...grpc.CallOption) (*kvrpcpb.BatchRollbackResponse, error)
	KvScanLock(ctx context.Context, in *kvrpcpb.ScanLockRequest, opts ...grpc.CallOption) (*kvrpcpb.ScanLockResponse, error)
	KvResolveLock(ctx context.Context, in *kvrpcpb.ResolveLockRequest, opts ...grpc.CallOption) (*kvrpcpb.ResolveLockResponse, error)
	KvGC(ctx context.Context, in *kvrpcpb.GCRequest, opts ...grpc.CallOption) (*kvrpcpb.GCResponse, error)
	// RawKV commands.
	RawGet(ctx context.Context, in *kvrpcpb.RawGetRequest, opts ...grpc.CallOption) (*kvrpcpb.RawGetResponse, error)
	RawPut(ctx context.Context, in *kvrpcpb.RawPutRequest, opts ...grpc.CallOption) (*kvrpcpb.RawPutResponse, error)
	RawDelete(ctx context.Context, in *kvrpcpb.RawDeleteRequest, opts ...grpc.CallOption) (*kvrpcpb.RawDeleteResponse, error)
	RawScan(ctx context.Context, in *kvrpcpb.RawScanRequest, opts ...grpc.CallOption) (*kvrpcpb.RawScanResponse, error)
	// SQL push down commands.
	Coprocessor(ctx context.Context, in *coprocessor.Request, opts ...grpc.CallOption) (*coprocessor.Response, error)
	// Raft commands (tikv <-> tikv).
	Raft(ctx context.Context, opts ...grpc.CallOption) (Tikv_RaftClient, error)
	Snapshot(ctx context.Context, opts ...grpc.CallOption) (Tikv_SnapshotClient, error)
	// transaction debugger commands.
	MvccGetByKey(ctx context.Context, in *kvrpcpb.MvccGetByKeyRequest, opts ...grpc.CallOption) (*kvrpcpb.MvccGetByKeyResponse, error)
	MvccGetByStartTs(ctx context.Context, in *kvrpcpb.MvccGetByStartTsRequest, opts ...grpc.CallOption) (*kvrpcpb.MvccGetByStartTsResponse, error)
}

type tikvClient struct {
	cc *grpc.ClientConn
}

func NewTikvClient(cc *grpc.ClientConn) TikvClient {
	return &tikvClient{cc}
}

func (c *tikvClient) KvGet(ctx context.Context, in *kvrpcpb.GetRequest, opts ...grpc.CallOption) (*kvrpcpb.GetResponse, error) {
	out := new(kvrpcpb.GetResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvScan(ctx context.Context, in *kvrpcpb.ScanRequest, opts ...grpc.CallOption) (*kvrpcpb.ScanResponse, error) {
	out := new(kvrpcpb.ScanResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvScan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvPrewrite(ctx context.Context, in *kvrpcpb.PrewriteRequest, opts ...grpc.CallOption) (*kvrpcpb.PrewriteResponse, error) {
	out := new(kvrpcpb.PrewriteResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvPrewrite", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvCommit(ctx context.Context, in *kvrpcpb.CommitRequest, opts ...grpc.CallOption) (*kvrpcpb.CommitResponse, error) {
	out := new(kvrpcpb.CommitResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvCommit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvImport(ctx context.Context, in *kvrpcpb.ImportRequest, opts ...grpc.CallOption) (*kvrpcpb.ImportResponse, error) {
	out := new(kvrpcpb.ImportResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvImport", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvCleanup(ctx context.Context, in *kvrpcpb.CleanupRequest, opts ...grpc.CallOption) (*kvrpcpb.CleanupResponse, error) {
	out := new(kvrpcpb.CleanupResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvCleanup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvBatchGet(ctx context.Context, in *kvrpcpb.BatchGetRequest, opts ...grpc.CallOption) (*kvrpcpb.BatchGetResponse, error) {
	out := new(kvrpcpb.BatchGetResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvBatchGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvBatchRollback(ctx context.Context, in *kvrpcpb.BatchRollbackRequest, opts ...grpc.CallOption) (*kvrpcpb.BatchRollbackResponse, error) {
	out := new(kvrpcpb.BatchRollbackResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvBatchRollback", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvScanLock(ctx context.Context, in *kvrpcpb.ScanLockRequest, opts ...grpc.CallOption) (*kvrpcpb.ScanLockResponse, error) {
	out := new(kvrpcpb.ScanLockResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvScanLock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvResolveLock(ctx context.Context, in *kvrpcpb.ResolveLockRequest, opts ...grpc.CallOption) (*kvrpcpb.ResolveLockResponse, error) {
	out := new(kvrpcpb.ResolveLockResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvResolveLock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvGC(ctx context.Context, in *kvrpcpb.GCRequest, opts ...grpc.CallOption) (*kvrpcpb.GCResponse, error) {
	out := new(kvrpcpb.GCResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/KvGC", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) RawGet(ctx context.Context, in *kvrpcpb.RawGetRequest, opts ...grpc.CallOption) (*kvrpcpb.RawGetResponse, error) {
	out := new(kvrpcpb.RawGetResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/RawGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) RawPut(ctx context.Context, in *kvrpcpb.RawPutRequest, opts ...grpc.CallOption) (*kvrpcpb.RawPutResponse, error) {
	out := new(kvrpcpb.RawPutResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/RawPut", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) RawDelete(ctx context.Context, in *kvrpcpb.RawDeleteRequest, opts ...grpc.CallOption) (*kvrpcpb.RawDeleteResponse, error) {
	out := new(kvrpcpb.RawDeleteResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/RawDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) RawScan(ctx context.Context, in *kvrpcpb.RawScanRequest, opts ...grpc.CallOption) (*kvrpcpb.RawScanResponse, error) {
	out := new(kvrpcpb.RawScanResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/RawScan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) Coprocessor(ctx context.Context, in *coprocessor.Request, opts ...grpc.CallOption) (*coprocessor.Response, error) {
	out := new(coprocessor.Response)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/Coprocessor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) Raft(ctx context.Context, opts ...grpc.CallOption) (Tikv_RaftClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Tikv_serviceDesc.Streams[0], c.cc, "/tikvpb.Tikv/Raft", opts...)
	if err != nil {
		return nil, err
	}
	x := &tikvRaftClient{stream}
	return x, nil
}

type Tikv_RaftClient interface {
	Send(*raft_serverpb.RaftMessage) error
	CloseAndRecv() (*raft_serverpb.Done, error)
	grpc.ClientStream
}

type tikvRaftClient struct {
	grpc.ClientStream
}

func (x *tikvRaftClient) Send(m *raft_serverpb.RaftMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tikvRaftClient) CloseAndRecv() (*raft_serverpb.Done, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(raft_serverpb.Done)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tikvClient) Snapshot(ctx context.Context, opts ...grpc.CallOption) (Tikv_SnapshotClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Tikv_serviceDesc.Streams[1], c.cc, "/tikvpb.Tikv/Snapshot", opts...)
	if err != nil {
		return nil, err
	}
	x := &tikvSnapshotClient{stream}
	return x, nil
}

type Tikv_SnapshotClient interface {
	Send(*raft_serverpb.SnapshotChunk) error
	CloseAndRecv() (*raft_serverpb.Done, error)
	grpc.ClientStream
}

type tikvSnapshotClient struct {
	grpc.ClientStream
}

func (x *tikvSnapshotClient) Send(m *raft_serverpb.SnapshotChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tikvSnapshotClient) CloseAndRecv() (*raft_serverpb.Done, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(raft_serverpb.Done)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tikvClient) MvccGetByKey(ctx context.Context, in *kvrpcpb.MvccGetByKeyRequest, opts ...grpc.CallOption) (*kvrpcpb.MvccGetByKeyResponse, error) {
	out := new(kvrpcpb.MvccGetByKeyResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/MvccGetByKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) MvccGetByStartTs(ctx context.Context, in *kvrpcpb.MvccGetByStartTsRequest, opts ...grpc.CallOption) (*kvrpcpb.MvccGetByStartTsResponse, error) {
	out := new(kvrpcpb.MvccGetByStartTsResponse)
	err := grpc.Invoke(ctx, "/tikvpb.Tikv/MvccGetByStartTs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Tikv service

type TikvServer interface {
	// KV commands with mvcc/txn supported.
	KvGet(context.Context, *kvrpcpb.GetRequest) (*kvrpcpb.GetResponse, error)
	KvScan(context.Context, *kvrpcpb.ScanRequest) (*kvrpcpb.ScanResponse, error)
	KvPrewrite(context.Context, *kvrpcpb.PrewriteRequest) (*kvrpcpb.PrewriteResponse, error)
	KvCommit(context.Context, *kvrpcpb.CommitRequest) (*kvrpcpb.CommitResponse, error)
	KvImport(context.Context, *kvrpcpb.ImportRequest) (*kvrpcpb.ImportResponse, error)
	KvCleanup(context.Context, *kvrpcpb.CleanupRequest) (*kvrpcpb.CleanupResponse, error)
	KvBatchGet(context.Context, *kvrpcpb.BatchGetRequest) (*kvrpcpb.BatchGetResponse, error)
	KvBatchRollback(context.Context, *kvrpcpb.BatchRollbackRequest) (*kvrpcpb.BatchRollbackResponse, error)
	KvScanLock(context.Context, *kvrpcpb.ScanLockRequest) (*kvrpcpb.ScanLockResponse, error)
	KvResolveLock(context.Context, *kvrpcpb.ResolveLockRequest) (*kvrpcpb.ResolveLockResponse, error)
	KvGC(context.Context, *kvrpcpb.GCRequest) (*kvrpcpb.GCResponse, error)
	// RawKV commands.
	RawGet(context.Context, *kvrpcpb.RawGetRequest) (*kvrpcpb.RawGetResponse, error)
	RawPut(context.Context, *kvrpcpb.RawPutRequest) (*kvrpcpb.RawPutResponse, error)
	RawDelete(context.Context, *kvrpcpb.RawDeleteRequest) (*kvrpcpb.RawDeleteResponse, error)
	RawScan(context.Context, *kvrpcpb.RawScanRequest) (*kvrpcpb.RawScanResponse, error)
	// SQL push down commands.
	Coprocessor(context.Context, *coprocessor.Request) (*coprocessor.Response, error)
	// Raft commands (tikv <-> tikv).
	Raft(Tikv_RaftServer) error
	Snapshot(Tikv_SnapshotServer) error
	// transaction debugger commands.
	MvccGetByKey(context.Context, *kvrpcpb.MvccGetByKeyRequest) (*kvrpcpb.MvccGetByKeyResponse, error)
	MvccGetByStartTs(context.Context, *kvrpcpb.MvccGetByStartTsRequest) (*kvrpcpb.MvccGetByStartTsResponse, error)
}

func RegisterTikvServer(s *grpc.Server, srv TikvServer) {
	s.RegisterService(&_Tikv_serviceDesc, srv)
}

func _Tikv_KvGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvGet(ctx, req.(*kvrpcpb.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.ScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvScan(ctx, req.(*kvrpcpb.ScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvPrewrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.PrewriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvPrewrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvPrewrite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvPrewrite(ctx, req.(*kvrpcpb.PrewriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.CommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvCommit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvCommit(ctx, req.(*kvrpcpb.CommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvImport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.ImportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvImport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvImport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvImport(ctx, req.(*kvrpcpb.ImportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvCleanup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.CleanupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvCleanup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvCleanup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvCleanup(ctx, req.(*kvrpcpb.CleanupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvBatchGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.BatchGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvBatchGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvBatchGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvBatchGet(ctx, req.(*kvrpcpb.BatchGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvBatchRollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.BatchRollbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvBatchRollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvBatchRollback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvBatchRollback(ctx, req.(*kvrpcpb.BatchRollbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvScanLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.ScanLockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvScanLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvScanLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvScanLock(ctx, req.(*kvrpcpb.ScanLockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvResolveLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.ResolveLockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvResolveLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvResolveLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvResolveLock(ctx, req.(*kvrpcpb.ResolveLockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvGC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.GCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvGC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvGC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvGC(ctx, req.(*kvrpcpb.GCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_RawGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).RawGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/RawGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).RawGet(ctx, req.(*kvrpcpb.RawGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_RawPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).RawPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/RawPut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).RawPut(ctx, req.(*kvrpcpb.RawPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_RawDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).RawDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/RawDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).RawDelete(ctx, req.(*kvrpcpb.RawDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_RawScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).RawScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/RawScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).RawScan(ctx, req.(*kvrpcpb.RawScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_Coprocessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(coprocessor.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).Coprocessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/Coprocessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).Coprocessor(ctx, req.(*coprocessor.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_Raft_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TikvServer).Raft(&tikvRaftServer{stream})
}

type Tikv_RaftServer interface {
	SendAndClose(*raft_serverpb.Done) error
	Recv() (*raft_serverpb.RaftMessage, error)
	grpc.ServerStream
}

type tikvRaftServer struct {
	grpc.ServerStream
}

func (x *tikvRaftServer) SendAndClose(m *raft_serverpb.Done) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tikvRaftServer) Recv() (*raft_serverpb.RaftMessage, error) {
	m := new(raft_serverpb.RaftMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Tikv_Snapshot_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TikvServer).Snapshot(&tikvSnapshotServer{stream})
}

type Tikv_SnapshotServer interface {
	SendAndClose(*raft_serverpb.Done) error
	Recv() (*raft_serverpb.SnapshotChunk, error)
	grpc.ServerStream
}

type tikvSnapshotServer struct {
	grpc.ServerStream
}

func (x *tikvSnapshotServer) SendAndClose(m *raft_serverpb.Done) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tikvSnapshotServer) Recv() (*raft_serverpb.SnapshotChunk, error) {
	m := new(raft_serverpb.SnapshotChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Tikv_MvccGetByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.MvccGetByKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).MvccGetByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/MvccGetByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).MvccGetByKey(ctx, req.(*kvrpcpb.MvccGetByKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_MvccGetByStartTs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.MvccGetByStartTsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).MvccGetByStartTs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/MvccGetByStartTs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).MvccGetByStartTs(ctx, req.(*kvrpcpb.MvccGetByStartTsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tikv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tikvpb.Tikv",
	HandlerType: (*TikvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KvGet",
			Handler:    _Tikv_KvGet_Handler,
		},
		{
			MethodName: "KvScan",
			Handler:    _Tikv_KvScan_Handler,
		},
		{
			MethodName: "KvPrewrite",
			Handler:    _Tikv_KvPrewrite_Handler,
		},
		{
			MethodName: "KvCommit",
			Handler:    _Tikv_KvCommit_Handler,
		},
		{
			MethodName: "KvImport",
			Handler:    _Tikv_KvImport_Handler,
		},
		{
			MethodName: "KvCleanup",
			Handler:    _Tikv_KvCleanup_Handler,
		},
		{
			MethodName: "KvBatchGet",
			Handler:    _Tikv_KvBatchGet_Handler,
		},
		{
			MethodName: "KvBatchRollback",
			Handler:    _Tikv_KvBatchRollback_Handler,
		},
		{
			MethodName: "KvScanLock",
			Handler:    _Tikv_KvScanLock_Handler,
		},
		{
			MethodName: "KvResolveLock",
			Handler:    _Tikv_KvResolveLock_Handler,
		},
		{
			MethodName: "KvGC",
			Handler:    _Tikv_KvGC_Handler,
		},
		{
			MethodName: "RawGet",
			Handler:    _Tikv_RawGet_Handler,
		},
		{
			MethodName: "RawPut",
			Handler:    _Tikv_RawPut_Handler,
		},
		{
			MethodName: "RawDelete",
			Handler:    _Tikv_RawDelete_Handler,
		},
		{
			MethodName: "RawScan",
			Handler:    _Tikv_RawScan_Handler,
		},
		{
			MethodName: "Coprocessor",
			Handler:    _Tikv_Coprocessor_Handler,
		},
		{
			MethodName: "MvccGetByKey",
			Handler:    _Tikv_MvccGetByKey_Handler,
		},
		{
			MethodName: "MvccGetByStartTs",
			Handler:    _Tikv_MvccGetByStartTs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Raft",
			Handler:       _Tikv_Raft_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Snapshot",
			Handler:       _Tikv_Snapshot_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "tikvpb.proto",
}

func init() { proto.RegisterFile("tikvpb.proto", fileDescriptorTikvpb) }

var fileDescriptorTikvpb = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0x5b, 0xa9, 0x94, 0xcd, 0x6c, 0x02, 0xdc, 0x02, 0x5d, 0xd8, 0x2a, 0xd8, 0x09, 0x71,
	0x08, 0xe2, 0x45, 0xe2, 0xc0, 0x8b, 0xa0, 0xa9, 0x54, 0xa1, 0xac, 0x52, 0x95, 0xee, 0xc2, 0x09,
	0xb9, 0xd6, 0xb3, 0xb6, 0x4a, 0x1b, 0x07, 0xdb, 0x75, 0xb5, 0x6f, 0xc2, 0xd7, 0xe1, 0xc6, 0x91,
	0x8f, 0x80, 0xca, 0x17, 0x41, 0x49, 0xf6, 0x38, 0x49, 0xd3, 0x72, 0x4a, 0xf2, 0xfb, 0xbf, 0xc4,
	0xb1, 0x63, 0x93, 0x23, 0x3d, 0x0f, 0x4d, 0x3c, 0x71, 0x63, 0x29, 0xb4, 0xa0, 0xcd, 0xec, 0xc9,
	0xb9, 0xcf, 0x45, 0x2c, 0x05, 0x07, 0xa5, 0x84, 0xcc, 0x24, 0xe7, 0x38, 0x34, 0x32, 0xe6, 0xe8,
	0x74, 0x5a, 0x92, 0x5d, 0xe9, 0x6f, 0x0a, 0xa4, 0x01, 0x69, 0x61, 0x7b, 0x2a, 0xa6, 0x22, 0xbd,
	0x7d, 0x91, 0xdc, 0x65, 0xf4, 0xd5, 0x4f, 0x42, 0x1a, 0x97, 0xf3, 0xd0, 0xd0, 0x37, 0xe4, 0x96,
	0x6f, 0x06, 0xa0, 0x69, 0xcb, 0xc5, 0xb2, 0x01, 0xe8, 0x00, 0xbe, 0xaf, 0x40, 0x69, 0xa7, 0x5d,
	0x86, 0x2a, 0x16, 0x91, 0x82, 0xf3, 0x1a, 0x7d, 0x4b, 0x9a, 0xbe, 0x19, 0x73, 0x16, 0xd1, 0xdc,
	0x91, 0x3c, 0x62, 0xee, 0xc1, 0x16, 0xb5, 0x41, 0x8f, 0x10, 0xdf, 0x8c, 0x24, 0xac, 0xe5, 0x5c,
	0x03, 0xed, 0x58, 0x1b, 0x22, 0x2c, 0x38, 0xd9, 0xa1, 0xd8, 0x92, 0x0f, 0xe4, 0xc0, 0x37, 0x9e,
	0x58, 0x2e, 0xe7, 0x9a, 0x3e, 0xb4, 0xc6, 0x0c, 0x60, 0xc1, 0xa3, 0x0a, 0x2f, 0xc7, 0xbf, 0x2c,
	0x63, 0x21, 0x8b, 0xf1, 0x0c, 0x54, 0xe3, 0xc8, 0x6d, 0xfc, 0x13, 0x39, 0xf4, 0x8d, 0xb7, 0x00,
	0x16, 0xad, 0x62, 0x5a, 0x78, 0x4d, 0x46, 0xb0, 0xa0, 0x53, 0x15, 0xca, 0x93, 0xd0, 0x63, 0x9a,
	0xcf, 0x92, 0x89, 0xcf, 0x9d, 0x88, 0xaa, 0x93, 0x90, 0x2b, 0xb6, 0x24, 0x20, 0x77, 0x6f, 0x4a,
	0x02, 0xb1, 0x58, 0x4c, 0x18, 0x0f, 0xe9, 0x59, 0xd9, 0x8f, 0x1c, 0xeb, 0xba, 0xfb, 0xe4, 0xf2,
	0xc0, 0x92, 0x15, 0xbb, 0x10, 0x3c, 0x2c, 0x0c, 0x0c, 0x51, 0x75, 0x60, 0xb9, 0x62, 0x4b, 0x2e,
	0xc8, 0xb1, 0x6f, 0x02, 0x50, 0x62, 0x61, 0x20, 0xed, 0x79, 0x6c, 0xdd, 0x05, 0x8a, 0x55, 0xa7,
	0xbb, 0x45, 0xdb, 0xf6, 0x92, 0x34, 0x7c, 0x33, 0xf0, 0x28, 0xcd, 0xff, 0x44, 0x0f, 0xb3, 0xad,
	0x12, 0xb3, 0x91, 0x77, 0xa4, 0x19, 0xb0, 0x75, 0x32, 0xb5, 0xf9, 0xea, 0x66, 0xa0, 0xba, 0xba,
	0xc8, 0xb7, 0xc2, 0xa3, 0xd5, 0x56, 0x78, 0xb4, 0xda, 0x1d, 0x4e, 0xb9, 0x0d, 0xf7, 0xc9, 0x61,
	0xc0, 0xd6, 0x7d, 0x58, 0x80, 0x06, 0x7a, 0x52, 0xf4, 0x65, 0x0c, 0x2b, 0x9c, 0x5d, 0x92, 0x6d,
	0xf9, 0x48, 0x6e, 0x07, 0x6c, 0x9d, 0xee, 0xae, 0xd2, 0xbb, 0x8a, 0x1b, 0xac, 0x53, 0x15, 0x6c,
	0xfe, 0x3d, 0xb9, 0xe3, 0xe5, 0x47, 0x05, 0x6d, 0xbb, 0xc5, 0x83, 0x23, 0xdf, 0xa1, 0x65, 0x5a,
	0x98, 0x80, 0x46, 0xc0, 0xae, 0x34, 0x75, 0xdc, 0xf2, 0x69, 0x92, 0xc0, 0x21, 0x28, 0xc5, 0xa6,
	0xe0, 0xb4, 0xb6, 0xb4, 0xbe, 0x88, 0xe0, 0xbc, 0xf6, 0xac, 0x4e, 0x3f, 0x93, 0x83, 0x71, 0xc4,
	0x62, 0x35, 0x13, 0x9a, 0x9e, 0x6e, 0x99, 0x50, 0xf0, 0x66, 0xab, 0x28, 0xdc, 0x5f, 0x31, 0x24,
	0x47, 0x43, 0xc3, 0xf9, 0x00, 0x74, 0xef, 0xda, 0x87, 0x6b, 0x9a, 0xff, 0x20, 0x45, 0x8c, 0x9f,
	0x71, 0xb6, 0x47, 0xb5, 0x9f, 0xf3, 0x95, 0xdc, 0xb3, 0xca, 0x58, 0x33, 0xa9, 0x2f, 0x15, 0x7d,
	0x52, 0x0d, 0xdd, 0x48, 0x58, 0xfb, 0xf4, 0x3f, 0x0e, 0xac, 0xee, 0x3d, 0xff, 0xb5, 0xe9, 0xd6,
	0x7f, 0x6f, 0xba, 0xf5, 0x3f, 0x9b, 0x6e, 0xfd, 0xc7, 0xdf, 0x6e, 0x8d, 0x74, 0xb8, 0x58, 0xba,
	0xf1, 0x3c, 0x9a, 0x72, 0x16, 0xbb, 0xc9, 0xb1, 0xed, 0x86, 0x26, 0x3d, 0x6f, 0x27, 0xcd, 0xf4,
	0xf2, 0xfa, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xbe, 0x8c, 0x6b, 0xdb, 0x05, 0x00, 0x00,
}

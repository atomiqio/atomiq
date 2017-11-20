// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/appcelerator/amp/api/rpc/cluster/cluster.proto

/*
Package cluster is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/cluster/cluster.proto

It has these top-level messages:
	CreateRequest
	CreateReply
	ListRequest
	ListReply
	StatusRequest
	StatusReply
	UpdateRequest
	UpdateReply
	RemoveRequest
	RemoveReply
	NodeListRequest
	NodeListReply
	NodeReply
*/
package cluster

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
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

type CreateRequest struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Compose []byte `protobuf:"bytes,2,opt,name=compose,proto3" json:"compose,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetCompose() []byte {
	if m != nil {
		return m.Compose
	}
	return nil
}

type CreateReply struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateReply) Reset()                    { *m = CreateReply{} }
func (m *CreateReply) String() string            { return proto.CompactTextString(m) }
func (*CreateReply) ProtoMessage()               {}
func (*CreateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ListReply struct {
}

func (m *ListReply) Reset()                    { *m = ListReply{} }
func (m *ListReply) String() string            { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()               {}
func (*ListReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type StatusRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *StatusRequest) Reset()                    { *m = StatusRequest{} }
func (m *StatusRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()               {}
func (*StatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StatusRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type StatusReply struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *StatusReply) Reset()                    { *m = StatusReply{} }
func (m *StatusReply) String() string            { return proto.CompactTextString(m) }
func (*StatusReply) ProtoMessage()               {}
func (*StatusReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *StatusReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateReply struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *UpdateReply) Reset()                    { *m = UpdateReply{} }
func (m *UpdateReply) String() string            { return proto.CompactTextString(m) }
func (*UpdateReply) ProtoMessage()               {}
func (*UpdateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UpdateReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RemoveRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *RemoveRequest) Reset()                    { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()               {}
func (*RemoveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RemoveRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RemoveReply struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *RemoveReply) Reset()                    { *m = RemoveReply{} }
func (m *RemoveReply) String() string            { return proto.CompactTextString(m) }
func (*RemoveReply) ProtoMessage()               {}
func (*RemoveReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RemoveReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type NodeListRequest struct {
}

func (m *NodeListRequest) Reset()                    { *m = NodeListRequest{} }
func (m *NodeListRequest) String() string            { return proto.CompactTextString(m) }
func (*NodeListRequest) ProtoMessage()               {}
func (*NodeListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type NodeListReply struct {
	Nodes []*NodeReply `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *NodeListReply) Reset()                    { *m = NodeListReply{} }
func (m *NodeListReply) String() string            { return proto.CompactTextString(m) }
func (*NodeListReply) ProtoMessage()               {}
func (*NodeListReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *NodeListReply) GetNodes() []*NodeReply {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type NodeReply struct {
	Id            string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Hostname      string `protobuf:"bytes,2,opt,name=hostname" json:"hostname,omitempty"`
	Status        string `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	Availability  string `protobuf:"bytes,4,opt,name=availability" json:"availability,omitempty"`
	Role          string `protobuf:"bytes,5,opt,name=role" json:"role,omitempty"`
	ManagerLeader bool   `protobuf:"varint,6,opt,name=manager_leader,json=managerLeader" json:"manager_leader,omitempty"`
}

func (m *NodeReply) Reset()                    { *m = NodeReply{} }
func (m *NodeReply) String() string            { return proto.CompactTextString(m) }
func (*NodeReply) ProtoMessage()               {}
func (*NodeReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *NodeReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NodeReply) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *NodeReply) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *NodeReply) GetAvailability() string {
	if m != nil {
		return m.Availability
	}
	return ""
}

func (m *NodeReply) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *NodeReply) GetManagerLeader() bool {
	if m != nil {
		return m.ManagerLeader
	}
	return false
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "cluster.CreateRequest")
	proto.RegisterType((*CreateReply)(nil), "cluster.CreateReply")
	proto.RegisterType((*ListRequest)(nil), "cluster.ListRequest")
	proto.RegisterType((*ListReply)(nil), "cluster.ListReply")
	proto.RegisterType((*StatusRequest)(nil), "cluster.StatusRequest")
	proto.RegisterType((*StatusReply)(nil), "cluster.StatusReply")
	proto.RegisterType((*UpdateRequest)(nil), "cluster.UpdateRequest")
	proto.RegisterType((*UpdateReply)(nil), "cluster.UpdateReply")
	proto.RegisterType((*RemoveRequest)(nil), "cluster.RemoveRequest")
	proto.RegisterType((*RemoveReply)(nil), "cluster.RemoveReply")
	proto.RegisterType((*NodeListRequest)(nil), "cluster.NodeListRequest")
	proto.RegisterType((*NodeListReply)(nil), "cluster.NodeListReply")
	proto.RegisterType((*NodeReply)(nil), "cluster.NodeReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Cluster service

type ClusterClient interface {
	ClusterCreate(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error)
	ClusterList(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
	ClusterNodeList(ctx context.Context, in *NodeListRequest, opts ...grpc.CallOption) (*NodeListReply, error)
	ClusterStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error)
	ClusterUpdate(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateReply, error)
	ClusterRemove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error)
}

type clusterClient struct {
	cc *grpc.ClientConn
}

func NewClusterClient(cc *grpc.ClientConn) ClusterClient {
	return &clusterClient{cc}
}

func (c *clusterClient) ClusterCreate(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error) {
	out := new(CreateReply)
	err := grpc.Invoke(ctx, "/cluster.Cluster/ClusterCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) ClusterList(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := grpc.Invoke(ctx, "/cluster.Cluster/ClusterList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) ClusterNodeList(ctx context.Context, in *NodeListRequest, opts ...grpc.CallOption) (*NodeListReply, error) {
	out := new(NodeListReply)
	err := grpc.Invoke(ctx, "/cluster.Cluster/ClusterNodeList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) ClusterStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := grpc.Invoke(ctx, "/cluster.Cluster/ClusterStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) ClusterUpdate(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateReply, error) {
	out := new(UpdateReply)
	err := grpc.Invoke(ctx, "/cluster.Cluster/ClusterUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) ClusterRemove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error) {
	out := new(RemoveReply)
	err := grpc.Invoke(ctx, "/cluster.Cluster/ClusterRemove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cluster service

type ClusterServer interface {
	ClusterCreate(context.Context, *CreateRequest) (*CreateReply, error)
	ClusterList(context.Context, *ListRequest) (*ListReply, error)
	ClusterNodeList(context.Context, *NodeListRequest) (*NodeListReply, error)
	ClusterStatus(context.Context, *StatusRequest) (*StatusReply, error)
	ClusterUpdate(context.Context, *UpdateRequest) (*UpdateReply, error)
	ClusterRemove(context.Context, *RemoveRequest) (*RemoveReply, error)
}

func RegisterClusterServer(s *grpc.Server, srv ClusterServer) {
	s.RegisterService(&_Cluster_serviceDesc, srv)
}

func _Cluster_ClusterCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).ClusterCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.Cluster/ClusterCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).ClusterCreate(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_ClusterList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).ClusterList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.Cluster/ClusterList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).ClusterList(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_ClusterNodeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).ClusterNodeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.Cluster/ClusterNodeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).ClusterNodeList(ctx, req.(*NodeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_ClusterStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).ClusterStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.Cluster/ClusterStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).ClusterStatus(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_ClusterUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).ClusterUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.Cluster/ClusterUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).ClusterUpdate(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_ClusterRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).ClusterRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.Cluster/ClusterRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).ClusterRemove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cluster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cluster.Cluster",
	HandlerType: (*ClusterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClusterCreate",
			Handler:    _Cluster_ClusterCreate_Handler,
		},
		{
			MethodName: "ClusterList",
			Handler:    _Cluster_ClusterList_Handler,
		},
		{
			MethodName: "ClusterNodeList",
			Handler:    _Cluster_ClusterNodeList_Handler,
		},
		{
			MethodName: "ClusterStatus",
			Handler:    _Cluster_ClusterStatus_Handler,
		},
		{
			MethodName: "ClusterUpdate",
			Handler:    _Cluster_ClusterUpdate_Handler,
		},
		{
			MethodName: "ClusterRemove",
			Handler:    _Cluster_ClusterRemove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/appcelerator/amp/api/rpc/cluster/cluster.proto",
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/cluster/cluster.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xd1, 0x6a, 0xdb, 0x30,
	0x14, 0x86, 0x71, 0x9a, 0x26, 0xcd, 0x71, 0x9c, 0x92, 0xb3, 0x90, 0x79, 0xa6, 0xa5, 0xc1, 0x30,
	0x08, 0xbd, 0x88, 0x59, 0x77, 0xd5, 0xc1, 0xae, 0x7a, 0xdb, 0xed, 0xc2, 0x63, 0x63, 0xb0, 0x8b,
	0xa2, 0xc4, 0x22, 0x15, 0xd8, 0x96, 0x66, 0x29, 0x81, 0x30, 0x76, 0xb3, 0x57, 0xd8, 0x1b, 0xec,
	0x01, 0xf6, 0x32, 0x7b, 0x85, 0x3d, 0xc8, 0xb0, 0x6c, 0x25, 0x72, 0x92, 0xf5, 0x2a, 0x3a, 0x47,
	0xe7, 0xff, 0x2c, 0x9d, 0xf3, 0x2b, 0x70, 0xbb, 0x64, 0xea, 0x71, 0x35, 0x9f, 0x2d, 0x78, 0x16,
	0x11, 0x21, 0x16, 0x34, 0xa5, 0x05, 0x51, 0xbc, 0x88, 0x48, 0x26, 0x22, 0x22, 0x58, 0x54, 0x88,
	0x45, 0xb4, 0x48, 0x57, 0x52, 0xd1, 0xc2, 0xfc, 0xce, 0x44, 0xc1, 0x15, 0xc7, 0x6e, 0x1d, 0x06,
	0x17, 0x4b, 0xce, 0x97, 0x29, 0xd5, 0xe5, 0x24, 0xcf, 0xb9, 0x22, 0x8a, 0xf1, 0x5c, 0x56, 0x65,
	0xe1, 0x5b, 0xf0, 0xee, 0x0a, 0x4a, 0x14, 0x8d, 0xe9, 0xd7, 0x15, 0x95, 0x0a, 0x11, 0xda, 0x39,
	0xc9, 0xa8, 0xef, 0x4c, 0x9c, 0x69, 0x2f, 0xd6, 0x6b, 0xf4, 0xa1, 0xbb, 0xe0, 0x99, 0xe0, 0x92,
	0xfa, 0xad, 0x89, 0x33, 0xed, 0xc7, 0x26, 0x0c, 0x2f, 0xc1, 0x35, 0x72, 0x91, 0x6e, 0x70, 0x00,
	0x2d, 0x96, 0xd4, 0xd2, 0x16, 0x4b, 0x42, 0x0f, 0xdc, 0x7b, 0x26, 0x55, 0xcd, 0x0e, 0x5d, 0xe8,
	0x55, 0xa1, 0x48, 0x37, 0xe1, 0x15, 0x78, 0x1f, 0x14, 0x51, 0x2b, 0x69, 0xbe, 0xbc, 0x2f, 0xbe,
	0x04, 0xd7, 0x14, 0x1c, 0x63, 0x5f, 0x81, 0xf7, 0x51, 0x24, 0xd6, 0xc9, 0x8f, 0xe8, 0x4d, 0xc1,
	0x7f, 0xf4, 0x31, 0xcd, 0xf8, 0xfa, 0x29, 0xbd, 0x29, 0x38, 0xa6, 0x1f, 0xc2, 0xf9, 0x7b, 0x9e,
	0x50, 0xfb, 0x7e, 0xb7, 0xe0, 0xed, 0x52, 0xa5, 0x66, 0x0a, 0xa7, 0x39, 0x4f, 0xa8, 0xf4, 0x9d,
	0xc9, 0xc9, 0xd4, 0xbd, 0xc1, 0x99, 0x99, 0x51, 0x59, 0xa6, 0x4b, 0xe2, 0xaa, 0x20, 0xfc, 0xed,
	0x40, 0x6f, 0x9b, 0xdc, 0xff, 0x16, 0x06, 0x70, 0xf6, 0xc8, 0xa5, 0xd2, 0x83, 0x69, 0xe9, 0xec,
	0x36, 0xc6, 0x31, 0x74, 0xa4, 0x6e, 0x93, 0x7f, 0xa2, 0x77, 0xea, 0x08, 0x43, 0xe8, 0x93, 0x35,
	0x61, 0x29, 0x99, 0xb3, 0x94, 0xa9, 0x8d, 0xdf, 0xd6, 0xbb, 0x8d, 0x5c, 0x39, 0xec, 0x82, 0xa7,
	0xd4, 0x3f, 0xad, 0x86, 0x5d, 0xae, 0xf1, 0x25, 0x0c, 0x32, 0x92, 0x93, 0x25, 0x2d, 0x1e, 0x52,
	0x4a, 0x12, 0x5a, 0xf8, 0x9d, 0x89, 0x33, 0x3d, 0x8b, 0xbd, 0x3a, 0x7b, 0xaf, 0x93, 0x37, 0xbf,
	0xda, 0xd0, 0xbd, 0xab, 0x6e, 0x83, 0x9f, 0xc0, 0xab, 0x97, 0x95, 0x19, 0x70, 0xbc, 0xbd, 0x68,
	0xc3, 0x5c, 0xc1, 0xe8, 0x20, 0x5f, 0x3a, 0xe1, 0xf9, 0x8f, 0x3f, 0x7f, 0x7f, 0xb6, 0x86, 0x61,
	0x3f, 0x5a, 0xbf, 0x32, 0x2e, 0x96, 0x6f, 0x9c, 0x6b, 0x7c, 0x07, 0x6e, 0xcd, 0x2d, 0x5b, 0x8a,
	0x3b, 0xb5, 0xd5, 0xf4, 0x00, 0xf7, 0xb2, 0x25, 0x71, 0xa4, 0x89, 0x03, 0x6c, 0x10, 0xf1, 0x01,
	0xce, 0x6b, 0x9c, 0x99, 0x12, 0xfa, 0x8d, 0x89, 0xd8, 0xd8, 0xf1, 0x91, 0x9d, 0x12, 0x1d, 0x68,
	0xf4, 0x08, 0xd1, 0x46, 0x47, 0x7a, 0x88, 0xf8, 0x79, 0xdb, 0x87, 0xca, 0xb8, 0x56, 0x1f, 0x1a,
	0x56, 0xb7, 0xfa, 0x60, 0x39, 0x3c, 0x7c, 0xa1, 0xd1, 0xcf, 0x70, 0xd8, 0x40, 0x7f, 0x63, 0xc9,
	0x77, 0xfc, 0xb2, 0x25, 0x57, 0x96, 0xb6, 0xc8, 0x8d, 0x47, 0x60, 0x91, 0x2d, 0xef, 0x87, 0x17,
	0x9a, 0x3c, 0x0e, 0x0e, 0xc9, 0x65, 0x9b, 0x77, 0xc7, 0xae, 0xfc, 0x6e, 0xc1, 0x1b, 0x2f, 0xc4,
	0x82, 0x5b, 0x0f, 0xc3, 0x1c, 0xfb, 0xfa, 0x10, 0x3e, 0xef, 0xe8, 0x3f, 0x99, 0xd7, 0xff, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x8b, 0xe6, 0x60, 0x40, 0xc8, 0x04, 0x00, 0x00,
}

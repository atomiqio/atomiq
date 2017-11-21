// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/appcelerator/amp/api/rpc/stack/stack.proto

/*
Package stack is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/stack/stack.proto

It has these top-level messages:
	DeployRequest
	DeployReply
	ListRequest
	ListReply
	StackEntry
	RemoveRequest
	RemoveReply
	ServicesRequest
	ServiceEntry
	ServicesReply
*/
package stack

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import stacks "github.com/appcelerator/amp/data/stacks"
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

type DeployRequest struct {
	Name        string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Compose     []byte   `protobuf:"bytes,2,opt,name=compose,proto3" json:"compose,omitempty"`
	Environment []string `protobuf:"bytes,3,rep,name=environment" json:"environment,omitempty"`
	Config      []byte   `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
}

func (m *DeployRequest) Reset()                    { *m = DeployRequest{} }
func (m *DeployRequest) String() string            { return proto.CompactTextString(m) }
func (*DeployRequest) ProtoMessage()               {}
func (*DeployRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeployRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeployRequest) GetCompose() []byte {
	if m != nil {
		return m.Compose
	}
	return nil
}

func (m *DeployRequest) GetEnvironment() []string {
	if m != nil {
		return m.Environment
	}
	return nil
}

func (m *DeployRequest) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

type DeployReply struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	FullName string `protobuf:"bytes,2,opt,name=full_name,json=fullName" json:"full_name,omitempty"`
	Answer   string `protobuf:"bytes,3,opt,name=answer" json:"answer,omitempty"`
}

func (m *DeployReply) Reset()                    { *m = DeployReply{} }
func (m *DeployReply) String() string            { return proto.CompactTextString(m) }
func (*DeployReply) ProtoMessage()               {}
func (*DeployReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeployReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeployReply) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *DeployReply) GetAnswer() string {
	if m != nil {
		return m.Answer
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
	Entries []*StackEntry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *ListReply) Reset()                    { *m = ListReply{} }
func (m *ListReply) String() string            { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()               {}
func (*ListReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListReply) GetEntries() []*StackEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type StackEntry struct {
	Stack             *stacks.Stack `protobuf:"bytes,1,opt,name=stack" json:"stack,omitempty"`
	RunningServices   int32         `protobuf:"varint,2,opt,name=running_services,json=runningServices" json:"running_services,omitempty"`
	TotalServices     int32         `protobuf:"varint,3,opt,name=total_services,json=totalServices" json:"total_services,omitempty"`
	Status            string        `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	CompleteServices  int32         `protobuf:"varint,5,opt,name=complete_services,json=completeServices" json:"complete_services,omitempty"`
	PreparingServices int32         `protobuf:"varint,6,opt,name=preparing_services,json=preparingServices" json:"preparing_services,omitempty"`
}

func (m *StackEntry) Reset()                    { *m = StackEntry{} }
func (m *StackEntry) String() string            { return proto.CompactTextString(m) }
func (*StackEntry) ProtoMessage()               {}
func (*StackEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StackEntry) GetStack() *stacks.Stack {
	if m != nil {
		return m.Stack
	}
	return nil
}

func (m *StackEntry) GetRunningServices() int32 {
	if m != nil {
		return m.RunningServices
	}
	return 0
}

func (m *StackEntry) GetTotalServices() int32 {
	if m != nil {
		return m.TotalServices
	}
	return 0
}

func (m *StackEntry) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *StackEntry) GetCompleteServices() int32 {
	if m != nil {
		return m.CompleteServices
	}
	return 0
}

func (m *StackEntry) GetPreparingServices() int32 {
	if m != nil {
		return m.PreparingServices
	}
	return 0
}

type RemoveRequest struct {
	Stack string `protobuf:"bytes,1,opt,name=stack" json:"stack,omitempty"`
}

func (m *RemoveRequest) Reset()                    { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()               {}
func (*RemoveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RemoveRequest) GetStack() string {
	if m != nil {
		return m.Stack
	}
	return ""
}

type RemoveReply struct {
	Answer string `protobuf:"bytes,1,opt,name=answer" json:"answer,omitempty"`
}

func (m *RemoveReply) Reset()                    { *m = RemoveReply{} }
func (m *RemoveReply) String() string            { return proto.CompactTextString(m) }
func (*RemoveReply) ProtoMessage()               {}
func (*RemoveReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RemoveReply) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

type ServicesRequest struct {
	Stack string `protobuf:"bytes,1,opt,name=stack" json:"stack,omitempty"`
}

func (m *ServicesRequest) Reset()                    { *m = ServicesRequest{} }
func (m *ServicesRequest) String() string            { return proto.CompactTextString(m) }
func (*ServicesRequest) ProtoMessage()               {}
func (*ServicesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ServicesRequest) GetStack() string {
	if m != nil {
		return m.Stack
	}
	return ""
}

type ServiceEntry struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Mode     string `protobuf:"bytes,3,opt,name=mode" json:"mode,omitempty"`
	Replicas string `protobuf:"bytes,4,opt,name=replicas" json:"replicas,omitempty"`
	Image    string `protobuf:"bytes,5,opt,name=image" json:"image,omitempty"`
}

func (m *ServiceEntry) Reset()                    { *m = ServiceEntry{} }
func (m *ServiceEntry) String() string            { return proto.CompactTextString(m) }
func (*ServiceEntry) ProtoMessage()               {}
func (*ServiceEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ServiceEntry) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServiceEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceEntry) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *ServiceEntry) GetReplicas() string {
	if m != nil {
		return m.Replicas
	}
	return ""
}

func (m *ServiceEntry) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

type ServicesReply struct {
	Services []*ServiceEntry `protobuf:"bytes,1,rep,name=services" json:"services,omitempty"`
}

func (m *ServicesReply) Reset()                    { *m = ServicesReply{} }
func (m *ServicesReply) String() string            { return proto.CompactTextString(m) }
func (*ServicesReply) ProtoMessage()               {}
func (*ServicesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ServicesReply) GetServices() []*ServiceEntry {
	if m != nil {
		return m.Services
	}
	return nil
}

func init() {
	proto.RegisterType((*DeployRequest)(nil), "stack.DeployRequest")
	proto.RegisterType((*DeployReply)(nil), "stack.DeployReply")
	proto.RegisterType((*ListRequest)(nil), "stack.ListRequest")
	proto.RegisterType((*ListReply)(nil), "stack.ListReply")
	proto.RegisterType((*StackEntry)(nil), "stack.StackEntry")
	proto.RegisterType((*RemoveRequest)(nil), "stack.RemoveRequest")
	proto.RegisterType((*RemoveReply)(nil), "stack.RemoveReply")
	proto.RegisterType((*ServicesRequest)(nil), "stack.ServicesRequest")
	proto.RegisterType((*ServiceEntry)(nil), "stack.ServiceEntry")
	proto.RegisterType((*ServicesReply)(nil), "stack.ServicesReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Stack service

type StackClient interface {
	StackDeploy(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*DeployReply, error)
	StackList(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
	StackRemove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error)
	StackServices(ctx context.Context, in *ServicesRequest, opts ...grpc.CallOption) (*ServicesReply, error)
}

type stackClient struct {
	cc *grpc.ClientConn
}

func NewStackClient(cc *grpc.ClientConn) StackClient {
	return &stackClient{cc}
}

func (c *stackClient) StackDeploy(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*DeployReply, error) {
	out := new(DeployReply)
	err := grpc.Invoke(ctx, "/stack.Stack/StackDeploy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackClient) StackList(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := grpc.Invoke(ctx, "/stack.Stack/StackList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackClient) StackRemove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error) {
	out := new(RemoveReply)
	err := grpc.Invoke(ctx, "/stack.Stack/StackRemove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackClient) StackServices(ctx context.Context, in *ServicesRequest, opts ...grpc.CallOption) (*ServicesReply, error) {
	out := new(ServicesReply)
	err := grpc.Invoke(ctx, "/stack.Stack/StackServices", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Stack service

type StackServer interface {
	StackDeploy(context.Context, *DeployRequest) (*DeployReply, error)
	StackList(context.Context, *ListRequest) (*ListReply, error)
	StackRemove(context.Context, *RemoveRequest) (*RemoveReply, error)
	StackServices(context.Context, *ServicesRequest) (*ServicesReply, error)
}

func RegisterStackServer(s *grpc.Server, srv StackServer) {
	s.RegisterService(&_Stack_serviceDesc, srv)
}

func _Stack_StackDeploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServer).StackDeploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.Stack/StackDeploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServer).StackDeploy(ctx, req.(*DeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stack_StackList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServer).StackList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.Stack/StackList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServer).StackList(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stack_StackRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServer).StackRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.Stack/StackRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServer).StackRemove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stack_StackServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServer).StackServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.Stack/StackServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServer).StackServices(ctx, req.(*ServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stack_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stack.Stack",
	HandlerType: (*StackServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StackDeploy",
			Handler:    _Stack_StackDeploy_Handler,
		},
		{
			MethodName: "StackList",
			Handler:    _Stack_StackList_Handler,
		},
		{
			MethodName: "StackRemove",
			Handler:    _Stack_StackRemove_Handler,
		},
		{
			MethodName: "StackServices",
			Handler:    _Stack_StackServices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/appcelerator/amp/api/rpc/stack/stack.proto",
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/stack/stack.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xcd, 0x6e, 0x9c, 0x48,
	0x10, 0x16, 0x8c, 0x67, 0x6c, 0x0a, 0xe3, 0x9f, 0xda, 0x59, 0x0b, 0xe1, 0x5d, 0x69, 0x84, 0x65,
	0xed, 0xac, 0xad, 0x0c, 0x8a, 0x93, 0x48, 0x51, 0x4e, 0x39, 0xc4, 0xb7, 0xc8, 0x07, 0xfc, 0x00,
	0x56, 0x9b, 0x69, 0x4f, 0x50, 0xa0, 0x9b, 0x40, 0xcf, 0x38, 0xa3, 0x28, 0x97, 0xbc, 0x42, 0x1e,
	0x2d, 0x87, 0xbc, 0x40, 0xde, 0x21, 0xd7, 0x88, 0xa2, 0x1b, 0x18, 0x5b, 0xf2, 0x05, 0xba, 0xaa,
	0xab, 0xbe, 0xfe, 0xfa, 0xab, 0x0f, 0xe0, 0xd5, 0x22, 0x55, 0x1f, 0x96, 0xb7, 0xb3, 0x44, 0xe6,
	0x11, 0x2b, 0x8a, 0x84, 0x67, 0xbc, 0x64, 0x4a, 0x96, 0x11, 0xcb, 0x8b, 0x88, 0x15, 0x69, 0x54,
	0x16, 0x49, 0x54, 0x29, 0x96, 0x7c, 0x6c, 0x9e, 0xb3, 0xa2, 0x94, 0x4a, 0xe2, 0x90, 0x82, 0xe0,
	0xe5, 0x53, 0xdd, 0x73, 0xa6, 0x58, 0xd3, 0x54, 0xe9, 0x57, 0xd3, 0x1c, 0xfc, 0xb3, 0x90, 0x72,
	0x91, 0x71, 0x82, 0x67, 0x42, 0x48, 0xc5, 0x54, 0x2a, 0x85, 0xde, 0x0d, 0xef, 0xc1, 0x7b, 0xc7,
	0x8b, 0x4c, 0xae, 0x63, 0xfe, 0x69, 0xc9, 0x2b, 0x85, 0x08, 0x5b, 0x82, 0xe5, 0xdc, 0xb7, 0x26,
	0xd6, 0xd4, 0x89, 0x69, 0x8d, 0x3e, 0x6c, 0x27, 0x32, 0x2f, 0x64, 0xc5, 0x7d, 0x7b, 0x62, 0x4d,
	0x77, 0x63, 0x13, 0xe2, 0x04, 0x5c, 0x2e, 0x56, 0x69, 0x29, 0x45, 0xce, 0x85, 0xf2, 0x07, 0x93,
	0xc1, 0xd4, 0x89, 0xfb, 0x29, 0x3c, 0x82, 0x51, 0x22, 0xc5, 0x5d, 0xba, 0xf0, 0xb7, 0xa8, 0x55,
	0x47, 0x61, 0x0c, 0xae, 0x39, 0xb8, 0xc8, 0xd6, 0xb8, 0x07, 0x76, 0x3a, 0xd7, 0x87, 0xda, 0xe9,
	0x1c, 0x8f, 0xc1, 0xb9, 0x5b, 0x66, 0xd9, 0x0d, 0x71, 0xb1, 0x29, 0xbd, 0x53, 0x27, 0xae, 0x6a,
	0x3e, 0x47, 0x30, 0x62, 0xa2, 0xba, 0xe7, 0xa5, 0x3f, 0xa0, 0x1d, 0x1d, 0x85, 0x1e, 0xb8, 0xef,
	0xd3, 0x4a, 0xe9, 0xab, 0x84, 0xaf, 0xc1, 0x69, 0xc2, 0xfa, 0x80, 0x73, 0xd8, 0xe6, 0x42, 0x95,
	0x29, 0xaf, 0x7c, 0x6b, 0x32, 0x98, 0xba, 0x17, 0x87, 0xb3, 0x46, 0xe2, 0xeb, 0xfa, 0x79, 0x29,
	0x54, 0xb9, 0x8e, 0x4d, 0x45, 0xf8, 0xdb, 0x02, 0xe8, 0xf2, 0x78, 0x02, 0xcd, 0x04, 0x88, 0x9f,
	0x7b, 0xe1, 0xcd, 0xb4, 0xc0, 0x54, 0x12, 0x37, 0x7b, 0xf8, 0x3f, 0x1c, 0x94, 0x4b, 0x21, 0x52,
	0xb1, 0xb8, 0xa9, 0x78, 0xb9, 0x4a, 0x13, 0x5e, 0x11, 0xf1, 0x61, 0xbc, 0xaf, 0xf3, 0xd7, 0x3a,
	0x8d, 0xa7, 0xb0, 0xa7, 0xa4, 0x62, 0x59, 0x57, 0x38, 0xa0, 0x42, 0x8f, 0xb2, 0x6d, 0xd9, 0x11,
	0x8c, 0x2a, 0xc5, 0xd4, 0xb2, 0x22, 0xe9, 0x9c, 0x58, 0x47, 0x78, 0x0e, 0x87, 0xb5, 0xfe, 0x19,
	0x57, 0xbc, 0x43, 0x18, 0x12, 0xc2, 0x81, 0xd9, 0x68, 0x41, 0x9e, 0x01, 0x16, 0x25, 0x2f, 0x58,
	0xb9, 0x41, 0x6c, 0x44, 0xd5, 0x87, 0xed, 0x8e, 0x29, 0x0f, 0x4f, 0xc1, 0x8b, 0x79, 0x2e, 0x57,
	0xdc, 0xf8, 0x61, 0xdc, 0xbf, 0xbb, 0xa3, 0x2f, 0x1b, 0x9e, 0x82, 0x6b, 0xca, 0x6a, 0x71, 0xbb,
	0x81, 0x58, 0x1b, 0x03, 0xf9, 0x0f, 0xf6, 0x0d, 0xf2, 0xd3, 0x78, 0x9f, 0x61, 0x57, 0x17, 0x36,
	0x8a, 0x3f, 0xb4, 0x83, 0x71, 0xa5, 0xdd, 0x73, 0x25, 0xc2, 0x56, 0x2e, 0xe7, 0x5c, 0x7b, 0x80,
	0xd6, 0x18, 0xc0, 0x4e, 0xc9, 0x8b, 0x2c, 0x4d, 0x98, 0x11, 0xad, 0x8d, 0xeb, 0x93, 0xd3, 0x9c,
	0x2d, 0x38, 0x49, 0xe5, 0xc4, 0x4d, 0x10, 0xbe, 0x05, 0xaf, 0xa3, 0x58, 0xdf, 0x25, 0x82, 0x9d,
	0x56, 0xa6, 0xc6, 0x29, 0x7f, 0x19, 0xa7, 0xf4, 0x18, 0xc6, 0x6d, 0xd1, 0xc5, 0x4f, 0x1b, 0x86,
	0xe4, 0x04, 0xbc, 0x02, 0x97, 0x16, 0x8d, 0xb1, 0x71, 0xac, 0xfb, 0x36, 0x3e, 0xb0, 0x00, 0x1f,
	0x64, 0x8b, 0x6c, 0x1d, 0xfe, 0xfd, 0xed, 0xc7, 0xaf, 0xef, 0xf6, 0x7e, 0x08, 0xd1, 0xea, 0xb9,
	0xfe, 0x7a, 0xdf, 0x58, 0x67, 0x78, 0x09, 0x0e, 0xe1, 0xd5, 0x2e, 0x46, 0xd3, 0xd7, 0x73, 0x78,
	0x70, 0xb0, 0x91, 0xab, 0x91, 0x90, 0x90, 0x76, 0xb1, 0x87, 0x84, 0xd7, 0x9a, 0x56, 0x33, 0xb1,
	0x96, 0xd6, 0xc6, 0x9c, 0x5b, 0x5a, 0xbd, 0xb1, 0x86, 0x01, 0x81, 0x8d, 0xcf, 0xb0, 0x03, 0x8b,
	0xbe, 0xd0, 0xfb, 0x2b, 0xde, 0x82, 0x47, 0xa0, 0x9d, 0x5b, 0x37, 0x55, 0x32, 0x03, 0x0f, 0xc6,
	0x8f, 0xf2, 0x35, 0xf4, 0x09, 0x41, 0xff, 0x8b, 0xc7, 0x8f, 0xa1, 0x23, 0xa3, 0xec, 0xed, 0x88,
	0xfe, 0x51, 0x2f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x5c, 0x6c, 0xc6, 0x37, 0x05, 0x00,
	0x00,
}

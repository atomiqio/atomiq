// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/cmd/adm-server/servergrpc/cluster-server.proto
// DO NOT EDIT!

/*
Package servergrpc is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/cmd/adm-server/servergrpc/cluster-server.proto

It has these top-level messages:
	ClientMes
	RegisterRequest
	AgentHealthRequest
	ServerRet
	ServerRets
	GetNodesInfoRequest
	NodesInfo
	NodeInfo
	PurgeNodesRequest
	PurgeNodeAnswer
	PurgeNodesAnswers
	AmpRequest
	AmpRet
	TypedOutput
	AmpMonitorAnswers
	AmpStatusAnswer
*/
package servergrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

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

type ClientMes struct {
	ClientId string       `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	Function string       `protobuf:"bytes,2,opt,name=function" json:"function,omitempty"`
	Output   *TypedOutput `protobuf:"bytes,3,opt,name=output" json:"output,omitempty"`
}

func (m *ClientMes) Reset()                    { *m = ClientMes{} }
func (m *ClientMes) String() string            { return proto.CompactTextString(m) }
func (*ClientMes) ProtoMessage()               {}
func (*ClientMes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClientMes) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *ClientMes) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func (m *ClientMes) GetOutput() *TypedOutput {
	if m != nil {
		return m.Output
	}
	return nil
}

type RegisterRequest struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	NodeId   string `protobuf:"bytes,2,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Hostname string `protobuf:"bytes,3,opt,name=hostname" json:"hostname,omitempty"`
	Address  string `protobuf:"bytes,4,opt,name=address" json:"address,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RegisterRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *RegisterRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *RegisterRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type AgentHealthRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *AgentHealthRequest) Reset()                    { *m = AgentHealthRequest{} }
func (m *AgentHealthRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentHealthRequest) ProtoMessage()               {}
func (*AgentHealthRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AgentHealthRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ServerRet struct {
	AgentId string `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *ServerRet) Reset()                    { *m = ServerRet{} }
func (m *ServerRet) String() string            { return proto.CompactTextString(m) }
func (*ServerRet) ProtoMessage()               {}
func (*ServerRet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ServerRet) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *ServerRet) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ServerRets struct {
	Rets []*ServerRet `protobuf:"bytes,1,rep,name=rets" json:"rets,omitempty"`
}

func (m *ServerRets) Reset()                    { *m = ServerRets{} }
func (m *ServerRets) String() string            { return proto.CompactTextString(m) }
func (*ServerRets) ProtoMessage()               {}
func (*ServerRets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ServerRets) GetRets() []*ServerRet {
	if m != nil {
		return m.Rets
	}
	return nil
}

type GetNodesInfoRequest struct {
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	Node     string `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
	More     bool   `protobuf:"varint,3,opt,name=more" json:"more,omitempty"`
}

func (m *GetNodesInfoRequest) Reset()                    { *m = GetNodesInfoRequest{} }
func (m *GetNodesInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNodesInfoRequest) ProtoMessage()               {}
func (*GetNodesInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetNodesInfoRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *GetNodesInfoRequest) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *GetNodesInfoRequest) GetMore() bool {
	if m != nil {
		return m.More
	}
	return false
}

type NodesInfo struct {
	Nodes []*NodeInfo `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *NodesInfo) Reset()                    { *m = NodesInfo{} }
func (m *NodesInfo) String() string            { return proto.CompactTextString(m) }
func (*NodesInfo) ProtoMessage()               {}
func (*NodesInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *NodesInfo) GetNodes() []*NodeInfo {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type NodeInfo struct {
	Id                  string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Role                string `protobuf:"bytes,2,opt,name=role" json:"role,omitempty"`
	Availability        string `protobuf:"bytes,3,opt,name=availability" json:"availability,omitempty"`
	Hostname            string `protobuf:"bytes,4,opt,name=hostname" json:"hostname,omitempty"`
	HostArchitecture    string `protobuf:"bytes,5,opt,name=host_architecture,json=hostArchitecture" json:"host_architecture,omitempty"`
	HostOs              string `protobuf:"bytes,6,opt,name=host_os,json=hostOs" json:"host_os,omitempty"`
	Cpu                 int64  `protobuf:"varint,7,opt,name=cpu" json:"cpu,omitempty"`
	Memory              int64  `protobuf:"varint,8,opt,name=memory" json:"memory,omitempty"`
	DockerVersion       string `protobuf:"bytes,9,opt,name=docker_version,json=dockerVersion" json:"docker_version,omitempty"`
	Status              string `protobuf:"bytes,10,opt,name=status" json:"status,omitempty"`
	Leader              bool   `protobuf:"varint,11,opt,name=leader" json:"leader,omitempty"`
	Reachability        string `protobuf:"bytes,12,opt,name=reachability" json:"reachability,omitempty"`
	Address             string `protobuf:"bytes,13,opt,name=address" json:"address,omitempty"`
	NbContainers        int64  `protobuf:"varint,14,opt,name=nb_containers,json=nbContainers" json:"nb_containers,omitempty"`
	NbContainersRunning int64  `protobuf:"varint,15,opt,name=nb_containers_running,json=nbContainersRunning" json:"nb_containers_running,omitempty"`
	NbContainersPaused  int64  `protobuf:"varint,16,opt,name=nb_containers_paused,json=nbContainersPaused" json:"nb_containers_paused,omitempty"`
	NbContainersStopped int64  `protobuf:"varint,17,opt,name=nb_containers_stopped,json=nbContainersStopped" json:"nb_containers_stopped,omitempty"`
	Images              int64  `protobuf:"varint,18,opt,name=images" json:"images,omitempty"`
	Error               string `protobuf:"bytes,19,opt,name=error" json:"error,omitempty"`
	AgentId             string `protobuf:"bytes,20,opt,name=agentId" json:"agentId,omitempty"`
	NodeName            string `protobuf:"bytes,21,opt,name=nodeName" json:"nodeName,omitempty"`
}

func (m *NodeInfo) Reset()                    { *m = NodeInfo{} }
func (m *NodeInfo) String() string            { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()               {}
func (*NodeInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *NodeInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NodeInfo) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *NodeInfo) GetAvailability() string {
	if m != nil {
		return m.Availability
	}
	return ""
}

func (m *NodeInfo) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *NodeInfo) GetHostArchitecture() string {
	if m != nil {
		return m.HostArchitecture
	}
	return ""
}

func (m *NodeInfo) GetHostOs() string {
	if m != nil {
		return m.HostOs
	}
	return ""
}

func (m *NodeInfo) GetCpu() int64 {
	if m != nil {
		return m.Cpu
	}
	return 0
}

func (m *NodeInfo) GetMemory() int64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *NodeInfo) GetDockerVersion() string {
	if m != nil {
		return m.DockerVersion
	}
	return ""
}

func (m *NodeInfo) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *NodeInfo) GetLeader() bool {
	if m != nil {
		return m.Leader
	}
	return false
}

func (m *NodeInfo) GetReachability() string {
	if m != nil {
		return m.Reachability
	}
	return ""
}

func (m *NodeInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *NodeInfo) GetNbContainers() int64 {
	if m != nil {
		return m.NbContainers
	}
	return 0
}

func (m *NodeInfo) GetNbContainersRunning() int64 {
	if m != nil {
		return m.NbContainersRunning
	}
	return 0
}

func (m *NodeInfo) GetNbContainersPaused() int64 {
	if m != nil {
		return m.NbContainersPaused
	}
	return 0
}

func (m *NodeInfo) GetNbContainersStopped() int64 {
	if m != nil {
		return m.NbContainersStopped
	}
	return 0
}

func (m *NodeInfo) GetImages() int64 {
	if m != nil {
		return m.Images
	}
	return 0
}

func (m *NodeInfo) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *NodeInfo) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *NodeInfo) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type PurgeNodesRequest struct {
	ClientId  string `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	Node      string `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
	Container bool   `protobuf:"varint,3,opt,name=container" json:"container,omitempty"`
	Volume    bool   `protobuf:"varint,4,opt,name=volume" json:"volume,omitempty"`
	Image     bool   `protobuf:"varint,5,opt,name=image" json:"image,omitempty"`
	Force     bool   `protobuf:"varint,6,opt,name=force" json:"force,omitempty"`
}

func (m *PurgeNodesRequest) Reset()                    { *m = PurgeNodesRequest{} }
func (m *PurgeNodesRequest) String() string            { return proto.CompactTextString(m) }
func (*PurgeNodesRequest) ProtoMessage()               {}
func (*PurgeNodesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PurgeNodesRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *PurgeNodesRequest) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *PurgeNodesRequest) GetContainer() bool {
	if m != nil {
		return m.Container
	}
	return false
}

func (m *PurgeNodesRequest) GetVolume() bool {
	if m != nil {
		return m.Volume
	}
	return false
}

func (m *PurgeNodesRequest) GetImage() bool {
	if m != nil {
		return m.Image
	}
	return false
}

func (m *PurgeNodesRequest) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

type PurgeNodeAnswer struct {
	ClientId     string `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	AgentId      string `protobuf:"bytes,2,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	NbContainers int32  `protobuf:"varint,3,opt,name=nb_containers,json=nbContainers" json:"nb_containers,omitempty"`
	NbVolumes    int32  `protobuf:"varint,4,opt,name=nb_volumes,json=nbVolumes" json:"nb_volumes,omitempty"`
	NbImages     int32  `protobuf:"varint,5,opt,name=nb_images,json=nbImages" json:"nb_images,omitempty"`
	Error        string `protobuf:"bytes,6,opt,name=error" json:"error,omitempty"`
}

func (m *PurgeNodeAnswer) Reset()                    { *m = PurgeNodeAnswer{} }
func (m *PurgeNodeAnswer) String() string            { return proto.CompactTextString(m) }
func (*PurgeNodeAnswer) ProtoMessage()               {}
func (*PurgeNodeAnswer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *PurgeNodeAnswer) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *PurgeNodeAnswer) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *PurgeNodeAnswer) GetNbContainers() int32 {
	if m != nil {
		return m.NbContainers
	}
	return 0
}

func (m *PurgeNodeAnswer) GetNbVolumes() int32 {
	if m != nil {
		return m.NbVolumes
	}
	return 0
}

func (m *PurgeNodeAnswer) GetNbImages() int32 {
	if m != nil {
		return m.NbImages
	}
	return 0
}

func (m *PurgeNodeAnswer) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type PurgeNodesAnswers struct {
	Agents []*PurgeNodeAnswer `protobuf:"bytes,1,rep,name=agents" json:"agents,omitempty"`
}

func (m *PurgeNodesAnswers) Reset()                    { *m = PurgeNodesAnswers{} }
func (m *PurgeNodesAnswers) String() string            { return proto.CompactTextString(m) }
func (*PurgeNodesAnswers) ProtoMessage()               {}
func (*PurgeNodesAnswers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *PurgeNodesAnswers) GetAgents() []*PurgeNodeAnswer {
	if m != nil {
		return m.Agents
	}
	return nil
}

type AmpRequest struct {
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	Silence  bool   `protobuf:"varint,2,opt,name=silence" json:"silence,omitempty"`
	Verbose  bool   `protobuf:"varint,3,opt,name=verbose" json:"verbose,omitempty"`
	Force    bool   `protobuf:"varint,4,opt,name=force" json:"force,omitempty"`
	Local    bool   `protobuf:"varint,5,opt,name=local" json:"local,omitempty"`
	Node     string `protobuf:"bytes,6,opt,name=node" json:"node,omitempty"`
}

func (m *AmpRequest) Reset()                    { *m = AmpRequest{} }
func (m *AmpRequest) String() string            { return proto.CompactTextString(m) }
func (*AmpRequest) ProtoMessage()               {}
func (*AmpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *AmpRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *AmpRequest) GetSilence() bool {
	if m != nil {
		return m.Silence
	}
	return false
}

func (m *AmpRequest) GetVerbose() bool {
	if m != nil {
		return m.Verbose
	}
	return false
}

func (m *AmpRequest) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

func (m *AmpRequest) GetLocal() bool {
	if m != nil {
		return m.Local
	}
	return false
}

func (m *AmpRequest) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

type AmpRet struct {
}

func (m *AmpRet) Reset()                    { *m = AmpRet{} }
func (m *AmpRet) String() string            { return proto.CompactTextString(m) }
func (*AmpRet) ProtoMessage()               {}
func (*AmpRet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type TypedOutput struct {
	Output     string `protobuf:"bytes,1,opt,name=output" json:"output,omitempty"`
	OutputType int32  `protobuf:"varint,2,opt,name=output_type,json=outputType" json:"output_type,omitempty"`
}

func (m *TypedOutput) Reset()                    { *m = TypedOutput{} }
func (m *TypedOutput) String() string            { return proto.CompactTextString(m) }
func (*TypedOutput) ProtoMessage()               {}
func (*TypedOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *TypedOutput) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *TypedOutput) GetOutputType() int32 {
	if m != nil {
		return m.OutputType
	}
	return 0
}

type AmpMonitorAnswers struct {
	Outputs []*TypedOutput `protobuf:"bytes,1,rep,name=outputs" json:"outputs,omitempty"`
}

func (m *AmpMonitorAnswers) Reset()                    { *m = AmpMonitorAnswers{} }
func (m *AmpMonitorAnswers) String() string            { return proto.CompactTextString(m) }
func (*AmpMonitorAnswers) ProtoMessage()               {}
func (*AmpMonitorAnswers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *AmpMonitorAnswers) GetOutputs() []*TypedOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

type AmpStatusAnswer struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *AmpStatusAnswer) Reset()                    { *m = AmpStatusAnswer{} }
func (m *AmpStatusAnswer) String() string            { return proto.CompactTextString(m) }
func (*AmpStatusAnswer) ProtoMessage()               {}
func (*AmpStatusAnswer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *AmpStatusAnswer) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*ClientMes)(nil), "servergrpc.ClientMes")
	proto.RegisterType((*RegisterRequest)(nil), "servergrpc.RegisterRequest")
	proto.RegisterType((*AgentHealthRequest)(nil), "servergrpc.AgentHealthRequest")
	proto.RegisterType((*ServerRet)(nil), "servergrpc.ServerRet")
	proto.RegisterType((*ServerRets)(nil), "servergrpc.ServerRets")
	proto.RegisterType((*GetNodesInfoRequest)(nil), "servergrpc.GetNodesInfoRequest")
	proto.RegisterType((*NodesInfo)(nil), "servergrpc.NodesInfo")
	proto.RegisterType((*NodeInfo)(nil), "servergrpc.NodeInfo")
	proto.RegisterType((*PurgeNodesRequest)(nil), "servergrpc.PurgeNodesRequest")
	proto.RegisterType((*PurgeNodeAnswer)(nil), "servergrpc.PurgeNodeAnswer")
	proto.RegisterType((*PurgeNodesAnswers)(nil), "servergrpc.PurgeNodesAnswers")
	proto.RegisterType((*AmpRequest)(nil), "servergrpc.AmpRequest")
	proto.RegisterType((*AmpRet)(nil), "servergrpc.AmpRet")
	proto.RegisterType((*TypedOutput)(nil), "servergrpc.TypedOutput")
	proto.RegisterType((*AmpMonitorAnswers)(nil), "servergrpc.AmpMonitorAnswers")
	proto.RegisterType((*AmpStatusAnswer)(nil), "servergrpc.AmpStatusAnswer")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ClusterServerService service

type ClusterServerServiceClient interface {
	GetClientStream(ctx context.Context, opts ...grpc.CallOption) (ClusterServerService_GetClientStreamClient, error)
	RegisterAgent(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*ServerRet, error)
	AgentHealth(ctx context.Context, in *AgentHealthRequest, opts ...grpc.CallOption) (*ServerRet, error)
	GetNodesInfo(ctx context.Context, in *GetNodesInfoRequest, opts ...grpc.CallOption) (*NodesInfo, error)
	PurgeNodes(ctx context.Context, in *PurgeNodesRequest, opts ...grpc.CallOption) (*PurgeNodesAnswers, error)
}

type clusterServerServiceClient struct {
	cc *grpc.ClientConn
}

func NewClusterServerServiceClient(cc *grpc.ClientConn) ClusterServerServiceClient {
	return &clusterServerServiceClient{cc}
}

func (c *clusterServerServiceClient) GetClientStream(ctx context.Context, opts ...grpc.CallOption) (ClusterServerService_GetClientStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ClusterServerService_serviceDesc.Streams[0], c.cc, "/servergrpc.ClusterServerService/GetClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &clusterServerServiceGetClientStreamClient{stream}
	return x, nil
}

type ClusterServerService_GetClientStreamClient interface {
	Send(*ClientMes) error
	Recv() (*ClientMes, error)
	grpc.ClientStream
}

type clusterServerServiceGetClientStreamClient struct {
	grpc.ClientStream
}

func (x *clusterServerServiceGetClientStreamClient) Send(m *ClientMes) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clusterServerServiceGetClientStreamClient) Recv() (*ClientMes, error) {
	m := new(ClientMes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clusterServerServiceClient) RegisterAgent(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*ServerRet, error) {
	out := new(ServerRet)
	err := grpc.Invoke(ctx, "/servergrpc.ClusterServerService/RegisterAgent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServerServiceClient) AgentHealth(ctx context.Context, in *AgentHealthRequest, opts ...grpc.CallOption) (*ServerRet, error) {
	out := new(ServerRet)
	err := grpc.Invoke(ctx, "/servergrpc.ClusterServerService/AgentHealth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServerServiceClient) GetNodesInfo(ctx context.Context, in *GetNodesInfoRequest, opts ...grpc.CallOption) (*NodesInfo, error) {
	out := new(NodesInfo)
	err := grpc.Invoke(ctx, "/servergrpc.ClusterServerService/GetNodesInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServerServiceClient) PurgeNodes(ctx context.Context, in *PurgeNodesRequest, opts ...grpc.CallOption) (*PurgeNodesAnswers, error) {
	out := new(PurgeNodesAnswers)
	err := grpc.Invoke(ctx, "/servergrpc.ClusterServerService/PurgeNodes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ClusterServerService service

type ClusterServerServiceServer interface {
	GetClientStream(ClusterServerService_GetClientStreamServer) error
	RegisterAgent(context.Context, *RegisterRequest) (*ServerRet, error)
	AgentHealth(context.Context, *AgentHealthRequest) (*ServerRet, error)
	GetNodesInfo(context.Context, *GetNodesInfoRequest) (*NodesInfo, error)
	PurgeNodes(context.Context, *PurgeNodesRequest) (*PurgeNodesAnswers, error)
}

func RegisterClusterServerServiceServer(s *grpc.Server, srv ClusterServerServiceServer) {
	s.RegisterService(&_ClusterServerService_serviceDesc, srv)
}

func _ClusterServerService_GetClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClusterServerServiceServer).GetClientStream(&clusterServerServiceGetClientStreamServer{stream})
}

type ClusterServerService_GetClientStreamServer interface {
	Send(*ClientMes) error
	Recv() (*ClientMes, error)
	grpc.ServerStream
}

type clusterServerServiceGetClientStreamServer struct {
	grpc.ServerStream
}

func (x *clusterServerServiceGetClientStreamServer) Send(m *ClientMes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clusterServerServiceGetClientStreamServer) Recv() (*ClientMes, error) {
	m := new(ClientMes)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ClusterServerService_RegisterAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServerServiceServer).RegisterAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servergrpc.ClusterServerService/RegisterAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServerServiceServer).RegisterAgent(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterServerService_AgentHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentHealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServerServiceServer).AgentHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servergrpc.ClusterServerService/AgentHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServerServiceServer).AgentHealth(ctx, req.(*AgentHealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterServerService_GetNodesInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodesInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServerServiceServer).GetNodesInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servergrpc.ClusterServerService/GetNodesInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServerServiceServer).GetNodesInfo(ctx, req.(*GetNodesInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterServerService_PurgeNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurgeNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServerServiceServer).PurgeNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servergrpc.ClusterServerService/PurgeNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServerServiceServer).PurgeNodes(ctx, req.(*PurgeNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClusterServerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "servergrpc.ClusterServerService",
	HandlerType: (*ClusterServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterAgent",
			Handler:    _ClusterServerService_RegisterAgent_Handler,
		},
		{
			MethodName: "AgentHealth",
			Handler:    _ClusterServerService_AgentHealth_Handler,
		},
		{
			MethodName: "GetNodesInfo",
			Handler:    _ClusterServerService_GetNodesInfo_Handler,
		},
		{
			MethodName: "PurgeNodes",
			Handler:    _ClusterServerService_PurgeNodes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetClientStream",
			Handler:       _ClusterServerService_GetClientStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/appcelerator/amp/cmd/adm-server/servergrpc/cluster-server.proto",
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/cmd/adm-server/servergrpc/cluster-server.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 1045 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0x4b, 0x6f, 0xdc, 0x36,
	0x10, 0xb6, 0xec, 0x7d, 0x68, 0x67, 0xfd, 0xa4, 0xed, 0x46, 0xdd, 0x24, 0x8d, 0xc1, 0xb6, 0x80,
	0xd3, 0xa2, 0xde, 0xd4, 0x39, 0xe4, 0xd2, 0xcb, 0xc2, 0x88, 0x63, 0x17, 0xcd, 0x03, 0x72, 0x91,
	0x43, 0x2f, 0x0b, 0xae, 0x34, 0x96, 0x85, 0x4a, 0xa4, 0x4a, 0x52, 0x2e, 0x7c, 0xef, 0x6f, 0xe8,
	0xad, 0xd7, 0xfe, 0x8e, 0xfe, 0xb4, 0x82, 0xa4, 0xa4, 0x95, 0xfc, 0x08, 0x8a, 0x5e, 0x6c, 0x7e,
	0x33, 0x43, 0xf1, 0xe3, 0xcc, 0xc7, 0x99, 0x85, 0x1f, 0x93, 0x54, 0x5f, 0x95, 0x8b, 0xa3, 0x48,
	0xe4, 0x53, 0x56, 0x14, 0x11, 0x66, 0x28, 0x99, 0x16, 0x72, 0xca, 0xf2, 0x62, 0x1a, 0xe5, 0xf1,
	0x94, 0xc5, 0xf9, 0x77, 0x0a, 0xe5, 0x35, 0xca, 0xa9, 0xfb, 0x97, 0xc8, 0x22, 0x9a, 0x46, 0x59,
	0xa9, 0x34, 0xca, 0xca, 0x73, 0x54, 0x48, 0xa1, 0x05, 0x81, 0x65, 0xc0, 0xe4, 0x49, 0x22, 0x44,
	0x92, 0xe1, 0x94, 0x15, 0xe9, 0x94, 0x71, 0x2e, 0x34, 0xd3, 0xa9, 0xe0, 0xca, 0x45, 0xd2, 0x12,
	0x46, 0x27, 0x59, 0x8a, 0x5c, 0xbf, 0x45, 0x45, 0x1e, 0xc3, 0x28, 0xb2, 0x60, 0x9e, 0xc6, 0x81,
	0x77, 0xe0, 0x1d, 0x8e, 0x42, 0xdf, 0x19, 0xce, 0x63, 0x32, 0x01, 0xff, 0xb2, 0xe4, 0x91, 0xd9,
	0x1c, 0xac, 0x3a, 0x5f, 0x8d, 0xc9, 0x14, 0x06, 0xa2, 0xd4, 0x45, 0xa9, 0x83, 0xb5, 0x03, 0xef,
	0x70, 0x7c, 0xfc, 0xe8, 0x68, 0x49, 0xe0, 0xe8, 0xe7, 0x9b, 0x02, 0xe3, 0xf7, 0xd6, 0x1d, 0x56,
	0x61, 0xb4, 0x80, 0xad, 0x10, 0x93, 0xd4, 0x30, 0x0f, 0xf1, 0xb7, 0x12, 0x95, 0x26, 0x9b, 0xb0,
	0xda, 0x9c, 0xba, 0x9a, 0xc6, 0xe4, 0x11, 0x0c, 0xb9, 0x88, 0xd1, 0x50, 0x71, 0xc7, 0x0d, 0x0c,
	0x74, 0x44, 0xae, 0x84, 0xd2, 0x9c, 0xe5, 0x68, 0x8f, 0x1b, 0x85, 0x0d, 0x26, 0x01, 0x0c, 0x59,
	0x1c, 0x4b, 0x54, 0x2a, 0xe8, 0x59, 0x57, 0x0d, 0xe9, 0x57, 0x40, 0x66, 0x09, 0x72, 0x7d, 0x86,
	0x2c, 0xd3, 0x57, 0x0f, 0x1c, 0x4a, 0x7f, 0x80, 0xd1, 0x85, 0x65, 0x1e, 0xa2, 0x26, 0x9f, 0x83,
	0xcf, 0x92, 0x4e, 0x36, 0x86, 0x16, 0x9f, 0xc7, 0x64, 0x0f, 0xfa, 0x28, 0xa5, 0x90, 0x15, 0x35,
	0x07, 0xe8, 0x2b, 0x80, 0x66, 0xb7, 0x22, 0xcf, 0xa1, 0x27, 0x51, 0xab, 0xc0, 0x3b, 0x58, 0x3b,
	0x1c, 0x1f, 0xef, 0xb7, 0x53, 0xd2, 0x44, 0x85, 0x36, 0x84, 0xfe, 0x02, 0xbb, 0x6f, 0x50, 0xbf,
	0x13, 0x31, 0xaa, 0x73, 0x7e, 0x29, 0x6a, 0x76, 0x9f, 0xac, 0x07, 0x81, 0x9e, 0x49, 0x48, 0xc5,
	0xc0, 0xae, 0x8d, 0x2d, 0x17, 0xd2, 0xa5, 0xc5, 0x0f, 0xed, 0x9a, 0xbe, 0x82, 0x51, 0xf3, 0x61,
	0xf2, 0x0d, 0xf4, 0x4d, 0x60, 0x4d, 0x6a, 0xaf, 0x4d, 0xca, 0x44, 0xd9, 0xd3, 0x5d, 0x08, 0xfd,
	0xb3, 0x0f, 0x7e, 0x6d, 0xbb, 0x53, 0x1d, 0x02, 0x3d, 0x29, 0xb2, 0xe6, 0x74, 0xb3, 0x26, 0x14,
	0xd6, 0xd9, 0x35, 0x4b, 0x33, 0xb6, 0x48, 0xb3, 0x54, 0xdf, 0x54, 0xc5, 0xe9, 0xd8, 0x3a, 0xc5,
	0xeb, 0xdd, 0x2a, 0xde, 0xb7, 0xb0, 0x63, 0xd6, 0x73, 0x26, 0xa3, 0xab, 0x54, 0x63, 0xa4, 0x4b,
	0x89, 0x41, 0xdf, 0x06, 0x6d, 0x1b, 0xc7, 0xac, 0x65, 0x37, 0xf2, 0xb0, 0xc1, 0x42, 0x05, 0x03,
	0x27, 0x0f, 0x03, 0xdf, 0x2b, 0xb2, 0x0d, 0x6b, 0x51, 0x51, 0x06, 0xc3, 0x03, 0xef, 0x70, 0x2d,
	0x34, 0x4b, 0xf2, 0x19, 0x0c, 0x72, 0xcc, 0x85, 0xbc, 0x09, 0x7c, 0x6b, 0xac, 0x10, 0xf9, 0x1a,
	0x36, 0x63, 0x11, 0xfd, 0x8a, 0x72, 0x7e, 0x8d, 0x52, 0x19, 0x5d, 0x8f, 0xec, 0x97, 0x36, 0x9c,
	0xf5, 0xa3, 0x33, 0x9a, 0xed, 0x4a, 0x33, 0x5d, 0xaa, 0x00, 0xdc, 0x41, 0x0e, 0x19, 0x7b, 0x86,
	0x2c, 0x46, 0x19, 0x8c, 0x6d, 0xba, 0x2b, 0x64, 0xd2, 0x20, 0x91, 0x45, 0x57, 0x75, 0x1a, 0xd6,
	0x5d, 0x1a, 0xda, 0xb6, 0xb6, 0x4e, 0x37, 0x3a, 0x3a, 0x25, 0x5f, 0xc2, 0x06, 0x5f, 0xcc, 0x23,
	0xc1, 0x35, 0x4b, 0x39, 0x4a, 0x15, 0x6c, 0x5a, 0xce, 0xeb, 0x7c, 0x71, 0xd2, 0xd8, 0xc8, 0x31,
	0xec, 0x77, 0x82, 0xe6, 0xb2, 0xe4, 0x3c, 0xe5, 0x49, 0xb0, 0x65, 0x83, 0x77, 0xdb, 0xc1, 0xa1,
	0x73, 0x91, 0x17, 0xb0, 0xd7, 0xdd, 0x53, 0xb0, 0x52, 0x61, 0x1c, 0x6c, 0xdb, 0x2d, 0xa4, 0xbd,
	0xe5, 0x83, 0xf5, 0xdc, 0x3d, 0x45, 0x69, 0x51, 0x14, 0x18, 0x07, 0x3b, 0x77, 0x4f, 0xb9, 0x70,
	0x2e, 0x93, 0x94, 0x34, 0x67, 0x09, 0xaa, 0x80, 0xb8, 0x5c, 0x3b, 0xb4, 0x7c, 0x30, 0xbb, 0xad,
	0x07, 0x63, 0xd3, 0xe0, 0x5e, 0x54, 0xb0, 0xd7, 0x7d, 0x60, 0x13, 0xf0, 0x8d, 0x0a, 0xdf, 0x19,
	0x9d, 0xec, 0x3b, 0x9d, 0xd4, 0x98, 0xfe, 0xed, 0xc1, 0xce, 0x87, 0x52, 0x26, 0x68, 0x75, 0xfd,
	0xbf, 0x1f, 0xcb, 0x13, 0x18, 0x35, 0x77, 0xab, 0x5e, 0xcc, 0xd2, 0x60, 0x2e, 0x72, 0x2d, 0xb2,
	0xb2, 0x92, 0xa9, 0x1f, 0x56, 0xc8, 0x5c, 0xc4, 0x5e, 0xc9, 0x0a, 0xd3, 0x0f, 0x1d, 0x30, 0xd6,
	0x4b, 0x21, 0x23, 0xb4, 0x5a, 0xf4, 0x43, 0x07, 0xe8, 0x3f, 0x1e, 0x6c, 0x35, 0x44, 0x67, 0x5c,
	0xfd, 0x8e, 0xf2, 0xd3, 0x34, 0xdb, 0x1d, 0x67, 0xb5, 0x9b, 0x90, 0x3b, 0xba, 0x30, 0x8c, 0xfb,
	0xb7, 0x74, 0xf1, 0x14, 0x80, 0x2f, 0xe6, 0x8e, 0xa9, 0xeb, 0x80, 0xfd, 0x70, 0xc4, 0x17, 0x1f,
	0x9d, 0xc1, 0x9c, 0xcd, 0x17, 0xf3, 0xaa, 0x3e, 0x7d, 0xeb, 0xf5, 0xf9, 0xe2, 0xfc, 0x56, 0x85,
	0x06, 0xed, 0x96, 0x76, 0xd6, 0x4e, 0xb5, 0xbb, 0x82, 0x22, 0x2f, 0x61, 0x60, 0x69, 0xd5, 0x6d,
	0xe4, 0x71, 0xbb, 0x8d, 0xdc, 0xba, 0x70, 0x58, 0x85, 0xd2, 0xbf, 0x3c, 0x80, 0x59, 0x5e, 0xfc,
	0xa7, 0x72, 0x05, 0x30, 0x54, 0x69, 0x86, 0x3c, 0x72, 0x15, 0xf3, 0xc3, 0x1a, 0x1a, 0xcf, 0x35,
	0xca, 0x85, 0x50, 0x75, 0x93, 0xab, 0xe1, 0xb2, 0x04, 0xbd, 0x56, 0x09, 0x8c, 0x35, 0x13, 0x11,
	0xcb, 0xea, 0x72, 0x59, 0xd0, 0xc8, 0x61, 0xb0, 0x94, 0x03, 0xf5, 0x61, 0x60, 0xe9, 0x69, 0x7a,
	0x0a, 0xe3, 0xd6, 0xcc, 0x32, 0x4a, 0xa8, 0x86, 0x9b, 0xa3, 0x59, 0x21, 0xf2, 0x0c, 0xc6, 0x6e,
	0x35, 0xd7, 0x37, 0x85, 0x23, 0xda, 0x0f, 0xc1, 0x99, 0xcc, 0x7e, 0x7a, 0x0a, 0x3b, 0xb3, 0xbc,
	0x78, 0x2b, 0x78, 0xaa, 0x85, 0xac, 0x73, 0xf7, 0x3d, 0x0c, 0x5d, 0x48, 0x9d, 0xbc, 0x07, 0x67,
	0x65, 0x1d, 0x47, 0x9f, 0xc3, 0xd6, 0x2c, 0x2f, 0x2e, 0x6c, 0xd7, 0xa9, 0x54, 0xb4, 0xec, 0x49,
	0x5e, 0xbb, 0x27, 0x1d, 0xff, 0xb1, 0x06, 0x7b, 0x27, 0xee, 0x17, 0x81, 0x9b, 0x31, 0xe6, 0x6f,
	0x1a, 0x21, 0x79, 0x0d, 0x5b, 0x6f, 0x50, 0xbb, 0x51, 0x7f, 0xa1, 0x25, 0xb2, 0x9c, 0x74, 0x26,
	0x52, 0xf3, 0x23, 0x60, 0x72, 0xbf, 0x99, 0xae, 0x1c, 0x7a, 0x2f, 0x3c, 0xf2, 0x1a, 0x36, 0xea,
	0xb9, 0x6d, 0xa7, 0x29, 0xe9, 0x94, 0xfe, 0xd6, 0x48, 0x9f, 0xdc, 0x3f, 0xf3, 0xe8, 0x0a, 0x39,
	0x85, 0x71, 0x6b, 0x18, 0x93, 0x2f, 0xda, 0x71, 0x77, 0xa7, 0xf4, 0xc3, 0xdf, 0x39, 0x83, 0xf5,
	0xf6, 0xdc, 0x24, 0xcf, 0xda, 0x81, 0xf7, 0x4c, 0xd4, 0xee, 0x97, 0x1a, 0x2f, 0x5d, 0x21, 0x3f,
	0x01, 0x2c, 0x75, 0x4e, 0x9e, 0xde, 0x2b, 0xe8, 0xba, 0xd5, 0x4c, 0x1e, 0x70, 0x57, 0x25, 0xa6,
	0x2b, 0x8b, 0x81, 0xfd, 0x71, 0xf5, 0xf2, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x7a, 0x48,
	0xaa, 0xd4, 0x09, 0x00, 0x00,
}

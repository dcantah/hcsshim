// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/Microsoft/hcsshim/cmd/ncproxy/ncproxygrpc/networkconfigproxy.proto

package ncproxygrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RegisterNetworkConfigAgentRequest struct {
	AgentAddress         string   `protobuf:"bytes,1,opt,name=agent_address,json=agentAddress,proto3" json:"agent_address,omitempty"`
	NetworkID            string   `protobuf:"bytes,2,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterNetworkConfigAgentRequest) Reset()      { *m = RegisterNetworkConfigAgentRequest{} }
func (*RegisterNetworkConfigAgentRequest) ProtoMessage() {}
func (*RegisterNetworkConfigAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4dbe7e533383a60, []int{0}
}
func (m *RegisterNetworkConfigAgentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterNetworkConfigAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterNetworkConfigAgentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterNetworkConfigAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterNetworkConfigAgentRequest.Merge(m, src)
}
func (m *RegisterNetworkConfigAgentRequest) XXX_Size() int {
	return m.Size()
}
func (m *RegisterNetworkConfigAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterNetworkConfigAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterNetworkConfigAgentRequest proto.InternalMessageInfo

type RegisterNetworkConfigAgentResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterNetworkConfigAgentResponse) Reset()      { *m = RegisterNetworkConfigAgentResponse{} }
func (*RegisterNetworkConfigAgentResponse) ProtoMessage() {}
func (*RegisterNetworkConfigAgentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4dbe7e533383a60, []int{1}
}
func (m *RegisterNetworkConfigAgentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterNetworkConfigAgentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterNetworkConfigAgentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterNetworkConfigAgentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterNetworkConfigAgentResponse.Merge(m, src)
}
func (m *RegisterNetworkConfigAgentResponse) XXX_Size() int {
	return m.Size()
}
func (m *RegisterNetworkConfigAgentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterNetworkConfigAgentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterNetworkConfigAgentResponse proto.InternalMessageInfo

type AddNICRequest struct {
	NamespaceID          string   `protobuf:"bytes,1,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	NicID                string   `protobuf:"bytes,2,opt,name=nic_id,json=nicId,proto3" json:"nic_id,omitempty"`
	EndpointID           string   `protobuf:"bytes,3,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddNICRequest) Reset()      { *m = AddNICRequest{} }
func (*AddNICRequest) ProtoMessage() {}
func (*AddNICRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4dbe7e533383a60, []int{2}
}
func (m *AddNICRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddNICRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddNICRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddNICRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddNICRequest.Merge(m, src)
}
func (m *AddNICRequest) XXX_Size() int {
	return m.Size()
}
func (m *AddNICRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddNICRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddNICRequest proto.InternalMessageInfo

type AddNICResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddNICResponse) Reset()      { *m = AddNICResponse{} }
func (*AddNICResponse) ProtoMessage() {}
func (*AddNICResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4dbe7e533383a60, []int{3}
}
func (m *AddNICResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddNICResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddNICResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddNICResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddNICResponse.Merge(m, src)
}
func (m *AddNICResponse) XXX_Size() int {
	return m.Size()
}
func (m *AddNICResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddNICResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddNICResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegisterNetworkConfigAgentRequest)(nil), "RegisterNetworkConfigAgentRequest")
	proto.RegisterType((*RegisterNetworkConfigAgentResponse)(nil), "RegisterNetworkConfigAgentResponse")
	proto.RegisterType((*AddNICRequest)(nil), "AddNICRequest")
	proto.RegisterType((*AddNICResponse)(nil), "AddNICResponse")
}

func init() {
	proto.RegisterFile("github.com/Microsoft/hcsshim/cmd/ncproxy/ncproxygrpc/networkconfigproxy.proto", fileDescriptor_b4dbe7e533383a60)
}

var fileDescriptor_b4dbe7e533383a60 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbb, 0x6e, 0xd4, 0x40,
	0x14, 0x86, 0x3d, 0xa0, 0xac, 0xe4, 0xb3, 0xd9, 0x0d, 0x1a, 0x51, 0x44, 0x2e, 0xec, 0xe0, 0x50,
	0x20, 0x81, 0x6c, 0x29, 0x94, 0x34, 0xec, 0xc6, 0x14, 0x53, 0xc4, 0x42, 0xae, 0x50, 0x9a, 0xc8,
	0x99, 0x99, 0xcc, 0x8e, 0x90, 0x67, 0x8c, 0x67, 0xc2, 0xa5, 0xe3, 0x29, 0x78, 0x08, 0x9e, 0x64,
	0x4b, 0x4a, 0x2a, 0x8b, 0x9d, 0x27, 0x41, 0xbe, 0x01, 0x5b, 0x70, 0x11, 0x95, 0x8f, 0x7f, 0xfd,
	0x67, 0xce, 0x77, 0x7e, 0x1d, 0xb8, 0x10, 0xd2, 0x6e, 0x6e, 0xaf, 0x13, 0xaa, 0xab, 0xf4, 0x42,
	0xd2, 0x46, 0x1b, 0x7d, 0x63, 0xd3, 0x0d, 0x35, 0x66, 0x23, 0xab, 0x94, 0x56, 0x2c, 0x55, 0xb4,
	0x6e, 0xf4, 0xfb, 0x0f, 0xd3, 0x57, 0x34, 0x35, 0x4d, 0x15, 0xb7, 0xef, 0x74, 0xf3, 0x9a, 0x6a,
	0x75, 0x23, 0x45, 0x2f, 0x27, 0x75, 0xa3, 0xad, 0x0e, 0xee, 0x0b, 0x2d, 0x74, 0x5f, 0xa6, 0x5d,
	0x35, 0xa8, 0xf1, 0x5b, 0x78, 0x50, 0x70, 0x21, 0x8d, 0xe5, 0x4d, 0x3e, 0x74, 0x9e, 0xf7, 0x9d,
	0x2b, 0xc1, 0x95, 0x2d, 0xf8, 0x9b, 0x5b, 0x6e, 0x2c, 0x3e, 0x85, 0x45, 0xd9, 0xfd, 0x5f, 0x95,
	0x8c, 0x35, 0xdc, 0x98, 0x63, 0x74, 0x82, 0x1e, 0xf9, 0xc5, 0x61, 0x2f, 0xae, 0x06, 0x0d, 0x3f,
	0x01, 0x18, 0x67, 0x5f, 0x49, 0x76, 0x7c, 0xa7, 0x73, 0xac, 0x17, 0xae, 0x8d, 0xfc, 0xf1, 0x5d,
	0x92, 0x15, 0xfe, 0x68, 0x20, 0x2c, 0x7e, 0x08, 0xf1, 0x9f, 0xe6, 0x9a, 0x5a, 0x2b, 0xc3, 0xe3,
	0x4f, 0x08, 0x16, 0x2b, 0xc6, 0x72, 0x72, 0x3e, 0xa1, 0x9c, 0xc1, 0xa1, 0x2a, 0x2b, 0x6e, 0xea,
	0x92, 0xf2, 0x6e, 0x4e, 0x4f, 0xb2, 0x3e, 0x72, 0x6d, 0x34, 0xcf, 0x27, 0x9d, 0x64, 0xc5, 0xfc,
	0x87, 0x89, 0x30, 0x7c, 0x02, 0x33, 0x25, 0xe9, 0x4f, 0x2a, 0xdf, 0xb5, 0xd1, 0x41, 0x2e, 0x29,
	0xc9, 0x8a, 0x03, 0x25, 0x29, 0x61, 0x38, 0x85, 0x39, 0x57, 0xac, 0xd6, 0x52, 0xd9, 0xce, 0x76,
	0xb7, 0xb7, 0x2d, 0x5d, 0x1b, 0xc1, 0x8b, 0x51, 0x26, 0x59, 0x01, 0x93, 0x85, 0xb0, 0xf8, 0x1e,
	0x2c, 0x27, 0xae, 0x01, 0xf5, 0xec, 0x33, 0x02, 0xbc, 0xb7, 0xc9, 0xcb, 0x2e, 0x7b, 0x2c, 0x20,
	0xf8, 0xfd, 0x9e, 0x38, 0x4e, 0xfe, 0x1a, 0x7e, 0x70, 0x9a, 0xfc, 0x43, 0x50, 0x1e, 0x7e, 0x0c,
	0xb3, 0x81, 0x08, 0x2f, 0x93, 0xbd, 0xc8, 0x82, 0xa3, 0x64, 0x1f, 0x35, 0xf6, 0xd6, 0x97, 0xdb,
	0x5d, 0xe8, 0x7d, 0xdd, 0x85, 0xde, 0x47, 0x17, 0xa2, 0xad, 0x0b, 0xd1, 0x17, 0x17, 0xa2, 0x6f,
	0x2e, 0x44, 0x97, 0xcf, 0xff, 0xe7, 0xe8, 0x9e, 0xfd, 0x52, 0xbf, 0xf2, 0xae, 0x67, 0xfd, 0x69,
	0x3d, 0xfd, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x25, 0xfa, 0x9f, 0x26, 0xc1, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkConfigProxyClient is the client API for NetworkConfigProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkConfigProxyClient interface {
	RegisterNetworkConfigAgent(ctx context.Context, in *RegisterNetworkConfigAgentRequest, opts ...grpc.CallOption) (*RegisterNetworkConfigAgentResponse, error)
	AddNIC(ctx context.Context, in *AddNICRequest, opts ...grpc.CallOption) (*AddNICResponse, error)
}

type networkConfigProxyClient struct {
	cc *grpc.ClientConn
}

func NewNetworkConfigProxyClient(cc *grpc.ClientConn) NetworkConfigProxyClient {
	return &networkConfigProxyClient{cc}
}

func (c *networkConfigProxyClient) RegisterNetworkConfigAgent(ctx context.Context, in *RegisterNetworkConfigAgentRequest, opts ...grpc.CallOption) (*RegisterNetworkConfigAgentResponse, error) {
	out := new(RegisterNetworkConfigAgentResponse)
	err := c.cc.Invoke(ctx, "/NetworkConfigProxy/RegisterNetworkConfigAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkConfigProxyClient) AddNIC(ctx context.Context, in *AddNICRequest, opts ...grpc.CallOption) (*AddNICResponse, error) {
	out := new(AddNICResponse)
	err := c.cc.Invoke(ctx, "/NetworkConfigProxy/AddNIC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkConfigProxyServer is the server API for NetworkConfigProxy service.
type NetworkConfigProxyServer interface {
	RegisterNetworkConfigAgent(context.Context, *RegisterNetworkConfigAgentRequest) (*RegisterNetworkConfigAgentResponse, error)
	AddNIC(context.Context, *AddNICRequest) (*AddNICResponse, error)
}

func RegisterNetworkConfigProxyServer(s *grpc.Server, srv NetworkConfigProxyServer) {
	s.RegisterService(&_NetworkConfigProxy_serviceDesc, srv)
}

func _NetworkConfigProxy_RegisterNetworkConfigAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNetworkConfigAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkConfigProxyServer).RegisterNetworkConfigAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NetworkConfigProxy/RegisterNetworkConfigAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkConfigProxyServer).RegisterNetworkConfigAgent(ctx, req.(*RegisterNetworkConfigAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkConfigProxy_AddNIC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNICRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkConfigProxyServer).AddNIC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NetworkConfigProxy/AddNIC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkConfigProxyServer).AddNIC(ctx, req.(*AddNICRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkConfigProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "NetworkConfigProxy",
	HandlerType: (*NetworkConfigProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNetworkConfigAgent",
			Handler:    _NetworkConfigProxy_RegisterNetworkConfigAgent_Handler,
		},
		{
			MethodName: "AddNIC",
			Handler:    _NetworkConfigProxy_AddNIC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/Microsoft/hcsshim/cmd/ncproxy/ncproxygrpc/networkconfigproxy.proto",
}

func (m *RegisterNetworkConfigAgentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterNetworkConfigAgentRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AgentAddress) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNetworkconfigproxy(dAtA, i, uint64(len(m.AgentAddress)))
		i += copy(dAtA[i:], m.AgentAddress)
	}
	if len(m.NetworkID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNetworkconfigproxy(dAtA, i, uint64(len(m.NetworkID)))
		i += copy(dAtA[i:], m.NetworkID)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *RegisterNetworkConfigAgentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterNetworkConfigAgentResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *AddNICRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddNICRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.NamespaceID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNetworkconfigproxy(dAtA, i, uint64(len(m.NamespaceID)))
		i += copy(dAtA[i:], m.NamespaceID)
	}
	if len(m.NicID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNetworkconfigproxy(dAtA, i, uint64(len(m.NicID)))
		i += copy(dAtA[i:], m.NicID)
	}
	if len(m.EndpointID) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintNetworkconfigproxy(dAtA, i, uint64(len(m.EndpointID)))
		i += copy(dAtA[i:], m.EndpointID)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *AddNICResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddNICResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintNetworkconfigproxy(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RegisterNetworkConfigAgentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AgentAddress)
	if l > 0 {
		n += 1 + l + sovNetworkconfigproxy(uint64(l))
	}
	l = len(m.NetworkID)
	if l > 0 {
		n += 1 + l + sovNetworkconfigproxy(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RegisterNetworkConfigAgentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AddNICRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NamespaceID)
	if l > 0 {
		n += 1 + l + sovNetworkconfigproxy(uint64(l))
	}
	l = len(m.NicID)
	if l > 0 {
		n += 1 + l + sovNetworkconfigproxy(uint64(l))
	}
	l = len(m.EndpointID)
	if l > 0 {
		n += 1 + l + sovNetworkconfigproxy(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AddNICResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovNetworkconfigproxy(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNetworkconfigproxy(x uint64) (n int) {
	return sovNetworkconfigproxy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *RegisterNetworkConfigAgentRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RegisterNetworkConfigAgentRequest{`,
		`AgentAddress:` + fmt.Sprintf("%v", this.AgentAddress) + `,`,
		`NetworkID:` + fmt.Sprintf("%v", this.NetworkID) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RegisterNetworkConfigAgentResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RegisterNetworkConfigAgentResponse{`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AddNICRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AddNICRequest{`,
		`NamespaceID:` + fmt.Sprintf("%v", this.NamespaceID) + `,`,
		`NicID:` + fmt.Sprintf("%v", this.NicID) + `,`,
		`EndpointID:` + fmt.Sprintf("%v", this.EndpointID) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AddNICResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AddNICResponse{`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringNetworkconfigproxy(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *RegisterNetworkConfigAgentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkconfigproxy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterNetworkConfigAgentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterNetworkConfigAgentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AgentAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkconfigproxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AgentAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkconfigproxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NetworkID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkconfigproxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RegisterNetworkConfigAgentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkconfigproxy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterNetworkConfigAgentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterNetworkConfigAgentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkconfigproxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AddNICRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkconfigproxy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AddNICRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddNICRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NamespaceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkconfigproxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NamespaceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NicID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkconfigproxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NicID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndpointID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkconfigproxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndpointID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkconfigproxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AddNICResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkconfigproxy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AddNICResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddNICResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkconfigproxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetworkconfigproxy
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNetworkconfigproxy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNetworkconfigproxy
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNetworkconfigproxy
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNetworkconfigproxy
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthNetworkconfigproxy
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthNetworkconfigproxy
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNetworkconfigproxy
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipNetworkconfigproxy(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthNetworkconfigproxy
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthNetworkconfigproxy = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNetworkconfigproxy   = fmt.Errorf("proto: integer overflow")
)

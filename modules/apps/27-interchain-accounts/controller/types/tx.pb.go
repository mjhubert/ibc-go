// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/interchain_accounts/controller/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgRegisterInterchainAccount defines the payload for Msg/RegisterAccount
type MsgRegisterInterchainAccount struct {
	Owner        string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	ConnectionId string `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	Version      string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *MsgRegisterInterchainAccount) Reset()         { *m = MsgRegisterInterchainAccount{} }
func (m *MsgRegisterInterchainAccount) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterInterchainAccount) ProtoMessage()    {}
func (*MsgRegisterInterchainAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_7def041328c84a30, []int{0}
}
func (m *MsgRegisterInterchainAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterInterchainAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterInterchainAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterInterchainAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterInterchainAccount.Merge(m, src)
}
func (m *MsgRegisterInterchainAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterInterchainAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterInterchainAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterInterchainAccount proto.InternalMessageInfo

// MsgRegisterInterchainAccountResponse defines the response for Msg/RegisterAccount
type MsgRegisterInterchainAccountResponse struct {
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	PortId    string `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
}

func (m *MsgRegisterInterchainAccountResponse) Reset()         { *m = MsgRegisterInterchainAccountResponse{} }
func (m *MsgRegisterInterchainAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterInterchainAccountResponse) ProtoMessage()    {}
func (*MsgRegisterInterchainAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7def041328c84a30, []int{1}
}
func (m *MsgRegisterInterchainAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterInterchainAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterInterchainAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterInterchainAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterInterchainAccountResponse.Merge(m, src)
}
func (m *MsgRegisterInterchainAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterInterchainAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterInterchainAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterInterchainAccountResponse proto.InternalMessageInfo

func (m *MsgRegisterInterchainAccountResponse) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *MsgRegisterInterchainAccountResponse) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

// MsgSendTx defines the payload for Msg/SendTx
type MsgSendTx struct {
	Owner        string                            `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	ConnectionId string                            `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	PacketData   types.InterchainAccountPacketData `protobuf:"bytes,3,opt,name=packet_data,json=packetData,proto3" json:"packet_data"`
	// Relative timeout timestamp provided will be added to the current block time during transaction execution.
	// The timeout timestamp must be non-zero.
	RelativeTimeout uint64 `protobuf:"varint,4,opt,name=relative_timeout,json=relativeTimeout,proto3" json:"relative_timeout,omitempty"`
}

func (m *MsgSendTx) Reset()         { *m = MsgSendTx{} }
func (m *MsgSendTx) String() string { return proto.CompactTextString(m) }
func (*MsgSendTx) ProtoMessage()    {}
func (*MsgSendTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_7def041328c84a30, []int{2}
}
func (m *MsgSendTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendTx.Merge(m, src)
}
func (m *MsgSendTx) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendTx) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendTx.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendTx proto.InternalMessageInfo

// MsgSendTxResponse defines the response for MsgSendTx
type MsgSendTxResponse struct {
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *MsgSendTxResponse) Reset()         { *m = MsgSendTxResponse{} }
func (m *MsgSendTxResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendTxResponse) ProtoMessage()    {}
func (*MsgSendTxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7def041328c84a30, []int{3}
}
func (m *MsgSendTxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendTxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendTxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendTxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendTxResponse.Merge(m, src)
}
func (m *MsgSendTxResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendTxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendTxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendTxResponse proto.InternalMessageInfo

func (m *MsgSendTxResponse) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgRegisterInterchainAccount)(nil), "ibc.applications.interchain_accounts.controller.v1.MsgRegisterInterchainAccount")
	proto.RegisterType((*MsgRegisterInterchainAccountResponse)(nil), "ibc.applications.interchain_accounts.controller.v1.MsgRegisterInterchainAccountResponse")
	proto.RegisterType((*MsgSendTx)(nil), "ibc.applications.interchain_accounts.controller.v1.MsgSendTx")
	proto.RegisterType((*MsgSendTxResponse)(nil), "ibc.applications.interchain_accounts.controller.v1.MsgSendTxResponse")
}

func init() {
	proto.RegisterFile("ibc/applications/interchain_accounts/controller/v1/tx.proto", fileDescriptor_7def041328c84a30)
}

var fileDescriptor_7def041328c84a30 = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xf6, 0xb5, 0x69, 0x4a, 0xae, 0x20, 0xc0, 0xaa, 0xd4, 0x60, 0x81, 0x53, 0x05, 0x86, 0x52,
	0x29, 0x3e, 0x25, 0x20, 0x55, 0x02, 0x31, 0x50, 0x95, 0x21, 0x43, 0xa4, 0xc8, 0x74, 0x40, 0x0c,
	0x44, 0x97, 0xf3, 0x93, 0x73, 0xd4, 0xbe, 0x33, 0xbe, 0x8b, 0x29, 0x1b, 0x62, 0x81, 0x09, 0xf1,
	0x13, 0xfa, 0x13, 0xfa, 0x2f, 0xe8, 0xd8, 0x91, 0x09, 0xa1, 0x64, 0x08, 0x33, 0xbf, 0x00, 0xd9,
	0x4e, 0x1c, 0x24, 0x42, 0x55, 0x28, 0x9b, 0xdf, 0xf7, 0xf4, 0x7d, 0xfe, 0xbe, 0xf7, 0xee, 0x0e,
	0x3f, 0xe4, 0x7d, 0x46, 0x68, 0x14, 0x05, 0x9c, 0x51, 0xcd, 0xa5, 0x50, 0x84, 0x0b, 0x0d, 0x31,
	0x1b, 0x50, 0x2e, 0x7a, 0x94, 0x31, 0x39, 0x14, 0x5a, 0x11, 0x26, 0x85, 0x8e, 0x65, 0x10, 0x40,
	0x4c, 0x92, 0x26, 0xd1, 0x87, 0x4e, 0x14, 0x4b, 0x2d, 0xcd, 0x16, 0xef, 0x33, 0xe7, 0x57, 0xb2,
	0xb3, 0x80, 0xec, 0xcc, 0xc9, 0x4e, 0xd2, 0xb4, 0xd6, 0x7d, 0xe9, 0xcb, 0x8c, 0x4e, 0xd2, 0xaf,
	0x5c, 0xc9, 0xba, 0x7f, 0x2e, 0x1b, 0x49, 0x93, 0x44, 0x94, 0x1d, 0x80, 0x9e, 0xb2, 0x36, 0x98,
	0x54, 0xa1, 0x54, 0x24, 0x54, 0x7e, 0xda, 0x0b, 0x95, 0x9f, 0x37, 0xea, 0xef, 0x11, 0xbe, 0xd9,
	0x51, 0xbe, 0x0b, 0x3e, 0x57, 0x1a, 0xe2, 0x76, 0xa1, 0xf5, 0x38, 0x97, 0x32, 0xd7, 0xf1, 0x8a,
	0x7c, 0x2d, 0x20, 0xae, 0xa2, 0x4d, 0xb4, 0x55, 0x71, 0xf3, 0xc2, 0xbc, 0x8d, 0xaf, 0x30, 0x29,
	0x04, 0xb0, 0xd4, 0x42, 0x8f, 0x7b, 0xd5, 0xa5, 0xac, 0x7b, 0x79, 0x0e, 0xb6, 0x3d, 0xb3, 0x8a,
	0x57, 0x13, 0x88, 0x15, 0x97, 0xa2, 0xba, 0x9c, 0xb5, 0x67, 0xe5, 0x03, 0xf3, 0xc3, 0x51, 0xcd,
	0xf8, 0x7e, 0x54, 0x33, 0xde, 0x4d, 0x8e, 0xb7, 0x73, 0xc9, 0xfa, 0x0b, 0x7c, 0xe7, 0x2c, 0x23,
	0x2e, 0xa8, 0x48, 0x0a, 0x05, 0xe6, 0x2d, 0x8c, 0xd9, 0x80, 0x0a, 0x01, 0x41, 0xfa, 0xdf, 0xdc,
	0x55, 0x65, 0x8a, 0xb4, 0x3d, 0x73, 0x03, 0xaf, 0x46, 0x32, 0xd6, 0x73, 0x4f, 0xe5, 0xb4, 0x6c,
	0x7b, 0xf5, 0x1f, 0x08, 0x57, 0x3a, 0xca, 0x7f, 0x0a, 0xc2, 0xdb, 0x3f, 0xbc, 0x48, 0xac, 0x03,
	0xbc, 0x96, 0xcf, 0xb6, 0xe7, 0x51, 0x4d, 0xb3, 0x68, 0x6b, 0xad, 0x3d, 0xe7, 0x5c, 0x1b, 0x4e,
	0x9a, 0xce, 0x6f, 0xc9, 0xba, 0x99, 0xd8, 0x1e, 0xd5, 0x74, 0xb7, 0x74, 0xf2, 0xb5, 0x66, 0xb8,
	0x38, 0x2a, 0x10, 0xf3, 0x2e, 0xbe, 0x16, 0x43, 0x40, 0x35, 0x4f, 0xa0, 0xa7, 0x79, 0x08, 0x72,
	0xa8, 0xab, 0xa5, 0x4d, 0xb4, 0x55, 0x72, 0xaf, 0xce, 0xf0, 0xfd, 0x1c, 0x5e, 0x38, 0x54, 0x82,
	0xaf, 0x17, 0x99, 0x8b, 0x09, 0x5a, 0xf8, 0x92, 0x82, 0x57, 0x43, 0x10, 0x0c, 0xb2, 0xf8, 0x25,
	0xb7, 0xa8, 0x5b, 0x93, 0x25, 0xbc, 0xdc, 0x51, 0xbe, 0xf9, 0x19, 0xe1, 0x1b, 0x7f, 0x3e, 0x14,
	0x5d, 0xe7, 0xef, 0xcf, 0xb3, 0x73, 0xd6, 0x76, 0xad, 0x67, 0xff, 0x5b, 0xb1, 0x48, 0xfb, 0x11,
	0xe1, 0xf2, 0x74, 0xe9, 0x8f, 0xfe, 0xf1, 0x27, 0x39, 0xdd, 0x7a, 0x72, 0x21, 0xfa, 0xcc, 0x90,
	0xb5, 0xf2, 0x76, 0x72, 0xbc, 0x8d, 0x76, 0x5f, 0x9e, 0x8c, 0x6c, 0x74, 0x3a, 0xb2, 0xd1, 0xb7,
	0x91, 0x8d, 0x3e, 0x8d, 0x6d, 0xe3, 0x74, 0x6c, 0x1b, 0x5f, 0xc6, 0xb6, 0xf1, 0xbc, 0xeb, 0x73,
	0x3d, 0x18, 0xf6, 0x1d, 0x26, 0x43, 0x32, 0xbd, 0xb7, 0xbc, 0xcf, 0x1a, 0xbe, 0x24, 0xc9, 0x0e,
	0x09, 0xa5, 0x37, 0x0c, 0x40, 0xa5, 0x4f, 0x80, 0x22, 0xad, 0x9d, 0xc6, 0xdc, 0x41, 0x63, 0xd1,
	0x23, 0xa4, 0xdf, 0x44, 0xa0, 0xfa, 0xe5, 0xec, 0xb2, 0xdf, 0xfb, 0x19, 0x00, 0x00, 0xff, 0xff,
	0x77, 0x5e, 0x71, 0x7a, 0xc4, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// RegisterInterchainAccount defines a rpc handler for MsgRegisterInterchainAccount.
	RegisterInterchainAccount(ctx context.Context, in *MsgRegisterInterchainAccount, opts ...grpc.CallOption) (*MsgRegisterInterchainAccountResponse, error)
	// SendTx defines a rpc handler for MsgSendTx.
	SendTx(ctx context.Context, in *MsgSendTx, opts ...grpc.CallOption) (*MsgSendTxResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterInterchainAccount(ctx context.Context, in *MsgRegisterInterchainAccount, opts ...grpc.CallOption) (*MsgRegisterInterchainAccountResponse, error) {
	out := new(MsgRegisterInterchainAccountResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.interchain_accounts.controller.v1.Msg/RegisterInterchainAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SendTx(ctx context.Context, in *MsgSendTx, opts ...grpc.CallOption) (*MsgSendTxResponse, error) {
	out := new(MsgSendTxResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.interchain_accounts.controller.v1.Msg/SendTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// RegisterInterchainAccount defines a rpc handler for MsgRegisterInterchainAccount.
	RegisterInterchainAccount(context.Context, *MsgRegisterInterchainAccount) (*MsgRegisterInterchainAccountResponse, error)
	// SendTx defines a rpc handler for MsgSendTx.
	SendTx(context.Context, *MsgSendTx) (*MsgSendTxResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RegisterInterchainAccount(ctx context.Context, req *MsgRegisterInterchainAccount) (*MsgRegisterInterchainAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterInterchainAccount not implemented")
}
func (*UnimplementedMsgServer) SendTx(ctx context.Context, req *MsgSendTx) (*MsgSendTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTx not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RegisterInterchainAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterInterchainAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterInterchainAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.interchain_accounts.controller.v1.Msg/RegisterInterchainAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterInterchainAccount(ctx, req.(*MsgRegisterInterchainAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SendTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.interchain_accounts.controller.v1.Msg/SendTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendTx(ctx, req.(*MsgSendTx))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.applications.interchain_accounts.controller.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterInterchainAccount",
			Handler:    _Msg_RegisterInterchainAccount_Handler,
		},
		{
			MethodName: "SendTx",
			Handler:    _Msg_SendTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/applications/interchain_accounts/controller/v1/tx.proto",
}

func (m *MsgRegisterInterchainAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterInterchainAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterInterchainAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegisterInterchainAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterInterchainAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterInterchainAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RelativeTimeout != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.RelativeTimeout))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.PacketData.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendTxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendTxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendTxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRegisterInterchainAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRegisterInterchainAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSendTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.PacketData.Size()
	n += 1 + l + sovTx(uint64(l))
	if m.RelativeTimeout != 0 {
		n += 1 + sovTx(uint64(m.RelativeTimeout))
	}
	return n
}

func (m *MsgSendTxResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sequence != 0 {
		n += 1 + sovTx(uint64(m.Sequence))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRegisterInterchainAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRegisterInterchainAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterInterchainAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRegisterInterchainAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRegisterInterchainAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterInterchainAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSendTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSendTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PacketData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelativeTimeout", wireType)
			}
			m.RelativeTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RelativeTimeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSendTxResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSendTxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendTxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)

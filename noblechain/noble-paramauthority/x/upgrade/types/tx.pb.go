// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos2/upgrade/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_x_upgrade_types "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgSoftwareUpgrade is the Msg/SoftwareUpgrade request type.
//
// Since: cosmos-sdk 0.46
type MsgSoftwareUpgrade struct {
	// authority is the address of the governance account.
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// plan is the upgrade plan.
	Plan github_com_cosmos_cosmos_sdk_x_upgrade_types.Plan `protobuf:"bytes,2,opt,name=plan,proto3,casttype=github.com/cosmos/cosmos-sdk/x/upgrade/types.Plan" json:"plan"`
}

func (m *MsgSoftwareUpgrade) Reset()         { *m = MsgSoftwareUpgrade{} }
func (m *MsgSoftwareUpgrade) String() string { return proto.CompactTextString(m) }
func (*MsgSoftwareUpgrade) ProtoMessage()    {}
func (*MsgSoftwareUpgrade) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b0759bf79280790, []int{0}
}
func (m *MsgSoftwareUpgrade) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSoftwareUpgrade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSoftwareUpgrade.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSoftwareUpgrade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSoftwareUpgrade.Merge(m, src)
}
func (m *MsgSoftwareUpgrade) XXX_Size() int {
	return m.Size()
}
func (m *MsgSoftwareUpgrade) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSoftwareUpgrade.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSoftwareUpgrade proto.InternalMessageInfo

func (m *MsgSoftwareUpgrade) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgSoftwareUpgrade) GetPlan() github_com_cosmos_cosmos_sdk_x_upgrade_types.Plan {
	if m != nil {
		return m.Plan
	}
	return github_com_cosmos_cosmos_sdk_x_upgrade_types.Plan{}
}

// MsgSoftwareUpgradeResponse is the Msg/SoftwareUpgrade response type.
//
// Since: cosmos-sdk 0.46
type MsgSoftwareUpgradeResponse struct {
}

func (m *MsgSoftwareUpgradeResponse) Reset()         { *m = MsgSoftwareUpgradeResponse{} }
func (m *MsgSoftwareUpgradeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSoftwareUpgradeResponse) ProtoMessage()    {}
func (*MsgSoftwareUpgradeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b0759bf79280790, []int{1}
}
func (m *MsgSoftwareUpgradeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSoftwareUpgradeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSoftwareUpgradeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSoftwareUpgradeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSoftwareUpgradeResponse.Merge(m, src)
}
func (m *MsgSoftwareUpgradeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSoftwareUpgradeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSoftwareUpgradeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSoftwareUpgradeResponse proto.InternalMessageInfo

// MsgCancelUpgrade is the Msg/CancelUpgrade request type.
//
// Since: cosmos-sdk 0.46
type MsgCancelUpgrade struct {
	// authority is the address of the governance account.
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
}

func (m *MsgCancelUpgrade) Reset()         { *m = MsgCancelUpgrade{} }
func (m *MsgCancelUpgrade) String() string { return proto.CompactTextString(m) }
func (*MsgCancelUpgrade) ProtoMessage()    {}
func (*MsgCancelUpgrade) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b0759bf79280790, []int{2}
}
func (m *MsgCancelUpgrade) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCancelUpgrade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCancelUpgrade.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCancelUpgrade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCancelUpgrade.Merge(m, src)
}
func (m *MsgCancelUpgrade) XXX_Size() int {
	return m.Size()
}
func (m *MsgCancelUpgrade) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCancelUpgrade.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCancelUpgrade proto.InternalMessageInfo

func (m *MsgCancelUpgrade) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

// MsgCancelUpgradeResponse is the Msg/CancelUpgrade response type.
//
// Since: cosmos-sdk 0.46
type MsgCancelUpgradeResponse struct {
}

func (m *MsgCancelUpgradeResponse) Reset()         { *m = MsgCancelUpgradeResponse{} }
func (m *MsgCancelUpgradeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCancelUpgradeResponse) ProtoMessage()    {}
func (*MsgCancelUpgradeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b0759bf79280790, []int{3}
}
func (m *MsgCancelUpgradeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCancelUpgradeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCancelUpgradeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCancelUpgradeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCancelUpgradeResponse.Merge(m, src)
}
func (m *MsgCancelUpgradeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCancelUpgradeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCancelUpgradeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCancelUpgradeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSoftwareUpgrade)(nil), "cosmos2.upgrade.v1beta1.MsgSoftwareUpgrade")
	proto.RegisterType((*MsgSoftwareUpgradeResponse)(nil), "cosmos2.upgrade.v1beta1.MsgSoftwareUpgradeResponse")
	proto.RegisterType((*MsgCancelUpgrade)(nil), "cosmos2.upgrade.v1beta1.MsgCancelUpgrade")
	proto.RegisterType((*MsgCancelUpgradeResponse)(nil), "cosmos2.upgrade.v1beta1.MsgCancelUpgradeResponse")
}

func init() { proto.RegisterFile("cosmos2/upgrade/v1beta1/tx.proto", fileDescriptor_1b0759bf79280790) }

var fileDescriptor_1b0759bf79280790 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0x36, 0xd2, 0x2f, 0x2d, 0x48, 0x2f, 0x4a, 0x4c, 0x49, 0xd5, 0x2f, 0x33, 0x4c, 0x4a,
	0x2d, 0x49, 0x34, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x87, 0xaa,
	0xd0, 0x83, 0xaa, 0xd0, 0x83, 0xaa, 0x90, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xab, 0xd1, 0x07,
	0xb1, 0x20, 0xca, 0xa5, 0xb4, 0x20, 0xca, 0xe3, 0x21, 0x12, 0xe8, 0xa6, 0xc2, 0xcc, 0x00, 0xcb,
	0x2a, 0x2d, 0x65, 0xe4, 0x12, 0xf2, 0x2d, 0x4e, 0x0f, 0xce, 0x4f, 0x2b, 0x29, 0x4f, 0x2c, 0x4a,
	0x0d, 0x85, 0x48, 0x0a, 0xc9, 0x70, 0x71, 0x26, 0x96, 0x96, 0x64, 0xe4, 0x17, 0x65, 0x96, 0x54,
	0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x21, 0x04, 0x84, 0x0a, 0xb9, 0x58, 0x0a, 0x72, 0x12,
	0xf3, 0x24, 0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0x94, 0xf4, 0x90, 0xed, 0x43, 0x77, 0xa3, 0x5e,
	0x40, 0x4e, 0x62, 0x9e, 0x93, 0xe5, 0x89, 0x7b, 0xf2, 0x0c, 0xbf, 0xee, 0xc9, 0x1b, 0xa6, 0x67,
	0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x43, 0x74, 0x41, 0x29, 0xdd, 0xe2, 0x94,
	0x6c, 0xfd, 0x0a, 0xb8, 0x63, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0xc1, 0x5a, 0x83, 0xc0, 0x56, 0x29,
	0xc9, 0x70, 0x49, 0x61, 0x3a, 0x33, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0xc9, 0x80,
	0x4b, 0xc0, 0xb7, 0x38, 0xdd, 0x39, 0x31, 0x2f, 0x39, 0x35, 0x87, 0x28, 0x2f, 0x28, 0x49, 0x71,
	0x49, 0xa0, 0xeb, 0x80, 0x99, 0x66, 0xf4, 0x92, 0x91, 0x8b, 0xd9, 0xb7, 0x38, 0x5d, 0xa8, 0x98,
	0x8b, 0x1f, 0x3d, 0x5c, 0xb4, 0xf5, 0x70, 0x44, 0x85, 0x1e, 0xa6, 0xeb, 0xa4, 0x8c, 0x49, 0x50,
	0x0c, 0xb3, 0x5c, 0x28, 0x97, 0x8b, 0x17, 0xd5, 0x1f, 0x9a, 0xf8, 0x4c, 0x41, 0x51, 0x2a, 0x65,
	0x48, 0xb4, 0x52, 0x98, 0x75, 0x4e, 0x61, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8,
	0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7,
	0x10, 0x65, 0x83, 0x14, 0x55, 0xe5, 0x69, 0x49, 0x39, 0xf9, 0xc9, 0xd9, 0xc9, 0x19, 0x89, 0x99,
	0x79, 0xfa, 0x79, 0xf9, 0x49, 0x39, 0xa9, 0xba, 0x05, 0x89, 0x45, 0x89, 0xb9, 0xf0, 0xf0, 0x44,
	0x8f, 0xba, 0x24, 0x36, 0x70, 0x62, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x64, 0x14, 0x06,
	0xd8, 0xdd, 0x02, 0x00, 0x00,
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
	// SoftwareUpgrade is a governance operation for initiating a software upgrade.
	//
	// Since: cosmos-sdk 0.46
	SoftwareUpgrade(ctx context.Context, in *MsgSoftwareUpgrade, opts ...grpc.CallOption) (*MsgSoftwareUpgradeResponse, error)
	// CancelUpgrade is a governance operation for cancelling a previously
	// approvid software upgrade.
	//
	// Since: cosmos-sdk 0.46
	CancelUpgrade(ctx context.Context, in *MsgCancelUpgrade, opts ...grpc.CallOption) (*MsgCancelUpgradeResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SoftwareUpgrade(ctx context.Context, in *MsgSoftwareUpgrade, opts ...grpc.CallOption) (*MsgSoftwareUpgradeResponse, error) {
	out := new(MsgSoftwareUpgradeResponse)
	err := c.cc.Invoke(ctx, "/cosmos2.upgrade.v1beta1.Msg/SoftwareUpgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelUpgrade(ctx context.Context, in *MsgCancelUpgrade, opts ...grpc.CallOption) (*MsgCancelUpgradeResponse, error) {
	out := new(MsgCancelUpgradeResponse)
	err := c.cc.Invoke(ctx, "/cosmos2.upgrade.v1beta1.Msg/CancelUpgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SoftwareUpgrade is a governance operation for initiating a software upgrade.
	//
	// Since: cosmos-sdk 0.46
	SoftwareUpgrade(context.Context, *MsgSoftwareUpgrade) (*MsgSoftwareUpgradeResponse, error)
	// CancelUpgrade is a governance operation for cancelling a previously
	// approvid software upgrade.
	//
	// Since: cosmos-sdk 0.46
	CancelUpgrade(context.Context, *MsgCancelUpgrade) (*MsgCancelUpgradeResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SoftwareUpgrade(ctx context.Context, req *MsgSoftwareUpgrade) (*MsgSoftwareUpgradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SoftwareUpgrade not implemented")
}
func (*UnimplementedMsgServer) CancelUpgrade(ctx context.Context, req *MsgCancelUpgrade) (*MsgCancelUpgradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelUpgrade not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SoftwareUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSoftwareUpgrade)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SoftwareUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos2.upgrade.v1beta1.Msg/SoftwareUpgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SoftwareUpgrade(ctx, req.(*MsgSoftwareUpgrade))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelUpgrade)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos2.upgrade.v1beta1.Msg/CancelUpgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelUpgrade(ctx, req.(*MsgCancelUpgrade))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos2.upgrade.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SoftwareUpgrade",
			Handler:    _Msg_SoftwareUpgrade_Handler,
		},
		{
			MethodName: "CancelUpgrade",
			Handler:    _Msg_CancelUpgrade_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos2/upgrade/v1beta1/tx.proto",
}

func (m *MsgSoftwareUpgrade) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSoftwareUpgrade) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSoftwareUpgrade) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Plan.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSoftwareUpgradeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSoftwareUpgradeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSoftwareUpgradeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCancelUpgrade) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCancelUpgrade) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCancelUpgrade) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCancelUpgradeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCancelUpgradeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCancelUpgradeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *MsgSoftwareUpgrade) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Plan.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgSoftwareUpgradeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCancelUpgrade) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCancelUpgradeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSoftwareUpgrade) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSoftwareUpgrade: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSoftwareUpgrade: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
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
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plan", wireType)
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
			if err := m.Plan.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgSoftwareUpgradeResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSoftwareUpgradeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSoftwareUpgradeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *MsgCancelUpgrade) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCancelUpgrade: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCancelUpgrade: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
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
			m.Authority = string(dAtA[iNdEx:postIndex])
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
func (m *MsgCancelUpgradeResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCancelUpgradeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCancelUpgradeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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

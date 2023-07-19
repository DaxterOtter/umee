// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/ugov/v1/tx.proto

package ugov

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
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

// MsgGovUpdateMinGasPrice request type.
type MsgGovUpdateMinGasPrice struct {
	// authority must be the address of the governance account.
	Authority   string        `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	MinGasPrice types.DecCoin `protobuf:"bytes,2,opt,name=min_gas_price,json=minGasPrice,proto3" json:"min_gas_price"`
}

func (m *MsgGovUpdateMinGasPrice) Reset()      { *m = MsgGovUpdateMinGasPrice{} }
func (*MsgGovUpdateMinGasPrice) ProtoMessage() {}
func (*MsgGovUpdateMinGasPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ffc07de1c6ee91b, []int{0}
}
func (m *MsgGovUpdateMinGasPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGovUpdateMinGasPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGovUpdateMinGasPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGovUpdateMinGasPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGovUpdateMinGasPrice.Merge(m, src)
}
func (m *MsgGovUpdateMinGasPrice) XXX_Size() int {
	return m.Size()
}
func (m *MsgGovUpdateMinGasPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGovUpdateMinGasPrice.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGovUpdateMinGasPrice proto.InternalMessageInfo

// MsgGovUpdateMinGasPriceResponse response type.
type MsgGovUpdateMinGasPriceResponse struct {
}

func (m *MsgGovUpdateMinGasPriceResponse) Reset()         { *m = MsgGovUpdateMinGasPriceResponse{} }
func (m *MsgGovUpdateMinGasPriceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgGovUpdateMinGasPriceResponse) ProtoMessage()    {}
func (*MsgGovUpdateMinGasPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ffc07de1c6ee91b, []int{1}
}
func (m *MsgGovUpdateMinGasPriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGovUpdateMinGasPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGovUpdateMinGasPriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGovUpdateMinGasPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGovUpdateMinGasPriceResponse.Merge(m, src)
}
func (m *MsgGovUpdateMinGasPriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgGovUpdateMinGasPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGovUpdateMinGasPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGovUpdateMinGasPriceResponse proto.InternalMessageInfo

// MsgGovSetEmergencyGroup request type.
type MsgGovSetEmergencyGroup struct {
	// authority must be the address of the governance account.
	Authority      string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	EmergencyGroup string `protobuf:"bytes,2,opt,name=emergency_group,json=emergencyGroup,proto3" json:"emergency_group,omitempty"`
}

func (m *MsgGovSetEmergencyGroup) Reset()         { *m = MsgGovSetEmergencyGroup{} }
func (m *MsgGovSetEmergencyGroup) String() string { return proto.CompactTextString(m) }
func (*MsgGovSetEmergencyGroup) ProtoMessage()    {}
func (*MsgGovSetEmergencyGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ffc07de1c6ee91b, []int{2}
}
func (m *MsgGovSetEmergencyGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGovSetEmergencyGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGovSetEmergencyGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGovSetEmergencyGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGovSetEmergencyGroup.Merge(m, src)
}
func (m *MsgGovSetEmergencyGroup) XXX_Size() int {
	return m.Size()
}
func (m *MsgGovSetEmergencyGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGovSetEmergencyGroup.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGovSetEmergencyGroup proto.InternalMessageInfo

// MsgGovSetEmergencyGroupResponse response type.
type MsgGovSetEmergencyGroupResponse struct {
}

func (m *MsgGovSetEmergencyGroupResponse) Reset()         { *m = MsgGovSetEmergencyGroupResponse{} }
func (m *MsgGovSetEmergencyGroupResponse) String() string { return proto.CompactTextString(m) }
func (*MsgGovSetEmergencyGroupResponse) ProtoMessage()    {}
func (*MsgGovSetEmergencyGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ffc07de1c6ee91b, []int{3}
}
func (m *MsgGovSetEmergencyGroupResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGovSetEmergencyGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGovSetEmergencyGroupResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGovSetEmergencyGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGovSetEmergencyGroupResponse.Merge(m, src)
}
func (m *MsgGovSetEmergencyGroupResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgGovSetEmergencyGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGovSetEmergencyGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGovSetEmergencyGroupResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgGovUpdateMinGasPrice)(nil), "umee.ugov.v1.MsgGovUpdateMinGasPrice")
	proto.RegisterType((*MsgGovUpdateMinGasPriceResponse)(nil), "umee.ugov.v1.MsgGovUpdateMinGasPriceResponse")
	proto.RegisterType((*MsgGovSetEmergencyGroup)(nil), "umee.ugov.v1.MsgGovSetEmergencyGroup")
	proto.RegisterType((*MsgGovSetEmergencyGroupResponse)(nil), "umee.ugov.v1.MsgGovSetEmergencyGroupResponse")
}

func init() { proto.RegisterFile("umee/ugov/v1/tx.proto", fileDescriptor_9ffc07de1c6ee91b) }

var fileDescriptor_9ffc07de1c6ee91b = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x8b, 0x13, 0x31,
	0x18, 0xc6, 0x27, 0x2a, 0xc2, 0x66, 0x75, 0x85, 0xa1, 0xb2, 0xb5, 0x48, 0xba, 0x16, 0x94, 0x45,
	0x68, 0x42, 0x57, 0xf4, 0xb0, 0xb7, 0xad, 0xab, 0x3d, 0x15, 0xa4, 0x8b, 0x17, 0x2f, 0x65, 0x66,
	0xfa, 0x92, 0x0d, 0x9a, 0x64, 0x48, 0x32, 0xe3, 0xee, 0xd5, 0x4f, 0xe0, 0xd1, 0xa3, 0x07, 0xaf,
	0x82, 0x07, 0x3f, 0x44, 0x8f, 0x8b, 0x27, 0x4f, 0x62, 0xdb, 0x83, 0x5f, 0x43, 0x32, 0x33, 0xb5,
	0x4a, 0x3b, 0x22, 0xde, 0xe6, 0xfd, 0x93, 0xe7, 0x79, 0x7f, 0x99, 0x37, 0xf8, 0x66, 0x26, 0x01,
	0x58, 0xc6, 0x75, 0xce, 0xf2, 0x1e, 0x73, 0x67, 0x34, 0x35, 0xda, 0xe9, 0xf0, 0x9a, 0x4f, 0x53,
	0x9f, 0xa6, 0x79, 0xaf, 0x45, 0x12, 0x6d, 0xa5, 0xb6, 0x2c, 0x8e, 0x2c, 0xb0, 0xbc, 0x17, 0x83,
	0x8b, 0x7a, 0x2c, 0xd1, 0x42, 0x95, 0xdd, 0xad, 0xdd, 0xaa, 0x2e, 0x2d, 0xf7, 0x2a, 0xd2, 0xf2,
	0xaa, 0x70, 0xab, 0x2c, 0x8c, 0x8b, 0x88, 0x95, 0x41, 0x55, 0x6a, 0x70, 0xcd, 0x75, 0x99, 0xf7,
	0x5f, 0x65, 0xb6, 0xf3, 0x11, 0xe1, 0xdd, 0xa1, 0xe5, 0x03, 0x9d, 0x3f, 0x4f, 0x27, 0x91, 0x83,
	0xa1, 0x50, 0x83, 0xc8, 0x3e, 0x33, 0x22, 0x81, 0xf0, 0x11, 0xde, 0x8a, 0x32, 0x77, 0xaa, 0x8d,
	0x70, 0xe7, 0x4d, 0xb4, 0x87, 0xf6, 0xb7, 0xfa, 0xcd, 0x2f, 0x9f, 0xbb, 0x8d, 0x4a, 0xf6, 0x68,
	0x32, 0x31, 0x60, 0xed, 0x89, 0x33, 0x42, 0xf1, 0xd1, 0xaa, 0x35, 0x7c, 0x8a, 0xaf, 0x4b, 0xa1,
	0xc6, 0x3c, 0xf2, 0x73, 0x88, 0x04, 0x9a, 0x97, 0xf6, 0xd0, 0xfe, 0xf6, 0xc1, 0x6d, 0x5a, 0x1d,
	0xf4, 0x54, 0xb4, 0xa2, 0xa2, 0xc7, 0x90, 0x3c, 0xd6, 0x42, 0xf5, 0xaf, 0x4c, 0xbf, 0xb5, 0x83,
	0xd1, 0xb6, 0x5c, 0xf9, 0x1f, 0x86, 0xef, 0xde, 0xb7, 0x83, 0x37, 0x3f, 0x3e, 0xdd, 0x5f, 0x69,
	0x77, 0xee, 0xe0, 0x76, 0xcd, 0xb8, 0x23, 0xb0, 0xa9, 0x56, 0x16, 0x3a, 0x1f, 0x7e, 0x21, 0x9d,
	0x80, 0x7b, 0x22, 0xc1, 0x70, 0x50, 0xc9, 0xf9, 0xc0, 0xe8, 0x2c, 0xfd, 0x6f, 0xa4, 0x23, 0x7c,
	0x03, 0x96, 0x4a, 0x63, 0xee, 0xa5, 0x0a, 0xa8, 0xbf, 0x9d, 0xde, 0x81, 0x3f, 0xac, 0x0f, 0x77,
	0xea, 0x48, 0xd6, 0xa6, 0x5c, 0x92, 0x1c, 0xcc, 0x10, 0xbe, 0x3c, 0xb4, 0x3c, 0x7c, 0x85, 0x1b,
	0x1b, 0x7f, 0xd0, 0x5d, 0xfa, 0xfb, 0xd6, 0xd0, 0x9a, 0x8b, 0x69, 0x75, 0xff, 0xa9, 0x6d, 0xe9,
	0x5a, 0xb9, 0xad, 0xdf, 0xdd, 0x46, 0xb7, 0xb5, 0xb6, 0xcd, 0x6e, 0xb5, 0x8c, 0xfd, 0xe3, 0xe9,
	0x8c, 0x04, 0xd3, 0x39, 0x41, 0x17, 0x73, 0x82, 0xbe, 0xcf, 0x09, 0x7a, 0xbb, 0x20, 0xc1, 0xc5,
	0x82, 0x04, 0x5f, 0x17, 0x24, 0x78, 0x71, 0x8f, 0x0b, 0x77, 0x9a, 0xc5, 0x34, 0xd1, 0x92, 0x79,
	0xd9, 0xae, 0x02, 0xf7, 0x5a, 0x9b, 0x97, 0x45, 0xc0, 0xf2, 0x87, 0xec, 0xac, 0x78, 0x4a, 0xf1,
	0xd5, 0x62, 0x9b, 0x1f, 0xfc, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xa9, 0xbb, 0x4f, 0x5e, 0x03,
	0x00, 0x00,
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
	// GovUpdateMinGasPrice sets protocol controlled tx min fees.
	GovUpdateMinGasPrice(ctx context.Context, in *MsgGovUpdateMinGasPrice, opts ...grpc.CallOption) (*MsgGovUpdateMinGasPriceResponse, error)
	// GovSetEmergencyGroup sets emergency group address.
	GovSetEmergencyGroup(ctx context.Context, in *MsgGovSetEmergencyGroup, opts ...grpc.CallOption) (*MsgGovSetEmergencyGroupResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) GovUpdateMinGasPrice(ctx context.Context, in *MsgGovUpdateMinGasPrice, opts ...grpc.CallOption) (*MsgGovUpdateMinGasPriceResponse, error) {
	out := new(MsgGovUpdateMinGasPriceResponse)
	err := c.cc.Invoke(ctx, "/umee.ugov.v1.Msg/GovUpdateMinGasPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GovSetEmergencyGroup(ctx context.Context, in *MsgGovSetEmergencyGroup, opts ...grpc.CallOption) (*MsgGovSetEmergencyGroupResponse, error) {
	out := new(MsgGovSetEmergencyGroupResponse)
	err := c.cc.Invoke(ctx, "/umee.ugov.v1.Msg/GovSetEmergencyGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// GovUpdateMinGasPrice sets protocol controlled tx min fees.
	GovUpdateMinGasPrice(context.Context, *MsgGovUpdateMinGasPrice) (*MsgGovUpdateMinGasPriceResponse, error)
	// GovSetEmergencyGroup sets emergency group address.
	GovSetEmergencyGroup(context.Context, *MsgGovSetEmergencyGroup) (*MsgGovSetEmergencyGroupResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) GovUpdateMinGasPrice(ctx context.Context, req *MsgGovUpdateMinGasPrice) (*MsgGovUpdateMinGasPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovUpdateMinGasPrice not implemented")
}
func (*UnimplementedMsgServer) GovSetEmergencyGroup(ctx context.Context, req *MsgGovSetEmergencyGroup) (*MsgGovSetEmergencyGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovSetEmergencyGroup not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_GovUpdateMinGasPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGovUpdateMinGasPrice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GovUpdateMinGasPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umee.ugov.v1.Msg/GovUpdateMinGasPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GovUpdateMinGasPrice(ctx, req.(*MsgGovUpdateMinGasPrice))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GovSetEmergencyGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGovSetEmergencyGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GovSetEmergencyGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umee.ugov.v1.Msg/GovSetEmergencyGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GovSetEmergencyGroup(ctx, req.(*MsgGovSetEmergencyGroup))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "umee.ugov.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GovUpdateMinGasPrice",
			Handler:    _Msg_GovUpdateMinGasPrice_Handler,
		},
		{
			MethodName: "GovSetEmergencyGroup",
			Handler:    _Msg_GovSetEmergencyGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "umee/ugov/v1/tx.proto",
}

func (m *MsgGovUpdateMinGasPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGovUpdateMinGasPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGovUpdateMinGasPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.MinGasPrice.MarshalToSizedBuffer(dAtA[:i])
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

func (m *MsgGovUpdateMinGasPriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGovUpdateMinGasPriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGovUpdateMinGasPriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgGovSetEmergencyGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGovSetEmergencyGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGovSetEmergencyGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EmergencyGroup) > 0 {
		i -= len(m.EmergencyGroup)
		copy(dAtA[i:], m.EmergencyGroup)
		i = encodeVarintTx(dAtA, i, uint64(len(m.EmergencyGroup)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgGovSetEmergencyGroupResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGovSetEmergencyGroupResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGovSetEmergencyGroupResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgGovUpdateMinGasPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.MinGasPrice.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgGovUpdateMinGasPriceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgGovSetEmergencyGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.EmergencyGroup)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgGovSetEmergencyGroupResponse) Size() (n int) {
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
func (m *MsgGovUpdateMinGasPrice) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgGovUpdateMinGasPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGovUpdateMinGasPrice: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field MinGasPrice", wireType)
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
			if err := m.MinGasPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgGovUpdateMinGasPriceResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgGovUpdateMinGasPriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGovUpdateMinGasPriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgGovSetEmergencyGroup) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgGovSetEmergencyGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGovSetEmergencyGroup: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field EmergencyGroup", wireType)
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
			m.EmergencyGroup = string(dAtA[iNdEx:postIndex])
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
func (m *MsgGovSetEmergencyGroupResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgGovSetEmergencyGroupResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGovSetEmergencyGroupResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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

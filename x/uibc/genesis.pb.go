// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/uibc/v1/genesis.proto

package uibc

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the uibc module's genesis state.
type GenesisState struct {
	Params Params                                      `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Quotas github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,2,rep,name=quotas,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"quotas"`
	// total_outflow_sum defines the total outflow sum of ibc-transfer in USD
	TotalOutflowSum github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=total_outflow_sum,json=totalOutflowSum,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"total_outflow_sum"`
	// quota_expires defines quota expires for  ibc-transfer denom in seconds
	QuotaExpires time.Time `protobuf:"bytes,4,opt,name=quota_expires,json=quotaExpires,proto3,stdtime" json:"quota_duration,omitempty" yaml:"quota_expires"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0196ecf2d08401fb, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "umee.uibc.v1.GenesisState")
}

func init() { proto.RegisterFile("umee/uibc/v1/genesis.proto", fileDescriptor_0196ecf2d08401fb) }

var fileDescriptor_0196ecf2d08401fb = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x6d, 0x52, 0x45, 0xc2, 0x0d, 0x42, 0x58, 0x19, 0x4c, 0x84, 0xec, 0x28, 0x03, 0x8a,
	0x04, 0xb9, 0x53, 0x52, 0x26, 0xc4, 0x14, 0x82, 0x18, 0x41, 0x29, 0x13, 0x4b, 0x74, 0x76, 0xaf,
	0xee, 0xa9, 0x39, 0x3f, 0xe3, 0x7b, 0x97, 0x36, 0x03, 0xdf, 0xa1, 0x13, 0x1f, 0x82, 0x99, 0x0f,
	0x91, 0xb1, 0x62, 0x42, 0x0c, 0x2e, 0x24, 0x1b, 0x23, 0x9f, 0x00, 0xf9, 0xee, 0x22, 0xb5, 0x1b,
	0x93, 0xfd, 0xee, 0xff, 0xde, 0xff, 0x7e, 0xf7, 0x7f, 0x41, 0x4f, 0x4b, 0xce, 0xa9, 0x16, 0x69,
	0x46, 0x57, 0x63, 0x9a, 0xf3, 0x82, 0x2b, 0xa1, 0x48, 0x59, 0x01, 0x42, 0xd8, 0x69, 0x34, 0xd2,
	0x68, 0x64, 0x35, 0xee, 0x75, 0x73, 0xc8, 0xc1, 0x08, 0xb4, 0xf9, 0xb3, 0x3d, 0xbd, 0xc7, 0x19,
	0x28, 0x09, 0x6a, 0x61, 0x05, 0x5b, 0x38, 0x29, 0xb6, 0x15, 0x4d, 0x99, 0xe2, 0x74, 0x35, 0x4e,
	0x39, 0xb2, 0x31, 0xcd, 0x40, 0x14, 0x4e, 0x4f, 0x72, 0x80, 0x7c, 0xc9, 0xa9, 0xa9, 0x52, 0x7d,
	0x4a, 0x51, 0x48, 0xae, 0x90, 0xc9, 0xd2, 0x35, 0x44, 0x77, 0xd8, 0x3e, 0x69, 0x40, 0x66, 0x95,
	0xc1, 0x97, 0x56, 0xd0, 0x79, 0x6b, 0x59, 0x8f, 0x91, 0x21, 0x0f, 0x27, 0x41, 0xbb, 0x64, 0x15,
	0x93, 0x2a, 0xf2, 0xfb, 0xfe, 0xf0, 0x70, 0xd2, 0x25, 0xb7, 0xd9, 0xc9, 0x7b, 0xa3, 0x4d, 0x0f,
	0x36, 0x75, 0xe2, 0xcd, 0x5d, 0x67, 0x28, 0x82, 0xb6, 0xf1, 0x54, 0xd1, 0xbd, 0x7e, 0x6b, 0x78,
	0x38, 0x79, 0x42, 0x1c, 0x7e, 0x03, 0x4c, 0x1c, 0x30, 0x99, 0xf1, 0xec, 0x35, 0x88, 0x62, 0x7a,
	0xd4, 0xcc, 0x7e, 0xbd, 0x49, 0x9e, 0xe5, 0x02, 0xcf, 0x74, 0x4a, 0x32, 0x90, 0xee, 0xb9, 0xee,
	0x33, 0x52, 0x27, 0xe7, 0x14, 0xd7, 0x25, 0x57, 0xfb, 0x19, 0x35, 0x77, 0x17, 0x84, 0x67, 0xc1,
	0x23, 0x04, 0x64, 0xcb, 0x05, 0x68, 0x3c, 0x5d, 0xc2, 0xc5, 0x42, 0x69, 0x19, 0xb5, 0xfa, 0xfe,
	0xf0, 0xfe, 0xf4, 0x55, 0xe3, 0xfb, 0xb3, 0x4e, 0x9e, 0xfe, 0x9f, 0xef, 0xf7, 0x6f, 0xa3, 0xc0,
	0x61, 0xce, 0x78, 0x36, 0x7f, 0x68, 0x6c, 0xdf, 0x59, 0xd7, 0x63, 0x2d, 0xc3, 0xcf, 0xc1, 0x03,
	0x73, 0xe7, 0x82, 0x5f, 0x96, 0xa2, 0xe2, 0x2a, 0x3a, 0x30, 0x79, 0xf4, 0x88, 0x0d, 0x9b, 0xec,
	0xc3, 0x26, 0x1f, 0xf6, 0x61, 0x5b, 0x82, 0x3f, 0x75, 0x12, 0xd9, 0xc1, 0x13, 0x5d, 0x31, 0x14,
	0x50, 0x3c, 0x07, 0x29, 0x90, 0xcb, 0x12, 0xd7, 0x7f, 0xeb, 0xa4, 0xbb, 0x66, 0x72, 0xf9, 0x72,
	0x70, 0xc7, 0x7a, 0x70, 0x75, 0x93, 0xf8, 0xf3, 0x8e, 0x39, 0x7b, 0x63, 0x8f, 0xa6, 0xb3, 0xcd,
	0xef, 0xd8, 0xdb, 0x6c, 0x63, 0xff, 0x7a, 0x1b, 0xfb, 0xbf, 0xb6, 0xb1, 0x7f, 0xb5, 0x8b, 0xbd,
	0xeb, 0x5d, 0xec, 0xfd, 0xd8, 0xc5, 0xde, 0xc7, 0xdb, 0x6f, 0x6c, 0xf6, 0x33, 0x2a, 0x38, 0x5e,
	0x40, 0x75, 0x6e, 0x0a, 0xba, 0x7a, 0x41, 0x2f, 0xcd, 0xb6, 0xd3, 0xb6, 0xa1, 0x3c, 0xfa, 0x17,
	0x00, 0x00, 0xff, 0xff, 0x49, 0x76, 0x65, 0xe9, 0x9d, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.QuotaExpires, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.QuotaExpires):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGenesis(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	{
		size := m.TotalOutflowSum.Size()
		i -= size
		if _, err := m.TotalOutflowSum.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Quotas) > 0 {
		for iNdEx := len(m.Quotas) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Quotas[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Quotas) > 0 {
		for _, e := range m.Quotas {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.TotalOutflowSum.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.QuotaExpires)
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quotas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Quotas = append(m.Quotas, types.DecCoin{})
			if err := m.Quotas[len(m.Quotas)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalOutflowSum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalOutflowSum.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuotaExpires", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.QuotaExpires, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)

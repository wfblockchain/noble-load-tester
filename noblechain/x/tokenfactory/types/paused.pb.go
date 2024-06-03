// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tokenfactory/paused.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type Paused struct {
	Paused bool `protobuf:"varint,1,opt,name=paused,proto3" json:"paused,omitempty"`
}

func (m *Paused) Reset()         { *m = Paused{} }
func (m *Paused) String() string { return proto.CompactTextString(m) }
func (*Paused) ProtoMessage()    {}
func (*Paused) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80e08031f66ef0e, []int{0}
}
func (m *Paused) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Paused) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Paused.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Paused) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Paused.Merge(m, src)
}
func (m *Paused) XXX_Size() int {
	return m.Size()
}
func (m *Paused) XXX_DiscardUnknown() {
	xxx_messageInfo_Paused.DiscardUnknown(m)
}

var xxx_messageInfo_Paused proto.InternalMessageInfo

func (m *Paused) GetPaused() bool {
	if m != nil {
		return m.Paused
	}
	return false
}

func init() {
	proto.RegisterType((*Paused)(nil), "tokenfactory.Paused")
}

func init() { proto.RegisterFile("tokenfactory/paused.proto", fileDescriptor_f80e08031f66ef0e) }

var fileDescriptor_f80e08031f66ef0e = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0xc9, 0xcf, 0x4e,
	0xcd, 0x4b, 0x4b, 0x4c, 0x2e, 0xc9, 0x2f, 0xaa, 0xd4, 0x2f, 0x48, 0x2c, 0x2d, 0x4e, 0x4d, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x41, 0x96, 0x52, 0x52, 0xe0, 0x62, 0x0b, 0x00, 0xcb,
	0x0a, 0x89, 0x71, 0xb1, 0x41, 0xd4, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x04, 0x41, 0x79, 0x4e,
	0x21, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7,
	0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x95, 0x9e, 0x59, 0x92,
	0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x5f, 0x9e, 0x96, 0x94, 0x93, 0x9f, 0x9c, 0x9d, 0x9c,
	0x91, 0x98, 0x99, 0xa7, 0x9f, 0x97, 0x9f, 0x94, 0x93, 0x0a, 0x61, 0x96, 0x99, 0xea, 0x57, 0xe8,
	0xa3, 0x38, 0xa6, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x18, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x10, 0xd6, 0x92, 0xb4, 0xa9, 0x00, 0x00, 0x00,
}

func (m *Paused) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Paused) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Paused) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Paused {
		i--
		if m.Paused {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPaused(dAtA []byte, offset int, v uint64) int {
	offset -= sovPaused(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Paused) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Paused {
		n += 2
	}
	return n
}

func sovPaused(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPaused(x uint64) (n int) {
	return sovPaused(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Paused) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPaused
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
			return fmt.Errorf("proto: Paused: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Paused: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paused", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaused
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Paused = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPaused(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPaused
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
func skipPaused(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPaused
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
					return 0, ErrIntOverflowPaused
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
					return 0, ErrIntOverflowPaused
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
				return 0, ErrInvalidLengthPaused
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPaused
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPaused
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPaused        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPaused          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPaused = fmt.Errorf("proto: unexpected end of group")
)

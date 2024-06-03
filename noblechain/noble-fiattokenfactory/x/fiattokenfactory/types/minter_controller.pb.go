// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fiattokenfactory/minter_controller.proto

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

type MinterController struct {
	Minter     string `protobuf:"bytes,1,opt,name=minter,proto3" json:"minter,omitempty"`
	Controller string `protobuf:"bytes,2,opt,name=controller,proto3" json:"controller,omitempty"`
}

func (m *MinterController) Reset()         { *m = MinterController{} }
func (m *MinterController) String() string { return proto.CompactTextString(m) }
func (*MinterController) ProtoMessage()    {}
func (*MinterController) Descriptor() ([]byte, []int) {
	return fileDescriptor_1affbda666df4b99, []int{0}
}
func (m *MinterController) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MinterController) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MinterController.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MinterController) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MinterController.Merge(m, src)
}
func (m *MinterController) XXX_Size() int {
	return m.Size()
}
func (m *MinterController) XXX_DiscardUnknown() {
	xxx_messageInfo_MinterController.DiscardUnknown(m)
}

var xxx_messageInfo_MinterController proto.InternalMessageInfo

func (m *MinterController) GetMinter() string {
	if m != nil {
		return m.Minter
	}
	return ""
}

func (m *MinterController) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

func init() {
	proto.RegisterType((*MinterController)(nil), "fiattokenfactory.MinterController")
}

func init() {
	proto.RegisterFile("fiattokenfactory/minter_controller.proto", fileDescriptor_1affbda666df4b99)
}

var fileDescriptor_1affbda666df4b99 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0xcb, 0x4c, 0x2c,
	0x29, 0xc9, 0xcf, 0x4e, 0xcd, 0x4b, 0x4b, 0x4c, 0x2e, 0xc9, 0x2f, 0xaa, 0xd4, 0xcf, 0xcd, 0xcc,
	0x2b, 0x49, 0x2d, 0x8a, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0xca, 0xcf, 0xc9, 0x49, 0x2d, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x40, 0x57, 0xa9, 0xe4, 0xc5, 0x25, 0xe0, 0x0b, 0x56, 0xec,
	0x0c, 0x57, 0x2b, 0x24, 0xc6, 0xc5, 0x06, 0x31, 0x40, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08,
	0xca, 0x13, 0x92, 0xe3, 0xe2, 0x42, 0x98, 0x28, 0xc1, 0x04, 0x96, 0x43, 0x12, 0x71, 0x4a, 0x3c,
	0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63,
	0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xf7, 0xf4, 0xcc, 0x92, 0x8c, 0xd2,
	0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xf2, 0xb4, 0xa4, 0x9c, 0xfc, 0xe4, 0xec, 0xe4, 0x8c, 0xc4,
	0xcc, 0x3c, 0xfd, 0xbc, 0xfc, 0xa4, 0x9c, 0x54, 0x5d, 0x0c, 0xf7, 0x57, 0xe8, 0x63, 0x08, 0x95,
	0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0xfd, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x45,
	0x0e, 0xdc, 0x63, 0xf3, 0x00, 0x00, 0x00,
}

func (m *MinterController) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MinterController) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MinterController) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintMinterController(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Minter) > 0 {
		i -= len(m.Minter)
		copy(dAtA[i:], m.Minter)
		i = encodeVarintMinterController(dAtA, i, uint64(len(m.Minter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMinterController(dAtA []byte, offset int, v uint64) int {
	offset -= sovMinterController(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MinterController) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Minter)
	if l > 0 {
		n += 1 + l + sovMinterController(uint64(l))
	}
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovMinterController(uint64(l))
	}
	return n
}

func sovMinterController(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMinterController(x uint64) (n int) {
	return sovMinterController(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MinterController) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMinterController
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
			return fmt.Errorf("proto: MinterController: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MinterController: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinterController
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
				return ErrInvalidLengthMinterController
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinterController
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Minter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinterController
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
				return ErrInvalidLengthMinterController
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinterController
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMinterController(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMinterController
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
func skipMinterController(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMinterController
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
					return 0, ErrIntOverflowMinterController
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
					return 0, ErrIntOverflowMinterController
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
				return 0, ErrInvalidLengthMinterController
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMinterController
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMinterController
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMinterController        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMinterController          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMinterController = fmt.Errorf("proto: unexpected end of group")
)

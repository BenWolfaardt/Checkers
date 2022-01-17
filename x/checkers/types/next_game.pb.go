// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: checkers/next_game.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type NextGame struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	IdValue uint64 `protobuf:"varint,2,opt,name=idValue,proto3" json:"idValue,omitempty"`
}

func (m *NextGame) Reset()         { *m = NextGame{} }
func (m *NextGame) String() string { return proto.CompactTextString(m) }
func (*NextGame) ProtoMessage()    {}
func (*NextGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_c12fafd6c218bc9c, []int{0}
}
func (m *NextGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NextGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NextGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NextGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NextGame.Merge(m, src)
}
func (m *NextGame) XXX_Size() int {
	return m.Size()
}
func (m *NextGame) XXX_DiscardUnknown() {
	xxx_messageInfo_NextGame.DiscardUnknown(m)
}

var xxx_messageInfo_NextGame proto.InternalMessageInfo

func (m *NextGame) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *NextGame) GetIdValue() uint64 {
	if m != nil {
		return m.IdValue
	}
	return 0
}

func init() {
	proto.RegisterType((*NextGame)(nil), "BenWolfaardt.checkers.checkers.NextGame")
}

func init() { proto.RegisterFile("checkers/next_game.proto", fileDescriptor_c12fafd6c218bc9c) }

var fileDescriptor_c12fafd6c218bc9c = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xce, 0x48, 0x4d,
	0xce, 0x4e, 0x2d, 0x2a, 0xd6, 0xcf, 0x4b, 0xad, 0x28, 0x89, 0x4f, 0x4f, 0xcc, 0x4d, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x73, 0x4a, 0xcd, 0x0b, 0xcf, 0xcf, 0x49, 0x4b, 0x4c, 0x2c,
	0x4a, 0x29, 0xd1, 0x83, 0x29, 0x83, 0x33, 0xa4, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x4a, 0xf5,
	0x41, 0x2c, 0x88, 0x2e, 0x25, 0x3b, 0x2e, 0x0e, 0xbf, 0xd4, 0x8a, 0x12, 0xf7, 0xc4, 0xdc, 0x54,
	0x21, 0x09, 0x2e, 0xf6, 0xe4, 0xa2, 0xd4, 0xc4, 0x92, 0xfc, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x18, 0x17, 0x24, 0x93, 0x99, 0x12, 0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0xa4, 0xc0,
	0xa8, 0xc1, 0x12, 0x04, 0xe3, 0x3a, 0x79, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3,
	0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c,
	0x43, 0x94, 0x61, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xb2, 0xd3,
	0xf4, 0x9d, 0x61, 0x3e, 0xa8, 0xd0, 0x87, 0x7b, 0xa6, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d,
	0xec, 0x26, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x46, 0xda, 0x16, 0x5a, 0xe5, 0x00, 0x00,
	0x00,
}

func (m *NextGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NextGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NextGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IdValue != 0 {
		i = encodeVarintNextGame(dAtA, i, uint64(m.IdValue))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintNextGame(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNextGame(dAtA []byte, offset int, v uint64) int {
	offset -= sovNextGame(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NextGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovNextGame(uint64(l))
	}
	if m.IdValue != 0 {
		n += 1 + sovNextGame(uint64(m.IdValue))
	}
	return n
}

func sovNextGame(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNextGame(x uint64) (n int) {
	return sovNextGame(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NextGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNextGame
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
			return fmt.Errorf("proto: NextGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NextGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNextGame
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
				return ErrInvalidLengthNextGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNextGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdValue", wireType)
			}
			m.IdValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNextGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IdValue |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNextGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNextGame
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
func skipNextGame(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNextGame
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
					return 0, ErrIntOverflowNextGame
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
					return 0, ErrIntOverflowNextGame
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
				return 0, ErrInvalidLengthNextGame
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNextGame
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNextGame
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNextGame        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNextGame          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNextGame = fmt.Errorf("proto: unexpected end of group")
)

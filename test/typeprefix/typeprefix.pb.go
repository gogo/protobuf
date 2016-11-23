// Code generated by protoc-gen-gogo.
// source: typeprefix.proto
// DO NOT EDIT!

/*
Package typeprefix is a generated protocol buffer package.

It is generated from these files:
	typeprefix.proto

It has these top-level messages:
	Foo
*/
package typeprefix

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Foo struct {
	Bar string `protobuf:"bytes,1,opt,name=bar,proto3" json:"bar,omitempty"`
}

func (m *Foo) Reset()                    { *m = Foo{} }
func (m *Foo) String() string            { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()               {}
func (*Foo) Descriptor() ([]byte, []int) { return fileDescriptorTypeprefix, []int{0} }

func init() {
	proto.RegisterType((*Foo)(nil), "FooBar.Foo")
}
func (m *Foo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Foo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Bar) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTypeprefix(dAtA, i, uint64(len(m.Bar)))
		i += copy(dAtA[i:], m.Bar)
	}
	return i, nil
}

func encodeFixed64Typeprefix(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Typeprefix(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTypeprefix(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Foo) Size() (n int) {
	var l int
	_ = l
	l = len(m.Bar)
	if l > 0 {
		n += 1 + l + sovTypeprefix(uint64(l))
	}
	return n
}

func sovTypeprefix(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTypeprefix(x uint64) (n int) {
	return sovTypeprefix(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Foo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeprefix
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Foo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Foo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeprefix
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypeprefix
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bar = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypeprefix(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypeprefix
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
func skipTypeprefix(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypeprefix
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
					return 0, ErrIntOverflowTypeprefix
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
					return 0, ErrIntOverflowTypeprefix
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTypeprefix
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTypeprefix
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
				next, err := skipTypeprefix(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthTypeprefix = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypeprefix   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("typeprefix.proto", fileDescriptorTypeprefix) }

var fileDescriptorTypeprefix = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xa9, 0x2c, 0x48,
	0x2d, 0x28, 0x4a, 0x4d, 0xcb, 0xac, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xd2, 0x4d, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x0b,
	0x27, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0x51, 0xae, 0x24, 0xce, 0xc5, 0xec, 0x96,
	0x9f, 0x2f, 0x24, 0xc0, 0xc5, 0x9c, 0x94, 0x58, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04,
	0x62, 0x3a, 0x89, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c,
	0x9b, 0x1e, 0xcb, 0xb1, 0xb9, 0xe5, 0xe7, 0x3b, 0x25, 0x16, 0x25, 0xb1, 0x81, 0xf5, 0x19, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x21, 0x97, 0x83, 0x73, 0x7a, 0x00, 0x00, 0x00,
}

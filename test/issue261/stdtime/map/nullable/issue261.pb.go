// Code generated by protoc-gen-gogo.
// source: issue261.proto
// DO NOT EDIT!

/*
	Package issue261 is a generated protocol buffer package.

	It is generated from these files:
		issue261.proto

	It has these top-level messages:
		MapStdTypes
*/
package issue261

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

import time "time"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

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

type MapStdTypes struct {
	Timestamp map[int32]*time.Time `protobuf:"bytes,2,rep,name=timestamp,stdtime" json:"timestamp,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *MapStdTypes) Reset()                    { *m = MapStdTypes{} }
func (*MapStdTypes) ProtoMessage()               {}
func (*MapStdTypes) Descriptor() ([]byte, []int) { return fileDescriptorIssue261, []int{0} }

func (m *MapStdTypes) GetTimestamp() map[int32]*time.Time {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*MapStdTypes)(nil), "issue261.MapStdTypes")
}
func (this *MapStdTypes) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*MapStdTypes)
	if !ok {
		that2, ok := that.(MapStdTypes)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Timestamp) != len(that1.Timestamp) {
		return false
	}
	for i := range this.Timestamp {
		if !this.Timestamp[i].Equal(*that1.Timestamp[i]) {
			return false
		}
	}
	return true
}
func (this *MapStdTypes) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&issue261.MapStdTypes{")
	keysForTimestamp := make([]int32, 0, len(this.Timestamp))
	for k := range this.Timestamp {
		keysForTimestamp = append(keysForTimestamp, k)
	}
	github_com_gogo_protobuf_sortkeys.Int32s(keysForTimestamp)
	mapStringForTimestamp := "map[int32]*time.Time{"
	for _, k := range keysForTimestamp {
		mapStringForTimestamp += fmt.Sprintf("%#v: %#v,", k, this.Timestamp[k])
	}
	mapStringForTimestamp += "}"
	if this.Timestamp != nil {
		s = append(s, "Timestamp: "+mapStringForTimestamp+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringIssue261(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *MapStdTypes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MapStdTypes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Timestamp) > 0 {
		for k := range m.Timestamp {
			dAtA[i] = 0x12
			i++
			v := m.Timestamp[k]
			msgSize := 0
			if v != nil {
				msgSize = github_com_gogo_protobuf_types.SizeOfStdTime(*v)
				msgSize += 1 + sovIssue261(uint64(msgSize))
			}
			mapSize := 1 + sovIssue261(uint64(k)) + msgSize
			i = encodeVarintIssue261(dAtA, i, uint64(mapSize))
			dAtA[i] = 0x8
			i++
			i = encodeVarintIssue261(dAtA, i, uint64(k))
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintIssue261(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*v)))
				n1, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*v, dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n1
			}
		}
	}
	return i, nil
}

func encodeFixed64Issue261(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Issue261(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintIssue261(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MapStdTypes) Size() (n int) {
	var l int
	_ = l
	if len(m.Timestamp) > 0 {
		for k, v := range m.Timestamp {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = github_com_gogo_protobuf_types.SizeOfStdTime(*v)
				l += 1 + sovIssue261(uint64(l))
			}
			mapEntrySize := 1 + sovIssue261(uint64(k)) + l
			n += mapEntrySize + 1 + sovIssue261(uint64(mapEntrySize))
		}
	}
	return n
}

func sovIssue261(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozIssue261(x uint64) (n int) {
	return sovIssue261(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *MapStdTypes) String() string {
	if this == nil {
		return "nil"
	}
	keysForTimestamp := make([]int32, 0, len(this.Timestamp))
	for k := range this.Timestamp {
		keysForTimestamp = append(keysForTimestamp, k)
	}
	github_com_gogo_protobuf_sortkeys.Int32s(keysForTimestamp)
	mapStringForTimestamp := "map[int32]*time.Time{"
	for _, k := range keysForTimestamp {
		mapStringForTimestamp += fmt.Sprintf("%v: %v,", k, this.Timestamp[k])
	}
	mapStringForTimestamp += "}"
	s := strings.Join([]string{`&MapStdTypes{`,
		`Timestamp:` + mapStringForTimestamp + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringIssue261(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *MapStdTypes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIssue261
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
			return fmt.Errorf("proto: MapStdTypes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MapStdTypes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue261
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIssue261
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue261
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var mapkey int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue261
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				mapkey |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if m.Timestamp == nil {
				m.Timestamp = make(map[int32]*time.Time)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowIssue261
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapmsglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowIssue261
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					mapmsglen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if mapmsglen < 0 {
					return ErrInvalidLengthIssue261
				}
				postmsgIndex := iNdEx + mapmsglen
				if mapmsglen < 0 {
					return ErrInvalidLengthIssue261
				}
				if postmsgIndex > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := new(time.Time)
				if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(mapvalue, dAtA[iNdEx:postmsgIndex]); err != nil {
					return err
				}
				iNdEx = postmsgIndex
				m.Timestamp[mapkey] = mapvalue
			} else {
				var mapvalue = new(time.Time)
				m.Timestamp[mapkey] = mapvalue
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIssue261(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIssue261
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
func skipIssue261(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIssue261
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
					return 0, ErrIntOverflowIssue261
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
					return 0, ErrIntOverflowIssue261
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
				return 0, ErrInvalidLengthIssue261
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowIssue261
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
				next, err := skipIssue261(dAtA[start:])
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
	ErrInvalidLengthIssue261 = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIssue261   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("issue261.proto", fileDescriptorIssue261) }

var fileDescriptorIssue261 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0x32, 0x33, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x74,
	0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5,
	0xc1, 0x0a, 0x92, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x68, 0x94, 0x92, 0x4f, 0xcf,
	0xcf, 0x4f, 0xcf, 0x49, 0x45, 0xa8, 0x2a, 0xc9, 0xcc, 0x4d, 0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0x80,
	0x28, 0x50, 0xda, 0xc8, 0xc8, 0xc5, 0xed, 0x9b, 0x58, 0x10, 0x5c, 0x92, 0x12, 0x52, 0x59, 0x90,
	0x5a, 0x2c, 0xe4, 0xc1, 0xc5, 0x09, 0x57, 0x22, 0xc1, 0xa4, 0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0xa2,
	0x07, 0x77, 0x0d, 0x92, 0x4a, 0xbd, 0x10, 0x98, 0x32, 0xd7, 0xbc, 0x92, 0xa2, 0x4a, 0x27, 0x96,
	0x09, 0xf7, 0xe5, 0x19, 0x83, 0x10, 0x9a, 0xa5, 0x22, 0xb8, 0xf8, 0x50, 0x95, 0x08, 0x09, 0x70,
	0x31, 0x67, 0xa7, 0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x81, 0x98, 0x42, 0x06, 0x5c,
	0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x52, 0x7a, 0x10,
	0xe7, 0xea, 0xc1, 0x9c, 0x8b, 0xb0, 0x24, 0x08, 0xa2, 0xd0, 0x8a, 0xc9, 0x82, 0xd1, 0x49, 0xe7,
	0xc2, 0x43, 0x39, 0x86, 0x1b, 0x0f, 0xe5, 0x18, 0x3e, 0x3c, 0x94, 0x63, 0x6c, 0x78, 0x24, 0xc7,
	0xb8, 0xe2, 0x91, 0x1c, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24,
	0xc7, 0xf8, 0xe2, 0x91, 0x1c, 0xc3, 0x87, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x24, 0xb1,
	0x81, 0x0d, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xf2, 0x76, 0xdd, 0x54, 0x01, 0x00,
	0x00,
}

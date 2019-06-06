// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: field_mask_test_object.proto

package jsonpb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	math "math"
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

type FieldMaskWKT struct {
	FieldMask            *types.FieldMask `protobuf:"bytes,1,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *FieldMaskWKT) Reset()         { *m = FieldMaskWKT{} }
func (m *FieldMaskWKT) String() string { return proto.CompactTextString(m) }
func (*FieldMaskWKT) ProtoMessage()    {}
func (*FieldMaskWKT) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e5516809704a705, []int{0}
}
func (m *FieldMaskWKT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldMaskWKT.Unmarshal(m, b)
}
func (m *FieldMaskWKT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldMaskWKT.Marshal(b, m, deterministic)
}
func (m *FieldMaskWKT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldMaskWKT.Merge(m, src)
}
func (m *FieldMaskWKT) XXX_Size() int {
	return xxx_messageInfo_FieldMaskWKT.Size(m)
}
func (m *FieldMaskWKT) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldMaskWKT.DiscardUnknown(m)
}

var xxx_messageInfo_FieldMaskWKT proto.InternalMessageInfo

func (m *FieldMaskWKT) GetFieldMask() *types.FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func init() {
	proto.RegisterType((*FieldMaskWKT)(nil), "jsonpb.FieldMaskWKT")
}

func init() { proto.RegisterFile("field_mask_test_object.proto", fileDescriptor_8e5516809704a705) }

var fileDescriptor_8e5516809704a705 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0xcb, 0x4c, 0xcd,
	0x49, 0x89, 0xcf, 0x4d, 0x2c, 0xce, 0x8e, 0x2f, 0x49, 0x2d, 0x2e, 0x89, 0xcf, 0x4f, 0xca, 0x4a,
	0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcb, 0x2a, 0xce, 0xcf, 0x2b, 0x48,
	0x92, 0x52, 0x48, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x26, 0x95, 0xa6, 0xe9, 0x23,
	0x74, 0x41, 0x54, 0x2a, 0x79, 0x72, 0xf1, 0xb8, 0x81, 0xc4, 0x7c, 0x13, 0x8b, 0xb3, 0xc3, 0xbd,
	0x43, 0x84, 0x2c, 0xb9, 0xb8, 0x10, 0x6a, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xa4, 0xf4,
	0x20, 0xc6, 0xe8, 0xc1, 0x8c, 0xd1, 0x83, 0x6b, 0x09, 0xe2, 0x4c, 0x83, 0x31, 0x93, 0xd8, 0xc0,
	0xd2, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x41, 0x44, 0x7b, 0x9b, 0x00, 0x00, 0x00,
}

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue411.proto

package issue411

import (
	fmt "fmt"
	_ "github.com/buptbill220/protobuf/gogoproto"
	proto "github.com/buptbill220/protobuf/proto"
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

type Span struct {
	TraceID              TraceID  `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3,customtype=TraceID" json:"trace_id"`
	SpanID               SpanID   `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3,customtype=SpanID" json:"span_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Span) Reset()         { *m = Span{} }
func (m *Span) String() string { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()    {}
func (*Span) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e1ed5cde895f96f, []int{0}
}
func (m *Span) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span.Unmarshal(m, b)
}
func (m *Span) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span.Marshal(b, m, deterministic)
}
func (m *Span) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span.Merge(m, src)
}
func (m *Span) XXX_Size() int {
	return xxx_messageInfo_Span.Size(m)
}
func (m *Span) XXX_DiscardUnknown() {
	xxx_messageInfo_Span.DiscardUnknown(m)
}

var xxx_messageInfo_Span proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Span)(nil), "issue411.Span")
}

func init() { proto.RegisterFile("issue411.proto", fileDescriptor_7e1ed5cde895f96f) }

var fileDescriptor_7e1ed5cde895f96f = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0x31, 0x34, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x74,
	0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5,
	0xc1, 0x0a, 0x92, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x68, 0x54, 0x2a, 0xe0, 0x62,
	0x09, 0x2e, 0x48, 0xcc, 0x13, 0x32, 0xe5, 0xe2, 0x28, 0x29, 0x4a, 0x4c, 0x4e, 0x8d, 0xcf, 0x4c,
	0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x71, 0x92, 0x3a, 0x71, 0x4f, 0x9e, 0xe1, 0xd6, 0x3d, 0x79,
	0xf6, 0x10, 0x90, 0xb8, 0xa7, 0xcb, 0x23, 0x04, 0x33, 0x88, 0x1d, 0xac, 0xd6, 0x33, 0x45, 0xc8,
	0x90, 0x8b, 0xbd, 0xb8, 0x20, 0x31, 0x0f, 0xa4, 0x8b, 0x09, 0xac, 0x4b, 0x02, 0xaa, 0x8b, 0x0d,
	0x64, 0x2a, 0x58, 0x13, 0x94, 0x15, 0xc4, 0x06, 0x52, 0xe8, 0x99, 0x92, 0xc4, 0x06, 0xb6, 0xd8,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x68, 0x60, 0x2b, 0xc3, 0x00, 0x00, 0x00,
}

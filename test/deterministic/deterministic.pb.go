// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: deterministic.proto

package deterministic

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type OrderedMap struct {
	StringMap            map[string]string `protobuf:"bytes,1,rep,name=StringMap" json:"StringMap,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *OrderedMap) Reset()         { *m = OrderedMap{} }
func (m *OrderedMap) String() string { return proto.CompactTextString(m) }
func (*OrderedMap) ProtoMessage()    {}
func (*OrderedMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_deterministic_4064ea6d3f97db40, []int{0}
}
func (m *OrderedMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderedMap.Unmarshal(m, b)
}
func (m *OrderedMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *OrderedMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderedMap.Merge(dst, src)
}
func (m *OrderedMap) XXX_Size() int {
	return m.Size()
}
func (m *OrderedMap) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderedMap.DiscardUnknown(m)
}

var xxx_messageInfo_OrderedMap proto.InternalMessageInfo

func (m *OrderedMap) GetStringMap() map[string]string {
	if m != nil {
		return m.StringMap
	}
	return nil
}

type UnorderedMap struct {
	StringMap            map[string]string `protobuf:"bytes,1,rep,name=StringMap" json:"StringMap,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UnorderedMap) Reset()         { *m = UnorderedMap{} }
func (m *UnorderedMap) String() string { return proto.CompactTextString(m) }
func (*UnorderedMap) ProtoMessage()    {}
func (*UnorderedMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_deterministic_4064ea6d3f97db40, []int{1}
}
func (m *UnorderedMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnorderedMap.Unmarshal(m, b)
}
func (m *UnorderedMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnorderedMap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UnorderedMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnorderedMap.Merge(dst, src)
}
func (m *UnorderedMap) XXX_Size() int {
	return m.Size()
}
func (m *UnorderedMap) XXX_DiscardUnknown() {
	xxx_messageInfo_UnorderedMap.DiscardUnknown(m)
}

var xxx_messageInfo_UnorderedMap proto.InternalMessageInfo

func (m *UnorderedMap) GetStringMap() map[string]string {
	if m != nil {
		return m.StringMap
	}
	return nil
}

type MapNoMarshaler struct {
	StringMap            map[string]string `protobuf:"bytes,1,rep,name=StringMap" json:"StringMap,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MapNoMarshaler) Reset()         { *m = MapNoMarshaler{} }
func (m *MapNoMarshaler) String() string { return proto.CompactTextString(m) }
func (*MapNoMarshaler) ProtoMessage()    {}
func (*MapNoMarshaler) Descriptor() ([]byte, []int) {
	return fileDescriptor_deterministic_4064ea6d3f97db40, []int{2}
}
func (m *MapNoMarshaler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapNoMarshaler.Unmarshal(m, b)
}
func (m *MapNoMarshaler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapNoMarshaler.Marshal(b, m, deterministic)
}
func (dst *MapNoMarshaler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapNoMarshaler.Merge(dst, src)
}
func (m *MapNoMarshaler) XXX_Size() int {
	return xxx_messageInfo_MapNoMarshaler.Size(m)
}
func (m *MapNoMarshaler) XXX_DiscardUnknown() {
	xxx_messageInfo_MapNoMarshaler.DiscardUnknown(m)
}

var xxx_messageInfo_MapNoMarshaler proto.InternalMessageInfo

func (m *MapNoMarshaler) GetStringMap() map[string]string {
	if m != nil {
		return m.StringMap
	}
	return nil
}

type NestedOrderedMap struct {
	StringMap            map[string]string `protobuf:"bytes,1,rep,name=StringMap" json:"StringMap,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NestedMap            *NestedMap1       `protobuf:"bytes,2,opt,name=NestedMap" json:"NestedMap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NestedOrderedMap) Reset()         { *m = NestedOrderedMap{} }
func (m *NestedOrderedMap) String() string { return proto.CompactTextString(m) }
func (*NestedOrderedMap) ProtoMessage()    {}
func (*NestedOrderedMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_deterministic_4064ea6d3f97db40, []int{3}
}
func (m *NestedOrderedMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NestedOrderedMap.Unmarshal(m, b)
}
func (m *NestedOrderedMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *NestedOrderedMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NestedOrderedMap.Merge(dst, src)
}
func (m *NestedOrderedMap) XXX_Size() int {
	return m.Size()
}
func (m *NestedOrderedMap) XXX_DiscardUnknown() {
	xxx_messageInfo_NestedOrderedMap.DiscardUnknown(m)
}

var xxx_messageInfo_NestedOrderedMap proto.InternalMessageInfo

func (m *NestedOrderedMap) GetStringMap() map[string]string {
	if m != nil {
		return m.StringMap
	}
	return nil
}

func (m *NestedOrderedMap) GetNestedMap() *NestedMap1 {
	if m != nil {
		return m.NestedMap
	}
	return nil
}

type NestedMap1 struct {
	NestedStringMap      map[string]string `protobuf:"bytes,1,rep,name=NestedStringMap" json:"NestedStringMap,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NestedMap1) Reset()         { *m = NestedMap1{} }
func (m *NestedMap1) String() string { return proto.CompactTextString(m) }
func (*NestedMap1) ProtoMessage()    {}
func (*NestedMap1) Descriptor() ([]byte, []int) {
	return fileDescriptor_deterministic_4064ea6d3f97db40, []int{4}
}
func (m *NestedMap1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NestedMap1.Unmarshal(m, b)
}
func (m *NestedMap1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *NestedMap1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NestedMap1.Merge(dst, src)
}
func (m *NestedMap1) XXX_Size() int {
	return m.Size()
}
func (m *NestedMap1) XXX_DiscardUnknown() {
	xxx_messageInfo_NestedMap1.DiscardUnknown(m)
}

var xxx_messageInfo_NestedMap1 proto.InternalMessageInfo

func (m *NestedMap1) GetNestedStringMap() map[string]string {
	if m != nil {
		return m.NestedStringMap
	}
	return nil
}

type NestedUnorderedMap struct {
	StringMap            map[string]string `protobuf:"bytes,1,rep,name=StringMap" json:"StringMap,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NestedMap            *NestedMap2       `protobuf:"bytes,2,opt,name=NestedMap" json:"NestedMap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NestedUnorderedMap) Reset()         { *m = NestedUnorderedMap{} }
func (m *NestedUnorderedMap) String() string { return proto.CompactTextString(m) }
func (*NestedUnorderedMap) ProtoMessage()    {}
func (*NestedUnorderedMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_deterministic_4064ea6d3f97db40, []int{5}
}
func (m *NestedUnorderedMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NestedUnorderedMap.Unmarshal(m, b)
}
func (m *NestedUnorderedMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NestedUnorderedMap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *NestedUnorderedMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NestedUnorderedMap.Merge(dst, src)
}
func (m *NestedUnorderedMap) XXX_Size() int {
	return m.Size()
}
func (m *NestedUnorderedMap) XXX_DiscardUnknown() {
	xxx_messageInfo_NestedUnorderedMap.DiscardUnknown(m)
}

var xxx_messageInfo_NestedUnorderedMap proto.InternalMessageInfo

func (m *NestedUnorderedMap) GetStringMap() map[string]string {
	if m != nil {
		return m.StringMap
	}
	return nil
}

func (m *NestedUnorderedMap) GetNestedMap() *NestedMap2 {
	if m != nil {
		return m.NestedMap
	}
	return nil
}

type NestedMap2 struct {
	NestedStringMap      map[string]string `protobuf:"bytes,1,rep,name=NestedStringMap" json:"NestedStringMap,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NestedMap2) Reset()         { *m = NestedMap2{} }
func (m *NestedMap2) String() string { return proto.CompactTextString(m) }
func (*NestedMap2) ProtoMessage()    {}
func (*NestedMap2) Descriptor() ([]byte, []int) {
	return fileDescriptor_deterministic_4064ea6d3f97db40, []int{6}
}
func (m *NestedMap2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NestedMap2.Unmarshal(m, b)
}
func (m *NestedMap2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NestedMap2.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *NestedMap2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NestedMap2.Merge(dst, src)
}
func (m *NestedMap2) XXX_Size() int {
	return m.Size()
}
func (m *NestedMap2) XXX_DiscardUnknown() {
	xxx_messageInfo_NestedMap2.DiscardUnknown(m)
}

var xxx_messageInfo_NestedMap2 proto.InternalMessageInfo

func (m *NestedMap2) GetNestedStringMap() map[string]string {
	if m != nil {
		return m.NestedStringMap
	}
	return nil
}

func init() {
	proto.RegisterType((*OrderedMap)(nil), "deterministic.OrderedMap")
	proto.RegisterMapType((map[string]string)(nil), "deterministic.OrderedMap.StringMapEntry")
	proto.RegisterType((*UnorderedMap)(nil), "deterministic.UnorderedMap")
	proto.RegisterMapType((map[string]string)(nil), "deterministic.UnorderedMap.StringMapEntry")
	proto.RegisterType((*MapNoMarshaler)(nil), "deterministic.MapNoMarshaler")
	proto.RegisterMapType((map[string]string)(nil), "deterministic.MapNoMarshaler.StringMapEntry")
	proto.RegisterType((*NestedOrderedMap)(nil), "deterministic.NestedOrderedMap")
	proto.RegisterMapType((map[string]string)(nil), "deterministic.NestedOrderedMap.StringMapEntry")
	proto.RegisterType((*NestedMap1)(nil), "deterministic.NestedMap1")
	proto.RegisterMapType((map[string]string)(nil), "deterministic.NestedMap1.NestedStringMapEntry")
	proto.RegisterType((*NestedUnorderedMap)(nil), "deterministic.NestedUnorderedMap")
	proto.RegisterMapType((map[string]string)(nil), "deterministic.NestedUnorderedMap.StringMapEntry")
	proto.RegisterType((*NestedMap2)(nil), "deterministic.NestedMap2")
	proto.RegisterMapType((map[string]string)(nil), "deterministic.NestedMap2.NestedStringMapEntry")
}
func (this *OrderedMap) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*OrderedMap)
	if !ok {
		that2, ok := that.(OrderedMap)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *OrderedMap")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *OrderedMap but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *OrderedMap but is not nil && this == nil")
	}
	if len(this.StringMap) != len(that1.StringMap) {
		return fmt.Errorf("StringMap this(%v) Not Equal that(%v)", len(this.StringMap), len(that1.StringMap))
	}
	for i := range this.StringMap {
		if this.StringMap[i] != that1.StringMap[i] {
			return fmt.Errorf("StringMap this[%v](%v) Not Equal that[%v](%v)", i, this.StringMap[i], i, that1.StringMap[i])
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *OrderedMap) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*OrderedMap)
	if !ok {
		that2, ok := that.(OrderedMap)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.StringMap) != len(that1.StringMap) {
		return false
	}
	for i := range this.StringMap {
		if this.StringMap[i] != that1.StringMap[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *UnorderedMap) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*UnorderedMap)
	if !ok {
		that2, ok := that.(UnorderedMap)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *UnorderedMap")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *UnorderedMap but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *UnorderedMap but is not nil && this == nil")
	}
	if len(this.StringMap) != len(that1.StringMap) {
		return fmt.Errorf("StringMap this(%v) Not Equal that(%v)", len(this.StringMap), len(that1.StringMap))
	}
	for i := range this.StringMap {
		if this.StringMap[i] != that1.StringMap[i] {
			return fmt.Errorf("StringMap this[%v](%v) Not Equal that[%v](%v)", i, this.StringMap[i], i, that1.StringMap[i])
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *UnorderedMap) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UnorderedMap)
	if !ok {
		that2, ok := that.(UnorderedMap)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.StringMap) != len(that1.StringMap) {
		return false
	}
	for i := range this.StringMap {
		if this.StringMap[i] != that1.StringMap[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MapNoMarshaler) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*MapNoMarshaler)
	if !ok {
		that2, ok := that.(MapNoMarshaler)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *MapNoMarshaler")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *MapNoMarshaler but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *MapNoMarshaler but is not nil && this == nil")
	}
	if len(this.StringMap) != len(that1.StringMap) {
		return fmt.Errorf("StringMap this(%v) Not Equal that(%v)", len(this.StringMap), len(that1.StringMap))
	}
	for i := range this.StringMap {
		if this.StringMap[i] != that1.StringMap[i] {
			return fmt.Errorf("StringMap this[%v](%v) Not Equal that[%v](%v)", i, this.StringMap[i], i, that1.StringMap[i])
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *MapNoMarshaler) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MapNoMarshaler)
	if !ok {
		that2, ok := that.(MapNoMarshaler)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.StringMap) != len(that1.StringMap) {
		return false
	}
	for i := range this.StringMap {
		if this.StringMap[i] != that1.StringMap[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *NestedOrderedMap) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*NestedOrderedMap)
	if !ok {
		that2, ok := that.(NestedOrderedMap)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *NestedOrderedMap")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *NestedOrderedMap but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *NestedOrderedMap but is not nil && this == nil")
	}
	if len(this.StringMap) != len(that1.StringMap) {
		return fmt.Errorf("StringMap this(%v) Not Equal that(%v)", len(this.StringMap), len(that1.StringMap))
	}
	for i := range this.StringMap {
		if this.StringMap[i] != that1.StringMap[i] {
			return fmt.Errorf("StringMap this[%v](%v) Not Equal that[%v](%v)", i, this.StringMap[i], i, that1.StringMap[i])
		}
	}
	if !this.NestedMap.Equal(that1.NestedMap) {
		return fmt.Errorf("NestedMap this(%v) Not Equal that(%v)", this.NestedMap, that1.NestedMap)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *NestedOrderedMap) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NestedOrderedMap)
	if !ok {
		that2, ok := that.(NestedOrderedMap)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.StringMap) != len(that1.StringMap) {
		return false
	}
	for i := range this.StringMap {
		if this.StringMap[i] != that1.StringMap[i] {
			return false
		}
	}
	if !this.NestedMap.Equal(that1.NestedMap) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *NestedMap1) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*NestedMap1)
	if !ok {
		that2, ok := that.(NestedMap1)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *NestedMap1")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *NestedMap1 but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *NestedMap1 but is not nil && this == nil")
	}
	if len(this.NestedStringMap) != len(that1.NestedStringMap) {
		return fmt.Errorf("NestedStringMap this(%v) Not Equal that(%v)", len(this.NestedStringMap), len(that1.NestedStringMap))
	}
	for i := range this.NestedStringMap {
		if this.NestedStringMap[i] != that1.NestedStringMap[i] {
			return fmt.Errorf("NestedStringMap this[%v](%v) Not Equal that[%v](%v)", i, this.NestedStringMap[i], i, that1.NestedStringMap[i])
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *NestedMap1) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NestedMap1)
	if !ok {
		that2, ok := that.(NestedMap1)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.NestedStringMap) != len(that1.NestedStringMap) {
		return false
	}
	for i := range this.NestedStringMap {
		if this.NestedStringMap[i] != that1.NestedStringMap[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *NestedUnorderedMap) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*NestedUnorderedMap)
	if !ok {
		that2, ok := that.(NestedUnorderedMap)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *NestedUnorderedMap")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *NestedUnorderedMap but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *NestedUnorderedMap but is not nil && this == nil")
	}
	if len(this.StringMap) != len(that1.StringMap) {
		return fmt.Errorf("StringMap this(%v) Not Equal that(%v)", len(this.StringMap), len(that1.StringMap))
	}
	for i := range this.StringMap {
		if this.StringMap[i] != that1.StringMap[i] {
			return fmt.Errorf("StringMap this[%v](%v) Not Equal that[%v](%v)", i, this.StringMap[i], i, that1.StringMap[i])
		}
	}
	if !this.NestedMap.Equal(that1.NestedMap) {
		return fmt.Errorf("NestedMap this(%v) Not Equal that(%v)", this.NestedMap, that1.NestedMap)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *NestedUnorderedMap) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NestedUnorderedMap)
	if !ok {
		that2, ok := that.(NestedUnorderedMap)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.StringMap) != len(that1.StringMap) {
		return false
	}
	for i := range this.StringMap {
		if this.StringMap[i] != that1.StringMap[i] {
			return false
		}
	}
	if !this.NestedMap.Equal(that1.NestedMap) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *NestedMap2) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*NestedMap2)
	if !ok {
		that2, ok := that.(NestedMap2)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *NestedMap2")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *NestedMap2 but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *NestedMap2 but is not nil && this == nil")
	}
	if len(this.NestedStringMap) != len(that1.NestedStringMap) {
		return fmt.Errorf("NestedStringMap this(%v) Not Equal that(%v)", len(this.NestedStringMap), len(that1.NestedStringMap))
	}
	for i := range this.NestedStringMap {
		if this.NestedStringMap[i] != that1.NestedStringMap[i] {
			return fmt.Errorf("NestedStringMap this[%v](%v) Not Equal that[%v](%v)", i, this.NestedStringMap[i], i, that1.NestedStringMap[i])
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *NestedMap2) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NestedMap2)
	if !ok {
		that2, ok := that.(NestedMap2)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.NestedStringMap) != len(that1.NestedStringMap) {
		return false
	}
	for i := range this.NestedStringMap {
		if this.NestedStringMap[i] != that1.NestedStringMap[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *OrderedMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderedMap) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.StringMap) > 0 {
		keysForStringMap := make([]string, 0, len(m.StringMap))
		for k := range m.StringMap {
			keysForStringMap = append(keysForStringMap, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForStringMap)
		for _, k := range keysForStringMap {
			dAtA[i] = 0xa
			i++
			v := m.StringMap[string(k)]
			mapSize := 1 + len(k) + sovDeterministic(uint64(len(k))) + 1 + len(v) + sovDeterministic(uint64(len(v)))
			i = encodeVarintDeterministic(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintDeterministic(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintDeterministic(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *UnorderedMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnorderedMap) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.StringMap) > 0 {
		for k := range m.StringMap {
			dAtA[i] = 0xa
			i++
			v := m.StringMap[k]
			mapSize := 1 + len(k) + sovDeterministic(uint64(len(k))) + 1 + len(v) + sovDeterministic(uint64(len(v)))
			i = encodeVarintDeterministic(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintDeterministic(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintDeterministic(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *NestedOrderedMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NestedOrderedMap) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.StringMap) > 0 {
		keysForStringMap := make([]string, 0, len(m.StringMap))
		for k := range m.StringMap {
			keysForStringMap = append(keysForStringMap, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForStringMap)
		for _, k := range keysForStringMap {
			dAtA[i] = 0xa
			i++
			v := m.StringMap[string(k)]
			mapSize := 1 + len(k) + sovDeterministic(uint64(len(k))) + 1 + len(v) + sovDeterministic(uint64(len(v)))
			i = encodeVarintDeterministic(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintDeterministic(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintDeterministic(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.NestedMap != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDeterministic(dAtA, i, uint64(m.NestedMap.Size()))
		n1, err := m.NestedMap.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *NestedMap1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NestedMap1) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.NestedStringMap) > 0 {
		keysForNestedStringMap := make([]string, 0, len(m.NestedStringMap))
		for k := range m.NestedStringMap {
			keysForNestedStringMap = append(keysForNestedStringMap, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForNestedStringMap)
		for _, k := range keysForNestedStringMap {
			dAtA[i] = 0xa
			i++
			v := m.NestedStringMap[string(k)]
			mapSize := 1 + len(k) + sovDeterministic(uint64(len(k))) + 1 + len(v) + sovDeterministic(uint64(len(v)))
			i = encodeVarintDeterministic(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintDeterministic(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintDeterministic(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *NestedUnorderedMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NestedUnorderedMap) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.StringMap) > 0 {
		for k := range m.StringMap {
			dAtA[i] = 0xa
			i++
			v := m.StringMap[k]
			mapSize := 1 + len(k) + sovDeterministic(uint64(len(k))) + 1 + len(v) + sovDeterministic(uint64(len(v)))
			i = encodeVarintDeterministic(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintDeterministic(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintDeterministic(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.NestedMap != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDeterministic(dAtA, i, uint64(m.NestedMap.Size()))
		n2, err := m.NestedMap.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *NestedMap2) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NestedMap2) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.NestedStringMap) > 0 {
		for k := range m.NestedStringMap {
			dAtA[i] = 0xa
			i++
			v := m.NestedStringMap[k]
			mapSize := 1 + len(k) + sovDeterministic(uint64(len(k))) + 1 + len(v) + sovDeterministic(uint64(len(v)))
			i = encodeVarintDeterministic(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintDeterministic(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintDeterministic(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintDeterministic(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *OrderedMap) Size() (n int) {
	var l int
	_ = l
	if len(m.StringMap) > 0 {
		for k, v := range m.StringMap {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovDeterministic(uint64(len(k))) + 1 + len(v) + sovDeterministic(uint64(len(v)))
			n += mapEntrySize + 1 + sovDeterministic(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UnorderedMap) Size() (n int) {
	var l int
	_ = l
	if len(m.StringMap) > 0 {
		for k, v := range m.StringMap {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovDeterministic(uint64(len(k))) + 1 + len(v) + sovDeterministic(uint64(len(v)))
			n += mapEntrySize + 1 + sovDeterministic(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MapNoMarshaler) Size() (n int) {
	var l int
	_ = l
	if len(m.StringMap) > 0 {
		for k, v := range m.StringMap {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovDeterministic(uint64(len(k))) + 1 + len(v) + sovDeterministic(uint64(len(v)))
			n += mapEntrySize + 1 + sovDeterministic(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NestedOrderedMap) Size() (n int) {
	var l int
	_ = l
	if len(m.StringMap) > 0 {
		for k, v := range m.StringMap {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovDeterministic(uint64(len(k))) + 1 + len(v) + sovDeterministic(uint64(len(v)))
			n += mapEntrySize + 1 + sovDeterministic(uint64(mapEntrySize))
		}
	}
	if m.NestedMap != nil {
		l = m.NestedMap.Size()
		n += 1 + l + sovDeterministic(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NestedMap1) Size() (n int) {
	var l int
	_ = l
	if len(m.NestedStringMap) > 0 {
		for k, v := range m.NestedStringMap {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovDeterministic(uint64(len(k))) + 1 + len(v) + sovDeterministic(uint64(len(v)))
			n += mapEntrySize + 1 + sovDeterministic(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NestedUnorderedMap) Size() (n int) {
	var l int
	_ = l
	if len(m.StringMap) > 0 {
		for k, v := range m.StringMap {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovDeterministic(uint64(len(k))) + 1 + len(v) + sovDeterministic(uint64(len(v)))
			n += mapEntrySize + 1 + sovDeterministic(uint64(mapEntrySize))
		}
	}
	if m.NestedMap != nil {
		l = m.NestedMap.Size()
		n += 1 + l + sovDeterministic(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NestedMap2) Size() (n int) {
	var l int
	_ = l
	if len(m.NestedStringMap) > 0 {
		for k, v := range m.NestedStringMap {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovDeterministic(uint64(len(k))) + 1 + len(v) + sovDeterministic(uint64(len(v)))
			n += mapEntrySize + 1 + sovDeterministic(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDeterministic(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDeterministic(x uint64) (n int) {
	return sovDeterministic(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func init() { proto.RegisterFile("deterministic.proto", fileDescriptor_deterministic_4064ea6d3f97db40) }

var fileDescriptor_deterministic_4064ea6d3f97db40 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x49, 0x2d, 0x49,
	0x2d, 0xca, 0xcd, 0xcc, 0xcb, 0x2c, 0x2e, 0xc9, 0x4c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x45, 0x11, 0x94, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5,
	0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x4a, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b,
	0xa2, 0x5b, 0x69, 0x0e, 0x23, 0x17, 0x97, 0x7f, 0x51, 0x4a, 0x6a, 0x51, 0x6a, 0x8a, 0x6f, 0x62,
	0x81, 0x90, 0x1b, 0x17, 0x67, 0x70, 0x49, 0x51, 0x66, 0x5e, 0xba, 0x6f, 0x62, 0x81, 0x04, 0xa3,
	0x02, 0xb3, 0x06, 0xb7, 0x91, 0x86, 0x1e, 0xaa, 0xad, 0x08, 0xd5, 0x7a, 0x70, 0xa5, 0xae, 0x79,
	0x25, 0x45, 0x95, 0x41, 0x08, 0xad, 0x52, 0x36, 0x5c, 0x7c, 0xa8, 0x92, 0x42, 0x02, 0x5c, 0xcc,
	0xd9, 0xa9, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x90, 0x08, 0x17, 0x6b,
	0x59, 0x62, 0x4e, 0x69, 0xaa, 0x04, 0x13, 0x58, 0x0c, 0xc2, 0xb1, 0x62, 0xb2, 0x60, 0xb4, 0xe2,
	0xe8, 0x58, 0x28, 0xcf, 0x38, 0x63, 0xa1, 0x3c, 0xa3, 0xd2, 0x02, 0x46, 0x2e, 0x9e, 0xd0, 0xbc,
	0x7c, 0x84, 0x03, 0x3d, 0x30, 0x1d, 0xa8, 0x85, 0xe6, 0x40, 0x64, 0xf5, 0x34, 0x77, 0x22, 0x03,
	0xc8, 0x89, 0x7c, 0xbe, 0x89, 0x05, 0x7e, 0xf9, 0xbe, 0x89, 0x45, 0xc5, 0x19, 0x89, 0x39, 0xa9,
	0x45, 0x42, 0x5e, 0x98, 0x8e, 0xd4, 0x41, 0x73, 0x24, 0xaa, 0x0e, 0x9a, 0x39, 0x93, 0xa5, 0x03,
	0xe4, 0xc4, 0x87, 0x8c, 0x5c, 0x02, 0x7e, 0xa9, 0xc5, 0x25, 0xa9, 0x29, 0x48, 0x51, 0xed, 0x83,
	0xe9, 0x48, 0x3d, 0x34, 0x47, 0xa2, 0xeb, 0xc1, 0xed, 0x4c, 0x21, 0x73, 0x2e, 0x4e, 0x88, 0x6a,
	0x90, 0x69, 0x20, 0x67, 0x70, 0x1b, 0x49, 0x62, 0x35, 0xcd, 0x37, 0xb1, 0xc0, 0x30, 0x08, 0xa1,
	0x96, 0x6a, 0x29, 0x65, 0x0b, 0x23, 0x17, 0x17, 0xc2, 0x06, 0xa1, 0x08, 0x2e, 0x7e, 0x08, 0x8f,
	0x38, 0x3f, 0x82, 0xf4, 0xe8, 0xa1, 0x69, 0x80, 0xf8, 0x11, 0xdd, 0x18, 0x29, 0x27, 0x2e, 0x11,
	0x6c, 0x0a, 0xc9, 0x74, 0xf6, 0x53, 0x46, 0x2e, 0x21, 0x88, 0x71, 0x28, 0xc9, 0xdc, 0x0f, 0x33,
	0x72, 0x0c, 0xb0, 0x3a, 0x9c, 0xb8, 0xc4, 0x4e, 0x52, 0xf4, 0x18, 0x51, 0x3f, 0x7a, 0x18, 0x50,
	0xa3, 0xc7, 0x88, 0x8c, 0xe8, 0x31, 0x1a, 0x80, 0xe8, 0x61, 0x70, 0x12, 0x78, 0xf0, 0x50, 0x8e,
	0x71, 0xc5, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x3f, 0x3c, 0x92, 0x63, 0x04, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x66, 0x28, 0x8e, 0x1b, 0x84, 0x05, 0x00, 0x00,
}

// Copyright (c) 2013, Vastech SA (PTY) LTD. All rights reserved.
// http://code.google.com/p/gogoprotobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

/*
The marshalto plugin generates a Marshal and MarshalTo method for each message.
The `Marshal() ([]byte, error)` method results in the fact that the message
implements the Marshaler interface.
This allows proto.Marshal to be faster by calling the generated Marshal method rather than using reflect to Marshal the struct.

If is enabled by the following extensions:

  - marshaler
  - marshaler_all

Or the following extensions:

  - unsafe_marshaler
  - unsafe_marshaler_all

That is if you want to use the unsafe package in your generated code.
The speed up using the unsafe package is not very significant.

The generation of marshalling tests are enabled using one of the following extensions:

  - testgen
  - testgen_all

And benchmarks given it is enabled using one of the following extensions:

  - benchgen
  - benchgen_all

Let us look at:

  code.google.com/p/gogoprotobuf/test/example/example.proto

Btw all the output can be seen at:

  code.google.com/p/gogoprotobuf/test/example/*

The following message:

option (gogoproto.marshaler_all) = true;

message B {
	option (gogoproto.description) = true;
	optional A A = 1 [(gogoproto.nullable) = false, (gogoproto.embed) = true];
	repeated bytes G = 2 [(gogoproto.customtype) = "code.google.com/p/gogoprotobuf/test/custom.Uint128", (gogoproto.nullable) = false];
}

given to the marshalto plugin, will generate the following code:

  func (m *B) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
  }

  func (m *B) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintExample(data, i, uint64(m.A.Size()))
	n2, err := m.A.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.G) > 0 {
		for _, msg := range m.G {
			data[i] = 0x12
			i++
			i = encodeVarintExample(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
  }

As shown above Marshal calculates the size of the not yet marshalled message
and allocates the appropriate buffer.
This is followed by calling the MarshalTo method which requires a preallocated buffer.
The MarshalTo method allows a user to rather preallocated a reusable buffer.

The Size method is generated using the size plugin and the gogoproto.sizer, gogoproto.sizer_all extensions.
The user can also using the generated Size method to check that his reusable buffer is still big enough.

The generated tests and benchmarks will keep you safe and show that this is really a significant speed improvement.

*/
package marshalto

import (
	"code.google.com/p/gogoprotobuf/gogoproto"
	"code.google.com/p/gogoprotobuf/proto"
	descriptor "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"
	"code.google.com/p/gogoprotobuf/protoc-gen-gogo/generator"
	"fmt"
	"strconv"
	"strings"
)

type NumGen interface {
	Next() string
	Current() string
}

type numGen struct {
	index int
}

func NewNumGen() NumGen {
	return &numGen{0}
}

func (this *numGen) Next() string {
	this.index++
	return this.Current()
}

func (this *numGen) Current() string {
	return strconv.Itoa(this.index)
}

type marshalto struct {
	*generator.Generator
	generator.PluginImports
	atleastOne bool
	unsafePkg  generator.Single
	localName  string
	unsafe     bool
}

func NewMarshal() *marshalto {
	return &marshalto{}
}

func NewUnsafeMarshal() *marshalto {
	return &marshalto{unsafe: true}
}

func (p *marshalto) Name() string {
	if p.unsafe {
		return "unsafemarshaler"
	}
	return "marshalto"
}

func (p *marshalto) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *marshalto) callFixed64(varName ...string) {
	p.P(`i = encodeFixed64`, p.localName, `(data, i, uint64(`, strings.Join(varName, ""), `))`)
}

func (p *marshalto) callFixed32(varName ...string) {
	p.P(`i = encodeFixed32`, p.localName, `(data, i, uint32(`, strings.Join(varName, ""), `))`)
}

func (p *marshalto) callVarint(varName ...string) {
	p.P(`i = encodeVarint`, p.localName, `(data, i, uint64(`, strings.Join(varName, ""), `))`)
}

func (p *marshalto) callInt32Varint(varName ...string) {
	p.P(`i = encodeVarint`, p.localName, `(data, i, uint64(uint32(`, strings.Join(varName, ""), `)))`)
}

func (p *marshalto) encodeVarint(varName string) {
	p.P(`for `, varName, ` >= 1<<7 {`)
	p.In()
	p.P(`data[i] = uint8(uint64(`, varName, `)&0x7f|0x80)`)
	p.P(varName, ` >>= 7`)
	p.P(`i++`)
	p.Out()
	p.P(`}`)
	p.P(`data[i] = uint8(`, varName, `)`)
	p.P(`i++`)
}

func (p *marshalto) encodeFixed64(varName string) {
	p.P(`data[i] = uint8(`, varName, `)`)
	p.P(`i++`)
	p.P(`data[i] = uint8(`, varName, ` >> 8)`)
	p.P(`i++`)
	p.P(`data[i] = uint8(`, varName, ` >> 16)`)
	p.P(`i++`)
	p.P(`data[i] = uint8(`, varName, ` >> 24)`)
	p.P(`i++`)
	p.P(`data[i] = uint8(`, varName, ` >> 32)`)
	p.P(`i++`)
	p.P(`data[i] = uint8(`, varName, ` >> 40)`)
	p.P(`i++`)
	p.P(`data[i] = uint8(`, varName, ` >> 48)`)
	p.P(`i++`)
	p.P(`data[i] = uint8(`, varName, ` >> 56)`)
	p.P(`i++`)
}

func (p *marshalto) unsafeFixed64(varName string, someType string) {
	p.P(`*(*`, someType, `)(`, p.unsafePkg.Use(), `.Pointer(&data[i])) = `, varName)
	p.P(`i+=8`)
}

func (p *marshalto) encodeFixed32(varName string) {
	p.P(`data[i] = uint8(`, varName, `)`)
	p.P(`i++`)
	p.P(`data[i] = uint8(`, varName, ` >> 8)`)
	p.P(`i++`)
	p.P(`data[i] = uint8(`, varName, ` >> 16)`)
	p.P(`i++`)
	p.P(`data[i] = uint8(`, varName, ` >> 24)`)
	p.P(`i++`)
}

func (p *marshalto) unsafeFixed32(varName string, someType string) {
	p.P(`*(*`, someType, `)(`, p.unsafePkg.Use(), `.Pointer(&data[i])) = `, varName)
	p.P(`i+=4`)
}

func (p *marshalto) encodeKey(fieldNumber int32, wireType int) {
	x := uint32(fieldNumber)<<3 | uint32(wireType)
	i := 0
	keybuf := make([]byte, 0)
	for i = 0; x > 127; i++ {
		keybuf = append(keybuf, 0x80|uint8(x&0x7F))
		x >>= 7
	}
	keybuf = append(keybuf, uint8(x))
	for _, b := range keybuf {
		p.P(`data[i] = `, fmt.Sprintf("%#v", b))
		p.P(`i++`)
	}
}

func (p *marshalto) Generate(file *generator.FileDescriptor) {
	numGen := NewNumGen()
	p.PluginImports = generator.NewPluginImports(p.Generator)
	p.atleastOne = false
	p.localName = generator.FileName(file)

	mathPkg := p.NewImport("math")
	protoPkg := p.NewImport("code.google.com/p/gogoprotobuf/proto")
	p.unsafePkg = p.NewImport("unsafe")

	for _, message := range file.Messages() {
		ccTypeName := generator.CamelCaseSlice(message.TypeName())
		if p.unsafe {
			if !gogoproto.IsUnsafeMarshaler(file.FileDescriptorProto, message.DescriptorProto) {
				continue
			}
			if gogoproto.IsMarshaler(file.FileDescriptorProto, message.DescriptorProto) {
				panic(fmt.Sprintf("unsafe_marshaler and marshalto enabled for %v", ccTypeName))
			}
		}
		if !p.unsafe {
			if !gogoproto.IsMarshaler(file.FileDescriptorProto, message.DescriptorProto) {
				continue
			}
			if gogoproto.IsUnsafeMarshaler(file.FileDescriptorProto, message.DescriptorProto) {
				panic(fmt.Sprintf("unsafe_marshaler and marshalto enabled for %v", ccTypeName))
			}
		}
		p.atleastOne = true

		p.P(`func (m *`, ccTypeName, `) Marshal() (data []byte, err error) {`)
		p.In()
		p.P(`size := m.Size()`)
		p.P(`data = make([]byte, size)`)
		p.P(`n, err := m.MarshalTo(data)`)
		p.P(`if err != nil {`)
		p.In()
		p.P(`return nil, err`)
		p.Out()
		p.P(`}`)
		p.P(`return data[:n], nil`)
		p.Out()
		p.P(`}`)
		p.P(``)
		p.P(`func (m *`, ccTypeName, `) MarshalTo(data []byte) (n int, err error) {`)
		p.In()
		p.P(`var i int`)
		p.P(`_ = i`)
		p.P(`var l int`)
		p.P(`_ = l`)
		for _, field := range message.Field {
			fieldname := p.GetFieldName(message, field)
			nullable := gogoproto.IsNullable(field)
			repeated := field.IsRepeated()
			if repeated {
				p.P(`if len(m.`, fieldname, `) > 0 {`)
				p.In()
			} else if nullable {
				p.P(`if m.`, fieldname, ` != nil {`)
				p.In()
			}
			packed := field.IsPacked()
			wireType := field.WireType()
			fieldNumber := field.GetNumber()
			if packed {
				wireType = proto.WireBytes
			}
			switch *field.Type {
			case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
				if !p.unsafe {
					if packed {
						p.encodeKey(fieldNumber, wireType)
						p.callVarint(`len(m.`, fieldname, `) * 8`)
						p.P(`for _, num := range m.`, fieldname, ` {`)
						p.In()
						p.P(`f`, numGen.Next(), ` := `, mathPkg.Use(), `.Float64bits(num)`)
						p.encodeFixed64("f" + numGen.Current())
						p.Out()
						p.P(`}`)
					} else if repeated {
						p.P(`for _, num := range m.`, fieldname, ` {`)
						p.In()
						p.encodeKey(fieldNumber, wireType)
						p.P(`f`, numGen.Next(), ` := `, mathPkg.Use(), `.Float64bits(num)`)
						p.encodeFixed64("f" + numGen.Current())
						p.Out()
						p.P(`}`)
					} else if nullable {
						p.encodeKey(fieldNumber, wireType)
						p.callFixed64(mathPkg.Use(), `.Float64bits(*m.`+fieldname, `)`)
					} else {
						p.encodeKey(fieldNumber, wireType)
						p.callFixed64(mathPkg.Use(), `.Float64bits(m.`+fieldname, `)`)
					}
				} else {
					if packed {
						p.encodeKey(fieldNumber, wireType)
						p.callVarint(`len(m.`, fieldname, `) * 8`)
						p.P(`for _, num := range m.`, fieldname, ` {`)
						p.In()
						p.unsafeFixed64("num", "float64")
						p.Out()
						p.P(`}`)
					} else if repeated {
						p.P(`for _, num := range m.`, fieldname, ` {`)
						p.In()
						p.encodeKey(fieldNumber, wireType)
						p.unsafeFixed64("num", "float64")
						p.Out()
						p.P(`}`)
					} else if nullable {
						p.encodeKey(fieldNumber, wireType)
						p.unsafeFixed64(`*m.`+fieldname, `float64`)
					} else {
						p.encodeKey(fieldNumber, wireType)
						p.unsafeFixed64(`m.`+fieldname, "float64")
					}
				}
			case descriptor.FieldDescriptorProto_TYPE_FLOAT:
				if !p.unsafe {
					if packed {
						p.encodeKey(fieldNumber, wireType)
						p.callVarint(`len(m.`, fieldname, `) * 4`)
						p.P(`for _, num := range m.`, fieldname, ` {`)
						p.In()
						p.P(`f`, numGen.Next(), ` := `, mathPkg.Use(), `.Float32bits(num)`)
						p.encodeFixed32("f" + numGen.Current())
						p.Out()
						p.P(`}`)
					} else if repeated {
						p.P(`for _, num := range m.`, fieldname, ` {`)
						p.In()
						p.encodeKey(fieldNumber, wireType)
						p.P(`f`, numGen.Next(), ` := `, mathPkg.Use(), `.Float32bits(num)`)
						p.encodeFixed32("f" + numGen.Current())
						p.Out()
						p.P(`}`)
					} else if nullable {
						p.encodeKey(fieldNumber, wireType)
						p.callFixed32(mathPkg.Use(), `.Float32bits(*m.`+fieldname, `)`)
					} else {
						p.encodeKey(fieldNumber, wireType)
						p.callFixed32(mathPkg.Use(), `.Float32bits(m.`+fieldname, `)`)
					}
				} else {
					if packed {
						p.encodeKey(fieldNumber, wireType)
						p.callVarint(`len(m.`, fieldname, `) * 4`)
						p.P(`for _, num := range m.`, fieldname, ` {`)
						p.In()
						p.unsafeFixed32("num", "float32")
						p.Out()
						p.P(`}`)
					} else if repeated {
						p.P(`for _, num := range m.`, fieldname, ` {`)
						p.In()
						p.encodeKey(fieldNumber, wireType)
						p.unsafeFixed32("num", "float32")
						p.Out()
						p.P(`}`)
					} else if nullable {
						p.encodeKey(fieldNumber, wireType)
						p.unsafeFixed32(`*m.`+fieldname, "float32")
					} else {
						p.encodeKey(fieldNumber, wireType)
						p.unsafeFixed32(`m.`+fieldname, `float32`)
					}
				}
			case descriptor.FieldDescriptorProto_TYPE_INT64,
				descriptor.FieldDescriptorProto_TYPE_UINT64,
				descriptor.FieldDescriptorProto_TYPE_INT32,
				descriptor.FieldDescriptorProto_TYPE_UINT32,
				descriptor.FieldDescriptorProto_TYPE_ENUM:
				if packed {
					jvar := "j" + numGen.Next()
					p.P(`data`, numGen.Next(), ` := make([]byte, len(m.`, fieldname, `)*10)`)
					p.P(`var `, jvar, ` int`)
					if *field.Type == descriptor.FieldDescriptorProto_TYPE_INT32 {
						p.P(`for _, num1 := range m.`, fieldname, ` {`)
						p.In()
						p.P(`num := uint32(num1)`)
					} else if *field.Type == descriptor.FieldDescriptorProto_TYPE_INT64 {
						p.P(`for _, num1 := range m.`, fieldname, ` {`)
						p.In()
						p.P(`num := uint64(num1)`)
					} else {
						p.P(`for _, num := range m.`, fieldname, ` {`)
						p.In()
					}
					p.P(`for num >= 1<<7 {`)
					p.In()
					p.P(`data`, numGen.Current(), `[`, jvar, `] = uint8(uint64(num)&0x7f|0x80)`)
					p.P(`num >>= 7`)
					p.P(jvar, `++`)
					p.Out()
					p.P(`}`)
					p.P(`data`, numGen.Current(), `[`, jvar, `] = uint8(num)`)
					p.P(jvar, `++`)
					p.Out()
					p.P(`}`)
					p.encodeKey(fieldNumber, wireType)
					p.callVarint(jvar)
					p.P(`i += copy(data[i:], data`, numGen.Current(), `[:`, jvar, `])`)
				} else if repeated {
					p.P(`for _, num := range m.`, fieldname, ` {`)
					p.In()
					p.encodeKey(fieldNumber, wireType)
					if *field.Type == descriptor.FieldDescriptorProto_TYPE_INT32 {
						p.callInt32Varint("num")
					} else {
						p.callVarint("num")
					}
					p.Out()
					p.P(`}`)
				} else if nullable {
					p.encodeKey(fieldNumber, wireType)
					if *field.Type == descriptor.FieldDescriptorProto_TYPE_INT32 {
						p.callInt32Varint(`*m.`, fieldname)
					} else {
						p.callVarint(`*m.`, fieldname)
					}
				} else {
					p.encodeKey(fieldNumber, wireType)
					if *field.Type == descriptor.FieldDescriptorProto_TYPE_INT32 {
						p.callInt32Varint(`m.`, fieldname)
					} else {
						p.callVarint(`m.`, fieldname)
					}
				}
			case descriptor.FieldDescriptorProto_TYPE_FIXED64,
				descriptor.FieldDescriptorProto_TYPE_SFIXED64:
				if !p.unsafe {
					if packed {
						p.encodeKey(fieldNumber, wireType)
						p.callVarint(`len(m.`, fieldname, `) * 8`)
						p.P(`for _, num := range m.`, fieldname, ` {`)
						p.In()
						p.encodeFixed64("num")
						p.Out()
						p.P(`}`)
					} else if repeated {
						p.P(`for _, num := range m.`, fieldname, ` {`)
						p.In()
						p.encodeKey(fieldNumber, wireType)
						p.encodeFixed64("num")
						p.Out()
						p.P(`}`)
					} else if nullable {
						p.encodeKey(fieldNumber, wireType)
						p.callFixed64("*m." + fieldname)
					} else {
						p.encodeKey(fieldNumber, wireType)
						p.callFixed64("m." + fieldname)
					}
				} else {
					typeName := "int64"
					if *field.Type == descriptor.FieldDescriptorProto_TYPE_FIXED64 {
						typeName = "uint64"
					}
					if packed {
						p.encodeKey(fieldNumber, wireType)
						p.callVarint(`len(m.`, fieldname, `) * 8`)
						p.P(`for _, num := range m.`, fieldname, ` {`)
						p.In()
						p.unsafeFixed64("num", typeName)
						p.Out()
						p.P(`}`)
					} else if repeated {
						p.P(`for _, num := range m.`, fieldname, ` {`)
						p.In()
						p.encodeKey(fieldNumber, wireType)
						p.unsafeFixed64("num", typeName)
						p.Out()
						p.P(`}`)
					} else if nullable {
						p.encodeKey(fieldNumber, wireType)
						p.unsafeFixed64("*m."+fieldname, typeName)
					} else {
						p.encodeKey(fieldNumber, wireType)
						p.unsafeFixed64("m."+fieldname, typeName)
					}
				}
			case descriptor.FieldDescriptorProto_TYPE_FIXED32,
				descriptor.FieldDescriptorProto_TYPE_SFIXED32:
				if !p.unsafe {
					if packed {
						p.encodeKey(fieldNumber, wireType)
						p.callVarint(`len(m.`, fieldname, `) * 4`)
						p.P(`for _, num := range m.`, fieldname, ` {`)
						p.In()
						p.encodeFixed32("num")
						p.Out()
						p.P(`}`)
					} else if repeated {
						p.P(`for _, num := range m.`, fieldname, ` {`)
						p.In()
						p.encodeKey(fieldNumber, wireType)
						p.encodeFixed32("num")
						p.Out()
						p.P(`}`)
					} else if nullable {
						p.encodeKey(fieldNumber, wireType)
						p.callFixed32("*m." + fieldname)
					} else {
						p.encodeKey(fieldNumber, wireType)
						p.callFixed32("m." + fieldname)
					}
				} else {
					typeName := "int32"
					if *field.Type == descriptor.FieldDescriptorProto_TYPE_FIXED32 {
						typeName = "uint32"
					}
					if packed {
						p.encodeKey(fieldNumber, wireType)
						p.callVarint(`len(m.`, fieldname, `) * 4`)
						p.P(`for _, num := range m.`, fieldname, ` {`)
						p.In()
						p.unsafeFixed32("num", typeName)
						p.Out()
						p.P(`}`)
					} else if repeated {
						p.P(`for _, num := range m.`, fieldname, ` {`)
						p.In()
						p.encodeKey(fieldNumber, wireType)
						p.unsafeFixed32("num", typeName)
						p.Out()
						p.P(`}`)
					} else if nullable {
						p.encodeKey(fieldNumber, wireType)
						p.unsafeFixed32("*m."+fieldname, typeName)
					} else {
						p.encodeKey(fieldNumber, wireType)
						p.unsafeFixed32("m."+fieldname, typeName)
					}
				}
			case descriptor.FieldDescriptorProto_TYPE_BOOL:
				if packed {
					p.encodeKey(fieldNumber, wireType)
					p.callVarint(`len(m.`, fieldname, `)`)
					p.P(`for _, b := range m.`, fieldname, ` {`)
					p.In()
					p.P(`if b {`)
					p.In()
					p.P(`data[i] = 1`)
					p.Out()
					p.P(`} else {`)
					p.In()
					p.P(`data[i] = 0`)
					p.Out()
					p.P(`}`)
					p.P(`i++`)
					p.Out()
					p.P(`}`)
				} else if repeated {
					p.P(`for _, b := range m.`, fieldname, ` {`)
					p.In()
					p.encodeKey(fieldNumber, wireType)
					p.P(`if b {`)
					p.In()
					p.P(`data[i] = 1`)
					p.Out()
					p.P(`} else {`)
					p.In()
					p.P(`data[i] = 0`)
					p.Out()
					p.P(`}`)
					p.P(`i++`)
					p.Out()
					p.P(`}`)
				} else if nullable {
					p.encodeKey(fieldNumber, wireType)
					p.P(`if *m.`, fieldname, ` {`)
					p.In()
					p.P(`data[i] = 1`)
					p.Out()
					p.P(`} else {`)
					p.In()
					p.P(`data[i] = 0`)
					p.Out()
					p.P(`}`)
					p.P(`i++`)
				} else {
					p.encodeKey(fieldNumber, wireType)
					p.P(`if m.`, fieldname, ` {`)
					p.In()
					p.P(`data[i] = 1`)
					p.Out()
					p.P(`} else {`)
					p.In()
					p.P(`data[i] = 0`)
					p.Out()
					p.P(`}`)
					p.P(`i++`)
				}
			case descriptor.FieldDescriptorProto_TYPE_STRING:
				if repeated {
					p.P(`for _, s := range m.`, fieldname, ` {`)
					p.In()
					p.encodeKey(fieldNumber, wireType)
					p.P(`l = len(s)`)
					p.encodeVarint("l")
					p.P(`i+=copy(data[i:], s)`)
					p.Out()
					p.P(`}`)
				} else if nullable {
					p.encodeKey(fieldNumber, wireType)
					p.callVarint(`len(*m.`, fieldname, `)`)
					p.P(`i+=copy(data[i:], *m.`, fieldname, `)`)
				} else {
					p.encodeKey(fieldNumber, wireType)
					p.callVarint(`len(m.`, fieldname, `)`)
					p.P(`i+=copy(data[i:], m.`, fieldname, `)`)
				}
			case descriptor.FieldDescriptorProto_TYPE_GROUP:
				panic(fmt.Errorf("marshaler does not support group %v", fieldname))
			case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
				if repeated {
					p.P(`for _, msg := range m.`, fieldname, ` {`)
					p.In()
					p.encodeKey(fieldNumber, wireType)
					p.callVarint("msg.Size()")
					p.P(`n, err := msg.MarshalTo(data[i:])`)
					p.P(`if err != nil {`)
					p.In()
					p.P(`return 0, err`)
					p.Out()
					p.P(`}`)
					p.P(`i+=n`)
					p.Out()
					p.P(`}`)
				} else {
					p.encodeKey(fieldNumber, wireType)
					p.callVarint(`m.`, fieldname, `.Size()`)
					p.P(`n`, numGen.Next(), `, err := m.`, fieldname, `.MarshalTo(data[i:])`)
					p.P(`if err != nil {`)
					p.In()
					p.P(`return 0, err`)
					p.Out()
					p.P(`}`)
					p.P(`i+=n`, numGen.Current())
				}
			case descriptor.FieldDescriptorProto_TYPE_BYTES:
				if !gogoproto.IsCustomType(field) {
					if repeated {
						p.P(`for _, b := range m.`, fieldname, ` {`)
						p.In()
						p.encodeKey(fieldNumber, wireType)
						p.callVarint("len(b)")
						p.P(`i+=copy(data[i:], b)`)
						p.Out()
						p.P(`}`)
					} else {
						p.encodeKey(fieldNumber, wireType)
						p.callVarint(`len(m.`, fieldname, `)`)
						p.P(`i+=copy(data[i:], m.`, fieldname, `)`)
					}
				} else {
					if repeated {
						p.P(`for _, msg := range m.`, fieldname, ` {`)
						p.In()
						p.encodeKey(fieldNumber, wireType)
						p.callVarint(`msg.Size()`)
						p.P(`n, err := msg.MarshalTo(data[i:])`)
						p.P(`if err != nil {`)
						p.In()
						p.P(`return 0, err`)
						p.Out()
						p.P(`}`)
						p.P(`i+=n`)
						p.Out()
						p.P(`}`)
					} else {
						p.encodeKey(fieldNumber, wireType)
						p.callVarint(`m.`, fieldname, `.Size()`)
						p.P(`n`, numGen.Next(), `, err := m.`, fieldname, `.MarshalTo(data[i:])`)
						p.P(`if err != nil {`)
						p.In()
						p.P(`return 0, err`)
						p.Out()
						p.P(`}`)
						p.P(`i+=n`, numGen.Current())
					}
				}
			case descriptor.FieldDescriptorProto_TYPE_SINT32:
				if packed {
					datavar := "data" + numGen.Next()
					jvar := "j" + numGen.Next()
					p.P(datavar, ` := make([]byte, len(m.`, fieldname, ")*5)")
					p.P(`var `, jvar, ` int`)
					p.P(`for _, num := range m.`, fieldname, ` {`)
					p.In()
					xvar := "x" + numGen.Next()
					p.P(xvar, ` := (uint32(num) << 1) ^ uint32((num >> 31))`)
					p.P(`for `, xvar, ` >= 1<<7 {`)
					p.In()
					p.P(datavar, `[`, jvar, `] = uint8(uint64(`, xvar, `)&0x7f|0x80)`)
					p.P(jvar, `++`)
					p.P(xvar, ` >>= 7`)
					p.Out()
					p.P(`}`)
					p.P(datavar, `[`, jvar, `] = uint8(`, xvar, `)`)
					p.P(jvar, `++`)
					p.Out()
					p.P(`}`)
					p.encodeKey(fieldNumber, wireType)
					p.callVarint(jvar)
					p.P(`i+=copy(data[i:], `, datavar, `[:`, jvar, `])`)
				} else if repeated {
					p.P(`for _, num := range m.`, fieldname, ` {`)
					p.In()
					p.encodeKey(fieldNumber, wireType)
					p.P(`x`, numGen.Next(), ` := (uint32(num) << 1) ^ uint32((num >> 31))`)
					p.encodeVarint("x" + numGen.Current())
					p.Out()
					p.P(`}`)
				} else if nullable {
					p.encodeKey(fieldNumber, wireType)
					p.callVarint(`(uint32(*m.`, fieldname, `) << 1) ^ uint32((*m.`, fieldname, ` >> 31))`)
				} else {
					p.encodeKey(fieldNumber, wireType)
					p.callVarint(`(uint32(m.`, fieldname, `) << 1) ^ uint32((m.`, fieldname, ` >> 31))`)
				}
			case descriptor.FieldDescriptorProto_TYPE_SINT64:
				if packed {
					jvar := "j" + numGen.Next()
					xvar := "x" + numGen.Next()
					datavar := "data" + numGen.Next()
					p.P(`var `, jvar, ` int`)
					p.P(datavar, ` := make([]byte, len(m.`, fieldname, `)*10)`)
					p.P(`for _, num := range m.`, fieldname, ` {`)
					p.In()
					p.P(xvar, ` := (uint64(num) << 1) ^ uint64((num >> 63))`)
					p.P(`for `, xvar, ` >= 1<<7 {`)
					p.In()
					p.P(datavar, `[`, jvar, `] = uint8(uint64(`, xvar, `)&0x7f|0x80)`)
					p.P(jvar, `++`)
					p.P(xvar, ` >>= 7`)
					p.Out()
					p.P(`}`)
					p.P(datavar, `[`, jvar, `] = uint8(`, xvar, `)`)
					p.P(jvar, `++`)
					p.Out()
					p.P(`}`)
					p.encodeKey(fieldNumber, wireType)
					p.callVarint(jvar)
					p.P(`i+=copy(data[i:], `, datavar, `[:`, jvar, `])`)
				} else if repeated {
					p.P(`for _, num := range m.`, fieldname, ` {`)
					p.In()
					p.encodeKey(fieldNumber, wireType)
					p.P(`x`, numGen.Next(), ` := (uint64(num) << 1) ^ uint64((num >> 63))`)
					p.encodeVarint("x" + numGen.Current())
					p.Out()
					p.P(`}`)
				} else if nullable {
					p.encodeKey(fieldNumber, wireType)
					p.callVarint(`(uint64(*m.`, fieldname, `) << 1) ^ uint64((*m.`, fieldname, ` >> 63))`)
				} else {
					p.encodeKey(fieldNumber, wireType)
					p.callVarint(`(uint64(m.`, fieldname, `) << 1) ^ uint64((m.`, fieldname, ` >> 63))`)
				}
			default:
				panic("not implemented")
			}
			if nullable || repeated {
				p.Out()
				p.P(`}`)
			}
		}
		if message.DescriptorProto.HasExtension() {
			if gogoproto.HasExtensionsMap(file.FileDescriptorProto, message.DescriptorProto) {
				p.P(`if len(m.XXX_extensions) > 0 {`)
				p.In()
				p.P(`n, err := `, protoPkg.Use(), `.EncodeExtensionMap(m.XXX_extensions, data[i:])`)
				p.P(`if err != nil {`)
				p.In()
				p.P(`return 0, err`)
				p.Out()
				p.P(`}`)
				p.P(`i+=n`)
				p.Out()
				p.P(`}`)
			} else {
				p.P(`if m.XXX_extensions != nil {`)
				p.In()
				p.P(`i+=copy(data[i:], m.XXX_extensions)`)
				p.Out()
				p.P(`}`)
			}
		}
		p.P(`if m.XXX_unrecognized != nil {`)
		p.In()
		p.P(`i+=copy(data[i:], m.XXX_unrecognized)`)
		p.Out()
		p.P(`}`)
		p.P(`return i, nil`)
		p.Out()
		p.P(`}`)
	}

	if p.atleastOne {
		p.P(`func encodeFixed64`, p.localName, `(data []byte, offset int, v uint64) int {`)
		p.In()
		p.P(`data[offset] = uint8(v)`)
		p.P(`data[offset+1] = uint8(v >> 8)`)
		p.P(`data[offset+2] = uint8(v >> 16)`)
		p.P(`data[offset+3] = uint8(v >> 24)`)
		p.P(`data[offset+4] = uint8(v >> 32)`)
		p.P(`data[offset+5] = uint8(v >> 40)`)
		p.P(`data[offset+6] = uint8(v >> 48)`)
		p.P(`data[offset+7] = uint8(v >> 56)`)
		p.P(`return offset+8`)
		p.Out()
		p.P(`}`)

		p.P(`func encodeFixed32`, p.localName, `(data []byte, offset int, v uint32) int {`)
		p.In()
		p.P(`data[offset] = uint8(v)`)
		p.P(`data[offset+1] = uint8(v >> 8)`)
		p.P(`data[offset+2] = uint8(v >> 16)`)
		p.P(`data[offset+3] = uint8(v >> 24)`)
		p.P(`return offset+4`)
		p.Out()
		p.P(`}`)

		p.P(`func encodeVarint`, p.localName, `(data []byte, offset int, v uint64) int {`)
		p.In()
		p.P(`for v >= 1<<7 {`)
		p.In()
		p.P(`data[offset] = uint8(v&0x7f|0x80)`)
		p.P(`v >>= 7`)
		p.P(`offset++`)
		p.Out()
		p.P(`}`)
		p.P(`data[offset] = uint8(v)`)
		p.P(`return offset+1`)
		p.Out()
		p.P(`}`)
	}

}

func init() {
	generator.RegisterPlugin(NewMarshal())
	generator.RegisterPlugin(NewUnsafeMarshal())
}

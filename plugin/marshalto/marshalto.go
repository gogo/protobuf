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
	localNum   string
}

func NewMarshal() *marshalto {
	return &marshalto{}
}

func (p *marshalto) Name() string {
	return "marshalto"
}

func (p *marshalto) Init(g *generator.Generator) {
	p.Generator = g
}

func wireToType(wire string) int {
	switch wire {
	case "fixed64":
		return proto.WireFixed64
	case "fixed32":
		return proto.WireFixed32
	case "varint":
		return proto.WireVarint
	case "bytes":
		return proto.WireBytes
	case "group":
		return proto.WireBytes
	case "zigzag32":
		return proto.WireVarint
	case "zigzag64":
		return proto.WireVarint
	}
	panic("unreachable")
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

	mathPkg := p.NewImport("math")

	for _, message := range file.Messages() {
		if !gogoproto.IsMarshaler(file.FileDescriptorProto, message.DescriptorProto) {
			continue
		}
		p.atleastOne = true

		ccTypeName := generator.CamelCaseSlice(message.TypeName())
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
			fieldname := generator.CamelCase(*field.Name)
			if field.IsMessage() {
				desc := p.ObjectNamed(field.GetTypeName())
				msgname := p.TypeName(desc)
				msgnames := strings.Split(msgname, ".")
				typeName := msgnames[len(msgnames)-1]
				if gogoproto.IsEmbed(field) {
					fieldname = typeName
				}
			}
			nullable := gogoproto.IsNullable(field)
			repeated := field.IsRepeated()
			if nullable || repeated {
				p.P(`if m.`, fieldname, ` != nil {`)
				p.In()
			}
			packed := field.IsPacked()
			_, wire := p.GoType(message, field)
			wireType := wireToType(wire)
			fieldNumber := field.GetNumber()
			if packed {
				wireType = proto.WireBytes
			}
			switch *field.Type {
			case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
				if packed {
					p.P(`l = len(m.`, fieldname, `) * 8`)
					p.encodeKey(fieldNumber, wireType)
					p.encodeVarint("l")
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
					p.P(`f`, numGen.Next(), ` := `, mathPkg.Use(), `.Float64bits(*m.`+fieldname, `)`)
					p.encodeFixed64("f" + numGen.Current())
				} else {
					p.encodeKey(fieldNumber, wireType)
					p.P(`f`, numGen.Next(), ` := `, mathPkg.Use(), `.Float64bits(m.`+fieldname, `)`)
					p.encodeFixed64("f" + numGen.Current())
				}
			case descriptor.FieldDescriptorProto_TYPE_FLOAT:
				if packed {
					p.P(`l = len(m.`, fieldname, `) * 4`)
					p.encodeKey(fieldNumber, wireType)
					p.encodeVarint("l")
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
					p.P(`f`, numGen.Next(), ` := `, mathPkg.Use(), `.Float32bits(*m.`+fieldname, `)`)
					p.encodeFixed32("f" + numGen.Current())
				} else {
					p.encodeKey(fieldNumber, wireType)
					p.P(`f`, numGen.Next(), ` := `, mathPkg.Use(), `.Float32bits(m.`+fieldname, `)`)
					p.encodeFixed32("f" + numGen.Current())
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
					p.P(`for _, num := range m.`, fieldname, ` {`)
					p.In()
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
					p.P(`l = `, jvar)
					p.encodeKey(fieldNumber, wireType)
					p.encodeVarint("l")
					p.P(`copy(data[i:], data`, numGen.Current(), `[:`, jvar, `])`)
					p.P(`i += `, jvar)
				} else if repeated {
					p.P(`for _, num := range m.`, fieldname, ` {`)
					p.In()
					p.encodeKey(fieldNumber, wireType)
					p.encodeVarint("num")
					p.Out()
					p.P(`}`)
				} else if nullable {
					p.encodeKey(fieldNumber, wireType)
					p.P(`x`, numGen.Next(), ` := uint64(*m.`, fieldname, `)`)
					p.encodeVarint("x" + numGen.Current())
				} else {
					p.encodeKey(fieldNumber, wireType)
					p.P(`x`, numGen.Next(), ` := uint64(m.`, fieldname, `)`)
					p.encodeVarint("x" + numGen.Current())
				}
			case descriptor.FieldDescriptorProto_TYPE_FIXED64,
				descriptor.FieldDescriptorProto_TYPE_SFIXED64:
				if packed {
					p.P(`l = len(m.`, fieldname, `) * 8`)
					p.encodeKey(fieldNumber, wireType)
					p.encodeVarint("l")
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
					p.encodeFixed64("*m." + fieldname)
				} else {
					p.encodeKey(fieldNumber, wireType)
					p.encodeFixed64("m." + fieldname)
				}
			case descriptor.FieldDescriptorProto_TYPE_FIXED32,
				descriptor.FieldDescriptorProto_TYPE_SFIXED32:
				if packed {
					p.P(`l = len(m.`, fieldname, `) * 4`)
					p.encodeKey(fieldNumber, wireType)
					p.encodeVarint("l")
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
					p.encodeFixed32("*m." + fieldname)
				} else {
					p.encodeKey(fieldNumber, wireType)
					p.encodeFixed32("m." + fieldname)
				}
			case descriptor.FieldDescriptorProto_TYPE_BOOL:
				if packed {
					p.P(`l = len(m.`, fieldname, `)`)
					p.encodeKey(fieldNumber, wireType)
					p.encodeVarint("l")
					p.P(`for _, b := range m.`, fieldname, ` {`)
					p.In()
					p.P(`if b {`)
					p.In()
					p.P(`data[i] = 1`)
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
					p.P(`}`)
					p.P(`i++`)
				} else {
					p.encodeKey(fieldNumber, wireType)
					p.P(`if m.`, fieldname, ` {`)
					p.In()
					p.P(`data[i] = 1`)
					p.Out()
					p.P(`}`)
					p.P(`i++`)
				}
			case descriptor.FieldDescriptorProto_TYPE_STRING:
				if repeated {
					p.P(`for _, s := range m.`, fieldname, ` {`)
					p.In()
					p.P(`s`, numGen.Next(), ` := []byte(s)`)
					p.P(`l = len(s`, numGen.Current(), `)`)
					p.encodeKey(fieldNumber, wireType)
					p.encodeVarint("l")
					p.P(`copy(data[i:], s`, numGen.Current(), `)`)
					p.P(`i+=len(s`, numGen.Current(), `)`)
					p.Out()
					p.P(`}`)
				} else if nullable {
					p.P(`s`, numGen.Next(), ` := []byte(*m.`, fieldname, `)`)
					p.P(`l = len(s`, numGen.Current(), `)`)
					p.encodeKey(fieldNumber, wireType)
					p.encodeVarint("l")
					p.P(`copy(data[i:], s`, numGen.Current(), `)`)
					p.P(`i+=len(s`, numGen.Current(), `)`)
				} else {
					p.P(`s`, numGen.Next(), ` := []byte(m.`, fieldname, `)`)
					p.P(`l = len(s`, numGen.Current(), `)`)
					p.encodeKey(fieldNumber, wireType)
					p.encodeVarint("l")
					p.P(`copy(data[i:], s`, numGen.Current(), `)`)
					p.P(`i+=len(s`, numGen.Current(), `)`)
				}
			case descriptor.FieldDescriptorProto_TYPE_GROUP:
				panic("not supported")
			case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
				if repeated {
					p.P(`for _, msg := range m.`, fieldname, ` {`)
					p.In()
					p.encodeKey(fieldNumber, wireType)
					p.P(`l = msg.Size()`)
					p.encodeVarint("l")
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
					p.P(`l = m.`, fieldname, `.Size()`)
					p.encodeVarint("l")
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
						p.P(`l = len(b)`)
						p.encodeKey(fieldNumber, wireType)
						p.encodeVarint("l")
						p.P(`copy(data[i:], b)`)
						p.P(`i+=len(b)`)
						p.Out()
						p.P(`}`)
					} else {
						p.P(`l = len(m.`, fieldname, `)`)
						p.encodeKey(fieldNumber, wireType)
						p.encodeVarint("l")
						p.P(`copy(data[i:], m.`, fieldname, `)`)
						p.P(`i+=len(m.`, fieldname, `)`)
					}
				} else {
					if repeated {
						p.P(`for _, msg := range m.`, fieldname, ` {`)
						p.In()
						p.encodeKey(fieldNumber, wireType)
						p.P(`l = msg.Size()`)
						p.encodeVarint("l")
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
						p.P(`l = m.`, fieldname, `.Size()`)
						p.encodeVarint("l")
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
					p.P(`l = `, jvar)
					p.encodeKey(fieldNumber, wireType)
					p.encodeVarint("l")
					p.P(`copy(data[i:], `, datavar, `[:`, jvar, `])`)
					p.P(`i+=`, jvar)
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
					p.P(`x`, numGen.Next(), ` := (uint32(*m.`, fieldname, `) << 1) ^ uint32((*m.`, fieldname, ` >> 31))`)
					p.encodeVarint("x" + numGen.Current())
				} else {
					p.encodeKey(fieldNumber, wireType)
					p.P(`x`, numGen.Next(), ` := (uint32(m.`, fieldname, `) << 1) ^ uint32((m.`, fieldname, ` >> 31))`)
					p.encodeVarint("x" + numGen.Current())
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
					p.P(`l = `, jvar)
					p.encodeKey(fieldNumber, wireType)
					p.encodeVarint("l")
					p.P(`copy(data[i:], `, datavar, `[:`, jvar, `])`)
					p.P(`i+=`, jvar)
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
					p.P(`x`, numGen.Next(), ` := (uint64(*m.`, fieldname, `) << 1) ^ uint64((*m.`, fieldname, ` >> 63))`)
					p.encodeVarint("x" + numGen.Current())
				} else {
					p.encodeKey(fieldNumber, wireType)
					p.P(`x`, numGen.Next(), ` := (uint64(m.`, fieldname, `) << 1) ^ uint64((m.`, fieldname, ` >> 63))`)
					p.encodeVarint("x" + numGen.Current())
				}
			default:
				panic("not implemented")
			}
			if nullable || repeated {
				p.Out()
				p.P(`}`)
			}
		}
		p.P(`return i, nil`)
		p.Out()
		p.P(`}`)
	}

	if !p.atleastOne {
		return
	}

}

func init() {
	generator.RegisterPlugin(NewMarshal())
}

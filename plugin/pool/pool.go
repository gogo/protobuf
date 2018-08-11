// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
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
Package pool provides a plugin that generates message pools.
*/
package pool

import (
	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

func init() {
	generator.RegisterPlugin(newPool())
}

type pool struct {
	*generator.Generator
	generator.PluginImports

	memPkg generator.Single
}

func newPool() *pool {
	return &pool{}
}

func (p *pool) Name() string {
	return "pool"
}

func (p *pool) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *pool) Generate(file *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)

	if !gogoproto.HasPool(file.FileDescriptorProto) {
		return
	}

	p.memPkg = p.NewImport("github.com/gogo/protobuf/mem")

	var messages []*generator.Descriptor
	for _, message := range file.Messages() {
		// Don't generate for map entry messages.
		if !message.GetOptions().GetMapEntry() {
			messages = append(messages, message)
		}
	}
	if len(messages) == 0 {
		return
	}

	// GetMessage
	for _, message := range messages {
		messageGoType := getMessageGoType(message)
		p.P(`// Get`, messageGoType, ` gets a reset *`, messageGoType, `.`)
		p.P(`func Get`, messageGoType, `() *`, messageGoType, `{ `)
		p.In()
		p.P(`if !`, p.memPkg.Use(), `.PoolingEnabled() {`)
		p.In()
		p.P(`return &`, messageGoType, `{}`)
		p.Out()
		p.P(`}`)
		p.P(`return global`, messageGoType, `ObjectPool.Get().(*`, messageGoType, `)`)
		p.Out()
		p.P(`}`)
	}

	// message.Recycle
	for _, message := range messages {
		messageGoType := getMessageGoType(message)
		p.P(`// Recycle puts the message back in the Pool that created it.`)
		p.P(`//`)
		p.P(`// Any non-nil fields that are messages will be recycled as well, including elements of repeated fields and values of map fields.`)
		p.P(`// Once Recycle is called, the message should no longer be used.`)
		p.P(`//`)
		p.P(`// If the message is nil, the message was not allocated from a Pool, or pooling is disabled, this is a no-op.`)
		p.P(`func (m *`, messageGoType, `) Recycle() {`)
		p.In()
		p.P(`if !`, p.memPkg.Use(), `.PoolingEnabled() {`)
		p.In()
		p.P(`return`)
		p.Out()
		p.P(`}`)
		p.P(`if m == nil {`)
		p.In()
		p.P(`return`)
		p.Out()
		p.P(`}`)
		for _, field := range message.Field {
			if *field.Type == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
				fieldName := p.GetFieldName(message, field)
				if p.IsMap(field) {
					valueField := p.GoMapType(nil, field).ValueField
					if *valueField.Type == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
						if gogoproto.HasPool(p.ObjectNamed(valueField.GetTypeName()).File().FileDescriptorProto) {
							p.P(`for _, elem := range m.`, fieldName, ` {`)
							p.In()
							p.P(`elem.Recycle()`)
							p.Out()
							p.P(`}`)
						}
					}
				} else if field.IsRepeated() {
					if gogoproto.HasPool(p.ObjectNamed(field.GetTypeName()).File().FileDescriptorProto) {
						p.P(`for _, value := range m.`, fieldName, ` {`)
						p.In()
						p.P(`value.Recycle()`)
						p.Out()
						p.P(`}`)
					}
				} else {
					if gogoproto.HasPool(p.ObjectNamed(field.GetTypeName()).File().FileDescriptorProto) {
						p.P(`m.`, fieldName, `.Recycle()`)
					}
				}
			}
		}
		p.P(`global`, messageGoType, `ObjectPool.Put(m)`)
		p.Out()
		p.P(`}`)
		p.P()
	}

	// constructors
	for _, message := range messages {
		messageGoType := getMessageGoType(message)
		p.P(`func new`, messageGoType, `() `, p.memPkg.Use(), `.Object {`)
		p.In()
		p.P(`return &`, messageGoType, `{}`)
		p.Out()
		p.P(`}`)
		p.P()
	}

	// object pools
	p.P(`var (`)
	p.In()
	for _, message := range messages {
		messageGoType := getMessageGoType(message)
		p.P(`global`, messageGoType, `ObjectPool = `, p.memPkg.Use(), `.NewObjectPool(new`, messageGoType, `)`)
	}
	p.Out()
	p.P(`)`)
	p.P()

	// init
	p.P(`func init() {`)
	p.In()
	for _, message := range messages {
		messageGoType := getMessageGoType(message)
		p.P(p.memPkg.Use(), `.RegisterGlobalObjectPool(`)
		p.In()
		p.P(`func() *`, p.memPkg.Use(), `.ObjectPool { return global`, messageGoType, `ObjectPool },`)
		p.P(`func(options ...`, p.memPkg.Use(), `.ObjectPoolOption) { global`, messageGoType, `ObjectPool = `, p.memPkg.Use(), `.NewObjectPool(new`, messageGoType, `, options...) },`)
		p.Out()
		p.P(`)`)
	}
	p.Out()
	p.P(`}`)
	p.P()
}

func getMessageGoType(message *generator.Descriptor) string {
	return generator.CamelCaseSlice(message.TypeName())
}

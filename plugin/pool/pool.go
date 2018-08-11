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
	"strings"

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

	memPkg    generator.Single
	localName string
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
	p.localName = generator.FileName(file)

	if !gogoproto.GeneratesPool(file.FileDescriptorProto) {
		return
	}

	p.memPkg = p.NewImport("github.com/gogo/protobuf/mem")

	// RegisterToPoolFile
	p.P(`// RegisterToPool`, p.localName, ` registers constructors for all messages in this file to the given Pool.`)
	p.P(`//`)
	p.P(`// Users must call this function at initialization for the Pool to pool messages in this file.`)
	p.P(`func RegisterToPool`, p.localName, ` (pool *`, p.memPkg.Use(), `.Pool) {`)
	p.In()
	for _, message := range file.Messages() {
		// Don't generate for map entry messages.
		if message.GetOptions().GetMapEntry() {
			continue
		}
		p.P(`pool.RegisterConstructor("`, getMessageType(file, message), `", registerToPool`, getMessageGoType(message), `)`)
	}
	p.Out()
	p.P(`}`)
	p.P()

	// GetMessage, GetMessageFromPool
	for _, message := range file.Messages() {
		// Don't generate for map entry messages.
		if message.GetOptions().GetMapEntry() {
			continue
		}
		messageGoType := getMessageGoType(message)

		p.P(`// Get`, messageGoType, ` gets a reset *`, messageGoType, ` from the global Pool.`)
		p.P(`//`)
		p.P(`// It is generally preferable to create a Pool instance for your application and`)
		p.P(`// use Get`, messageGoType, `FromPool instead, however this will prove difficult`)
		p.P(`// for existing applications wishing to migrate to the Pool API. Additionally,`)
		p.P(`// RegisterToPool`, p.localName, `is implicitly called on the global Pool so there`)
		p.P(`// is no need for additional setup to use this function.`)
		p.P(`func Get`, messageGoType, `() *`, messageGoType, `{ `)
		p.In()
		p.P(`return Get`, messageGoType, `FromPool(`, p.memPkg.Use(), `.Global())`)
		p.Out()
		p.P(`}`)

		p.P(`// Get`, messageGoType, `FromPool gets a reset *`, messageGoType, ` from the given Pool.`)
		p.P(`//`)
		p.P(`// If the Pool is nil, or `, p.localName, `PoolRegister was not called on this Pool, this will return a new message not managed by any Pool.`)
		p.P(`func Get`, messageGoType, `FromPool(pool *`, p.memPkg.Use(), `.Pool) *`, messageGoType, `{ `)
		p.In()
		p.P(`if pool == nil {`)
		p.In()
		p.P(`return &`, messageGoType, `{}`)
		p.Out()
		p.P(`}`)
		p.P(`pooledMessage := pool.Get("`, getMessageType(file, message), `")`)
		p.P(`if pooledMessage == nil {`)
		p.In()
		p.P(`return &`, messageGoType, `{}`)
		p.Out()
		p.P(`}`)
		p.P(`return pooledMessage.(*`, messageGoType, `)`)
		p.Out()
		p.P(`}`)
		p.P()
	}

	// message.Recycle
	for _, message := range file.Messages() {
		// Don't generate for map entry messages.
		if message.GetOptions().GetMapEntry() {
			continue
		}
		p.P(`// Recycle puts the message back in the Pool that created it.`)
		p.P(`//`)
		p.P(`// Any non-nil fields that are messages will be recycled as well, including elements of repeated fields and values of map fields.`)
		p.P(`//`)
		p.P(`// If the message is nil, or the message was not allocated from a Pool, this is a no-op.`)
		p.P(`func (this *`, getMessageGoType(message), `) Recycle() {`)
		p.In()
		p.P(`if this == nil {`)
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
						p.P(`for _, elem := range this.`, fieldName, ` {`)
						p.In()
						// TODO: If there is a way to go from FieldDescriptorProto to FileDescriptorProto,
						// we can discover whether or not gogoproto.GeneratesPool returns true for the file
						// that defines this message and we can skip this check and call Recycle directly.
						p.P(`var iface interface{} = elem`)
						p.P(`if pooledMessage, ok := iface.(`, p.memPkg.Use(), `.PooledMessage); ok {`)
						p.In()
						p.P(`pooledMessage.Recycle()`)
						p.Out()
						p.P(`}`)
						p.Out()
						p.P(`}`)
					}
				} else if field.IsRepeated() {
					p.P(`for _, elem := range this.`, fieldName, ` {`)
					p.In()
					// TODO: If there is a way to go from FieldDescriptorProto to FileDescriptorProto,
					// we can discover whether or not gogoproto.GeneratesPool returns true for the file
					// that defines this message and we can skip this check and call Recycle directly.
					p.P(`var iface interface{} = elem`)
					p.P(`if pooledMessage, ok := iface.(`, p.memPkg.Use(), `.PooledMessage); ok {`)
					p.In()
					p.P(`pooledMessage.Recycle()`)
					p.Out()
					p.P(`}`)
					p.Out()
					p.P(`}`)
				} else {
					// TODO: If there is a way to go from FieldDescriptorProto to FileDescriptorProto,
					// we can discover whether or not gogoproto.GeneratesPool returns true for the file
					// that defines this message and we can skip this check and call Recycle directly.
					p.P(`var iface interface{} = this.`, fieldName)
					p.P(`if pooledMessage, ok := iface.(`, p.memPkg.Use(), `.PooledMessage); ok {`)
					p.In()
					p.P(`pooledMessage.Recycle()`)
					p.Out()
					p.P(`}`)
				}
			}
		}
		p.P(`if this.pool == nil {`)
		p.In()
		p.P(`return`)
		p.Out()
		p.P(`}`)
		p.P(`this.pool.Put("`, getMessageType(file, message), `", this)`)
		p.Out()
		p.P(`}`)
		p.P()
	}

	// registerToPoolMessage
	for _, message := range file.Messages() {
		// Don't generate for map entry messages.
		if message.GetOptions().GetMapEntry() {
			continue
		}
		messageGoType := getMessageGoType(message)
		p.P(`func registerToPool`, messageGoType, `(pool *`, p.memPkg.Use(), `.Pool) `, p.memPkg.Use(), `.PooledMessage {`)
		p.In()
		p.P(`return &`, messageGoType, `{pool: pool}`)
		p.Out()
		p.P(`}`)
		p.P()
	}

	// init
	p.P(`func init() {`)
	p.In()
	p.P(`RegisterToPool`, p.localName, `(`, p.memPkg.Use(), `.Global())`)
	p.Out()
	p.P(`}`)
	p.P()
}

func getMessageType(file *generator.FileDescriptor, message *generator.Descriptor) string {
	fullName := strings.Join(message.TypeName(), ".")
	if file.Package != nil {
		fullName = *file.Package + "." + fullName
	}
	return fullName
}

func getMessageGoType(message *generator.Descriptor) string {
	return generator.CamelCaseSlice(message.TypeName())
}

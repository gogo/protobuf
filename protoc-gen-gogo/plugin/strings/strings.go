// Extensions for Protocol Buffers to create more go like structures.
//
// Copyright (c) 2013, Vastech SA (PTY) LTD. All rights reserved.
// http://code.google.com/p/gogoprotobuf/gogoproto
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

package strings

import (
	"code.google.com/p/gogoprotobuf/protoc-gen-gogo/generator"
)

type plugin struct {
	*generator.Generator
	imports map[string]bool
}

func NewPlugin() *plugin {
	return &plugin{
		imports: make(map[string]bool),
	}
}

func (p *plugin) Name() string {
	return "strings"
}

func (p *plugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *plugin) Generate(file *generator.FileDescriptor) {
	p.imports = make(map[string]bool)
	fileOf := p.FileOf(file.FileDescriptorProto)
	for _, msg := range fileOf.Messages() {
		p.generateMessage(msg)
	}
}

func (p *plugin) generateMessage(message *generator.Descriptor) {
	typeName := message.TypeName()
	ccTypeName := generator.CamelCaseSlice(typeName)
	if !message.IsGroup() {
		p.P("func (m *", ccTypeName, ") String() string { return ", p.Pkg["proto"], ".CompactTextString(m) }")
	}
}

func (p *plugin) GenerateImports(file *generator.FileDescriptor) {
}

func init() {
	generator.RegisterPlugin(NewPlugin())
}

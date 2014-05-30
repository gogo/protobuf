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
The defaultcheck plugin is used to check whether nullable is not used incorrectly.
For instance:
An error is caused if a nullable field:
  - has a default value.
  - is an enum which does not start at zero.
  - is used for an extension.

An error is also caused if a field with a default value is used in a message:
  - which is a face.
  - without getters.

It is enabled by the following extensions:

  - nullable

For incorrect usage of nullable with tests see:

  code.google.com/p/gogoprotobuf/test/nullableconflict

*/
package defaultcheck

import (
	"code.google.com/p/gogoprotobuf/gogoproto"
	"code.google.com/p/gogoprotobuf/protoc-gen-gogo/generator"
	"fmt"
	"os"
)

type plugin struct {
	*generator.Generator
}

func NewPlugin() *plugin {
	return &plugin{}
}

func (p *plugin) Name() string {
	return "defaultcheck"
}

func (p *plugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *plugin) Generate(file *generator.FileDescriptor) {
	for _, msg := range file.Messages() {
		getters := gogoproto.HasGoGetters(file.FileDescriptorProto, msg.DescriptorProto)
		face := gogoproto.IsFace(file.FileDescriptorProto, msg.DescriptorProto)
		for _, field := range msg.GetField() {
			if len(field.GetDefaultValue()) > 0 {
				if !getters {
					fmt.Fprintf(os.Stderr, "ERROR: field %v.%v cannot have a default value and not have a getter method", generator.CamelCase(*msg.Name), generator.CamelCase(*field.Name))
					os.Exit(1)
				}
				if face {
					fmt.Fprintf(os.Stderr, "ERROR: field %v.%v cannot have a default value be in a face", generator.CamelCase(*msg.Name), generator.CamelCase(*field.Name))
					os.Exit(1)
				}
			}
			if gogoproto.IsNullable(field) {
				continue
			}
			if len(field.GetDefaultValue()) > 0 {
				fmt.Fprintf(os.Stderr, "ERROR: field %v.%v cannot be non-nullable and have a default value", generator.CamelCase(*msg.Name), generator.CamelCase(*field.Name))
				os.Exit(1)
			}
			if !field.IsEnum() {
				continue
			}
			enum := p.ObjectNamed(field.GetTypeName()).(*generator.EnumDescriptor)
			if len(enum.Value) == 0 || enum.Value[0].GetNumber() != 0 {
				fmt.Fprintf(os.Stderr, "ERROR: field %v.%v cannot be non-nullable and be an enum type %v which does not start with zero", generator.CamelCase(*msg.Name), generator.CamelCase(*field.Name), enum.GetName())
				os.Exit(1)
			}
		}
	}
	for _, e := range file.GetExtension() {
		if !gogoproto.IsNullable(e) {
			fmt.Fprintf(os.Stderr, "ERROR: extended field %v cannot be nullable %v", generator.CamelCase(e.GetName()), generator.CamelCase(*e.Name))
			os.Exit(1)
		}
	}
}

func (p *plugin) GenerateImports(*generator.FileDescriptor) {}

func init() {
	generator.RegisterPlugin(NewPlugin())
}

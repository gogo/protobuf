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
The description plugin generates a Description method for each message.
The Description method returns a populated google_protobuf.DescriptorProto struct.
This contains the description of the specific message.

It is enabled by the following extensions:

  - description
  - description_all

The description plugin also generates a test given it is enabled using one of the following extensions:

  - testgen
  - testgen_all

Let us look at:

  code.google.com/p/gogoprotobuf/test/example/example.proto

Btw all the output can be seen at:

  code.google.com/p/gogoprotobuf/test/example/*

The following message:

  message B {
	option (gogoproto.description) = true;
	optional A A = 1 [(gogoproto.nullable) = false, (gogoproto.embed) = true];
	repeated bytes G = 2 [(gogoproto.customtype) = "code.google.com/p/gogoprotobuf/test/custom.Uint128", (gogoproto.nullable) = false];
  }

given to the description plugin, will generate the following code:

  func (this *B) Description() (desc *google_protobuf.DescriptorProto) {
	return &google_protobuf.DescriptorProto{Name: func(v string) *string { return &v }("B"), Field: []*google_protobuf.FieldDescriptorProto{{Name: func(v string) *string { return &v }("A"), Number: func(v int32) *int32 { return &v }(1), Label: func(v google_protobuf.FieldDescriptorProto_Label) *google_protobuf.FieldDescriptorProto_Label {
		return &v
	}(1), Type: func(v google_protobuf.FieldDescriptorProto_Type) *google_protobuf.FieldDescriptorProto_Type {
		return &v
	}(11), TypeName: func(v string) *string { return &v }(".test.A"), Extendee: nil, DefaultValue: nil, Options: &google_protobuf.FieldOptions{Ctype: nil, Packed: nil, Lazy: nil, Deprecated: nil, ExperimentalMapKey: nil, Weak: nil, UninterpretedOption: []*google_protobuf.UninterpretedOption(nil)}}, {Name: func(v string) *string { return &v }("G"), Number: func(v int32) *int32 { return &v }(2), Label: func(v google_protobuf.FieldDescriptorProto_Label) *google_protobuf.FieldDescriptorProto_Label {
		return &v
	}(3), Type: func(v google_protobuf.FieldDescriptorProto_Type) *google_protobuf.FieldDescriptorProto_Type {
		return &v
	}(12), TypeName: nil, Extendee: nil, DefaultValue: nil, Options: &google_protobuf.FieldOptions{Ctype: nil, Packed: nil, Lazy: nil, Deprecated: nil, ExperimentalMapKey: nil, Weak: nil, UninterpretedOption: []*google_protobuf.UninterpretedOption(nil)}}}, Extension: []*google_protobuf.FieldDescriptorProto(nil), NestedType: []*google_protobuf.DescriptorProto(nil), EnumType: []*google_protobuf.EnumDescriptorProto(nil), ExtensionRange: []*google_protobuf.DescriptorProto_ExtensionRange(nil), Options: &google_protobuf.MessageOptions{MessageSetWireFormat: nil, NoStandardDescriptorAccessor: nil, UninterpretedOption: []*google_protobuf.UninterpretedOption(nil)}}
  }

and the following test code:

  func TestBDescription(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedB(popr)
	p.Description()
  }

The hope is to use this struct in some way instead of reflect.
This package is subject to change, since a use has not been figured out yet and might require a bigger desciptor, like a FileDescriptorSet.

*/
package description

import (
	"code.google.com/p/gogoprotobuf/gogoproto"
	"code.google.com/p/gogoprotobuf/protoc-gen-gogo/generator"
	"fmt"
)

type plugin struct {
	*generator.Generator
	used bool
}

func NewPlugin() *plugin {
	return &plugin{}
}

func (p *plugin) Name() string {
	return "description"
}

func (p *plugin) Init(g *generator.Generator) {
	p.used = false
	p.Generator = g
}

func (p *plugin) Generate(file *generator.FileDescriptor) {
	for _, message := range file.Messages() {
		if !gogoproto.HasDescription(file.FileDescriptorProto, message.DescriptorProto) {
			continue
		}
		p.used = true
	}

	if p.used {
		p.P(`func Description() (desc *google_protobuf.FileDescriptorSet) {`)
		p.In()
		s := fmt.Sprintf("%#v", p.Generator.AllFiles())
		p.P(`return `, s)
		p.Out()
		p.P(`}`)
	}
}

func (this *plugin) GenerateImports(file *generator.FileDescriptor) {
	if this.used {
		this.P(`import google_protobuf "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"`)
	}
}

func init() {
	generator.RegisterPlugin(NewPlugin())
}

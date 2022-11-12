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
This plugin will mark a generated type as "well known" to the gogo proto JSON marshaler/unmarshaler which will
serialize the value without the wrapper. This happens by default for google.protobuf.*Value message types.

# Without jsonwkt
message UUID {
	string value = 1;
}

message MyType {
	UUID user_uuid = 1;
}

MyType would be serialized to JSON as:

{
	"user_uuid": {
		"value": "8dfab8c2-7916-477f-bbd6-c9fc00dc4158"
	}
}

# With jsonwkt
message UUID {
	option (gogoproto.json_well_known_type) = "StringValue";
	string value = 1;
}

message MyType {
	UUID user_uuid = 1;
}

MyType would be serialized to JSON as:

{
	"user_uuid": "8dfab8c2-7916-477f-bbd6-c9fc00dc4158"
}
*/

package jsonwkt

import (
	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

var wellKnownTypes = map[string]bool{
	"Duration":    true,
	"Timestamp":   true,
	"Value":       true,
	"ListValue":   true,
	"DoubleValue": true,
	"FloatValue":  true,
	"Int64Value":  true,
	"UInt64Value": true,
	"Int32Value":  true,
	"UInt32Value": true,
	"BoolValue":   true,
	"StringValue": true,
	"BytesValue":  true,
}

type jsonwkt struct {
	*generator.Generator
	generator.PluginImports
}

func NewJsonWKT() *jsonwkt {
	return &jsonwkt{}
}

func (p *jsonwkt) Name() string {
	return "jsonwkt"
}

func (p *jsonwkt) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *jsonwkt) Generate(file *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)
	for _, message := range file.Messages() {
		wellKnownType := gogoproto.GetWellKnownType(message.DescriptorProto)
		if wellKnownType == nil {
			continue
		}
		if _, ok := wellKnownTypes[*wellKnownType]; !ok {
			p.Generator.Fail("invalid json_well_known_type option %s", *wellKnownType)
			continue
		}

		ccTypeName := generator.CamelCaseSlice(message.TypeName())
		p.Generator.P("func (*", ccTypeName, `) XXX_WellKnownType() string { return "`, *wellKnownType, `" }`)
	}
}

func init() {
	generator.RegisterPlugin(NewJsonWKT())
}

/**
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

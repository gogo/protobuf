package jsonwkt

import (
	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/plugin/testgen"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

type test struct {
	*generator.Generator
}

func NewTest(g *generator.Generator) testgen.TestPlugin {
	return &test{g}
}

func (p *test) Generate(imports generator.PluginImports, file *generator.FileDescriptor) bool {
	used := false
	randPkg := imports.NewImport("math/rand")
	timePkg := imports.NewImport("time")
	testingPkg := imports.NewImport("testing")
	for _, message := range file.Messages() {
		wellKnownType := gogoproto.GetWellKnownType(message.DescriptorProto)
		if wellKnownType == nil {
			continue
		}

		if gogoproto.HasTestGen(file.FileDescriptorProto, message.DescriptorProto) {
			ccTypeName := generator.CamelCaseSlice(message.TypeName())
			used = true
			p.P(`func Test`, ccTypeName, `JsonWKT(t *`, testingPkg.Use(), `.T) {`)
			p.In()
			p.P(`popr := `, randPkg.Use(), `.New(`, randPkg.Use(), `.NewSource(`, timePkg.Use(), `.Now().UnixNano()))`)
			p.P(`p := NewPopulated`, ccTypeName, `(popr, false)`)

			p.P(`_, ok := interface{}(p).(interface { XXX_WellKnownType() string })`)
			p.P(`if !ok {`)
			p.In()
			p.P(`t.Fatalf("Type `, ccTypeName, ` should implement XXX_WellKnownType but did not")`)
			p.Out()
			p.P(`}`)
			p.Out()
			p.P(`}`)
		}
	}
	return used
}

func init() {
	testgen.RegisterTestPlugin(NewTest)
}

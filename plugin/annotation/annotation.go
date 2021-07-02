package annotation

import (
	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

type annotation struct {
	*generator.Generator
	generator.PluginImports
}

func NewAnnotation() *annotation {
	return &annotation{}
}

func (p *annotation) Name() string {
	return "annotation"
}

func (p *annotation) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *annotation) Generate(file *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)
	for _, message := range file.Messages() {
		if gogoproto.GetStringAnnotation(message.DescriptorProto) == "" {
			continue
		}

		ccTypeName := generator.CamelCaseSlice(message.TypeName())
		p.P(`func init(){ proto.RegisterAnnotation(`,
			`(*`, ccTypeName, ")(nil), `", gogoproto.GetStringAnnotation(message.DescriptorProto), "`)}")
	}
}

func init() {
	generator.RegisterPlugin(NewAnnotation())
}

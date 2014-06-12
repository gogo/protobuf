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
The populate plugin generates a NewPopulated function.
This function returns a newly populated structure.

It is enabled by the following extensions:

  - populate
  - populate_all

Let us look at:

  code.google.com/p/gogoprotobuf/test/example/example.proto

Btw all the output can be seen at:

  code.google.com/p/gogoprotobuf/test/example/*

The following message:

  option (gogoproto.populate_all) = true;

  message B {
	optional A A = 1 [(gogoproto.nullable) = false, (gogoproto.embed) = true];
	repeated bytes G = 2 [(gogoproto.customtype) = "code.google.com/p/gogoprotobuf/test/custom.Uint128", (gogoproto.nullable) = false];
  }

given to the populate plugin, will generate code the following code:

  func NewPopulatedB(r randyExample, easy bool) *B {
	this := &B{}
	v2 := NewPopulatedA(r, easy)
	this.A = *v2
	if r.Intn(10) != 0 {
		v3 := r.Intn(10)
		this.G = make([]code_google_com_p_gogoprotobuf_test_custom.Uint128, v3)
		for i := 0; i < v3; i++ {
			v4 := code_google_com_p_gogoprotobuf_test_custom.NewPopulatedUint128(r)
			this.G[i] = *v4
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedExample(r, 3)
	}
	return this
  }

The idea that is useful for testing.
Most of the other plugins' generated test code uses it.
You will still be able to use the generated test code of other packages
if you turn off the popluate plugin and write your own custom NewPopulated function.

If the easy flag is not set the XXX_unrecognized and XXX_extensions fields are also populated.
These have caused problems with JSON marshalling and unmarshalling tests.

*/
package populate

import (
	"code.google.com/p/gogoprotobuf/gogoproto"
	descriptor "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"
	"code.google.com/p/gogoprotobuf/protoc-gen-gogo/generator"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	maxRune      = '\U0010FFFF'
	surrogateMin = 0xD800
	surrogateMax = 0xDFFF
)

type VarGen interface {
	Next() string
	Current() string
}

type varGen struct {
	index int64
}

func NewVarGen() VarGen {
	return &varGen{0}
}

func (this *varGen) Next() string {
	this.index++
	return fmt.Sprintf("v%d", this.index)
}

func (this *varGen) Current() string {
	return fmt.Sprintf("v%d", this.index)
}

type plugin struct {
	*generator.Generator
	generator.PluginImports
	varGen     VarGen
	atleastOne bool
	localName  string
}

func NewPlugin() *plugin {
	return &plugin{}
}

func (p *plugin) Name() string {
	return "populate"
}

func (p *plugin) Init(g *generator.Generator) {
	p.Generator = g
}

func value(typName string) string {
	switch typName {
	case "float64":
		return "r.Float64()"
	case "float32":
		return "r.Float32()"
	case "int64":
		return "r.Int63()"
	case "int32":
		return "r.Int31()"
	case "uint32":
		return "r.Uint32()"
	case "uint64":
		return "uint64(r.Uint32())"
	case "bool":
		return `bool(r.Intn(2) == 0)`
	}
	panic(fmt.Errorf("unexpected type %v", typName))
}

func negative(typeName string) bool {
	switch typeName {
	case "uint32", "uint64", "bool":
		return false
	}
	return true
}

func (p *plugin) GenerateField(message *generator.Descriptor, field *descriptor.FieldDescriptorProto) {
	goTyp, _ := p.GoType(message, field)
	fieldname := p.GetFieldName(message, field)
	goTypName := generator.GoTypeToName(goTyp)
	if field.IsMessage() || p.IsGroup(field) {
		funcName := "NewPopulated" + goTypName
		goTypNames := strings.Split(goTypName, ".")
		if len(goTypNames) == 2 {
			funcName = goTypNames[0] + ".NewPopulated" + goTypNames[1]
		} else if len(goTypNames) != 1 {
			panic(fmt.Errorf("unreachable: too many dots in %v", goTypName))
		}
		funcCall := funcName + "(r, easy)"
		if field.IsRepeated() {
			p.P(p.varGen.Next(), ` := r.Intn(10)`)
			p.P(`this.`, fieldname, ` = make(`, goTyp, `, `, p.varGen.Current(), `)`)
			p.P(`for i := 0; i < `, p.varGen.Current(), `; i++ {`)
			p.In()
			if gogoproto.IsNullable(field) {
				p.P(`this.`, fieldname, `[i] = `, funcCall)
			} else {
				p.P(p.varGen.Next(), `:= `, funcCall)
				p.P(`this.`, fieldname, `[i] = *`, p.varGen.Current())
			}
			p.Out()
			p.P(`}`)
		} else {
			if gogoproto.IsNullable(field) {
				p.P(`this.`, fieldname, ` = `, funcCall)
			} else {
				p.P(p.varGen.Next(), `:= `, funcCall)
				p.P(`this.`, fieldname, ` = *`, p.varGen.Current())
			}
		}
	} else {
		if field.IsEnum() {
			enum := p.ObjectNamed(field.GetTypeName()).(*generator.EnumDescriptor)
			l := len(enum.Value)
			values := make([]string, l)
			for i := range enum.Value {
				values[i] = strconv.Itoa(int(*enum.Value[i].Number))
			}
			arr := "[]int32{" + strings.Join(values, ",") + "}"
			val := strings.Join([]string{generator.GoTypeToName(goTyp), `(`, arr, `[r.Intn(`, fmt.Sprintf("%d", l), `)])`}, "")
			if field.IsRepeated() {
				p.P(p.varGen.Next(), ` := r.Intn(10)`)
				p.P(`this.`, fieldname, ` = make(`, goTyp, `, `, p.varGen.Current(), `)`)
				p.P(`for i := 0; i < `, p.varGen.Current(), `; i++ {`)
				p.In()
				p.P(`this.`, fieldname, `[i] = `, val)
				p.Out()
				p.P(`}`)
			} else if gogoproto.IsNullable(field) {
				p.P(p.varGen.Next(), ` := `, val)
				p.P(`this.`, fieldname, ` = &`, p.varGen.Current())
			} else {
				p.P(`this.`, fieldname, ` = `, val)
			}
		} else if gogoproto.IsCustomType(field) {
			funcName := "NewPopulated" + goTypName
			goTypNames := strings.Split(goTypName, ".")
			if len(goTypNames) == 2 {
				funcName = goTypNames[0] + ".NewPopulated" + goTypNames[1]
			} else if len(goTypNames) != 1 {
				panic(fmt.Errorf("unreachable: too many dots in %v", goTypName))
			}
			funcCall := funcName + "(r)"
			if field.IsRepeated() {
				p.P(p.varGen.Next(), ` := r.Intn(10)`)
				p.P(`this.`, fieldname, ` = make(`, goTyp, `, `, p.varGen.Current(), `)`)
				p.P(`for i := 0; i < `, p.varGen.Current(), `; i++ {`)
				p.In()
				p.P(p.varGen.Next(), `:= `, funcCall)
				p.P(`this.`, fieldname, `[i] = *`, p.varGen.Current())
				p.Out()
				p.P(`}`)
			} else if gogoproto.IsNullable(field) {
				p.P(`this.`, fieldname, ` = `, funcCall)
			} else {
				p.P(p.varGen.Next(), `:= `, funcCall)
				p.P(`this.`, fieldname, ` = *`, p.varGen.Current())
			}
		} else if field.IsBytes() {
			if field.IsRepeated() {
				p.P(p.varGen.Next(), ` := r.Intn(100)`)
				p.P(`this.`, fieldname, ` = make(`, goTyp, `, `, p.varGen.Current(), `)`)
				p.P(`for i := 0; i < `, p.varGen.Current(), `; i++ {`)
				p.In()
				p.P(p.varGen.Next(), ` := r.Intn(100)`)
				p.P(`this.`, fieldname, `[i] = make([]byte,`, p.varGen.Current(), `)`)
				p.P(`for j := 0; j < `, p.varGen.Current(), `; j++ {`)
				p.In()
				p.P(`this.`, fieldname, `[i][j] = byte(r.Intn(256))`)
				p.Out()
				p.P(`}`)
				p.Out()
				p.P(`}`)
			} else {
				p.P(p.varGen.Next(), ` := r.Intn(100)`)
				p.P(`this.`, fieldname, ` = make(`, goTyp, `, `, p.varGen.Current(), `)`)
				p.P(`for i := 0; i < `, p.varGen.Current(), `; i++ {`)
				p.In()
				p.P(`this.`, fieldname, `[i] = byte(r.Intn(256))`)
				p.Out()
				p.P(`}`)
			}
		} else if field.IsString() {
			val := fmt.Sprintf("randString%v(r)", p.localName)
			if field.IsRepeated() {
				p.P(p.varGen.Next(), ` := r.Intn(10)`)
				p.P(`this.`, fieldname, ` = make(`, goTyp, `, `, p.varGen.Current(), `)`)
				p.P(`for i := 0; i < `, p.varGen.Current(), `; i++ {`)
				p.In()
				p.P(`this.`, fieldname, `[i] = `, val)
				p.Out()
				p.P(`}`)
			} else if gogoproto.IsNullable(field) {
				p.P(p.varGen.Next(), `:= `, val)
				p.P(`this.`, fieldname, ` = &`, p.varGen.Current())
			} else {
				p.P(`this.`, fieldname, ` = `, val)
			}
		} else {
			typName := generator.GoTypeToName(goTyp)
			if field.IsRepeated() {
				p.P(p.varGen.Next(), ` := r.Intn(100)`)
				p.P(`this.`, fieldname, ` = make(`, goTyp, `, `, p.varGen.Current(), `)`)
				p.P(`for i := 0; i < `, p.varGen.Current(), `; i++ {`)
				p.In()
				p.P(`this.`, fieldname, `[i] = `, value(typName))
				if negative(typName) {
					p.P(`if r.Intn(2) == 0 {`)
					p.In()
					p.P(`this.`, fieldname, `[i] *= -1`)
					p.Out()
					p.P(`}`)
				}
				p.Out()
				p.P(`}`)
			} else if gogoproto.IsNullable(field) {
				p.P(p.varGen.Next(), ` := `, value(typName))
				if negative(typName) {
					p.P(`if r.Intn(2) == 0 {`)
					p.In()
					p.P(p.varGen.Current(), ` *= -1`)
					p.Out()
					p.P(`}`)
				}
				p.P(`this.`, fieldname, ` = &`, p.varGen.Current())
			} else {
				p.P(`this.`, fieldname, ` = `, value(typName))
				if negative(typName) {
					p.P(`if r.Intn(2) == 0 {`)
					p.In()
					p.P(`this.`, fieldname, ` *= -1`)
					p.Out()
					p.P(`}`)
				}
			}
		}
	}
}

func (p *plugin) hasLoop(field *descriptor.FieldDescriptorProto, visited []*generator.Descriptor, excludes []*generator.Descriptor) *generator.Descriptor {
	if field.IsMessage() || p.IsGroup(field) {
		fieldMessage := p.ObjectNamed(field.GetTypeName()).(*generator.Descriptor)
		fieldTypeName := generator.CamelCaseSlice(fieldMessage.TypeName())
		for _, message := range visited {
			messageTypeName := generator.CamelCaseSlice(message.TypeName())
			if fieldTypeName == messageTypeName {
				for _, e := range excludes {
					if fieldTypeName == generator.CamelCaseSlice(e.TypeName()) {
						return nil
					}
				}
				return fieldMessage
			}
		}
		for _, f := range fieldMessage.Field {
			visited = append(visited, fieldMessage)
			loopTo := p.hasLoop(f, visited, excludes)
			if loopTo != nil {
				return loopTo
			}
		}
	}
	return nil
}

func (p *plugin) loops(field *descriptor.FieldDescriptorProto, message *generator.Descriptor) int {
	//fmt.Fprintf(os.Stderr, "loops %v %v\n", field.GetTypeName(), generator.CamelCaseSlice(message.TypeName()))
	excludes := []*generator.Descriptor{}
	loops := 0
	for {
		visited := []*generator.Descriptor{}
		loopTo := p.hasLoop(field, visited, excludes)
		if loopTo == nil {
			break
		}
		//fmt.Fprintf(os.Stderr, "loopTo %v\n", generator.CamelCaseSlice(loopTo.TypeName()))
		excludes = append(excludes, loopTo)
		loops++
	}
	return loops
}

func (p *plugin) Generate(file *generator.FileDescriptor) {
	p.atleastOne = false
	p.PluginImports = generator.NewPluginImports(p.Generator)
	p.varGen = NewVarGen()

	p.localName = generator.FileName(file)
	protoPkg := p.NewImport("code.google.com/p/gogoprotobuf/proto")

	for _, message := range file.Messages() {
		if !gogoproto.HasPopulate(file.FileDescriptorProto, message.DescriptorProto) {
			continue
		}
		p.atleastOne = true
		ccTypeName := generator.CamelCaseSlice(message.TypeName())
		p.P(`func NewPopulated`, ccTypeName, `(r randy`, p.localName, `, easy bool) *`, ccTypeName, ` {`)
		p.In()
		p.P(`this := &`, ccTypeName, `{}`)
		if gogoproto.IsUnion(message.File(), message.DescriptorProto) && len(message.Field) > 0 {
			loopLevels := make([]int, len(message.Field))
			maxLoopLevel := 0
			for i, field := range message.Field {
				loopLevels[i] = p.loops(field, message)
				if loopLevels[i] > maxLoopLevel {
					maxLoopLevel = loopLevels[i]
				}
			}
			ran := 0
			for i := range loopLevels {
				ran += int(math.Pow10(maxLoopLevel - loopLevels[i]))
			}
			p.P(`fieldNum := r.Intn(`, fmt.Sprintf("%d", ran), `)`)
			p.P(`switch fieldNum {`)
			k := 0
			for i, field := range message.Field {
				is := []string{}
				ran := int(math.Pow10(maxLoopLevel - loopLevels[i]))
				for j := 0; j < ran; j++ {
					is = append(is, fmt.Sprintf("%d", j+k))
				}
				k += ran
				p.P(`case `, strings.Join(is, ","), `:`)
				p.In()
				p.GenerateField(message, field)
				p.Out()
			}
			p.P(`}`)
		} else {
			var maxFieldNumber int32
			for _, field := range message.Field {
				if field.IsRequired() || (!gogoproto.IsNullable(field) && !field.IsRepeated()) {
					p.GenerateField(message, field)
				} else {
					p.P(`if r.Intn(10) != 0 {`)
					p.In()
					p.GenerateField(message, field)
					p.Out()
					p.P(`}`)
				}
				if field.GetNumber() > maxFieldNumber {
					maxFieldNumber = field.GetNumber()
				}
			}
			if message.DescriptorProto.HasExtension() {
				p.P(`if !easy && r.Intn(10) != 0 {`)
				p.In()
				p.P(`l := r.Intn(5)`)
				p.P(`for i := 0; i < l; i++ {`)
				p.In()
				if len(message.DescriptorProto.GetExtensionRange()) > 1 {
					p.P(`eIndex := r.Intn(`, strconv.Itoa(len(message.DescriptorProto.GetExtensionRange())), `)`)
					p.P(`fieldNumber := 0`)
					p.P(`switch eIndex {`)
					for i, e := range message.DescriptorProto.GetExtensionRange() {
						p.P(`case `, strconv.Itoa(i), `:`)
						p.In()
						p.P(`fieldNumber = r.Intn(`, strconv.Itoa(int(e.GetEnd()-e.GetStart())), `) + `, strconv.Itoa(int(e.GetStart())))
						p.Out()
						if e.GetEnd() > maxFieldNumber {
							maxFieldNumber = e.GetEnd()
						}
					}
					p.P(`}`)
				} else {
					e := message.DescriptorProto.GetExtensionRange()[0]
					p.P(`fieldNumber := r.Intn(`, strconv.Itoa(int(e.GetEnd()-e.GetStart())), `) + `, strconv.Itoa(int(e.GetStart())))
					if e.GetEnd() > maxFieldNumber {
						maxFieldNumber = e.GetEnd()
					}
				}
				p.P(`wire := r.Intn(4)`)
				p.P(`if wire == 3 { wire = 5 }`)
				p.P(`data := randField`, p.localName, `(nil, r, fieldNumber, wire)`)
				p.P(protoPkg.Use(), `.SetRawExtension(this, int32(fieldNumber), data)`)
				p.Out()
				p.P(`}`)
				p.Out()
				p.P(`}`)
			}

			if maxFieldNumber < (1 << 10) {
				p.P(`if !easy && r.Intn(10) != 0 {`)
				p.In()
				p.P(`this.XXX_unrecognized = randUnrecognized`, p.localName, `(r, `, strconv.Itoa(int(maxFieldNumber+1)), `)`)
				p.Out()
				p.P(`}`)
			}
		}
		p.P(`return this`)
		p.Out()
		p.P(`}`)
		p.P(``)
	}

	if !p.atleastOne {
		return
	}

	p.P(`type randy`, p.localName, ` interface {`)
	p.In()
	p.P(`Float32() float32`)
	p.P(`Float64() float64`)
	p.P(`Int63() int64`)
	p.P(`Int31() int32`)
	p.P(`Uint32() uint32`)
	p.P(`Intn(n int) int`)
	p.Out()
	p.P(`}`)

	surrogateRange := surrogateMax - surrogateMin
	maxRand := maxRune - surrogateRange

	p.P(`func randUTF8Rune`, p.localName, `(r randy`, p.localName, `) rune {`)
	p.In()
	p.P(`res := rune(r.Uint32() % `, fmt.Sprintf("%d", maxRand), `)`)
	p.P(`if `, fmt.Sprintf("%d", surrogateMin), ` <= res {`)
	p.In()
	p.P(`res += `, fmt.Sprintf("%d", surrogateRange))
	p.Out()
	p.P(`}`)
	p.P(`return res`)
	p.Out()
	p.P(`}`)

	p.P(`func randString`, p.localName, `(r randy`, p.localName, `) string {`)
	p.In()
	p.P(p.varGen.Next(), ` := r.Intn(100)`)
	p.P(`tmps := make([]rune, `, p.varGen.Current(), `)`)
	p.P(`for i := 0; i < `, p.varGen.Current(), `; i++ {`)
	p.In()
	p.P(`tmps[i] = randUTF8Rune`, p.localName, `(r)`)
	p.Out()
	p.P(`}`)
	p.P(`return string(tmps)`)
	p.Out()
	p.P(`}`)

	p.P(`func randUnrecognized`, p.localName, `(r randy`, p.localName, `, maxFieldNumber int) (data []byte) {`)
	p.In()
	p.P(`l := r.Intn(5)`)
	p.P(`for i := 0; i < l; i++ {`)
	p.In()
	p.P(`wire := r.Intn(4)`)
	p.P(`if wire == 3 { wire = 5 }`)
	p.P(`fieldNumber := maxFieldNumber + r.Intn(100)`)
	p.P(`data = randField`, p.localName, `(data, r, fieldNumber, wire)`)
	p.Out()
	p.P(`}`)
	p.P(`return data`)
	p.Out()
	p.P(`}`)

	p.P(`func randField`, p.localName, `(data []byte, r randy`, p.localName, `, fieldNumber int, wire int) []byte {`)
	p.In()
	p.P(`key := uint32(fieldNumber)<<3 | uint32(wire)`)
	p.P(`switch wire {`)
	p.P(`case 0:`)
	p.In()
	p.P(`data = encodeVarintPopulate`, p.localName, `(data, uint64(key))`)
	p.P(p.varGen.Next(), ` := r.Int63()`)
	p.P(`if r.Intn(2) == 0 {`)
	p.In()
	p.P(p.varGen.Current(), ` *= -1`)
	p.Out()
	p.P(`}`)
	p.P(`data = encodeVarintPopulate`, p.localName, `(data, uint64(`, p.varGen.Current(), `))`)
	p.Out()
	p.P(`case 1:`)
	p.In()
	p.P(`data = encodeVarintPopulate`, p.localName, `(data, uint64(key))`)
	p.P(`data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))`)
	p.Out()
	p.P(`case 2:`)
	p.In()
	p.P(`data = encodeVarintPopulate`, p.localName, `(data, uint64(key))`)
	p.P(`ll := r.Intn(100)`)
	p.P(`data = encodeVarintPopulate`, p.localName, `(data, uint64(ll))`)
	p.P(`for j := 0; j < ll; j++ {`)
	p.In()
	p.P(`data = append(data, byte(r.Intn(256)))`)
	p.Out()
	p.P(`}`)
	p.Out()
	p.P(`default:`)
	p.In()
	p.P(`data = encodeVarintPopulate`, p.localName, `(data, uint64(key))`)
	p.P(`data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))`)
	p.Out()
	p.P(`}`)
	p.P(`return data`)
	p.Out()
	p.P(`}`)

	p.P(`func encodeVarintPopulate`, p.localName, `(data []byte, v uint64) []byte {`)
	p.In()
	p.P(`for v >= 1<<7 {`)
	p.In()
	p.P(`data = append(data, uint8(uint64(v)&0x7f|0x80))`)
	p.P(`v >>= 7`)
	p.Out()
	p.P(`}`)
	p.P(`data = append(data, uint8(v))`)
	p.P(`return data`)
	p.Out()
	p.P(`}`)

}

func init() {
	generator.RegisterPlugin(NewPlugin())
}

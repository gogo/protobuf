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

The populate plugin also generates a benchmark given it is enabled using one of the following extensions:

  - benchgen
  - benchgen_all

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

  func NewPopulatedB(r randy) *B {
	this := &B{}
	v2 := NewPopulatedA(r)
	this.A = *v2
	if r.Intn(10) != 0 {
		v3 := r.Intn(10)
		this.G = make([]code_google_com_p_gogoprotobuf_test_custom.Uint128, v3)
		for i := 0; i < v3; i++ {
			v4 := code_google_com_p_gogoprotobuf_test_custom.NewPopulatedUint128(r)
			this.G[i] = *v4
		}
	}
	return this
  }

and the following test code:

  func BenchmarkBPopulate(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewPopulatedB(popr)
	}
  }

The idea that is useful for testing.
Most of the other plugins' generated test code uses it.
You will still be able to use the generated test code of other packages
if you tunn off the popluate plugin and write your own custom NewPopulated function.

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
	localNum   string
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
	if typName == "float64" {
		return "r.Float64()"
	} else if typName == "float32" {
		return "r.Float32()"
	} else if typName == "int64" {
		return "r.Int63()"
	} else if typName == "int32" {
		return "r.Int31()"
	} else if typName == "uint32" {
		return "r.Uint32()"
	} else if typName == "uint64" {
		return "uint64(r.Uint32())"
	} else if typName == "bool" {
		return `bool(r.Intn(2) == 0)`
	}
	panic(fmt.Errorf("unexpected type %v", typName))
}

func (p *plugin) GenerateField(message *generator.Descriptor, field *descriptor.FieldDescriptorProto) {
	goTyp, _ := p.GoType(message, field)
	fieldname := generator.CamelCase(*field.Name)
	if field.IsMessage() && gogoproto.IsEmbed(field) {
		fieldname = generator.EmbedFieldName(goTyp)
	}
	goTypName := generator.GoTypeToName(goTyp)
	if field.IsMessage() {
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
			val := fmt.Sprintf("randString%v(r)", p.localNum)
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
				p.Out()
				p.P(`}`)
			} else if gogoproto.IsNullable(field) {
				p.P(p.varGen.Next(), ` := `, value(typName))
				p.P(`this.`, fieldname, ` = &`, p.varGen.Current())
			} else {
				p.P(`this.`, fieldname, ` = `, value(typName))
			}
		}
	}
}

func (p *plugin) hasLoop(field *descriptor.FieldDescriptorProto, visited []*generator.Descriptor, excludes []*generator.Descriptor) *generator.Descriptor {
	if field.IsMessage() {
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

	index := p.PackageIndex(file)
	p.localNum = ""
	if index > 0 {
		p.localNum = fmt.Sprintf("%d", index)
	}

	for _, message := range file.Messages() {
		if !gogoproto.HasPopulate(file.FileDescriptorProto, message.DescriptorProto) {
			continue
		}
		p.atleastOne = true
		ccTypeName := generator.CamelCaseSlice(message.TypeName())
		p.P(`func NewPopulated`, ccTypeName, `(r randy`, p.localNum, `) *`, ccTypeName, ` {`)
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

	p.P(`type randy`, p.localNum, ` interface {`)
	p.In()
	p.P(`Float32() float32`)
	p.P(`Float64() float64`)
	p.P(`Int63() int64`)
	p.P(`Int31() int32`)
	p.P(`Uint32() uint32`)
	p.P(`Intn(n int) int`)
	p.Out()
	p.P(`}`)

	/*
		// unicode/utf8
		// ValidRune reports whether r can be legally encoded as UTF-8.
		// Code points that are out of range or a surrogate half are illegal.
		func ValidRune(r rune) bool {
			switch {
			case r < 0:
				return false
			case surrogateMin <= r && r <= surrogateMax:
				return false
			case r > MaxRune:
				return false
			}
			return true
		}
	*/
	surrogateRange := surrogateMax - surrogateMin
	maxRand := maxRune - surrogateRange

	p.P(`func randUTF8Rune`, p.localNum, `(r randy`, p.localNum, `) rune {`)
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

	p.P(`func randString`, p.localNum, `(r randy`, p.localNum, `) string {`)
	p.In()
	p.P(p.varGen.Next(), ` := r.Intn(100)`)
	p.P(`tmps := make([]rune, `, p.varGen.Current(), `)`)
	p.P(`for i := 0; i < `, p.varGen.Current(), `; i++ {`)
	p.In()
	p.P(`tmps[i] = randUTF8Rune`, p.localNum, `(r)`)
	p.Out()
	p.P(`}`)
	p.P(`return string(tmps)`)
	p.Out()
	p.P(`}`)
}

func init() {
	generator.RegisterPlugin(NewPlugin())
}

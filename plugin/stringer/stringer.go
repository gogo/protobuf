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
The stringgen plugin generates a String method for each message.

It is enabled by the following extensions:

  - stringgen
  - stringgen_all

The stringgen plugin also generates a test given it is enabled using one of the following extensions:

  - testgen
  - testgen_all

And a benchmark given it is enabled using one of the following extensions:

  - benchgen
  - benchgen_all

Let us look at:

  code.google.com/p/gogoprotobuf/test/example/example.proto

Btw all the output can be seen at:

  code.google.com/p/gogoprotobuf/test/example/*

The following message:

option (gogoproto.msgstringmethod_all) = false;
option (gogoproto.stringgen_all) =  true;

message A {
	optional string Description = 1 [(gogoproto.nullable) = false];
	optional int64 Number = 2 [(gogoproto.nullable) = false];
	optional bytes Id = 3 [(gogoproto.customtype) = "code.google.com/p/gogoprotobuf/test/custom.Uuid", (gogoproto.nullable) = false];
}

given to the stringgen stringgen, will generate the following code:

  func (this *A) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&A{`,
		`Description:` + fmt.Sprintf("%v", this.Description) + `,`,
		`Number:` + fmt.Sprintf("%v", this.Number) + `,`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`}`,
	}, "")
	return s
  }

and the following test code:

  func TestAStringGen(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedA(popr)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
  }
  func BenchmarkAStringGen(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedA(popr)
	data := p.String()
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.String()
	}
  }

Typically fmt.Printf("%v") will stop to print when it reaches a pointer and
not print their values, while the generated String method will always print all values, recursively.

*/
package stringer

import (
	"code.google.com/p/gogoprotobuf/gogoproto"
	"code.google.com/p/gogoprotobuf/protoc-gen-gogo/generator"
	"strings"
)

type stringgen struct {
	*generator.Generator
	generator.PluginImports
	atleastOne bool
	localName  string
}

func NewStringGen() *stringgen {
	return &stringgen{}
}

func (p *stringgen) Name() string {
	return "stringgen"
}

func (p *stringgen) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *stringgen) Generate(file *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)
	p.atleastOne = false

	p.localName = generator.FileName(file)

	fmtPkg := p.NewImport("fmt")
	stringsPkg := p.NewImport("strings")
	reflectPkg := p.NewImport("reflect")
	for _, message := range file.Messages() {
		if !gogoproto.IsStringer(file.FileDescriptorProto, message.DescriptorProto) {
			continue
		}
		if gogoproto.HasOldString(file.FileDescriptorProto, message.DescriptorProto) {
			panic("old string method needs to be disabled, please use gogoproto.msgstringmethod or gogoproto.msgstringmethod_all and set it to false")
		}
		p.atleastOne = true
		ccTypeName := generator.CamelCaseSlice(message.TypeName())
		p.P(`func (this *`, ccTypeName, `) String() string {`)
		p.In()
		p.P(`if this == nil {`)
		p.In()
		p.P(`return "nil"`)
		p.Out()
		p.P(`}`)
		p.P("s := ", stringsPkg.Use(), ".Join([]string{`&", ccTypeName, "{`,")
		for _, field := range message.Field {
			nullable := gogoproto.IsNullable(field)
			repeated := field.IsRepeated()
			fieldname := generator.CamelCase(*field.Name)
			if field.IsMessage() || p.IsGroup(field) {
				desc := p.ObjectNamed(field.GetTypeName())
				msgname := p.TypeName(desc)
				msgnames := strings.Split(msgname, ".")
				typeName := msgnames[len(msgnames)-1]
				if gogoproto.IsEmbed(field) {
					fieldname = typeName
				}
				if nullable {
					p.P("`", fieldname, ":`", ` + `, stringsPkg.Use(), `.Replace(`, fmtPkg.Use(), `.Sprintf("%v", this.`, fieldname, `), "`, typeName, `","`, msgname, `"`, ", 1) + `,", "`,")
				} else if repeated {
					p.P("`", fieldname, ":`", ` + `, stringsPkg.Use(), `.Replace(`, stringsPkg.Use(), `.Replace(`, fmtPkg.Use(), `.Sprintf("%v", this.`, fieldname, `), "`, typeName, `","`, msgname, `"`, ", 1),`&`,``,1) + `,", "`,")
				} else {
					p.P("`", fieldname, ":`", ` + `, stringsPkg.Use(), `.Replace(`, stringsPkg.Use(), `.Replace(this.`, fieldname, `.String(), "`, typeName, `","`, msgname, `"`, ", 1),`&`,``,1) + `,", "`,")
				}
			} else {
				if nullable && !repeated {
					p.P("`", fieldname, ":`", ` + valueToString`, p.localName, `(this.`, fieldname, ") + `,", "`,")
				} else {
					p.P("`", fieldname, ":`", ` + `, fmtPkg.Use(), `.Sprintf("%v", this.`, fieldname, ") + `,", "`,")
				}
			}

		}
		p.P("`}`,")
		p.P(`}`, `,""`, ")")
		p.P(`return s`)
		p.Out()
		p.P(`}`)
	}

	if !p.atleastOne {
		return
	}

	p.P(`func valueToString`, p.localName, `(v interface{}) string {`)
	p.In()
	p.P(`rv := `, reflectPkg.Use(), `.ValueOf(v)`)
	p.P(`if rv.IsNil() {`)
	p.In()
	p.P(`return "nil"`)
	p.Out()
	p.P(`}`)
	p.P(`pv := `, reflectPkg.Use(), `.Indirect(rv).Interface()`)
	p.P(`return `, fmtPkg.Use(), `.Sprintf("*%v", pv)`)
	p.Out()
	p.P(`}`)

}

func init() {
	generator.RegisterPlugin(NewStringGen())
}

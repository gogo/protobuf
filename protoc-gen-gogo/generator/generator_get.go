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

package generator

import (
	"bytes"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"

	"code.google.com/p/gogoprotobuf/gogoproto"
	"code.google.com/p/gogoprotobuf/proto"
	descriptor "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"
	plugin "code.google.com/p/gogoprotobuf/protoc-gen-gogo/plugin"
	"errors"
)

func (d *FileDescriptor) Messages() []*Descriptor {
	return d.desc
}

func (d *Descriptor) IsGroup() bool {
	return d.group
}

func (g *Generator) TypeNameByObject(typeName string) Object {
	o, ok := g.typeNameToObject[typeName]
	if !ok {
		g.Fail("can't find object with type", typeName)
	}
	return o
}

func (g *Generator) GeneratePlugin(p Plugin) {
	p.Init(g)
	// Generate the output. The generator runs for every file, even the files
	// that we don't generate output for, so that we can collate the full list
	// of exported symbols to support public imports.
	genFileMap := make(map[*FileDescriptor]bool, len(g.genFiles))
	for _, file := range g.genFiles {
		genFileMap[file] = true
	}
	i := 0
	for _, file := range g.allFiles {
		g.Reset()
		g.generatePlugin(file, p)
		if _, ok := genFileMap[file]; !ok {
			continue
		}
		g.Response.File[i] = new(plugin.CodeGeneratorResponse_File)
		g.Response.File[i].Name = proto.String(goFileName(*file.Name))
		g.Response.File[i].Content = proto.String(g.String())
		i++
	}
}

func (g *Generator) generatePlugin(file *FileDescriptor, p Plugin) {
	g.file = g.FileOf(file.FileDescriptorProto)
	g.usedPackages = make(map[string]bool)

	// Run the plugins before the imports so we know which imports are necessary.
	p.Generate(file)

	// Generate header and imports last, though they appear first in the output.
	rem := g.Buffer
	g.Buffer = new(bytes.Buffer)
	g.generateHeader()
	p.GenerateImports(g.file)
	g.Write(rem.Bytes())

	// Reformat generated code.
	fset := token.NewFileSet()
	ast, err := parser.ParseFile(fset, "", g, parser.ParseComments)
	if err != nil {
		g.Fail("bad Go source code was generated:", err.Error())
		return
	}
	g.Reset()
	err = (&printer.Config{Mode: printer.TabIndent | printer.UseSpaces, Tabwidth: 8}).Fprint(g, fset, ast)
	if err != nil {
		g.Fail("generated Go source code could not be reformatted:", err.Error())
	}
}

func isNullable(field *descriptor.FieldDescriptorProto) bool {
	return proto.GetBoolExtension(field.Options, gogoproto.E_Nullable, true)
}

func isEmbed(field *descriptor.FieldDescriptorProto) bool {
	return proto.GetBoolExtension(field.Options, gogoproto.E_Embed, false)
}

func isCustomType(field *descriptor.FieldDescriptorProto) bool {
	if field.Options != nil {
		v, err := proto.GetExtension(field.Options, gogoproto.E_Customtype)
		if err == nil && v.(*string) != nil {
			if len(*(v.(*string))) > 0 {
				return true
			}
		}
	}
	return false
}

func getCustomType(field *descriptor.FieldDescriptorProto) (packageName string, typ string, err error) {
	if field.Options != nil {
		v, err := proto.GetExtension(field.Options, gogoproto.E_Customtype)
		if err == nil && v.(*string) != nil {
			ctype := *(v.(*string))
			ss := strings.Split(ctype, ".")
			if len(ss) == 1 {
				return "", typ, nil
			} else if len(ss) >= 2 {
				packageName := strings.Join(ss[0:len(ss)-1], ".")
				typeName := ss[len(ss)-1]
				importStr := strings.Replace(strings.Replace(packageName, "/", "_", -1), ".", "_", -1)
				typ = importStr + "." + typeName
				return typ, packageName, nil
			} else {
				err = errors.New("custom type requires a dot-separated import and type")
			}
		}
	}
	return "", "", err
}

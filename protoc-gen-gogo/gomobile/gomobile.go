// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2015 The Go Authors.  All rights reserved.
// https://github.com/golang/protobuf
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
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
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

// Package gomobile outputs gRPC service descriptions in Go code.
// It runs as a plugin for the Go protocol buffer compiler plugin.
// It is linked in to protoc-gen-go.
package gomobile

import (
	"fmt"
	"os"
	"strings"

	pb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

// generatedCodeVersion indicates a version of the generated code.
// It is incremented whenever an incompatibility between the generated code and
// the gomobile package is introduced; the generated code references
// a constant, gomobile.SupportPackageIsVersionN (where N is generatedCodeVersion).
const generatedCodeVersion = 4

func init() {
	generator.RegisterPlugin(new(gomobile))
}

// gomobile is an implementation of the Go protocol buffer compiler's
// plugin architecture.  It generates bindings for gRPC support.
type gomobile struct {
	gen *generator.Generator
}

// Name returns the name of this plugin, "gomobile".
func (g *gomobile) Name() string {
	return "gomobile"
}

var (
	pbPkg string
)

// Init initializes the plugin.
func (g *gomobile) Init(gen *generator.Generator) {
	g.gen = gen
}

// Given a type name defined in a .proto, return its object.
// Also record that we're using it, to guarantee the associated import.
func (g *gomobile) objectNamed(name string) generator.Object {
	g.gen.RecordTypeUse(name)
	return g.gen.ObjectNamed(name)
}

// Given a type name defined in a .proto, return its name as we will print it.
func (g *gomobile) typeName(str string) string {
	return g.gen.TypeName(g.objectNamed(str))
}

// P forwards to g.gen.P.
func (g *gomobile) P(args ...interface{}) { g.gen.P(args...) }

// Generate generates code for the services in the given file.
func (g *gomobile) Generate(file *generator.FileDescriptor) {
	//fmt.Printf("importPath: %s\npackageName: %s\n", file.SourceCodeInfo.Location.)
	if len(file.FileDescriptorProto.Service) == 0 {
		return
	}

	if os.Getenv("PACKAGE_PATH") == "" {
		g.gen.Fail("neeed to set PACKAGE_PATH env var")
	}

	pbPkg = string(g.gen.AddImport(generator.GoImportPath(os.Getenv("PACKAGE_PATH"))))

	// Assert version compatibility.
	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible with the gomobile package it is being compiled against.")
	g.P()

	for i, service := range file.FileDescriptorProto.Service {
		g.generateService(file, service, i)
	}
}

// GenerateImports generates the import declaration for this file.
func (g *gomobile) GenerateImports(file *generator.FileDescriptor) {}

// reservedClientName records whether a client name is reserved on the client side.
var reservedClientName = map[string]bool{
	// TODO: do we need any in gRPC?
}

func unexport(s string) string { return strings.ToLower(s[:1]) + s[1:] }

// deprecationComment is the standard comment added to deprecated
// messages, fields, enums, and enum values.
var deprecationComment = "// Deprecated: Do not use."

// generateService generates all the code for the named service.
func (g *gomobile) generateService(file *generator.FileDescriptor, service *pb.ServiceDescriptorProto, index int) {
	path := fmt.Sprintf("6,%d", index) // 6 means service.

	origServName := service.GetName()
	fullServName := origServName
	if pkg := file.GetPackage(); pkg != "" {
		fullServName = pkg + "." + fullServName
	}
	handlerName := generator.CamelCase(origServName)
	deprecated := service.GetOptions().GetDeprecated()

	// Handler interface.
	handlerType := handlerName + "Handler"
	g.P("// ", handlerType, " is the handler API for ", handlerName, " service.")
	if deprecated {
		g.P("//")
		g.P(deprecationComment)
	}

	handlerVar := strings.ToLower(handlerType[0:1]) + handlerType[1:]
	g.P("var ", handlerVar, " ", handlerType)

	g.P("type ", handlerType, " interface {")
	for i, method := range service.Method {
		g.gen.PrintComments(fmt.Sprintf("%s,2,%d", path, i)) // 2 means method in a service.
		g.P(g.generateHandlerSignature(handlerName, method))
	}
	g.P("}")
	g.P()

	// Handler Unimplemented struct for forward compatability.
	if deprecated {
		g.P(deprecationComment)
	}

	// Handler registration.
	if deprecated {
		g.P(deprecationComment)
	}

	g.P("func register", handlerName, "Handler(srv ", handlerType, ") {")
	g.P(handlerVar, " = srv")
	g.P("}")
	g.P()

	// Handler handler implementations.
	var handlerNames []string
	for _, method := range service.Method {
		hname := g.generateHandlerMethod(handlerVar, method)
		handlerNames = append(handlerNames, hname)
	}

	g.P()

	g.P("func CommandAsync(cmd string, data []byte, callback func(data []byte)) {")
	g.P("go func() {")
	g.P("var cd []byte")
	g.P("switch cmd {")
	for _, method := range service.Method {
		methName := generator.CamelCase(method.GetName())

		g.P("case \"", methName, "\":")
		g.P("cd = ", methName, "(data)")
	}
	g.P(`default: log.Errorf("unknown command type: %s\n", cmd)`)
	g.P("}")
	g.P("if callback != nil { callback(cd) }")
	g.P("}()")
	g.P("}")

	g.P()

	g.P("type MessageHandler interface {")
	g.P("Handle(b []byte)")
	g.P("}")
	g.P()
	g.P("func CommandMobile(cmd string, data []byte, callback MessageHandler) {")
	g.P("CommandAsync(cmd, data, callback.Handle)")
	g.P("}")
	g.P()
}

// generateHandlerSignatureWithParamNames returns the handler-side signature for a method with parameter names.
func (g *gomobile) generateHandlerSignatureWithParamNames(handlerName string, method *pb.MethodDescriptorProto) string {
	origMethName := method.GetName()
	methName := generator.CamelCase(origMethName)
	if reservedClientName[methName] {
		methName += "_"
	}

	var reqArgs []string
	ret := "error"
	if !method.GetServerStreaming() && !method.GetClientStreaming() {
		ret = "*" + pbPkg + "." + g.typeName(method.GetOutputType()) + ""
	}
	if !method.GetClientStreaming() {
		reqArgs = append(reqArgs, "req *"+pbPkg+"."+g.typeName(method.GetInputType()))
	}

	return methName + "(" + strings.Join(reqArgs, ", ") + ") " + ret
}

// generateHandlerSignature returns the handler-side signature for a method.
func (g *gomobile) generateHandlerSignature(handlerName string, method *pb.MethodDescriptorProto) string {
	origMethName := method.GetName()
	methName := generator.CamelCase(origMethName)
	if reservedClientName[methName] {
		methName += "_"
	}

	var reqArgs []string
	ret := ""
	if !method.GetServerStreaming() && !method.GetClientStreaming() {
		ret = "*" + pbPkg + "." + g.typeName(method.GetOutputType())
	}

	if !method.GetClientStreaming() {
		reqArgs = append(reqArgs, "*"+pbPkg+"."+g.typeName(method.GetInputType()))
	}

	return methName + "(" + strings.Join(reqArgs, ", ") + ") " + ret
}

func (g *gomobile) generateHandlerMethod(handlerVar string, method *pb.MethodDescriptorProto) string {
	methName := generator.CamelCase(method.GetName())
	hname := fmt.Sprintf("%s", methName)
	inType := g.typeName(method.GetInputType())
	outType := g.typeName(method.GetOutputType())
	errorPrefix := "_Error"
	if os.Getenv("GOGO_NO_UNDERSCORE") == "1" {
		errorPrefix = "Error"
	}
	if !method.GetServerStreaming() && !method.GetClientStreaming() {
		g.P("func ", hname, "(b []byte) ([]byte) {")
		g.P("in := new(", pbPkg+"."+inType, ")")
		g.P("if err := in.Unmarshal(b); err != nil { ")
		g.P("resp, _ := (&", pbPkg, ".", outType, "{Error: &", pbPkg, ".", outType, errorPrefix, "{Code: ", pbPkg, ".", outType, errorPrefix, "_BAD_INPUT, Description: err.Error()}}).Marshal()")
		g.P("return resp")
		g.P("}")

		g.P("resp, _ := ", handlerVar, ".", methName, "(in).Marshal()")
		g.P("return resp")

		g.P("}")
		g.P()
		return hname
	}
	return hname
}

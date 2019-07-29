// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2015, The GoGo Authors. All rights reserved.
// http://github.com/buptbill220/protobuf
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

package command

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"strings"

	_ "github.com/buptbill220/protobuf/plugin/compare"
	_ "github.com/buptbill220/protobuf/plugin/defaultcheck"
	_ "github.com/buptbill220/protobuf/plugin/description"
	_ "github.com/buptbill220/protobuf/plugin/embedcheck"
	_ "github.com/buptbill220/protobuf/plugin/enumstringer"
	_ "github.com/buptbill220/protobuf/plugin/equal"
	_ "github.com/buptbill220/protobuf/plugin/face"
	_ "github.com/buptbill220/protobuf/plugin/gostring"
	_ "github.com/buptbill220/protobuf/plugin/marshalto"
	_ "github.com/buptbill220/protobuf/plugin/oneofcheck"
	_ "github.com/buptbill220/protobuf/plugin/populate"
	_ "github.com/buptbill220/protobuf/plugin/size"
	_ "github.com/buptbill220/protobuf/plugin/stringer"
	"github.com/buptbill220/protobuf/plugin/testgen"
	_ "github.com/buptbill220/protobuf/plugin/union"
	_ "github.com/buptbill220/protobuf/plugin/unmarshal"
	_ "github.com/buptbill220/protobuf/plugin/jsonmarshal"
	_ "github.com/buptbill220/protobuf/plugin/validate"
	_ "github.com/buptbill220/protobuf/plugin/pyjsonmarshal"
	"github.com/buptbill220/protobuf/proto"
	"github.com/buptbill220/protobuf/protoc-gen-gogo/generator"
	_ "github.com/buptbill220/protobuf/protoc-gen-gogo/grpc"
	plugin "github.com/buptbill220/protobuf/protoc-gen-gogo/plugin"
)


func Read() *plugin.CodeGeneratorRequest {
	g := generator.New()
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		g.Error(err, "reading input")
	}

	if err := proto.Unmarshal(data, g.Request); err != nil {
		g.Error(err, "parsing input proto")
	}

	if len(g.Request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}
	return g.Request
}

// filenameSuffix replaces the .pb.go at the end of each filename.
func GeneratePlugin(req *plugin.CodeGeneratorRequest, p generator.Plugin, filenameSuffix string) *plugin.CodeGeneratorResponse {
	g := generator.New()
	g.Request = req
	if len(g.Request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}

	g.CommandLineParameters(g.Request.GetParameter())

	g.WrapTypes()
	g.SetPackageNames()
	g.BuildTypeNameMap()
	g.GeneratePlugin(p)

	for i := 0; i < len(g.Response.File); i++ {
		g.Response.File[i].Name = proto.String(
			strings.Replace(*g.Response.File[i].Name, ".pb.go", filenameSuffix, -1),
		)
	}
	if err := goformat(g.Response); err != nil {
		g.Error(err)
	}
	return g.Response
}

func goformat(resp *plugin.CodeGeneratorResponse) error {
	for i := 0; i < len(resp.File); i++ {
		var formatted []byte
		var err error
		if !generator.IsPyFmq {
			formatted, err = format.Source([]byte(resp.File[i].GetContent()))
			if err != nil {
				return fmt.Errorf("go format error: %v", err)
			}
		} else {
			formatted = []byte(resp.File[i].GetContent())
		}
		fmts := string(formatted)
		resp.File[i].Content = &fmts
	}
	return nil
}

func Generate(req *plugin.CodeGeneratorRequest) *plugin.CodeGeneratorResponse {
	// Begin by allocating a generator. The request and response structures are stored there
	// so we can do error handling easily - the response structure contains the field to
	// report failure.
	g := generator.New()
	g.Request = req

	g.CommandLineParameters(g.Request.GetParameter())

	// Create a wrapped version of the Descriptors and EnumDescriptors that
	// point to the file that defines them.
	g.WrapTypes()

	g.SetPackageNames()
	g.BuildTypeNameMap()

	g.GenerateAllFiles()

	if err := goformat(g.Response); err != nil {
		g.Error(err)
	}

	if !generator.IsPyFmq {
		testReq := proto.Clone(req).(*plugin.CodeGeneratorRequest)
		
		testResp := GeneratePlugin(testReq, testgen.NewPlugin(), "pb_test.go")
		
		for i := 0; i < len(testResp.File); i++ {
			if strings.Contains(*testResp.File[i].Content, `//These tests are generated by github.com/buptbill220/protobuf/plugin/testgen`) {
				g.Response.File = append(g.Response.File, testResp.File[i])
			}
		}
	}

	return g.Response
}

func Write(resp *plugin.CodeGeneratorResponse) {
	g := generator.New()
	// Send back the results.
	data, err := proto.Marshal(resp)
	if err != nil {
		g.Error(err, "failed to marshal output proto")
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		g.Error(err, "failed to write output proto")
	}
}

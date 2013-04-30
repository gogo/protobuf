// Extensions for Protocol Buffers to create more go like structures.
//
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

package google_protobuf

import (
	"fmt"
	"reflect"
	"strings"
)

func goStringValue(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

func (this *FileDescriptorSet) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&google_protobuf.FileDescriptorSet{` + `File:` + fmt.Sprintf("%#v", this.File) + `}`}, ", ")
	return s
}
func (this *FileDescriptorProto) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&google_protobuf.FileDescriptorProto{` + `Name:` + goStringValue(this.Name, "string"), `Package:` + goStringValue(this.Package, "string"), `Dependency:` + fmt.Sprintf("%#v", this.Dependency), `PublicDependency:` + fmt.Sprintf("%#v", this.PublicDependency), `WeakDependency:` + fmt.Sprintf("%#v", this.WeakDependency), `MessageType:` + fmt.Sprintf("%#v", this.MessageType), `EnumType:` + fmt.Sprintf("%#v", this.EnumType), `Service:` + fmt.Sprintf("%#v", this.Service), `Extension:` + fmt.Sprintf("%#v", this.Extension), `Options:` + fmt.Sprintf("%#v", this.Options), `SourceCodeInfo:` + fmt.Sprintf("%#v", this.SourceCodeInfo) + `}`}, ", ")
	return s
}
func (this *DescriptorProto) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&google_protobuf.DescriptorProto{` + `Name:` + goStringValue(this.Name, "string"), `Field:` + fmt.Sprintf("%#v", this.Field), `Extension:` + fmt.Sprintf("%#v", this.Extension), `NestedType:` + fmt.Sprintf("%#v", this.NestedType), `EnumType:` + fmt.Sprintf("%#v", this.EnumType), `ExtensionRange:` + fmt.Sprintf("%#v", this.ExtensionRange), `Options:` + fmt.Sprintf("%#v", this.Options) + `}`}, ", ")
	return s
}
func (this *DescriptorProto_ExtensionRange) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&google_protobuf.DescriptorProto_ExtensionRange{` + `Start:` + goStringValue(this.Start, "int32"), `End:` + goStringValue(this.End, "int32") + `}`}, ", ")
	return s
}
func (this *FieldDescriptorProto) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&google_protobuf.FieldDescriptorProto{` + `Name:` + goStringValue(this.Name, "string"), `Number:` + goStringValue(this.Number, "int32"), `Label:` + goStringValue(this.Label, "google_protobuf.FieldDescriptorProto_Label"), `Type:` + goStringValue(this.Type, "google_protobuf.FieldDescriptorProto_Type"), `TypeName:` + goStringValue(this.TypeName, "string"), `Extendee:` + goStringValue(this.Extendee, "string"), `DefaultValue:` + goStringValue(this.DefaultValue, "string"), `Options:` + fmt.Sprintf("%#v", this.Options) + `}`}, ", ")
	return s
}
func (this *EnumDescriptorProto) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&google_protobuf.EnumDescriptorProto{` + `Name:` + goStringValue(this.Name, "string"), `Value:` + fmt.Sprintf("%#v", this.Value), `Options:` + fmt.Sprintf("%#v", this.Options) + `}`}, ", ")
	return s
}
func (this *EnumValueDescriptorProto) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&google_protobuf.EnumValueDescriptorProto{` + `Name:` + goStringValue(this.Name, "string"), `Number:` + goStringValue(this.Number, "int32"), `Options:` + fmt.Sprintf("%#v", this.Options) + `}`}, ", ")
	return s
}
func (this *ServiceDescriptorProto) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&google_protobuf.ServiceDescriptorProto{` + `Name:` + goStringValue(this.Name, "string"), `Method:` + fmt.Sprintf("%#v", this.Method), `Options:` + fmt.Sprintf("%#v", this.Options) + `}`}, ", ")
	return s
}
func (this *MethodDescriptorProto) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&google_protobuf.MethodDescriptorProto{` + `Name:` + goStringValue(this.Name, "string"), `InputType:` + goStringValue(this.InputType, "string"), `OutputType:` + goStringValue(this.OutputType, "string"), `Options:` + fmt.Sprintf("%#v", this.Options) + `}`}, ", ")
	return s
}
func (this *FileOptions) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&google_protobuf.FileOptions{` + `JavaPackage:` + goStringValue(this.JavaPackage, "string"), `JavaOuterClassname:` + goStringValue(this.JavaOuterClassname, "string"), `JavaMultipleFiles:` + goStringValue(this.JavaMultipleFiles, "bool"), `JavaGenerateEqualsAndHash:` + goStringValue(this.JavaGenerateEqualsAndHash, "bool"), `OptimizeFor:` + goStringValue(this.OptimizeFor, "google_protobuf.FileOptions_OptimizeMode"), `GoPackage:` + goStringValue(this.GoPackage, "string"), `CcGenericServices:` + goStringValue(this.CcGenericServices, "bool"), `JavaGenericServices:` + goStringValue(this.JavaGenericServices, "bool"), `PyGenericServices:` + goStringValue(this.PyGenericServices, "bool"), `UninterpretedOption:` + fmt.Sprintf("%#v", this.UninterpretedOption) + `}`}, ", ")
	return s
}
func (this *MessageOptions) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&google_protobuf.MessageOptions{` + `MessageSetWireFormat:` + goStringValue(this.MessageSetWireFormat, "bool"), `NoStandardDescriptorAccessor:` + goStringValue(this.NoStandardDescriptorAccessor, "bool"), `UninterpretedOption:` + fmt.Sprintf("%#v", this.UninterpretedOption) + `}`}, ", ")
	return s
}
func (this *FieldOptions) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&google_protobuf.FieldOptions{` + `Ctype:` + goStringValue(this.Ctype, "google_protobuf.FieldOptions_CType"), `Packed:` + goStringValue(this.Packed, "bool"), `Lazy:` + goStringValue(this.Lazy, "bool"), `Deprecated:` + goStringValue(this.Deprecated, "bool"), `ExperimentalMapKey:` + goStringValue(this.ExperimentalMapKey, "string"), `Weak:` + goStringValue(this.Weak, "bool"), `UninterpretedOption:` + fmt.Sprintf("%#v", this.UninterpretedOption) + `}`}, ", ")
	return s
}
func (this *EnumOptions) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&google_protobuf.EnumOptions{` + `AllowAlias:` + goStringValue(this.AllowAlias, "bool"), `UninterpretedOption:` + fmt.Sprintf("%#v", this.UninterpretedOption) + `}`}, ", ")
	return s
}
func (this *EnumValueOptions) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&google_protobuf.EnumValueOptions{` + `UninterpretedOption:` + fmt.Sprintf("%#v", this.UninterpretedOption) + `}`}, ", ")
	return s
}
func (this *ServiceOptions) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&google_protobuf.ServiceOptions{` + `UninterpretedOption:` + fmt.Sprintf("%#v", this.UninterpretedOption) + `}`}, ", ")
	return s
}
func (this *MethodOptions) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&google_protobuf.MethodOptions{` + `UninterpretedOption:` + fmt.Sprintf("%#v", this.UninterpretedOption) + `}`}, ", ")
	return s
}
func (this *UninterpretedOption) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&google_protobuf.UninterpretedOption{` + `Name:` + fmt.Sprintf("%#v", this.Name), `IdentifierValue:` + goStringValue(this.IdentifierValue, "string"), `PositiveIntValue:` + goStringValue(this.PositiveIntValue, "uint64"), `NegativeIntValue:` + goStringValue(this.NegativeIntValue, "int64"), `DoubleValue:` + goStringValue(this.DoubleValue, "float64"), `StringValue:` + goStringValue(this.StringValue, "[]byte"), `AggregateValue:` + goStringValue(this.AggregateValue, "string") + `}`}, ", ")
	return s
}
func (this *UninterpretedOption_NamePart) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&google_protobuf.UninterpretedOption_NamePart{` + `NamePart:` + goStringValue(this.NamePart, "string"), `IsExtension:` + goStringValue(this.IsExtension, "bool") + `}`}, ", ")
	return s
}
func (this *SourceCodeInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&google_protobuf.SourceCodeInfo{` + `Location:` + fmt.Sprintf("%#v", this.Location) + `}`}, ", ")
	return s
}
func (this *SourceCodeInfo_Location) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&google_protobuf.SourceCodeInfo_Location{` + `Path:` + fmt.Sprintf("%#v", this.Path), `Span:` + fmt.Sprintf("%#v", this.Span), `LeadingComments:` + goStringValue(this.LeadingComments, "string"), `TrailingComments:` + goStringValue(this.TrailingComments, "string") + `}`}, ", ")
	return s
}

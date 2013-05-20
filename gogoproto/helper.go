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

package gogoproto

import google_protobuf "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"
import proto "code.google.com/p/gogoprotobuf/proto"

func IsEmbed(field *google_protobuf.FieldDescriptorProto) bool {
	return proto.GetBoolExtension(field.Options, E_Embed, false)
}

func IsNullable(field *google_protobuf.FieldDescriptorProto) bool {
	return proto.GetBoolExtension(field.Options, E_Nullable, true)
}

func IsCustomType(field *google_protobuf.FieldDescriptorProto) bool {
	typ := GetCustomType(field)
	if len(typ) > 0 {
		return true
	}
	return false
}

func GetCustomType(field *google_protobuf.FieldDescriptorProto) string {
	if field.Options != nil {
		v, err := proto.GetExtension(field.Options, E_Customtype)
		if err == nil && v.(*string) != nil {
			return *(v.(*string))
		}
	}
	return ""
}

func HasGetters(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Getters, proto.GetBoolExtension(file.Options, E_GettersAll, true))
}

func IsUnion(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Union, proto.GetBoolExtension(file.Options, E_UnionAll, false))
}

func HasGoString(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Gostring, proto.GetBoolExtension(file.Options, E_GostringAll, false))
}

func HasEqual(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Equal, proto.GetBoolExtension(file.Options, E_EqualAll, false))
}

func HasVerboseEqual(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_VerboseEqual, proto.GetBoolExtension(file.Options, E_VerboseEqualAll, false))
}

func HasString(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Stringgen, proto.GetBoolExtension(file.Options, E_StringgenAll, false))
}

func HasOldString(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Msgstringmethod, proto.GetBoolExtension(file.Options, E_MsgstringmethodAll, true))
}

func IsFace(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Face, proto.GetBoolExtension(file.Options, E_FaceAll, false))
}

func HasDescription(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Description, proto.GetBoolExtension(file.Options, E_DescriptionAll, false))
}

func HasPopulate(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Populate, proto.GetBoolExtension(file.Options, E_PopulateAll, false))
}

func HasTestGen(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Testgen, proto.GetBoolExtension(file.Options, E_TestgenAll, false))
}

func HasBenchGen(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Benchgen, proto.GetBoolExtension(file.Options, E_BenchgenAll, false))
}

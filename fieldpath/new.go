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

package fieldpath

import (
	descriptor "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"
	"strings"
)

type fdesc struct {
	rootPkg string
	rootMsg string
	fields  []*descriptor.FieldDescriptorProto
	field   *descriptor.FieldDescriptorProto
	path    []uint64
}

func new(rootPackage string, rootMessage string, descSet *descriptor.FileDescriptorSet, path string) (*fdesc, error) {
	fieldpaths := strings.Split(path, ".")
	keys := make([]uint64, len(fieldpaths))
	fields := make([]*descriptor.FieldDescriptorProto, len(fieldpaths))

	curPackage := rootPackage
	curMessage := rootMessage
	last := len(fieldpaths) - 1
	var fieldDesc *descriptor.FieldDescriptorProto
	for i, f := range fieldpaths {
		fieldName := f
		fieldDesc = descSet.GetField(curPackage, curMessage, fieldName)
		if fieldDesc == nil {
			curPackage, fieldDesc = descSet.FindExtension(curPackage, curMessage, fieldName)
			if fieldDesc == nil {
				return nil, &errChild{fieldName: fieldName, pkg: curPackage, msg: curMessage}
			}
			typeNames := strings.Split(fieldDesc.GetTypeName(), ".")
			curMessage = fieldDesc.GetTypeName()
			if len(typeNames) > 1 {
				curPackage = typeNames[1]
				curMessage = typeNames[2]
			}
			fieldKey := fieldDesc.GetKeyUint64()
			keys[i] = fieldKey
			fields[i] = fieldDesc
		} else {
			fieldKey := fieldDesc.GetKeyUint64()
			if fieldDesc.IsMessage() {
				curPackage, curMessage = descSet.FindMessage(curPackage, curMessage, fieldName)
			} else if i != last {
				return nil, &errMessage{fieldName}
			}
			keys[i] = fieldKey
			fields[i] = fieldDesc
		}
	}
	return &fdesc{curPackage, curMessage, fields, fieldDesc, keys}, nil
}

//only for testing
func TestNew(rootPackage string, rootMessage string, descSet *descriptor.FileDescriptorSet, path string) ([]uint64, *descriptor.FieldDescriptorProto, error) {
	fd, err := new(rootPackage, rootMessage, descSet, path)
	if err != nil {
		return nil, nil, err
	}
	return fd.path, fd.field, nil
}

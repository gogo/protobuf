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

package google_protobuf

func (field *FieldDescriptorProto) WireType() (wire int) {
	switch *field.Type {
	case FieldDescriptorProto_TYPE_DOUBLE:
		return 1
	case FieldDescriptorProto_TYPE_FLOAT:
		return 5
	case FieldDescriptorProto_TYPE_INT64:
		return 0
	case FieldDescriptorProto_TYPE_UINT64:
		return 0
	case FieldDescriptorProto_TYPE_INT32:
		return 0
	case FieldDescriptorProto_TYPE_UINT32:
		return 0
	case FieldDescriptorProto_TYPE_FIXED64:
		return 1
	case FieldDescriptorProto_TYPE_FIXED32:
		return 5
	case FieldDescriptorProto_TYPE_BOOL:
		return 0
	case FieldDescriptorProto_TYPE_STRING:
		return 2
	case FieldDescriptorProto_TYPE_GROUP:
		return 2
	case FieldDescriptorProto_TYPE_MESSAGE:
		return 2
	case FieldDescriptorProto_TYPE_BYTES:
		return 2
	case FieldDescriptorProto_TYPE_ENUM:
		return 0
	case FieldDescriptorProto_TYPE_SFIXED32:
		return 5
	case FieldDescriptorProto_TYPE_SFIXED64:
		return 1
	case FieldDescriptorProto_TYPE_SINT32:
		return 0
	case FieldDescriptorProto_TYPE_SINT64:
		return 0
	}
	panic("unreachable")
}

func (field *FieldDescriptorProto) GetKey() []byte {
	packed := field.IsPacked()
	wireType := field.WireType()
	fieldNumber := field.GetNumber()
	if packed {
		wireType = 2
	}
	x := uint32(fieldNumber)<<3 | uint32(wireType)
	i := 0
	keybuf := make([]byte, 0)
	for i = 0; x > 127; i++ {
		keybuf = append(keybuf, 0x80|uint8(x&0x7F))
		x >>= 7
	}
	keybuf = append(keybuf, uint8(x))
	return keybuf
}

func (desc *FileDescriptorSet) GetField(packageName, messageName, fieldName string) *FieldDescriptorProto {
	msg := desc.GetMessage(packageName, messageName)
	if msg == nil {
		return nil
	}
	for _, field := range msg.GetField() {
		if field.GetName() == fieldName {
			return field
		}
	}
	return nil
}

func (file *FileDescriptorProto) GetMessage(typeName string) *DescriptorProto {
	for _, msg := range file.GetMessageType() {
		if msg.GetName() == typeName {
			return msg
		}
	}
	return nil
}

func (desc *FileDescriptorSet) GetMessage(packageName string, typeName string) *DescriptorProto {
	for _, file := range desc.GetFile() {
		if file.GetPackage() != packageName {
			continue
		}
		for _, msg := range file.GetMessageType() {
			if msg.GetName() == typeName {
				return msg
			}
		}
	}
	return nil
}

func (desc *FileDescriptorSet) GetEnum(packageName string, typeName string) *EnumDescriptorProto {
	for _, file := range desc.GetFile() {
		if file.GetPackage() != packageName {
			continue
		}
		for _, enum := range file.GetEnumType() {
			if enum.GetName() == typeName {
				return enum
			}
		}
	}
	return nil
}

func (f *FieldDescriptorProto) IsEnum() bool {
	return *f.Type == FieldDescriptorProto_TYPE_ENUM
}

func (f *FieldDescriptorProto) IsMessage() bool {
	return *f.Type == FieldDescriptorProto_TYPE_MESSAGE
}

func (f *FieldDescriptorProto) IsBytes() bool {
	return *f.Type == FieldDescriptorProto_TYPE_BYTES
}

func (f *FieldDescriptorProto) IsRepeated() bool {
	return f.Label != nil && *f.Label == FieldDescriptorProto_LABEL_REPEATED
}

func (f *FieldDescriptorProto) IsString() bool {
	return *f.Type == FieldDescriptorProto_TYPE_STRING
}

func (f *FieldDescriptorProto) IsRequired() bool {
	return f.Label != nil && *f.Label == FieldDescriptorProto_LABEL_REQUIRED
}

func (f *FieldDescriptorProto) IsPacked() bool {
	return f.Options != nil && f.GetOptions().GetPacked()
}

func (m *DescriptorProto) HasExtension() bool {
	return len(m.ExtensionRange) > 0
}

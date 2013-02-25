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

package test

import (
	"testing"
)

func TestGobGoGoSimple(t *testing.T) {
	testGobGoGo{
		msg:    gogosimplemsg,
		newMsg: newGoGoProtoSimpleMsg,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestGobVisNidGoGoOptNative(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNidOptNative,
		newMsg: newGoGoProtoNidOptNative,
	}.test(t)
}

//Optional Native
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestGobVimNidGoGoOptNative(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNidOptNative,
		newMsg: newGoGoProtoNidOptNative,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestGobVisNinGoGoOptNative(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNinOptNative,
		newMsg: newGoGoProtoNinOptNative,
	}.test(t)
}

//Optional Native
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestGobVimNinGoGoOptNative(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNinOptNative,
		newMsg: newGoGoProtoNinOptNative,
	}.test(t)
}

//Repeated Native
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestGobVisNidGoGoRepNative(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNidRepNative,
		newMsg: newGoGoProtoNidRepNative,
	}.test(t)
}

//Repeated Native
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestGobVimNidGoGoRepNative(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNidRepNative,
		newMsg: newGoGoProtoNidRepNative,
	}.test(t)
}

//Repeated Native
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestGobVisNinGoGoRepNative(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNinRepNative,
		newMsg: newGoGoProtoNinRepNative,
	}.test(t)
}

//Repeated Native
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestGobVimNinGoGoRepNative(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNinRepNative,
		newMsg: newGoGoProtoNinRepNative,
	}.test(t)
}

//Packed Repeated Native
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestGobVisNidGoGoRepPackedNative(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNidRepPackedNative,
		newMsg: newGoGoProtoNidRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestGobVimNidGoGoRepPackedNative(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNidRepPackedNative,
		newMsg: newGoGoProtoNidRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestGobVisNinGoGoRepPackedNative(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNinRepPackedNative,
		newMsg: newGoGoProtoNinRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestGobVimNinGoGoRepPackedNative(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNinRepPackedNative,
		newMsg: newGoGoProtoNinRepPackedNative,
	}.test(t)
}

//Optional Struct
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestGobVisNidGoGoOptStruct(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNidOptStruct,
		newMsg: newGoGoProtoNidOptStruct,
	}.test(t)
}

//Optional Struct
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestGobVimNidGoGoOptStruct(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNidOptStruct,
		newMsg: newGoGoProtoNidOptStruct,
	}.test(t)
}

//Optional Struct
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestGobVisNinGoGoOptStruct(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNinOptStruct,
		newMsg: newGoGoProtoNinOptStruct,
	}.test(t)
}

//Optional Struct
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestGobVimNinGoGoOptStruct(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNinOptStruct,
		newMsg: newGoGoProtoNinOptStruct,
	}.test(t)
}

//Repeated Struct
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestGobVisNidGoGoRepStruct(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNidRepStruct,
		newMsg: newGoGoProtoNidRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestGobVimNidGoGoRepStruct(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNidRepStruct,
		newMsg: newGoGoProtoNidRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestGobVisNinGoGoRepStruct(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNinRepStruct,
		newMsg: newGoGoProtoNinRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestGobVimNinGoGoRepStruct(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNinRepStruct,
		newMsg: newGoGoProtoNinRepStruct,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestGobVisNidGoGoEmbeddedStruct(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNidEmbeddedStruct,
		newMsg: newGoGoProtoNidEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestGobVimNidGoGoEmbeddedStruct(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNidEmbeddedStruct,
		newMsg: newGoGoProtoNidEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestGobVisNinGoGoEmbeddedStruct(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNinEmbeddedStruct,
		newMsg: newGoGoProtoNinEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestGobVimNinGoGoEmbeddedStruct(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNinEmbeddedStruct,
		newMsg: newGoGoProtoNinEmbeddedStruct,
	}.test(t)
}

//Optional Uuid
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestGobVisNidGoGoOptUuid(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNidOptUuid,
		newMsg: newGoGoProtoNidOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestGobVimNidGoGoOptUuid(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNidOptUuid,
		newMsg: newGoGoProtoNidOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestGobVisNinGoGoOptUuid(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNinOptUuid,
		newMsg: newGoGoProtoNinOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestGobVimNinGoGoOptUuid(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNinOptUuid,
		newMsg: newGoGoProtoNinOptUuid,
	}.test(t)
}

//Repeated Uuid
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestGobVisNidGoGoRepUuid(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNidRepUuid,
		newMsg: newGoGoProtoNidRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestGobVimNidGoGoRepUuid(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNidRepUuid,
		newMsg: newGoGoProtoNidRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestGobVisNinGoGoRepUuid(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNinRepUuid,
		newMsg: newGoGoProtoNinRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestGobVimNinGoGoRepUuid(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNinRepUuid,
		newMsg: newGoGoProtoNinRepUuid,
	}.test(t)
}

//Nested Struct
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestGobVisNidGoGoNestedStruct(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNidNestedStruct,
		newMsg: newGoGoProtoNidNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestGobVimNidGoGoNestedStruct(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNidNestedStruct,
		newMsg: newGoGoProtoNidNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestGobVisNinGoGoNestedStruct(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNinNestedStruct,
		newMsg: newGoGoProtoNinNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestGobVimNinGoGoNestedStruct(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNinNestedStruct,
		newMsg: newGoGoProtoNinNestedStruct,
	}.test(t)
}

//Optional Enum
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestGobVisNidGoGoOptEnum(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNidOptEnum,
		newMsg: newGoGoProtoNidOptEnum,
	}.test(t)
}

//Optional Enum
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestGobVimNidGoGoOptEnum(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNidOptEnum,
		newMsg: newGoGoProtoNidOptEnum,
	}.test(t)
}

//Optional Enum
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestGobVisNinGoGoOptEnum(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNinOptEnum,
		newMsg: newGoGoProtoNinOptEnum,
	}.test(t)
}

//Optional Enum
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestGobVimNinGoGoOptEnum(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNinOptEnum,
		newMsg: newGoGoProtoNinOptEnum,
	}.test(t)
}

//Repeated Enum
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestGobVisNidGoGoRepEnum(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNidRepEnum,
		newMsg: newGoGoProtoNidRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestGobVimNidGoGoRepEnum(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNidRepEnum,
		newMsg: newGoGoProtoNidRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestGobVisNinGoGoRepEnum(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNinRepEnum,
		newMsg: newGoGoProtoNinRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestGobVimNinGoGoRepEnum(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNinRepEnum,
		newMsg: newGoGoProtoNinRepEnum,
	}.test(t)
}

//Optional Uint128
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestGobVisNidGoGoOptUint128(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNidOptUint128,
		newMsg: newGoGoProtoNidOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestGobVimNidGoGoOptUint128(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNidOptUint128,
		newMsg: newGoGoProtoNidOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestGobVisNinGoGoOptUint128(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNinOptUint128,
		newMsg: newGoGoProtoNinOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestGobVimNinGoGoOptUint128(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNinOptUint128,
		newMsg: newGoGoProtoNinOptUint128,
	}.test(t)
}

//Repeated Uint128
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestGobVisNidGoGoRepUint128(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNidRepUint128,
		newMsg: newGoGoProtoNidRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestGobVimNidGoGoRepUint128(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNidRepUint128,
		newMsg: newGoGoProtoNidRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestGobVisNinGoGoRepUint128(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNinRepUint128,
		newMsg: newGoGoProtoNinRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestGobVimNinGoGoRepUint128(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNinRepUint128,
		newMsg: newGoGoProtoNinRepUint128,
	}.test(t)
}

//Optional Uint16
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestGobVisNidGoGoOptUint16(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNidOptUint16,
		newMsg: newGoGoProtoNidOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestGobVimNidGoGoOptUint16(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNidOptUint16,
		newMsg: newGoGoProtoNidOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestGobVisNinGoGoOptUint16(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNinOptUint16,
		newMsg: newGoGoProtoNinOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestGobVimNinGoGoOptUint16(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNinOptUint16,
		newMsg: newGoGoProtoNinOptUint16,
	}.test(t)
}

//Repeated Uint16
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestGobVisNidGoGoRepUint16(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNidRepUint16,
		newMsg: newGoGoProtoNidRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestGobVimNidGoGoRepUint16(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNidRepUint16,
		newMsg: newGoGoProtoNidRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestGobVisNinGoGoRepUint16(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNinRepUint16,
		newMsg: newGoGoProtoNinRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestGobVimNinGoGoRepUint16(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNinRepUint16,
		newMsg: newGoGoProtoNinRepUint16,
	}.test(t)
}

//Optional EnumUint32
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestGobVisNidGoGoOptEnumUint32(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNidOptEnumUint32,
		newMsg: newGoGoProtoNidOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestGobVimNidGoGoOptEnumUint32(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNidOptEnumUint32,
		newMsg: newGoGoProtoNidOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestGobVisNinGoGoOptEnumUint32(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNinOptEnumUint32,
		newMsg: newGoGoProtoNinOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestGobVimNinGoGoOptEnumUint32(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNinOptEnumUint32,
		newMsg: newGoGoProtoNinOptEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestGobVisNidGoGoRepEnumUint32(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNidRepEnumUint32,
		newMsg: newGoGoProtoNidRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestGobVimNidGoGoRepEnumUint32(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNidRepEnumUint32,
		newMsg: newGoGoProtoNidRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestGobVisNinGoGoRepEnumUint32(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNinRepEnumUint32,
		newMsg: newGoGoProtoNinRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestGobVimNinGoGoRepEnumUint32(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNinRepEnumUint32,
		newMsg: newGoGoProtoNinRepEnumUint32,
	}.test(t)
}

//Optional EnumUint16
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestGobVisNidGoGoOptEnumUint16(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNidOptEnumUint16,
		newMsg: newGoGoProtoNidOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestGobVimNidGoGoOptEnumUint16(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNidOptEnumUint16,
		newMsg: newGoGoProtoNidOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestGobVisNinGoGoOptEnumUint16(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNinOptEnumUint16,
		newMsg: newGoGoProtoNinOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestGobVimNinGoGoOptEnumUint16(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNinOptEnumUint16,
		newMsg: newGoGoProtoNinOptEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestGobVisNidGoGoRepEnumUint16(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNidRepEnumUint16,
		newMsg: newGoGoProtoNidRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestGobVimNidGoGoRepEnumUint16(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNidRepEnumUint16,
		newMsg: newGoGoProtoNidRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestGobVisNinGoGoRepEnumUint16(t *testing.T) {
	testGobGoGo{
		msg:    gogovisNinRepEnumUint16,
		newMsg: newGoGoProtoNinRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestGobVimNinGoGoRepEnumUint16(t *testing.T) {
	testGobGoGo{
		msg:    gogovimNinRepEnumUint16,
		newMsg: newGoGoProtoNinRepEnumUint16,
	}.test(t)
}

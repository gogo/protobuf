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

package test

import (
	"testing"
)

func TestJsonGoSimple(t *testing.T) {
	testJsonGo{
		msg:    gosimplemsg,
		newMsg: newGoProtoSimpleMsg,
	}.test(t)
}

func TestJsonGoGoSimple(t *testing.T) {
	testJsonGoGo{
		msg:    gogosimplemsg,
		newMsg: newGoGoProtoSimpleMsg,
	}.test(t)
}

func TestJsonGoToGoGoSimple(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      gosimplemsg,
		newGogomsg: newGoGoProtoSimpleMsg,
		gogomsg:    gogosimplemsg,
	}.test(t)
}

func TestJsonGoGoToGoSimple(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogosimplemsg,
		newGomsg: newGoProtoSimpleMsg,
		gomsg:    gosimplemsg,
		vis:      true,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is false
//Go -> Go
func TestJsonVisNidGoOptNative(t *testing.T) {
	testJsonGo{
		msg:    govisNidOptNative,
		newMsg: newGoProtoNidOptNative,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestJsonVisNidGoGoOptNative(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNidOptNative,
		newMsg: newGoGoProtoNidOptNative,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is false
//Go -> GoGo
func TestJsonVisNidGoToGoGoOptNative(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNidOptNative,
		newGogomsg: newGoGoProtoNidOptNative,
		gogomsg:    gogovisNidOptNative,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is false
//GoGo -> Go
func TestJsonVisNidGoGoToGoOptNative(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNidOptNative,
		newGomsg: newGoProtoNidOptNative,
		gomsg:    govisNidOptNative,
		vis:      true,
	}.test(t)
}

//

//Optional Native
//Values are not available
//Nullable is false
//Go -> Go
func TestJsonVimNidGoOptNative(t *testing.T) {
	testJsonGo{
		msg:    govimNidOptNative,
		newMsg: newGoProtoNidOptNative,
	}.test(t)
}

//Optional Native
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestJsonVimNidGoGoOptNative(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNidOptNative,
		newMsg: newGoGoProtoNidOptNative,
	}.test(t)
}

//Optional Native
//Values are not available
//Nullable is false
//Go -> GoGo
func TestJsonVimNidGoToGoGoOptNative(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNidOptNative,
		newGogomsg: newGoGoProtoNidOptNative,
		gogomsg:    gogovimNidOptNative,
	}.test(t)
}

//Optional Native
//Values are not available
//Nullable is false
//GoGo -> Go
func TestJsonVimNidGoGoToGoOptNative(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNidOptNative,
		newGomsg: newGoProtoNidOptNative,
		gomsg:    gozeroNidOptNative,
		vis:      false,
	}.test(t)
}

//

//Optional Native
//Values are available
//Nullable is true
//Go -> Go
func TestJsonVisNinGoOptNative(t *testing.T) {
	testJsonGo{
		msg:    govisNinOptNative,
		newMsg: newGoProtoNinOptNative,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestJsonVisNinGoGoOptNative(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNinOptNative,
		newMsg: newGoGoProtoNinOptNative,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is true
//Go -> GoGo
func TestJsonVisNinGoToGoGoOptNative(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNinOptNative,
		newGogomsg: newGoGoProtoNinOptNative,
		gogomsg:    gogovisNinOptNative,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is true
//GoGo -> Go
func TestJsonVisNinGoGoToGoOptNative(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNinOptNative,
		newGomsg: newGoProtoNinOptNative,
		gomsg:    govisNinOptNative,
		vis:      true,
	}.test(t)
}

//

//Optional Native
//Values are not available
//Nullable is true
//Go -> Go
func TestJsonVimNinGoOptNative(t *testing.T) {
	testJsonGo{
		msg:    govimNinOptNative,
		newMsg: newGoProtoNinOptNative,
	}.test(t)
}

//Optional Native
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestJsonVimNinGoGoOptNative(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNinOptNative,
		newMsg: newGoGoProtoNinOptNative,
	}.test(t)
}

//Optional Native
//Values are not available
//Nullable is true
//Go -> GoGo
func TestJsonVimNinGoToGoGoOptNative(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNinOptNative,
		newGogomsg: newGoGoProtoNinOptNative,
		gogomsg:    gogovimNinOptNative,
	}.test(t)
}

//Optional Native
//Values are not available
//Nullable is true
//GoGo -> Go
func TestJsonVimNinGoGoToGoOptNative(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNinOptNative,
		newGomsg: newGoProtoNinOptNative,
		gomsg:    govimNinOptNative,
		vis:      false,
	}.test(t)
}

//

//Repeated Native
//Values are available
//Nullable is false
//Go -> Go
func TestJsonVisNidGoRepNative(t *testing.T) {
	testJsonGo{
		msg:    govisNidRepNative,
		newMsg: newGoProtoNidRepNative,
	}.test(t)
}

//Repeated Native
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestJsonVisNidGoGoRepNative(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNidRepNative,
		newMsg: newGoGoProtoNidRepNative,
	}.test(t)
}

//Repeated Native
//Values are available
//Nullable is false
//Go -> GoGo
func TestJsonVisNidGoToGoGoRepNative(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNidRepNative,
		newGogomsg: newGoGoProtoNidRepNative,
		gogomsg:    gogovisNidRepNative,
	}.test(t)
}

//Repeated Native
//Values are available
//Nullable is false
//GoGo -> Go
func TestJsonVisNidGoGoToGoRepNative(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNidRepNative,
		newGomsg: newGoProtoNidRepNative,
		gomsg:    govisNidRepNative,
		vis:      true,
	}.test(t)
}

//

//Repeated Native
//Values are not available
//Nullable is false
//Go -> Go
func TestJsonVimNidGoRepNative(t *testing.T) {
	testJsonGo{
		msg:    govimNidRepNative,
		newMsg: newGoProtoNidRepNative,
	}.test(t)
}

//Repeated Native
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestJsonVimNidGoGoRepNative(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNidRepNative,
		newMsg: newGoGoProtoNidRepNative,
	}.test(t)
}

//Repeated Native
//Values are not available
//Nullable is false
//Go -> GoGo
func TestJsonVimNidGoToGoGoRepNative(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNidRepNative,
		newGogomsg: newGoGoProtoNidRepNative,
		gogomsg:    gogovimNidRepNative,
	}.test(t)
}

//Repeated Native
//Values are not available
//Nullable is false
//GoGo -> Go
func TestJsonVimNidGoGoToGoRepNative(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNidRepNative,
		newGomsg: newGoProtoNidRepNative,
		gomsg:    govimNidRepNative,
		vis:      false,
	}.test(t)
}

//

//Repeated Native
//Values are available
//Nullable is true
//Go -> Go
func TestJsonVisNinGoRepNative(t *testing.T) {
	testJsonGo{
		msg:    govisNinRepNative,
		newMsg: newGoProtoNinRepNative,
	}.test(t)
}

//Repeated Native
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestJsonVisNinGoGoRepNative(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNinRepNative,
		newMsg: newGoGoProtoNinRepNative,
	}.test(t)
}

//Repeated Native
//Values are available
//Nullable is true
//Go -> GoGo
func TestJsonVisNinGoToGoGoRepNative(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNinRepNative,
		newGogomsg: newGoGoProtoNinRepNative,
		gogomsg:    gogovisNinRepNative,
	}.test(t)
}

//Repeated Native
//Values are available
//Nullable is true
//GoGo -> Go
func TestJsonVisNinGoGoToGoRepNative(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNinRepNative,
		newGomsg: newGoProtoNinRepNative,
		gomsg:    govisNinRepNative,
		vis:      true,
	}.test(t)
}

//

//Repeated Native
//Values are not available
//Nullable is true
//Go -> Go
func TestJsonVimNinGoRepNative(t *testing.T) {
	testJsonGo{
		msg:    govimNinRepNative,
		newMsg: newGoProtoNinRepNative,
	}.test(t)
}

//Repeated Native
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestJsonVimNinGoGoRepNative(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNinRepNative,
		newMsg: newGoGoProtoNinRepNative,
	}.test(t)
}

//Repeated Native
//Values are not available
//Nullable is true
//Go -> GoGo
func TestJsonVimNinGoToGoGoRepNative(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNinRepNative,
		newGogomsg: newGoGoProtoNinRepNative,
		gogomsg:    gogovimNinRepNative,
	}.test(t)
}

//Repeated Native
//Values are not available
//Nullable is true
//GoGo -> Go
func TestJsonVimNinGoGoToGoRepNative(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNinRepNative,
		newGomsg: newGoProtoNinRepNative,
		gomsg:    govimNinRepNative,
		vis:      false,
	}.test(t)
}

//

//Packed Repeated Native
//Values are available
//Nullable is false
//Go -> Go
func TestJsonVisNidGoRepPackedNative(t *testing.T) {
	testJsonGo{
		msg:    govisNidRepPackedNative,
		newMsg: newGoProtoNidRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestJsonVisNidGoGoRepPackedNative(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNidRepPackedNative,
		newMsg: newGoGoProtoNidRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are available
//Nullable is false
//Go -> GoGo
func TestJsonVisNidGoToGoGoRepPackedNative(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNidRepPackedNative,
		newGogomsg: newGoGoProtoNidRepPackedNative,
		gogomsg:    gogovisNidRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are available
//Nullable is false
//GoGo -> Go
func TestJsonVisNidGoGoToGoRepPackedNative(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNidRepPackedNative,
		newGomsg: newGoProtoNidRepPackedNative,
		gomsg:    govisNidRepPackedNative,
		vis:      true,
	}.test(t)
}

//

//Packed Repeated Native
//Values are not available
//Nullable is false
//Go -> Go
func TestJsonVimNidGoRepPackedNative(t *testing.T) {
	testJsonGo{
		msg:    govimNidRepPackedNative,
		newMsg: newGoProtoNidRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestJsonVimNidGoGoRepPackedNative(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNidRepPackedNative,
		newMsg: newGoGoProtoNidRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are not available
//Nullable is false
//Go -> GoGo
func TestJsonVimNidGoToGoGoRepPackedNative(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNidRepPackedNative,
		newGogomsg: newGoGoProtoNidRepPackedNative,
		gogomsg:    gogovimNidRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are not available
//Nullable is false
//GoGo -> Go
func TestJsonVimNidGoGoToGoRepPackedNative(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNidRepPackedNative,
		newGomsg: newGoProtoNidRepPackedNative,
		gomsg:    govimNidRepPackedNative,
		vis:      false,
	}.test(t)
}

//

//Packed Repeated Native
//Values are available
//Nullable is true
//Go -> Go
func TestJsonVisNinGoRepPackedNative(t *testing.T) {
	testJsonGo{
		msg:    govisNinRepPackedNative,
		newMsg: newGoProtoNinRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestJsonVisNinGoGoRepPackedNative(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNinRepPackedNative,
		newMsg: newGoGoProtoNinRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are available
//Nullable is true
//Go -> GoGo
func TestJsonVisNinGoToGoGoRepPackedNative(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNinRepPackedNative,
		newGogomsg: newGoGoProtoNinRepPackedNative,
		gogomsg:    gogovisNinRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are available
//Nullable is true
//GoGo -> Go
func TestJsonVisNinGoGoToGoRepPackedNative(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNinRepPackedNative,
		newGomsg: newGoProtoNinRepPackedNative,
		gomsg:    govisNinRepPackedNative,
		vis:      true,
	}.test(t)
}

//

//Packed Repeated Native
//Values are not available
//Nullable is true
//Go -> Go
func TestJsonVimNinGoRepPackedNative(t *testing.T) {
	testJsonGo{
		msg:    govimNinRepPackedNative,
		newMsg: newGoProtoNinRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestJsonVimNinGoGoRepPackedNative(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNinRepPackedNative,
		newMsg: newGoGoProtoNinRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are not available
//Nullable is true
//Go -> GoGo
func TestJsonVimNinGoToGoGoRepPackedNative(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNinRepPackedNative,
		newGogomsg: newGoGoProtoNinRepPackedNative,
		gogomsg:    gogovimNinRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are not available
//Nullable is true
//GoGo -> Go
func TestJsonVimNinGoGoToGoRepPackedNative(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNinRepPackedNative,
		newGomsg: newGoProtoNinRepPackedNative,
		gomsg:    govimNinRepPackedNative,
		vis:      false,
	}.test(t)
}

//

//Optional Struct
//Values are available
//Nullable is false
//Go -> Go
func TestJsonVisNidGoOptStruct(t *testing.T) {
	testJsonGo{
		msg:    govisNidOptStruct,
		newMsg: newGoProtoNidOptStruct,
	}.test(t)
}

//Optional Struct
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestJsonVisNidGoGoOptStruct(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNidOptStruct,
		newMsg: newGoGoProtoNidOptStruct,
	}.test(t)
}

//Optional Struct
//Values are available
//Nullable is false
//Go -> GoGo
func TestJsonVisNidGoToGoGoOptStruct(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNidOptStruct,
		newGogomsg: newGoGoProtoNidOptStruct,
		gogomsg:    gogovisNidOptStruct,
	}.test(t)
}

//Optional Struct
//Values are available
//Nullable is false
//GoGo -> Go
func TestJsonVisNidGoGoToGoOptStruct(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNidOptStruct,
		newGomsg: newGoProtoNidOptStruct,
		gomsg:    govisNidOptStruct,
		vis:      true,
	}.test(t)
}

//

//Optional Struct
//Values are not available
//Nullable is false
//Go -> Go
func TestJsonVimNidGoOptStruct(t *testing.T) {
	testJsonGo{
		msg:    govimNidOptStruct,
		newMsg: newGoProtoNidOptStruct,
	}.test(t)
}

//Optional Struct
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestJsonVimNidGoGoOptStruct(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNidOptStruct,
		newMsg: newGoGoProtoNidOptStruct,
	}.test(t)
}

//Optional Struct
//Values are not available
//Nullable is false
//Go -> GoGo
func TestJsonVimNidGoToGoGoOptStruct(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      gozeroNidOptStruct,
		newGogomsg: newGoGoProtoNidOptStruct,
		gogomsg:    gogovimNidOptStruct,
	}.test(t)
}

//Optional Struct
//Values are not available
//Nullable is false
//GoGo -> Go
func TestJsonVimNidGoGoToGoOptStruct(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNidOptStruct,
		newGomsg: newGoProtoNidOptStruct,
		gomsg:    gozeroNidOptStruct,
		vis:      false,
	}.test(t)
}

//

//Optional Struct
//Values are available
//Nullable is true
//Go -> Go
func TestJsonVisNinGoOptStruct(t *testing.T) {
	testJsonGo{
		msg:    govisNinOptStruct,
		newMsg: newGoProtoNinOptStruct,
	}.test(t)
}

//Optional Struct
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestJsonVisNinGoGoOptStruct(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNinOptStruct,
		newMsg: newGoGoProtoNinOptStruct,
	}.test(t)
}

//Optional Struct
//Values are available
//Nullable is true
//Go -> GoGo
func TestJsonVisNinGoToGoGoOptStruct(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNinOptStruct,
		newGogomsg: newGoGoProtoNinOptStruct,
		gogomsg:    gogovisNinOptStruct,
	}.test(t)
}

//Optional Struct
//Values are available
//Nullable is true
//GoGo -> Go
func TestJsonVisNinGoGoToGoOptStruct(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNinOptStruct,
		newGomsg: newGoProtoNinOptStruct,
		gomsg:    govisNinOptStruct,
		vis:      true,
	}.test(t)
}

//

//Optional Struct
//Values are not available
//Nullable is true
//Go -> Go
func TestJsonVimNinGoOptStruct(t *testing.T) {
	testJsonGo{
		msg:    govimNinOptStruct,
		newMsg: newGoProtoNinOptStruct,
	}.test(t)
}

//Optional Struct
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestJsonVimNinGoGoOptStruct(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNinOptStruct,
		newMsg: newGoGoProtoNinOptStruct,
	}.test(t)
}

//Optional Struct
//Values are not available
//Nullable is true
//Go -> GoGo
func TestJsonVimNinGoToGoGoOptStruct(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNinOptStruct,
		newGogomsg: newGoGoProtoNinOptStruct,
		gogomsg:    gogovimNinOptStruct,
	}.test(t)
}

//Optional Struct
//Values are not available
//Nullable is true
//GoGo -> Go
func TestJsonVimNinGoGoToGoOptStruct(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNinOptStruct,
		newGomsg: newGoProtoNinOptStruct,
		gomsg:    govimNinOptStruct,
		vis:      false,
	}.test(t)
}

//

//Repeated Struct
//Values are available
//Nullable is false
//Go -> Go
func TestJsonVisNidGoRepStruct(t *testing.T) {
	testJsonGo{
		msg:    govisNidRepStruct,
		newMsg: newGoProtoNidRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestJsonVisNidGoGoRepStruct(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNidRepStruct,
		newMsg: newGoGoProtoNidRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are available
//Nullable is false
//Go -> GoGo
func TestJsonVisNidGoToGoGoRepStruct(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNidRepStruct,
		newGogomsg: newGoGoProtoNidRepStruct,
		gogomsg:    gogovisNidRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are available
//Nullable is false
//GoGo -> Go
func TestJsonVisNidGoGoToGoRepStruct(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNidRepStruct,
		newGomsg: newGoProtoNidRepStruct,
		gomsg:    govisNidRepStruct,
		vis:      true,
	}.test(t)
}

//

//Repeated Struct
//Values are not available
//Nullable is false
//Go -> Go
func TestJsonVimNidGoRepStruct(t *testing.T) {
	testJsonGo{
		msg:    govimNidRepStruct,
		newMsg: newGoProtoNidRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestJsonVimNidGoGoRepStruct(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNidRepStruct,
		newMsg: newGoGoProtoNidRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are not available
//Nullable is false
//Go -> GoGo
func TestJsonVimNidGoToGoGoRepStruct(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNidRepStruct,
		newGogomsg: newGoGoProtoNidRepStruct,
		gogomsg:    gogovimNidRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are not available
//Nullable is false
//GoGo -> Go
func TestJsonVimNidGoGoToGoRepStruct(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNidRepStruct,
		newGomsg: newGoProtoNidRepStruct,
		gomsg:    govimNidRepStruct,
		vis:      false,
	}.test(t)
}

//

//Repeated Struct
//Values are available
//Nullable is true
//Go -> Go
func TestJsonVisNinGoRepStruct(t *testing.T) {
	testJsonGo{
		msg:    govisNinRepStruct,
		newMsg: newGoProtoNinRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestJsonVisNinGoGoRepStruct(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNinRepStruct,
		newMsg: newGoGoProtoNinRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are available
//Nullable is true
//Go -> GoGo
func TestJsonVisNinGoToGoGoRepStruct(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNinRepStruct,
		newGogomsg: newGoGoProtoNinRepStruct,
		gogomsg:    gogovisNinRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are available
//Nullable is true
//GoGo -> Go
func TestJsonVisNinGoGoToGoRepStruct(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNinRepStruct,
		newGomsg: newGoProtoNinRepStruct,
		gomsg:    govisNinRepStruct,
		vis:      true,
	}.test(t)
}

//

//Repeated Struct
//Values are not available
//Nullable is true
//Go -> Go
func TestJsonVimNinGoRepStruct(t *testing.T) {
	testJsonGo{
		msg:    govimNinRepStruct,
		newMsg: newGoProtoNinRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestJsonVimNinGoGoRepStruct(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNinRepStruct,
		newMsg: newGoGoProtoNinRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are not available
//Nullable is true
//Go -> GoGo
func TestJsonVimNinGoToGoGoRepStruct(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNinRepStruct,
		newGogomsg: newGoGoProtoNinRepStruct,
		gogomsg:    gogovimNinRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are not available
//Nullable is true
//GoGo -> Go
func TestJsonVimNinGoGoToGoRepStruct(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNinRepStruct,
		newGomsg: newGoProtoNinRepStruct,
		gomsg:    govimNinRepStruct,
		vis:      false,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is false
//Go -> Go
func TestJsonVisNidGoEmbeddedStruct(t *testing.T) {
	testJsonGo{
		msg:    govisNidEmbeddedStruct,
		newMsg: newGoProtoNidEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestJsonVisNidGoGoEmbeddedStruct(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNidEmbeddedStruct,
		newMsg: newGoGoProtoNidEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is false
//Go -> GoGo
func TestJsonVisNidGoToGoGoEmbeddedStruct(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNidEmbeddedStruct,
		newGogomsg: newGoGoProtoNidEmbeddedStruct,
		gogomsg:    gogovisNidEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is false
//GoGo -> Go
func TestJsonVisNidGoGoToGoEmbeddedStruct(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNidEmbeddedStruct,
		newGomsg: newGoProtoNidEmbeddedStruct,
		gomsg:    govisNidEmbeddedStruct,
		vis:      true,
	}.test(t)
}

//

//Embedded Struct
//Values are not available
//Nullable is false
//Go -> Go
func TestJsonVimNidGoEmbeddedStruct(t *testing.T) {
	testJsonGo{
		msg:    govimNidEmbeddedStruct,
		newMsg: newGoProtoNidEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestJsonVimNidGoGoEmbeddedStruct(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNidEmbeddedStruct,
		newMsg: newGoGoProtoNidEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are not available
//Nullable is false
//Go -> GoGo
func TestJsonVimNidGoToGoGoEmbeddedStruct(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      gozeroNidEmbeddedStruct,
		newGogomsg: newGoGoProtoNidEmbeddedStruct,
		gogomsg:    gogovimNidEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are not available
//Nullable is false
//GoGo -> Go
func TestJsonVimNidGoGoToGoEmbeddedStruct(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNidEmbeddedStruct,
		newGomsg: newGoProtoNidEmbeddedStruct,
		gomsg:    gozeroNidEmbeddedStruct,
		vis:      false,
	}.test(t)
}

//

//Embedded Struct
//Values are available
//Nullable is true
//Go -> Go
func TestJsonVisNinGoEmbeddedStruct(t *testing.T) {
	testJsonGo{
		msg:    govisNinEmbeddedStruct,
		newMsg: newGoProtoNinEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestJsonVisNinGoGoEmbeddedStruct(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNinEmbeddedStruct,
		newMsg: newGoGoProtoNinEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is true
//Go -> GoGo
func TestJsonVisNinGoToGoGoEmbeddedStruct(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNinEmbeddedStruct,
		newGogomsg: newGoGoProtoNinEmbeddedStruct,
		gogomsg:    gogovisNinEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is true
//GoGo -> Go
func TestJsonVisNinGoGoToGoEmbeddedStruct(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNinEmbeddedStruct,
		newGomsg: newGoProtoNinEmbeddedStruct,
		gomsg:    govisNinEmbeddedStruct,
		vis:      true,
	}.test(t)
}

//

//Embedded Struct
//Values are not available
//Nullable is true
//Go -> Go
func TestJsonVimNinGoEmbeddedStruct(t *testing.T) {
	testJsonGo{
		msg:    govimNinEmbeddedStruct,
		newMsg: newGoProtoNinEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestJsonVimNinGoGoEmbeddedStruct(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNinEmbeddedStruct,
		newMsg: newGoGoProtoNinEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are not available
//Nullable is true
//Go -> GoGo
func TestJsonVimNinGoToGoGoEmbeddedStruct(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNinEmbeddedStruct,
		newGogomsg: newGoGoProtoNinEmbeddedStruct,
		gogomsg:    gogovimNinEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are not available
//Nullable is true
//GoGo -> Go
func TestJsonVimNinGoGoToGoEmbeddedStruct(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNinEmbeddedStruct,
		newGomsg: newGoProtoNinEmbeddedStruct,
		gomsg:    govimNinEmbeddedStruct,
		vis:      false,
	}.test(t)
}

//

//Nested Struct
//Values are available
//Nullable is false
//Go -> Go
func TestJsonVisNidGoNestedStruct(t *testing.T) {
	testJsonGo{
		msg:    govisNidNestedStruct,
		newMsg: newGoProtoNidNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestJsonVisNidGoGoNestedStruct(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNidNestedStruct,
		newMsg: newGoGoProtoNidNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are available
//Nullable is false
//Go -> GoGo
func TestJsonVisNidGoToGoGoNestedStruct(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNidNestedStruct,
		newGogomsg: newGoGoProtoNidNestedStruct,
		gogomsg:    gogovisNidNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are available
//Nullable is false
//GoGo -> Go
func TestJsonVisNidGoGoToGoNestedStruct(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNidNestedStruct,
		newGomsg: newGoProtoNidNestedStruct,
		gomsg:    govisNidNestedStruct,
		vis:      true,
	}.test(t)
}

//

//Nested Struct
//Values are not available
//Nullable is false
//Go -> Go
func TestJsonVimNidGoNestedStruct(t *testing.T) {
	testJsonGo{
		msg:    govimNidNestedStruct,
		newMsg: newGoProtoNidNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestJsonVimNidGoGoNestedStruct(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNidNestedStruct,
		newMsg: newGoGoProtoNidNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are not available
//Nullable is false
//Go -> GoGo
func TestJsonVimNidGoToGoGoNestedStruct(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      gozeroNidNestedStruct,
		newGogomsg: newGoGoProtoNidNestedStruct,
		gogomsg:    gogovimNidNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are not available
//Nullable is false
//GoGo -> Go
func TestJsonVimNidGoGoToGoNestedStruct(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNidNestedStruct,
		newGomsg: newGoProtoNidNestedStruct,
		gomsg:    gozeroNidNestedStruct,
		vis:      false,
	}.test(t)
}

//

//Nested Struct
//Values are available
//Nullable is true
//Go -> Go
func TestJsonVisNinGoNestedStruct(t *testing.T) {
	testJsonGo{
		msg:    govisNinNestedStruct,
		newMsg: newGoProtoNinNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestJsonVisNinGoGoNestedStruct(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNinNestedStruct,
		newMsg: newGoGoProtoNinNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are available
//Nullable is true
//Go -> GoGo
func TestJsonVisNinGoToGoGoNestedStruct(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNinNestedStruct,
		newGogomsg: newGoGoProtoNinNestedStruct,
		gogomsg:    gogovisNinNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are available
//Nullable is true
//GoGo -> Go
func TestJsonVisNinGoGoToGoNestedStruct(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNinNestedStruct,
		newGomsg: newGoProtoNinNestedStruct,
		gomsg:    govisNinNestedStruct,
		vis:      true,
	}.test(t)
}

//

//Nested Struct
//Values are not available
//Nullable is true
//Go -> Go
func TestJsonVimNinGoNestedStruct(t *testing.T) {
	testJsonGo{
		msg:    govimNinNestedStruct,
		newMsg: newGoProtoNinNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestJsonVimNinGoGoNestedStruct(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNinNestedStruct,
		newMsg: newGoGoProtoNinNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are not available
//Nullable is true
//Go -> GoGo
func TestJsonVimNinGoToGoGoNestedStruct(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNinNestedStruct,
		newGogomsg: newGoGoProtoNinNestedStruct,
		gogomsg:    gogovimNinNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are not available
//Nullable is true
//GoGo -> Go
func TestJsonVimNinGoGoToGoNestedStruct(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNinNestedStruct,
		newGomsg: newGoProtoNinNestedStruct,
		gomsg:    govimNinNestedStruct,
		vis:      false,
	}.test(t)
}

//

//Optional Enum
//Values are available
//Nullable is false
//Go -> Go
func TestJsonVisNidGoOptEnum(t *testing.T) {
	testJsonGo{
		msg:    govisNidOptEnum,
		newMsg: newGoProtoNidOptEnum,
	}.test(t)
}

//Optional Enum
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestJsonVisNidGoGoOptEnum(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNidOptEnum,
		newMsg: newGoGoProtoNidOptEnum,
	}.test(t)
}

//Optional Enum
//Values are available
//Nullable is false
//Go -> GoGo
func TestJsonVisNidGoToGoGoOptEnum(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNidOptEnum,
		newGogomsg: newGoGoProtoNidOptEnum,
		gogomsg:    gogovisNidOptEnum,
	}.test(t)
}

//Optional Enum
//Values are available
//Nullable is false
//GoGo -> Go
func TestJsonVisNidGoGoToGoOptEnum(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNidOptEnum,
		newGomsg: newGoProtoNidOptEnum,
		gomsg:    govisNidOptEnum,
		vis:      true,
	}.test(t)
}

//

//Optional Enum
//Values are not available
//Nullable is false
//Go -> Go
func TestJsonVimNidGoOptEnum(t *testing.T) {
	testJsonGo{
		msg:    govimNidOptEnum,
		newMsg: newGoProtoNidOptEnum,
	}.test(t)
}

//Optional Enum
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestJsonVimNidGoGoOptEnum(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNidOptEnum,
		newMsg: newGoGoProtoNidOptEnum,
	}.test(t)
}

//Optional Enum
//Values are not available
//Nullable is false
//Go -> GoGo
func TestJsonVimNidGoToGoGoOptEnum(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      gozeroNidOptEnum,
		newGogomsg: newGoGoProtoNidOptEnum,
		gogomsg:    gogovimNidOptEnum,
	}.test(t)
}

//Optional Enum
//Values are not available
//Nullable is false
//GoGo -> Go
func TestJsonVimNidGoGoToGoOptEnum(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNidOptEnum,
		newGomsg: newGoProtoNidOptEnum,
		gomsg:    gozeroNidOptEnum,
		vis:      false,
	}.test(t)
}

//

//Optional Enum
//Values are available
//Nullable is true
//Go -> Go
func TestJsonVisNinGoOptEnum(t *testing.T) {
	testJsonGo{
		msg:    govisNinOptEnum,
		newMsg: newGoProtoNinOptEnum,
	}.test(t)
}

//Optional Enum
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestJsonVisNinGoGoOptEnum(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNinOptEnum,
		newMsg: newGoGoProtoNinOptEnum,
	}.test(t)
}

//Optional Enum
//Values are available
//Nullable is true
//Go -> GoGo
func TestJsonVisNinGoToGoGoOptEnum(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNinOptEnum,
		newGogomsg: newGoGoProtoNinOptEnum,
		gogomsg:    gogovisNinOptEnum,
	}.test(t)
}

//Optional Enum
//Values are available
//Nullable is true
//GoGo -> Go
func TestJsonVisNinGoGoToGoOptEnum(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNinOptEnum,
		newGomsg: newGoProtoNinOptEnum,
		gomsg:    govisNinOptEnum,
		vis:      true,
	}.test(t)
}

//

//Optional Enum
//Values are not available
//Nullable is true
//Go -> Go
func TestJsonVimNinGoOptEnum(t *testing.T) {
	testJsonGo{
		msg:    govimNinOptEnum,
		newMsg: newGoProtoNinOptEnum,
	}.test(t)
}

//Optional Enum
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestJsonVimNinGoGoOptEnum(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNinOptEnum,
		newMsg: newGoGoProtoNinOptEnum,
	}.test(t)
}

//Optional Enum
//Values are not available
//Nullable is true
//Go -> GoGo
func TestJsonVimNinGoToGoGoOptEnum(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNinOptEnum,
		newGogomsg: newGoGoProtoNinOptEnum,
		gogomsg:    gogovimNinOptEnum,
	}.test(t)
}

//Optional Enum
//Values are not available
//Nullable is true
//GoGo -> Go
func TestJsonVimNinGoGoToGoOptEnum(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNinOptEnum,
		newGomsg: newGoProtoNinOptEnum,
		gomsg:    govimNinOptEnum,
		vis:      false,
	}.test(t)
}

//

//Repeated Enum
//Values are available
//Nullable is false
//Go -> Go
func TestJsonVisNidGoRepEnum(t *testing.T) {
	testJsonGo{
		msg:    govisNidRepEnum,
		newMsg: newGoProtoNidRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestJsonVisNidGoGoRepEnum(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNidRepEnum,
		newMsg: newGoGoProtoNidRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are available
//Nullable is false
//Go -> GoGo
func TestJsonVisNidGoToGoGoRepEnum(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNidRepEnum,
		newGogomsg: newGoGoProtoNidRepEnum,
		gogomsg:    gogovisNidRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are available
//Nullable is false
//GoGo -> Go
func TestJsonVisNidGoGoToGoRepEnum(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNidRepEnum,
		newGomsg: newGoProtoNidRepEnum,
		gomsg:    govisNidRepEnum,
		vis:      true,
	}.test(t)
}

//

//Repeated Enum
//Values are not available
//Nullable is false
//Go -> Go
func TestJsonVimNidGoRepEnum(t *testing.T) {
	testJsonGo{
		msg:    govimNidRepEnum,
		newMsg: newGoProtoNidRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestJsonVimNidGoGoRepEnum(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNidRepEnum,
		newMsg: newGoGoProtoNidRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are not available
//Nullable is false
//Go -> GoGo
func TestJsonVimNidGoToGoGoRepEnum(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNidRepEnum,
		newGogomsg: newGoGoProtoNidRepEnum,
		gogomsg:    gogovimNidRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are not available
//Nullable is false
//GoGo -> Go
func TestJsonVimNidGoGoToGoRepEnum(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNidRepEnum,
		newGomsg: newGoProtoNidRepEnum,
		gomsg:    govimNidRepEnum,
		vis:      false,
	}.test(t)
}

//

//Repeated Enum
//Values are available
//Nullable is true
//Go -> Go
func TestJsonVisNinGoRepEnum(t *testing.T) {
	testJsonGo{
		msg:    govisNinRepEnum,
		newMsg: newGoProtoNinRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestJsonVisNinGoGoRepEnum(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNinRepEnum,
		newMsg: newGoGoProtoNinRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are available
//Nullable is true
//Go -> GoGo
func TestJsonVisNinGoToGoGoRepEnum(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNinRepEnum,
		newGogomsg: newGoGoProtoNinRepEnum,
		gogomsg:    gogovisNinRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are available
//Nullable is true
//GoGo -> Go
func TestJsonVisNinGoGoToGoRepEnum(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNinRepEnum,
		newGomsg: newGoProtoNinRepEnum,
		gomsg:    govisNinRepEnum,
		vis:      true,
	}.test(t)
}

//

//Repeated Enum
//Values are not available
//Nullable is true
//Go -> Go
func TestJsonVimNinGoRepEnum(t *testing.T) {
	testJsonGo{
		msg:    govimNinRepEnum,
		newMsg: newGoProtoNinRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestJsonVimNinGoGoRepEnum(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNinRepEnum,
		newMsg: newGoGoProtoNinRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are not available
//Nullable is true
//Go -> GoGo
func TestJsonVimNinGoToGoGoRepEnum(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNinRepEnum,
		newGogomsg: newGoGoProtoNinRepEnum,
		gogomsg:    gogovimNinRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are not available
//Nullable is true
//GoGo -> Go
func TestJsonVimNinGoGoToGoRepEnum(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNinRepEnum,
		newGomsg: newGoProtoNinRepEnum,
		gomsg:    govimNinRepEnum,
		vis:      false,
	}.test(t)
}

//

//Optional UuidAsBytes
//Values are available
//Nullable is false
//Go -> Go
func TestJsonVisNidGoOptUuidAsBytes(t *testing.T) {
	testJsonGo{
		msg:    govisNidOptUuidAsBytes,
		newMsg: newGoProtoNidOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestJsonVisNidGoGoOptUuidAsBytes(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNidOptUuidAsBytes,
		newMsg: newGoGoProtoNidOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are available
//Nullable is false
//Go -> GoGo
func TestJsonVisNidGoToGoGoOptUuidAsBytes(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNidOptUuidAsBytes,
		newGogomsg: newGoGoProtoNidOptUuidAsBytes,
		gogomsg:    gogovisNidOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are available
//Nullable is false
//GoGo -> Go
func TestJsonVisNidGoGoToGoOptUuidAsBytes(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNidOptUuidAsBytes,
		newGomsg: newGoProtoNidOptUuidAsBytes,
		gomsg:    govisNidOptUuidAsBytes,
		vis:      true,
	}.test(t)
}

//

//Optional UuidAsBytes
//Values are not available
//Nullable is false
//Go -> Go
func TestJsonVimNidGoOptUuidAsBytes(t *testing.T) {
	testJsonGo{
		msg:    govimNidOptUuidAsBytes,
		newMsg: newGoProtoNidOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestJsonVimNidGoGoOptUuidAsBytes(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNidOptUuidAsBytes,
		newMsg: newGoGoProtoNidOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are not available
//Nullable is false
//Go -> GoGo
func TestJsonVimNidGoToGoGoOptUuidAsBytes(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNidOptUuidAsBytes,
		newGogomsg: newGoGoProtoNidOptUuidAsBytes,
		gogomsg:    gogovimNidOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are not available
//Nullable is false
//GoGo -> Go
func TestJsonVimNidGoGoToGoOptUuidAsBytes(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNidOptUuidAsBytes,
		newGomsg: newGoProtoNidOptUuidAsBytes,
		gomsg:    govimNidOptUuidAsBytes,
		vis:      false,
	}.test(t)
}

//

//Optional UuidAsBytes
//Values are available
//Nullable is true
//Go -> Go
func TestJsonVisNinGoOptUuidAsBytes(t *testing.T) {
	testJsonGo{
		msg:    govisNinOptUuidAsBytes,
		newMsg: newGoProtoNinOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestJsonVisNinGoGoOptUuidAsBytes(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNinOptUuidAsBytes,
		newMsg: newGoGoProtoNinOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are available
//Nullable is true
//Go -> GoGo
func TestJsonVisNinGoToGoGoOptUuidAsBytes(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNinOptUuidAsBytes,
		newGogomsg: newGoGoProtoNinOptUuidAsBytes,
		gogomsg:    gogovisNinOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are available
//Nullable is true
//GoGo -> Go
func TestJsonVisNinGoGoToGoOptUuidAsBytes(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNinOptUuidAsBytes,
		newGomsg: newGoProtoNinOptUuidAsBytes,
		gomsg:    govisNinOptUuidAsBytes,
		vis:      true,
	}.test(t)
}

//

//Optional UuidAsBytes
//Values are not available
//Nullable is true
//Go -> Go
func TestJsonVimNinGoOptUuidAsBytes(t *testing.T) {
	testJsonGo{
		msg:    govimNinOptUuidAsBytes,
		newMsg: newGoProtoNinOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestJsonVimNinGoGoOptUuidAsBytes(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNinOptUuidAsBytes,
		newMsg: newGoGoProtoNinOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are not available
//Nullable is true
//Go -> GoGo
func TestJsonVimNinGoToGoGoOptUuidAsBytes(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNinOptUuidAsBytes,
		newGogomsg: newGoGoProtoNinOptUuidAsBytes,
		gogomsg:    gogovimNinOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are not available
//Nullable is true
//GoGo -> Go
func TestJsonVimNinGoGoToGoOptUuidAsBytes(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNinOptUuidAsBytes,
		newGomsg: newGoProtoNinOptUuidAsBytes,
		gomsg:    govimNinOptUuidAsBytes,
		vis:      false,
	}.test(t)
}

//

//Repeated UuidAsBytes
//Values are available
//Nullable is false
//Go -> Go
func TestJsonVisNidGoRepUuidAsBytes(t *testing.T) {
	testJsonGo{
		msg:    govisNidRepUuidAsBytes,
		newMsg: newGoProtoNidRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestJsonVisNidGoGoRepUuidAsBytes(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNidRepUuidAsBytes,
		newMsg: newGoGoProtoNidRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are available
//Nullable is false
//Go -> GoGo
func TestJsonVisNidGoToGoGoRepUuidAsBytes(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNidRepUuidAsBytes,
		newGogomsg: newGoGoProtoNidRepUuidAsBytes,
		gogomsg:    gogovisNidRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are available
//Nullable is false
//GoGo -> Go
func TestJsonVisNidGoGoToGoRepUuidAsBytes(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNidRepUuidAsBytes,
		newGomsg: newGoProtoNidRepUuidAsBytes,
		gomsg:    govisNidRepUuidAsBytes,
		vis:      true,
	}.test(t)
}

//

//Repeated UuidAsBytes
//Values are not available
//Nullable is false
//Go -> Go
func TestJsonVimNidGoRepUuidAsBytes(t *testing.T) {
	testJsonGo{
		msg:    govimNidRepUuidAsBytes,
		newMsg: newGoProtoNidRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestJsonVimNidGoGoRepUuidAsBytes(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNidRepUuidAsBytes,
		newMsg: newGoGoProtoNidRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are not available
//Nullable is false
//Go -> GoGo
func TestJsonVimNidGoToGoGoRepUuidAsBytes(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNidRepUuidAsBytes,
		newGogomsg: newGoGoProtoNidRepUuidAsBytes,
		gogomsg:    gogovimNidRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are not available
//Nullable is false
//GoGo -> Go
func TestJsonVimNidGoGoToGoRepUuidAsBytes(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNidRepUuidAsBytes,
		newGomsg: newGoProtoNidRepUuidAsBytes,
		gomsg:    govimNidRepUuidAsBytes,
		vis:      false,
	}.test(t)
}

//

//Repeated UuidAsBytes
//Values are available
//Nullable is true
//Go -> Go
func TestJsonVisNinGoRepUuidAsBytes(t *testing.T) {
	testJsonGo{
		msg:    govisNinRepUuidAsBytes,
		newMsg: newGoProtoNinRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestJsonVisNinGoGoRepUuidAsBytes(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNinRepUuidAsBytes,
		newMsg: newGoGoProtoNinRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are available
//Nullable is true
//Go -> GoGo
func TestJsonVisNinGoToGoGoRepUuidAsBytes(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNinRepUuidAsBytes,
		newGogomsg: newGoGoProtoNinRepUuidAsBytes,
		gogomsg:    gogovisNinRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are available
//Nullable is true
//GoGo -> Go
func TestJsonVisNinGoGoToGoRepUuidAsBytes(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNinRepUuidAsBytes,
		newGomsg: newGoProtoNinRepUuidAsBytes,
		gomsg:    govisNinRepUuidAsBytes,
		vis:      true,
	}.test(t)
}

//

//Repeated UuidAsBytes
//Values are not available
//Nullable is true
//Go -> Go
func TestJsonVimNinGoRepUuidAsBytes(t *testing.T) {
	testJsonGo{
		msg:    govimNinRepUuidAsBytes,
		newMsg: newGoProtoNinRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestJsonVimNinGoGoRepUuidAsBytes(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNinRepUuidAsBytes,
		newMsg: newGoGoProtoNinRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are not available
//Nullable is true
//Go -> GoGo
func TestJsonVimNinGoToGoGoRepUuidAsBytes(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNinRepUuidAsBytes,
		newGogomsg: newGoGoProtoNinRepUuidAsBytes,
		gogomsg:    gogovimNinRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are not available
//Nullable is true
//GoGo -> Go
func TestJsonVimNinGoGoToGoRepUuidAsBytes(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNinRepUuidAsBytes,
		newGomsg: newGoProtoNinRepUuidAsBytes,
		gomsg:    govimNinRepUuidAsBytes,
		vis:      false,
	}.test(t)
}

//

//Optional Uint128AsBytes
//Values are available
//Nullable is false
//Go -> Go
func TestJsonVisNidGoOptUint128AsBytes(t *testing.T) {
	testJsonGo{
		msg:    govisNidOptUint128AsBytes,
		newMsg: newGoProtoNidOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestJsonVisNidGoGoOptUint128AsBytes(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNidOptUint128AsBytes,
		newMsg: newGoGoProtoNidOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are available
//Nullable is false
//Go -> GoGo
func TestJsonVisNidGoToGoGoOptUint128AsBytes(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNidOptUint128AsBytes,
		newGogomsg: newGoGoProtoNidOptUint128AsBytes,
		gogomsg:    gogovisNidOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are available
//Nullable is false
//GoGo -> Go
func TestJsonVisNidGoGoToGoOptUint128AsBytes(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNidOptUint128AsBytes,
		newGomsg: newGoProtoNidOptUint128AsBytes,
		gomsg:    govisNidOptUint128AsBytes,
		vis:      true,
	}.test(t)
}

//

//Optional Uint128AsBytes
//Values are not available
//Nullable is false
//Go -> Go
func TestJsonVimNidGoOptUint128AsBytes(t *testing.T) {
	testJsonGo{
		msg:    govimNidOptUint128AsBytes,
		newMsg: newGoProtoNidOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestJsonVimNidGoGoOptUint128AsBytes(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNidOptUint128AsBytes,
		newMsg: newGoGoProtoNidOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are not available
//Nullable is false
//Go -> GoGo
func TestJsonVimNidGoToGoGoOptUint128AsBytes(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNidOptUint128AsBytes,
		newGogomsg: newGoGoProtoNidOptUint128AsBytes,
		gogomsg:    gogovimNidOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are not available
//Nullable is false
//GoGo -> Go
func TestJsonVimNidGoGoToGoOptUint128AsBytes(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNidOptUint128AsBytes,
		newGomsg: newGoProtoNidOptUint128AsBytes,
		gomsg:    gozeroOptUint128AsBytes,
		vis:      false,
	}.test(t)
}

//

//Optional Uint128AsBytes
//Values are available
//Nullable is true
//Go -> Go
func TestJsonVisNinGoOptUint128AsBytes(t *testing.T) {
	testJsonGo{
		msg:    govisNinOptUint128AsBytes,
		newMsg: newGoProtoNinOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestJsonVisNinGoGoOptUint128AsBytes(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNinOptUint128AsBytes,
		newMsg: newGoGoProtoNinOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are available
//Nullable is true
//Go -> GoGo
func TestJsonVisNinGoToGoGoOptUint128AsBytes(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNinOptUint128AsBytes,
		newGogomsg: newGoGoProtoNinOptUint128AsBytes,
		gogomsg:    gogovisNinOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are available
//Nullable is true
//GoGo -> Go
func TestJsonVisNinGoGoToGoOptUint128AsBytes(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNinOptUint128AsBytes,
		newGomsg: newGoProtoNinOptUint128AsBytes,
		gomsg:    govisNinOptUint128AsBytes,
		vis:      true,
	}.test(t)
}

//

//Optional Uint128AsBytes
//Values are not available
//Nullable is true
//Go -> Go
func TestJsonVimNinGoOptUint128AsBytes(t *testing.T) {
	testJsonGo{
		msg:    govimNinOptUint128AsBytes,
		newMsg: newGoProtoNinOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestJsonVimNinGoGoOptUint128AsBytes(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNinOptUint128AsBytes,
		newMsg: newGoGoProtoNinOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are not available
//Nullable is true
//Go -> GoGo
func TestJsonVimNinGoToGoGoOptUint128AsBytes(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNinOptUint128AsBytes,
		newGogomsg: newGoGoProtoNinOptUint128AsBytes,
		gogomsg:    gogovimNinOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are not available
//Nullable is true
//GoGo -> Go
func TestJsonVimNinGoGoToGoOptUint128AsBytes(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNinOptUint128AsBytes,
		newGomsg: newGoProtoNinOptUint128AsBytes,
		gomsg:    govimNinOptUint128AsBytes,
		vis:      false,
	}.test(t)
}

//

//Repeated Uint128AsBytes
//Values are available
//Nullable is false
//Go -> Go
func TestJsonVisNidGoRepUint128AsBytes(t *testing.T) {
	testJsonGo{
		msg:    govisNidRepUint128AsBytes,
		newMsg: newGoProtoNidRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestJsonVisNidGoGoRepUint128AsBytes(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNidRepUint128AsBytes,
		newMsg: newGoGoProtoNidRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are available
//Nullable is false
//Go -> GoGo
func TestJsonVisNidGoToGoGoRepUint128AsBytes(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNidRepUint128AsBytes,
		newGogomsg: newGoGoProtoNidRepUint128AsBytes,
		gogomsg:    gogovisNidRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are available
//Nullable is false
//GoGo -> Go
func TestJsonVisNidGoGoToGoRepUint128AsBytes(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNidRepUint128AsBytes,
		newGomsg: newGoProtoNidRepUint128AsBytes,
		gomsg:    govisNidRepUint128AsBytes,
		vis:      true,
	}.test(t)
}

//

//Repeated Uint128AsBytes
//Values are not available
//Nullable is false
//Go -> Go
func TestJsonVimNidGoRepUint128AsBytes(t *testing.T) {
	testJsonGo{
		msg:    govimNidRepUint128AsBytes,
		newMsg: newGoProtoNidRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestJsonVimNidGoGoRepUint128AsBytes(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNidRepUint128AsBytes,
		newMsg: newGoGoProtoNidRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are not available
//Nullable is false
//Go -> GoGo
func TestJsonVimNidGoToGoGoRepUint128AsBytes(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNidRepUint128AsBytes,
		newGogomsg: newGoGoProtoNidRepUint128AsBytes,
		gogomsg:    gogovimNidRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are not available
//Nullable is false
//GoGo -> Go
func TestJsonVimNidGoGoToGoRepUint128AsBytes(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNidRepUint128AsBytes,
		newGomsg: newGoProtoNidRepUint128AsBytes,
		gomsg:    govimNidRepUint128AsBytes,
		vis:      false,
	}.test(t)
}

//

//Repeated Uint128AsBytes
//Values are available
//Nullable is true
//Go -> Go
func TestJsonVisNinGoRepUint128AsBytes(t *testing.T) {
	testJsonGo{
		msg:    govisNinRepUint128AsBytes,
		newMsg: newGoProtoNinRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestJsonVisNinGoGoRepUint128AsBytes(t *testing.T) {
	testJsonGoGo{
		msg:    gogovisNinRepUint128AsBytes,
		newMsg: newGoGoProtoNinRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are available
//Nullable is true
//Go -> GoGo
func TestJsonVisNinGoToGoGoRepUint128AsBytes(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govisNinRepUint128AsBytes,
		newGogomsg: newGoGoProtoNinRepUint128AsBytes,
		gogomsg:    gogovisNinRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are available
//Nullable is true
//GoGo -> Go
func TestJsonVisNinGoGoToGoRepUint128AsBytes(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovisNinRepUint128AsBytes,
		newGomsg: newGoProtoNinRepUint128AsBytes,
		gomsg:    govisNinRepUint128AsBytes,
		vis:      true,
	}.test(t)
}

//

//Repeated Uint128AsBytes
//Values are not available
//Nullable is true
//Go -> Go
func TestJsonVimNinGoRepUint128AsBytes(t *testing.T) {
	testJsonGo{
		msg:    govimNinRepUint128AsBytes,
		newMsg: newGoProtoNinRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestJsonVimNinGoGoRepUint128AsBytes(t *testing.T) {
	testJsonGoGo{
		msg:    gogovimNinRepUint128AsBytes,
		newMsg: newGoGoProtoNinRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are not available
//Nullable is true
//Go -> GoGo
func TestJsonVimNinGoToGoGoRepUint128AsBytes(t *testing.T) {
	testJsonGoToGoGo{
		gomsg:      govimNinRepUint128AsBytes,
		newGogomsg: newGoGoProtoNinRepUint128AsBytes,
		gogomsg:    gogovimNinRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are not available
//Nullable is true
//GoGo -> Go
func TestJsonVimNinGoGoToGoRepUint128AsBytes(t *testing.T) {
	testJsonGoGoToGo{
		gogomsg:  gogovimNinRepUint128AsBytes,
		newGomsg: newGoProtoNinRepUint128AsBytes,
		gomsg:    govimNinRepUint128AsBytes,
		vis:      false,
	}.test(t)
}

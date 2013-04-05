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

func TestProtoTextGoSimple(t *testing.T) {
	testProtoTextGo{
		msg:    gosimplemsg,
		newMsg: newGoProtoSimpleMsg,
	}.test(t)
}

func TestProtoTextGoGoSimple(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogosimplemsg,
		newMsg: newGoGoProtoSimpleMsg,
	}.test(t)
}

func TestProtoTextGoToGoGoSimple(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      gosimplemsg,
		newGogomsg: newGoGoProtoSimpleMsg,
		gogomsg:    gogosimplemsg,
	}.test(t)
}

func TestProtoTextGoGoToGoSimple(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNidGoOptNative(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidOptNative,
		newMsg: newGoProtoNidOptNative,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoOptNative(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidOptNative,
		newMsg: newGoGoProtoNidOptNative,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoOptNative(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidOptNative,
		newGogomsg: newGoGoProtoNidOptNative,
		gogomsg:    gogovisNidOptNative,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoOptNative(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNidGoOptNative(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidOptNative,
		newMsg: newGoProtoNidOptNative,
	}.test(t)
}

//Optional Native
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoOptNative(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidOptNative,
		newMsg: newGoGoProtoNidOptNative,
	}.test(t)
}

//Optional Native
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoOptNative(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidOptNative,
		newGogomsg: newGoGoProtoNidOptNative,
		gogomsg:    gogovimNidOptNative,
	}.test(t)
}

//Optional Native
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoOptNative(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNinGoOptNative(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinOptNative,
		newMsg: newGoProtoNinOptNative,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoOptNative(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinOptNative,
		newMsg: newGoGoProtoNinOptNative,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoOptNative(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinOptNative,
		newGogomsg: newGoGoProtoNinOptNative,
		gogomsg:    gogovisNinOptNative,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoOptNative(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNinGoOptNative(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinOptNative,
		newMsg: newGoProtoNinOptNative,
	}.test(t)
}

//Optional Native
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoOptNative(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinOptNative,
		newMsg: newGoGoProtoNinOptNative,
	}.test(t)
}

//Optional Native
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoOptNative(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinOptNative,
		newGogomsg: newGoGoProtoNinOptNative,
		gogomsg:    gogovimNinOptNative,
	}.test(t)
}

//Optional Native
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoOptNative(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNidGoRepNative(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidRepNative,
		newMsg: newGoProtoNidRepNative,
	}.test(t)
}

//Repeated Native
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoRepNative(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidRepNative,
		newMsg: newGoGoProtoNidRepNative,
	}.test(t)
}

//Repeated Native
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoRepNative(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidRepNative,
		newGogomsg: newGoGoProtoNidRepNative,
		gogomsg:    gogovisNidRepNative,
	}.test(t)
}

//Repeated Native
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoRepNative(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNidGoRepNative(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidRepNative,
		newMsg: newGoProtoNidRepNative,
	}.test(t)
}

//Repeated Native
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoRepNative(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidRepNative,
		newMsg: newGoGoProtoNidRepNative,
	}.test(t)
}

//Repeated Native
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoRepNative(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidRepNative,
		newGogomsg: newGoGoProtoNidRepNative,
		gogomsg:    gogovimNidRepNative,
	}.test(t)
}

//Repeated Native
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoRepNative(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNinGoRepNative(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinRepNative,
		newMsg: newGoProtoNinRepNative,
	}.test(t)
}

//Repeated Native
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoRepNative(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinRepNative,
		newMsg: newGoGoProtoNinRepNative,
	}.test(t)
}

//Repeated Native
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoRepNative(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinRepNative,
		newGogomsg: newGoGoProtoNinRepNative,
		gogomsg:    gogovisNinRepNative,
	}.test(t)
}

//Repeated Native
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoRepNative(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNinGoRepNative(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinRepNative,
		newMsg: newGoProtoNinRepNative,
	}.test(t)
}

//Repeated Native
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoRepNative(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinRepNative,
		newMsg: newGoGoProtoNinRepNative,
	}.test(t)
}

//Repeated Native
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoRepNative(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinRepNative,
		newGogomsg: newGoGoProtoNinRepNative,
		gogomsg:    gogovimNinRepNative,
	}.test(t)
}

//Repeated Native
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoRepNative(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNidGoRepPackedNative(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidRepPackedNative,
		newMsg: newGoProtoNidRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoRepPackedNative(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidRepPackedNative,
		newMsg: newGoGoProtoNidRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoRepPackedNative(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidRepPackedNative,
		newGogomsg: newGoGoProtoNidRepPackedNative,
		gogomsg:    gogovisNidRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoRepPackedNative(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNidGoRepPackedNative(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidRepPackedNative,
		newMsg: newGoProtoNidRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoRepPackedNative(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidRepPackedNative,
		newMsg: newGoGoProtoNidRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoRepPackedNative(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidRepPackedNative,
		newGogomsg: newGoGoProtoNidRepPackedNative,
		gogomsg:    gogovimNidRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoRepPackedNative(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNinGoRepPackedNative(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinRepPackedNative,
		newMsg: newGoProtoNinRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoRepPackedNative(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinRepPackedNative,
		newMsg: newGoGoProtoNinRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoRepPackedNative(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinRepPackedNative,
		newGogomsg: newGoGoProtoNinRepPackedNative,
		gogomsg:    gogovisNinRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoRepPackedNative(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNinGoRepPackedNative(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinRepPackedNative,
		newMsg: newGoProtoNinRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoRepPackedNative(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinRepPackedNative,
		newMsg: newGoGoProtoNinRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoRepPackedNative(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinRepPackedNative,
		newGogomsg: newGoGoProtoNinRepPackedNative,
		gogomsg:    gogovimNinRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoRepPackedNative(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNidGoOptStruct(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidOptStruct,
		newMsg: newGoProtoNidOptStruct,
	}.test(t)
}

//Optional Struct
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoOptStruct(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidOptStruct,
		newMsg: newGoGoProtoNidOptStruct,
	}.test(t)
}

//Optional Struct
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoOptStruct(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidOptStruct,
		newGogomsg: newGoGoProtoNidOptStruct,
		gogomsg:    gogovisNidOptStruct,
	}.test(t)
}

//Optional Struct
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoOptStruct(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNidGoOptStruct(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidOptStruct,
		newMsg: newGoProtoNidOptStruct,
	}.test(t)
}

//Optional Struct
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoOptStruct(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidOptStruct,
		newMsg: newGoGoProtoNidOptStruct,
	}.test(t)
}

//Optional Struct
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoOptStruct(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      gozeroNidOptStruct,
		newGogomsg: newGoGoProtoNidOptStruct,
		gogomsg:    gogovimNidOptStruct,
	}.test(t)
}

//Optional Struct
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoOptStruct(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNinGoOptStruct(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinOptStruct,
		newMsg: newGoProtoNinOptStruct,
	}.test(t)
}

//Optional Struct
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoOptStruct(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinOptStruct,
		newMsg: newGoGoProtoNinOptStruct,
	}.test(t)
}

//Optional Struct
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoOptStruct(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinOptStruct,
		newGogomsg: newGoGoProtoNinOptStruct,
		gogomsg:    gogovisNinOptStruct,
	}.test(t)
}

//Optional Struct
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoOptStruct(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNinGoOptStruct(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinOptStruct,
		newMsg: newGoProtoNinOptStruct,
	}.test(t)
}

//Optional Struct
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoOptStruct(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinOptStruct,
		newMsg: newGoGoProtoNinOptStruct,
	}.test(t)
}

//Optional Struct
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoOptStruct(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinOptStruct,
		newGogomsg: newGoGoProtoNinOptStruct,
		gogomsg:    gogovimNinOptStruct,
	}.test(t)
}

//Optional Struct
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoOptStruct(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNidGoRepStruct(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidRepStruct,
		newMsg: newGoProtoNidRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoRepStruct(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidRepStruct,
		newMsg: newGoGoProtoNidRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoRepStruct(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidRepStruct,
		newGogomsg: newGoGoProtoNidRepStruct,
		gogomsg:    gogovisNidRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoRepStruct(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNidGoRepStruct(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidRepStruct,
		newMsg: newGoProtoNidRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoRepStruct(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidRepStruct,
		newMsg: newGoGoProtoNidRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoRepStruct(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidRepStruct,
		newGogomsg: newGoGoProtoNidRepStruct,
		gogomsg:    gogovimNidRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoRepStruct(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNinGoRepStruct(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinRepStruct,
		newMsg: newGoProtoNinRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoRepStruct(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinRepStruct,
		newMsg: newGoGoProtoNinRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoRepStruct(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinRepStruct,
		newGogomsg: newGoGoProtoNinRepStruct,
		gogomsg:    gogovisNinRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoRepStruct(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNinGoRepStruct(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinRepStruct,
		newMsg: newGoProtoNinRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoRepStruct(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinRepStruct,
		newMsg: newGoGoProtoNinRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoRepStruct(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinRepStruct,
		newGogomsg: newGoGoProtoNinRepStruct,
		gogomsg:    gogovimNinRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoRepStruct(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNidGoEmbeddedStruct(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidEmbeddedStruct,
		newMsg: newGoProtoNidEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoEmbeddedStruct(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidEmbeddedStruct,
		newMsg: newGoGoProtoNidEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoEmbeddedStruct(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidEmbeddedStruct,
		newGogomsg: newGoGoProtoNidEmbeddedStruct,
		gogomsg:    gogovisNidEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoEmbeddedStruct(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNidGoEmbeddedStruct(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidEmbeddedStruct,
		newMsg: newGoProtoNidEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoEmbeddedStruct(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidEmbeddedStruct,
		newMsg: newGoGoProtoNidEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoEmbeddedStruct(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      gozeroNidEmbeddedStruct,
		newGogomsg: newGoGoProtoNidEmbeddedStruct,
		gogomsg:    gogovimNidEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoEmbeddedStruct(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNinGoEmbeddedStruct(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinEmbeddedStruct,
		newMsg: newGoProtoNinEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoEmbeddedStruct(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinEmbeddedStruct,
		newMsg: newGoGoProtoNinEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoEmbeddedStruct(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinEmbeddedStruct,
		newGogomsg: newGoGoProtoNinEmbeddedStruct,
		gogomsg:    gogovisNinEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoEmbeddedStruct(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNinGoEmbeddedStruct(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinEmbeddedStruct,
		newMsg: newGoProtoNinEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoEmbeddedStruct(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinEmbeddedStruct,
		newMsg: newGoGoProtoNinEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoEmbeddedStruct(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinEmbeddedStruct,
		newGogomsg: newGoGoProtoNinEmbeddedStruct,
		gogomsg:    gogovimNinEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoEmbeddedStruct(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNidGoNestedStruct(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidNestedStruct,
		newMsg: newGoProtoNidNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoNestedStruct(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidNestedStruct,
		newMsg: newGoGoProtoNidNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoNestedStruct(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidNestedStruct,
		newGogomsg: newGoGoProtoNidNestedStruct,
		gogomsg:    gogovisNidNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoNestedStruct(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNidGoNestedStruct(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidNestedStruct,
		newMsg: newGoProtoNidNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoNestedStruct(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidNestedStruct,
		newMsg: newGoGoProtoNidNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoNestedStruct(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      gozeroNidNestedStruct,
		newGogomsg: newGoGoProtoNidNestedStruct,
		gogomsg:    gogovimNidNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoNestedStruct(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNinGoNestedStruct(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinNestedStruct,
		newMsg: newGoProtoNinNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoNestedStruct(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinNestedStruct,
		newMsg: newGoGoProtoNinNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoNestedStruct(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinNestedStruct,
		newGogomsg: newGoGoProtoNinNestedStruct,
		gogomsg:    gogovisNinNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoNestedStruct(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNinGoNestedStruct(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinNestedStruct,
		newMsg: newGoProtoNinNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoNestedStruct(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinNestedStruct,
		newMsg: newGoGoProtoNinNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoNestedStruct(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinNestedStruct,
		newGogomsg: newGoGoProtoNinNestedStruct,
		gogomsg:    gogovimNinNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoNestedStruct(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNidGoOptEnum(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidOptEnum,
		newMsg: newGoProtoNidOptEnum,
	}.test(t)
}

//Optional Enum
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoOptEnum(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidOptEnum,
		newMsg: newGoGoProtoNidOptEnum,
	}.test(t)
}

//Optional Enum
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoOptEnum(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidOptEnum,
		newGogomsg: newGoGoProtoNidOptEnum,
		gogomsg:    gogovisNidOptEnum,
	}.test(t)
}

//Optional Enum
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoOptEnum(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNidGoOptEnum(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidOptEnum,
		newMsg: newGoProtoNidOptEnum,
	}.test(t)
}

//Optional Enum
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoOptEnum(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidOptEnum,
		newMsg: newGoGoProtoNidOptEnum,
	}.test(t)
}

//Optional Enum
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoOptEnum(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      gozeroNidOptEnum,
		newGogomsg: newGoGoProtoNidOptEnum,
		gogomsg:    gogovimNidOptEnum,
	}.test(t)
}

//Optional Enum
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoOptEnum(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNinGoOptEnum(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinOptEnum,
		newMsg: newGoProtoNinOptEnum,
	}.test(t)
}

//Optional Enum
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoOptEnum(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinOptEnum,
		newMsg: newGoGoProtoNinOptEnum,
	}.test(t)
}

//Optional Enum
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoOptEnum(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinOptEnum,
		newGogomsg: newGoGoProtoNinOptEnum,
		gogomsg:    gogovisNinOptEnum,
	}.test(t)
}

//Optional Enum
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoOptEnum(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNinGoOptEnum(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinOptEnum,
		newMsg: newGoProtoNinOptEnum,
	}.test(t)
}

//Optional Enum
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoOptEnum(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinOptEnum,
		newMsg: newGoGoProtoNinOptEnum,
	}.test(t)
}

//Optional Enum
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoOptEnum(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinOptEnum,
		newGogomsg: newGoGoProtoNinOptEnum,
		gogomsg:    gogovimNinOptEnum,
	}.test(t)
}

//Optional Enum
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoOptEnum(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNidGoRepEnum(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidRepEnum,
		newMsg: newGoProtoNidRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoRepEnum(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidRepEnum,
		newMsg: newGoGoProtoNidRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoRepEnum(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidRepEnum,
		newGogomsg: newGoGoProtoNidRepEnum,
		gogomsg:    gogovisNidRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoRepEnum(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNidGoRepEnum(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidRepEnum,
		newMsg: newGoProtoNidRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoRepEnum(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidRepEnum,
		newMsg: newGoGoProtoNidRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoRepEnum(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidRepEnum,
		newGogomsg: newGoGoProtoNidRepEnum,
		gogomsg:    gogovimNidRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoRepEnum(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNinGoRepEnum(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinRepEnum,
		newMsg: newGoProtoNinRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoRepEnum(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinRepEnum,
		newMsg: newGoGoProtoNinRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoRepEnum(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinRepEnum,
		newGogomsg: newGoGoProtoNinRepEnum,
		gogomsg:    gogovisNinRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoRepEnum(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNinGoRepEnum(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinRepEnum,
		newMsg: newGoProtoNinRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoRepEnum(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinRepEnum,
		newMsg: newGoGoProtoNinRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoRepEnum(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinRepEnum,
		newGogomsg: newGoGoProtoNinRepEnum,
		gogomsg:    gogovimNinRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoRepEnum(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNidGoOptUuidAsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidOptUuidAsBytes,
		newMsg: newGoProtoNidOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoOptUuidAsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidOptUuidAsBytes,
		newMsg: newGoGoProtoNidOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoOptUuidAsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidOptUuidAsBytes,
		newGogomsg: newGoGoProtoNidOptUuidAsBytes,
		gogomsg:    gogovisNidOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoOptUuidAsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNidGoOptUuidAsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidOptUuidAsBytes,
		newMsg: newGoProtoNidOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoOptUuidAsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidOptUuidAsBytes,
		newMsg: newGoGoProtoNidOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoOptUuidAsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidOptUuidAsBytes,
		newGogomsg: newGoGoProtoNidOptUuidAsBytes,
		gogomsg:    gogovimNidOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoOptUuidAsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNinGoOptUuidAsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinOptUuidAsBytes,
		newMsg: newGoProtoNinOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoOptUuidAsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinOptUuidAsBytes,
		newMsg: newGoGoProtoNinOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoOptUuidAsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinOptUuidAsBytes,
		newGogomsg: newGoGoProtoNinOptUuidAsBytes,
		gogomsg:    gogovisNinOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoOptUuidAsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNinGoOptUuidAsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinOptUuidAsBytes,
		newMsg: newGoProtoNinOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoOptUuidAsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinOptUuidAsBytes,
		newMsg: newGoGoProtoNinOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoOptUuidAsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinOptUuidAsBytes,
		newGogomsg: newGoGoProtoNinOptUuidAsBytes,
		gogomsg:    gogovimNinOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoOptUuidAsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNidGoRepUuidAsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidRepUuidAsBytes,
		newMsg: newGoProtoNidRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoRepUuidAsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidRepUuidAsBytes,
		newMsg: newGoGoProtoNidRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoRepUuidAsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidRepUuidAsBytes,
		newGogomsg: newGoGoProtoNidRepUuidAsBytes,
		gogomsg:    gogovisNidRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoRepUuidAsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNidGoRepUuidAsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidRepUuidAsBytes,
		newMsg: newGoProtoNidRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoRepUuidAsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidRepUuidAsBytes,
		newMsg: newGoGoProtoNidRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoRepUuidAsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidRepUuidAsBytes,
		newGogomsg: newGoGoProtoNidRepUuidAsBytes,
		gogomsg:    gogovimNidRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoRepUuidAsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNinGoRepUuidAsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinRepUuidAsBytes,
		newMsg: newGoProtoNinRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoRepUuidAsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinRepUuidAsBytes,
		newMsg: newGoGoProtoNinRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoRepUuidAsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinRepUuidAsBytes,
		newGogomsg: newGoGoProtoNinRepUuidAsBytes,
		gogomsg:    gogovisNinRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoRepUuidAsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNinGoRepUuidAsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinRepUuidAsBytes,
		newMsg: newGoProtoNinRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoRepUuidAsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinRepUuidAsBytes,
		newMsg: newGoGoProtoNinRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoRepUuidAsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinRepUuidAsBytes,
		newGogomsg: newGoGoProtoNinRepUuidAsBytes,
		gogomsg:    gogovimNinRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoRepUuidAsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNidGoOptUint128AsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidOptUint128AsBytes,
		newMsg: newGoProtoNidOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoOptUint128AsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidOptUint128AsBytes,
		newMsg: newGoGoProtoNidOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoOptUint128AsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidOptUint128AsBytes,
		newGogomsg: newGoGoProtoNidOptUint128AsBytes,
		gogomsg:    gogovisNidOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoOptUint128AsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNidGoOptUint128AsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidOptUint128AsBytes,
		newMsg: newGoProtoNidOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoOptUint128AsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidOptUint128AsBytes,
		newMsg: newGoGoProtoNidOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoOptUint128AsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidOptUint128AsBytes,
		newGogomsg: newGoGoProtoNidOptUint128AsBytes,
		gogomsg:    gogovimNidOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoOptUint128AsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNinGoOptUint128AsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinOptUint128AsBytes,
		newMsg: newGoProtoNinOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoOptUint128AsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinOptUint128AsBytes,
		newMsg: newGoGoProtoNinOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoOptUint128AsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinOptUint128AsBytes,
		newGogomsg: newGoGoProtoNinOptUint128AsBytes,
		gogomsg:    gogovisNinOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoOptUint128AsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNinGoOptUint128AsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinOptUint128AsBytes,
		newMsg: newGoProtoNinOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoOptUint128AsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinOptUint128AsBytes,
		newMsg: newGoGoProtoNinOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoOptUint128AsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinOptUint128AsBytes,
		newGogomsg: newGoGoProtoNinOptUint128AsBytes,
		gogomsg:    gogovimNinOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoOptUint128AsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNidGoRepUint128AsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidRepUint128AsBytes,
		newMsg: newGoProtoNidRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoRepUint128AsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidRepUint128AsBytes,
		newMsg: newGoGoProtoNidRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoRepUint128AsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidRepUint128AsBytes,
		newGogomsg: newGoGoProtoNidRepUint128AsBytes,
		gogomsg:    gogovisNidRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoRepUint128AsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNidGoRepUint128AsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidRepUint128AsBytes,
		newMsg: newGoProtoNidRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoRepUint128AsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidRepUint128AsBytes,
		newMsg: newGoGoProtoNidRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoRepUint128AsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidRepUint128AsBytes,
		newGogomsg: newGoGoProtoNidRepUint128AsBytes,
		gogomsg:    gogovimNidRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoRepUint128AsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVisNinGoRepUint128AsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinRepUint128AsBytes,
		newMsg: newGoProtoNinRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoRepUint128AsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinRepUint128AsBytes,
		newMsg: newGoGoProtoNinRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoRepUint128AsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinRepUint128AsBytes,
		newGogomsg: newGoGoProtoNinRepUint128AsBytes,
		gogomsg:    gogovisNinRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoRepUint128AsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
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
func TestProtoTextVimNinGoRepUint128AsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinRepUint128AsBytes,
		newMsg: newGoProtoNinRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoRepUint128AsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinRepUint128AsBytes,
		newMsg: newGoGoProtoNinRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoRepUint128AsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinRepUint128AsBytes,
		newGogomsg: newGoGoProtoNinRepUint128AsBytes,
		gogomsg:    gogovimNinRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoRepUint128AsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNinRepUint128AsBytes,
		newGomsg: newGoProtoNinRepUint128AsBytes,
		gomsg:    govimNinRepUint128AsBytes,
		vis:      false,
	}.test(t)
}

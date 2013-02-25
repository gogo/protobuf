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

//Optional Uuid
//Values are available
//Nullable is false
//Go -> Go
func TestProtoTextVisNidGoOptUuid(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidOptUuid,
		newMsg: newGoProtoNidOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoOptUuid(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidOptUuid,
		newMsg: newGoGoProtoNidOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoOptUuid(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidOptUuid,
		newGogomsg: newGoGoProtoNidOptUuid,
		gogomsg:    gogovisNidOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoOptUuid(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNidOptUuid,
		newGomsg: newGoProtoNidOptUuid,
		gomsg:    govisNidOptUuid,
		vis:      true,
	}.test(t)
}

//

//Optional Uuid
//Values are not available
//Nullable is false
//Go -> Go
func TestProtoTextVimNidGoOptUuid(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidOptUuid,
		newMsg: newGoProtoNidOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoOptUuid(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidOptUuid,
		newMsg: newGoGoProtoNidOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoOptUuid(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidOptUuid,
		newGogomsg: newGoGoProtoNidOptUuid,
		gogomsg:    gogovimNidOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoOptUuid(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNidOptUuid,
		newGomsg: newGoProtoNidOptUuid,
		gomsg:    govimNidOptUuid,
		vis:      false,
	}.test(t)
}

//

//Optional Uuid
//Values are available
//Nullable is true
//Go -> Go
func TestProtoTextVisNinGoOptUuid(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinOptUuid,
		newMsg: newGoProtoNinOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoOptUuid(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinOptUuid,
		newMsg: newGoGoProtoNinOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoOptUuid(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinOptUuid,
		newGogomsg: newGoGoProtoNinOptUuid,
		gogomsg:    gogovisNinOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoOptUuid(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNinOptUuid,
		newGomsg: newGoProtoNinOptUuid,
		gomsg:    govisNinOptUuid,
		vis:      true,
	}.test(t)
}

//

//Optional Uuid
//Values are not available
//Nullable is true
//Go -> Go
func TestProtoTextVimNinGoOptUuid(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinOptUuid,
		newMsg: newGoProtoNinOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoOptUuid(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinOptUuid,
		newMsg: newGoGoProtoNinOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoOptUuid(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinOptUuid,
		newGogomsg: newGoGoProtoNinOptUuid,
		gogomsg:    gogovimNinOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoOptUuid(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNinOptUuid,
		newGomsg: newGoProtoNinOptUuid,
		gomsg:    govimNinOptUuid,
		vis:      false,
	}.test(t)
}

//

//Repeated Uuid
//Values are available
//Nullable is false
//Go -> Go
func TestProtoTextVisNidGoRepUuid(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidRepUuid,
		newMsg: newGoProtoNidRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoRepUuid(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidRepUuid,
		newMsg: newGoGoProtoNidRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoRepUuid(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidRepUuid,
		newGogomsg: newGoGoProtoNidRepUuid,
		gogomsg:    gogovisNidRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoRepUuid(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNidRepUuid,
		newGomsg: newGoProtoNidRepUuid,
		gomsg:    govisNidRepUuid,
		vis:      true,
	}.test(t)
}

//

//Repeated Uuid
//Values are not available
//Nullable is false
//Go -> Go
func TestProtoTextVimNidGoRepUuid(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidRepUuid,
		newMsg: newGoProtoNidRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoRepUuid(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidRepUuid,
		newMsg: newGoGoProtoNidRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoRepUuid(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidRepUuid,
		newGogomsg: newGoGoProtoNidRepUuid,
		gogomsg:    gogovimNidRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoRepUuid(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNidRepUuid,
		newGomsg: newGoProtoNidRepUuid,
		gomsg:    govimNidRepUuid,
		vis:      false,
	}.test(t)
}

//

//Repeated Uuid
//Values are available
//Nullable is true
//Go -> Go
func TestProtoTextVisNinGoRepUuid(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinRepUuid,
		newMsg: newGoProtoNinRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoRepUuid(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinRepUuid,
		newMsg: newGoGoProtoNinRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoRepUuid(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinRepUuid,
		newGogomsg: newGoGoProtoNinRepUuid,
		gogomsg:    gogovisNinRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoRepUuid(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNinRepUuid,
		newGomsg: newGoProtoNinRepUuid,
		gomsg:    govisNinRepUuid,
		vis:      true,
	}.test(t)
}

//

//Repeated Uuid
//Values are not available
//Nullable is true
//Go -> Go
func TestProtoTextVimNinGoRepUuid(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinRepUuid,
		newMsg: newGoProtoNinRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoRepUuid(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinRepUuid,
		newMsg: newGoGoProtoNinRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoRepUuid(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinRepUuid,
		newGogomsg: newGoGoProtoNinRepUuid,
		gogomsg:    gogovimNinRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoRepUuid(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNinRepUuid,
		newGomsg: newGoProtoNinRepUuid,
		gomsg:    govimNinRepUuid,
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

//Optional Uint128
//Values are available
//Nullable is false
//Go -> Go
func TestProtoTextVisNidGoOptUint128(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidOptUint128,
		newMsg: newGoProtoNidOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoOptUint128(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidOptUint128,
		newMsg: newGoGoProtoNidOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoOptUint128(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidOptUint128,
		newGogomsg: newGoGoProtoNidOptUint128,
		gogomsg:    gogovisNidOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoOptUint128(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNidOptUint128,
		newGomsg: newGoProtoNidOptUint128,
		gomsg:    govisNidOptUint128,
		vis:      true,
	}.test(t)
}

//

//Optional Uint128
//Values are not available
//Nullable is false
//Go -> Go
func TestProtoTextVimNidGoOptUint128(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidOptUint128,
		newMsg: newGoProtoNidOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoOptUint128(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidOptUint128,
		newMsg: newGoGoProtoNidOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoOptUint128(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidOptUint128,
		newGogomsg: newGoGoProtoNidOptUint128,
		gogomsg:    gogovimNidOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoOptUint128(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNidOptUint128,
		newGomsg: newGoProtoNidOptUint128,
		gomsg:    gozeroOptUint128,
		vis:      false,
	}.test(t)
}

//

//Optional Uint128
//Values are available
//Nullable is true
//Go -> Go
func TestProtoTextVisNinGoOptUint128(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinOptUint128,
		newMsg: newGoProtoNinOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoOptUint128(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinOptUint128,
		newMsg: newGoGoProtoNinOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoOptUint128(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinOptUint128,
		newGogomsg: newGoGoProtoNinOptUint128,
		gogomsg:    gogovisNinOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoOptUint128(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNinOptUint128,
		newGomsg: newGoProtoNinOptUint128,
		gomsg:    govisNinOptUint128,
		vis:      true,
	}.test(t)
}

//

//Optional Uint128
//Values are not available
//Nullable is true
//Go -> Go
func TestProtoTextVimNinGoOptUint128(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinOptUint128,
		newMsg: newGoProtoNinOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoOptUint128(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinOptUint128,
		newMsg: newGoGoProtoNinOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoOptUint128(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinOptUint128,
		newGogomsg: newGoGoProtoNinOptUint128,
		gogomsg:    gogovimNinOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoOptUint128(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNinOptUint128,
		newGomsg: newGoProtoNinOptUint128,
		gomsg:    govimNinOptUint128,
		vis:      false,
	}.test(t)
}

//

//Repeated Uint128
//Values are available
//Nullable is false
//Go -> Go
func TestProtoTextVisNidGoRepUint128(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidRepUint128,
		newMsg: newGoProtoNidRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoRepUint128(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidRepUint128,
		newMsg: newGoGoProtoNidRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoRepUint128(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidRepUint128,
		newGogomsg: newGoGoProtoNidRepUint128,
		gogomsg:    gogovisNidRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoRepUint128(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNidRepUint128,
		newGomsg: newGoProtoNidRepUint128,
		gomsg:    govisNidRepUint128,
		vis:      true,
	}.test(t)
}

//

//Repeated Uint128
//Values are not available
//Nullable is false
//Go -> Go
func TestProtoTextVimNidGoRepUint128(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidRepUint128,
		newMsg: newGoProtoNidRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoRepUint128(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidRepUint128,
		newMsg: newGoGoProtoNidRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoRepUint128(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidRepUint128,
		newGogomsg: newGoGoProtoNidRepUint128,
		gogomsg:    gogovimNidRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoRepUint128(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNidRepUint128,
		newGomsg: newGoProtoNidRepUint128,
		gomsg:    govimNidRepUint128,
		vis:      false,
	}.test(t)
}

//

//Repeated Uint128
//Values are available
//Nullable is true
//Go -> Go
func TestProtoTextVisNinGoRepUint128(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinRepUint128,
		newMsg: newGoProtoNinRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoRepUint128(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinRepUint128,
		newMsg: newGoGoProtoNinRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoRepUint128(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinRepUint128,
		newGogomsg: newGoGoProtoNinRepUint128,
		gogomsg:    gogovisNinRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoRepUint128(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNinRepUint128,
		newGomsg: newGoProtoNinRepUint128,
		gomsg:    govisNinRepUint128,
		vis:      true,
	}.test(t)
}

//

//Repeated Uint128
//Values are not available
//Nullable is true
//Go -> Go
func TestProtoTextVimNinGoRepUint128(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinRepUint128,
		newMsg: newGoProtoNinRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoRepUint128(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinRepUint128,
		newMsg: newGoGoProtoNinRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoRepUint128(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinRepUint128,
		newGogomsg: newGoGoProtoNinRepUint128,
		gogomsg:    gogovimNinRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoRepUint128(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNinRepUint128,
		newGomsg: newGoProtoNinRepUint128,
		gomsg:    govimNinRepUint128,
		vis:      false,
	}.test(t)
}

//

//Optional Uint16
//Values are available
//Nullable is false
//Go -> Go
func TestProtoTextVisNidGoOptUint16(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidOptUint16,
		newMsg: newGoProtoNidOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoOptUint16(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidOptUint16,
		newMsg: newGoGoProtoNidOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoOptUint16(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidOptUint16,
		newGogomsg: newGoGoProtoNidOptUint16,
		gogomsg:    gogovisNidOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoOptUint16(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNidOptUint16,
		newGomsg: newGoProtoNidOptUint16,
		gomsg:    govisNidOptUint16,
		vis:      true,
	}.test(t)
}

//

//Optional Uint16
//Values are not available
//Nullable is false
//Go -> Go
func TestProtoTextVimNidGoOptUint16(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidOptUint16,
		newMsg: newGoProtoNidOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoOptUint16(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidOptUint16,
		newMsg: newGoGoProtoNidOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoOptUint16(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidOptUint16,
		newGogomsg: newGoGoProtoNidOptUint16,
		gogomsg:    gogovimNidOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoOptUint16(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNidOptUint16,
		newGomsg: newGoProtoNidOptUint16,
		gomsg:    gozeroOptUint16,
		vis:      false,
	}.test(t)
}

//

//Optional Uint16
//Values are available
//Nullable is true
//Go -> Go
func TestProtoTextVisNinGoOptUint16(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinOptUint16,
		newMsg: newGoProtoNinOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoOptUint16(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinOptUint16,
		newMsg: newGoGoProtoNinOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoOptUint16(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinOptUint16,
		newGogomsg: newGoGoProtoNinOptUint16,
		gogomsg:    gogovisNinOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoOptUint16(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNinOptUint16,
		newGomsg: newGoProtoNinOptUint16,
		gomsg:    govisNinOptUint16,
		vis:      true,
	}.test(t)
}

//

//Optional Uint16
//Values are not available
//Nullable is true
//Go -> Go
func TestProtoTextVimNinGoOptUint16(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinOptUint16,
		newMsg: newGoProtoNinOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoOptUint16(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinOptUint16,
		newMsg: newGoGoProtoNinOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoOptUint16(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinOptUint16,
		newGogomsg: newGoGoProtoNinOptUint16,
		gogomsg:    gogovimNinOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoOptUint16(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNinOptUint16,
		newGomsg: newGoProtoNinOptUint16,
		gomsg:    govimNinOptUint16,
		vis:      false,
	}.test(t)
}

//

//Repeated Uint16
//Values are available
//Nullable is false
//Go -> Go
func TestProtoTextVisNidGoRepUint16(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidRepUint16,
		newMsg: newGoProtoNidRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoRepUint16(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidRepUint16,
		newMsg: newGoGoProtoNidRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoRepUint16(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidRepUint16,
		newGogomsg: newGoGoProtoNidRepUint16,
		gogomsg:    gogovisNidRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoRepUint16(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNidRepUint16,
		newGomsg: newGoProtoNidRepUint16,
		gomsg:    govisNidRepUint16,
		vis:      true,
	}.test(t)
}

//

//Repeated Uint16
//Values are not available
//Nullable is false
//Go -> Go
func TestProtoTextVimNidGoRepUint16(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidRepUint16,
		newMsg: newGoProtoNidRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoRepUint16(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidRepUint16,
		newMsg: newGoGoProtoNidRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoRepUint16(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidRepUint16,
		newGogomsg: newGoGoProtoNidRepUint16,
		gogomsg:    gogovimNidRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoRepUint16(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNidRepUint16,
		newGomsg: newGoProtoNidRepUint16,
		gomsg:    govimNidRepUint16,
		vis:      false,
	}.test(t)
}

//

//Repeated Uint16
//Values are available
//Nullable is true
//Go -> Go
func TestProtoTextVisNinGoRepUint16(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinRepUint16,
		newMsg: newGoProtoNinRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoRepUint16(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinRepUint16,
		newMsg: newGoGoProtoNinRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoRepUint16(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinRepUint16,
		newGogomsg: newGoGoProtoNinRepUint16,
		gogomsg:    gogovisNinRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoRepUint16(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNinRepUint16,
		newGomsg: newGoProtoNinRepUint16,
		gomsg:    govisNinRepUint16,
		vis:      true,
	}.test(t)
}

//

//Repeated Uint16
//Values are not available
//Nullable is true
//Go -> Go
func TestProtoTextVimNinGoRepUint16(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinRepUint16,
		newMsg: newGoProtoNinRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoRepUint16(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinRepUint16,
		newMsg: newGoGoProtoNinRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoRepUint16(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinRepUint16,
		newGogomsg: newGoGoProtoNinRepUint16,
		gogomsg:    gogovimNinRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoRepUint16(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNinRepUint16,
		newGomsg: newGoProtoNinRepUint16,
		gomsg:    govimNinRepUint16,
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

//

//Optional Uint16AsBytes
//Values are available
//Nullable is false
//Go -> Go
func TestProtoTextVisNidGoOptUint16AsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidOptUint16AsBytes,
		newMsg: newGoProtoNidOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoOptUint16AsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidOptUint16AsBytes,
		newMsg: newGoGoProtoNidOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoOptUint16AsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidOptUint16AsBytes,
		newGogomsg: newGoGoProtoNidOptUint16AsBytes,
		gogomsg:    gogovisNidOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoOptUint16AsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNidOptUint16AsBytes,
		newGomsg: newGoProtoNidOptUint16AsBytes,
		gomsg:    govisNidOptUint16AsBytes,
		vis:      true,
	}.test(t)
}

//

//Optional Uint16AsBytes
//Values are not available
//Nullable is false
//Go -> Go
func TestProtoTextVimNidGoOptUint16AsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidOptUint16AsBytes,
		newMsg: newGoProtoNidOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoOptUint16AsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidOptUint16AsBytes,
		newMsg: newGoGoProtoNidOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoOptUint16AsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidOptUint16AsBytes,
		newGogomsg: newGoGoProtoNidOptUint16AsBytes,
		gogomsg:    gogovimNidOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoOptUint16AsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNidOptUint16AsBytes,
		newGomsg: newGoProtoNidOptUint16AsBytes,
		gomsg:    gozeroOptUint16AsBytes,
		vis:      false,
	}.test(t)
}

//

//Optional Uint16AsBytes
//Values are available
//Nullable is true
//Go -> Go
func TestProtoTextVisNinGoOptUint16AsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinOptUint16AsBytes,
		newMsg: newGoProtoNinOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoOptUint16AsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinOptUint16AsBytes,
		newMsg: newGoGoProtoNinOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoOptUint16AsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinOptUint16AsBytes,
		newGogomsg: newGoGoProtoNinOptUint16AsBytes,
		gogomsg:    gogovisNinOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoOptUint16AsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNinOptUint16AsBytes,
		newGomsg: newGoProtoNinOptUint16AsBytes,
		gomsg:    govisNinOptUint16AsBytes,
		vis:      true,
	}.test(t)
}

//

//Optional Uint16AsBytes
//Values are not available
//Nullable is true
//Go -> Go
func TestProtoTextVimNinGoOptUint16AsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinOptUint16AsBytes,
		newMsg: newGoProtoNinOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoOptUint16AsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinOptUint16AsBytes,
		newMsg: newGoGoProtoNinOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoOptUint16AsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinOptUint16AsBytes,
		newGogomsg: newGoGoProtoNinOptUint16AsBytes,
		gogomsg:    gogovimNinOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoOptUint16AsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNinOptUint16AsBytes,
		newGomsg: newGoProtoNinOptUint16AsBytes,
		gomsg:    govimNinOptUint16AsBytes,
		vis:      false,
	}.test(t)
}

//

//Repeated Uint16AsBytes
//Values are available
//Nullable is false
//Go -> Go
func TestProtoTextVisNidGoRepUint16AsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidRepUint16AsBytes,
		newMsg: newGoProtoNidRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoRepUint16AsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidRepUint16AsBytes,
		newMsg: newGoGoProtoNidRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoRepUint16AsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidRepUint16AsBytes,
		newGogomsg: newGoGoProtoNidRepUint16AsBytes,
		gogomsg:    gogovisNidRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoRepUint16AsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNidRepUint16AsBytes,
		newGomsg: newGoProtoNidRepUint16AsBytes,
		gomsg:    govisNidRepUint16AsBytes,
		vis:      true,
	}.test(t)
}

//

//Repeated Uint16AsBytes
//Values are not available
//Nullable is false
//Go -> Go
func TestProtoTextVimNidGoRepUint16AsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidRepUint16AsBytes,
		newMsg: newGoProtoNidRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoRepUint16AsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidRepUint16AsBytes,
		newMsg: newGoGoProtoNidRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoRepUint16AsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidRepUint16AsBytes,
		newGogomsg: newGoGoProtoNidRepUint16AsBytes,
		gogomsg:    gogovimNidRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoRepUint16AsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNidRepUint16AsBytes,
		newGomsg: newGoProtoNidRepUint16AsBytes,
		gomsg:    govimNidRepUint16AsBytes,
		vis:      false,
	}.test(t)
}

//

//Repeated Uint16AsBytes
//Values are available
//Nullable is true
//Go -> Go
func TestProtoTextVisNinGoRepUint16AsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinRepUint16AsBytes,
		newMsg: newGoProtoNinRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoRepUint16AsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinRepUint16AsBytes,
		newMsg: newGoGoProtoNinRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoRepUint16AsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinRepUint16AsBytes,
		newGogomsg: newGoGoProtoNinRepUint16AsBytes,
		gogomsg:    gogovisNinRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoRepUint16AsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNinRepUint16AsBytes,
		newGomsg: newGoProtoNinRepUint16AsBytes,
		gomsg:    govisNinRepUint16AsBytes,
		vis:      true,
	}.test(t)
}

//

//Repeated Uint16AsBytes
//Values are not available
//Nullable is true
//Go -> Go
func TestProtoTextVimNinGoRepUint16AsBytes(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinRepUint16AsBytes,
		newMsg: newGoProtoNinRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoRepUint16AsBytes(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinRepUint16AsBytes,
		newMsg: newGoGoProtoNinRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoRepUint16AsBytes(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinRepUint16AsBytes,
		newGogomsg: newGoGoProtoNinRepUint16AsBytes,
		gogomsg:    gogovimNinRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoRepUint16AsBytes(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNinRepUint16AsBytes,
		newGomsg: newGoProtoNinRepUint16AsBytes,
		gomsg:    govimNinRepUint16AsBytes,
		vis:      false,
	}.test(t)
}

//

//Optional Uint16AsUint32
//Values are available
//Nullable is false
//Go -> Go
func TestProtoTextVisNidGoOptUint16AsUint32(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidOptUint16AsUint32,
		newMsg: newGoProtoNidOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoOptUint16AsUint32(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidOptUint16AsUint32,
		newMsg: newGoGoProtoNidOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoOptUint16AsUint32(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidOptUint16AsUint32,
		newGogomsg: newGoGoProtoNidOptUint16AsUint32,
		gogomsg:    gogovisNidOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoOptUint16AsUint32(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNidOptUint16AsUint32,
		newGomsg: newGoProtoNidOptUint16AsUint32,
		gomsg:    govisNidOptUint16AsUint32,
		vis:      true,
	}.test(t)
}

//

//Optional Uint16AsUint32
//Values are not available
//Nullable is false
//Go -> Go
func TestProtoTextVimNidGoOptUint16AsUint32(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidOptUint16AsUint32,
		newMsg: newGoProtoNidOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoOptUint16AsUint32(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidOptUint16AsUint32,
		newMsg: newGoGoProtoNidOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoOptUint16AsUint32(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidOptUint16AsUint32,
		newGogomsg: newGoGoProtoNidOptUint16AsUint32,
		gogomsg:    gogovimNidOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoOptUint16AsUint32(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNidOptUint16AsUint32,
		newGomsg: newGoProtoNidOptUint16AsUint32,
		gomsg:    gozeroOptUint16AsUint32,
		vis:      false,
	}.test(t)
}

//

//Optional Uint16AsUint32
//Values are available
//Nullable is true
//Go -> Go
func TestProtoTextVisNinGoOptUint16AsUint32(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinOptUint16AsUint32,
		newMsg: newGoProtoNinOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoOptUint16AsUint32(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinOptUint16AsUint32,
		newMsg: newGoGoProtoNinOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoOptUint16AsUint32(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinOptUint16AsUint32,
		newGogomsg: newGoGoProtoNinOptUint16AsUint32,
		gogomsg:    gogovisNinOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoOptUint16AsUint32(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNinOptUint16AsUint32,
		newGomsg: newGoProtoNinOptUint16AsUint32,
		gomsg:    govisNinOptUint16AsUint32,
		vis:      true,
	}.test(t)
}

//

//Optional Uint16AsUint32
//Values are not available
//Nullable is true
//Go -> Go
func TestProtoTextVimNinGoOptUint16AsUint32(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinOptUint16AsUint32,
		newMsg: newGoProtoNinOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoOptUint16AsUint32(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinOptUint16AsUint32,
		newMsg: newGoGoProtoNinOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoOptUint16AsUint32(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinOptUint16AsUint32,
		newGogomsg: newGoGoProtoNinOptUint16AsUint32,
		gogomsg:    gogovimNinOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoOptUint16AsUint32(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNinOptUint16AsUint32,
		newGomsg: newGoProtoNinOptUint16AsUint32,
		gomsg:    govimNinOptUint16AsUint32,
		vis:      false,
	}.test(t)
}

//

//Repeated Uint16AsUint32
//Values are available
//Nullable is false
//Go -> Go
func TestProtoTextVisNidGoRepUint16AsUint32(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidRepUint16AsUint32,
		newMsg: newGoProtoNidRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoRepUint16AsUint32(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidRepUint16AsUint32,
		newMsg: newGoGoProtoNidRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoRepUint16AsUint32(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidRepUint16AsUint32,
		newGogomsg: newGoGoProtoNidRepUint16AsUint32,
		gogomsg:    gogovisNidRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoRepUint16AsUint32(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNidRepUint16AsUint32,
		newGomsg: newGoProtoNidRepUint16AsUint32,
		gomsg:    govisNidRepUint16AsUint32,
		vis:      true,
	}.test(t)
}

//

//Repeated Uint16AsUint32
//Values are not available
//Nullable is false
//Go -> Go
func TestProtoTextVimNidGoRepUint16AsUint32(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidRepUint16AsUint32,
		newMsg: newGoProtoNidRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoRepUint16AsUint32(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidRepUint16AsUint32,
		newMsg: newGoGoProtoNidRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoRepUint16AsUint32(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidRepUint16AsUint32,
		newGogomsg: newGoGoProtoNidRepUint16AsUint32,
		gogomsg:    gogovimNidRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoRepUint16AsUint32(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNidRepUint16AsUint32,
		newGomsg: newGoProtoNidRepUint16AsUint32,
		gomsg:    govimNidRepUint16AsUint32,
		vis:      false,
	}.test(t)
}

//

//Repeated Uint16AsUint32
//Values are available
//Nullable is true
//Go -> Go
func TestProtoTextVisNinGoRepUint16AsUint32(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinRepUint16AsUint32,
		newMsg: newGoProtoNinRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoRepUint16AsUint32(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinRepUint16AsUint32,
		newMsg: newGoGoProtoNinRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoRepUint16AsUint32(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinRepUint16AsUint32,
		newGogomsg: newGoGoProtoNinRepUint16AsUint32,
		gogomsg:    gogovisNinRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoRepUint16AsUint32(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNinRepUint16AsUint32,
		newGomsg: newGoProtoNinRepUint16AsUint32,
		gomsg:    govisNinRepUint16AsUint32,
		vis:      true,
	}.test(t)
}

//

//Repeated Uint16AsUint32
//Values are not available
//Nullable is true
//Go -> Go
func TestProtoTextVimNinGoRepUint16AsUint32(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinRepUint16AsUint32,
		newMsg: newGoProtoNinRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoRepUint16AsUint32(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinRepUint16AsUint32,
		newMsg: newGoGoProtoNinRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoRepUint16AsUint32(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinRepUint16AsUint32,
		newGogomsg: newGoGoProtoNinRepUint16AsUint32,
		gogomsg:    gogovimNinRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoRepUint16AsUint32(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNinRepUint16AsUint32,
		newGomsg: newGoProtoNinRepUint16AsUint32,
		gomsg:    govimNinRepUint16AsUint32,
		vis:      false,
	}.test(t)
}

//

//Optional EnumUint32
//Values are available
//Nullable is false
//Go -> Go
func TestProtoTextVisNidGoOptEnumUint32(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidOptEnumUint32,
		newMsg: newGoProtoNidOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoOptEnumUint32(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidOptEnumUint32,
		newMsg: newGoGoProtoNidOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoOptEnumUint32(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidOptEnumUint32,
		newGogomsg: newGoGoProtoNidOptEnumUint32,
		gogomsg:    gogovisNidOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoOptEnumUint32(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNidOptEnumUint32,
		newGomsg: newGoProtoNidOptEnumUint32,
		gomsg:    govisNidOptEnumUint32,
		vis:      true,
	}.test(t)
}

//

//Optional EnumUint32
//Values are not available
//Nullable is false
//Go -> Go
func TestProtoTextVimNidGoOptEnumUint32(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidOptEnumUint32,
		newMsg: newGoProtoNidOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoOptEnumUint32(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidOptEnumUint32,
		newMsg: newGoGoProtoNidOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoOptEnumUint32(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      gozeroNidOptEnumUint32,
		newGogomsg: newGoGoProtoNidOptEnumUint32,
		gogomsg:    gogovimNidOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoOptEnumUint32(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNidOptEnumUint32,
		newGomsg: newGoProtoNidOptEnumUint32,
		gomsg:    gozeroNidOptEnumUint32,
		vis:      false,
	}.test(t)
}

//

//Optional EnumUint32
//Values are available
//Nullable is true
//Go -> Go
func TestProtoTextVisNinGoOptEnumUint32(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinOptEnumUint32,
		newMsg: newGoProtoNinOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoOptEnumUint32(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinOptEnumUint32,
		newMsg: newGoGoProtoNinOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoOptEnumUint32(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinOptEnumUint32,
		newGogomsg: newGoGoProtoNinOptEnumUint32,
		gogomsg:    gogovisNinOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoOptEnumUint32(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNinOptEnumUint32,
		newGomsg: newGoProtoNinOptEnumUint32,
		gomsg:    govisNinOptEnumUint32,
		vis:      true,
	}.test(t)
}

//

//Optional EnumUint32
//Values are not available
//Nullable is true
//Go -> Go
func TestProtoTextVimNinGoOptEnumUint32(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinOptEnumUint32,
		newMsg: newGoProtoNinOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoOptEnumUint32(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinOptEnumUint32,
		newMsg: newGoGoProtoNinOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoOptEnumUint32(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinOptEnumUint32,
		newGogomsg: newGoGoProtoNinOptEnumUint32,
		gogomsg:    gogovimNinOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoOptEnumUint32(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNinOptEnumUint32,
		newGomsg: newGoProtoNinOptEnumUint32,
		gomsg:    govimNinOptEnumUint32,
		vis:      false,
	}.test(t)
}

//

//Repeated EnumUint32
//Values are available
//Nullable is false
//Go -> Go
func TestProtoTextVisNidGoRepEnumUint32(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidRepEnumUint32,
		newMsg: newGoProtoNidRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoRepEnumUint32(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidRepEnumUint32,
		newMsg: newGoGoProtoNidRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoRepEnumUint32(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidRepEnumUint32,
		newGogomsg: newGoGoProtoNidRepEnumUint32,
		gogomsg:    gogovisNidRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoRepEnumUint32(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNidRepEnumUint32,
		newGomsg: newGoProtoNidRepEnumUint32,
		gomsg:    govisNidRepEnumUint32,
		vis:      true,
	}.test(t)
}

//

//Repeated EnumUint32
//Values are not available
//Nullable is false
//Go -> Go
func TestProtoTextVimNidGoRepEnumUint32(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidRepEnumUint32,
		newMsg: newGoProtoNidRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoRepEnumUint32(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidRepEnumUint32,
		newMsg: newGoGoProtoNidRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoRepEnumUint32(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidRepEnumUint32,
		newGogomsg: newGoGoProtoNidRepEnumUint32,
		gogomsg:    gogovimNidRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoRepEnumUint32(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNidRepEnumUint32,
		newGomsg: newGoProtoNidRepEnumUint32,
		gomsg:    govimNidRepEnumUint32,
		vis:      false,
	}.test(t)
}

//

//Repeated EnumUint32
//Values are available
//Nullable is true
//Go -> Go
func TestProtoTextVisNinGoRepEnumUint32(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinRepEnumUint32,
		newMsg: newGoProtoNinRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoRepEnumUint32(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinRepEnumUint32,
		newMsg: newGoGoProtoNinRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoRepEnumUint32(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinRepEnumUint32,
		newGogomsg: newGoGoProtoNinRepEnumUint32,
		gogomsg:    gogovisNinRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoRepEnumUint32(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNinRepEnumUint32,
		newGomsg: newGoProtoNinRepEnumUint32,
		gomsg:    govisNinRepEnumUint32,
		vis:      true,
	}.test(t)
}

//

//Repeated EnumUint32
//Values are not available
//Nullable is true
//Go -> Go
func TestProtoTextVimNinGoRepEnumUint32(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinRepEnumUint32,
		newMsg: newGoProtoNinRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoRepEnumUint32(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinRepEnumUint32,
		newMsg: newGoGoProtoNinRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoRepEnumUint32(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinRepEnumUint32,
		newGogomsg: newGoGoProtoNinRepEnumUint32,
		gogomsg:    gogovimNinRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoRepEnumUint32(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNinRepEnumUint32,
		newGomsg: newGoProtoNinRepEnumUint32,
		gomsg:    govimNinRepEnumUint32,
		vis:      false,
	}.test(t)
}

//

//Optional EnumUint16
//Values are available
//Nullable is false
//Go -> Go
func TestProtoTextVisNidGoOptEnumUint16(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidOptEnumUint16,
		newMsg: newGoProtoNidOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoOptEnumUint16(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidOptEnumUint16,
		newMsg: newGoGoProtoNidOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoOptEnumUint16(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidOptEnumUint16,
		newGogomsg: newGoGoProtoNidOptEnumUint16,
		gogomsg:    gogovisNidOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoOptEnumUint16(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNidOptEnumUint16,
		newGomsg: newGoProtoNidOptEnumUint16,
		gomsg:    govisNidOptEnumUint16,
		vis:      true,
	}.test(t)
}

//

//Optional EnumUint16
//Values are not available
//Nullable is false
//Go -> Go
func TestProtoTextVimNidGoOptEnumUint16(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidOptEnumUint16,
		newMsg: newGoProtoNidOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoOptEnumUint16(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidOptEnumUint16,
		newMsg: newGoGoProtoNidOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoOptEnumUint16(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      gozeroNidOptEnumUint16,
		newGogomsg: newGoGoProtoNidOptEnumUint16,
		gogomsg:    gogovimNidOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoOptEnumUint16(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNidOptEnumUint16,
		newGomsg: newGoProtoNidOptEnumUint16,
		gomsg:    gozeroNidOptEnumUint16,
		vis:      false,
	}.test(t)
}

//

//Optional EnumUint16
//Values are available
//Nullable is true
//Go -> Go
func TestProtoTextVisNinGoOptEnumUint16(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinOptEnumUint16,
		newMsg: newGoProtoNinOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoOptEnumUint16(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinOptEnumUint16,
		newMsg: newGoGoProtoNinOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoOptEnumUint16(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinOptEnumUint16,
		newGogomsg: newGoGoProtoNinOptEnumUint16,
		gogomsg:    gogovisNinOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoOptEnumUint16(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNinOptEnumUint16,
		newGomsg: newGoProtoNinOptEnumUint16,
		gomsg:    govisNinOptEnumUint16,
		vis:      true,
	}.test(t)
}

//

//Optional EnumUint16
//Values are not available
//Nullable is true
//Go -> Go
func TestProtoTextVimNinGoOptEnumUint16(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinOptEnumUint16,
		newMsg: newGoProtoNinOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoOptEnumUint16(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinOptEnumUint16,
		newMsg: newGoGoProtoNinOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoOptEnumUint16(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinOptEnumUint16,
		newGogomsg: newGoGoProtoNinOptEnumUint16,
		gogomsg:    gogovimNinOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoOptEnumUint16(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNinOptEnumUint16,
		newGomsg: newGoProtoNinOptEnumUint16,
		gomsg:    govimNinOptEnumUint16,
		vis:      false,
	}.test(t)
}

//

//Repeated EnumUint16
//Values are available
//Nullable is false
//Go -> Go
func TestProtoTextVisNidGoRepEnumUint16(t *testing.T) {
	testProtoTextGo{
		msg:    govisNidRepEnumUint16,
		newMsg: newGoProtoNidRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVisNidGoGoRepEnumUint16(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNidRepEnumUint16,
		newMsg: newGoGoProtoNidRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoTextVisNidGoToGoGoRepEnumUint16(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNidRepEnumUint16,
		newGogomsg: newGoGoProtoNidRepEnumUint16,
		gogomsg:    gogovisNidRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoTextVisNidGoGoToGoRepEnumUint16(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNidRepEnumUint16,
		newGomsg: newGoProtoNidRepEnumUint16,
		gomsg:    govisNidRepEnumUint16,
		vis:      true,
	}.test(t)
}

//

//Repeated EnumUint16
//Values are not available
//Nullable is false
//Go -> Go
func TestProtoTextVimNidGoRepEnumUint16(t *testing.T) {
	testProtoTextGo{
		msg:    govimNidRepEnumUint16,
		newMsg: newGoProtoNidRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoTextVimNidGoGoRepEnumUint16(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNidRepEnumUint16,
		newMsg: newGoGoProtoNidRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoTextVimNidGoToGoGoRepEnumUint16(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNidRepEnumUint16,
		newGogomsg: newGoGoProtoNidRepEnumUint16,
		gogomsg:    gogovimNidRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoTextVimNidGoGoToGoRepEnumUint16(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNidRepEnumUint16,
		newGomsg: newGoProtoNidRepEnumUint16,
		gomsg:    govimNidRepEnumUint16,
		vis:      false,
	}.test(t)
}

//

//Repeated EnumUint16
//Values are available
//Nullable is true
//Go -> Go
func TestProtoTextVisNinGoRepEnumUint16(t *testing.T) {
	testProtoTextGo{
		msg:    govisNinRepEnumUint16,
		newMsg: newGoProtoNinRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVisNinGoGoRepEnumUint16(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovisNinRepEnumUint16,
		newMsg: newGoGoProtoNinRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoTextVisNinGoToGoGoRepEnumUint16(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govisNinRepEnumUint16,
		newGogomsg: newGoGoProtoNinRepEnumUint16,
		gogomsg:    gogovisNinRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoTextVisNinGoGoToGoRepEnumUint16(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovisNinRepEnumUint16,
		newGomsg: newGoProtoNinRepEnumUint16,
		gomsg:    govisNinRepEnumUint16,
		vis:      true,
	}.test(t)
}

//

//Repeated EnumUint16
//Values are not available
//Nullable is true
//Go -> Go
func TestProtoTextVimNinGoRepEnumUint16(t *testing.T) {
	testProtoTextGo{
		msg:    govimNinRepEnumUint16,
		newMsg: newGoProtoNinRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoTextVimNinGoGoRepEnumUint16(t *testing.T) {
	testProtoTextGoGo{
		msg:    gogovimNinRepEnumUint16,
		newMsg: newGoGoProtoNinRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoTextVimNinGoToGoGoRepEnumUint16(t *testing.T) {
	testProtoTextGoToGoGo{
		gomsg:      govimNinRepEnumUint16,
		newGogomsg: newGoGoProtoNinRepEnumUint16,
		gogomsg:    gogovimNinRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoTextVimNinGoGoToGoRepEnumUint16(t *testing.T) {
	testProtoTextGoGoToGo{
		gogomsg:  gogovimNinRepEnumUint16,
		newGomsg: newGoProtoNinRepEnumUint16,
		gomsg:    govimNinRepEnumUint16,
		vis:      false,
	}.test(t)
}

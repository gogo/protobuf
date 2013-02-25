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

func TestProtoGoSimple(t *testing.T) {
	testProtoGo{
		msg:    gosimplemsg,
		newMsg: newGoProtoSimpleMsg,
	}.test(t)
}

func TestProtoGoGoSimple(t *testing.T) {
	testProtoGoGo{
		msg:    gogosimplemsg,
		newMsg: newGoGoProtoSimpleMsg,
	}.test(t)
}

func TestProtoGoToGoGoSimple(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      gosimplemsg,
		newGogomsg: newGoGoProtoSimpleMsg,
		gogomsg:    gogosimplemsg,
	}.test(t)
}

func TestProtoGoGoToGoSimple(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoOptNative(t *testing.T) {
	testProtoGo{
		msg:    govisNidOptNative,
		newMsg: newGoProtoNidOptNative,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoOptNative(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidOptNative,
		newMsg: newGoGoProtoNidOptNative,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoOptNative(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidOptNative,
		newGogomsg: newGoGoProtoNidOptNative,
		gogomsg:    gogovisNidOptNative,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoOptNative(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoOptNative(t *testing.T) {
	testProtoGo{
		msg:    govimNidOptNative,
		newMsg: newGoProtoNidOptNative,
	}.test(t)
}

//Optional Native
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoOptNative(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidOptNative,
		newMsg: newGoGoProtoNidOptNative,
	}.test(t)
}

//Optional Native
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoOptNative(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidOptNative,
		newGogomsg: newGoGoProtoNidOptNative,
		gogomsg:    gogovimNidOptNative,
	}.test(t)
}

//Optional Native
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoOptNative(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoOptNative(t *testing.T) {
	testProtoGo{
		msg:    govisNinOptNative,
		newMsg: newGoProtoNinOptNative,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoOptNative(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinOptNative,
		newMsg: newGoGoProtoNinOptNative,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoOptNative(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinOptNative,
		newGogomsg: newGoGoProtoNinOptNative,
		gogomsg:    gogovisNinOptNative,
	}.test(t)
}

//Optional Native
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoOptNative(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoOptNative(t *testing.T) {
	testProtoGo{
		msg:    govimNinOptNative,
		newMsg: newGoProtoNinOptNative,
	}.test(t)
}

//Optional Native
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoOptNative(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinOptNative,
		newMsg: newGoGoProtoNinOptNative,
	}.test(t)
}

//Optional Native
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoOptNative(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinOptNative,
		newGogomsg: newGoGoProtoNinOptNative,
		gogomsg:    gogovimNinOptNative,
	}.test(t)
}

//Optional Native
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoOptNative(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoRepNative(t *testing.T) {
	testProtoGo{
		msg:    govisNidRepNative,
		newMsg: newGoProtoNidRepNative,
	}.test(t)
}

//Repeated Native
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoRepNative(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidRepNative,
		newMsg: newGoGoProtoNidRepNative,
	}.test(t)
}

//Repeated Native
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoRepNative(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidRepNative,
		newGogomsg: newGoGoProtoNidRepNative,
		gogomsg:    gogovisNidRepNative,
	}.test(t)
}

//Repeated Native
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoRepNative(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoRepNative(t *testing.T) {
	testProtoGo{
		msg:    govimNidRepNative,
		newMsg: newGoProtoNidRepNative,
	}.test(t)
}

//Repeated Native
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoRepNative(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidRepNative,
		newMsg: newGoGoProtoNidRepNative,
	}.test(t)
}

//Repeated Native
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoRepNative(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidRepNative,
		newGogomsg: newGoGoProtoNidRepNative,
		gogomsg:    gogovimNidRepNative,
	}.test(t)
}

//Repeated Native
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoRepNative(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoRepNative(t *testing.T) {
	testProtoGo{
		msg:    govisNinRepNative,
		newMsg: newGoProtoNinRepNative,
	}.test(t)
}

//Repeated Native
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoRepNative(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinRepNative,
		newMsg: newGoGoProtoNinRepNative,
	}.test(t)
}

//Repeated Native
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoRepNative(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinRepNative,
		newGogomsg: newGoGoProtoNinRepNative,
		gogomsg:    gogovisNinRepNative,
	}.test(t)
}

//Repeated Native
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoRepNative(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoRepNative(t *testing.T) {
	testProtoGo{
		msg:    govimNinRepNative,
		newMsg: newGoProtoNinRepNative,
	}.test(t)
}

//Repeated Native
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoRepNative(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinRepNative,
		newMsg: newGoGoProtoNinRepNative,
	}.test(t)
}

//Repeated Native
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoRepNative(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinRepNative,
		newGogomsg: newGoGoProtoNinRepNative,
		gogomsg:    gogovimNinRepNative,
	}.test(t)
}

//Repeated Native
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoRepNative(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoRepPackedNative(t *testing.T) {
	testProtoGo{
		msg:    govisNidRepPackedNative,
		newMsg: newGoProtoNidRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoRepPackedNative(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidRepPackedNative,
		newMsg: newGoGoProtoNidRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoRepPackedNative(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidRepPackedNative,
		newGogomsg: newGoGoProtoNidRepPackedNative,
		gogomsg:    gogovisNidRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoRepPackedNative(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoRepPackedNative(t *testing.T) {
	testProtoGo{
		msg:    govimNidRepPackedNative,
		newMsg: newGoProtoNidRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoRepPackedNative(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidRepPackedNative,
		newMsg: newGoGoProtoNidRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoRepPackedNative(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidRepPackedNative,
		newGogomsg: newGoGoProtoNidRepPackedNative,
		gogomsg:    gogovimNidRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoRepPackedNative(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoRepPackedNative(t *testing.T) {
	testProtoGo{
		msg:    govisNinRepPackedNative,
		newMsg: newGoProtoNinRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoRepPackedNative(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinRepPackedNative,
		newMsg: newGoGoProtoNinRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoRepPackedNative(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinRepPackedNative,
		newGogomsg: newGoGoProtoNinRepPackedNative,
		gogomsg:    gogovisNinRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoRepPackedNative(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoRepPackedNative(t *testing.T) {
	testProtoGo{
		msg:    govimNinRepPackedNative,
		newMsg: newGoProtoNinRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoRepPackedNative(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinRepPackedNative,
		newMsg: newGoGoProtoNinRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoRepPackedNative(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinRepPackedNative,
		newGogomsg: newGoGoProtoNinRepPackedNative,
		gogomsg:    gogovimNinRepPackedNative,
	}.test(t)
}

//Packed Repeated Native
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoRepPackedNative(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoOptStruct(t *testing.T) {
	testProtoGo{
		msg:    govisNidOptStruct,
		newMsg: newGoProtoNidOptStruct,
	}.test(t)
}

//Optional Struct
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoOptStruct(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidOptStruct,
		newMsg: newGoGoProtoNidOptStruct,
	}.test(t)
}

//Optional Struct
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoOptStruct(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidOptStruct,
		newGogomsg: newGoGoProtoNidOptStruct,
		gogomsg:    gogovisNidOptStruct,
	}.test(t)
}

//Optional Struct
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoOptStruct(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoOptStruct(t *testing.T) {
	testProtoGo{
		msg:    govimNidOptStruct,
		newMsg: newGoProtoNidOptStruct,
	}.test(t)
}

//Optional Struct
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoOptStruct(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidOptStruct,
		newMsg: newGoGoProtoNidOptStruct,
	}.test(t)
}

//Optional Struct
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoOptStruct(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      gozeroNidOptStruct,
		newGogomsg: newGoGoProtoNidOptStruct,
		gogomsg:    gogovimNidOptStruct,
	}.test(t)
}

//Optional Struct
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoOptStruct(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoOptStruct(t *testing.T) {
	testProtoGo{
		msg:    govisNinOptStruct,
		newMsg: newGoProtoNinOptStruct,
	}.test(t)
}

//Optional Struct
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoOptStruct(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinOptStruct,
		newMsg: newGoGoProtoNinOptStruct,
	}.test(t)
}

//Optional Struct
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoOptStruct(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinOptStruct,
		newGogomsg: newGoGoProtoNinOptStruct,
		gogomsg:    gogovisNinOptStruct,
	}.test(t)
}

//Optional Struct
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoOptStruct(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoOptStruct(t *testing.T) {
	testProtoGo{
		msg:    govimNinOptStruct,
		newMsg: newGoProtoNinOptStruct,
	}.test(t)
}

//Optional Struct
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoOptStruct(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinOptStruct,
		newMsg: newGoGoProtoNinOptStruct,
	}.test(t)
}

//Optional Struct
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoOptStruct(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinOptStruct,
		newGogomsg: newGoGoProtoNinOptStruct,
		gogomsg:    gogovimNinOptStruct,
	}.test(t)
}

//Optional Struct
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoOptStruct(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoRepStruct(t *testing.T) {
	testProtoGo{
		msg:    govisNidRepStruct,
		newMsg: newGoProtoNidRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoRepStruct(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidRepStruct,
		newMsg: newGoGoProtoNidRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoRepStruct(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidRepStruct,
		newGogomsg: newGoGoProtoNidRepStruct,
		gogomsg:    gogovisNidRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoRepStruct(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoRepStruct(t *testing.T) {
	testProtoGo{
		msg:    govimNidRepStruct,
		newMsg: newGoProtoNidRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoRepStruct(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidRepStruct,
		newMsg: newGoGoProtoNidRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoRepStruct(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidRepStruct,
		newGogomsg: newGoGoProtoNidRepStruct,
		gogomsg:    gogovimNidRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoRepStruct(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoRepStruct(t *testing.T) {
	testProtoGo{
		msg:    govisNinRepStruct,
		newMsg: newGoProtoNinRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoRepStruct(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinRepStruct,
		newMsg: newGoGoProtoNinRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoRepStruct(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinRepStruct,
		newGogomsg: newGoGoProtoNinRepStruct,
		gogomsg:    gogovisNinRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoRepStruct(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoRepStruct(t *testing.T) {
	testProtoGo{
		msg:    govimNinRepStruct,
		newMsg: newGoProtoNinRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoRepStruct(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinRepStruct,
		newMsg: newGoGoProtoNinRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoRepStruct(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinRepStruct,
		newGogomsg: newGoGoProtoNinRepStruct,
		gogomsg:    gogovimNinRepStruct,
	}.test(t)
}

//Repeated Struct
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoRepStruct(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoEmbeddedStruct(t *testing.T) {
	testProtoGo{
		msg:    govisNidEmbeddedStruct,
		newMsg: newGoProtoNidEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoEmbeddedStruct(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidEmbeddedStruct,
		newMsg: newGoGoProtoNidEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoEmbeddedStruct(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidEmbeddedStruct,
		newGogomsg: newGoGoProtoNidEmbeddedStruct,
		gogomsg:    gogovisNidEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoEmbeddedStruct(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoEmbeddedStruct(t *testing.T) {
	testProtoGo{
		msg:    govimNidEmbeddedStruct,
		newMsg: newGoProtoNidEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoEmbeddedStruct(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidEmbeddedStruct,
		newMsg: newGoGoProtoNidEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoEmbeddedStruct(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      gozeroNidEmbeddedStruct,
		newGogomsg: newGoGoProtoNidEmbeddedStruct,
		gogomsg:    gogovimNidEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoEmbeddedStruct(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoEmbeddedStruct(t *testing.T) {
	testProtoGo{
		msg:    govisNinEmbeddedStruct,
		newMsg: newGoProtoNinEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoEmbeddedStruct(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinEmbeddedStruct,
		newMsg: newGoGoProtoNinEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoEmbeddedStruct(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinEmbeddedStruct,
		newGogomsg: newGoGoProtoNinEmbeddedStruct,
		gogomsg:    gogovisNinEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoEmbeddedStruct(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoEmbeddedStruct(t *testing.T) {
	testProtoGo{
		msg:    govimNinEmbeddedStruct,
		newMsg: newGoProtoNinEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoEmbeddedStruct(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinEmbeddedStruct,
		newMsg: newGoGoProtoNinEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoEmbeddedStruct(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinEmbeddedStruct,
		newGogomsg: newGoGoProtoNinEmbeddedStruct,
		gogomsg:    gogovimNinEmbeddedStruct,
	}.test(t)
}

//Embedded Struct
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoEmbeddedStruct(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoOptUuid(t *testing.T) {
	testProtoGo{
		msg:    govisNidOptUuid,
		newMsg: newGoProtoNidOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoOptUuid(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidOptUuid,
		newMsg: newGoGoProtoNidOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoOptUuid(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidOptUuid,
		newGogomsg: newGoGoProtoNidOptUuid,
		gogomsg:    gogovisNidOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoOptUuid(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoOptUuid(t *testing.T) {
	testProtoGo{
		msg:    govimNidOptUuid,
		newMsg: newGoProtoNidOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoOptUuid(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidOptUuid,
		newMsg: newGoGoProtoNidOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoOptUuid(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidOptUuid,
		newGogomsg: newGoGoProtoNidOptUuid,
		gogomsg:    gogovimNidOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoOptUuid(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoOptUuid(t *testing.T) {
	testProtoGo{
		msg:    govisNinOptUuid,
		newMsg: newGoProtoNinOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoOptUuid(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinOptUuid,
		newMsg: newGoGoProtoNinOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoOptUuid(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinOptUuid,
		newGogomsg: newGoGoProtoNinOptUuid,
		gogomsg:    gogovisNinOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoOptUuid(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoOptUuid(t *testing.T) {
	testProtoGo{
		msg:    govimNinOptUuid,
		newMsg: newGoProtoNinOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoOptUuid(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinOptUuid,
		newMsg: newGoGoProtoNinOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoOptUuid(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinOptUuid,
		newGogomsg: newGoGoProtoNinOptUuid,
		gogomsg:    gogovimNinOptUuid,
	}.test(t)
}

//Optional Uuid
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoOptUuid(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoRepUuid(t *testing.T) {
	testProtoGo{
		msg:    govisNidRepUuid,
		newMsg: newGoProtoNidRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoRepUuid(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidRepUuid,
		newMsg: newGoGoProtoNidRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoRepUuid(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidRepUuid,
		newGogomsg: newGoGoProtoNidRepUuid,
		gogomsg:    gogovisNidRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoRepUuid(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoRepUuid(t *testing.T) {
	testProtoGo{
		msg:    govimNidRepUuid,
		newMsg: newGoProtoNidRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoRepUuid(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidRepUuid,
		newMsg: newGoGoProtoNidRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoRepUuid(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidRepUuid,
		newGogomsg: newGoGoProtoNidRepUuid,
		gogomsg:    gogovimNidRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoRepUuid(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoRepUuid(t *testing.T) {
	testProtoGo{
		msg:    govisNinRepUuid,
		newMsg: newGoProtoNinRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoRepUuid(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinRepUuid,
		newMsg: newGoGoProtoNinRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoRepUuid(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinRepUuid,
		newGogomsg: newGoGoProtoNinRepUuid,
		gogomsg:    gogovisNinRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoRepUuid(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoRepUuid(t *testing.T) {
	testProtoGo{
		msg:    govimNinRepUuid,
		newMsg: newGoProtoNinRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoRepUuid(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinRepUuid,
		newMsg: newGoGoProtoNinRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoRepUuid(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinRepUuid,
		newGogomsg: newGoGoProtoNinRepUuid,
		gogomsg:    gogovimNinRepUuid,
	}.test(t)
}

//Repeated Uuid
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoRepUuid(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoNestedStruct(t *testing.T) {
	testProtoGo{
		msg:    govisNidNestedStruct,
		newMsg: newGoProtoNidNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoNestedStruct(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidNestedStruct,
		newMsg: newGoGoProtoNidNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoNestedStruct(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidNestedStruct,
		newGogomsg: newGoGoProtoNidNestedStruct,
		gogomsg:    gogovisNidNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoNestedStruct(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoNestedStruct(t *testing.T) {
	testProtoGo{
		msg:    govimNidNestedStruct,
		newMsg: newGoProtoNidNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoNestedStruct(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidNestedStruct,
		newMsg: newGoGoProtoNidNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoNestedStruct(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      gozeroNidNestedStruct,
		newGogomsg: newGoGoProtoNidNestedStruct,
		gogomsg:    gogovimNidNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoNestedStruct(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoNestedStruct(t *testing.T) {
	testProtoGo{
		msg:    govisNinNestedStruct,
		newMsg: newGoProtoNinNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoNestedStruct(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinNestedStruct,
		newMsg: newGoGoProtoNinNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoNestedStruct(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinNestedStruct,
		newGogomsg: newGoGoProtoNinNestedStruct,
		gogomsg:    gogovisNinNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoNestedStruct(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoNestedStruct(t *testing.T) {
	testProtoGo{
		msg:    govimNinNestedStruct,
		newMsg: newGoProtoNinNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoNestedStruct(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinNestedStruct,
		newMsg: newGoGoProtoNinNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoNestedStruct(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinNestedStruct,
		newGogomsg: newGoGoProtoNinNestedStruct,
		gogomsg:    gogovimNinNestedStruct,
	}.test(t)
}

//Nested Struct
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoNestedStruct(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoOptEnum(t *testing.T) {
	testProtoGo{
		msg:    govisNidOptEnum,
		newMsg: newGoProtoNidOptEnum,
	}.test(t)
}

//Optional Enum
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoOptEnum(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidOptEnum,
		newMsg: newGoGoProtoNidOptEnum,
	}.test(t)
}

//Optional Enum
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoOptEnum(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidOptEnum,
		newGogomsg: newGoGoProtoNidOptEnum,
		gogomsg:    gogovisNidOptEnum,
	}.test(t)
}

//Optional Enum
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoOptEnum(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoOptEnum(t *testing.T) {
	testProtoGo{
		msg:    govimNidOptEnum,
		newMsg: newGoProtoNidOptEnum,
	}.test(t)
}

//Optional Enum
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoOptEnum(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidOptEnum,
		newMsg: newGoGoProtoNidOptEnum,
	}.test(t)
}

//Optional Enum
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoOptEnum(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      gozeroNidOptEnum,
		newGogomsg: newGoGoProtoNidOptEnum,
		gogomsg:    gogovimNidOptEnum,
	}.test(t)
}

//Optional Enum
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoOptEnum(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoOptEnum(t *testing.T) {
	testProtoGo{
		msg:    govisNinOptEnum,
		newMsg: newGoProtoNinOptEnum,
	}.test(t)
}

//Optional Enum
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoOptEnum(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinOptEnum,
		newMsg: newGoGoProtoNinOptEnum,
	}.test(t)
}

//Optional Enum
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoOptEnum(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinOptEnum,
		newGogomsg: newGoGoProtoNinOptEnum,
		gogomsg:    gogovisNinOptEnum,
	}.test(t)
}

//Optional Enum
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoOptEnum(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoOptEnum(t *testing.T) {
	testProtoGo{
		msg:    govimNinOptEnum,
		newMsg: newGoProtoNinOptEnum,
	}.test(t)
}

//Optional Enum
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoOptEnum(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinOptEnum,
		newMsg: newGoGoProtoNinOptEnum,
	}.test(t)
}

//Optional Enum
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoOptEnum(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinOptEnum,
		newGogomsg: newGoGoProtoNinOptEnum,
		gogomsg:    gogovimNinOptEnum,
	}.test(t)
}

//Optional Enum
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoOptEnum(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoRepEnum(t *testing.T) {
	testProtoGo{
		msg:    govisNidRepEnum,
		newMsg: newGoProtoNidRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoRepEnum(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidRepEnum,
		newMsg: newGoGoProtoNidRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoRepEnum(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidRepEnum,
		newGogomsg: newGoGoProtoNidRepEnum,
		gogomsg:    gogovisNidRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoRepEnum(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoRepEnum(t *testing.T) {
	testProtoGo{
		msg:    govimNidRepEnum,
		newMsg: newGoProtoNidRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoRepEnum(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidRepEnum,
		newMsg: newGoGoProtoNidRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoRepEnum(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidRepEnum,
		newGogomsg: newGoGoProtoNidRepEnum,
		gogomsg:    gogovimNidRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoRepEnum(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoRepEnum(t *testing.T) {
	testProtoGo{
		msg:    govisNinRepEnum,
		newMsg: newGoProtoNinRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoRepEnum(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinRepEnum,
		newMsg: newGoGoProtoNinRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoRepEnum(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinRepEnum,
		newGogomsg: newGoGoProtoNinRepEnum,
		gogomsg:    gogovisNinRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoRepEnum(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoRepEnum(t *testing.T) {
	testProtoGo{
		msg:    govimNinRepEnum,
		newMsg: newGoProtoNinRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoRepEnum(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinRepEnum,
		newMsg: newGoGoProtoNinRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoRepEnum(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinRepEnum,
		newGogomsg: newGoGoProtoNinRepEnum,
		gogomsg:    gogovimNinRepEnum,
	}.test(t)
}

//Repeated Enum
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoRepEnum(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoOptUint128(t *testing.T) {
	testProtoGo{
		msg:    govisNidOptUint128,
		newMsg: newGoProtoNidOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoOptUint128(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidOptUint128,
		newMsg: newGoGoProtoNidOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoOptUint128(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidOptUint128,
		newGogomsg: newGoGoProtoNidOptUint128,
		gogomsg:    gogovisNidOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoOptUint128(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoOptUint128(t *testing.T) {
	testProtoGo{
		msg:    govimNidOptUint128,
		newMsg: newGoProtoNidOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoOptUint128(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidOptUint128,
		newMsg: newGoGoProtoNidOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoOptUint128(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidOptUint128,
		newGogomsg: newGoGoProtoNidOptUint128,
		gogomsg:    gogovimNidOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoOptUint128(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoOptUint128(t *testing.T) {
	testProtoGo{
		msg:    govisNinOptUint128,
		newMsg: newGoProtoNinOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoOptUint128(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinOptUint128,
		newMsg: newGoGoProtoNinOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoOptUint128(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinOptUint128,
		newGogomsg: newGoGoProtoNinOptUint128,
		gogomsg:    gogovisNinOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoOptUint128(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoOptUint128(t *testing.T) {
	testProtoGo{
		msg:    govimNinOptUint128,
		newMsg: newGoProtoNinOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoOptUint128(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinOptUint128,
		newMsg: newGoGoProtoNinOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoOptUint128(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinOptUint128,
		newGogomsg: newGoGoProtoNinOptUint128,
		gogomsg:    gogovimNinOptUint128,
	}.test(t)
}

//Optional Uint128
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoOptUint128(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoRepUint128(t *testing.T) {
	testProtoGo{
		msg:    govisNidRepUint128,
		newMsg: newGoProtoNidRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoRepUint128(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidRepUint128,
		newMsg: newGoGoProtoNidRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoRepUint128(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidRepUint128,
		newGogomsg: newGoGoProtoNidRepUint128,
		gogomsg:    gogovisNidRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoRepUint128(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoRepUint128(t *testing.T) {
	testProtoGo{
		msg:    govimNidRepUint128,
		newMsg: newGoProtoNidRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoRepUint128(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidRepUint128,
		newMsg: newGoGoProtoNidRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoRepUint128(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidRepUint128,
		newGogomsg: newGoGoProtoNidRepUint128,
		gogomsg:    gogovimNidRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoRepUint128(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoRepUint128(t *testing.T) {
	testProtoGo{
		msg:    govisNinRepUint128,
		newMsg: newGoProtoNinRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoRepUint128(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinRepUint128,
		newMsg: newGoGoProtoNinRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoRepUint128(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinRepUint128,
		newGogomsg: newGoGoProtoNinRepUint128,
		gogomsg:    gogovisNinRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoRepUint128(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoRepUint128(t *testing.T) {
	testProtoGo{
		msg:    govimNinRepUint128,
		newMsg: newGoProtoNinRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoRepUint128(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinRepUint128,
		newMsg: newGoGoProtoNinRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoRepUint128(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinRepUint128,
		newGogomsg: newGoGoProtoNinRepUint128,
		gogomsg:    gogovimNinRepUint128,
	}.test(t)
}

//Repeated Uint128
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoRepUint128(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoOptUint16(t *testing.T) {
	testProtoGo{
		msg:    govisNidOptUint16,
		newMsg: newGoProtoNidOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoOptUint16(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidOptUint16,
		newMsg: newGoGoProtoNidOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoOptUint16(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidOptUint16,
		newGogomsg: newGoGoProtoNidOptUint16,
		gogomsg:    gogovisNidOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoOptUint16(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoOptUint16(t *testing.T) {
	testProtoGo{
		msg:    govimNidOptUint16,
		newMsg: newGoProtoNidOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoOptUint16(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidOptUint16,
		newMsg: newGoGoProtoNidOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoOptUint16(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidOptUint16,
		newGogomsg: newGoGoProtoNidOptUint16,
		gogomsg:    gogovimNidOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoOptUint16(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoOptUint16(t *testing.T) {
	testProtoGo{
		msg:    govisNinOptUint16,
		newMsg: newGoProtoNinOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoOptUint16(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinOptUint16,
		newMsg: newGoGoProtoNinOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoOptUint16(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinOptUint16,
		newGogomsg: newGoGoProtoNinOptUint16,
		gogomsg:    gogovisNinOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoOptUint16(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoOptUint16(t *testing.T) {
	testProtoGo{
		msg:    govimNinOptUint16,
		newMsg: newGoProtoNinOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoOptUint16(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinOptUint16,
		newMsg: newGoGoProtoNinOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoOptUint16(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinOptUint16,
		newGogomsg: newGoGoProtoNinOptUint16,
		gogomsg:    gogovimNinOptUint16,
	}.test(t)
}

//Optional Uint16
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoOptUint16(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoRepUint16(t *testing.T) {
	testProtoGo{
		msg:    govisNidRepUint16,
		newMsg: newGoProtoNidRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoRepUint16(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidRepUint16,
		newMsg: newGoGoProtoNidRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoRepUint16(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidRepUint16,
		newGogomsg: newGoGoProtoNidRepUint16,
		gogomsg:    gogovisNidRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoRepUint16(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoRepUint16(t *testing.T) {
	testProtoGo{
		msg:    govimNidRepUint16,
		newMsg: newGoProtoNidRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoRepUint16(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidRepUint16,
		newMsg: newGoGoProtoNidRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoRepUint16(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidRepUint16,
		newGogomsg: newGoGoProtoNidRepUint16,
		gogomsg:    gogovimNidRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoRepUint16(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoRepUint16(t *testing.T) {
	testProtoGo{
		msg:    govisNinRepUint16,
		newMsg: newGoProtoNinRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoRepUint16(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinRepUint16,
		newMsg: newGoGoProtoNinRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoRepUint16(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinRepUint16,
		newGogomsg: newGoGoProtoNinRepUint16,
		gogomsg:    gogovisNinRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoRepUint16(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoRepUint16(t *testing.T) {
	testProtoGo{
		msg:    govimNinRepUint16,
		newMsg: newGoProtoNinRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoRepUint16(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinRepUint16,
		newMsg: newGoGoProtoNinRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoRepUint16(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinRepUint16,
		newGogomsg: newGoGoProtoNinRepUint16,
		gogomsg:    gogovimNinRepUint16,
	}.test(t)
}

//Repeated Uint16
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoRepUint16(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoOptUuidAsBytes(t *testing.T) {
	testProtoGo{
		msg:    govisNidOptUuidAsBytes,
		newMsg: newGoProtoNidOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoOptUuidAsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidOptUuidAsBytes,
		newMsg: newGoGoProtoNidOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoOptUuidAsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidOptUuidAsBytes,
		newGogomsg: newGoGoProtoNidOptUuidAsBytes,
		gogomsg:    gogovisNidOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoOptUuidAsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoOptUuidAsBytes(t *testing.T) {
	testProtoGo{
		msg:    govimNidOptUuidAsBytes,
		newMsg: newGoProtoNidOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoOptUuidAsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidOptUuidAsBytes,
		newMsg: newGoGoProtoNidOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoOptUuidAsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidOptUuidAsBytes,
		newGogomsg: newGoGoProtoNidOptUuidAsBytes,
		gogomsg:    gogovimNidOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoOptUuidAsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoOptUuidAsBytes(t *testing.T) {
	testProtoGo{
		msg:    govisNinOptUuidAsBytes,
		newMsg: newGoProtoNinOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoOptUuidAsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinOptUuidAsBytes,
		newMsg: newGoGoProtoNinOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoOptUuidAsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinOptUuidAsBytes,
		newGogomsg: newGoGoProtoNinOptUuidAsBytes,
		gogomsg:    gogovisNinOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoOptUuidAsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoOptUuidAsBytes(t *testing.T) {
	testProtoGo{
		msg:    govimNinOptUuidAsBytes,
		newMsg: newGoProtoNinOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoOptUuidAsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinOptUuidAsBytes,
		newMsg: newGoGoProtoNinOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoOptUuidAsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinOptUuidAsBytes,
		newGogomsg: newGoGoProtoNinOptUuidAsBytes,
		gogomsg:    gogovimNinOptUuidAsBytes,
	}.test(t)
}

//Optional UuidAsBytes
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoOptUuidAsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoRepUuidAsBytes(t *testing.T) {
	testProtoGo{
		msg:    govisNidRepUuidAsBytes,
		newMsg: newGoProtoNidRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoRepUuidAsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidRepUuidAsBytes,
		newMsg: newGoGoProtoNidRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoRepUuidAsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidRepUuidAsBytes,
		newGogomsg: newGoGoProtoNidRepUuidAsBytes,
		gogomsg:    gogovisNidRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoRepUuidAsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoRepUuidAsBytes(t *testing.T) {
	testProtoGo{
		msg:    govimNidRepUuidAsBytes,
		newMsg: newGoProtoNidRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoRepUuidAsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidRepUuidAsBytes,
		newMsg: newGoGoProtoNidRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoRepUuidAsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidRepUuidAsBytes,
		newGogomsg: newGoGoProtoNidRepUuidAsBytes,
		gogomsg:    gogovimNidRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoRepUuidAsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoRepUuidAsBytes(t *testing.T) {
	testProtoGo{
		msg:    govisNinRepUuidAsBytes,
		newMsg: newGoProtoNinRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoRepUuidAsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinRepUuidAsBytes,
		newMsg: newGoGoProtoNinRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoRepUuidAsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinRepUuidAsBytes,
		newGogomsg: newGoGoProtoNinRepUuidAsBytes,
		gogomsg:    gogovisNinRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoRepUuidAsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoRepUuidAsBytes(t *testing.T) {
	testProtoGo{
		msg:    govimNinRepUuidAsBytes,
		newMsg: newGoProtoNinRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoRepUuidAsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinRepUuidAsBytes,
		newMsg: newGoGoProtoNinRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoRepUuidAsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinRepUuidAsBytes,
		newGogomsg: newGoGoProtoNinRepUuidAsBytes,
		gogomsg:    gogovimNinRepUuidAsBytes,
	}.test(t)
}

//Repeated UuidAsBytes
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoRepUuidAsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoOptUint128AsBytes(t *testing.T) {
	testProtoGo{
		msg:    govisNidOptUint128AsBytes,
		newMsg: newGoProtoNidOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoOptUint128AsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidOptUint128AsBytes,
		newMsg: newGoGoProtoNidOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoOptUint128AsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidOptUint128AsBytes,
		newGogomsg: newGoGoProtoNidOptUint128AsBytes,
		gogomsg:    gogovisNidOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoOptUint128AsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoOptUint128AsBytes(t *testing.T) {
	testProtoGo{
		msg:    govimNidOptUint128AsBytes,
		newMsg: newGoProtoNidOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoOptUint128AsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidOptUint128AsBytes,
		newMsg: newGoGoProtoNidOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoOptUint128AsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidOptUint128AsBytes,
		newGogomsg: newGoGoProtoNidOptUint128AsBytes,
		gogomsg:    gogovimNidOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoOptUint128AsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoOptUint128AsBytes(t *testing.T) {
	testProtoGo{
		msg:    govisNinOptUint128AsBytes,
		newMsg: newGoProtoNinOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoOptUint128AsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinOptUint128AsBytes,
		newMsg: newGoGoProtoNinOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoOptUint128AsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinOptUint128AsBytes,
		newGogomsg: newGoGoProtoNinOptUint128AsBytes,
		gogomsg:    gogovisNinOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoOptUint128AsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoOptUint128AsBytes(t *testing.T) {
	testProtoGo{
		msg:    govimNinOptUint128AsBytes,
		newMsg: newGoProtoNinOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoOptUint128AsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinOptUint128AsBytes,
		newMsg: newGoGoProtoNinOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoOptUint128AsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinOptUint128AsBytes,
		newGogomsg: newGoGoProtoNinOptUint128AsBytes,
		gogomsg:    gogovimNinOptUint128AsBytes,
	}.test(t)
}

//Optional Uint128AsBytes
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoOptUint128AsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoRepUint128AsBytes(t *testing.T) {
	testProtoGo{
		msg:    govisNidRepUint128AsBytes,
		newMsg: newGoProtoNidRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoRepUint128AsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidRepUint128AsBytes,
		newMsg: newGoGoProtoNidRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoRepUint128AsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidRepUint128AsBytes,
		newGogomsg: newGoGoProtoNidRepUint128AsBytes,
		gogomsg:    gogovisNidRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoRepUint128AsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoRepUint128AsBytes(t *testing.T) {
	testProtoGo{
		msg:    govimNidRepUint128AsBytes,
		newMsg: newGoProtoNidRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoRepUint128AsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidRepUint128AsBytes,
		newMsg: newGoGoProtoNidRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoRepUint128AsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidRepUint128AsBytes,
		newGogomsg: newGoGoProtoNidRepUint128AsBytes,
		gogomsg:    gogovimNidRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoRepUint128AsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoRepUint128AsBytes(t *testing.T) {
	testProtoGo{
		msg:    govisNinRepUint128AsBytes,
		newMsg: newGoProtoNinRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoRepUint128AsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinRepUint128AsBytes,
		newMsg: newGoGoProtoNinRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoRepUint128AsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinRepUint128AsBytes,
		newGogomsg: newGoGoProtoNinRepUint128AsBytes,
		gogomsg:    gogovisNinRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoRepUint128AsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoRepUint128AsBytes(t *testing.T) {
	testProtoGo{
		msg:    govimNinRepUint128AsBytes,
		newMsg: newGoProtoNinRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoRepUint128AsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinRepUint128AsBytes,
		newMsg: newGoGoProtoNinRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoRepUint128AsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinRepUint128AsBytes,
		newGogomsg: newGoGoProtoNinRepUint128AsBytes,
		gogomsg:    gogovimNinRepUint128AsBytes,
	}.test(t)
}

//Repeated Uint128AsBytes
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoRepUint128AsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoOptUint16AsBytes(t *testing.T) {
	testProtoGo{
		msg:    govisNidOptUint16AsBytes,
		newMsg: newGoProtoNidOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoOptUint16AsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidOptUint16AsBytes,
		newMsg: newGoGoProtoNidOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoOptUint16AsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidOptUint16AsBytes,
		newGogomsg: newGoGoProtoNidOptUint16AsBytes,
		gogomsg:    gogovisNidOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoOptUint16AsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoOptUint16AsBytes(t *testing.T) {
	testProtoGo{
		msg:    govimNidOptUint16AsBytes,
		newMsg: newGoProtoNidOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoOptUint16AsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidOptUint16AsBytes,
		newMsg: newGoGoProtoNidOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoOptUint16AsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidOptUint16AsBytes,
		newGogomsg: newGoGoProtoNidOptUint16AsBytes,
		gogomsg:    gogovimNidOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoOptUint16AsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoOptUint16AsBytes(t *testing.T) {
	testProtoGo{
		msg:    govisNinOptUint16AsBytes,
		newMsg: newGoProtoNinOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoOptUint16AsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinOptUint16AsBytes,
		newMsg: newGoGoProtoNinOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoOptUint16AsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinOptUint16AsBytes,
		newGogomsg: newGoGoProtoNinOptUint16AsBytes,
		gogomsg:    gogovisNinOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoOptUint16AsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoOptUint16AsBytes(t *testing.T) {
	testProtoGo{
		msg:    govimNinOptUint16AsBytes,
		newMsg: newGoProtoNinOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoOptUint16AsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinOptUint16AsBytes,
		newMsg: newGoGoProtoNinOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoOptUint16AsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinOptUint16AsBytes,
		newGogomsg: newGoGoProtoNinOptUint16AsBytes,
		gogomsg:    gogovimNinOptUint16AsBytes,
	}.test(t)
}

//Optional Uint16AsBytes
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoOptUint16AsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoRepUint16AsBytes(t *testing.T) {
	testProtoGo{
		msg:    govisNidRepUint16AsBytes,
		newMsg: newGoProtoNidRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoRepUint16AsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidRepUint16AsBytes,
		newMsg: newGoGoProtoNidRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoRepUint16AsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidRepUint16AsBytes,
		newGogomsg: newGoGoProtoNidRepUint16AsBytes,
		gogomsg:    gogovisNidRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoRepUint16AsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoRepUint16AsBytes(t *testing.T) {
	testProtoGo{
		msg:    govimNidRepUint16AsBytes,
		newMsg: newGoProtoNidRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoRepUint16AsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidRepUint16AsBytes,
		newMsg: newGoGoProtoNidRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoRepUint16AsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidRepUint16AsBytes,
		newGogomsg: newGoGoProtoNidRepUint16AsBytes,
		gogomsg:    gogovimNidRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoRepUint16AsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoRepUint16AsBytes(t *testing.T) {
	testProtoGo{
		msg:    govisNinRepUint16AsBytes,
		newMsg: newGoProtoNinRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoRepUint16AsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinRepUint16AsBytes,
		newMsg: newGoGoProtoNinRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoRepUint16AsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinRepUint16AsBytes,
		newGogomsg: newGoGoProtoNinRepUint16AsBytes,
		gogomsg:    gogovisNinRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoRepUint16AsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoRepUint16AsBytes(t *testing.T) {
	testProtoGo{
		msg:    govimNinRepUint16AsBytes,
		newMsg: newGoProtoNinRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoRepUint16AsBytes(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinRepUint16AsBytes,
		newMsg: newGoGoProtoNinRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoRepUint16AsBytes(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinRepUint16AsBytes,
		newGogomsg: newGoGoProtoNinRepUint16AsBytes,
		gogomsg:    gogovimNinRepUint16AsBytes,
	}.test(t)
}

//Repeated Uint16AsBytes
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoRepUint16AsBytes(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoOptUint16AsUint32(t *testing.T) {
	testProtoGo{
		msg:    govisNidOptUint16AsUint32,
		newMsg: newGoProtoNidOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoOptUint16AsUint32(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidOptUint16AsUint32,
		newMsg: newGoGoProtoNidOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoOptUint16AsUint32(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidOptUint16AsUint32,
		newGogomsg: newGoGoProtoNidOptUint16AsUint32,
		gogomsg:    gogovisNidOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoOptUint16AsUint32(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoOptUint16AsUint32(t *testing.T) {
	testProtoGo{
		msg:    govimNidOptUint16AsUint32,
		newMsg: newGoProtoNidOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoOptUint16AsUint32(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidOptUint16AsUint32,
		newMsg: newGoGoProtoNidOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoOptUint16AsUint32(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidOptUint16AsUint32,
		newGogomsg: newGoGoProtoNidOptUint16AsUint32,
		gogomsg:    gogovimNidOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoOptUint16AsUint32(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoOptUint16AsUint32(t *testing.T) {
	testProtoGo{
		msg:    govisNinOptUint16AsUint32,
		newMsg: newGoProtoNinOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoOptUint16AsUint32(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinOptUint16AsUint32,
		newMsg: newGoGoProtoNinOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoOptUint16AsUint32(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinOptUint16AsUint32,
		newGogomsg: newGoGoProtoNinOptUint16AsUint32,
		gogomsg:    gogovisNinOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoOptUint16AsUint32(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoOptUint16AsUint32(t *testing.T) {
	testProtoGo{
		msg:    govimNinOptUint16AsUint32,
		newMsg: newGoProtoNinOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoOptUint16AsUint32(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinOptUint16AsUint32,
		newMsg: newGoGoProtoNinOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoOptUint16AsUint32(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinOptUint16AsUint32,
		newGogomsg: newGoGoProtoNinOptUint16AsUint32,
		gogomsg:    gogovimNinOptUint16AsUint32,
	}.test(t)
}

//Optional Uint16AsUint32
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoOptUint16AsUint32(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoRepUint16AsUint32(t *testing.T) {
	testProtoGo{
		msg:    govisNidRepUint16AsUint32,
		newMsg: newGoProtoNidRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoRepUint16AsUint32(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidRepUint16AsUint32,
		newMsg: newGoGoProtoNidRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoRepUint16AsUint32(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidRepUint16AsUint32,
		newGogomsg: newGoGoProtoNidRepUint16AsUint32,
		gogomsg:    gogovisNidRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoRepUint16AsUint32(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoRepUint16AsUint32(t *testing.T) {
	testProtoGo{
		msg:    govimNidRepUint16AsUint32,
		newMsg: newGoProtoNidRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoRepUint16AsUint32(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidRepUint16AsUint32,
		newMsg: newGoGoProtoNidRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoRepUint16AsUint32(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidRepUint16AsUint32,
		newGogomsg: newGoGoProtoNidRepUint16AsUint32,
		gogomsg:    gogovimNidRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoRepUint16AsUint32(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoRepUint16AsUint32(t *testing.T) {
	testProtoGo{
		msg:    govisNinRepUint16AsUint32,
		newMsg: newGoProtoNinRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoRepUint16AsUint32(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinRepUint16AsUint32,
		newMsg: newGoGoProtoNinRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoRepUint16AsUint32(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinRepUint16AsUint32,
		newGogomsg: newGoGoProtoNinRepUint16AsUint32,
		gogomsg:    gogovisNinRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoRepUint16AsUint32(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoRepUint16AsUint32(t *testing.T) {
	testProtoGo{
		msg:    govimNinRepUint16AsUint32,
		newMsg: newGoProtoNinRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoRepUint16AsUint32(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinRepUint16AsUint32,
		newMsg: newGoGoProtoNinRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoRepUint16AsUint32(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinRepUint16AsUint32,
		newGogomsg: newGoGoProtoNinRepUint16AsUint32,
		gogomsg:    gogovimNinRepUint16AsUint32,
	}.test(t)
}

//Repeated Uint16AsUint32
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoRepUint16AsUint32(t *testing.T) {
	testProtoGoGoToGo{
		gogomsg:  gogovimNinRepUint16AsUint32,
		newGomsg: newGoProtoNinRepUint16AsUint32,
		gomsg:    govimNinRepUint16AsUint32,
		vis:      false,
	}.test(t)
}

//Optional EnumUint32
//Values are available
//Nullable is false
//Go -> Go
func TestProtoVisNidGoOptEnumUint32(t *testing.T) {
	testProtoGo{
		msg:    govisNidOptEnumUint32,
		newMsg: newGoProtoNidOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoOptEnumUint32(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidOptEnumUint32,
		newMsg: newGoGoProtoNidOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoOptEnumUint32(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidOptEnumUint32,
		newGogomsg: newGoGoProtoNidOptEnumUint32,
		gogomsg:    gogovisNidOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoOptEnumUint32(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoOptEnumUint32(t *testing.T) {
	testProtoGo{
		msg:    govimNidOptEnumUint32,
		newMsg: newGoProtoNidOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoOptEnumUint32(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidOptEnumUint32,
		newMsg: newGoGoProtoNidOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoOptEnumUint32(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      gozeroNidOptEnumUint32,
		newGogomsg: newGoGoProtoNidOptEnumUint32,
		gogomsg:    gogovimNidOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoOptEnumUint32(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoOptEnumUint32(t *testing.T) {
	testProtoGo{
		msg:    govisNinOptEnumUint32,
		newMsg: newGoProtoNinOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoOptEnumUint32(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinOptEnumUint32,
		newMsg: newGoGoProtoNinOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoOptEnumUint32(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinOptEnumUint32,
		newGogomsg: newGoGoProtoNinOptEnumUint32,
		gogomsg:    gogovisNinOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoOptEnumUint32(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoOptEnumUint32(t *testing.T) {
	testProtoGo{
		msg:    govimNinOptEnumUint32,
		newMsg: newGoProtoNinOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoOptEnumUint32(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinOptEnumUint32,
		newMsg: newGoGoProtoNinOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoOptEnumUint32(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinOptEnumUint32,
		newGogomsg: newGoGoProtoNinOptEnumUint32,
		gogomsg:    gogovimNinOptEnumUint32,
	}.test(t)
}

//Optional EnumUint32
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoOptEnumUint32(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoRepEnumUint32(t *testing.T) {
	testProtoGo{
		msg:    govisNidRepEnumUint32,
		newMsg: newGoProtoNidRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoRepEnumUint32(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidRepEnumUint32,
		newMsg: newGoGoProtoNidRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoRepEnumUint32(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidRepEnumUint32,
		newGogomsg: newGoGoProtoNidRepEnumUint32,
		gogomsg:    gogovisNidRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoRepEnumUint32(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoRepEnumUint32(t *testing.T) {
	testProtoGo{
		msg:    govimNidRepEnumUint32,
		newMsg: newGoProtoNidRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoRepEnumUint32(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidRepEnumUint32,
		newMsg: newGoGoProtoNidRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoRepEnumUint32(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidRepEnumUint32,
		newGogomsg: newGoGoProtoNidRepEnumUint32,
		gogomsg:    gogovimNidRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoRepEnumUint32(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoRepEnumUint32(t *testing.T) {
	testProtoGo{
		msg:    govisNinRepEnumUint32,
		newMsg: newGoProtoNinRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoRepEnumUint32(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinRepEnumUint32,
		newMsg: newGoGoProtoNinRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoRepEnumUint32(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinRepEnumUint32,
		newGogomsg: newGoGoProtoNinRepEnumUint32,
		gogomsg:    gogovisNinRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoRepEnumUint32(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoRepEnumUint32(t *testing.T) {
	testProtoGo{
		msg:    govimNinRepEnumUint32,
		newMsg: newGoProtoNinRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoRepEnumUint32(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinRepEnumUint32,
		newMsg: newGoGoProtoNinRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoRepEnumUint32(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinRepEnumUint32,
		newGogomsg: newGoGoProtoNinRepEnumUint32,
		gogomsg:    gogovimNinRepEnumUint32,
	}.test(t)
}

//Repeated EnumUint32
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoRepEnumUint32(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoOptEnumUint16(t *testing.T) {
	testProtoGo{
		msg:    govisNidOptEnumUint16,
		newMsg: newGoProtoNidOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoOptEnumUint16(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidOptEnumUint16,
		newMsg: newGoGoProtoNidOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoOptEnumUint16(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidOptEnumUint16,
		newGogomsg: newGoGoProtoNidOptEnumUint16,
		gogomsg:    gogovisNidOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoOptEnumUint16(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoOptEnumUint16(t *testing.T) {
	testProtoGo{
		msg:    govimNidOptEnumUint16,
		newMsg: newGoProtoNidOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoOptEnumUint16(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidOptEnumUint16,
		newMsg: newGoGoProtoNidOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoOptEnumUint16(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      gozeroNidOptEnumUint16,
		newGogomsg: newGoGoProtoNidOptEnumUint16,
		gogomsg:    gogovimNidOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoOptEnumUint16(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoOptEnumUint16(t *testing.T) {
	testProtoGo{
		msg:    govisNinOptEnumUint16,
		newMsg: newGoProtoNinOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoOptEnumUint16(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinOptEnumUint16,
		newMsg: newGoGoProtoNinOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoOptEnumUint16(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinOptEnumUint16,
		newGogomsg: newGoGoProtoNinOptEnumUint16,
		gogomsg:    gogovisNinOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoOptEnumUint16(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoOptEnumUint16(t *testing.T) {
	testProtoGo{
		msg:    govimNinOptEnumUint16,
		newMsg: newGoProtoNinOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoOptEnumUint16(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinOptEnumUint16,
		newMsg: newGoGoProtoNinOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoOptEnumUint16(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinOptEnumUint16,
		newGogomsg: newGoGoProtoNinOptEnumUint16,
		gogomsg:    gogovimNinOptEnumUint16,
	}.test(t)
}

//Optional EnumUint16
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoOptEnumUint16(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNidGoRepEnumUint16(t *testing.T) {
	testProtoGo{
		msg:    govisNidRepEnumUint16,
		newMsg: newGoProtoNidRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are available
//Nullable is false
//GoGo -> GoGo
func TestProtoVisNidGoGoRepEnumUint16(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNidRepEnumUint16,
		newMsg: newGoGoProtoNidRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are available
//Nullable is false
//Go -> GoGo
func TestProtoVisNidGoToGoGoRepEnumUint16(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNidRepEnumUint16,
		newGogomsg: newGoGoProtoNidRepEnumUint16,
		gogomsg:    gogovisNidRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are available
//Nullable is false
//GoGo -> Go
func TestProtoVisNidGoGoToGoRepEnumUint16(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNidGoRepEnumUint16(t *testing.T) {
	testProtoGo{
		msg:    govimNidRepEnumUint16,
		newMsg: newGoProtoNidRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are not available
//Nullable is false
//GoGo -> GoGo
func TestProtoVimNidGoGoRepEnumUint16(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNidRepEnumUint16,
		newMsg: newGoGoProtoNidRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are not available
//Nullable is false
//Go -> GoGo
func TestProtoVimNidGoToGoGoRepEnumUint16(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNidRepEnumUint16,
		newGogomsg: newGoGoProtoNidRepEnumUint16,
		gogomsg:    gogovimNidRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are not available
//Nullable is false
//GoGo -> Go
func TestProtoVimNidGoGoToGoRepEnumUint16(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVisNinGoRepEnumUint16(t *testing.T) {
	testProtoGo{
		msg:    govisNinRepEnumUint16,
		newMsg: newGoProtoNinRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are available
//Nullable is true
//GoGo -> GoGo
func TestProtoVisNinGoGoRepEnumUint16(t *testing.T) {
	testProtoGoGo{
		msg:    gogovisNinRepEnumUint16,
		newMsg: newGoGoProtoNinRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are available
//Nullable is true
//Go -> GoGo
func TestProtoVisNinGoToGoGoRepEnumUint16(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govisNinRepEnumUint16,
		newGogomsg: newGoGoProtoNinRepEnumUint16,
		gogomsg:    gogovisNinRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are available
//Nullable is true
//GoGo -> Go
func TestProtoVisNinGoGoToGoRepEnumUint16(t *testing.T) {
	testProtoGoGoToGo{
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
func TestProtoVimNinGoRepEnumUint16(t *testing.T) {
	testProtoGo{
		msg:    govimNinRepEnumUint16,
		newMsg: newGoProtoNinRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are not available
//Nullable is true
//GoGo -> GoGo
func TestProtoVimNinGoGoRepEnumUint16(t *testing.T) {
	testProtoGoGo{
		msg:    gogovimNinRepEnumUint16,
		newMsg: newGoGoProtoNinRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are not available
//Nullable is true
//Go -> GoGo
func TestProtoVimNinGoToGoGoRepEnumUint16(t *testing.T) {
	testProtoGoToGoGo{
		gomsg:      govimNinRepEnumUint16,
		newGogomsg: newGoGoProtoNinRepEnumUint16,
		gogomsg:    gogovimNinRepEnumUint16,
	}.test(t)
}

//Repeated EnumUint16
//Values are not available
//Nullable is true
//GoGo -> Go
func TestProtoVimNinGoGoToGoRepEnumUint16(t *testing.T) {
	testProtoGoGoToGo{
		gogomsg:  gogovimNinRepEnumUint16,
		newGomsg: newGoProtoNinRepEnumUint16,
		gomsg:    govimNinRepEnumUint16,
		vis:      false,
	}.test(t)
}

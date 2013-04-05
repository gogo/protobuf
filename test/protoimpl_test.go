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
	gogoproto "code.google.com/p/gogoprotobuf/proto"
	goproto "code.google.com/p/goprotobuf/proto"
	"fmt"
	"io/ioutil"
	"path"
	"reflect"
	"strings"
	"testing"
)

type testProtoGo struct {
	msg    goproto.Message
	newMsg func() goproto.Message
}

func (this testProtoGo) test(t *testing.T) {
	data, err := goproto.Marshal(this.msg)
	if err != nil {
		panic(err)
	}
	msg := this.newMsg()
	if err := goproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(msg, this.msg) {
		t.Fatalf("%#v != %#v", msg, this.msg)
	}
}

type testProtoGoGo struct {
	msg    gogoproto.Message
	newMsg func() gogoproto.Message
}

func (this testProtoGoGo) test(t *testing.T) {
	data, err := gogoproto.Marshal(this.msg)
	if err != nil {
		panic(err)
	}
	msg := this.newMsg()
	if err := gogoproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(msg, this.msg) {
		t.Fatalf("%#v != %#v", msg, this.msg)
	}
}

type testProtoGoToGoGo struct {
	gomsg      goproto.Message
	newGogomsg func() gogoproto.Message
	gogomsg    gogoproto.Message
}

func (this testProtoGoToGoGo) test(t *testing.T) {
	data, err := goproto.Marshal(this.gomsg)
	if err != nil {
		panic(err)
	}
	msg := this.newGogomsg()
	if err := gogoproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(msg, this.gogomsg) {
		t.Fatalf("%#v != %#v", msg, this.gogomsg)
	}
}

type testProtoGoGoToGo struct {
	gogomsg  gogoproto.Message
	newGomsg func() goproto.Message
	gomsg    goproto.Message
	vis      bool
}

func (this testProtoGoGoToGo) test(t *testing.T) {
	data, err := gogoproto.Marshal(this.gogomsg)
	if err != nil {
		panic(err)
	}
	var data2 []byte
	if gogoFiles {
		filename := strings.Replace(path.Clean(fmt.Sprintf("%T.dat", this.gogomsg)), "*", "", -1)
		filename = strings.Replace(filename, "test.", "", 1)
		if this.vis {
			filename = "Vis" + filename
		} else {
			filename = "Vim" + filename
		}
		filename = "GoGo" + filename
		if err := ioutil.WriteFile(filename, data, 0777); err != nil {
			panic(err)
		}
		data2, err = ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
	} else {
		data2 = data
	}
	msg := this.newGomsg()
	if err := goproto.Unmarshal(data2, msg); err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(msg, this.gomsg) {
		t.Fatalf("%#v != %#v", msg, this.gomsg)
	}
}

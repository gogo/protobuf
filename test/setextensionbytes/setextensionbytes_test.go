// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2019, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
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

package setextensionbytes

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestSetextensionBytes(t *testing.T) {
	myExtendable := &MyExtendable{}

	foo := &Foo{IntFoo: 1}
	fmt.Printf("Foo: %v\n", foo)
	err := proto.SetUnsafeExtension(myExtendable, 2, foo)
	if err != nil {
		panic(err)
	}
	gotFoo1, err := proto.GetUnsafeExtension(myExtendable, 2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("From GetExtension Foo: %v\n", gotFoo1)
	gotFoo1.(*Foo).IntFoo = 10
	fmt.Printf("Changed Foo: %v\n", gotFoo1)

	err = proto.SetUnsafeExtension(myExtendable, 2, gotFoo1)
	if err != nil {
		panic(err)
	}
	gotFoo2, err := proto.GetUnsafeExtension(myExtendable, 2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("From Second GetExtension Foo: %v\n", gotFoo1)

	if gotFoo2.(*Foo).IntFoo != 10 {
		t.Fatalf("got %d want %d", gotFoo2.(*Foo).IntFoo, 10)
	}
}

// === RUN   TestSetextensionBytes
// Foo: &Foo{IntFoo:1,XXX_unrecognized:[],}
// SetExtension NewVal: [18 2 8 1]
// SetExtension Done: &[18 2 8 1]
// GetExtension: &[18 2 8 1]
// From GetExtension Foo: &Foo{IntFoo:1,XXX_unrecognized:[],}
// Changed Foo: &Foo{IntFoo:10,XXX_unrecognized:[],}
// SetExtension NewVal: [18 2 8 10]
// SetExtension Done: &[18 2 8 1 18 2 8 10]
// GetExtension: &[18 2 8 1 18 2 8 10]
// From Second GetExtension Foo: &Foo{IntFoo:10,XXX_unrecognized:[],}
// --- FAIL: TestSetextensionBytes (0.00s)
//     setextensionbytes_test.go:68: got 1 want 10
// FAIL

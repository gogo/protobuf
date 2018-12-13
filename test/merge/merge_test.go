// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2015, The GoGo Authors. All rights reserved.
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

package merge

import (
	"testing"

	"github.com/gogo/protobuf/proto"
)

// TestCloneNonNullable tests Clone on messages that use the nullable extension.
func TestCloneNonNullable(t *testing.T) {
	orig := &MyMessageNonNull{
		Count: 42,
		Name:  "Dave",
		Pet:   []string{"bunny", "kitty", "horsey"},
		Inner: InnerMessage{
			Host:      "niles",
			Port:      9099,
			Connected: true,
		},
		Others: []OtherMessage{
			{
				Value: []byte("some bytes"),
				Inner: &InnerMessage{
					Host:      "x",
					Port:      9098,
					Connected: true,
				},
			},
		},
		WeMustGoDeeper: RequiredInnerMessageNonNull{
			LeoFinallyWonAnOscar: InnerMessage{
				Host:      "y",
				Port:      9097,
				Connected: true,
			},
		},

		RepInner: []InnerMessage{
			{
				Host:      "a",
				Port:      1024,
				Connected: true,
			},
			{
				Host:      "b",
				Port:      1025,
				Connected: false,
			},
		},
		Bikeshed: MyMessageNonNull_GREEN,
		RepBytes: [][]byte{[]byte("sham"), []byte("wow")},
	}
	m := proto.Clone(orig).(*MyMessageNonNull)
	if !proto.Equal(m, orig) {
		t.Fatalf("Clone(%v) = %v", orig, m)
	}

	// Verify it was a deep copy.
	m.Inner.Port++
	if proto.Equal(m, orig) {
		t.Error("Mutating clone changed the original")
	}
	m.Inner.Port--

	m.WeMustGoDeeper.LeoFinallyWonAnOscar.Port++
	if proto.Equal(m, orig) {
		t.Error("Mutating clone changed the original")
	}
	m.WeMustGoDeeper.LeoFinallyWonAnOscar.Port--

	m.Others[0].Inner.Port++
	if proto.Equal(m, orig) {
		t.Error("Mutating clone changed the original")
	}
	m.Others[0].Inner.Port--

	m.RepInner[0].Port++
	if proto.Equal(m, orig) {
		t.Error("Mutating clone changed the original")
	}
	m.RepInner[0].Port--

	// Byte fields and repeated fields should be copied.
	if &m.Pet[0] == &orig.Pet[0] {
		t.Error("Pet: repeated field not copied")
	}
	if &m.Others[0] == &orig.Others[0] {
		t.Error("Others: repeated field not copied")
	}
	if &m.Others[0].Value[0] == &orig.Others[0].Value[0] {
		t.Error("Others[0].Value: bytes field not copied")
	}
	if &m.RepBytes[0] == &orig.RepBytes[0] {
		t.Error("RepBytes: repeated field not copied")
	}
	if &m.RepBytes[0][0] == &orig.RepBytes[0][0] {
		t.Error("RepBytes[0]: bytes field not copied")
	}
}

// TestCloneNonNullable tests Merge on messages that use the nullable extension.
func TestMergeNotNullable(t *testing.T) {
	src := &MyMessageNonNull{
		Count: 1,
		Pet:   []string{"horsey"},
		Inner: InnerMessage{
			Host: "a",
		},
		Others: []OtherMessage{
			{
				Key: 12345,
				Inner: &InnerMessage{
					Host: "b",
				},
			},
			{
				Key: 123456,
			},
		},
		WeMustGoDeeper: RequiredInnerMessageNonNull{
			LeoFinallyWonAnOscar: InnerMessage{
				Host: "c",
			},
		},
		RepInner: []InnerMessage{
			{
				Host: "d",
				Port: 1024,
			},
			{
				Host:      "e",
				Connected: true,
			},
		},
		Bikeshed: MyMessageNonNull_BLUE,
	}

	dst := &MyMessageNonNull{
		Count: 2,
		Name:  "name",
		Pet:   []string{"bunny", "kitty"},
		Inner: InnerMessage{
			Port: 666,
		},
		Others: []OtherMessage{
			{
				Key: 12,
				Inner: &InnerMessage{
					Host: "x",
				},
			},
		},
		WeMustGoDeeper: RequiredInnerMessageNonNull{
			LeoFinallyWonAnOscar: InnerMessage{
				Port: 1023,
			},
		},
		RepInner: []InnerMessage{
			{
				Host:      "x",
				Connected: true,
			},
		},
	}
	want := &MyMessageNonNull{
		Count: 1,
		Name:  "name",
		Pet:   []string{"bunny", "kitty", "horsey"},
		Inner: InnerMessage{
			Host: "a",
			Port: 666,
		},
		Others: []OtherMessage{
			{
				Key: 12,
				Inner: &InnerMessage{
					Host: "x",
				},
			},
			{
				Key: 12345,
				Inner: &InnerMessage{
					Host: "b",
				},
			},
			{
				Key: 123456,
			},
		},
		WeMustGoDeeper: RequiredInnerMessageNonNull{
			LeoFinallyWonAnOscar: InnerMessage{
				Host: "c",
				Port: 1023,
			},
		},
		RepInner: []InnerMessage{
			{
				Host:      "x",
				Connected: true,
			},
			{
				Host: "d",
				Port: 1024,
			},
			{
				Host:      "e",
				Connected: true,
			},
		},
		Bikeshed: MyMessageNonNull_BLUE,
	}

	got := proto.Clone(dst)
	if !proto.Equal(got, dst) {
		t.Fatalf("Clone()\ngot  %v\nwant %v", got, dst)
	}
	proto.Merge(got, src)
	if !proto.Equal(got, want) {
		t.Errorf("Merge(%v, %v)\ngot  %v\nwant %v", dst, src, got, want)
	}
}

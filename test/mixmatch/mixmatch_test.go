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

package mixmatch

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
)

type MixMatch struct {
	Old []string
	New []string
}

func (this MixMatch) Regenerate() {
	fmt.Printf("mixmatch\n")
	data, err := ioutil.ReadFile("../thetest.proto")
	if err != nil {
		panic(err)
	}
	content := string(data)
	for i, old := range this.Old {
		content = strings.Replace(content, old, this.New[i], -1)
	}
	if err := ioutil.WriteFile("./testdata/thetest.proto", []byte(content), 0666); err != nil {
		panic(err)
	}
	data2, err := ioutil.ReadFile("../uuid.go")
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("./testdata/uuid.go", data2, 0666); err != nil {
		panic(err)
	}
	var regenerate = exec.Command("protoc", "--gogo_out=.", "-I=../../:../../protobuf/:../../../../../:.", "./testdata/thetest.proto")
	fmt.Printf("regenerating\n")
	out, err := regenerate.CombinedOutput()
	fmt.Printf("regenerate output: %v\n", string(out))
	if err != nil {
		panic(err)
	}
}

func (this MixMatch) Test(t *testing.T) {
	if _, err := exec.LookPath("protoc"); err != nil {
		t.Skipf("cannot find protoc in PATH")
	}
	if _, err := exec.LookPath("go"); err != nil {
		t.Skipf("cannot find go in PATH")
	}
	if err := os.MkdirAll("./testdata", 0777); err != nil {
		panic(err)
	}
	this.Regenerate()
	var test = exec.Command("go", "test", "-v", "./testdata/")
	fmt.Printf("testing\n")
	out, err := test.CombinedOutput()
	fmt.Printf("test output: %v\n", string(out))
	if err != nil {
		panic(err)
	}
	if err := os.RemoveAll("./testdata"); err != nil {
		panic(err)
	}
}

func TestNeither(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	MixMatch{
		Old: []string{
			"option (gogoproto.unmarshaler_all) = true;",
			"option (gogoproto.marshaler_all) = true;",
			"option (gogoproto.unsafe_unmarshaler_all) = true;",
			"option (gogoproto.unsafe_marshaler_all) = true;",
		},
		New: []string{
			"option (gogoproto.unmarshaler_all) = false;",
			"option (gogoproto.marshaler_all) = false;",
			"option (gogoproto.unsafe_unmarshaler_all) = false;",
			"option (gogoproto.unsafe_marshaler_all) = false;",
		},
	}.Test(t)
}

func TestMarshaler(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	MixMatch{
		Old: []string{
			"option (gogoproto.marshaler_all) = false;",
			"option (gogoproto.unmarshaler_all) = true;",
			"option (gogoproto.unsafe_unmarshaler_all) = true;",
			"option (gogoproto.unsafe_marshaler_all) = true;",
		},
		New: []string{
			"option (gogoproto.marshaler_all) = true;",
			"option (gogoproto.unmarshaler_all) = false;",
			"option (gogoproto.unsafe_unmarshaler_all) = false;",
			"option (gogoproto.unsafe_marshaler_all) = false;",
		},
	}.Test(t)
}

func TestUnmarshaler(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	MixMatch{
		Old: []string{
			"option (gogoproto.unmarshaler_all) = false;",
			"option (gogoproto.marshaler_all) = true;",
			"option (gogoproto.unsafe_unmarshaler_all) = true;",
			"option (gogoproto.unsafe_marshaler_all) = true;",
		},
		New: []string{
			"option (gogoproto.unmarshaler_all) = true;",
			"option (gogoproto.marshaler_all) = false;",
			"option (gogoproto.unsafe_unmarshaler_all) = false;",
			"option (gogoproto.unsafe_marshaler_all) = false;",
		},
	}.Test(t)
}

func TestBoth(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	MixMatch{
		Old: []string{
			"option (gogoproto.unmarshaler_all) = false;",
			"option (gogoproto.marshaler_all) = false;",
			"option (gogoproto.unsafe_unmarshaler_all) = true;",
			"option (gogoproto.unsafe_marshaler_all) = true;",
		},
		New: []string{
			"option (gogoproto.unmarshaler_all) = true;",
			"option (gogoproto.marshaler_all) = true;",
			"option (gogoproto.unsafe_unmarshaler_all) = false;",
			"option (gogoproto.unsafe_marshaler_all) = false;",
		},
	}.Test(t)
}

func TestUnsafeMarshaler(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	MixMatch{
		Old: []string{
			"option (gogoproto.marshaler_all) = true;",
			"option (gogoproto.unmarshaler_all) = true;",
			"option (gogoproto.unsafe_unmarshaler_all) = true;",
			"option (gogoproto.unsafe_marshaler_all) = false;",
		},
		New: []string{
			"option (gogoproto.marshaler_all) = false;",
			"option (gogoproto.unmarshaler_all) = false;",
			"option (gogoproto.unsafe_unmarshaler_all) = false;",
			"option (gogoproto.unsafe_marshaler_all) = true;",
		},
	}.Test(t)
}

func TestUnsafeUnMarshaler(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	MixMatch{
		Old: []string{
			"option (gogoproto.marshaler_all) = true;",
			"option (gogoproto.unmarshaler_all) = true;",
			"option (gogoproto.unsafe_unmarshaler_all) = false;",
			"option (gogoproto.unsafe_marshaler_all) = true;",
		},
		New: []string{
			"option (gogoproto.marshaler_all) = false;",
			"option (gogoproto.unmarshaler_all) = false;",
			"option (gogoproto.unsafe_unmarshaler_all) = true;",
			"option (gogoproto.unsafe_marshaler_all) = false;",
		},
	}.Test(t)
}

func TestBothUnsafe(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	MixMatch{
		Old: []string{
			"option (gogoproto.unmarshaler_all) = true;",
			"option (gogoproto.marshaler_all) = true;",
			"option (gogoproto.unsafe_unmarshaler_all) = false;",
			"option (gogoproto.unsafe_marshaler_all) = false;",
		},
		New: []string{
			"option (gogoproto.unmarshaler_all) = false;",
			"option (gogoproto.marshaler_all) = false;",
			"option (gogoproto.unsafe_unmarshaler_all) = true;",
			"option (gogoproto.unsafe_marshaler_all) = true;",
		},
	}.Test(t)
}

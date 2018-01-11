// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
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

package copyjsontags

import (
	"bytes"
	"encoding/xml"
	"testing"
)

func TestXml(t *testing.T) {
	s := `<Outside><inside><MyField1>Field1Value</MyField1><XXX_unrecognized></XXX_unrecognized></inside><field_two>Field2Value</field_two><XXX_unrecognized></XXX_unrecognized></Outside>`
	field1 := "Field1Value"
	field2 := "Field2Value"
	msg1 := &Outside{}
	err := xml.Unmarshal([]byte(s), msg1)
	if err != nil {
		panic(err)
	}
	msg2 := &Outside{
		Inside: &Inside{
			FieldOne: &field1,
		},
		FieldTwo: &field2,
	}
	if msg1.GetInside().GetFieldOne() != msg2.GetInside().GetFieldOne() {
		t.Fatalf("field1 expected %s got %s", msg2.GetInside().GetFieldOne(), msg1.GetInside().GetFieldOne())
	}
	if err != nil {
		panic(err)
	}
	data, err := xml.Marshal(msg2)
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(data, []byte(s)) {
		t.Fatalf("expected %s got %s", s, string(data))
	}
}

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

package issue635_test

import (
	fmt "fmt"
	"testing"

	"github.com/gogo/protobuf/test"
	"github.com/gogo/protobuf/test/issue635"
)

func TestEnumNullableGoString(t *testing.T) {
	theEnum := issue635.TheEnum_TheEnum0
	enumNullable := &issue635.EnumNullable{TheEnum: &theEnum}
	expected := `&issue635.EnumNullable{TheEnum: func(v issue635.TheEnum) *issue635.TheEnum { return &v } ( 0 ),
}`
	testGostring(t, enumNullable, expected)
}

func TestEnumNonNullableGoString(t *testing.T) {
	enumNonNullable := &issue635.EnumNonNullable{TheEnum: issue635.TheEnum_TheEnum0}
	expected := `&issue635.EnumNonNullable{TheEnum: 0,
}`
	testGostring(t, enumNonNullable, expected)
}

func TestEnumNullableRepeatedGoString(t *testing.T) {
	enumNullableRepeated := &issue635.EnumNullableRepeated{TheEnum: []issue635.TheEnum{issue635.TheEnum_TheEnum0, issue635.TheEnum_TheEnum0}}
	expected := `&issue635.EnumNullableRepeated{TheEnum: []issue635.TheEnum{0, 0},
}`
	testGostring(t, enumNullableRepeated, expected)
}

func TestEnumNullableDifferentPackageGoString(t *testing.T) {
	theEnum := test.A
	enumNullableDifferentPackage := &issue635.EnumNullableDifferentPackage{TheEnum: &theEnum}
	expected := `&issue635.EnumNullableDifferentPackage{TheEnum: func(v test.TheTestEnum) *test.TheTestEnum { return &v } ( 0 ),
}`
	testGostring(t, enumNullableDifferentPackage, expected)
}

func TestEnumNullableNestedGoString(t *testing.T) {
	nestedEnum := issue635.NestedDefinition_TYPE_NESTED
	enumNullableNested := &issue635.EnumNullableNested{NestedEnum: &nestedEnum}
	expected := `&issue635.EnumNullableNested{NestedEnum: func(v issue635.NestedDefinition_NestedEnum) *issue635.NestedDefinition_NestedEnum { return &v } ( 1 ),
}`
	testGostring(t, enumNullableNested, expected)
}

func TestStringNullableGoString(t *testing.T) {
	theString := "theString"
	stringNullable := &issue635.StringNullable{TheString: &theString}
	expected := `&issue635.StringNullable{TheString: func(v string) *string { return &v } ( "theString" ),
}`
	testGostring(t, stringNullable, expected)
}
func TestStringNonNullableGoString(t *testing.T) {
	stringNonNullable := &issue635.StringNonNullable{TheString: "theString"}
	expected := `&issue635.StringNonNullable{TheString: "theString",
}`
	testGostring(t, stringNonNullable, expected)
}

func TestStringNullableRepeatedGoString(t *testing.T) {
	stringNullableRepeated := &issue635.StringNullableRepeated{TheString: []string{"theFirstString", "theSecondString"}}
	expected := `&issue635.StringNullableRepeated{TheString: []string{"theFirstString", "theSecondString"},
}`
	testGostring(t, stringNullableRepeated, expected)
}

func testGostring(t *testing.T, gs fmt.GoStringer, expected string) {
	actual := gs.GoString()
	if expected != actual {
		t.Fatalf("expected:\n%s\ngot:\n%s\n", expected, actual)
	}
}

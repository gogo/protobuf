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
	"bytes"
	gogoproto "code.google.com/p/gogoprotobuf/proto"
	"code.google.com/p/gogoprotobuf/test/custom"
	gotest "code.google.com/p/gogoprotobuf/test/go"
	gogotest "code.google.com/p/gogoprotobuf/test/gogo"
	goproto "code.google.com/p/goprotobuf/proto"
	"fmt"
	"math"
	"reflect"
	"testing"
)

//if true files are written with the encoded gogo messages, this is useful for testing against another language
var gogoFiles bool = false

func testStr(t *testing.T, a gogoproto.CustomStringType, astr string) {
	mastr, err := a.MarshalToString()
	if err != nil {
		panic(err)
	}
	if mastr == nil {
		t.Fatalf("MarshalToSring returned nil")
	}
	if *mastr != astr {
		t.Fatalf("%v != %v", mastr, astr)
	}
	ptr := reflect.New(reflect.TypeOf(a).Elem())
	b := ptr.Interface().(gogoproto.CustomStringType)
	if err := b.UnmarshalFromString(mastr); err != nil {
		panic(err)
	}
	mastr, err = b.MarshalToString()
	if err != nil {
		panic(err)
	}
	if mastr == nil {
		t.Fatalf("MarshalToSring returned nil after UnmarshalFromString")
	}
	if *mastr != astr {
		t.Fatalf("After UnmarshalFromString %v != %v", mastr, astr)
	}
}

func testBytes(t *testing.T, a gogoproto.CustomBytesType, astr []byte) {
	mastr, err := a.MarshalToBytes()
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(mastr, astr) {
		t.Fatalf("%v != %v", mastr, astr)
	}
	ptr := reflect.New(reflect.TypeOf(a).Elem())
	b := ptr.Interface().(gogoproto.CustomBytesType)
	if err := b.UnmarshalFromBytes(mastr); err != nil {
		panic(err)
	}
	mastr, err = b.MarshalToBytes()
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(mastr, astr) {
		t.Fatalf("After UnmarshalFromBytes %v != %v", mastr, astr)
	}
}

var name = "name"
var time int64 = 1

var id1 = custom.Uuid{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x46, 0x07, 0x88, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
var id2 = custom.Uuid{0, 1, 2, 3, 4, 5, 70, 7, 136, 9, 10, 11, 12, 15, 14, 15}

var idstr1 = "00010203-0405-4607-8809-0a0b0c0d0e0f"
var idstr2 = "00010203-0405-4607-8809-0a0b0c0f0e0f"

func TestVarUuidStr(t *testing.T) {
	testStr(t, &id1, idstr1)
	testStr(t, &id2, idstr2)
}

var idbytes1 = []byte(id1)
var idbytes2 = []byte(id2)

func TestVarUuidBytes(t *testing.T) {
	testBytes(t, &id1, idbytes1)
	testBytes(t, &id2, idbytes2)
}

var uint128a = custom.Uint128{math.MaxUint64, math.MaxUint64}
var uint128b = custom.Uint128{0, 2}
var zerouint128 = custom.Uint128{0, 0}

var uint128stra = fmt.Sprintf("%v-%v", uint64(math.MaxUint64), uint64(math.MaxUint64))
var uint128strb = "0-2"
var zerouint128str = "0-0"

func TestVarUint128Str(t *testing.T) {
	testStr(t, &uint128a, uint128stra)
	testStr(t, &uint128b, uint128strb)
}

var uint128bytesa = []byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
var uint128bytesb = []byte{0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0}
var zerouint128bytes = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

func TestVarUint128Bytes(t *testing.T) {
	testBytes(t, &uint128a, uint128bytesa)
	testBytes(t, &uint128b, uint128bytesb)
	testBytes(t, &zerouint128, zerouint128bytes)
}

var uint16a = custom.Uint16(uint16(math.MaxUint16))
var uint16b = custom.Uint16(uint16(2))
var zerouint16 = custom.Uint16(uint16(0))

var uint16stra = fmt.Sprintf("%v", math.MaxUint16)
var uint16strb = "2"
var zerouint16str = "0"

func TestVarUint16Str(t *testing.T) {
	testStr(t, &uint16a, uint16stra)
	testStr(t, &uint16b, uint16strb)
}

var uint16bytesa = []byte{255, 255}
var uint16bytesb = []byte{2, 0}
var zerouint16bytes = []byte{0, 0}

func TestVarUint16Bytes(t *testing.T) {
	testBytes(t, &uint16a, uint16bytesa)
	testBytes(t, &uint16b, uint16bytesb)
	testBytes(t, &zerouint16, zerouint16bytes)
}

var uint16u32a = uint32(math.MaxUint16)
var uint16u32b = uint32(2)
var zerouint16u32 = uint32(0)

func testUint32(t *testing.T, a gogoproto.CustomUint32Type, astr uint32) {
	mastr, err := a.MarshalToUint32()
	if err != nil {
		panic(err)
	}
	if mastr == nil {
		t.Fatalf("MarshalToUint32 returned nil")
	}
	if *mastr != astr {
		t.Fatalf("%v != %v", mastr, astr)
	}
	ptr := reflect.New(reflect.TypeOf(a).Elem())
	b := ptr.Interface().(gogoproto.CustomUint32Type)
	if err := b.UnmarshalFromUint32(mastr); err != nil {
		panic(err)
	}
	mastr, err = b.MarshalToUint32()
	if err != nil {
		panic(err)
	}
	if mastr == nil {
		t.Fatalf("MarshalToUint32 returned nil after UnmarshalFromUint32")
	}
	if *mastr != astr {
		t.Fatalf("After UnmarshalFromUint32 %v != %v", mastr, astr)
	}
}

func TestVarUint16u32(t *testing.T) {
	testUint32(t, &uint16a, uint16u32a)
	testUint32(t, &uint16b, uint16u32b)
}

var (
	newGoProtoSimpleMsg = func() goproto.Message { return &gotest.SimpleMessage{} }
)

var gosimplemsg = &gotest.SimpleMessage{
	Id:         &idstr1,
	SimpleName: &name,
	Time:       &time,
}

var (
	newGoGoProtoSimpleMsg = func() gogoproto.Message { return &gogotest.SimpleMessage{} }
)

var gogosimplemsg = &gogotest.SimpleMessage{
	Id:         id1,
	SimpleName: name,
	Time:       &time,
}

var (
	field1  float64 = math.MaxFloat64
	field2  float32 = math.MaxFloat32 //This should be MaxFloat32 when http://code.google.com/p/go/issues/detail?id=4282 is fixed
	field3  int32   = math.MaxInt32
	field4  int64   = math.MinInt64
	field5  uint32  = math.MaxUint32
	field6  uint64  = math.MaxUint64
	field7  int32   = math.MinInt32
	field8  int64   = math.MaxInt64
	field9  uint32  = math.MaxUint32 - 11
	field10 int32   = math.MaxInt32 - 11
	field11 uint64  = math.MaxUint64 - 11
	field12 int64   = math.MaxInt64 - 11
	field13 bool    = true
	field14 string  = "bla bla bla bla        \n"
	field15 []byte  = []byte{255, 0, 255, 0, 1, 2, 3, 4}
)

var (
	zero1  float64 = 0
	zero2  float32 = 0
	zero3  int32   = 0
	zero4  int64   = 0
	zero5  uint32  = 0
	zero6  uint64  = 0
	zero7  int32   = 0
	zero8  int64   = 0
	zero9  uint32  = 0
	zero10 int32   = 0
	zero11 uint64  = 0
	zero12 int64   = 0
	zero13 bool    = false
	zero14 string  = ""
	zero15 []byte  = nil
)

var (
	newGoProtoNidOptNative   = func() goproto.Message { return &gotest.NidOptNative{} }
	newGoGoProtoNidOptNative = func() gogoproto.Message { return &gogotest.NidOptNative{} }
)

var (
	govisNidOptNative = &gotest.NidOptNative{
		Field1:  &field1,
		Field2:  &field2,
		Field3:  &field3,
		Field4:  &field4,
		Field5:  &field5,
		Field6:  &field6,
		Field7:  &field7,
		Field8:  &field8,
		Field9:  &field9,
		Field10: &field10,
		Field11: &field11,
		Field12: &field12,
		Field13: &field13,
		Field14: &field14,
		Field15: field15,
	}
	gogovisNidOptNative = &gogotest.NidOptNative{
		Field1:  field1,
		Field2:  field2,
		Field3:  field3,
		Field4:  field4,
		Field5:  field5,
		Field6:  field6,
		Field7:  field7,
		Field8:  field8,
		Field9:  field9,
		Field10: field10,
		Field11: field11,
		Field12: field12,
		Field13: field13,
		Field14: field14,
		Field15: field15,
	}
	//swapped field4 and field8
	govisNidOptNative2 = &gotest.NidOptNative{
		Field1:  &field1,
		Field2:  &field2,
		Field3:  &field3,
		Field4:  &field8,
		Field5:  &field5,
		Field6:  &field6,
		Field7:  &field7,
		Field8:  &field4,
		Field9:  &field9,
		Field10: &field10,
		Field11: &field11,
		Field12: &field12,
		Field13: &field13,
		Field14: &field14,
		Field15: field15,
	}
	gogovisNidOptNative2 = &gogotest.NidOptNative{
		Field1:  field1,
		Field2:  field2,
		Field3:  field3,
		Field4:  field8,
		Field5:  field5,
		Field6:  field6,
		Field7:  field7,
		Field8:  field4,
		Field9:  field9,
		Field10: field10,
		Field11: field11,
		Field12: field12,
		Field13: field13,
		Field14: field14,
		Field15: field15,
	}
)

var (
	govimNidOptNative = &gotest.NidOptNative{
		Field1:  nil,
		Field2:  nil,
		Field3:  nil,
		Field4:  nil,
		Field5:  nil,
		Field6:  nil,
		Field7:  nil,
		Field8:  nil,
		Field9:  nil,
		Field10: nil,
		Field11: nil,
		Field12: nil,
		Field13: nil,
		Field14: nil,
		Field15: nil,
	}
	gozeroNidOptNative = &gotest.NidOptNative{
		Field1:  &zero1,
		Field2:  &zero2,
		Field3:  &zero3,
		Field4:  &zero4,
		Field5:  &zero5,
		Field6:  &zero6,
		Field7:  &zero7,
		Field8:  &zero8,
		Field9:  &zero9,
		Field10: &zero10,
		Field11: &zero11,
		Field12: &zero12,
		Field13: &zero13,
		Field14: &zero14,
		Field15: zero15,
	}
	gogovimNidOptNative = &gogotest.NidOptNative{
		Field1:  zero1,
		Field2:  zero2,
		Field3:  zero3,
		Field4:  zero4,
		Field5:  zero5,
		Field6:  zero6,
		Field7:  zero7,
		Field8:  zero8,
		Field9:  zero9,
		Field10: zero10,
		Field11: zero11,
		Field12: zero12,
		Field13: zero13,
		Field14: zero14,
		Field15: zero15,
	}
)

var (
	newGoProtoNinOptNative   = func() goproto.Message { return &gotest.NinOptNative{} }
	newGoGoProtoNinOptNative = func() gogoproto.Message { return &gogotest.NinOptNative{} }
)

var (
	govisNinOptNative = &gotest.NinOptNative{
		Field1:  &field1,
		Field2:  &field2,
		Field3:  &field3,
		Field4:  &field4,
		Field5:  &field5,
		Field6:  &field6,
		Field7:  &field7,
		Field8:  &field8,
		Field9:  &field9,
		Field10: &field10,
		Field11: &field11,
		Field12: &field12,
		Field13: &field13,
		Field14: &field14,
		Field15: field15,
	}
	gogovisNinOptNative = &gogotest.NinOptNative{
		Field1:  &field1,
		Field2:  &field2,
		Field3:  &field3,
		Field4:  &field4,
		Field5:  &field5,
		Field6:  &field6,
		Field7:  &field7,
		Field8:  &field8,
		Field9:  &field9,
		Field10: &field10,
		Field11: &field11,
		Field12: &field12,
		Field13: &field13,
		Field14: &field14,
		Field15: field15,
	}
	govimNinOptNative = &gotest.NinOptNative{
		Field1:  nil,
		Field2:  nil,
		Field3:  nil,
		Field4:  nil,
		Field5:  nil,
		Field6:  nil,
		Field7:  nil,
		Field8:  nil,
		Field9:  nil,
		Field10: nil,
		Field11: nil,
		Field12: nil,
		Field13: nil,
		Field14: nil,
		Field15: nil,
	}
	gogovimNinOptNative = &gogotest.NinOptNative{
		Field1:  nil,
		Field2:  nil,
		Field3:  nil,
		Field4:  nil,
		Field5:  nil,
		Field6:  nil,
		Field7:  nil,
		Field8:  nil,
		Field9:  nil,
		Field10: nil,
		Field11: nil,
		Field12: nil,
		Field13: nil,
		Field14: nil,
		Field15: nil,
	}
)

var (
	fields1  = []float64{math.MaxFloat64}
	fields2  = []float32{math.MaxFloat32} //This should be MaxFloat32 when http://code.google.com/p/go/issues/detail?id=4282 is fixed
	fields3  = []int32{math.MaxInt32}
	fields4  = []int64{math.MinInt64}
	fields5  = []uint32{math.MaxUint32}
	fields6  = []uint64{math.MaxUint64}
	fields7  = []int32{math.MinInt32}
	fields8  = []int64{math.MaxInt64}
	fields9  = []uint32{math.MaxUint32 - 11}
	fields10 = []int32{math.MaxInt32 - 11}
	fields11 = []uint64{math.MaxUint64 - 11}
	fields12 = []int64{math.MaxInt64 - 11}
	fields13 = []bool{true}
	fields14 = []string{"bla bla bla bla        \n"}
	fields15 = [][]byte{{255, 0, 255, 0, 1, 2, 3, 4}}
)

var (
	newGoProtoNidRepNative   = func() goproto.Message { return &gotest.NidRepNative{} }
	newGoGoProtoNidRepNative = func() gogoproto.Message { return &gogotest.NidRepNative{} }
	newGoProtoNinRepNative   = func() goproto.Message { return &gotest.NinRepNative{} }
	newGoGoProtoNinRepNative = func() gogoproto.Message { return &gogotest.NinRepNative{} }
)

var (
	govisNidRepNative = &gotest.NidRepNative{
		Field1:  fields1,
		Field2:  fields2,
		Field3:  fields3,
		Field4:  fields4,
		Field5:  fields5,
		Field6:  fields6,
		Field7:  fields7,
		Field8:  fields8,
		Field9:  fields9,
		Field10: fields10,
		Field11: fields11,
		Field12: fields12,
		Field13: fields13,
		Field14: fields14,
		Field15: fields15,
	}
	gogovisNidRepNative = &gogotest.NidRepNative{
		Field1:  fields1,
		Field2:  fields2,
		Field3:  fields3,
		Field4:  fields4,
		Field5:  fields5,
		Field6:  fields6,
		Field7:  fields7,
		Field8:  fields8,
		Field9:  fields9,
		Field10: fields10,
		Field11: fields11,
		Field12: fields12,
		Field13: fields13,
		Field14: fields14,
		Field15: fields15,
	}
	govimNidRepNative = &gotest.NidRepNative{
		Field1:  nil,
		Field2:  nil,
		Field3:  nil,
		Field4:  nil,
		Field5:  nil,
		Field6:  nil,
		Field7:  nil,
		Field8:  nil,
		Field9:  nil,
		Field10: nil,
		Field11: nil,
		Field12: nil,
		Field13: nil,
		Field14: nil,
		Field15: nil,
	}
	gogovimNidRepNative = &gogotest.NidRepNative{
		Field1:  nil,
		Field2:  nil,
		Field3:  nil,
		Field4:  nil,
		Field5:  nil,
		Field6:  nil,
		Field7:  nil,
		Field8:  nil,
		Field9:  nil,
		Field10: nil,
		Field11: nil,
		Field12: nil,
		Field13: nil,
		Field14: nil,
		Field15: nil,
	}
	govisNinRepNative = &gotest.NinRepNative{
		Field1:  fields1,
		Field2:  fields2,
		Field3:  fields3,
		Field4:  fields4,
		Field5:  fields5,
		Field6:  fields6,
		Field7:  fields7,
		Field8:  fields8,
		Field9:  fields9,
		Field10: fields10,
		Field11: fields11,
		Field12: fields12,
		Field13: fields13,
		Field14: fields14,
		Field15: fields15,
	}
	gogovisNinRepNative = &gogotest.NinRepNative{
		Field1:  fields1,
		Field2:  fields2,
		Field3:  fields3,
		Field4:  fields4,
		Field5:  fields5,
		Field6:  fields6,
		Field7:  fields7,
		Field8:  fields8,
		Field9:  fields9,
		Field10: fields10,
		Field11: fields11,
		Field12: fields12,
		Field13: fields13,
		Field14: fields14,
		Field15: fields15,
	}
	govimNinRepNative = &gotest.NinRepNative{
		Field1:  nil,
		Field2:  nil,
		Field3:  nil,
		Field4:  nil,
		Field5:  nil,
		Field6:  nil,
		Field7:  nil,
		Field8:  nil,
		Field9:  nil,
		Field10: nil,
		Field11: nil,
		Field12: nil,
		Field13: nil,
		Field14: nil,
		Field15: nil,
	}
	gogovimNinRepNative = &gogotest.NinRepNative{
		Field1:  nil,
		Field2:  nil,
		Field3:  nil,
		Field4:  nil,
		Field5:  nil,
		Field6:  nil,
		Field7:  nil,
		Field8:  nil,
		Field9:  nil,
		Field10: nil,
		Field11: nil,
		Field12: nil,
		Field13: nil,
		Field14: nil,
		Field15: nil,
	}
)

var (
	newGoProtoNidRepPackedNative   = func() goproto.Message { return &gotest.NidRepPackedNative{} }
	newGoGoProtoNidRepPackedNative = func() gogoproto.Message { return &gogotest.NidRepPackedNative{} }
	newGoProtoNinRepPackedNative   = func() goproto.Message { return &gotest.NinRepPackedNative{} }
	newGoGoProtoNinRepPackedNative = func() gogoproto.Message { return &gogotest.NinRepPackedNative{} }
)

var (
	govisNidRepPackedNative = &gotest.NidRepPackedNative{
		Field3:  fields3,
		Field4:  fields4,
		Field13: fields13,
	}
	gogovisNidRepPackedNative = &gogotest.NidRepPackedNative{
		Field3:  fields3,
		Field4:  fields4,
		Field13: fields13,
	}
	govimNidRepPackedNative = &gotest.NidRepPackedNative{
		Field3:  nil,
		Field4:  nil,
		Field13: nil,
	}
	gogovimNidRepPackedNative = &gogotest.NidRepPackedNative{
		Field3:  nil,
		Field4:  nil,
		Field13: nil,
	}
	govisNinRepPackedNative = &gotest.NinRepPackedNative{
		Field3:  fields3,
		Field4:  fields4,
		Field13: fields13,
	}
	gogovisNinRepPackedNative = &gogotest.NinRepPackedNative{
		Field3:  fields3,
		Field4:  fields4,
		Field13: fields13,
	}
	govimNinRepPackedNative = &gotest.NinRepPackedNative{
		Field3:  nil,
		Field4:  nil,
		Field13: nil,
	}
	gogovimNinRepPackedNative = &gogotest.NinRepPackedNative{
		Field3:  nil,
		Field4:  nil,
		Field13: nil,
	}
)

var (
	newGoProtoNidOptStruct   = func() goproto.Message { return &gotest.NidOptStruct{} }
	newGoGoProtoNidOptStruct = func() gogoproto.Message { return &gogotest.NidOptStruct{} }
	newGoProtoNinOptStruct   = func() goproto.Message { return &gotest.NinOptStruct{} }
	newGoGoProtoNinOptStruct = func() gogoproto.Message { return &gogotest.NinOptStruct{} }
)

var (
	govisNidOptStruct = &gotest.NidOptStruct{
		Field1:  &field1,
		Field2:  &field2,
		Field3:  govisNidOptNative,
		Field4:  govisNinOptNative,
		Field6:  &field6,
		Field7:  &field7,
		Field8:  govisNidOptNative2,
		Field13: &field13,
		Field14: &field14,
		Field15: field15,
	}
	gogovisNidOptStruct = &gogotest.NidOptStruct{
		Field1:  field1,
		Field2:  field2,
		Field3:  *gogovisNidOptNative,
		Field4:  *gogovisNinOptNative,
		Field6:  field6,
		Field7:  field7,
		Field8:  *gogovisNidOptNative2,
		Field13: field13,
		Field14: field14,
		Field15: field15,
	}
	govimNidOptStruct = &gotest.NidOptStruct{
		Field1:  nil,
		Field2:  nil,
		Field3:  nil,
		Field4:  nil,
		Field6:  nil,
		Field7:  nil,
		Field8:  nil,
		Field13: nil,
		Field14: nil,
		Field15: nil,
	}
	gozeroNidOptStruct = &gotest.NidOptStruct{
		Field1:  &zero1,
		Field2:  &zero2,
		Field3:  gozeroNidOptNative,
		Field4:  govimNinOptNative,
		Field6:  &zero6,
		Field7:  &zero7,
		Field8:  gozeroNidOptNative,
		Field13: &zero13,
		Field14: &zero14,
		Field15: zero15,
	}
	gogovimNidOptStruct = &gogotest.NidOptStruct{
		Field1:  zero1,
		Field2:  zero2,
		Field3:  *gogovimNidOptNative,
		Field4:  *gogovimNinOptNative,
		Field6:  zero6,
		Field7:  zero7,
		Field8:  *gogovimNidOptNative,
		Field13: zero13,
		Field14: zero14,
		Field15: zero15,
	}
	govisNinOptStruct = &gotest.NinOptStruct{
		Field1:  &field1,
		Field2:  &field2,
		Field3:  govisNidOptNative,
		Field4:  govisNinOptNative,
		Field6:  &field6,
		Field7:  &field7,
		Field8:  govisNidOptNative2,
		Field13: &field13,
		Field14: &field14,
		Field15: field15,
	}
	gogovisNinOptStruct = &gogotest.NinOptStruct{
		Field1:  &field1,
		Field2:  &field2,
		Field3:  gogovisNidOptNative,
		Field4:  gogovisNinOptNative,
		Field6:  &field6,
		Field7:  &field7,
		Field8:  gogovisNidOptNative2,
		Field13: &field13,
		Field14: &field14,
		Field15: field15,
	}
	govimNinOptStruct = &gotest.NinOptStruct{
		Field1:  nil,
		Field2:  nil,
		Field3:  nil,
		Field4:  nil,
		Field6:  nil,
		Field7:  nil,
		Field8:  nil,
		Field13: nil,
		Field14: nil,
		Field15: nil,
	}
	gogovimNinOptStruct = &gogotest.NinOptStruct{
		Field1:  nil,
		Field2:  nil,
		Field3:  nil,
		Field4:  nil,
		Field6:  nil,
		Field7:  nil,
		Field8:  nil,
		Field13: nil,
		Field14: nil,
		Field15: nil,
	}
)

var (
	govisNidOptNatives     = []*gotest.NidOptNative{govisNidOptNative}
	govisNinOptNatives     = []*gotest.NinOptNative{govisNinOptNative}
	govisNidOptNatives2    = []*gotest.NidOptNative{govisNidOptNative2, govisNidOptNative}
	gogovisNidOptNatives   = []gogotest.NidOptNative{*gogovisNidOptNative}
	gogovisNinOptNatives   = []gogotest.NinOptNative{*gogovisNinOptNative}
	gogovisNidOptNatives2  = []gogotest.NidOptNative{*gogovisNidOptNative2, *gogovisNidOptNative}
	gogovisNidOptNativess  = []*gogotest.NidOptNative{gogovisNidOptNative}
	gogovisNinOptNativess  = []*gogotest.NinOptNative{gogovisNinOptNative}
	gogovisNidOptNativess2 = []*gogotest.NidOptNative{gogovisNidOptNative2, gogovisNidOptNative}
)

var (
	newGoProtoNidRepStruct   = func() goproto.Message { return &gotest.NidRepStruct{} }
	newGoGoProtoNidRepStruct = func() gogoproto.Message { return &gogotest.NidRepStruct{} }
	newGoProtoNinRepStruct   = func() goproto.Message { return &gotest.NinRepStruct{} }
	newGoGoProtoNinRepStruct = func() gogoproto.Message { return &gogotest.NinRepStruct{} }
)

var (
	govisNidRepStruct = &gotest.NidRepStruct{
		Field1:  fields1,
		Field2:  fields2,
		Field3:  govisNidOptNatives,
		Field4:  govisNinOptNatives,
		Field6:  fields6,
		Field7:  fields7,
		Field8:  govisNidOptNatives2,
		Field13: fields13,
		Field14: fields14,
		Field15: fields15,
	}
	gogovisNidRepStruct = &gogotest.NidRepStruct{
		Field1:  fields1,
		Field2:  fields2,
		Field3:  gogovisNidOptNatives,
		Field4:  gogovisNinOptNatives,
		Field6:  fields6,
		Field7:  fields7,
		Field8:  gogovisNidOptNatives2,
		Field13: fields13,
		Field14: fields14,
		Field15: fields15,
	}
	govimNidRepStruct = &gotest.NidRepStruct{
		Field1:  nil,
		Field2:  nil,
		Field3:  nil,
		Field4:  nil,
		Field6:  nil,
		Field7:  nil,
		Field8:  nil,
		Field13: nil,
		Field14: nil,
		Field15: nil,
	}
	gogovimNidRepStruct = &gogotest.NidRepStruct{
		Field1:  nil,
		Field2:  nil,
		Field3:  nil,
		Field4:  nil,
		Field6:  nil,
		Field7:  nil,
		Field8:  nil,
		Field13: nil,
		Field14: nil,
		Field15: nil,
	}
	govisNinRepStruct = &gotest.NinRepStruct{
		Field1:  fields1,
		Field2:  fields2,
		Field3:  govisNidOptNatives,
		Field4:  govisNinOptNatives,
		Field6:  fields6,
		Field7:  fields7,
		Field8:  govisNidOptNatives2,
		Field13: fields13,
		Field14: fields14,
		Field15: fields15,
	}
	gogovisNinRepStruct = &gogotest.NinRepStruct{
		Field1:  fields1,
		Field2:  fields2,
		Field3:  gogovisNidOptNativess,
		Field4:  gogovisNinOptNativess,
		Field6:  fields6,
		Field7:  fields7,
		Field8:  gogovisNidOptNativess2,
		Field13: fields13,
		Field14: fields14,
		Field15: fields15,
	}
	govimNinRepStruct = &gotest.NinRepStruct{
		Field1:  nil,
		Field2:  nil,
		Field3:  nil,
		Field4:  nil,
		Field6:  nil,
		Field7:  nil,
		Field8:  nil,
		Field13: nil,
		Field14: nil,
		Field15: nil,
	}
	gogovimNinRepStruct = &gogotest.NinRepStruct{
		Field1:  nil,
		Field2:  nil,
		Field3:  nil,
		Field4:  nil,
		Field6:  nil,
		Field7:  nil,
		Field8:  nil,
		Field13: nil,
		Field14: nil,
		Field15: nil,
	}
)

var (
	newGoProtoNidEmbeddedStruct   = func() goproto.Message { return &gotest.NidEmbeddedStruct{} }
	newGoGoProtoNidEmbeddedStruct = func() gogoproto.Message { return &gogotest.NidEmbeddedStruct{} }
	newGoProtoNinEmbeddedStruct   = func() goproto.Message { return &gotest.NinEmbeddedStruct{} }
	newGoGoProtoNinEmbeddedStruct = func() gogoproto.Message { return &gogotest.NinEmbeddedStruct{} }
)

var (
	govisNidEmbeddedStruct = &gotest.NidEmbeddedStruct{
		Field1:   govisNidOptNative,
		Field200: govisNidOptNative,
		Field210: &field13,
	}
	gogovisNidEmbeddedStruct = &gogotest.NidEmbeddedStruct{
		NidOptNative: *gogovisNidOptNative,
		Field200:     *gogovisNidOptNative,
		Field210:     field13,
	}
	govimNidEmbeddedStruct = &gotest.NidEmbeddedStruct{
		Field1:   nil,
		Field200: nil,
		Field210: nil,
	}
	gozeroNidEmbeddedStruct = &gotest.NidEmbeddedStruct{
		Field1:   gozeroNidOptNative,
		Field200: gozeroNidOptNative,
		Field210: &zero13,
	}
	gogovimNidEmbeddedStruct = &gogotest.NidEmbeddedStruct{
		NidOptNative: *gogovimNidOptNative,
		Field200:     *gogovimNidOptNative,
		Field210:     zero13,
	}
	govisNinEmbeddedStruct = &gotest.NinEmbeddedStruct{
		Field1:   govisNidOptNative,
		Field200: govisNidOptNative,
		Field210: &field13,
	}
	gogovisNinEmbeddedStruct = &gogotest.NinEmbeddedStruct{
		NidOptNative: gogovisNidOptNative,
		Field200:     gogovisNidOptNative,
		Field210:     &field13,
	}
	govimNinEmbeddedStruct = &gotest.NinEmbeddedStruct{
		Field1:   nil,
		Field200: nil,
		Field210: nil,
	}
	gogovimNinEmbeddedStruct = &gogotest.NinEmbeddedStruct{
		NidOptNative: nil,
		Field200:     nil,
		Field210:     nil,
	}
)

var (
	newGoProtoNidOptUuid   = func() goproto.Message { return &gotest.NidOptUuid{} }
	newGoGoProtoNidOptUuid = func() gogoproto.Message { return &gogotest.NidOptUuid{} }
	newGoProtoNinOptUuid   = func() goproto.Message { return &gotest.NinOptUuid{} }
	newGoGoProtoNinOptUuid = func() gogoproto.Message { return &gogotest.NinOptUuid{} }
)

var (
	govisNidOptUuid = &gotest.NidOptUuid{
		Id: &idstr1,
	}
	gogovisNidOptUuid = &gogotest.NidOptUuid{
		Id: id1,
	}
	govimNidOptUuid = &gotest.NidOptUuid{
		Id: nil,
	}
	gogovimNidOptUuid = &gogotest.NidOptUuid{
		Id: nil,
	}
	govisNinOptUuid = &gotest.NinOptUuid{
		Id: &idstr1,
	}
	gogovisNinOptUuid = &gogotest.NinOptUuid{
		Id: &id1,
	}
	govimNinOptUuid = &gotest.NinOptUuid{
		Id: nil,
	}
	gogovimNinOptUuid = &gogotest.NinOptUuid{
		Id: nil,
	}
)

var (
	newGoProtoNidRepUuid   = func() goproto.Message { return &gotest.NidRepUuid{} }
	newGoGoProtoNidRepUuid = func() gogoproto.Message { return &gogotest.NidRepUuid{} }
	newGoProtoNinRepUuid   = func() goproto.Message { return &gotest.NinRepUuid{} }
	newGoGoProtoNinRepUuid = func() gogoproto.Message { return &gogotest.NinRepUuid{} }
)

var (
	govisNidRepUuid = &gotest.NidRepUuid{
		Id: []string{idstr1, idstr2},
	}
	gogovisNidRepUuid = &gogotest.NidRepUuid{
		Id: []custom.Uuid{id1, id2},
	}
	govimNidRepUuid = &gotest.NidRepUuid{
		Id: nil,
	}
	gogovimNidRepUuid = &gogotest.NidRepUuid{
		Id: nil,
	}
	govisNinRepUuid = &gotest.NinRepUuid{
		Id: []string{idstr1, idstr2},
	}
	gogovisNinRepUuid = &gogotest.NinRepUuid{
		Id: []custom.Uuid{id1, id2},
	}
	govimNinRepUuid = &gotest.NinRepUuid{
		Id: nil,
	}
	gogovimNinRepUuid = &gogotest.NinRepUuid{
		Id: nil,
	}
)

var (
	newGoProtoNidNestedStruct   = func() goproto.Message { return &gotest.NidNestedStruct{} }
	newGoGoProtoNidNestedStruct = func() gogoproto.Message { return &gogotest.NidNestedStruct{} }
	newGoProtoNinNestedStruct   = func() goproto.Message { return &gotest.NinNestedStruct{} }
	newGoGoProtoNinNestedStruct = func() gogoproto.Message { return &gogotest.NinNestedStruct{} }
)

var (
	govisNidNestedStruct = &gotest.NidNestedStruct{
		Field1: govisNidOptStruct,
		Field2: []*gotest.NidRepStruct{govisNidRepStruct, govisNidRepStruct},
	}
	gogovisNidNestedStruct = &gogotest.NidNestedStruct{
		Field1: *gogovisNidOptStruct,
		Field2: []gogotest.NidRepStruct{*gogovisNidRepStruct, *gogovisNidRepStruct},
	}
	govimNidNestedStruct = &gotest.NidNestedStruct{
		Field1: nil,
		Field2: nil,
	}
	gozeroNidNestedStruct = &gotest.NidNestedStruct{
		Field1: gozeroNidOptStruct,
		Field2: nil,
	}
	gogovimNidNestedStruct = &gogotest.NidNestedStruct{
		Field1: *gogovimNidOptStruct,
		Field2: nil,
	}
	govisNinNestedStruct = &gotest.NinNestedStruct{
		Field1: govisNidOptStruct,
		Field2: []*gotest.NidRepStruct{govisNidRepStruct, govisNidRepStruct},
	}
	gogovisNinNestedStruct = &gogotest.NinNestedStruct{
		Field1: gogovisNidOptStruct,
		Field2: []*gogotest.NidRepStruct{gogovisNidRepStruct, gogovisNidRepStruct},
	}
	govimNinNestedStruct = &gotest.NinNestedStruct{
		Field1: nil,
		Field2: nil,
	}
	gogovimNinNestedStruct = &gogotest.NinNestedStruct{
		Field1: nil,
		Field2: nil,
	}
)

var (
	newGoProtoNidOptEnum   = func() goproto.Message { return &gotest.NidOptEnum{} }
	newGoGoProtoNidOptEnum = func() gogoproto.Message { return &gogotest.NidOptEnum{} }
	newGoProtoNinOptEnum   = func() goproto.Message { return &gotest.NinOptEnum{} }
	newGoGoProtoNinOptEnum = func() gogoproto.Message { return &gogotest.NinOptEnum{} }
)

var (
	govisNidOptEnum = &gotest.NidOptEnum{
		Field1: gotest.TheTestEnum_B.Enum(),
	}
	gogovisNidOptEnum = &gogotest.NidOptEnum{
		Field1: gogotest.B,
	}
	govimNidOptEnum = &gotest.NidOptEnum{
		Field1: nil,
	}
	gozeroNidOptEnum = &gotest.NidOptEnum{
		Field1: gotest.TheTestEnum_A.Enum(),
	}
	gogovimNidOptEnum = &gogotest.NidOptEnum{
		Field1: gogotest.A,
	}
	govisNinOptEnum = &gotest.NinOptEnum{
		Field1: gotest.TheTestEnum_B.Enum(),
	}
	gogovisNinOptEnum = &gogotest.NinOptEnum{
		Field1: gogotest.B.Enum(),
	}
	govimNinOptEnum = &gotest.NinOptEnum{
		Field1: nil,
	}
	gogovimNinOptEnum = &gogotest.NinOptEnum{
		Field1: nil,
	}
)

var (
	newGoProtoNidRepEnum   = func() goproto.Message { return &gotest.NidRepEnum{} }
	newGoGoProtoNidRepEnum = func() gogoproto.Message { return &gogotest.NidRepEnum{} }
	newGoProtoNinRepEnum   = func() goproto.Message { return &gotest.NinRepEnum{} }
	newGoGoProtoNinRepEnum = func() gogoproto.Message { return &gogotest.NinRepEnum{} }
)

var (
	govisNidRepEnum = &gotest.NidRepEnum{
		Field1: []gotest.TheTestEnum{gotest.TheTestEnum_B, gotest.TheTestEnum_C},
	}
	gogovisNidRepEnum = &gogotest.NidRepEnum{
		Field1: []gogotest.TheTestEnum{gogotest.B, gogotest.C},
	}
	govimNidRepEnum = &gotest.NidRepEnum{
		Field1: nil,
	}
	gogovimNidRepEnum = &gogotest.NidRepEnum{
		Field1: nil,
	}
	govisNinRepEnum = &gotest.NinRepEnum{
		Field1: []gotest.TheTestEnum{gotest.TheTestEnum_B, gotest.TheTestEnum_C},
	}
	gogovisNinRepEnum = &gogotest.NinRepEnum{
		Field1: []gogotest.TheTestEnum{gogotest.B, gogotest.C},
	}
	govimNinRepEnum = &gotest.NinRepEnum{
		Field1: nil,
	}
	gogovimNinRepEnum = &gogotest.NinRepEnum{
		Field1: nil,
	}
)

var (
	newGoProtoNidOptUint128   = func() goproto.Message { return &gotest.NidOptUint128{} }
	newGoGoProtoNidOptUint128 = func() gogoproto.Message { return &gogotest.NidOptUint128{} }
	newGoProtoNinOptUint128   = func() goproto.Message { return &gotest.NinOptUint128{} }
	newGoGoProtoNinOptUint128 = func() gogoproto.Message { return &gogotest.NinOptUint128{} }
)

var (
	govisNidOptUint128 = &gotest.NidOptUint128{
		Value: &uint128stra,
	}
	gogovisNidOptUint128 = &gogotest.NidOptUint128{
		Value: uint128a,
	}
	govimNidOptUint128 = &gotest.NidOptUint128{
		Value: nil,
	}
	gozeroOptUint128 = &gotest.NidOptUint128{
		Value: &zerouint128str,
	}
	gogovimNidOptUint128 = &gogotest.NidOptUint128{
		Value: custom.Uint128{},
	}
	govisNinOptUint128 = &gotest.NinOptUint128{
		Value: &uint128stra,
	}
	gogovisNinOptUint128 = &gogotest.NinOptUint128{
		Value: &uint128a,
	}
	govimNinOptUint128 = &gotest.NinOptUint128{
		Value: nil,
	}
	gogovimNinOptUint128 = &gogotest.NinOptUint128{
		Value: nil,
	}
)

var (
	newGoProtoNidRepUint128   = func() goproto.Message { return &gotest.NidRepUint128{} }
	newGoGoProtoNidRepUint128 = func() gogoproto.Message { return &gogotest.NidRepUint128{} }
	newGoProtoNinRepUint128   = func() goproto.Message { return &gotest.NinRepUint128{} }
	newGoGoProtoNinRepUint128 = func() gogoproto.Message { return &gogotest.NinRepUint128{} }
)

var (
	govisNidRepUint128 = &gotest.NidRepUint128{
		Value: []string{uint128stra, uint128strb},
	}
	gogovisNidRepUint128 = &gogotest.NidRepUint128{
		Value: []custom.Uint128{uint128a, uint128b},
	}
	govimNidRepUint128 = &gotest.NidRepUint128{
		Value: nil,
	}
	gogovimNidRepUint128 = &gogotest.NidRepUint128{
		Value: nil,
	}
	govisNinRepUint128 = &gotest.NinRepUint128{
		Value: []string{uint128stra, uint128strb},
	}
	gogovisNinRepUint128 = &gogotest.NinRepUint128{
		Value: []custom.Uint128{uint128a, uint128b},
	}
	govimNinRepUint128 = &gotest.NinRepUint128{
		Value: nil,
	}
	gogovimNinRepUint128 = &gogotest.NinRepUint128{
		Value: nil,
	}
)

var (
	newGoProtoNidOptUint16   = func() goproto.Message { return &gotest.NidOptUint16{} }
	newGoGoProtoNidOptUint16 = func() gogoproto.Message { return &gogotest.NidOptUint16{} }
	newGoProtoNinOptUint16   = func() goproto.Message { return &gotest.NinOptUint16{} }
	newGoGoProtoNinOptUint16 = func() gogoproto.Message { return &gogotest.NinOptUint16{} }
)

var (
	govisNidOptUint16 = &gotest.NidOptUint16{
		Value: &uint16stra,
	}
	gogovisNidOptUint16 = &gogotest.NidOptUint16{
		Value: uint16a,
	}
	govimNidOptUint16 = &gotest.NidOptUint16{
		Value: nil,
	}
	gozeroOptUint16 = &gotest.NidOptUint16{
		Value: &zerouint16str,
	}
	gogovimNidOptUint16 = &gogotest.NidOptUint16{
		Value: 0,
	}
	govisNinOptUint16 = &gotest.NinOptUint16{
		Value: &uint16stra,
	}
	gogovisNinOptUint16 = &gogotest.NinOptUint16{
		Value: &uint16a,
	}
	govimNinOptUint16 = &gotest.NinOptUint16{
		Value: nil,
	}
	gogovimNinOptUint16 = &gogotest.NinOptUint16{
		Value: nil,
	}
)

var (
	newGoProtoNidRepUint16   = func() goproto.Message { return &gotest.NidRepUint16{} }
	newGoGoProtoNidRepUint16 = func() gogoproto.Message { return &gogotest.NidRepUint16{} }
	newGoProtoNinRepUint16   = func() goproto.Message { return &gotest.NinRepUint16{} }
	newGoGoProtoNinRepUint16 = func() gogoproto.Message { return &gogotest.NinRepUint16{} }
)

var (
	govisNidRepUint16 = &gotest.NidRepUint16{
		Value: []string{uint16stra, uint16strb},
	}
	gogovisNidRepUint16 = &gogotest.NidRepUint16{
		Value: []custom.Uint16{uint16a, uint16b},
	}
	govimNidRepUint16 = &gotest.NidRepUint16{
		Value: nil,
	}
	gogovimNidRepUint16 = &gogotest.NidRepUint16{
		Value: nil,
	}
	govisNinRepUint16 = &gotest.NinRepUint16{
		Value: []string{uint16stra, uint16strb},
	}
	gogovisNinRepUint16 = &gogotest.NinRepUint16{
		Value: []custom.Uint16{uint16a, uint16b},
	}
	govimNinRepUint16 = &gotest.NinRepUint16{
		Value: nil,
	}
	gogovimNinRepUint16 = &gogotest.NinRepUint16{
		Value: nil,
	}
)

var (
	newGoProtoNidOptUuidAsBytes   = func() goproto.Message { return &gotest.NidOptUuidAsBytes{} }
	newGoGoProtoNidOptUuidAsBytes = func() gogoproto.Message { return &gogotest.NidOptUuidAsBytes{} }
	newGoProtoNinOptUuidAsBytes   = func() goproto.Message { return &gotest.NinOptUuidAsBytes{} }
	newGoGoProtoNinOptUuidAsBytes = func() gogoproto.Message { return &gogotest.NinOptUuidAsBytes{} }
)

var (
	govisNidOptUuidAsBytes = &gotest.NidOptUuidAsBytes{
		Id: idbytes1,
	}
	gogovisNidOptUuidAsBytes = &gogotest.NidOptUuidAsBytes{
		Id: id1,
	}
	govimNidOptUuidAsBytes = &gotest.NidOptUuidAsBytes{
		Id: nil,
	}
	gogovimNidOptUuidAsBytes = &gogotest.NidOptUuidAsBytes{
		Id: nil,
	}
	govisNinOptUuidAsBytes = &gotest.NinOptUuidAsBytes{
		Id: idbytes1,
	}
	gogovisNinOptUuidAsBytes = &gogotest.NinOptUuidAsBytes{
		Id: &id1,
	}
	govimNinOptUuidAsBytes = &gotest.NinOptUuidAsBytes{
		Id: nil,
	}
	gogovimNinOptUuidAsBytes = &gogotest.NinOptUuidAsBytes{
		Id: nil,
	}
)

var (
	newGoProtoNidRepUuidAsBytes   = func() goproto.Message { return &gotest.NidRepUuidAsBytes{} }
	newGoGoProtoNidRepUuidAsBytes = func() gogoproto.Message { return &gogotest.NidRepUuidAsBytes{} }
	newGoProtoNinRepUuidAsBytes   = func() goproto.Message { return &gotest.NinRepUuidAsBytes{} }
	newGoGoProtoNinRepUuidAsBytes = func() gogoproto.Message { return &gogotest.NinRepUuidAsBytes{} }
)

var (
	govisNidRepUuidAsBytes = &gotest.NidRepUuidAsBytes{
		Id: [][]byte{idbytes1, idbytes2},
	}
	gogovisNidRepUuidAsBytes = &gogotest.NidRepUuidAsBytes{
		Id: []custom.Uuid{id1, id2},
	}
	govimNidRepUuidAsBytes = &gotest.NidRepUuidAsBytes{
		Id: nil,
	}
	gogovimNidRepUuidAsBytes = &gogotest.NidRepUuidAsBytes{
		Id: nil,
	}
	govisNinRepUuidAsBytes = &gotest.NinRepUuidAsBytes{
		Id: [][]byte{idbytes1, idbytes2},
	}
	gogovisNinRepUuidAsBytes = &gogotest.NinRepUuidAsBytes{
		Id: []custom.Uuid{id1, id2},
	}
	govimNinRepUuidAsBytes = &gotest.NinRepUuidAsBytes{
		Id: nil,
	}
	gogovimNinRepUuidAsBytes = &gogotest.NinRepUuidAsBytes{
		Id: nil,
	}
)

var (
	newGoProtoNidOptUint128AsBytes   = func() goproto.Message { return &gotest.NidOptUint128AsBytes{} }
	newGoGoProtoNidOptUint128AsBytes = func() gogoproto.Message { return &gogotest.NidOptUint128AsBytes{} }
	newGoProtoNinOptUint128AsBytes   = func() goproto.Message { return &gotest.NinOptUint128AsBytes{} }
	newGoGoProtoNinOptUint128AsBytes = func() gogoproto.Message { return &gogotest.NinOptUint128AsBytes{} }
)

var (
	govisNidOptUint128AsBytes = &gotest.NidOptUint128AsBytes{
		Value: uint128bytesa,
	}
	gogovisNidOptUint128AsBytes = &gogotest.NidOptUint128AsBytes{
		Value: uint128a,
	}
	govimNidOptUint128AsBytes = &gotest.NidOptUint128AsBytes{
		Value: nil,
	}
	gozeroOptUint128AsBytes = &gotest.NidOptUint128AsBytes{
		Value: zerouint128bytes,
	}
	gogovimNidOptUint128AsBytes = &gogotest.NidOptUint128AsBytes{
		Value: custom.Uint128{},
	}
	govisNinOptUint128AsBytes = &gotest.NinOptUint128AsBytes{
		Value: uint128bytesa,
	}
	gogovisNinOptUint128AsBytes = &gogotest.NinOptUint128AsBytes{
		Value: &uint128a,
	}
	govimNinOptUint128AsBytes = &gotest.NinOptUint128AsBytes{
		Value: nil,
	}
	gogovimNinOptUint128AsBytes = &gogotest.NinOptUint128AsBytes{
		Value: nil,
	}
)

var (
	newGoProtoNidRepUint128AsBytes   = func() goproto.Message { return &gotest.NidRepUint128AsBytes{} }
	newGoGoProtoNidRepUint128AsBytes = func() gogoproto.Message { return &gogotest.NidRepUint128AsBytes{} }
	newGoProtoNinRepUint128AsBytes   = func() goproto.Message { return &gotest.NinRepUint128AsBytes{} }
	newGoGoProtoNinRepUint128AsBytes = func() gogoproto.Message { return &gogotest.NinRepUint128AsBytes{} }
)

var (
	govisNidRepUint128AsBytes = &gotest.NidRepUint128AsBytes{
		Value: [][]byte{uint128bytesa, uint128bytesb},
	}
	gogovisNidRepUint128AsBytes = &gogotest.NidRepUint128AsBytes{
		Value: []custom.Uint128{uint128a, uint128b},
	}
	govimNidRepUint128AsBytes = &gotest.NidRepUint128AsBytes{
		Value: nil,
	}
	gogovimNidRepUint128AsBytes = &gogotest.NidRepUint128AsBytes{
		Value: nil,
	}
	govisNinRepUint128AsBytes = &gotest.NinRepUint128AsBytes{
		Value: [][]byte{uint128bytesa, uint128bytesb},
	}
	gogovisNinRepUint128AsBytes = &gogotest.NinRepUint128AsBytes{
		Value: []custom.Uint128{uint128a, uint128b},
	}
	govimNinRepUint128AsBytes = &gotest.NinRepUint128AsBytes{
		Value: nil,
	}
	gogovimNinRepUint128AsBytes = &gogotest.NinRepUint128AsBytes{
		Value: nil,
	}
)

var (
	newGoProtoNidOptUint16AsBytes   = func() goproto.Message { return &gotest.NidOptUint16AsBytes{} }
	newGoGoProtoNidOptUint16AsBytes = func() gogoproto.Message { return &gogotest.NidOptUint16AsBytes{} }
	newGoProtoNinOptUint16AsBytes   = func() goproto.Message { return &gotest.NinOptUint16AsBytes{} }
	newGoGoProtoNinOptUint16AsBytes = func() gogoproto.Message { return &gogotest.NinOptUint16AsBytes{} }
)

var (
	govisNidOptUint16AsBytes = &gotest.NidOptUint16AsBytes{
		Value: uint16bytesa,
	}
	gogovisNidOptUint16AsBytes = &gogotest.NidOptUint16AsBytes{
		Value: uint16a,
	}
	govimNidOptUint16AsBytes = &gotest.NidOptUint16AsBytes{
		Value: nil,
	}
	gozeroOptUint16AsBytes = &gotest.NidOptUint16AsBytes{
		Value: zerouint16bytes,
	}
	gogovimNidOptUint16AsBytes = &gogotest.NidOptUint16AsBytes{
		Value: 0,
	}
	govisNinOptUint16AsBytes = &gotest.NinOptUint16AsBytes{
		Value: uint16bytesa,
	}
	gogovisNinOptUint16AsBytes = &gogotest.NinOptUint16AsBytes{
		Value: &uint16a,
	}
	govimNinOptUint16AsBytes = &gotest.NinOptUint16AsBytes{
		Value: nil,
	}
	gogovimNinOptUint16AsBytes = &gogotest.NinOptUint16AsBytes{
		Value: nil,
	}
)

var (
	newGoProtoNidRepUint16AsBytes   = func() goproto.Message { return &gotest.NidRepUint16AsBytes{} }
	newGoGoProtoNidRepUint16AsBytes = func() gogoproto.Message { return &gogotest.NidRepUint16AsBytes{} }
	newGoProtoNinRepUint16AsBytes   = func() goproto.Message { return &gotest.NinRepUint16AsBytes{} }
	newGoGoProtoNinRepUint16AsBytes = func() gogoproto.Message { return &gogotest.NinRepUint16AsBytes{} }
)

var (
	govisNidRepUint16AsBytes = &gotest.NidRepUint16AsBytes{
		Value: [][]byte{uint16bytesa, uint16bytesb},
	}
	gogovisNidRepUint16AsBytes = &gogotest.NidRepUint16AsBytes{
		Value: []custom.Uint16{uint16a, uint16b},
	}
	govimNidRepUint16AsBytes = &gotest.NidRepUint16AsBytes{
		Value: nil,
	}
	gogovimNidRepUint16AsBytes = &gogotest.NidRepUint16AsBytes{
		Value: nil,
	}
	govisNinRepUint16AsBytes = &gotest.NinRepUint16AsBytes{
		Value: [][]byte{uint16bytesa, uint16bytesb},
	}
	gogovisNinRepUint16AsBytes = &gogotest.NinRepUint16AsBytes{
		Value: []custom.Uint16{uint16a, uint16b},
	}
	govimNinRepUint16AsBytes = &gotest.NinRepUint16AsBytes{
		Value: nil,
	}
	gogovimNinRepUint16AsBytes = &gogotest.NinRepUint16AsBytes{
		Value: nil,
	}
)

var (
	newGoProtoNidOptUint16AsUint32   = func() goproto.Message { return &gotest.NidOptUint16AsUint32{} }
	newGoGoProtoNidOptUint16AsUint32 = func() gogoproto.Message { return &gogotest.NidOptUint16AsUint32{} }
	newGoProtoNinOptUint16AsUint32   = func() goproto.Message { return &gotest.NinOptUint16AsUint32{} }
	newGoGoProtoNinOptUint16AsUint32 = func() gogoproto.Message { return &gogotest.NinOptUint16AsUint32{} }
)

var (
	govisNidOptUint16AsUint32 = &gotest.NidOptUint16AsUint32{
		Value: &uint16u32a,
	}
	gogovisNidOptUint16AsUint32 = &gogotest.NidOptUint16AsUint32{
		Value: uint16a,
	}
	govimNidOptUint16AsUint32 = &gotest.NidOptUint16AsUint32{
		Value: nil,
	}
	gozeroOptUint16AsUint32 = &gotest.NidOptUint16AsUint32{
		Value: &zerouint16u32,
	}
	gogovimNidOptUint16AsUint32 = &gogotest.NidOptUint16AsUint32{
		Value: 0,
	}
	govisNinOptUint16AsUint32 = &gotest.NinOptUint16AsUint32{
		Value: &uint16u32a,
	}
	gogovisNinOptUint16AsUint32 = &gogotest.NinOptUint16AsUint32{
		Value: &uint16a,
	}
	govimNinOptUint16AsUint32 = &gotest.NinOptUint16AsUint32{
		Value: nil,
	}
	gogovimNinOptUint16AsUint32 = &gogotest.NinOptUint16AsUint32{
		Value: nil,
	}
)

var (
	newGoProtoNidRepUint16AsUint32   = func() goproto.Message { return &gotest.NidRepUint16AsUint32{} }
	newGoGoProtoNidRepUint16AsUint32 = func() gogoproto.Message { return &gogotest.NidRepUint16AsUint32{} }
	newGoProtoNinRepUint16AsUint32   = func() goproto.Message { return &gotest.NinRepUint16AsUint32{} }
	newGoGoProtoNinRepUint16AsUint32 = func() gogoproto.Message { return &gogotest.NinRepUint16AsUint32{} }
)

var (
	govisNidRepUint16AsUint32 = &gotest.NidRepUint16AsUint32{
		Value: []uint32{uint16u32a, uint16u32b},
	}
	gogovisNidRepUint16AsUint32 = &gogotest.NidRepUint16AsUint32{
		Value: []custom.Uint16{uint16a, uint16b},
	}
	govimNidRepUint16AsUint32 = &gotest.NidRepUint16AsUint32{
		Value: nil,
	}
	gogovimNidRepUint16AsUint32 = &gogotest.NidRepUint16AsUint32{
		Value: nil,
	}
	govisNinRepUint16AsUint32 = &gotest.NinRepUint16AsUint32{
		Value: []uint32{uint16u32a, uint16u32b},
	}
	gogovisNinRepUint16AsUint32 = &gogotest.NinRepUint16AsUint32{
		Value: []custom.Uint16{uint16a, uint16b},
	}
	govimNinRepUint16AsUint32 = &gotest.NinRepUint16AsUint32{
		Value: nil,
	}
	gogovimNinRepUint16AsUint32 = &gogotest.NinRepUint16AsUint32{
		Value: nil,
	}
)

var (
	newGoProtoNidOptEnumUint32   = func() goproto.Message { return &gotest.NidOptEnumUint32{} }
	newGoGoProtoNidOptEnumUint32 = func() gogoproto.Message { return &gogotest.NidOptEnumUint32{} }
	newGoProtoNinOptEnumUint32   = func() goproto.Message { return &gotest.NinOptEnumUint32{} }
	newGoGoProtoNinOptEnumUint32 = func() gogoproto.Message { return &gogotest.NinOptEnumUint32{} }
)

var (
	govisNidOptEnumUint32 = &gotest.NidOptEnumUint32{
		Field1: gotest.TheTestEnumUint32_E.Enum(),
	}
	gogovisNidOptEnumUint32 = &gogotest.NidOptEnumUint32{
		Field1: gogotest.E,
	}
	govimNidOptEnumUint32 = &gotest.NidOptEnumUint32{
		Field1: nil,
	}
	gozeroNidOptEnumUint32 = &gotest.NidOptEnumUint32{
		Field1: gotest.TheTestEnumUint32_D.Enum(),
	}
	gogovimNidOptEnumUint32 = &gogotest.NidOptEnumUint32{
		Field1: gogotest.D,
	}
	govisNinOptEnumUint32 = &gotest.NinOptEnumUint32{
		Field1: gotest.TheTestEnumUint32_E.Enum(),
	}
	gogovisNinOptEnumUint32 = &gogotest.NinOptEnumUint32{
		Field1: gogotest.E.Enum(),
	}
	govimNinOptEnumUint32 = &gotest.NinOptEnumUint32{
		Field1: nil,
	}
	gogovimNinOptEnumUint32 = &gogotest.NinOptEnumUint32{
		Field1: nil,
	}
)

var (
	newGoProtoNidRepEnumUint32   = func() goproto.Message { return &gotest.NidRepEnumUint32{} }
	newGoGoProtoNidRepEnumUint32 = func() gogoproto.Message { return &gogotest.NidRepEnumUint32{} }
	newGoProtoNinRepEnumUint32   = func() goproto.Message { return &gotest.NinRepEnumUint32{} }
	newGoGoProtoNinRepEnumUint32 = func() gogoproto.Message { return &gogotest.NinRepEnumUint32{} }
)

var (
	govisNidRepEnumUint32 = &gotest.NidRepEnumUint32{
		Field1: []gotest.TheTestEnumUint32{gotest.TheTestEnumUint32_E, gotest.TheTestEnumUint32_F},
	}
	gogovisNidRepEnumUint32 = &gogotest.NidRepEnumUint32{
		Field1: []gogotest.TheTestEnumUint32{gogotest.E, gogotest.F},
	}
	govimNidRepEnumUint32 = &gotest.NidRepEnumUint32{
		Field1: nil,
	}
	gogovimNidRepEnumUint32 = &gogotest.NidRepEnumUint32{
		Field1: nil,
	}
	govisNinRepEnumUint32 = &gotest.NinRepEnumUint32{
		Field1: []gotest.TheTestEnumUint32{gotest.TheTestEnumUint32_E, gotest.TheTestEnumUint32_F},
	}
	gogovisNinRepEnumUint32 = &gogotest.NinRepEnumUint32{
		Field1: []gogotest.TheTestEnumUint32{gogotest.E, gogotest.F},
	}
	govimNinRepEnumUint32 = &gotest.NinRepEnumUint32{
		Field1: nil,
	}
	gogovimNinRepEnumUint32 = &gogotest.NinRepEnumUint32{
		Field1: nil,
	}
)

var (
	newGoProtoNidOptEnumUint16   = func() goproto.Message { return &gotest.NidOptEnumUint16{} }
	newGoGoProtoNidOptEnumUint16 = func() gogoproto.Message { return &gogotest.NidOptEnumUint16{} }
	newGoProtoNinOptEnumUint16   = func() goproto.Message { return &gotest.NinOptEnumUint16{} }
	newGoGoProtoNinOptEnumUint16 = func() gogoproto.Message { return &gogotest.NinOptEnumUint16{} }
)

var (
	govisNidOptEnumUint16 = &gotest.NidOptEnumUint16{
		Field1: gotest.TheTestEnumUint16_H.Enum(),
	}
	gogovisNidOptEnumUint16 = &gogotest.NidOptEnumUint16{
		Field1: gogotest.H,
	}
	govimNidOptEnumUint16 = &gotest.NidOptEnumUint16{
		Field1: nil,
	}
	gozeroNidOptEnumUint16 = &gotest.NidOptEnumUint16{
		Field1: gotest.TheTestEnumUint16_G.Enum(),
	}
	gogovimNidOptEnumUint16 = &gogotest.NidOptEnumUint16{
		Field1: gogotest.G,
	}
	govisNinOptEnumUint16 = &gotest.NinOptEnumUint16{
		Field1: gotest.TheTestEnumUint16_H.Enum(),
	}
	gogovisNinOptEnumUint16 = &gogotest.NinOptEnumUint16{
		Field1: gogotest.H.Enum(),
	}
	govimNinOptEnumUint16 = &gotest.NinOptEnumUint16{
		Field1: nil,
	}
	gogovimNinOptEnumUint16 = &gogotest.NinOptEnumUint16{
		Field1: nil,
	}
)

var (
	newGoProtoNidRepEnumUint16   = func() goproto.Message { return &gotest.NidRepEnumUint16{} }
	newGoGoProtoNidRepEnumUint16 = func() gogoproto.Message { return &gogotest.NidRepEnumUint16{} }
	newGoProtoNinRepEnumUint16   = func() goproto.Message { return &gotest.NinRepEnumUint16{} }
	newGoGoProtoNinRepEnumUint16 = func() gogoproto.Message { return &gogotest.NinRepEnumUint16{} }
)

var (
	govisNidRepEnumUint16 = &gotest.NidRepEnumUint16{
		Field1: []gotest.TheTestEnumUint16{gotest.TheTestEnumUint16_H, gotest.TheTestEnumUint16_I},
	}
	gogovisNidRepEnumUint16 = &gogotest.NidRepEnumUint16{
		Field1: []gogotest.TheTestEnumUint16{gogotest.H, gogotest.I},
	}
	govimNidRepEnumUint16 = &gotest.NidRepEnumUint16{
		Field1: nil,
	}
	gogovimNidRepEnumUint16 = &gogotest.NidRepEnumUint16{
		Field1: nil,
	}
	govisNinRepEnumUint16 = &gotest.NinRepEnumUint16{
		Field1: []gotest.TheTestEnumUint16{gotest.TheTestEnumUint16_H, gotest.TheTestEnumUint16_I},
	}
	gogovisNinRepEnumUint16 = &gogotest.NinRepEnumUint16{
		Field1: []gogotest.TheTestEnumUint16{gogotest.H, gogotest.I},
	}
	govimNinRepEnumUint16 = &gotest.NinRepEnumUint16{
		Field1: nil,
	}
	gogovimNinRepEnumUint16 = &gogotest.NinRepEnumUint16{
		Field1: nil,
	}
)

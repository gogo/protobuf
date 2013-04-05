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

type M interface {
	gogoproto.Marshaler
	gogoproto.Unmarshaler
}

func testBytes(t *testing.T, a M, astr []byte) {
	mastr, err := a.Marshal()
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(mastr, astr) {
		t.Fatalf("%v != %v", mastr, astr)
	}
	ptr := reflect.New(reflect.TypeOf(a).Elem())
	b := ptr.Interface().(M)
	if err := b.Unmarshal(mastr); err != nil {
		panic(err)
	}
	mastr, err = b.Marshal()
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

var uint128bytesa = []byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
var uint128bytesb = []byte{0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0}
var zerouint128bytes = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

func TestVarUint128Bytes(t *testing.T) {
	testBytes(t, &uint128a, uint128bytesa)
	testBytes(t, &uint128b, uint128bytesb)
	testBytes(t, &zerouint128, zerouint128bytes)
}

var (
	newGoProtoSimpleMsg = func() goproto.Message { return &gotest.SimpleMessage{} }
)

var gosimplemsg = &gotest.SimpleMessage{
	Id:         idbytes1,
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

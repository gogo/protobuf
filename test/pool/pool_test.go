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

package pool

import (
	"testing"

	"github.com/gogo/protobuf/mem"
	"github.com/gogo/protobuf/test/pool/poolfalse"
	"github.com/gogo/protobuf/test/pool/pooltrue"
)

const (
	numSliceElements = 1
	numMapElements   = 3
)

// The disabled functions should have the same performance as the
// false functions, this is just as verification. If this is correct,
// we can consolidate pooltrue and poolfalse into just pool, and do
// benchmarking just between enabled and disabled.

func BenchmarkPoolFalseWithoutMaps(b *testing.B) {
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolFalse(b, false)
	}
	b.StopTimer()
}

func BenchmarkPoolFalseWithoutMapsWithHeap(b *testing.B) {
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolFalseWithHeap(b, false)
	}
	b.StopTimer()
}

func BenchmarkPoolTrueDisabledWithoutMaps(b *testing.B) {
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolTrue(b, false, false)
	}
	b.StopTimer()
}

func BenchmarkPoolTrueDisabledWithSegListWithoutMaps(b *testing.B) {
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolTrue(b, false, true)
	}
	b.StopTimer()
}

func BenchmarkPoolTrueEnabledWithoutMapsSyncPool(b *testing.B) {
	mem.GlobalReset()
	mem.EnablePooling()
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolTrue(b, false, false)
	}
	b.StopTimer()
	mem.DisablePooling()
	checkAndPrintMemCounts(b)
}

func BenchmarkPoolTrueEnabledWithSegListWithoutMapsSyncPool(b *testing.B) {
	mem.GlobalReset()
	mem.EnablePooling()
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolTrue(b, false, true)
	}
	b.StopTimer()
	mem.DisablePooling()
	checkAndPrintMemCounts(b)
}

func BenchmarkPoolTrueEnabledWithoutMapsBufferedChannels(b *testing.B) {
	mem.SetBytesPoolChannelSize(256)
	mem.SetObjectPoolChannelSize(256)
	mem.SetBytesPoolNoSyncPool()
	mem.SetObjectPoolNoSyncPool()
	mem.GlobalReset()
	mem.EnablePooling()
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolTrue(b, false, false)
	}
	b.StopTimer()
	mem.DisablePooling()
	checkAndPrintMemCounts(b)
	mem.SetBytesPoolWithSyncPool()
	mem.SetObjectPoolWithSyncPool()
	mem.SetBytesPoolChannelSize(0)
	mem.SetObjectPoolChannelSize(0)
}

func BenchmarkPoolTrueEnabledWithSegListWithoutMapsBufferedChannels(b *testing.B) {
	mem.SetBytesPoolChannelSize(256)
	mem.SetObjectPoolChannelSize(256)
	mem.SetBytesPoolNoSyncPool()
	mem.SetObjectPoolNoSyncPool()
	mem.GlobalReset()
	mem.EnablePooling()
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolTrue(b, false, true)
	}
	b.StopTimer()
	mem.DisablePooling()
	checkAndPrintMemCounts(b)
	mem.SetBytesPoolWithSyncPool()
	mem.SetObjectPoolWithSyncPool()
	mem.SetBytesPoolChannelSize(0)
	mem.SetObjectPoolChannelSize(0)
}

func BenchmarkPoolFalseWithMaps(b *testing.B) {
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolFalse(b, true)
	}
	b.StopTimer()
}

func BenchmarkPoolFalseWithMapsWithHeap(b *testing.B) {
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolFalseWithHeap(b, true)
	}
	b.StopTimer()
}

func BenchmarkPoolTrueDisabledWithMaps(b *testing.B) {
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolTrue(b, true, false)
	}
	b.StopTimer()
}

func BenchmarkPoolTrueDisabledWithSegListAndMaps(b *testing.B) {
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolTrue(b, true, true)
	}
	b.StopTimer()
}

func BenchmarkPoolTrueEnabledWithMapsSyncPool(b *testing.B) {
	mem.GlobalReset()
	mem.EnablePooling()
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolTrue(b, true, false)
	}
	b.StopTimer()
	mem.DisablePooling()
	checkAndPrintMemCounts(b)
}

func BenchmarkPoolTrueEnabledWithSegListAndMapsSyncPool(b *testing.B) {
	mem.GlobalReset()
	mem.EnablePooling()
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolTrue(b, true, true)
	}
	b.StopTimer()
	mem.DisablePooling()
	checkAndPrintMemCounts(b)
}

func BenchmarkPoolTrueEnabledWithMapsBufferedChannels(b *testing.B) {
	mem.SetBytesPoolChannelSize(256)
	mem.SetObjectPoolChannelSize(256)
	mem.SetBytesPoolNoSyncPool()
	mem.SetObjectPoolNoSyncPool()
	mem.GlobalReset()
	mem.EnablePooling()
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolTrue(b, true, false)
	}
	b.StopTimer()
	mem.DisablePooling()
	checkAndPrintMemCounts(b)
	mem.SetBytesPoolWithSyncPool()
	mem.SetObjectPoolWithSyncPool()
	mem.SetBytesPoolChannelSize(0)
	mem.SetObjectPoolChannelSize(0)
}

func BenchmarkPoolTrueEnabledWithSegListAndMapsBufferedChannels(b *testing.B) {
	mem.SetBytesPoolChannelSize(256)
	mem.SetObjectPoolChannelSize(256)
	mem.SetBytesPoolNoSyncPool()
	mem.SetObjectPoolNoSyncPool()
	mem.GlobalReset()
	mem.EnablePooling()
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolTrue(b, true, true)
	}
	b.StopTimer()
	mem.DisablePooling()
	checkAndPrintMemCounts(b)
	mem.SetBytesPoolWithSyncPool()
	mem.SetObjectPoolWithSyncPool()
	mem.SetBytesPoolChannelSize(0)
	mem.SetObjectPoolChannelSize(0)
}

func benchmarkPoolFalse(b *testing.B, withMaps bool) {
	foo := &poolfalse.Foo{}
	foo.One = 1
	foo.Two = 2
	foo.Bar = &poolfalse.Bar{}
	foo.Bar.OneBar = 10
	foo.Bar.TwoBar = 20
	for i := 0; i < numSliceElements; i++ {
		barElem := &poolfalse.Bar{}
		barElem.OneBar = 100
		barElem.TwoBar = 200
		foo.RepeatedBar = append(foo.RepeatedBar, barElem)
	}
	if withMaps {
		foo.MapBar = make(map[uint32]*poolfalse.Bar, numMapElements)
		for i := 0; i < numMapElements; i++ {
			barElem := &poolfalse.Bar{}
			barElem.OneBar = 1000
			barElem.TwoBar = 2000
			foo.MapBar[uint32(i)] = barElem
		}
	}

	data, err := foo.Marshal()
	if err != nil {
		b.Fatal(err)
	}

	unmarshaledFoo := &poolfalse.Foo{}
	if err := unmarshaledFoo.Unmarshal(data); err != nil {
		b.Fatal(err)
	}
}

// we trick Golang into putting all elements we create onto the heap by returning the two
// Foos we create (that reference all the Bars) to more closely resemble real-world usage of Protobuf structs
func benchmarkPoolFalseWithHeap(b *testing.B, withMaps bool) (*poolfalse.Foo, *poolfalse.Foo) {
	foo := &poolfalse.Foo{}
	foo.One = 1
	foo.Two = 2
	foo.Bar = &poolfalse.Bar{}
	foo.Bar.OneBar = 10
	foo.Bar.TwoBar = 20
	for i := 0; i < numSliceElements; i++ {
		barElem := &poolfalse.Bar{}
		barElem.OneBar = 100
		barElem.TwoBar = 200
		foo.RepeatedBar = append(foo.RepeatedBar, barElem)
	}
	if withMaps {
		foo.MapBar = make(map[uint32]*poolfalse.Bar, numMapElements)
		for i := 0; i < numMapElements; i++ {
			barElem := &poolfalse.Bar{}
			barElem.OneBar = 1000
			barElem.TwoBar = 2000
			foo.MapBar[uint32(i)] = barElem
		}
	}

	data, err := foo.Marshal()
	if err != nil {
		b.Fatal(err)
	}

	unmarshaledFoo := &poolfalse.Foo{}
	if err := unmarshaledFoo.Unmarshal(data); err != nil {
		b.Fatal(err)
	}
	return foo, unmarshaledFoo
}

func benchmarkPoolTrue(b *testing.B, withMaps bool, withSegList bool) {
	foo := pooltrue.GetFoo()
	foo.One = 1
	foo.Two = 2
	foo.Bar = pooltrue.GetBar()
	foo.Bar.OneBar = 10
	foo.Bar.TwoBar = 20
	for i := 0; i < numSliceElements; i++ {
		barElem := pooltrue.GetBar()
		barElem.OneBar = 100
		barElem.TwoBar = 200
		foo.RepeatedBar = append(foo.RepeatedBar, barElem)
	}
	if withMaps {
		foo.MapBar = make(map[uint32]*pooltrue.Bar, numMapElements)
		for i := 0; i < numMapElements; i++ {
			barElem := pooltrue.GetBar()
			barElem.OneBar = 1000
			barElem.TwoBar = 2000
			foo.MapBar[uint32(i)] = barElem
		}
	}

	var data []byte
	var bytes *mem.Bytes
	var err error
	if withSegList {
		bytes, err = foo.MarshalPool()
		if err != nil {
			b.Fatal(err)
		}
		data = bytes.Value()
	} else {
		data, err = foo.Marshal()
		if err != nil {
			b.Fatal(err)
		}
	}

	unmarshaledFoo := pooltrue.GetFoo()
	if err := unmarshaledFoo.Unmarshal(data); err != nil {
		b.Fatal(err)
	}

	foo.Recycle()
	unmarshaledFoo.Recycle()
	bytes.Recycle()
}

func checkAndPrintMemCounts(tb testing.TB) {
	newCount := mem.NewCount()
	getCount := mem.GetCount()
	unrecycledCount := mem.UnrecycledCount()
	tb.Logf("MemNewCount: %d MemGetCount: %d MemUnrecycledCount: %d", newCount, getCount, unrecycledCount)
	if unrecycledCount != 0 {
		tb.Fatalf("Expected all memory to be recycled but had %d unrecycled objects", unrecycledCount)
	}
}

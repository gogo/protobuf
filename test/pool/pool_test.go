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

// The disabled functions should have the same performance as the
// false functions, this is just as verification. If this is correct,
// we can consolidate pooltrue and poolfalse into just pool, and do
// benchmarking just between enabled and disabled.

func BenchmarkPoolFalse(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolFalse(b, false)
	}
}

func BenchmarkPoolTrueDisabled(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolTrue(b, false, false)
	}
}

func BenchmarkPoolTrueDisabledWithSegList(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolTrue(b, false, true)
	}
}

func BenchmarkPoolTrueEnabled(b *testing.B) {
	mem.Global().Enable()
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolTrue(b, false, false)
	}
	b.StopTimer()
	mem.Global().Disable()
}

func BenchmarkPoolTrueEnabledWithSegList(b *testing.B) {
	mem.Global().Enable()
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolTrue(b, false, true)
	}
	b.StopTimer()
	mem.Global().Disable()
}

func BenchmarkPoolFalseWithMaps(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolFalse(b, true)
	}
}

func BenchmarkPoolTrueDisabledWithMaps(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolTrue(b, true, false)
	}
}

func BenchmarkPoolTrueDisabledWithSegListAndMaps(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolTrue(b, true, true)
	}
}

func BenchmarkPoolTrueEnabledWithMaps(b *testing.B) {
	mem.Global().Enable()
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolTrue(b, true, false)
	}
	b.StopTimer()
	mem.Global().Disable()
}

func BenchmarkPoolTrueEnabledWithSegListAndMaps(b *testing.B) {
	mem.Global().Enable()
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		benchmarkPoolTrue(b, true, true)
	}
	b.StopTimer()
	mem.Global().Disable()
}

func benchmarkPoolFalse(b *testing.B, withMaps bool) {
	foo := &poolfalse.Foo{}
	foo.One = 1
	foo.Two = 2
	if withMaps {
		foo.MapBar = make(map[uint32]*poolfalse.Bar)
	}
	foo.Bar = &poolfalse.Bar{}
	foo.Bar.OneBar = 10
	foo.Bar.TwoBar = 20
	for i := 0; i < 10; i++ {
		barElem := &poolfalse.Bar{}
		barElem.OneBar = 100
		barElem.TwoBar = 200
		foo.RepeatedBar = append(foo.RepeatedBar, barElem)
		if withMaps {
			barElem = &poolfalse.Bar{}
			barElem.OneBar = 1000
			barElem.TwoBar = 2000
			foo.MapBar[uint32(i)] = barElem
		}
	}

	// this is going to allocate the byte slice
	// if we had a byte slice pool, we could take this out of the benchmark
	// in terms of memory
	data, err := foo.Marshal()
	if err != nil {
		b.Fatal(err)
	}

	unmarshaledFoo := &poolfalse.Foo{}
	if err := unmarshaledFoo.Unmarshal(data); err != nil {
		b.Fatal(err)
	}
}

func benchmarkPoolTrue(b *testing.B, withMaps bool, withSegList bool) {
	foo := pooltrue.GetFoo()

	foo.One = 1
	foo.Two = 2
	if withMaps {
		foo.MapBar = make(map[uint32]*pooltrue.Bar)
	}
	foo.Bar = pooltrue.GetBar()
	foo.Bar.OneBar = 10
	foo.Bar.TwoBar = 20
	for i := 0; i < 10; i++ {
		barElem := pooltrue.GetBar()
		barElem.OneBar = 100
		barElem.TwoBar = 200
		foo.RepeatedBar = append(foo.RepeatedBar, barElem)
		if withMaps {
			barElem = pooltrue.GetBar()
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

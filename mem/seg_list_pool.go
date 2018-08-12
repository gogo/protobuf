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

package mem

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type segListPool struct {
	allocateLen int
	syncPool    *sync.Pool
	c           chan *Bytes
}

func newSegListPool(channelSize uint16, noSyncPool bool, allocateLen int) *segListPool {
	segListPool := &segListPool{allocateLen: allocateLen}
	// we check that at most one of channelSize == 0 and noSyncPool are true
	// so we do not have to check where both m.c == nil and m.syncPool == nil below
	if !noSyncPool {
		segListPool.syncPool = &sync.Pool{
			New: func() interface{} {
				atomic.AddUint64(&globalNewCount, 1)
				return newBytes(segListPool, allocateLen)
			},
		}
	}
	if channelSize > 0 {
		segListPool.c = make(chan *Bytes, channelSize)
	}
	return segListPool
}

func (m *segListPool) get(valueLen int) *Bytes {
	atomic.AddUint64(&globalGetCount, 1)
	if m.allocateLen < valueLen {
		panic(fmt.Sprintf("segListPool got request for size %d but has size %d", valueLen, m.allocateLen))
	}
	if m.c == nil {
		getBytes := m.syncPool.Get().(*Bytes)
		// this is only used for recycled or not
		// the existence of segListPool is used to denote whether this was allocated by a pool
		getBytes.poolMarker = PoolMarkerNone
		getBytes.valueLen = 0
		return getBytes
	}
	if m.syncPool == nil {
		select {
		case bytes := <-m.c:
			bytes.poolMarker = PoolMarkerNone
			bytes.valueLen = 0
			return bytes
		default:
			atomic.AddUint64(&globalNewCount, 1)
			return newBytes(m, m.allocateLen)
		}
	}
	select {
	case bytes := <-m.c:
		bytes.poolMarker = PoolMarkerNone
		bytes.valueLen = 0
		return bytes
	default:
		getBytes := m.syncPool.Get().(*Bytes)
		getBytes.poolMarker = PoolMarkerNone
		getBytes.valueLen = 0
		return getBytes
	}
}

func (m *segListPool) put(bytes *Bytes) {
	atomic.AddUint64(&globalRecycledCount, 1)
	if m.c == nil {
		m.syncPool.Put(bytes)
		return
	}
	if m.syncPool == nil {
		select {
		case m.c <- bytes:
			return
		default:
			return
		}
	}
	select {
	case m.c <- bytes:
		return
	default:
		m.syncPool.Put(bytes)
		return
	}
}

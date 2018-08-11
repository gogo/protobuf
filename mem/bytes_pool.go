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
	"sort"
)

// DefaultBytesPoolChannelSize is the default BytesPool channel size.
//
// See the BytesPoolWithChannelSize function for more details.
const DefaultBytesPoolChannelSize uint16 = 16

// DefaultBytesPoolSegListSizes are the default seg list sizes.
//
// See the WithSegListSizes function for more details.
var DefaultBytesPoolSegListSizes = []int{
	8,
	16,
	64,
	256,
	4096,
	65536,
	1048576,
	16777216,
}

// BytesPool is a pool of Bytes.
//
// Users should generally not interact with BytesPools directly, instead
// relying on the generated functions to call BytesPools.
type BytesPool struct {
	channelSize              uint16
	sortedSegListSizes       []int
	segListSizeToSegListPool map[int]*segListPool
}

// BytesPoolOption is an option when constructing a new BytesPool.
type BytesPoolOption func(*BytesPool)

// BytesPoolWithChannelSize returns a BytesPoolOption that sets the
// BytesPool channel size.
//
// If your workflow usually only has one object of a given type allocated,
// the garbage collector will generally garbage collect the object if it
// is in a sync.Pool before it can be allocated again from the Pool. If this
// option is used, a channel of allocated Bytes of this buffered size per seg
// list will  be used to recycle Bytes before using the sync.Pool.
//
// By default, use DefaultBytesPoolChannelSize. Set this to 0 to not use any channel.
func BytesPoolWithChannelSize(channelSize uint16) BytesPoolOption {
	return func(bytesPool *BytesPool) {
		bytesPool.channelSize = channelSize
	}
}

// BytesPoolWithSegListSizes returns a BytesPoolOption that sets the seg list sizes
// for allocating Bytes.
//
// By default, use DefaultSegListSizes.
func BytesPoolWithSegListSizes(segListSizes ...int) BytesPoolOption {
	return func(bytesPool *BytesPool) {
		segListSizesCopy := make([]int, len(segListSizes))
		copy(segListSizesCopy, segListSizes)
		sort.Ints(segListSizesCopy)
		bytesPool.sortedSegListSizes = segListSizesCopy
	}
}

// NewBytesPool creates a new BytesPool.
func NewBytesPool(options ...BytesPoolOption) *BytesPool {
	bytesPool := &BytesPool{
		channelSize:              DefaultBytesPoolChannelSize,
		sortedSegListSizes:       DefaultBytesPoolSegListSizes,
		segListSizeToSegListPool: make(map[int]*segListPool),
	}
	for _, option := range options {
		option(bytesPool)
	}
	for _, segListSize := range bytesPool.sortedSegListSizes {
		bytesPool.segListSizeToSegListPool[segListSize] = newSegListPool(bytesPool.channelSize, segListSize)
	}
	return bytesPool
}

// Get gets a Bytes of the given length.
//
// This Bytes can be written to by accessing the backing Value.
//
// Note that there may be existing data inside the Value. If you need the Value
// to be zeroed out, use the MemsetZero function.
func (p *BytesPool) Get(valueLen int) *Bytes {
	for _, segListSize := range p.sortedSegListSizes {
		if valueLen <= segListSize {
			segListPool, ok := p.segListSizeToSegListPool[segListSize]
			if !ok {
				panic(fmt.Sprintf("no segListPool of size %d", segListSize))
			}
			return segListPool.get(valueLen)
		}
	}
	// no seg list big enough, make a Bytes without a backing pool
	return newBytes(nil, valueLen)
}

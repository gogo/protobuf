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
	"sync"

	"github.com/gogo/protobuf/proto"
)

// DefaultChannelSize is the default Pool channel size.
//
// See the WithChannelSize function for more details.
const DefaultChannelSize uint16 = 4

// DefaultSegListSizes are the default seg list sizes.
//
// See the WithSegListSizes function for more details.
var DefaultSegListSizes = []int{
	8,
	16,
	64,
	256,
	4096,
	65536,
	1048576,
	16777216,
}

// Bytes represents a byte slice created from a Pool.
//
// Only create these from Pools.
type Bytes struct {
	segListPool *segListPool
	value       []byte
	valueLen    int
}

// NewUnmanagedBytes creates a new Bytes that is not backed by any Pool.
//
// Users should generally not call this directly.
func NewUnmanagedBytes(valueLen int) *Bytes {
	return &Bytes{
		value:    make([]byte, valueLen),
		valueLen: valueLen,
	}
}

// Value gets the value.
func (b *Bytes) Value() []byte {
	return b.value[:b.valueLen]
}

// Len gets the length.
//
// This is equivalent to len(b.Value()).
func (b *Bytes) Len() int {
	return b.valueLen
}

// Truncate truncates the value to the given length.
func (b *Bytes) Truncate(valueLen int) {
	b.valueLen = valueLen
}

// MemsetZero sets the data inside the value to zero.
//
// Optimized per https://golang.org/cl/137880043
func (b *Bytes) MemsetZero() {
	for i := range b.value {
		b.value[i] = 0
	}
}

// Recycle recycles the Bytes.
//
// If this is nil, this is a no-op.
func (b *Bytes) Recycle() {
	if b == nil {
		return
	}
	// note this is possible if we did not have a big enough seg list
	if b.segListPool != nil {
		b.segListPool.put(b)
	}
}

// PooledMessage is a proto.Message that can be pooled.
type PooledMessage interface {
	proto.Message

	Recycle()
}

// Pool is a pool of PooledMessages and Bytes.
type Pool struct {
	channelSize        uint16
	sortedSegListSizes []int

	messageTypeToMessagePool map[string]*messagePool
	segListSizeToSegListPool map[int]*segListPool

	lock *sync.RWMutex
}

// PoolOption is an option when constructing a new Pool.
type PoolOption func(*Pool)

// WithChannelSize returns a PoolOption that sets the Pool channel size.
//
// If your workflow usually only has one object of a given type allocated,
// the garbage collector will generally garbage collect the object if it
// is in a sync.Pool before it can be allocated again from the Pool. If this
// option is used, a channel of allocated PooledMessages of this buffered
// size will be used to recycle PooledMessages before using the sync.Pool.
//
//
// By default, use DefaultChannelSize. Set this to 0 to not use any channel.
func WithChannelSize(channelSize uint16) PoolOption {
	return func(pool *Pool) {
		pool.channelSize = channelSize
	}
}

// WithSegListSizes returns a PoolOption that sets the seg list sizes
// for allocating Bytes.
//
// By default, use DefaultSegListSizes.
func WithSegListSizes(segListSizes ...int) PoolOption {
	return func(pool *Pool) {
		segListSizesCopy := make([]int, len(segListSizes))
		copy(segListSizesCopy, segListSizes)
		sort.Ints(segListSizesCopy)
		pool.sortedSegListSizes = segListSizesCopy
	}
}

// NewPool creates a new Pool.
func NewPool(options ...PoolOption) *Pool {
	pool := &Pool{
		channelSize:              DefaultChannelSize,
		sortedSegListSizes:       DefaultSegListSizes,
		messageTypeToMessagePool: make(map[string]*messagePool),
		segListSizeToSegListPool: make(map[int]*segListPool),
		lock: &sync.RWMutex{},
	}
	for _, option := range options {
		option(pool)
	}
	for _, segListSize := range pool.sortedSegListSizes {
		pool.segListSizeToSegListPool[segListSize] = newSegListPool(pool, segListSize)
	}
	return pool
}

// RegisterConstructor registers the given constructor for the given message type.
//
// This will not register the constructor if called twice for the same message type.
//
// Users should not call this function directly, instead calling the generated functions
// for each file.
func (p *Pool) RegisterConstructor(messageType string, constructor func(*Pool) PooledMessage) {
	// TODO: If we require all constructors to be registered at initialization, we can
	// get rid of this locking. It's probably preferable to have the locking as this
	// provides a cheap guarantee (the read locking below has very little overhead).
	p.lock.Lock()
	if _, ok := p.messageTypeToMessagePool[messageType]; !ok {
		p.messageTypeToMessagePool[messageType] = newMessagePool(p, constructor)
	}
	p.lock.Unlock()
}

// Get gets a reset PooledMessage for the given message type.
//
// Returns nil if no constructor was registered for the PooledMessage.
//
// Users should not call this function directly, instead calling the generated functions
// for each file.
func (p *Pool) Get(messageType string) PooledMessage {
	p.lock.RLock()
	messagePool, ok := p.messageTypeToMessagePool[messageType]
	p.lock.RUnlock()
	if !ok {
		return nil
	}
	return messagePool.get()
}

// Put puts a PooledMessage back into the Pool.
//
// This PooledMessage must have been created by this Pool or this will panic.
//
// Users should call the generated Recycle function instead of calling this directly.
func (p *Pool) Put(messageType string, message PooledMessage) {
	p.lock.RLock()
	messagePool, ok := p.messageTypeToMessagePool[messageType]
	p.lock.RUnlock()
	if !ok {
		panic(fmt.Sprintf("message type %s was not registered with this Pool", messageType))
	}
	messagePool.put(message)
}

// GetBytes gets a Bytes of the given length.
//
// This Bytes can be written to by accessing the backing Value.
//
// Note that there may be existing data inside the Value. If you need the Value
// to be zeroed out, use the MemsetZero function.
func (p *Pool) GetBytes(valueLen int) *Bytes {
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
	return NewUnmanagedBytes(valueLen)
}

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
	"sync/atomic"
)

// DefaultChannelSize is the default Pool channel size.
//
// See the WithChannelSize function for more details.
const DefaultChannelSize uint16 = 128

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

var globalPool = NewPool(WithNoLocking(), WithDisabled())

// Global returns the global Pool instance.
//
// It is generally preferable to create a Pool instance for your application and
// pass it through to the places it is needed, however this will prove difficult for
// existing applications wishing to migrate to the Pool API.
//
// The Global Pool is disabled by default. To enable pooling in your application using
// the Get* generated functions for each message, call Global().Enable() at initialization.
func Global() *Pool {
	return globalPool
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

// PooledMessage is a message that can be pooled.
//
// This generally is meant for proto.Messages but is compatible with
// anything that has a Reset and Recycle method.
type PooledMessage interface {
	Reset()
	Recycle()
}

// Pool is a pool of PooledMessages and Bytes.
type Pool struct {
	enabled            uint32
	channelSize        uint16
	sortedSegListSizes []int

	messageTypeToMessagePool map[string]*MessagePool
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

// WithNoLocking returns a PoolOption that turns off locking around message constructor
// registration. If this option is turned on, you must make all RegisterConstructor calls
// at initialization, and not call RegisterConstructor afterwards.
//
// Generally, you should use this option as it provides a significant performance benefit,
// just make sure to only call registration functions at initialization.
//
// This option is called for the Global Pool.
func WithNoLocking() PoolOption {
	return func(pool *Pool) {
		pool.lock = nil
	}
}

// WithDisabled returns a PoolOption that disables the Pool at construction time.
//
// This is called on the global pool to make sure there are no ordering issues
// with other initialization functions.
func WithDisabled() PoolOption {
	return func(pool *Pool) {
		pool.enabled = 0
	}
}

// NewPool creates a new Pool.
func NewPool(options ...PoolOption) *Pool {
	pool := &Pool{
		enabled:                  1,
		channelSize:              DefaultChannelSize,
		sortedSegListSizes:       DefaultSegListSizes,
		messageTypeToMessagePool: make(map[string]*MessagePool),
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

// Enable enables all pooling.
func (p *Pool) Enable() {
	atomic.StoreUint32(&p.enabled, 1)
}

// Disable disables all pooling.
//
// This results in Get and GetBytes calls returning nil messages and unmanaged Bytes.
func (p *Pool) Disable() {
	atomic.StoreUint32(&p.enabled, 0)
}

// RegisterConstructor registers the given constructor for the given message type.
//
// This will not register the constructor if called twice for the same message type.
//
// Users should not call this function directly, instead calling the generated functions
// for each file.
func (p *Pool) RegisterConstructor(messageType string, constructor func(*Pool) PooledMessage) {
	if p.lock != nil {
		p.lock.Lock()
	}
	if _, ok := p.messageTypeToMessagePool[messageType]; !ok {
		p.messageTypeToMessagePool[messageType] = newMessagePool(p, constructor)
	}
	if p.lock != nil {
		p.lock.Unlock()
	}
}

// GetMessagePool gets a registered message pool.
//
// If there is no message pool for the message type, this will panic.
// This is an optimization so that no map access is needed for generated global functions.
//
// Users should never call this directly.
func (p *Pool) GetMessagePool(messageType string) *MessagePool {
	if p.lock != nil {
		p.lock.Lock()
	}
	messagePool, ok := p.messageTypeToMessagePool[messageType]
	if p.lock != nil {
		p.lock.Unlock()
	}
	if !ok {
		panic(fmt.Sprintf("no message pool for message type %s", messageType))
	}
	return messagePool
}

// Get gets a reset PooledMessage for the given message type.
//
// Returns nil if no constructor was registered for the PooledMessage, or if pooling is disabled.
//
// Users should not call this function directly, instead calling the generated functions
// for each file.
func (p *Pool) Get(messageType string) PooledMessage {
	if p.isDisabled() {
		return nil
	}
	if p.lock != nil {
		p.lock.RLock()
	}
	messagePool, ok := p.messageTypeToMessagePool[messageType]
	if p.lock != nil {
		p.lock.RUnlock()
	}
	if !ok {
		return nil
	}
	return messagePool.Get()
}

// Put puts a PooledMessage back into the Pool.
//
// This PooledMessage must have been created by this Pool or this will panic.
//
// Users should call the generated Recycle function instead of calling this directly.
func (p *Pool) Put(messageType string, message PooledMessage) {
	if p.isDisabled() {
		return
	}
	if p.lock != nil {
		p.lock.RLock()
	}
	messagePool, ok := p.messageTypeToMessagePool[messageType]
	if p.lock != nil {
		p.lock.RUnlock()
	}
	if !ok {
		panic(fmt.Sprintf("message type %s was not registered with this Pool", messageType))
	}
	messagePool.Put(message)
}

// GetBytes gets a Bytes of the given length.
//
// This Bytes can be written to by accessing the backing Value.
//
// Note that there may be existing data inside the Value. If you need the Value
// to be zeroed out, use the MemsetZero function.
func (p *Pool) GetBytes(valueLen int) *Bytes {
	if p.isDisabled() {
		return NewUnmanagedBytes(valueLen)
	}
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

func (p *Pool) isDisabled() bool {
	return atomic.LoadUint32(&p.enabled) == 0
}

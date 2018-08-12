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
	"sync"
)

// DefaultObjectPoolChannelSize is the default ObjectPool channel size.
//
// See the ObjectPoolWithChannelSize function for more details.
const DefaultObjectPoolChannelSize uint16 = 0

var globalObjectPool []ObjectPool

// Object is something that can be pooled with an ObjectPool.
type Object interface {
	// Reset the object. When constructed, an Object should already be reset.
	Reset()
}

// ObjectPool is a pool for an individual object type.
//
// Users should generally not interact with ObjectPools directly, instead
// relying on the generated functions to call ObjectPools.
type ObjectPool struct {
	channelSize uint16
	noSyncPool  bool
	constructor func() Object
	syncPool    *sync.Pool
	c           chan Object
}

// ObjectPoolOption is an option when constructing a new ObjectPool.
type ObjectPoolOption func(*ObjectPool)

// ObjectPoolWithChannelSize returns a ObjectPoolOption that sets the
// ObjectPool channel size.
//
// If your workflow usually only has one object of a given type allocated,
// the garbage collector will generally garbage collect the object if it
// is in a sync.Pool before it can be allocated again from the Pool. If this
// option is used, a channel of allocated Objects of this buffered size per seg
// list will  be used to recycle Objects before using the sync.Pool.
//
// By default, use DefaultObjectPoolChannelSize. Set this to 0 to not use any channel.
func ObjectPoolWithChannelSize(channelSize uint16) ObjectPoolOption {
	return func(objectPool *ObjectPool) {
		objectPool.channelSize = channelSize
	}
}

// ObjectPoolWithNoSyncPool returns a ObjectPoolOption that will not have a backing
// sync.Pool, instead discarding elements that cannot be put back into the channel.
//
// By default, this is false.
//
// Note if you set this and your channel size is 0 (which is the default value),
// you would not be doing any pooling at all, so the constructor will panic
// if both are set.
func ObjectPoolWithNoSyncPool() ObjectPoolOption {
	return func(objectPool *ObjectPool) {
		objectPool.noSyncPool = true
	}
}

// NewObjectPool creates a new ObjectPool for the given constructor and options.
func NewObjectPool(constructor func() Object, options ...ObjectPoolOption) *ObjectPool {
	objectPool := &ObjectPool{
		channelSize: DefaultObjectPoolChannelSize,
		constructor: constructor,
	}
	for _, option := range options {
		option(objectPool)
	}
	if objectPool.channelSize == 0 && objectPool.noSyncPool {
		panic("cannot have both a channel size of 0 and no sync pool")
	}
	if !objectPool.noSyncPool {
		objectPool.syncPool = &sync.Pool{
			New: func() interface{} {
				return constructor()
			},
		}
	}
	if objectPool.channelSize > 0 {
		objectPool.c = make(chan Object, objectPool.channelSize)
	}
	return objectPool
}

// Get returns a pooled object.
func (m *ObjectPool) Get() Object {
	if m.c == nil {
		getObject := m.syncPool.Get().(Object)
		getObject.Reset()
		return getObject
	}
	if m.syncPool == nil {
		select {
		case object := <-m.c:
			object.Reset()
			return object
		default:
			return m.constructor()
		}
	}
	select {
	case object := <-m.c:
		object.Reset()
		return object
	default:
		getObject := m.syncPool.Get().(Object)
		getObject.Reset()
		return getObject
	}
}

// Put puts a pooled object back into the object pool.
func (m *ObjectPool) Put(object Object) {
	if m.c == nil {
		m.syncPool.Put(object)
		return
	}
	if m.syncPool == nil {
		select {
		case m.c <- object:
			return
		default:
			return
		}
	}
	select {
	case m.c <- object:
		return
	default:
		m.syncPool.Put(object)
		return
	}
}

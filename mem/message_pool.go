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

import "sync"

// MessagePool is a a pool for an individual message type.
//
// These should only be created from a Pool.
type MessagePool struct {
	pool     *Pool
	syncPool *sync.Pool
	c        chan PooledMessage
}

func newMessagePool(pool *Pool, constructor func(*Pool) PooledMessage) *MessagePool {
	messagePool := &MessagePool{pool: pool}
	messagePool.syncPool = &sync.Pool{
		New: func() interface{} {
			return constructor(pool)
		},
	}
	if pool.channelSize > 0 {
		messagePool.c = make(chan PooledMessage, pool.channelSize)
	}
	return messagePool
}

// Get returns a pooled message.
//
// If the Pool that created this MessagePool is disabled, this will return nil.
func (m *MessagePool) Get() PooledMessage {
	if m.pool.isDisabled() {
		return nil
	}
	if m.c == nil {
		getMessage := m.syncPool.Get().(PooledMessage)
		getMessage.Reset()
		return getMessage
	}
	select {
	case message := <-m.c:
		message.Reset()
		return message
	default:
		getMessage := m.syncPool.Get().(PooledMessage)
		getMessage.Reset()
		return getMessage
	}
}

// Put puts a pooled message back into the message pool.
//
// If the Pool that created this MessagePool is disabled, this will return nil.
func (m *MessagePool) Put(message PooledMessage) {
	if m.pool.isDisabled() {
		return
	}
	if m.c == nil {
		m.syncPool.Put(message)
		return
	}
	select {
	case m.c <- message:
		return
	default:
		m.syncPool.Put(message)
		return
	}
}

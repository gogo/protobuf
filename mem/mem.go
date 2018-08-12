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

import "sync/atomic"

const (
	// PoolMarkerNone denots no pool marker.
	PoolMarkerNone PoolMarker = 0x0
	// PoolMarkerAllocatedByPool says that a message was allocated by a pool.
	PoolMarkerAllocatedByPool PoolMarker = 0x1
	// PoolMarkerRecycled says that a message has been recycled but not reset.
	PoolMarkerRecycled PoolMarker = 0x2

	// PanicUseAfterRecycle is the panic message used if a message is used after Recycle has been called.
	PanicUseAfterRecycle = "useAfterRecycle"
	// PanicDoubleRecycle is the panic message used if a message is Recycled twice.
	PanicDoubleRecycle = "doubleRecycle"
)

var (
	// one of the big benefits of globals here is we can cut functions
	// short without having to make a lot of function calls
	globalEnabled               bool
	globalBytesPoolChannelSize  = DefaultBytesPoolChannelSize
	globalBytesPoolNoSyncPool   = false
	globalBytesPoolSegListSizes = DefaultBytesPoolSegListSizes
	globalObjectPoolChannelSize = DefaultObjectPoolChannelSize
	globalObjectPoolNoSyncPool  = false

	globalBytesPool         = NewBytesPool()
	globalObjectPoolGetters []func() *ObjectPool
	globalObjectPoolSetters []func(...ObjectPoolOption)

	globalNewCount      uint64
	globalGetCount      uint64
	globalRecycledCount uint64
)

// PoolMarker is a pool marker.
//
// TODO: should this be a uint32 for better probability of word-alignment?
type PoolMarker uint8

// GlobalReset resets all counts, and makes new pools.
func GlobalReset() {
	globalNewCount = 0
	globalGetCount = 0
	globalRecycledCount = 0
	resetGlobalBytesPool()
	resetGlobalObjectPools()
}

// EnablePooling enables pooling.
//
// Pooling is disabled by default.
//
// This should only be called at initialization.
func EnablePooling() {
	globalEnabled = true
}

// DisablePooling disables pooling.
//
// Pooling is disabled by default.
//
// This should only be called at initialization.
func DisablePooling() {
	globalEnabled = false
}

// NewCount returns the global number of Bytes and Objects created.
func NewCount() uint64 {
	return atomic.LoadUint64(&globalNewCount)
}

// GetCount returns the global number of Bytes and Objects gotten.
func GetCount() uint64 {
	return atomic.LoadUint64(&globalGetCount)
}

// UnrecycledCount returns the global number of Bytes and Objects unrecycled.
func UnrecycledCount() uint64 {
	return GetCount() - atomic.LoadUint64(&globalRecycledCount)
}

// SetBytesPoolChannelSize results in the global BytesPool called by GetBytes
// being set to a new BytesPool with the BytesPoolWithChannelSize option
// set to the given value.
// Note you need to set this before calling SetBytesPoolNoSyncPool.
func SetBytesPoolChannelSize(channelSize uint16) {
	globalBytesPoolChannelSize = channelSize
	resetGlobalBytesPool()
}

// SetBytesPoolNoSyncPool results in the global BytesPools being set to a
// new BytesPool with the BytesPoolWithNoSyncPool option set.
func SetBytesPoolNoSyncPool() {
	globalBytesPoolNoSyncPool = true
	resetGlobalBytesPool()
}

// SetBytesPoolWithSyncPool results in the global BytesPools being set to a
// new BytesPool with the BytesPoolWithNoSyncPool option unset.
//
// This is set by default.
func SetBytesPoolWithSyncPool() {
	globalBytesPoolNoSyncPool = false
	resetGlobalBytesPool()
}

// SetBytesPoolSegListSizes results in the global BytesPool called by GetBytes
// being set to a new BytesPool with the BytesPoolWithSegListSizes option
// set to the given value.
func SetBytesPoolSegListSizes(seglistSizes ...int) {
	globalBytesPoolSegListSizes = seglistSizes
	resetGlobalBytesPool()
}

// SetObjectPoolChannelSize results in the global ObjectPools being set to a
// new ObjectPool with the ObjectPoolWithChannelSize option set to the given value.
// Note you need to set this before calling SetObjectPoolNoSyncPool.
func SetObjectPoolChannelSize(channelSize uint16) {
	globalObjectPoolChannelSize = channelSize
	resetGlobalObjectPools()
}

// SetObjectPoolNoSyncPool results in the global ObjectPools being set to a
// new ObjectPool with the ObjectPoolWithNoSyncPool option set.
func SetObjectPoolNoSyncPool() {
	globalObjectPoolNoSyncPool = true
	resetGlobalObjectPools()
}

// SetObjectPoolWithSyncPool results in the global ObjectPools being set to a
// new ObjectPool with the ObjectPoolWithNoSyncPool option unset.
//
// This is set by default.
func SetObjectPoolWithSyncPool() {
	globalObjectPoolNoSyncPool = false
	resetGlobalObjectPools()
}

// GetBytes gets a Bytes of the given length.
//
// This Bytes can be written to by accessing the backing Value.
// If pooling is not enabled via EnablePooling(), this will return an unmanaged Bytes.
//
// Note that there may be existing data inside the Value. If you need the Value
// to be zeroed out, use the MemsetZero function.
//
// This should not be called by users directly.
func GetBytes(valueLen int) *Bytes {
	if !globalEnabled {
		return newBytes(nil, valueLen)
	}
	return globalBytesPool.Get(valueLen)
}

// RegisterGlobalObjectPool registers the ObjectPool getter and setter.
//
// This should not be called by users directly.
//
// This should only be called at initialization.
func RegisterGlobalObjectPool(getter func() *ObjectPool, setter func(...ObjectPoolOption)) {
	globalObjectPoolGetters = append(globalObjectPoolGetters, getter)
	globalObjectPoolSetters = append(globalObjectPoolSetters, setter)
}

// PoolingEnabled returns true if pooling is enabled.
//
// This should not be called by users directly.
func PoolingEnabled() bool {
	return globalEnabled
}

func resetGlobalBytesPool() {
	if globalBytesPoolNoSyncPool {
		globalBytesPool = NewBytesPool(
			BytesPoolWithChannelSize(globalBytesPoolChannelSize),
			BytesPoolWithNoSyncPool(),
			BytesPoolWithSegListSizes(globalBytesPoolSegListSizes...),
		)
	} else {
		globalBytesPool = NewBytesPool(
			BytesPoolWithChannelSize(globalBytesPoolChannelSize),
			BytesPoolWithSegListSizes(globalBytesPoolSegListSizes...),
		)
	}
}

func resetGlobalObjectPools() {
	for _, setter := range globalObjectPoolSetters {
		if globalObjectPoolNoSyncPool {
			setter(
				ObjectPoolWithChannelSize(globalObjectPoolChannelSize),
				ObjectPoolWithNoSyncPool(),
			)
		} else {
			setter(
				ObjectPoolWithChannelSize(globalObjectPoolChannelSize),
			)
		}
	}
}

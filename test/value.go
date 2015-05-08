// Extensions for Protocol Buffers to create more go like structures.
//
// Copyright (c) 2013, Vastech SA (PTY) LTD. All rights reserved.
// http://github.com/gogo/protobuf/gogoproto
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
	"errors"

	"github.com/gogo/protobuf/proto"
)

type int31 interface {
	Int31() int32
}

type Value32 int32

func (v Value32) Marshal() ([]byte, error) {
	return proto.EncodeVarint(uint64(v)), nil
}

func (v *Value32) Unmarshal(bytes []byte) error {
	x, length := proto.DecodeVarint(bytes)
	if length != len(bytes) {
		return errors.New("invalid varint")
	}
	*v = Value32(x)
	return nil
}

func (v Value32) Equal(other Value32) bool {
	return v == other
}

func NewPopulatedValue32(r int31) *Value32 {
	v := Value32(r.Int31())
	return &v
}

type Value64 int64

func (v Value64) Marshal() ([]byte, error) {
	return proto.EncodeVarint(uint64(v)), nil
}

func (v *Value64) Unmarshal(bytes []byte) error {
	x, length := proto.DecodeVarint(bytes)
	if length != len(bytes) {
		return errors.New("invalid varint")
	}
	*v = Value64(x)
	return nil
}

func (v Value64) Equal(other Value64) bool {
	return v == other
}

func NewPopulatedValue64(r int63) *Value64 {
	v := Value64(r.Int63())
	return &v
}

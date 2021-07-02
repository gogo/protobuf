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

package types

import (
	fmt "fmt"
	"math/big"
)

type Caster struct{}

type Castee uint32

func (c *Caster) Equal(a, b *Castee) bool {
	if a == nil {
		return b == nil
	}
	return *a == *b
}

func (c *Caster) Size(a *Castee) int {
	return 1
}

func (c *Caster) MarshalTo(a *Castee, buf []byte) (int, error) {
	buf[0] = byte(*a)
	return 1, nil
}

func (c *Caster) Unmarshal(buf []byte) (*Castee, error) {
	ret := Castee(buf[0])
	return &ret, nil
}

func (c *Caster) NewPopulated() *Castee {
	t := Castee(0)
	return &t
}

type BigIntCaster struct{}

func (c *BigIntCaster) Equal(a, b *big.Int) bool {
	if a == nil {
		return b == nil
	}
	return a.Cmp(b) == 0
}

func (c *BigIntCaster) Size(a *big.Int) int {
	return len(a.Bytes()) + 1
}

func (c *BigIntCaster) MarshalTo(a *big.Int, buf []byte) (int, error) {
	bytes := a.Bytes()
	copy(buf[1:], bytes)
	if a.Sign() < 0 {
		buf[0] = 1
	} else {
		buf[0] = 0
	}
	return len(bytes) + 1, nil
}

func (c *BigIntCaster) Unmarshal(buf []byte) (*big.Int, error) {
	if len(buf) <= 1 {
		return big.NewInt(0), nil
	}
	ret := new(big.Int).SetBytes(buf[1:])
	switch buf[0] {
	case 0:

	case 1:
		ret = ret.Neg(ret)
	default:
		return nil, fmt.Errorf("invalid sign byte %x", buf[0])
	}

	return ret, nil
}

func (c *BigIntCaster) NewPopulated() *big.Int {
	return big.NewInt(0)
}

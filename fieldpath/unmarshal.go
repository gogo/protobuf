// Extensions for Protocol Buffers to create more go like structures.
//
// Copyright (c) 2013, Vastech SA (PTY) LTD. All rights reserved.
// http://code.google.com/p/gogoprotobuf/gogoproto
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

package fieldpath

type unmarshaler interface {
	unmarshal(buf []byte, offset int) (int, error)
}

//Used to Unmarshal a selected part of an input buffer.
type Unmarshaler struct {
	unmarshaler
	path     []uint64
	position int
	final    int
}

//Unmarshals the selected part of this input buffer.
func (this *Unmarshaler) Unmarshal(buf []byte) error {
	this.Reset()
	return Compile(this).Unmarshal(buf)
}

//Resets the position of the Unmarshaler in the compiled path.
func (this *Unmarshaler) Reset() {
	this.position = 0
	this.final = len(this.path) - 1
}

//One or more Unmarshalers compiled for single traversal of an input buffer.
type Compiled struct {
	children     map[uint64]*Compiled
	unmarshalers []unmarshaler
}

func newCompiled() *Compiled {
	return &Compiled{}
}

//Compiles a list of Unmarshalers into a object which can traverse an input buffer once and unmarshal all the selected parts.
func Compile(unmarshalers ...*Unmarshaler) *Compiled {
	c := newCompiled()
	for _, u := range unmarshalers {
		c.add(u.path, u.unmarshaler)
	}
	return c
}

func (this *Compiled) add(path []uint64, u unmarshaler) {
	if len(path) == 0 {
		this.unmarshalers = append(this.unmarshalers, u)
		return
	}
	if this.children == nil {
		this.children = make(map[uint64]*Compiled)
	}
	if c := this.children[path[0]]; c == nil {
		this.children[path[0]] = newCompiled()
	}
	this.children[path[0]].add(path[1:], u)
}

//Unmarshals all the slected parts with a single traversal of the input buffer.
func (this *Compiled) Unmarshal(buf []byte) error {
	i := 0
	for i < len(buf) {
		key, n, err := decodeVarint(buf, i)
		if err != nil {
			return err
		}
		i = i + n
		wireType := int(key & 0x7)
		c := this.children[key]
		if c == nil {
			i, err = skip(buf, i, wireType)
			if err != nil {
				return err
			}
			continue
		}
		if n, err := c.unmarshal(buf, i); err != nil {
			return err
		} else {
			i += n
		}
	}
	return nil
}

func (this *Compiled) unmarshal(buf []byte, offset int) (int, error) {
	var n int
	var err error
	i := offset
	for _, u := range this.unmarshalers {
		n, err = u.unmarshal(buf, i)
		if err != nil {
			return 0, err
		}
	}
	if this.children == nil {
		return n, nil
	}
	length, n, err := decodeVarint(buf, i)
	if err != nil {
		return 0, err
	}
	i += n
	nn := n + int(length)
	endOf := i + int(length)
	for i < endOf {
		key, n, err := decodeVarint(buf, i)
		if err != nil {
			return i - offset, err
		}
		i = i + n
		wireType := int(key & 0x7)
		c, ok := this.children[key]
		if !ok {
			i, err = skip(buf, i, wireType)
			if err != nil {
				return i - offset, err
			}
			continue
		}
		if n, err := c.unmarshal(buf, i); err != nil {
			return i - offset, err
		} else {
			i += n
		}
	}
	return nn, nil
}

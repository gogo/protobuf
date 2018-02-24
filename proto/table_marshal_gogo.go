// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2018, The GoGo Authors. All rights reserved.
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

package proto

// makeMessageRefMarshaler differs a bit from makeMessageMarshaler
// It marshal a message T instead of a *T
func makeMessageRefMarshaler(u *marshalInfo) (sizer, marshaler) {
	return func(ptr pointer, tagsize int) int {
			siz := u.size(ptr)
			return siz + SizeVarint(uint64(siz)) + tagsize
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
			b = appendVarint(b, wiretag)
			siz := u.cachedsize(ptr)
			b = appendVarint(b, uint64(siz))
			return u.marshal(b, ptr, deterministic)
		}
}

// makeMessageRefSliceMarshaler differs quite a lot from makeMessageSliceMarshaler
// It marshals a slice of messages []T instead of []*T
func makeMessageRefSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
	return func(ptr pointer, tagsize int) int {
			s := ptr.getSlice(u.typ)
			n := 0
			for i := 0; i < s.Len(); i++ {
				elem := s.Index(i)
				e := elem.Interface()
				v := toAddrPointer(&e, false)
				siz := u.size(v)
				n += siz + SizeVarint(uint64(siz)) + tagsize
			}
			return n
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
			s := ptr.getSlice(u.typ)
			var err, errreq error
			for i := 0; i < s.Len(); i++ {
				elem := s.Index(i)
				e := elem.Interface()
				v := toAddrPointer(&e, false)
				b = appendVarint(b, wiretag)
				siz := u.cachedsize(v)
				b = appendVarint(b, uint64(siz))
				b, err = u.marshal(b, v, deterministic)

				if err != nil {
					if _, ok := err.(*RequiredNotSetError); ok {
						// Required field in submessage is not set.
						// We record the error but keep going, to give a complete marshaling.
						if errreq == nil {
							errreq = err
						}
						continue
					}
					if err == ErrNil {
						err = errRepeatedHasNil
					}
					return b, err
				}
			}

			return b, errreq
		}
}

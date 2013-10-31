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

package proto

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
)

func (p *Properties) setCustomEncAndDec(typ reflect.Type) {
	p.ctype = typ
	if p.Repeated {
		p.enc = (*Buffer).enc_custom_slice_bytes
		p.dec = (*Buffer).dec_custom_slice_bytes
		p.size = size_custom_slice_bytes
	} else if typ.Kind() == reflect.Ptr {
		p.enc = (*Buffer).enc_custom_bytes
		p.dec = (*Buffer).dec_custom_bytes
		p.size = size_custom_bytes
	} else {
		p.enc = (*Buffer).enc_custom_ref_bytes
		p.dec = (*Buffer).dec_custom_ref_bytes
		p.size = size_custom_ref_bytes
	}
}

func (p *Properties) setNonNullableEncAndDec(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Bool:
		p.enc = (*Buffer).enc_ref_bool
		p.dec = (*Buffer).dec_ref_bool
		p.size = size_ref_bool
		if len(p.Default) > 0 {
			if p.Default == "1" {
				p.setDef = func(base structPointer, p *Properties) {
					setDefaultBool(base, p, true)
				}
			}
		}
	case reflect.Int32, reflect.Uint32:
		p.enc = (*Buffer).enc_ref_int32
		p.dec = (*Buffer).dec_ref_int32
		p.size = size_ref_int32
		if len(p.Default) > 0 {
			var v uint32
			if typ.Kind() == reflect.Int32 {
				v2, _ := strconv.ParseInt(p.Default, 0, 32)
				v = uint32(v2)
			} else {
				v2, _ := strconv.ParseUint(p.Default, 0, 32)
				v = uint32(v2)
			}
			if v != 0 {
				p.setDef = func(base structPointer, p *Properties) {
					setDefaultInt32(base, p, v)
				}
			}
		}
	case reflect.Int64, reflect.Uint64:
		p.enc = (*Buffer).enc_ref_int64
		p.dec = (*Buffer).dec_ref_int64
		p.size = size_ref_int64
		if len(p.Default) > 0 {
			var v uint64
			if typ.Kind() == reflect.Int64 {
				v2, _ := strconv.ParseInt(p.Default, 0, 64)
				v = uint64(v2)
			} else {
				v, _ = strconv.ParseUint(p.Default, 0, 64)
			}
			p.setDef = func(base structPointer, p *Properties) {
				setDefaultInt64(base, p, v)
			}
		}
	case reflect.Float32:
		p.enc = (*Buffer).enc_ref_int32 // can just treat them as bits
		p.dec = (*Buffer).dec_ref_int32
		p.size = size_ref_int32
		if len(p.Default) > 0 {
			f, _ := strconv.ParseFloat(p.Default, 32)
			p.setDef = func(base structPointer, p *Properties) {
				setDefaultInt32(base, p, math.Float32bits(float32(f)))
			}
		}
	case reflect.Float64:
		p.enc = (*Buffer).enc_ref_int64 // can just treat them as bits
		p.dec = (*Buffer).dec_ref_int64
		p.size = size_ref_int64
		if len(p.Default) > 0 {
			f, _ := strconv.ParseFloat(p.Default, 64)
			p.setDef = func(base structPointer, p *Properties) {
				setDefaultInt64(base, p, math.Float64bits(f))
			}
		}
	case reflect.String:
		p.dec = (*Buffer).dec_ref_string
		p.enc = (*Buffer).enc_ref_string
		p.size = size_ref_string
		if len(p.Default) > 0 {
			p.setDef = func(base structPointer, p *Properties) {
				setDefaultString(base, p, p.Default)
			}
		}
	case reflect.Struct:
		p.stype = typ
		p.isMarshaler = isMarshaler(typ)
		p.isUnmarshaler = isUnmarshaler(typ)
		if p.Wire == "bytes" {
			p.enc = (*Buffer).enc_ref_struct_message
			p.dec = (*Buffer).dec_ref_struct_message
			p.size = size_ref_struct_message
		} else {
			fmt.Fprintf(os.Stderr, "proto: no coders for struct %T\n", typ)
		}
	default:
		return false
	}
	return true
}

func (p *Properties) setSliceOfNonPointerStructs(typ reflect.Type) {
	t2 := typ.Elem()
	p.sstype = typ
	p.stype = t2
	p.isMarshaler = isMarshaler(t2)
	p.isUnmarshaler = isUnmarshaler(t2)
	p.enc = (*Buffer).enc_slice_ref_struct_message
	p.dec = (*Buffer).dec_slice_ref_struct_message
	p.size = size_slice_ref_struct_message
	if p.Wire != "bytes" {
		fmt.Fprintf(os.Stderr, "proto: no ptr oenc for %T -> %T \n", typ, t2)
	}
}

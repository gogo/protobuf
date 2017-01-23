// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2016, The GoGo Authors. All rights reserved.
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
	"fmt"
	"io"
	"time"
)

func NewPopulatedTimestamp(r interface {
	Int63() int64
}, easy bool) *Timestamp {
	this := &Timestamp{}
	ns := int64(r.Int63())
	this.Seconds = ns / 1e9
	this.Nanos = int32(ns % 1e9)
	return this
}

func (ts *Timestamp) String() string {
	return TimestampString(ts)
}

func NewPopulatedStdTime(r interface {
	Int63() int64
}, easy bool) *time.Time {
	timestamp := NewPopulatedTimestamp(r, easy)
	t, err := TimestampFromProto(timestamp)
	if err != nil {
		return nil
	}
	return &t
}

func SizeOfStdTime(t time.Time) int {
	ts, err := TimestampProto(t)
	if err != nil {
		return 0
	}
	return ts.Size()
}

func StdTimeMarshal(t time.Time) ([]byte, error) {
	size := SizeOfStdTime(t)
	buf := make([]byte, size)
	_, err := StdTimeMarshalTo(t, buf)
	return buf, err
}

func StdTimeMarshalTo(t time.Time, data []byte) (int, error) {
	ts, err := TimestampProto(t)
	if err != nil {
		return 0, err
	}
	return ts.MarshalTo(data)
}

func StdTimeUnmarshal(t *time.Time, data []byte) error {
	tt, err := unmarshalTimestampToStdTime(data)
	if err != nil {
		return err
	}
	*t = tt
	return nil
}

func unmarshalTimestampToStdTime(dAtA []byte) (time.Time, error) {
	var (
		seconds int64
		nanos   int32
	)

	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return time.Time{}, ErrIntOverflowTimestamp
			}
			if iNdEx >= l {
				return time.Time{}, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return time.Time{}, fmt.Errorf("proto: Timestamp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return time.Time{}, fmt.Errorf("proto: Timestamp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return time.Time{}, fmt.Errorf("proto: wrong wireType = %d for field Seconds", wireType)
			}
			seconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return time.Time{}, ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return time.Time{}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				seconds |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return time.Time{}, fmt.Errorf("proto: wrong wireType = %d for field Nanos", wireType)
			}
			nanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return time.Time{}, ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return time.Time{}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				nanos |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTimestamp(dAtA[iNdEx:])
			if err != nil {
				return time.Time{}, err
			}
			if skippy < 0 {
				return time.Time{}, ErrInvalidLengthTimestamp
			}
			if (iNdEx + skippy) > l {
				return time.Time{}, io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return time.Time{}, io.ErrUnexpectedEOF
	}
	return time.Unix(seconds, int64(nanos)).UTC(), nil
}

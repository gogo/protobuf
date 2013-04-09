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

/*
	Package gogoproto provides extensions for protocol buffers which help
	to create more canonical Go structures.

	All the extensions are FieldOptions:

	  - embed
	  - nullable (WARNING this breaks protocol buffer specifications, more below)
	  - customtype

	If you do not use any of these extension the code that is generated
	will be mostly the same as if goprotobuf has generated it.

	The most complete way to see examples is to look at

		code.google.com/p/gogoprotobuf/test/thetest.proto

	and its output generated in

		code.google.com/p/gogoprotobuf/test/gogo by gogoprotobuf

	and to compare

		code.google.com/p/gogoprotobuf/test/go by goprotobuf

	Let us look at:

		code.google.com/p/gogoprotobuf/test/example.proto

	for a quicker overview.

	The following message:

		message A {
			optional string Description = 1 [(gogoproto.nullable) = false];
			optional int64 Number = 2 [(gogoproto.nullable) = false];
			optional bytes Id = 3 [(gogoproto.customtype) = "code.google.com/p/gogoprotobuf/test/custom.Uuid", (gogoproto.nullable) = false];
		}

	Will generate a go struct which looks a lot like this:

		type A struct {
			Description string
			Number      int64
			Id          code_google_com_p_gogoprotobuf_test_custom.Uuid
		}

	You will see there are no pointers, since all fields are non-nullable.
	You will also see a custom type which marshals to a string.

	Next we will embed the message A in message B.

		message B {
			optional A A = 1 [(gogoproto.nullable) = false, (gogoproto.embed) = true];
			repeated bytes G = 2 [(gogoproto.customtype) = "code.google.com/p/gogoprotobuf/test/custom.Uint128", (gogoproto.nullable) = false];
		}

	See below that A is embedded in B.

		type B struct {
			A
			G []code_google_com_p_gogoprotobuf_test_custom.Uint128
		}

	Also see the repeated custom type.

		type Uint128 [2]uint64

	Gogoprotobuf also has some more subtle changes, these could be changed back:

	  - the generated package name for imports do not have the extra /filename.pb,
	  but are actually the imports specified in the .proto file.
	  - getters are not generated.
	  - enum value names are generated without the enumtype underscore as a prefix.

	Gogoprotobuf also has lost some features which should be brought back with time:

	  - Marshalling and unmarshalling with reflect and without the unsafe package,
	  this requires work in pointer_reflect.go

	Possible future plans:

	  - Generate all marshalling and unmarshalling code, this could be done with a plugin.

	Why does nullable break protocol buffer specifications:

	The protocol buffer specification states, somewhere, that you should be able to tell whether a
	field is set or unset.  With the option nullable=false this feature is lost,
	since your non-nullable fields will always be set.  It can be seen as a layer on top of
	protocol buffers, where before and after marshalling all non-nullable fields are set
	and they cannot be unset.

*/
package gogoproto

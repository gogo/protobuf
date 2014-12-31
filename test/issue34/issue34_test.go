package issue34

import (
	"bytes"
	"github.com/gogo/protobuf/proto"
	"testing"
)

func TestZeroLengthOptionalBytes(t *testing.T) {
	roundtrip := func(f *Foo) *Foo {
		data, err := proto.Marshal(f)
		if err != nil {
			panic(err)
		}
		newF := &Foo{}
		err = proto.Unmarshal(data, newF)
		if err != nil {
			panic(err)
		}
		return newF
	}

	f := &Foo{}
	roundtrippedF := roundtrip(f)
	if roundtrippedF.Bar != nil {
		t.Fatalf("should be nil")
	}

	f.Bar = []byte{}
	roundtrippedF = roundtrip(f)
	if roundtrippedF.Bar == nil {
		t.Fatalf("should not be nil")
	}
	if len(roundtrippedF.Bar) != 0 {
		t.Fatalf("should be empty")
	}
}

func TestRepeatedOptional(t *testing.T) {
	repeated := &FooWithRepeated{Bar: [][]byte{[]byte("a"), []byte("b")}}
	data, err := proto.Marshal(repeated)
	if err != nil {
		panic(err)
	}
	optional := &Foo{}
	err = proto.Unmarshal(data, optional)
	if err != nil {
		panic(err)
	}

	if !bytes.Equal(optional.Bar, []byte("b")) {
		t.Fatalf("should return the last entry")
	}
}

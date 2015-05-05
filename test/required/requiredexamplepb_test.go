package required

import (
	"github.com/gogo/protobuf/proto"
	"testing"
)

func TestMarshalToErrorsWhenRequiredFieldIsNotPresent(t *testing.T) {
	data := RequiredExample{}
	buf, err := proto.Marshal(&data)
	if err == nil {
		t.Fatalf("err == nil; was %v instead", err)
	}
	if err.Error() != `proto: required field "theRequiredString" not set` {
		t.Fatalf(`err.Error() != "proto: required field "theRequiredString" not set"; was "%s" instead`, err.Error())
	}
	if len(buf) != 0 {
		t.Fatalf(`len(buf) != 0; was %d instead`, len(buf))
	}
}

func TestMarshalToSucceedsWhenRequiredFieldIsPresent(t *testing.T) {
	data := RequiredExample{
		TheRequiredString: proto.String("present"),
	}
	buf, err := proto.Marshal(&data)
	if err != nil {
		t.Fatalf("err != nil; was %v instead", err)
	}
	if len(buf) == 0 {
		t.Fatalf(`len(buf) == 0; expected nonzero`)
	}
}

func TestUnmarshalErrorsWhenRequiredFieldIsNotPresent(t *testing.T) {
	missingRequiredField := []byte{0x12, 0x8, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c}
	data := RequiredExample{}
	err := proto.Unmarshal(missingRequiredField, &data)
	if err == nil {
		t.Fatalf("err == nil; was %v instead", err)
	}
	if err.Error() != `proto: required field "theRequiredString" not set` {
		t.Fatalf(`err.Error() != "proto: required field "theRequiredString" not set"; was "%s" instead`, err.Error())
	}
}

func TestUnmarshalSucceedsWhenRequiredIsNotPresent(t *testing.T) {
	dataOut := RequiredExample{
		TheRequiredString: proto.String("present"),
	}
	encodedMessage, _ := dataOut.Marshal()
	dataIn := RequiredExample{}
	err := proto.Unmarshal(encodedMessage, &dataIn)
	if err != nil {
		t.Fatalf("err != nil; was %v instead", err)
	}
}

package protoregisterprefix

import (
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestProtoRegisterPrefix(t *testing.T) {
	wantedMessageName := "FooBar.Foo"
	if proto.MessageName(&Foo{}) != wantedMessageName {
		t.Fatalf("proto.MessageName(&Foo{}) != %s", wantedMessageName)
	}

	wantedEnumName := "FooBar.Bar"
	valueMap := proto.EnumValueMap(wantedEnumName)
	if len(valueMap) != len(Bar_value) || valueMap["FOO"] != valueMap["FOO"] {
		t.Fatalf("proto.EnumValueMap(%s) != Bar_value", wantedEnumName)
	}
}

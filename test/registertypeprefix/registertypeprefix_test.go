package registertypeprefix

import (
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestRegisterTypePrefix(t *testing.T) {
	wantedName := "FooBar.Foo"
	if proto.MessageName(&Foo{}) != wantedName {
		t.Fatalf("proto.MessageName(&Foo{}) != %s", wantedName)
	}
}

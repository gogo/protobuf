package typeprefix

import (
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestTypePrefix(t *testing.T) {
	wantedName := "FooBar.Foo"
	if proto.MessageName(&Foo{}) != wantedName {
		t.Fatalf("proto.MessageName(&Foo{}) != %s", wantedName)
	}
}

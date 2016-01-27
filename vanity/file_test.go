package vanity

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

func TestNotGoogleProtobufDescriptorProto(t *testing.T) {
	testNotGoogleProtobufDescriptorProto(t, "google.protobuf", "descriptor.proto", false)
	testNotGoogleProtobufDescriptorProto(t, "google.protobuf", "google/protobuf/descriptor.proto", false)
	testNotGoogleProtobufDescriptorProto(t, "google.protobuf", "descriptor.protoo", true)
	testNotGoogleProtobufDescriptorProto(t, "google.protobuf", "any.proto", true)
	testNotGoogleProtobufDescriptorProto(t, "google.protobuff", "descriptor.proto", true)
}

func testNotGoogleProtobufDescriptorProto(t *testing.T, pkg string, name string, expected bool) {
	if NotGoogleProtobufDescriptorProto(
		&descriptor.FileDescriptorProto{
			Name:    proto.String(name),
			Package: proto.String(pkg),
		},
	) != expected {
		t.Errorf(
			"NotInGoogleProtobufDescriptorProto expected to return %v, returned %v, package was %s, name was %s",
			expected,
			!expected,
			pkg,
			name,
		)
	}
}

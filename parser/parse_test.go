package parser

import (
	"testing"
)

func TestParse(t *testing.T) {
	_, err := ParseFile("../protoc-gen-gogo/descriptor/descriptor.proto", "../protoc-gen-gogo/descriptor/")
	if err != nil {
		panic(err)
	}
}

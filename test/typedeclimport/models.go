package typedeclimport

import subpkg "github.com/frankee/protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}

package typedeclimport

import subpkg "github.com/buptbill220/protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}

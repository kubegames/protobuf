package typedeclimport

import subpkg "github.com/kubegames/protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}

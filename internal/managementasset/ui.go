package managementasset

import _ "embed"

//go:embed dist/index.html
var indexHTML []byte

func GetIndexHTML() []byte {
    return indexHTML
}

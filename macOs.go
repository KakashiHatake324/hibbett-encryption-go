//go:build !windows

package hibbettencryptiongo

import _ "embed"

//go:embed exec/encryptions
var program []byte

//go:build windows

package hibbettencryptiongo

import _ "embed"

//go:embed exec/encryptions.exe
var program []byte

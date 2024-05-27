package hash

import (
	"hash/crc64"
)

func CRC(content string) uint64 {
	// ref: https: //golang.hotexamples.com/examples/hash.crc64/-/New/golang-new-function-examples.html

	h := crc64.New(crc64.MakeTable(crc64.ECMA))
	h.Write([]byte(content))
	return h.Sum64()
}

package hash

import (
	jump "github.com/dgryski/go-jump"
)

// 一致性hash
// jump hash
// ref: https://blog.csdn.net/ltochange/article/details/121294393
func Jump(key uint64, numBuckets int) int32 {
	return jump.Hash(key, numBuckets)
}

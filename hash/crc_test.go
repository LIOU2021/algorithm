package hash

import (
	"hash/crc64"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	test_content = "hello world"
)

func Test_CRC(t *testing.T) {

	newSum1 := CRC(test_content)
	newSum2 := crc64.Checksum([]byte(test_content), crc64.MakeTable(crc64.ECMA))
	assert.Equal(t, newSum1, newSum2)
	t.Logf("CRC: %d", newSum1)
}

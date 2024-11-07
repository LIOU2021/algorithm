package hash

import (
	"hash/crc64"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	test_content = "hello world"
)

func Test_CRC_1(t *testing.T) {
	newSum1 := CRC(test_content)
	newSum2 := crc64.Checksum([]byte(test_content), crc64.MakeTable(crc64.ECMA))
	assert.Equal(t, newSum1, newSum2)
	t.Logf("CRC: %d", newSum1)
}

func Test_CRC_2(t *testing.T) {
	countSlice := [10]int{}
	for i := 0; i < 100; i++ {
		id := uuid.New().String()
		id2, err := uuid.Parse(id)

		assert.NoError(t, err)
		c := CRC(id) % 10
		countSlice[c]++
		t.Log(id, id == id2.String(), id2.Version(), id2.Variant(), CRC(id), c)
	}

	for key, val := range countSlice {
		t.Logf("key: %d, val: %d\n", key, val)
	}
}

// 测试相同input, output是否一致
func Test_CRC_3(t *testing.T) {
	newSum1 := CRC(test_content)
	for i := 0; i < 1000; i++ {
		newSum2 := CRC(test_content)
		assert.Equal(t, newSum1, newSum2)
	}
}

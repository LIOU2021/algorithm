package hash

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// buckets, 跟n越大，标准差就会明显越小，离散程度越低
// 尽量让每个bucket平均负载
func Test_Jump(t *testing.T) {
	buckets1 := 20
	buckets2 := 21
	diff := 0
	n := 1000000
	m1 := make(map[int32]int)
	m2 := make(map[int32]int)

	t1 := make(map[int]int32)

	for i := 0; i < n; i++ {
		b1 := Jump(uint64(i), buckets1)

		if v, exist := t1[i]; exist { // 测试是否buckets相同的情况下，hash 出来的bucket编号一致
			assert.Equal(t, v, b1)
		} else {
			t1[i] = b1
		}

		m1[b1]++

		b2 := Jump(uint64(i), buckets2)
		m2[b2]++

		// 计算离散程度
		if b1 != b2 { // 两个不同buckets 大小的jumpHash，hash出来结果的差异
			diff++
		}
	}
	log.Println("diff: ", diff)
	log.Println("diff%: ", diff*100/n)

	for i, v := range m1 {
		log.Printf("m1 bucket: %d, count: %d\n", i, v)
	}

	log.Println("========")

	for i, v := range m2 {
		log.Printf("m2 bucket: %d, count: %d\n", i, v)
	}

}

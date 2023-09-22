package sync

import (
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var TestDataLn = 100
var TestDataPool = make([]int, 0, TestDataLn)

func TestMain(t *testing.M) {
	for i := 0; i < TestDataLn; i++ {
		TestDataPool = append(TestDataPool, rand.Intn(TestDataLn))
	}
	code := t.Run()
	defer os.Exit(code)
}

func Test_PoolChain_FIFO(t *testing.T) {
	// t.Log(TestDataPool)
	p := NewPoolChain()

	for i := 0; i < TestDataLn; i++ {
		p.PushHead(TestDataPool[i])

	}

	for i := 0; i < TestDataLn; i++ {
		v, ok := p.PopTail()
		assert.True(t, ok)
		assert.Equal(t, TestDataPool[i], v)
	}
}

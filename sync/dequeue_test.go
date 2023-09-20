package sync

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Dequeue(t *testing.T) {
	q := NewPoolDequeue(4)

	assert.True(t, q.PushHead(1))
	assert.True(t, q.PushHead(2))
	assert.True(t, q.PushHead(3))
	assert.True(t, q.PushHead(4))
	assert.False(t, q.PushHead(5))

	t1, ok1 := q.PopHead()
	assert.Equal(t, 4, t1)
	assert.True(t, ok1)

	t2, ok2 := q.PopHead()
	assert.Equal(t, 3, t2)
	assert.True(t, ok2)

	t3, ok3 := q.PopTail()
	assert.Equal(t, 1, t3)
	assert.True(t, ok3)

	t4, ok4 := q.PopTail()
	assert.Equal(t, 2, t4)
	assert.True(t, ok4)

	t5, ok5 := q.PopTail()
	assert.Nil(t, t5)
	assert.False(t, ok5)

	assert.True(t, q.PushHead(10))
	assert.True(t, q.PushHead(13))
	assert.True(t, q.PushHead(19))

	t6, ok6 := q.PopHead()
	assert.Equal(t, 19, t6)
	assert.True(t, ok6)

	assert.True(t, q.PushHead(15))
	t7, ok7 := q.PopHead()
	assert.Equal(t, 15, t7)
	assert.True(t, ok7)
}

func Test_Dequeue_Official(t *testing.T) {
	testPoolDequeue(t, testNewPoolDequeue(16))
}

func testNewPoolDequeue(n int) PoolDequeueI {
	d := &PoolDequeue{
		vals: make([]eface, n),
	}
	// For testing purposes, set the head and tail indexes close
	// to wrapping around.
	d.headTail = d.pack(1<<dequeueBits-500, 1<<dequeueBits-500)
	return d
}

func testPoolDequeue(t *testing.T, d PoolDequeueI) {
	const P = 10
	var N int = 2e6
	if testing.Short() {
		N = 1e3
	}
	have := make([]int32, N)
	var stop int32
	var wg sync.WaitGroup
	record := func(val int) {
		atomic.AddInt32(&have[val], 1)
		if val == N-1 {
			atomic.StoreInt32(&stop, 1)
		}
	}

	// Start P-1 consumers.
	for i := 1; i < P; i++ {
		wg.Add(1)
		go func() {
			fail := 0
			for atomic.LoadInt32(&stop) == 0 {
				val, ok := d.PopTail()
				if ok {
					fail = 0
					record(val.(int))
				} else {
					// Speed up the test by
					// allowing the pusher to run.
					if fail++; fail%100 == 0 {
						runtime.Gosched()
					}
				}
			}
			wg.Done()
		}()
	}

	// Start 1 producer.
	nPopHead := 0
	wg.Add(1)
	go func() {
		for j := 0; j < N; j++ {
			for !d.PushHead(j) {
				// Allow a popper to run.
				runtime.Gosched()
			}
			if j%10 == 0 {
				val, ok := d.PopHead()
				if ok {
					nPopHead++
					record(val.(int))
				}
			}
		}
		wg.Done()
	}()
	wg.Wait()

	// Check results.
	for i, count := range have {
		if count != 1 {
			t.Errorf("expected have[%d] = 1, got %d", i, count)
		}
	}
	// Check that at least some PopHeads succeeded. We skip this
	// check in short mode because it's common enough that the
	// queue will stay nearly empty all the time and a PopTail
	// will happen during the window between every PushHead and
	// PopHead.
	if !testing.Short() && nPopHead == 0 {
		t.Errorf("popHead never succeeded")
	}
}

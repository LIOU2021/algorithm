package delayqueue

import (
	"container/heap"
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	pq := newPriorityQueue(5)
	if len(pq) != 0 {
		t.Errorf("Expected length of new queue to be 0, got %d", len(pq))
	}
	if cap(pq) != 5 {
		t.Errorf("Expected capacity of new queue to be 5, got %d", cap(pq))
	}
}

func TestPriorityQueueOperations(t *testing.T) {
	pq := newPriorityQueue(5)
	heap.Init(&pq)

	// Test Push and Len
	heap.Push(&pq, &item{Value: "test1", Priority: 30})
	heap.Push(&pq, &item{Value: "test2", Priority: 10})
	heap.Push(&pq, &item{Value: "test3", Priority: 40})
	heap.Push(&pq, &item{Value: "test4", Priority: 20})

	expectValue := []string{"test2", "test4", "test1", "test3"} // 根据Priority预期正确顺序
	t.Logf("cap: %d, len: %d\n", cap(pq), len(pq))

	pq.Range(func(i int, data *item) { // 会根据Priority由小到大打印
		t.Logf("index: %d, item: %+v\n", i, data)
	})

	if pq.Len() != 4 {
		t.Errorf("Expected length after pushes to be 4, got %d", pq.Len())
	}

	for i := 0; pq.Len() != 0; i++ {
		it := heap.Pop(&pq).(*item)
		if expectValue[i] != it.Value {
			t.Errorf("ordered - index: %d, expect: %s, actual: %s\n", i, expectValue[i], it.Value)
		}
	}
}
func TestPriorityQueueGrowAndShrink(t *testing.T) {
	pq := newPriorityQueue(20)
	heap.Init(&pq)

	// Test growing
	for i := 0; i < 30; i++ {
		heap.Push(&pq, &item{Value: i, Priority: int64(i)})
	}

	if cap(pq) <= 20 {
		t.Errorf("Queue did not grow as expected")
	}

	// Test shrinking
	for i := 0; i < 25; i++ {
		heap.Pop(&pq)
	}

	if cap(pq) > 20 {
		t.Errorf("Queue did not shrink as expected")
	}
}

func TestPeekAndShift(t *testing.T) {
	pq := newPriorityQueue(5)
	heap.Init(&pq)

	heap.Push(&pq, &item{Value: "test1", Priority: 10})
	heap.Push(&pq, &item{Value: "test2", Priority: 20})

	// Test when max is greater than the lowest priority
	item, delta := pq.PeekAndShift(15)
	if item == nil || item.Value != "test1" {
		t.Errorf("PeekAndShift did not return the expected item")
	}
	if delta != 0 {
		t.Errorf("Expected delta to be 0, got %d", delta)
	}

	// Test when max is less than the lowest priority
	item, delta = pq.PeekAndShift(5)
	if item != nil {
		t.Errorf("PeekAndShift should return nil when max is less than lowest priority")
	}
	if delta != 15 {
		t.Errorf("Expected delta to be 15, got %d", delta)
	}

	// Test on empty queue
	pq = newPriorityQueue(0)
	item, delta = pq.PeekAndShift(10)
	if item != nil || delta != 0 {
		t.Errorf("PeekAndShift on empty queue should return nil, 0")
	}
}

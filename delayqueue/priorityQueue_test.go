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
	heap.Push(&pq, &item{Value: "test1", Priority: 3})
	heap.Push(&pq, &item{Value: "test2", Priority: 1})
	heap.Push(&pq, &item{Value: "test3", Priority: 4})

	pq.Range(func(i int, data *item) { // 会根据Priority由小到大打印
		t.Logf("index: %d, item: %+v\n", i, data)
	})

	if pq.Len() != 3 {
		t.Errorf("Expected length after pushes to be 3, got %d", pq.Len())
	}

	// Test order (Less)
	if pq[0].Priority != 1 || pq[1].Priority != 3 || pq[2].Priority != 4 {
		t.Errorf("Queue is not correctly ordered")
	}

	// Test Pop
	item := heap.Pop(&pq).(*item)
	if item.Priority != 1 || item.Value != "test2" {
		t.Errorf("Expected to pop item with priority 1, got %d", item.Priority)
	}

	// Test Swap
	pq.Swap(0, 1)
	if pq[0].Priority != 4 || pq[1].Priority != 3 {
		t.Errorf("Swap did not work correctly")
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

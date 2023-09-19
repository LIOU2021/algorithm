package heap

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Heap_Min(t *testing.T) {
	h := &MinHeap{8, 5, 2} // 创建一个初始堆

	heap.Init(h) // 初始化堆

	// 插入元素到堆中
	heap.Push(h, 3)
	heap.Push(h, 1)

	// 从堆中取出并打印最小的元素
	assert.Equal(t, 1, heap.Pop(h))
	assert.Equal(t, 2, heap.Pop(h))
	assert.Equal(t, 3, heap.Pop(h))
	assert.Equal(t, 5, heap.Pop(h))
	assert.Equal(t, 8, heap.Pop(h))
}

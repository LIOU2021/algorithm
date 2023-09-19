package heap

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Heap_Max(t *testing.T) {
	h := &MaxHeap{2, 1, 5} // 创建一个初始堆

	heap.Init(h) // 初始化堆

	// 插入元素到堆中
	heap.Push(h, 3)
	heap.Push(h, 7)

	// 从堆中取出并打印最大的元素
	assert.Equal(t, 7, heap.Pop(h))
	assert.Equal(t, 5, heap.Pop(h))
	assert.Equal(t, 3, heap.Pop(h))
	assert.Equal(t, 2, heap.Pop(h))
	assert.Equal(t, 1, heap.Pop(h))
}

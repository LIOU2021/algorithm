package heap

// 最小堆

// 声明一个整数切片类型作为堆元素
type MinHeap []int

// 实现 heap.Interface 接口的方法：Len、Less、Swap
func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j] // 最小堆的比较条件是 h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 实现 heap.Interface 接口的方法：Push、Pop
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

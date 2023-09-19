package heap

// 最大堆

// 声明一个整数切片类型作为堆元素
type MaxHeap []int

// 实现 heap.Interface 接口的方法：Len、Less、Swap
func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j] // 最大堆的比较条件是 h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 实现 heap.Interface 接口的方法：Push、Pop
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

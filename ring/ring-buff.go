package ring

type CircularQueue struct {
	items []interface{} // 环形队列的存储空间
	n     int           // 环形队列的容量
	head  int           // 队首指针
	tail  int           // 队尾指针
}

func NewCircularQueue(n int) *CircularQueue {
	return &CircularQueue{
		items: make([]interface{}, n),
		n:     n,
		head:  0,
		tail:  0,
	}
}

func (q *CircularQueue) Enqueue(item interface{}) bool {
	if q.IsFull() {
		return false
	}
	q.items[q.tail] = item
	q.tail = (q.tail + 1) % q.n
	return true
}

func (q *CircularQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	item := q.items[q.head]
	q.items[q.head] = nil
	q.head = (q.head + 1) % q.n
	return item
}

func (q *CircularQueue) IsFull() bool {
	return (q.tail+1)%q.n == q.head
}

func (q *CircularQueue) IsEmpty() bool {
	return q.head == q.tail
}

func (q *CircularQueue) Size() int {
	return (q.tail - q.head + q.n) % q.n
}

// 遍历
func (q *CircularQueue) Each(fn func(node interface{})) {
	for i := q.head; i < q.head+q.Size(); i++ {
		fn(q.items[i%q.n])
	}
}

// 清空
func (q *CircularQueue) Clear() bool {
	q.n = 0
	q.head = 0
	q.tail = 0
	q.items = nil
	return true
}

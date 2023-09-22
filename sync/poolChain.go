package sync

import (
	"sync/atomic"
	"unsafe"
)

// poolChain is a dynamically-sized version of poolDequeue.
//
// This is implemented as a doubly-linked list queue of poolDequeues
// where each dequeue is double the size of the previous one. Once a
// dequeue fills up, this allocates a new one and only ever pushes to
// the latest dequeue. Pops happen from the other end of the list and
// once a dequeue is exhausted, it gets removed from the list.
type PoolChain struct {
	// head is the poolDequeue to push to. This is only accessed
	// by the producer, so doesn't need to be synchronized.
	head *PoolChainElt

	// tail is the poolDequeue to popTail from. This is accessed
	// by consumers, so reads and writes must be atomic.
	tail *PoolChainElt
}

// 把双向链结构，元素放dequeue，每个节点都是一个dequeues(PoolDequeue)
type PoolChainElt struct {
	PoolDequeue

	// next and prev link to the adjacent PoolChainElts in this
	// poolChain.
	//
	// next is written atomically by the producer and read
	// atomically by the consumer. It only transitions from nil to
	// non-nil.
	//
	// prev is written atomically by the consumer and read
	// atomically by the producer. It only transitions from
	// non-nil to nil.
	next, prev *PoolChainElt
}

func storePoolChainElt(pp **PoolChainElt, v *PoolChainElt) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(pp)), unsafe.Pointer(v))
}

func loadPoolChainElt(pp **PoolChainElt) *PoolChainElt {
	return (*PoolChainElt)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(pp))))
}

func (c *PoolChain) PushHead(val any) {
	// 选取head指向的dequeue
	d := c.head
	if d == nil {
		// Initialize the chain.
		// 初始化一个dequeue，容量是8
		const initSize = 8 // Must be a power of 2
		d = new(PoolChainElt)
		d.vals = make([]eface, initSize)
		// 将head指针指向创建的dequeue指标
		c.head = d
		storePoolChainElt(&c.tail, d)
	}

	// 塞资料到dequeue
	// c.head.PoolDequeue.PushHead(val)
	if d.PushHead(val) { // 成功塞入dequeue就会退出
		return
	}

	// 底下都是当head的dequeue容量满时才会执行到的策略

	// 执行到这里时代表d.PushHead(val)并未成功将资料塞入dequeue，因为dequeue满了
	// 所以要多new一个dequeue
	// The current dequeue is full. Allocate a new one of twice
	// the size.
	newSize := len(d.vals) * 2 // 新开的dequeue容量会是当前c.head的dequeue的容量的*2
	if newSize >= dequeueLimit {
		// Can't make it any bigger.
		newSize = dequeueLimit
	}

	d2 := &PoolChainElt{prev: d} // 因为head dequeue容量不足时而新开dequeue时，原本的head dequeue将改为 链表 的prev
	d2.vals = make([]eface, newSize)
	c.head = d2                    // 新创建的dequeue节点将成为 PoolChain的head
	storePoolChainElt(&d.next, d2) // 示意结果: node1 > node2(新创建的dequeue，目前PoolChain head的指向)
	d2.PushHead(val)               // 将值塞入新创建的dequeue
}

// 如果不断PopHead，当head指向的dequeue已经没元素时，head的参考指标也不会改变，也不会销毁当前的dequeue
// 也就是 dequeue的容量只会越开越大
func (c *PoolChain) PopHead() (any, bool) {
	d := c.head
	for d != nil { // 当*PoolChainElt已经找到处于PoolChain最左边的时，就代表没资料了
		// 执行c.head.PoolDequeue.PopHead()
		if val, ok := d.PopHead(); ok {
			return val, ok
		}
		// There may still be unconsumed elements in the
		// previous dequeue, so try backing up.
		d = loadPoolChainElt(&d.prev) // 如果当前head的dequeue没元素时，持续透过*PoolChainElt.prev找前一个dequeue
	}
	return nil, false
}

func (c *PoolChain) PopTail() (any, bool) {
	d := loadPoolChainElt(&c.tail)
	if d == nil {
		return nil, false
	}

	for {
		// It's important that we load the next pointer
		// *before* popping the tail. In general, d may be
		// transiently empty, but if next is non-nil before
		// the pop and the pop fails, then d is permanently
		// empty, which is the only condition under which it's
		// safe to drop d from the chain.
		d2 := loadPoolChainElt(&d.next)

		if val, ok := d.PopTail(); ok { // 有元素能弹出就结束function
			return val, ok
		}

		if d2 == nil { // 如果d.next是nil就当前的PoolChainElts是PoolChain上唯一的node了，也可以理解成*PoolChain.head指向的位置
			// This is the only dequeue. It's empty right
			// now, but could be pushed to in the future.
			return nil, false
		}

		// The tail of the chain has been drained, so move on
		// to the next dequeue. Try to drop it from the chain
		// so the next pop doesn't have to look at the empty
		// dequeue again.
		// 如果当前*PoolDequeue.PopTail取不到资料时，则表示*PoolChain.tail指向的dequeue已经没元素了
		// 此时会把*PoolChain.tail指向原本的*PoolChain.tail.next，就是tail指向的*PoolChainElt往右移
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&c.tail)), unsafe.Pointer(d), unsafe.Pointer(d2)) {
			// We won the race. Clear the prev pointer so
			// the garbage collector can collect the empty
			// dequeue and so popHead doesn't back up
			// further than necessary.
			storePoolChainElt(&d2.prev, nil) // 原本处于PoolChain左边数来第二个的node(*PoolChainElt)变成第一个，所以它的prev自然会是nil，毕竟是成为最左边了
			// 原本处于PoolChain最左边旧的node(*PoolChainElt)就会没有被参考，等GC回收
		}
		d = d2 // 用重新排列过后PoolChain最左边的*PoolChainElt继续PopTail
	}
}

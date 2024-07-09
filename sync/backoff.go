package sync

/*
当遇到失败情境时的重新尝试策略

如果每个重新尝试频率都设定固定的话，当遇到大量请求并发情境时，将会产生retry storms

"Retry storms"（重试风暴）是指在计算机系统或网络中，由于多个客户端或进程在同一时间段内进行重试操作，从而导致系统过载或拥塞的现象。
这种情况通常发生在高并发环境中，当多个客户端或进程遇到错误或失败时，它们会在相同或接近的时间间隔内进行重试。如果没有有效的退避策略，这些重试操作会同时进行，导致系统资源耗尽，进而引发更多的失败和重试，形成恶性循环。这个现象形象地被称为“重试风暴”。

当采用random sleep duration时，機器可能因為db query 開始慢了，或是因为大量访问等情境导致当前服务处理响应时间增加了，此时机器所需處理時間增加，但sleep duration range 不變，此時應該放寬sleep duration range，避免无效的高retry频率
*/

import (
	"math/rand"
	"time"
)

// Exponential Jitter Backoff
// 指数抖动退避
type EJB struct {
	t time.Duration // 所需等待时间
}

// 参数b为获取锁成功时，则代入true，反之获取锁失败时，则给false
func (e *EJB) Execute(b bool) bool {
	if !b {
		e.t += time.Duration((rand.Int31n(50) + 20) * int32(time.Millisecond)) // 沈睡时间的step为20~20+49毫秒间递增
		time.Sleep(e.t)
		return false
	}

	return true
}

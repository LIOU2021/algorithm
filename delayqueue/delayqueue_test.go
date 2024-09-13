package delayqueue

import (
	"fmt"
	"testing"
	"time"
)

func TestDelayQueueDemo(t *testing.T) {
	dq := New(10)
	exitC := make(chan struct{})

	go dq.Poll(exitC, func() int64 {
		return time.Now().UnixMilli()
	})

	dq.Offer("task1", time.Now().Add(time.Second*1).UnixMilli())
	dq.Offer("task3", time.Now().Add(time.Second*3).UnixMilli())
	dq.Offer("task2", time.Now().Add(time.Second*2).UnixMilli())

	for i := 0; i < 3; i++ {
		t.Log(<-dq.C) // 打印顺序根据到期时间 task1 > task2 > task3
	}

	close(exitC) // 关闭dp.Poll()
}

func TestDelayQueue(t *testing.T) {
	dq := New(10)
	exitC := make(chan struct{})

	go dq.Poll(exitC, func() int64 {
		return time.Now().UnixMilli()
	})

	// 添加三个任务，分别在1秒、2秒和3秒后过期
	dq.Offer("task1", time.Now().Add(time.Second*1).UnixMilli())
	dq.Offer("task3", time.Now().Add(time.Second*3).UnixMilli())
	dq.Offer("task2", time.Now().Add(time.Second*2).UnixMilli())

	// 验证任务按正确的顺序过期
	for i := 1; i <= 3; i++ {
		select {
		case item := <-dq.C:
			if item != "task"+fmt.Sprintf("%d", i) {
				t.Errorf("Expected task%d, got %v", i, item)
			}
		case <-time.After(4 * time.Second): // 最多只有添加了3秒过期的策略，所以不致于会等到4秒
			t.Errorf("Timed out waiting for task%d", i)
		}
	}

	// 测试在队列为空时的行为
	select {
	case item := <-dq.C:
		t.Errorf("Unexpected item received: %v", item)
	case <-time.After(1 * time.Second):
		// 这是预期的行为
	}

	// 清理
	close(exitC)
}

func TestDelayQueueConcurrency(t *testing.T) {
	dq := New(100)
	exitC := make(chan struct{})

	go dq.Poll(exitC, func() int64 {
		return time.Now().UnixMilli()
	})

	// 并发添加多个任务
	for i := 0; i < 100; i++ {
		go func(i int) {
			dq.Offer("task"+fmt.Sprintf("%d", i), time.Now().Add(time.Duration(i*100)*time.Millisecond).UnixMilli())
		}(i)
	}

	// 接收并验证所有任务
	received := make(map[string]bool)
	for i := 0; i < 100; i++ {
		select {
		case item := <-dq.C:
			if item.(string) != fmt.Sprintf("task%d", i) { // 确认讯息消费顺序是否根据时间延迟对列
				t.Errorf("expect: %s, actually: %s", fmt.Sprintf("task%d", i), item.(string))
			}
			received[item.(string)] = true
		case <-time.After(10 * time.Second): // 延迟最大的任务应为i=99，99*100*Millisecond＝9.9秒，所以自然不会延迟到第10秒还没发送完讯息
			t.Fatalf("Timed out waiting for all tasks")
		}
	}

	// 验证是否收到了所有任务
	for i := 0; i < 100; i++ {
		if !received["task"+fmt.Sprintf("%d", i)] {
			t.Errorf("Did not receive task%d", i)
		}
	}

	// 清理
	close(exitC)
}

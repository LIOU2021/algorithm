package sync

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Backoff(t *testing.T) {
	startTime := time.Now()
	count := 300
	testNum := 0
	wg := sync.WaitGroup{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(index int) {
			b := EJB{}
			for !b.Execute(Lock()) {
				fmt.Printf("index: %d, sleep duration: %s\n", index, b.t.String())
			}

			time.Sleep(100 * time.Millisecond)
			testNum++
			defer func() {
				if !Unlock() {
					log.Fatal("unlock unexpected error")
				}
				fmt.Printf("index done : %d \n", index)
				wg.Done()
			}()
		}(i)
	}

	wg.Wait()
	assert.Equal(t, count, testNum)
	fmt.Println("testNum: ", testNum)
	fmt.Println("process duration: ", time.Since(startTime))
}

// 模拟锁
var lock atomic.Bool

// 互斥锁
func Lock() bool {
	return lock.CompareAndSwap(false, true)
}

// 释放锁
func Unlock() bool {
	return lock.CompareAndSwap(true, false)
}

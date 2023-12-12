// https://github.com/line/garr
package adder

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	ga "go.linecorp.com/garr/adder"
)

func Test_LongAdder(t *testing.T) {
	// or ga.DefaultAdder() which uses jdk long-adder as default
	adder := ga.NewLongAdder(ga.JDKAdderType)

	wg := sync.WaitGroup{}

	n := 1000000
	wg.Add(n)

	now := time.Now()

	for i := 0; i < n; i++ {
		go func() {
			adder.Add(123)
			wg.Done()
		}()
	}

	wg.Wait()

	// get total added value
	fmt.Printf("Test_LongAdder - sum: %d, du: %s\n", adder.Sum(), time.Since(now))
}

func Test_NormalAdder(t *testing.T) {

	wg := sync.WaitGroup{}
	var sum int64

	n := 1000000
	wg.Add(n)

	now := time.Now()

	for i := 0; i < n; i++ {
		go func() {
			atomic.AddInt64(&sum, 123)
			wg.Done()
		}()
	}

	wg.Wait()

	// get total added value
	fmt.Printf("Test_NormalAdder - sum: %d, du: %s\n", sum, time.Since(now))
}

// https://github.com/line/garr
package adder

import (
	"fmt"
	"sync"
	"testing"

	ga "go.linecorp.com/garr/adder"
)

func Test_LongAdder(t *testing.T) {
	// or ga.DefaultAdder() which uses jdk long-adder as default
	adder := ga.NewLongAdder(ga.JDKAdderType)

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			adder.Add(123)
			wg.Done()
		}()
	}

	wg.Wait()

	// get total added value
	fmt.Println(adder.Sum())
}

// https://github.com/line/garr
package adder

import (
	"fmt"
	"time"

	ga "go.linecorp.com/garr/adder"
)

func Test_LongAdder() {
	// or ga.DefaultAdder() which uses jdk long-adder as default
	adder := ga.NewLongAdder(ga.JDKAdderType)

	for i := 0; i < 100; i++ {
		go func() {
			adder.Add(123)
		}()
	}

	time.Sleep(1 * time.Second)

	// get total added value
	fmt.Println(adder.Sum())
}

package list

import (
	"fmt"
	"testing"

	"github.com/huandu/skiplist"
)

// redis zset 底层就是用skip list
func Test_SkipList_01(t *testing.T) {
	// Create a skip list with int key.
	list := skiplist.New(skiplist.Int)

	// Add some values. Value can be anything.
	list.Set(12, "hello world")
	list.Set(34, 56)
	list.Set(78, 90.12)

	// Get element by index.
	elem := list.Get(34)                // Value is stored in elem.Value.
	fmt.Println(elem.Value)             // Output: 56
	next := elem.Next()                 // Get next element.
	prev := next.Prev()                 // Get previous element.
	fmt.Println(next.Value, prev.Value) // Output: 90.12    56

	// Or, directly get value just like a map
	val, ok := list.GetValue(34)
	fmt.Println(val, ok) // Output: 56  true

	// Find first elements with score greater or equal to key
	foundElem := list.Find(30)
	fmt.Println(foundElem.Key(), foundElem.Value) // Output: 34 56

	// Remove an element for key.
	list.Remove(34)
}

func Test_SkipList_02(t *testing.T) {
	// Create a skip list with int key.
	list := skiplist.New(skiplist.Int)

	// 随机增加值
	list.Set(3, "hello world")
	list.Set(2, 56)
	list.Set(4, 90.12)
	list.Set(1, "hi")

	n := list.Front()
	for {
		if n == nil {
			break
		}
		fmt.Printf("key: %v, value: %v\n", n.Key(), n.Value)
		n = n.Next()
	}
	// 有序产出
	// key: 1, value: hi
	// key: 2, value: 56
	// key: 3, value: hello world
	// key: 4, value: 90.12
}

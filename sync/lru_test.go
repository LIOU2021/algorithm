package sync

// https://github.com/hashicorp/golang-lru
import (
	"fmt"
	"testing"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/stretchr/testify/assert"
)

func Test_LRU_NORMAL(t *testing.T) {
	l, _ := lru.New[int, any](128)
	for i := 0; i < 256; i++ {
		l.Add(i, nil)
	}

	assert.Equal(t, 128, l.Len(), fmt.Sprintf("bad len: %v", l.Len()))
}

func Test_LRU_EXP(t *testing.T) {
	// make cache with 10ms TTL and 5 max keys
	cache := expirable.NewLRU[string, string](5, nil, time.Millisecond*10)

	key := "key1"
	v := "val1"

	// set value under key1.
	cache.Add(key, v)

	// get value under key1
	r, ok := cache.Get(key)

	assert.True(t, ok)
	assert.Equal(t, v, r)
	// check for OK value
	// if ok {
	// 	fmt.Printf("value before expiration is found: %v, value: %q\n", ok, r)
	// }

	// wait for cache to expire
	time.Sleep(time.Millisecond * 12)

	// get value under key1 after key expiration
	r, ok = cache.Get(key)
	// fmt.Printf("value after expiration is found: %v, value: %q\n", ok, r)

	assert.False(t, ok)
	assert.Empty(t, r)
	// set value under key2, would evict old entry because it is already expired.
	cache.Add("key2", "val2")
	assert.Equal(t, 1, cache.Len())
	// fmt.Printf("Cache len: %d\n", cache.Len())
	// Output:
	// value before expiration is found: true, value: "val1"
	// value after expiration is found: false, value: ""
	// Cache len: 1
}

package ring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Ring_Buff(t *testing.T) {
	type user struct {
		id int
	}

	n := 10
	q := NewCircularQueue(n)
	assert.Equal(t, true, q.IsEmpty())
	assert.Equal(t, false, q.IsFull())

	for i := 0; i < 20; i++ {
		result := q.Enqueue(&user{
			id: i,
		})

		if i < (n - 1) {
			assert.Equal(t, true, result)
		} else {
			assert.Equal(t, false, result)
		}

	}
	assert.Equal(t, true, q.IsFull())
	assert.Equal(t, false, q.IsEmpty())
	assert.Equal(t, n-1, q.Size())

	assert.Equal(t, 0, q.Dequeue().(*user).id)
	assert.Equal(t, 1, q.Dequeue().(*user).id)
	assert.Equal(t, 2, q.Dequeue().(*user).id)

	testEachCount := 3
	q.Each(func(node interface{}) {
		assert.Equal(t, testEachCount, node.(*user).id)
		testEachCount++
	})

	for i := 3; i < 20; i++ {
		result := q.Dequeue()

		if i < (n - 1) {
			assert.NotNil(t, result, i)
		} else {
			assert.Nil(t, result, i)
		}
	}
	assert.Equal(t, true, q.Clear())
}

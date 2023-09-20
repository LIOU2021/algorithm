package sync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Dequeue(t *testing.T) {
	q := NewPoolDequeue(4)

	assert.True(t, q.PushHead(1))
	assert.True(t, q.PushHead(2))
	assert.True(t, q.PushHead(3))
	assert.True(t, q.PushHead(4))
	assert.False(t, q.PushHead(5))

	t1, ok1 := q.PopHead()
	assert.Equal(t, 4, t1)
	assert.True(t, ok1)

	t2, ok2 := q.PopHead()
	assert.Equal(t, 3, t2)
	assert.True(t, ok2)

	t3, ok3 := q.PopTail()
	assert.Equal(t, 1, t3)
	assert.True(t, ok3)

	t4, ok4 := q.PopTail()
	assert.Equal(t, 2, t4)
	assert.True(t, ok4)

	t5, ok5 := q.PopTail()
	assert.Nil(t, t5)
	assert.False(t, ok5)

	assert.True(t, q.PushHead(10))
	assert.True(t, q.PushHead(13))
	assert.True(t, q.PushHead(19))

	t6, ok6 := q.PopHead()
	assert.Equal(t, 19, t6)
	assert.True(t, ok6)

	assert.True(t, q.PushHead(15))
	t7, ok7 := q.PopHead()
	assert.Equal(t, 15, t7)
	assert.True(t, ok7)
}

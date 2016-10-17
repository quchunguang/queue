package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPut(t *testing.T) {
	q := New()
	q.Put(`test`)
	q.Put(2)
	assert.Equal(t, 2, q.Len())
	s := q.Get().(string)
	assert.Equal(t, `test`, s)
	assert.Equal(t, 1, q.Len())
	i := q.Get().(int)
	assert.Equal(t, 2, i)
	assert.True(t, q.Empty())
}

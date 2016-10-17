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

func TestGet(t *testing.T) {
	q := New()
	q.Put(`test`)
	s := q.Get().(string)
	assert.Equal(t, `test`, s)

	assert.True(t, q.Empty())
	nothing := q.Get()
	assert.Nil(t, nothing)
}

func TestLen(t *testing.T) {
	q := New()
	assert.Equal(t, 0, q.Len())

	q.Put(`test`)
	assert.Equal(t, 1, q.Len())

	q.Get()
	assert.Equal(t, 0, q.Len())
}

func TestContain(t *testing.T) {
	q := New()
	assert.False(t, q.Contain(1))

	q.Put(`test`)
	q.Put(1)
	assert.True(t, q.Contain(`test`))
	assert.False(t, q.Contain(`test1`))
	assert.False(t, q.Contain(nil))

	assert.True(t, q.Contain(1))
	assert.False(t, q.Contain(0))
}

func TestMap(t *testing.T) {
	q := New()

	q.Put(2)
	q.Put(1)

	var value int
	mapFunc := func(a interface{}) bool {
		if v, ok := a.(int); ok {
			return v == value
		}
		return false
	}

	value = 1
	b := q.Map(mapFunc)
	assert.Equal(t, 1, b.(int))

	value = 2
	c := q.Map(mapFunc)
	assert.Equal(t, 2, c.(int))

	value = 3
	d := q.Map(mapFunc)
	assert.Nil(t, d)
}

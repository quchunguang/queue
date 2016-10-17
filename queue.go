package queue

import (
	"container/list"
)

// Queue is describe as a list with lock
type Queue struct {
	sem  chan int
	list *list.List
}

// CallbackFunc type describe any operation on each element in stack.
type CallbackFunc func(val interface{}) bool

// New create a new Queue.
func New() *Queue {
	sem := make(chan int, 1)
	list := list.New()
	return &Queue{sem, list}
}

// Put an element into queue.
func (q *Queue) Put(val interface{}) {
	q.sem <- 1
	q.list.PushFront(val)
	<-q.sem
}

// Get an element out of the queue.
func (q *Queue) Get() interface{} {
	q.sem <- 1
	e := q.list.Back()
	if e != nil {
		q.list.Remove(e)
	}
	<-q.sem

	if e != nil {
		return e.Value
	}
	return nil
}

// Len get the length of the queue.
func (q *Queue) Len() int {
	return q.list.Len()
}

// Empty tests if the queue is empty.
func (q *Queue) Empty() bool {
	return q.list.Len() == 0
}

// Map returns the first element in the queue causing mapFunc returns true.
func (q *Queue) Map(mapFunc CallbackFunc) interface{} {
	q.sem <- 1
	e := q.list.Front()
	for e != nil {
		if mapFunc(e.Value) {
			<-q.sem
			return e.Value
		}
		e = e.Next()
	}
	<-q.sem
	return nil
}

// Contain tests if this item in the queue.
func (q *Queue) Contain(val interface{}) bool {
	q.sem <- 1
	e := q.list.Front()
	for e != nil {
		if e.Value == val {
			<-q.sem
			return true
		}
		e = e.Next()
	}
	<-q.sem
	return false
}

package queue

import (
	"container/list"
	"reflect"
)

type Queue struct {
	sem  chan int
	list *list.List
}

var tFunc func(val interface{}) bool

// Create a new Queue and return.
func New() *Queue {
	sem := make(chan int, 1)
	list := list.New()
	return &Queue{sem, list}
}

// Get size of the queue
func (q *Queue) Size() int {
	return q.list.Len()
}

// Put an element into queue.
func (q *Queue) Put(val interface{}) *list.Element {
	q.sem <- 1
	e := q.list.PushFront(val)
	<-q.sem
	return e
}

// Get an element out of the queue.
func (q *Queue) Get() *list.Element {
	q.sem <- 1
	e := q.list.Back()
	q.list.Remove(e)
	<-q.sem
	return e
}

// Len get the length of the queue.
func (q *Queue) Len() int {
	return q.list.Len()
}

// Empty tests if the queue is empty.
func (q *Queue) Empty() bool {
	return q.list.Len() == 0
}

// Queue returns the element in the queue only if func queueFunc(element) returns true..
func (q *Queue) Query(queryFunc interface{}) *list.Element {
	q.sem <- 1
	e := q.list.Front()
	for e != nil {
		if reflect.TypeOf(queryFunc) == reflect.TypeOf(tFunc) {
			if queryFunc.(func(val interface{}) bool)(e.Value) {
				<-q.sem
				return e
			}
		} else {
			<-q.sem
			return nil
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
		} else {
			e = e.Next()
		}
	}
	<-q.sem
	return false
}

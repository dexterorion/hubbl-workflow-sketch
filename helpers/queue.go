package helpers

import (
	"errors"
	"fmt"
)

var (
	ErrPopingEmptyQueue = errors.New("cannot pop from empty queue")
)

type Queue interface {
	Push(element interface{})
	Pop() (interface{}, error)
	Size() int64
	String() string
	AsArray() []interface{}
}

type queueImpl struct {
	size     int64
	elements []interface{}
}

func StartQueue(size int64, elements []interface{}) Queue {
	return &queueImpl{
		size:     size,
		elements: elements,
	}
}

func StartEmptyQueue() Queue {
	return &queueImpl{
		size:     0,
		elements: []interface{}{},
	}
}

func (q *queueImpl) Push(element interface{}) {
	// npe safety
	if q.elements == nil {
		q.elements = make([]interface{}, 0)
	}

	q.elements = append(q.elements, element)
	q.size++
}

func (q *queueImpl) Pop() (element interface{}, err error) {
	if q.size == 0 {
		err = ErrPopingEmptyQueue
		return
	}

	element, q.elements = q.elements[0], q.elements[1:]

	q.size--

	return
}

func (q *queueImpl) Size() int64 {
	return q.size
}

func (q *queueImpl) String() string {
	return fmt.Sprint(q.elements)
}

func (q *queueImpl) AsArray() []interface{} {
	return q.elements
}

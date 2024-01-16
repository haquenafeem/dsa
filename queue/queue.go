package queue

import "errors"

type Queue[T any] struct {
	root   *node[T]
	length int
}

func (q *Queue[T]) Length() int {
	return q.length
}

func (q *Queue[T]) Enqueue(value int) {
	q.length++

	n := &node[T]{data: value}
	if q.root == nil {
		q.root = n
		return
	}

	cur := q.root
	for cur.hasNext() {
		cur = cur.next
	}

	cur.next = n
}

func (q *Queue[T]) Dequeue() (int, error) {
	if q.root == nil {
		return -1, errors.New("empty queue")
	}

	q.length--
	value := q.root.data

	if q.root.hasNext() {
		q.root = q.root.next
	} else {
		q.root = nil
	}

	return value, nil
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

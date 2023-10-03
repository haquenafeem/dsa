package stack

import "errors"

type Stack[T any] struct {
	root   *node[T]
	length int
}

func (s *Stack[T]) Length() int {
	return s.length
}

func (s *Stack[T]) Push(value T) {
	s.length++
	n := &node[T]{data: value}
	if s.root == nil {
		s.root = n
		return
	}

	n.next = s.root
	s.root = n
}

func (s *Stack[T]) Pop() (T, error) {
	var value T
	if s.root == nil {
		return value, errors.New("empty stack")
	}

	s.length--

	value = s.root.data
	s.root = s.root.next

	return value, nil
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

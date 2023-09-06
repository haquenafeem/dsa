package stack

import "errors"

type Stack struct {
	root   *node
	length int
}

func (s *Stack) Length() int {
	return s.length
}

func (s *Stack) Push(value int) {
	s.length++
	n := &node{data: value}
	if s.root == nil {
		s.root = n
		return
	}

	n.next = s.root
	s.root = n
}

func (s *Stack) Pop() (int, error) {
	if s.root == nil {
		return -1, errors.New("empty stack")
	}

	value := s.root.data
	s.root = s.root.next

	return value, nil
}

func New() *Stack {
	return &Stack{}
}

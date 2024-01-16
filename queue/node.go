package queue

type node[T any] struct {
	data int
	next *node[T]
}

func (n *node[T]) hasNext() bool {
	return n.next != nil
}

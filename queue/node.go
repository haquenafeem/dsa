package queue

type node struct {
	data int
	next *node
}

func (n *node) hasNext() bool {
	return n.next != nil
}

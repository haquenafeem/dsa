package hashtable

type node[T any] struct {
	key   string
	value T
	next  *node[T]
}

func (n *node[T]) insert(key string, value T) {
	nd := &node[T]{key: key, value: value}

	curr := n

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = nd
}

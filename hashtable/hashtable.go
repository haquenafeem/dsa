package hashtable

import "fmt"

const DefaultSize = 4

type HashTable[T any] struct {
	items []*node[T]
}

func New[T any]() *HashTable[T] {
	return &HashTable[T]{
		items: make([]*node[T], DefaultSize),
	}
}

func NewWithSize[T any](size int) *HashTable[T] {
	return &HashTable[T]{
		items: make([]*node[T], size),
	}
}

func (h *HashTable[T]) String() string {
	str := ""
	for i := 0; i < len(h.items); i++ {
		curr := h.items[i]
		str += fmt.Sprintf("[%d] -> ", i)

		for curr != nil {
			str += fmt.Sprintf("(%s,%v) -> ", curr.key, curr.value)
			curr = curr.next
		}

		str = str[:len(str)-3]
		str += "\n"
	}

	return str
}

func (h *HashTable[T]) hash(key string) int {
	return len(key) % DefaultSize
}

func (h *HashTable[T]) Add(key string, item T) {
	index := h.hash(key)

	head := h.items[index]

	if head == nil {
		h.items[index] = &node[T]{key: key, value: item}
		return
	}

	head.insert(key, item)
}

func (h *HashTable[T]) Get(key string) (T, bool) {
	index := h.hash(key)

	head := h.items[index]

	var value T
	if head == nil {
		return value, false
	}

	if head.key == key {
		return head.value, true
	}

	curr := head.next

	for curr != nil {
		if curr.key == key {
			return curr.value, true
		}
	}

	return value, false
}

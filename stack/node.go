package stack

type node[T any] struct {
	data T
	next *node[T]
}

package stack

type node[T any] struct {
	data int
	next *node[T]
}

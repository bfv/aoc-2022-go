package lib

type (
	Deque[T any] struct {
		first  *node[T]
		last   *node[T]
		length int
	}
	node[T any] struct {
		value T
		prev  *node[T]
	}
)

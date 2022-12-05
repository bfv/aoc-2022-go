package lib

// A Stack is a LIFO collection.
type Stack[T any] struct {
	values []T
	depth  int
	max    int
}

func (s *Stack[T]) Push(v T) error {

	s.values = append(s.values, v)
	s.depth++
	if s.depth > s.max {
		s.max = s.depth
	}

	return nil
}

// Pop returns the last inserted value of the stack and removes this from the stack
func (s *Stack[T]) Pop() T {
	v := s.values[len(s.values)-1:]
	s.values = s.values[:len(s.values)-1]
	s.depth--
	return v[0]
}

// Peek returns the last inserted value of the stack (no removal)
func (s *Stack[T]) Peek() T {
	return s.values[len(s.values)-1:][0]
}

// Returns true is the stack holds no values
func (s *Stack[T]) IsEmpty() bool {
	return s.depth == 0
}

// returns the current depth of the stack
func (s *Stack[T]) Depth() int {
	return s.depth
}

// returns the maximum depth of the stack
func (s *Stack[T]) MaxDepth() int {
	return s.max
}

// Content returns a copy of the slice holding the value
func (s *Stack[T]) Content() []T {
	c := make([]T, len(s.values))
	copy(c, s.values)
	return c
}

// Reset empties the Stack and reset depth & max depth (but NOT the datatype)
func (s *Stack[T]) Reset() {
	s.values = make([]T, 0)
	s.depth, s.max = 0, 0
}

func (s *Stack[T]) Copy() Stack[T] {
	b := Stack[T]{}
	b.values = append(b.values, s.values...)
	return b
}

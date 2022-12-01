package lib

import (
	"errors"
	"fmt"
	"reflect"
)

// A Stack is a LIFO collection.
type Stack struct {
	values   []interface{}
	depth    int
	max      int
	datatype string
}

// Push add a value to the last position of the stack
func (s *Stack) Push(v interface{}) error {

	dt := reflect.TypeOf(v).String()
	if s.datatype != "" && s.datatype != dt {
		return errors.New(fmt.Sprint("type should be: ", s.datatype, ", found: ", dt))
	}

	s.values = append(s.values, v)
	s.depth++
	if s.depth > s.max {
		s.max = s.depth
	}

	return nil
}

// Pop returns the last inserted value of the stack and removes this from the stack
func (s *Stack) Pop() interface{} {
	v := s.values[len(s.values)-1:]
	s.values = s.values[:len(s.values)-1]
	s.depth--
	return v[0]
}

// Peek returns the last inserted value of the stack (no removal)
func (s *Stack) Peek() interface{} {
	return s.values[len(s.values)-1:][0]
}

// Returns true is the stack holds no values
func (s *Stack) IsEmpty() bool {
	return s.depth == 0
}

// returns the current depth of the stack
func (s *Stack) Depth() int {
	return s.depth
}

// returns the maximum depth of the stack
func (s *Stack) MaxDepth() int {
	return s.max
}

// Content returns a copy of the slice holding the value
func (s *Stack) Content() []interface{} {
	c := make([]interface{}, len(s.values))
	copy(c, s.values)
	return c
}

// create a Stack which checks the type
func (s Stack) New(datatype string) Stack {
	s1 := Stack{datatype: datatype}
	return s1
}

// Reset empties the Stack and reset depth & max depth (but NOT the datatype)
func (s *Stack) Reset() {
	s.values = make([]interface{}, 0)
	s.depth, s.max = 0, 0
}

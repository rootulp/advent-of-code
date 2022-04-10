package main

type Stack[T string | int] struct {
	slice []T
}

func (s *Stack[T]) Push(element T) {
	s.slice = append(s.slice, element)
}
func (s *Stack[T]) Pop() (popped T) {
	popped = s.Peek()
	s.slice = s.slice[:len(s.slice)-1]
	return popped
}
func (s *Stack[T]) Peek() (top T) {
	return s.slice[len(s.slice)-1]
}
func (s *Stack[T]) Len() int {
	return len(s.slice)
}

// OperatorStack and OperandStack can be consolidated into a universal
// stack when Go supports generics
type OperatorStack struct {
	slice []string
}

func (s *OperatorStack) Push(element string) {
	s.slice = append(s.slice, element)
}
func (s *OperatorStack) Pop() (popped string) {
	popped = s.Peek()
	s.slice = s.slice[:len(s.slice)-1]
	return popped
}
func (s *OperatorStack) Peek() (top string) {
	return s.slice[len(s.slice)-1]
}
func (s *OperatorStack) Len() int {
	return len(s.slice)
}

type OperandStack struct {
	slice []int
}

func (s *OperandStack) Push(element int) {
	s.slice = append(s.slice, element)
}
func (s *OperandStack) Pop() (popped int) {
	popped = s.Peek()
	s.slice = s.slice[:len(s.slice)-1]
	return popped
}
func (s *OperandStack) Peek() (top int) {
	return s.slice[len(s.slice)-1]
}
func (s *OperandStack) Len() int {
	return len(s.slice)
}

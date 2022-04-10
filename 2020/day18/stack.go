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

package stack

type IStack[T any] interface {
	Push(item T)
	Pop() T
	Length() int
	IsEmpty() bool
	Peek() T
	PeekTail() T
	PushHead(item T)
}

type Stack[T any] struct {
	items []T
}

// New creates a new instance of Stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack[T]) Length() int {
	return len(s.items)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Peek() T {
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) PeekTail() T {
	return s.items[0]
}

func (s *Stack[T]) PushHead(item T) {
	s.items = append([]T{item}, s.items...)
}

package queue

type IQueue[T any] interface {
	Push(item T)
	Pop() T
	Length() int
	IsEmpty() bool
	Peek() T
	PeekTail() T
	PushHead(item T)
}

type Queue[T any] struct {
	items []T
}

// New creates a new instance of Queue.
//
// No parameters.
// Returns a pointer to the newly created Queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Push adds an element to the end of the queue.
//
// item T: the element to be added to the queue.
func (q *Queue[T]) Push(item T) {
	q.items = append(q.items, item)
}

// Pop removes and returns the first element from the queue.
//
// No parameters.
// Returns the element removed from the queue.
func (q *Queue[T]) Pop() T {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Length returns the number of elements in the queue.
//
// No parameters.
// Returns an integer representing the number of elements in the queue.
func (q *Queue[T]) Length() int {
	return len(q.items)
}

// IsEmpty checks if the queue is empty.
//
// No parameters.
// Returns a boolean indicating if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Peek returns the first element in the queue without removing it.
//
// No parameters.
// Returns the first element in the queue.
func (q *Queue[T]) Peek() T {
	return q.items[0]
}

// PeekTail returns the last element in the queue without removing it.
//
// No parameters.
// Returns the last element in the queue.
func (q *Queue[T]) PeekTail() T {
	return q.items[len(q.items)-1]
}

// PushHead adds an element to the beginning of the queue.
//
// item T: the element to be added to the front of the queue.
func (q *Queue[T]) PushHead(item T) {
	q.items = append([]T{item}, q.items...)
}

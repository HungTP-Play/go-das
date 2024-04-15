package sort

type HeapSort[T Number] struct {
	arr []T
}

func NewHeapSort[T Number](arr []T) *HeapSort[T] {
	return &HeapSort[T]{arr: arr}
}

// Sort sorts the elements using the HeapSort algorithm.
//
// direction: SORT_ASC or SORT_DESC
func (s *HeapSort[T]) Sort(direction int) []T {
	if direction == SORT_ASC {
		s.heapSortAsc()
	} else {
		s.heapSortDesc()
	}
	return s.arr
}

// heapSortAsc sorts the elements using the HeapSort algorithm in ascending order.
func (s *HeapSort[T]) heapSortAsc() {
	for i := len(s.arr) / 2; i >= 0; i-- {
		s.heapifyAsc(i, len(s.arr))
	}
	for i := len(s.arr) - 1; i >= 0; i-- {
		s.arr[0], s.arr[i] = s.arr[i], s.arr[0]
		s.heapifyAsc(0, i)
	}
}

// heapSortDesc sorts the elements using the HeapSort algorithm in descending order.
func (s *HeapSort[T]) heapSortDesc() {
	for i := len(s.arr) / 2; i >= 0; i-- {
		s.heapifyDesc(i, len(s.arr))
	}
	for i := len(s.arr) - 1; i >= 0; i-- {
		s.arr[0], s.arr[i] = s.arr[i], s.arr[0]
		s.heapifyDesc(0, i)
	}
}

// heapifyAsc sorts the elements using the HeapSort algorithm in ascending order.
func (s *HeapSort[T]) heapifyAsc(i, n int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && s.arr[left] > s.arr[largest] {
		largest = left
	}
	if right < n && s.arr[right] > s.arr[largest] {
		largest = right
	}
	if largest != i {
		s.arr[i], s.arr[largest] = s.arr[largest], s.arr[i]
		s.heapifyAsc(largest, n)
	}
}

func (s *HeapSort[T]) heapifyDesc(i, n int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && s.arr[left] < s.arr[largest] {
		largest = left
	}
	if right < n && s.arr[right] < s.arr[largest] {
		largest = right
	}
	if largest != i {
		s.arr[i], s.arr[largest] = s.arr[largest], s.arr[i]
		s.heapifyDesc(largest, n)
	}
}

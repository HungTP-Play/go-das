package sort

type QuickSort[T Number] struct {
	arr []T
}

func NewQuickSort[T Number](arr []T) *QuickSort[T] {
	return &QuickSort[T]{arr: arr}
}

// Sort sorts the elements using the QuickSort algorithm.
//
// direction: SORT_ASC or SORT_DESC
func (s *QuickSort[T]) Sort(direction int) []T {
	if direction == SORT_ASC {
		s.quickSortAsc(0, len(s.arr)-1)
	} else {
		s.quickSortDesc(0, len(s.arr)-1)
	}
	return s.arr
}

// quickSortAsc sorts the elements using the QuickSort algorithm in ascending order.
func (s *QuickSort[T]) quickSortAsc(low, high int) {
	if low < high {
		pivot := s.partitionAsc(low, high)
		s.quickSortAsc(low, pivot-1)
		s.quickSortAsc(pivot+1, high)
	}
}

// quickSortDesc sorts the elements using the QuickSort algorithm in descending order.
func (s *QuickSort[T]) quickSortDesc(low, high int) {
	if low < high {
		pivot := s.partitionDesc(low, high)
		s.quickSortDesc(low, pivot-1)
		s.quickSortDesc(pivot+1, high)
	}
}

// partitionAsc partitions the array in ascending order.
func (s *QuickSort[T]) partitionAsc(low, high int) int {
	pivot := s.arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if s.arr[j] < pivot {
			i++
			s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
		}
	}
	s.arr[i+1], s.arr[high] = s.arr[high], s.arr[i+1]
	return i + 1
}

func (s *QuickSort[T]) partitionDesc(low, high int) int {
	pivot := s.arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if s.arr[j] > pivot {
			i++
			s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
		}
	}
	s.arr[i+1], s.arr[high] = s.arr[high], s.arr[i+1]
	return i + 1
}

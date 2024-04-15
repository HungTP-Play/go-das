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

// ----------------------

type QuickSortBy[T any] struct {
	arr     []T
	compare Compare[T]
}

func NewQuickSortBy[T any](arr []T, compare Compare[T]) *QuickSortBy[T] {
	return &QuickSortBy[T]{arr: arr, compare: compare}
}

// Sort sorts the elements using the QuickSort algorithm.
func (s *QuickSortBy[T]) Sort() []T {
	s.quickSort(0, len(s.arr)-1)
	return s.arr
}

// quickSort sorts the elements using the QuickSort algorithm.
//
// low: the lower index of the subarray to be sorted.
// high: the higher index of the subarray to be sorted.
func (s *QuickSortBy[T]) quickSort(low, high int) {
	if low < high {
		pivot := s.partition(low, high)
		s.quickSort(low, pivot-1)
		s.quickSort(pivot+1, high)
	}
}

// partition partitions the array into two sections based on the pivot element.
//
// low: the lower index of the subarray to be partitioned.
// high: the higher index of the subarray to be partitioned.
// Returns the index where the pivot element is placed after partitioning.
func (s *QuickSortBy[T]) partition(low, high int) int {
	pivot := s.arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if s.compare(s.arr[j], pivot) < 0 {
			i++
			s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
		}
	}
	s.arr[i+1], s.arr[high] = s.arr[high], s.arr[i+1]
	return i + 1
}

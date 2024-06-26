package sort

type MergeSort[T Number] struct {
	arr       []T
	direction int
}

func NewMergeSort[T Number](arr []T) *MergeSort[T] {
	return &MergeSort[T]{arr: arr}
}

// Sort sorts the elements using the MergeSort algorithm.
//
// direction: SORT_ASC or SORT_DESC
func (s *MergeSort[T]) Sort(direction int) []T {
	s.direction = direction
	if direction == SORT_ASC {
		s.mergeSortAsc(0, len(s.arr)-1)
	} else {
		s.mergeSortDesc(0, len(s.arr)-1)
	}
	return s.arr
}

// mergeSortAsc sorts the elements using the MergeSort algorithm in ascending order.
func (s *MergeSort[T]) mergeSortAsc(low, high int) {
	if low < high {
		mid := low + (high-low)/2
		s.mergeSortAsc(low, mid)
		s.mergeSortAsc(mid+1, high)
		s.merge(low, mid, high)
	}
}

// mergeSortDesc sorts the elements using the MergeSort algorithm in descending order.
func (s *MergeSort[T]) mergeSortDesc(low, high int) {
	if low < high {
		mid := low + (high-low)/2
		s.mergeSortDesc(low, mid)
		s.mergeSortDesc(mid+1, high)
		s.merge(low, mid, high)
	}
}

func (s *MergeSort[T]) merge(low, mid, high int) {
	left, right := low, mid+1

	for left <= mid && right <= high {
		switch {
		case s.direction == SORT_ASC && s.arr[left] <= s.arr[right]:
			left++
		case s.direction != SORT_ASC && s.arr[left] >= s.arr[right]:
			left++
		default:
			s.arr[left], s.arr[right] = s.arr[right], s.arr[left]
			left++
			right++
		}
	}

	for i := left; i <= mid; i++ {
		s.arr[i] = s.arr[i+1]
	}
}

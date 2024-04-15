package sort

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

const (
	SORT_ASC = iota
	SORT_DESC
)

// ISort is an interface for sorting algorithms.
type ISort[T Number] interface {
	Sort(direction int) []T
}

// Compare is a function for comparing two elements.
//
// # Should return -1, 0, or 1
//
// In Ascending order: a < b -> -1
// In Descending order: a < b -> 1
type Compare[T any] func(a, b T) int

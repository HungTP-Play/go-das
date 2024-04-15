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

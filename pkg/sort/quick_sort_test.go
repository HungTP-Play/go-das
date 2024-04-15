package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name      string
		arr       []int
		direction int
		want      []int
	}{
		{
			name:      "Test Ascending",
			arr:       []int{4, 2, 1, 5, 3},
			direction: SORT_ASC,
			want:      []int{1, 2, 3, 4, 5},
		},
		{
			name:      "Test Descending",
			arr:       []int{4, 2, 1, 5, 3},
			direction: SORT_DESC,
			want:      []int{5, 4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewQuickSort(tt.arr)
			got := s.Sort(tt.direction)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

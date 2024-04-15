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

func compareInts(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

func TestQuickSortBy(t *testing.T) {
	arr := []int{5, 3, 8, 4, 2, 7, 1, 6}
	sorter := NewQuickSortBy(arr, compareInts)
	sortedArr := sorter.Sort()

	expectedArr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i, v := range expectedArr {
		if sortedArr[i] != v {
			t.Errorf("Expected %d, but got %d", v, sortedArr[i])
		}
	}
}

func TestQuickSortByWithEmptyArray(t *testing.T) {
	arr := []int{}
	sorter := NewQuickSortBy(arr, compareInts)
	sortedArr := sorter.Sort()

	if len(sortedArr) != 0 {
		t.Errorf("Expected empty array, but got %v", sortedArr)
	}
}

func TestQuickSortByWithOneElement(t *testing.T) {
	arr := []int{1}
	sorter := NewQuickSortBy(arr, compareInts)
	sortedArr := sorter.Sort()

	if sortedArr[0] != 1 {
		t.Errorf("Expected 1, but got %d", sortedArr[0])
	}
}

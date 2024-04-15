package sort

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	testCases := []struct {
		name      string
		input     []int
		direction int
		expected  []int
	}{
		{
			name:      "Test Case 1: Ascending order",
			input:     []int{5, 3, 1, 4, 2},
			direction: SORT_ASC,
			expected:  []int{1, 2, 3, 4, 5},
		},
		{
			name:      "Test Case 2: Descending order",
			input:     []int{5, 3, 1, 4, 2},
			direction: SORT_DESC,
			expected:  []int{5, 4, 3, 2, 1},
		},
		{
			name:      "Test Case 3: Ascending order with duplicate values",
			input:     []int{5, 3, 1, 4, 2, 2, 3, 1, 5},
			direction: SORT_ASC,
			expected:  []int{1, 1, 2, 2, 3, 3, 4, 5, 5},
		},
		{
			name:      "Test Case 4: Descending order with duplicate values",
			input:     []int{5, 3, 1, 4, 2, 2, 3, 1, 5},
			direction: SORT_DESC,
			expected:  []int{5, 5, 4, 3, 3, 2, 2, 1, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hs := NewHeapSort(tc.input)
			result := hs.Sort(tc.direction)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

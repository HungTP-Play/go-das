package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	testCases := []struct {
		name      string
		input     []int
		direction int
		expected  []int
	}{
		{
			name:      "Test Case 1: Ascending Order",
			input:     []int{4, 3, 2, 1},
			direction: SORT_ASC,
			expected:  []int{1, 2, 3, 4},
		},
		{
			name:      "Test Case 2: Descending Order",
			input:     []int{1, 2, 3, 4},
			direction: SORT_DESC,
			expected:  []int{4, 3, 2, 1},
		},
		{
			name:      "Test Case 3: Empty Array",
			input:     []int{},
			direction: SORT_ASC,
			expected:  []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sorter := NewMergeSort(tc.input)
			result := sorter.Sort(tc.direction)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %+v, got %+v", tc.expected, result)
			}
		})
	}
}

package arrays

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse sorted",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "mixed values",
			input:    []int{3, -1, 2, 0, 5, 2},
			expected: []int{-1, 0, 2, 2, 3, 5},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MergeSort(tc.input)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("MergeSort(%v) = %v; want %v", tc.input, got, tc.expected)
			}
		})
	}
}

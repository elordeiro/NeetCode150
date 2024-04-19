package main

import "testing"

func Test215KthLargestElementInAnArray(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}
	for _, test := range tests {
		nums, k, expected := test.nums, test.k, test.expected
		actual := findKthLargest(nums, k)
		t.Run("Test215KthLargestElementInAnArray", func(t *testing.T) {
			if actual != expected {
				t.Errorf("Test Failed:\n\tActual  :%v\n\tExpected:%v\n", actual, expected)
			}
		})
	}
}

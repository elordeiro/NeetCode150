package main

import (
	"math"
	"testing"

	. "github.com/elordeiro/NeetCode150/Go/utils"
)

const null = math.MaxInt

func Test124BinaryTreeMaximumPathSum(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{-10, 9, 20, null, null, 15, 7}, 42},
		{[]int{-3}, -3},
		{[]int{2, -1}, 2},
		{[]int{1, 2}, 3},
		{[]int{1, -2, -3, 1, 3, -2, null, -1}, 3},
		{[]int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, null, 1}, 48},
	}

	for _, test := range tests {
		actual := maxPathSum(CreatTree(test.nums))
		t.Run("Test124BinaryTreeMaximumPathSum", func(t *testing.T) {
			if actual != test.expected {
				t.Errorf("Acutal  : %v", actual)
				t.Errorf("Expected: %v", test.expected)
			}
		})
	}
}

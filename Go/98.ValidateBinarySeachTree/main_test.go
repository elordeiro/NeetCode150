package main

import (
	"math"
	"testing"

	. "github.com/elordeiro/NeetCode150/Go/utils"
)

const null = math.MaxInt

func Test98ValidateBinarySearchTree(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{2, 1, 3}, true},
		{[]int{5, 1, 4, null, null, 3, 6}, false},
		{[]int{2, 2, 2}, false},
		{[]int{1, 1}, false},
		{[]int{0, -1}, true},
		{[]int{5, 14, null, 1}, false},
	}

	for _, test := range tests {
		root := CreatTree(test.nums)
		t.Run("Test98ValidateBinarySearchTree", func(t *testing.T) {
			actual := isValidBST(root)
			if actual != test.expected {
				t.Errorf("Actual  : %v", actual)
				t.Errorf("Expected: %v", test.expected)
			}
		})
	}
}

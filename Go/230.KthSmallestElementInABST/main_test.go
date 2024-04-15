package main

import (
	"math"
	"testing"

	. "github.com/elordeiro/NeetCode150/Go/utils"
)

const null = math.MaxInt

func Test230KthSmallestElementInABST(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{3, 1, 4, null, 2}, 1, 1},
		{[]int{3, 1, 4, null, 2}, 2, 2},
		{[]int{5, 3, 6, 2, 4, null, null, 1}, 3, 3},
		{[]int{1, null, 2}, 2, 2},
		{[]int{2, 1, 3}, 3, 3},
	}

	for _, test := range tests {
		root := CreatTree(test.nums)
		actual := kthSmallest(root, test.k)
		t.Run("Test230KthSmallestElementInABST", func(t *testing.T) {
			if actual != test.expected {
				t.Errorf("Acutal  : %v", actual)
				t.Errorf("Expected: %v", test.expected)
			}
		})
	}
}

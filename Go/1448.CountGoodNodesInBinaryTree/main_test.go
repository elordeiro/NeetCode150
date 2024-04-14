package main

import (
	. "github.com/elordeiro/NeetCode150/Go/utils"
	"math"
	"testing"
)

const null = math.MaxInt

func Test1448CountGoodNodesInBinaryTree(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 1, 4, 3, null, 1, 5}, 4},
		{[]int{3, 3, null, 4, 2}, 3},
		{[]int{1}, 1},
		{[]int{2, null, 4, 10, 8, null, null, 4}, 4},
		{[]int{9, null, 3, 6}, 1},
		{[]int{-1, 5, -2, 4, 4, 2, -2, null, null, -4, null, -2, 3, null, -2, 0, null, -1, null, -3, null, -4, -3, 3, null, null, null, null, null, null, null, 3, -3}, 5},
	}
	for _, test := range tests {
		root := CreatTree(test.nums)
		actual := goodNodes(root)
		t.Run("Good Nodes", func(t *testing.T) {
			if actual != test.expected {
				t.Errorf("Actual  : %v", actual)
				t.Errorf("Expected: %v", test.expected)
			}
		})
	}
}

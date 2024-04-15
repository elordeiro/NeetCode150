package main

import (
	"math"
	"testing"

	. "github.com/elordeiro/NeetCode150/Go/utils"
)

const null = math.MaxInt

func Test105ConstructBinaryTreeFromPreorderAndInorderTraversal(t *testing.T) {
	tests := []struct {
		preorder  []int
		postorder []int
		nums      []int
	}{
		{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, []int{3, 9, 20, null, null, 15, 7}},
		{[]int{-1}, []int{-1}, []int{-1}},
	}

	for _, test := range tests {
		actual := buildTree(test.preorder, test.postorder)
		expected := CreatTree(test.nums)
		t.Run("Test230KthSmallestElementInABST", func(t *testing.T) {
			if !CompareTrees(actual, expected) {
				t.Errorf("Acutal  : %v", actual.ToString())
				t.Errorf("Expected: %v", expected.ToString())
			}
		})
	}
}

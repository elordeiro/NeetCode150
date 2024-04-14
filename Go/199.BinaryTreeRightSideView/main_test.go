package main

import (
	"math"
	"slices"
	"testing"

	. "github.com/elordeiro/NeetCode150/Go/utils"
)

const null = math.MaxInt

func Test199BinaryTreeRightSideView(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
	}{
		{[]int{1, 2, 3, null, 5, null, 4}, []int{1, 3, 4}},
		{[]int{1, null, 3}, []int{1, 3}},
		{[]int{}, []int{}},
	}
	testname := "199.Test199BinaryTreeRightSideView"
	for _, test := range tests {
		root := CreatTree(test.nums1)
		actual := rightSideView(root)
		expected := test.nums2
		t.Run(testname, func(t *testing.T) {
			if !slices.Equal(actual, expected) {
				t.Errorf("Actual  : %v", actual)
				t.Errorf("Expected: %v", expected)
			}
		})
	}
}

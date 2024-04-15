package main

import (
	"math"
	"testing"

	. "github.com/elordeiro/NeetCode150/Go/utils"
)

const null = math.MaxInt

func Test297SerializeAndDeserializeBinaryTree(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
	}{
		{[]int{1, 2, 3, null, null, 4, 5}, []int{1, 2, 3, null, null, 4, 5}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1}, []int{1}},
		// {[]int{}, []int{}},
	}

	ser := Constructor()
	deser := Constructor()
	for _, test := range tests {
		root, expected := CreatTree(test.nums1), CreatTree(test.nums2)
		t.Run("Test297SerializeAndDeserializeBinaryTree", func(t *testing.T) {
			actual := deser.deserialize(ser.serialize(root))
			if !CompareTrees(actual, expected) {
				t.Errorf("Acutal  : %v", actual.ToString())
				t.Errorf("Expected: %v", expected.ToString())
			}
		})
	}
}

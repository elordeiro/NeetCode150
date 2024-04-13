package main

import (
	"fmt"
	"math"
	"reflect"

	. "github.com/elordeiro/NeetCode150/Go/utils"
)

var null = math.MaxInt

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	queue := NewDeque()
	queue.PushRight(root)
	for !queue.IsEmpty() {
		level := []int{}
		for range queue.Length() {
			node := queue.PopLeft().(*TreeNode)
			level = append(level, node.Val)
			if node.Left != nil {
				queue.PushRight(node.Left)
			}
			if node.Right != nil {
				queue.PushRight(node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

func main() {
	tests := []struct {
		nums1 []int
		nums2 [][]int
	}{
		{[]int{3, 9, 20, null, null, 15, 7}, [][]int{{3}, {9, 20}, {15, 7}}},
		{[]int{1}, [][]int{{1}}},
		{[]int{}, [][]int{}},
	}
	passedAll := true
	testOnly := 0

	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		nums1, expected := test.nums1, test.nums2
		root := CreatTree(nums1)
		actual := levelOrder(root)
		if !reflect.DeepEqual(actual, expected) {
			fmt.Printf("Test %v Failed\n", i+1)
			fmt.Printf("\tActual  :  %v\n", actual)
			fmt.Printf("\tExpected:  %v\n", expected)
			passedAll = false
		}
	}
	if passedAll {
		fmt.Println("All Tests Passed")
	}
}

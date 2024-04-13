package main

import (
	"fmt"
	"math"
	"strconv"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func creatTree(nums []int, idx int) *TreeNode {
	if idx >= len(nums) || nums[idx] == math.MaxInt {
		return nil
	}
	root := &TreeNode{Val: nums[idx]}
	root.Left = creatTree(nums, 2*idx)
	root.Right = creatTree(nums, 2*idx+1)
	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Print("[]")
		return
	}
	res := "["
	var recur func(root *TreeNode)
	recur = func(root *TreeNode) {
		if root == nil {
			return
		}
		res += strconv.Itoa(root.Val) + ", "
		recur(root.Left)
		recur(root.Right)
	}
	recur(root)
	res = res[:len(res)-2]
	res += "]"
	fmt.Printf("%v\n", res)
}

func printTestFailed(actual int, expected int, testNum int) {
	fmt.Printf("Test %v Failed:\n", testNum)
	fmt.Printf("\tActual  : %v\n", actual)
	fmt.Printf("\tExpected: %v\n", expected)
}

func compareTrees(tree1 *TreeNode, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}
	if tree1 == nil || tree2 == nil {
		return false
	}
	if tree1.Val != tree2.Val {
		return false
	}
	if compareTrees(tree1.Left, tree2.Left) {
		return compareTrees(tree1.Right, tree2.Right)
	}
	return false
}

func main() {
	tests := []struct {
		nums1 []int
		nums2 int
	}{
		{[]int{-1, 3, 9, 20, math.MaxInt, math.MaxInt, 15, 7}, 3},
	}
	passedAll := true
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		nums1, expected := test.nums1, test.nums2
		root := creatTree(nums1, 1)
		actual := maxDepth(root)
		if actual != expected {
			printTestFailed(actual, expected, i+1)
			passedAll = false
		}
	}
	if passedAll {
		fmt.Println("All Tests Passed")
	}
}

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

func isBalanced(root *TreeNode) bool {
	isBalanced := true
	var recur func(root *TreeNode) int
	recur = func(root *TreeNode) int {
		if root == nil {
			return -1
		}
		left := recur(root.Left)
		right := recur(root.Right)
		if left != right && left+1 != right && left != right+1 {
			isBalanced = false
		}
		return 1 + max(left, right)
	}
	recur(root)
	return isBalanced
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

func printTestFailed(actual bool, expected bool, testNum int) {
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
		nums2 bool
	}{
		{[]int{-1, 3, 9, 20, math.MaxInt, math.MaxInt, 15, 7}, true},
		{[]int{-1, 1, 2, 2, 3, 3, math.MaxInt, math.MaxInt, 4, 4}, false},
	}
	passedAll := true
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		nums1, expected := test.nums1, test.nums2
		root := creatTree(nums1, 1)
		actual := isBalanced(root)
		if actual != expected {
			printTestFailed(actual, expected, i+1)
			passedAll = false
		}
	}
	if passedAll {
		fmt.Println("All Tests Passed")
	}
}

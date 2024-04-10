package utils

import (
	"fmt"
	"strconv"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeFunction(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = treeFunction(root.Right), treeFunction(root.Left)
	return root
}

func creatTree(nums []int, idx int) *TreeNode {
	if idx >= len(nums) {
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

func printTestFailed(actual *TreeNode, expected *TreeNode, testNum int) {
	fmt.Printf("Test %v Failed:\n", testNum)
	fmt.Printf("\tActual  : ")
	printTree(actual)
	fmt.Printf("\tExpected: ")
	printTree(expected)
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

func mainTree() {
	tests := []struct {
		nums1 []int
		nums2 []int
	}{
		{[]int{-1, 4, 2, 7, 1, 3, 6, 9}, []int{-1, 4, 7, 2, 9, 6, 3, 1}},
	}
	passedAll := true
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		nums1, nums2 := test.nums1, test.nums2
		root := creatTree(nums1, 1)
		actual := treeFunction(root)
		expected := creatTree(nums2, 1)
		if !compareTrees(actual, expected) {
			printTestFailed(actual, expected, i+1)
			passedAll = false
		}
	}
	if passedAll {
		fmt.Println("All Tests Passed")
	}
}

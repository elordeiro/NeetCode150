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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var recur func(p *TreeNode, q *TreeNode)
	recur = func(p *TreeNode, q *TreeNode) {
		if p == nil && q == nil {
			return
		}
		if p == nil || q == nil {
			panic("")
		}
		if p.Val != q.Val {
			panic("")
		}
		recur(p.Left, q.Left)
		recur(p.Right, q.Right)
	}
	defer func() bool {
		if r := recover(); r != nil {
			return false
		}
		return true
	}()
	recur(p, q)
	return true
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
		nums1    []int
		nums2    []int
		expected bool
	}{
		{[]int{-1, 1, 2, 3}, []int{-1, 1, 2, 3}, true},
		{[]int{-1, 1, 2}, []int{-1, 1, math.MaxInt, 2}, false},
		{[]int{-1, 1, 2, 1}, []int{-1, 1, 1, 2}, false},
	}
	passedAll := true
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		nums1, nums2, expected := test.nums1, test.nums2, test.expected
		root1 := creatTree(nums1, 1)
		root2 := creatTree(nums2, 1)
		actual := isSameTree(root1, root2)
		if actual != expected {
			printTestFailed(actual, expected, i+1)
			passedAll = false
		}
	}
	if passedAll {
		fmt.Println("All Tests Passed")
	}
}

package main

import (
	"fmt"
	"github.com/elordeiro/NeetCode150/Go/utils"
	"math"
	"strconv"
)

var null int = math.MaxInt

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}
	if root == nil {
		return false
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func creatTree(nums []int) *TreeNode {
	if nums == nil {
		return nil
	}
	idx := 0
	maxIdx := len(nums)
	head := &TreeNode{Val: nums[idx]}
	idx++
	queue := utils.NewDeque()
	queue.PushRight(head)

	for !queue.IsEmpty() {
		curr := queue.PopLeft()
		var left int
		var right int
		if idx < maxIdx {
			left = nums[idx]
			idx++
		} else {
			left = null
		}
		if idx < maxIdx {
			right = nums[idx]
			idx++
		} else {
			right = null
		}
		if left != null {
			curr.(*TreeNode).Left = &TreeNode{Val: left}
			queue.PushRight(curr.(*TreeNode).Left)
		}
		if right != null {
			curr.(*TreeNode).Right = &TreeNode{Val: right}
			queue.PushRight(curr.(*TreeNode).Right)
		}
	}
	return head
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Print("[]")
		return
	}

	res := "["
	toPrint := utils.NewDeque()
	queue := utils.NewDeque()
	queue.PushRight(root)

	for !queue.IsEmpty() {
		curr := queue.PopLeft()
		if curr.(*TreeNode) == nil {
			toPrint.PushRight(null)
			continue
		}
		toPrint.PushRight(curr.(*TreeNode).Val)
		queue.PushRight(curr.(*TreeNode).Left)
		queue.PushRight(curr.(*TreeNode).Right)
	}

	for !toPrint.IsEmpty() && toPrint.PeekRight().(int) == null {
		toPrint.PopRight()
	}

	for !toPrint.IsEmpty() {
		curr := toPrint.PopLeft()
		if curr.(int) == null {
			res += "null, "
			continue
		}
		res += strconv.Itoa(curr.(int)) + ", "
	}

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
		{[]int{3, 4, 5, 1, 2}, []int{4, 1, 2}, true},
		{[]int{3, 4, 5, 1, 2, null, null, null, null, 0}, []int{4, 1, 2}, false},
		{[]int{1, 1}, []int{1}, true},
		{[]int{3, 4, 5, 1, 2}, []int{3, 1, 2}, false},
		{[]int{
			-2, -3, -3, -4, -4, -4, -4, -5, -3, null, -5, -5, -5, -5, null, -4, -6, -2, null, -4, -6, -6, -4, -6, -4, -6, -4, null, -3, -5, -5, null, null, null, null, -7, -5, -7, -7, -5, -3, null, -7, -5, null, -5, -7, -5, -5, null, null, null, null, -6, -4, -6, -6, -4, -4, -8, -6, -8, -8, -6, -4, -4, -2, null, null, null, -4, -6, -4, null, -6, -4, -4, -6, -6, -7, -5, -5, null, null, -7, null, null, -3, -3, -5, null, null, null, -7, null, null, -7, -9, -9, -7, null, -5, null, null, -3, null, -1, null, null, -5, -5, -5, -3, -7, -5, null, null, -5, null, null, -5, null, -5, -8, -6, null, -6, -6, null, null, null, -2, null, null, null, null, null, -8, null, -6, -6, -10, null, -10, null, -6, null, null, -4, null, null, null, 0, -6, null, -6, null, -4, null, null, null, -8, -6, -6, -4, -6, null, null, null, null, null, -9, -7, -7, -5, null, null, null, null, null, null, -7, -7, -5, -5, null, null, -11, null, null, null, null, null, -3, -3, null, null, null, null, -7, -5, null, null, null, null, null, -5, null, -5, null, -3, null, null, null, null, -8, -6, -6, -6, null, null, null, null, -8, -6, null, -4, null, null, -12, null, -4, -4, -2, -2, null, -6, null, null, null, -6, null, null, null, null, null, null, -5, null, -7, -5, null, null, -7, -7, null, null, null, null, null, null, null, null, -3, null, null, null, null, -1, null, null, null, null, null, null, -8, -6, null, null, -8, null, null, null, -2, null, null, null, null, null, -7, null, null, null, null, -3, null, null, -4,
		}, []int{
			-3, -4, -4, null, null, -3, null, -2, null, null, -3, -4,
		}, true},
	}
	passedAll := true
	testOnly := 5
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		nums1, nums2, expected := test.nums1, test.nums2, test.expected
		root1 := creatTree(nums1)
		root2 := creatTree(nums2)
		actual := isSubtree(root1, root2)
		if actual != expected {
			printTestFailed(actual, expected, i+1)
			passedAll = false
		}
	}
	if passedAll {
		fmt.Println("All Tests Passed")
	}
}

package utils

import (
	"fmt"
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

func treeFunction(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = treeFunction(root.Right), treeFunction(root.Left)
	return root
}

func CreatTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	idx := 0
	maxIdx := len(nums)
	head := &TreeNode{Val: nums[idx]}
	idx++
	queue := NewDeque()
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

func PrintTree(root *TreeNode) {
	if root == nil {
		fmt.Print("[]")
		return
	}

	res := "["
	toPrint := NewDeque()
	queue := NewDeque()
	queue.PushRight(root)

	for !queue.IsEmpty() {
		curr := queue.PopLeft()
		if curr == null {
			toPrint.PushRight(null)
			continue
		}
		toPrint.PushRight(curr.(*TreeNode).Val)
		queue.PushRight(curr.(*TreeNode).Left)
		queue.PushRight(curr.(*TreeNode).Right)
	}

	for !toPrint.IsEmpty() && toPrint.PeekRight() != null {
		toPrint.PopRight()
	}

	for !toPrint.IsEmpty() {
		curr := toPrint.PopLeft()
		if curr == null {
			res += "null, "
			continue
		}
		res += strconv.Itoa(curr.(int))
	}

	res = res[:len(res)-2]
	res += "]"
	fmt.Printf("%v\n", res)
}

func PrintTestFailed(actual *TreeNode, expected *TreeNode, testNum int) {
	fmt.Printf("Test %v Failed:\n", testNum)
	fmt.Printf("\tActual  : ")
	PrintTree(actual)
	fmt.Printf("\tExpected: ")
	PrintTree(expected)
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
		root := CreatTree(nums1)
		actual := treeFunction(root)
		expected := CreatTree(nums2)
		if !compareTrees(actual, expected) {
			PrintTestFailed(actual, expected, i+1)
			passedAll = false
		}
	}
	if passedAll {
		fmt.Println("All Tests Passed")
	}
}

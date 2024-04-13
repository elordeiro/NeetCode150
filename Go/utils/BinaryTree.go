package utils

import (
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

func CreatTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	idx := 0
	maxIdx := len(nums)
	head := &TreeNode{Val: nums[idx]}
	idx++
	queue := Deque()
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

func (root *TreeNode) ToString() string {
	if root == nil {
		return "[]"
	}

	res := "["
	toPrint := Deque()
	queue := Deque()
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

	for !toPrint.IsEmpty() && toPrint.PeekRight() == null {
		toPrint.PopRight()
	}

	for !toPrint.IsEmpty() {
		curr := toPrint.PopLeft()
		if curr == null {
			res += "null, "
			continue
		}
		res += strconv.Itoa(curr.(int)) + ", "
	}

	res = res[:len(res)-2]
	res += "]"
	return res
}

func CompareTrees(tree1 *TreeNode, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}
	if tree1 == nil || tree2 == nil {
		return false
	}
	if tree1.Val != tree2.Val {
		return false
	}
	if CompareTrees(tree1.Left, tree2.Left) {
		return CompareTrees(tree1.Right, tree2.Right)
	}
	return false
}

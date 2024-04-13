package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/elordeiro/NeetCode150/Go/utils"
)

var null int = math.MaxInt

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var recur func(*TreeNode) *TreeNode
	recur = func(root *TreeNode) *TreeNode {
		if p.Val < root.Val && q.Val < root.Val {
			return recur(root.Left)
		}
		if p.Val > root.Val && q.Val > root.Val {
			return recur(root.Right)
		}
		return root
	}
	return recur(root)
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

func printTestFailed(actual *TreeNode, expected *TreeNode, testNum int) {
	fmt.Printf("Test %v Failed:\n", testNum)
	fmt.Printf("\tActual  : %v\n", actual.Val)
	fmt.Printf("\tExpected: %v\n", expected.Val)
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
		val1  int
		val2  int
		val3  int
	}{
		{[]int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5}, 2, 8, 6},
		{[]int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5}, 2, 4, 2},
		{[]int{2, 1}, 2, 1, 2},
	}
	passedAll := true
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		nums1, val1, val2, val3 := test.nums1, test.val1, test.val2, test.val3
		root := creatTree(nums1)
		actual := lowestCommonAncestor(root, &TreeNode{Val: val1}, &TreeNode{Val: val2})
		expected := &TreeNode{Val: val3}
		if actual.Val != expected.Val {
			printTestFailed(actual, expected, i+1)
			passedAll = false
		}
	}
	if passedAll {
		fmt.Println("All Tests Passed")
	}
}

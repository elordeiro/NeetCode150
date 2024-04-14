package main

import . "github.com/elordeiro/NeetCode150/Go/utils"

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	isGoodSum := 0
	var recur func(*TreeNode, int)
	recur = func(root *TreeNode, max int) {
		if root.Val >= max {
			isGoodSum++
			max = root.Val
		}
		if root.Left != nil {
			recur(root.Left, max)
		}
		if root.Right != nil {
			recur(root.Right, max)
		}
	}
	recur(root, root.Val)
	return isGoodSum
}

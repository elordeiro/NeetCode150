package main

import (
	. "github.com/elordeiro/NeetCode150/Go/utils"
)

var maxSum int

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := max(dfs(root.Left), 0), max(dfs(root.Right), 0)
	maxSum = max(maxSum, left+root.Val+right)
	return root.Val + max(left, right)
}

func maxPathSum(root *TreeNode) int {
	maxSum = -1001
	dfs(root)
	return maxSum
}

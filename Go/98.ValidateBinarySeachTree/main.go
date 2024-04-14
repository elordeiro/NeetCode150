package main

import (
	"math"

	. "github.com/elordeiro/NeetCode150/Go/utils"
)

const inf = math.MaxInt
const nInf = math.MinInt

func recur(root *TreeNode, least, greatest int) bool {
	if root == nil {
		return true
	}
	if least < root.Val && root.Val < greatest {
		return recur(root.Left, least, root.Val) && recur(root.Right, root.Val, greatest)
	}
	return false
}

func isValidBST(root *TreeNode) bool {
	return recur(root, nInf, inf)
}

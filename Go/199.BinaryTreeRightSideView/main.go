package main

import . "github.com/elordeiro/NeetCode150/Go/utils"

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	dq := Deque()
	dq.PushRight(root)
	for !dq.IsEmpty() {
		res = append(res, dq.PeekRight().(*TreeNode).Val)
		for range dq.Length() {
			node := dq.PopLeft().(*TreeNode)
			if node.Left != nil {
				dq.PushRight(node.Left)
			}
			if node.Right != nil {
				dq.PushRight(node.Right)
			}
		}
	}
	return res
}

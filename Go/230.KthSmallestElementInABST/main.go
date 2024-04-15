package main

import . "github.com/elordeiro/NeetCode150/Go/utils"

func kthSmallest(root *TreeNode, k int) int {
	curr := root
	stack := []*TreeNode{}
	for {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return curr.Val
		}
		curr = curr.Right
	}
}

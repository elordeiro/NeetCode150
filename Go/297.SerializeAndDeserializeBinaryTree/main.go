package main

import (
	"math"
	"strconv"
	"strings"

	. "github.com/elordeiro/NeetCode150/Go/utils"
)

const nul = ""

type Codec struct{}

func Constructor() Codec {
	return *new(Codec)
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	stack := []*TreeNode{}
	q := []*TreeNode{}
	q = append(q, root)
	for len(q) > 0 {
		for range len(q) {
			curr := q[0]
			q = q[1:]
			stack = append(stack, curr)
			if curr != nil {
				q = append(q, curr.Left)
				q = append(q, curr.Right)
			}
		}
	}

	for len(stack) > 0 && stack[len(stack)-1] == nil {
		stack = stack[:len(stack)-1]
	}

	res := ""

	for _, node := range stack {
		if node != nil {
			res += strconv.Itoa(node.Val) + ","
		} else {
			res += ","
		}
	}

	return res[:len(res)-1]
}

func quickInt(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return math.MaxInt
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	nodes := strings.Split(data, ",")
	n := len(nodes)

	root := &TreeNode{Val: quickInt(nodes[0])}
	q := []*TreeNode{}
	q = append(q, root)
	idx := 1
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		var left, right string
		if idx < n {
			left = nodes[idx]
		} else {
			left = nul
		}
		idx++
		if idx < n {
			right = nodes[idx]
		} else {
			right = nul
		}
		idx++

		if left != nul {
			curr.Left = &TreeNode{Val: quickInt(left)}
			q = append(q, curr.Left)
		}
		if right != nul {
			curr.Right = &TreeNode{Val: quickInt(right)}
			q = append(q, curr.Right)
		}
	}
	return root
}

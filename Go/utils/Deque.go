package utils

import (
	"fmt"
	"math"
)

type Node struct {
	val  any
	prev *Node
	next *Node
}

type deque struct {
	len   int
	left  *Node
	right *Node
}

func Deque() *deque {
	dq := new(deque)
	dq.left = new(Node)
	dq.right = new(Node)
	dq.left.next = dq.right
	dq.right.prev = dq.left
	return dq
}

func (dq *deque) PushLeft(val any) {
	if dq == nil {
		return
	}
	node := Node{val, nil, nil}
	node.prev = dq.left
	node.next = dq.left.next
	dq.left.next.prev = &node
	dq.left.next = &node
	dq.len++
}

func (dq *deque) PushRight(val any) {
	if dq == nil {
		return
	}
	node := Node{val, nil, nil}
	node.next = dq.right
	node.prev = dq.right.prev
	dq.right.prev.next = &node
	dq.right.prev = &node
	dq.len++
}

func (dq *deque) PopLeft() any {
	if dq == nil || dq.left.next == dq.right {
		return math.MinInt
	}
	first := dq.left.next
	dq.left.next = first.next
	first.next.prev = dq.left
	dq.len--
	return first.val
}

func (dq *deque) PopRight() any {
	if dq == nil || dq.left == dq.right {
		return math.MinInt
	}
	last := dq.right.prev
	dq.right.prev = last.prev
	last.prev.next = dq.right
	dq.len--
	return last.val
}

func (dq *deque) PeekLeft() any {
	return dq.left.next.val
}

func (dq *deque) PeekRight() any {
	return dq.right.prev.val
}

func (dq *deque) IsEmpty() bool {
	return dq.left.next == dq.right
}

func (dq *deque) Length() int {
	return dq.len
}

func (dq *deque) PrintDeque() {
	if dq == nil || dq.left.next == dq.right {
		return
	}

	var printList func(node *Node)
	printList = func(node *Node) {
		if node == dq.right {
			return
		}
		fmt.Println(node.val)
		printList(node.next)
	}
	printList(dq.left.next)
}

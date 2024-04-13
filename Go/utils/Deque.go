package utils

import (
	"fmt"
	"math"
)

type Node struct {
	val  interface{}
	prev *Node
	next *Node
}

type Deque struct {
	left  *Node
	right *Node
}

func NewDeque() *Deque {
	dq := new(Deque)
	dq.left = new(Node)
	dq.right = new(Node)
	dq.left.next = dq.right
	dq.right.prev = dq.left
	return dq
}

func (dq *Deque) PushLeft(val interface{}) {
	if dq == nil {
		return
	}
	node := Node{val, nil, nil}
	node.prev = dq.left
	node.next = dq.left.next
	dq.left.next.prev = &node
	dq.left.next = &node
}

func (dq *Deque) PushRight(val interface{}) {
	if dq == nil {
		return
	}
	node := Node{val, nil, nil}
	node.next = dq.right
	node.prev = dq.right.prev
	dq.right.prev.next = &node
	dq.right.prev = &node
}

func (dq *Deque) PopLeft() interface{} {
	if dq == nil || dq.left.next == dq.right {
		return math.MinInt
	}
	first := dq.left.next
	dq.left.next = first.next
	first.next.prev = dq.left
	return first.val
}

func (dq *Deque) PopRight() interface{} {
	if dq == nil || dq.left == dq.right {
		return math.MinInt
	}
	last := dq.right.prev
	dq.right.prev = last.prev
	last.prev.next = dq.right
	return last.val
}

func (dq *Deque) PeekLeft() interface{} {
	return dq.left.next.val
}

func (dq *Deque) PeekRight() interface{} {
	return dq.right.prev.val
}

func (dq *Deque) IsEmpty() bool {
	return dq.left.next == dq.right
}

func (dq *Deque) PrintDeque() {
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

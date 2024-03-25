package main

import (
	"fmt"
	"math"
	"slices"
)

type Node struct {
	val  int
	prev *Node
	next *Node
}

type Deque struct {
	left  *Node
	right *Node
}

func newDeque() *Deque {
	dq := new(Deque)
	dq.left = new(Node)
	dq.right = new(Node)
	dq.left.next = dq.right
	dq.right.prev = dq.left
	return dq
}

func (dq *Deque) pushLeft(val int) {
	if dq == nil {
		return
	}
	node := Node{val, nil, nil}
	node.prev = dq.left
	node.next = dq.left.next
	dq.left.next.prev = &node
	dq.left.next = &node
}

func (dq *Deque) pushRight(val int) {
	if dq == nil {
		return
	}
	node := Node{val, nil, nil}
	node.next = dq.right
	node.prev = dq.right.prev
	dq.right.prev.next = &node
	dq.right.prev = &node
}

func (dq *Deque) popLeft() int {
	if dq == nil || dq.left.next == dq.right {
		return math.MinInt
	}
	first := dq.left.next
	dq.left.next = first.next
	first.next.prev = dq.left
	return first.val
}

func (dq *Deque) popRight() int {
	if dq == nil || dq.left == dq.right {
		return math.MinInt
	}
	last := dq.right.prev
	dq.right.prev = last.prev
	last.prev.next = dq.right
	return last.val
}

func (dq *Deque) peekLeft() int {
	return dq.left.next.val
}

func (dq *Deque) peekRight() int {
	return dq.right.prev.val
}

func (dq *Deque) isEmpty() bool {
	return dq.left.next == dq.right
}

func (dq *Deque) printDeque() {
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

func maxSlidingWindow(nums []int, k int) []int {
	dq := newDeque()
	res := []int{}

	for j := range k {
		for !dq.isEmpty() && dq.peekRight() < nums[j] {
			dq.popRight()
		}
		dq.pushRight(nums[j])
	}

	res = append(res, dq.peekLeft())
	i := 0
	for j := k; j < len(nums); j++ {
		if dq.peekLeft() == nums[i] {
			dq.popLeft()
		}
		for !dq.isEmpty() && dq.peekRight() < nums[j] {
			dq.popRight()
		}
		dq.pushRight(nums[j])
		res = append(res, dq.peekLeft())
		i++
	}
	return res
}

func print_failed_test(i int, res []int, expected []int) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, -1}, 1, []int{1, -1}},
		{[]int{7, 2, 4}, 2, []int{7, 4}},
	}

	passed_all := true
	for i, test := range tests {
		nums, k, expected := test.nums, test.k, test.expected
		res := maxSlidingWindow(nums, k)
		if !slices.Equal(res, expected) {
			passed_all = false
			print_failed_test(i, res, expected)
		}
	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}

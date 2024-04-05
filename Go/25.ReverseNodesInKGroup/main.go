package main

import "fmt"

type Stack struct {
	stack []*ListNode
	len   int
	cap   int
}

func newStack() *Stack {
	return &Stack{
		stack: make([]*ListNode, 4, 4),
		len:   0,
		cap:   4,
	}
}

func (stk *Stack) push(node *ListNode) {
	if stk.len == stk.cap {
		stk.stack = append(stk.stack, make([]*ListNode, stk.len, stk.cap)...)
		stk.cap *= 2
	}
	stk.stack[stk.len] = node
	stk.len++
}

func (stk *Stack) pop() *ListNode {
	res := stk.stack[stk.len-1]
	stk.len--
	if stk.len < stk.cap/4 {
		stk.stack = stk.stack[:stk.cap/2]
		stk.cap /= 2
	}
	return res
}

func (stk *Stack) isEmpty() bool {
	return stk.len == 0
}

func (stk *Stack) peek() *ListNode {
	return stk.stack[stk.len-1]
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	curr := head
	stack := newStack()
	i := k
	for curr != nil && i > 0 {
		stack.push(curr)
		curr = curr.Next
		i--
	}
	top := stack.pop()
	newHead := top
	for !stack.isEmpty() {
		top.Next = stack.pop()
		top = top.Next
	}
	top.Next = curr
	oldTop := top

	for curr != nil {
		stack = newStack()
		i = k
		for curr != nil && i > 0 {
			stack.push(curr)
			curr = curr.Next
			i--
		}

		if stack.len < k {
			break
		}

		top = stack.pop()
		oldTop.Next = top

		for !stack.isEmpty() {
			top.Next = stack.pop()
			top = top.Next
		}
		oldTop = top
		top.Next = curr
	}
	return newHead
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func newNode(val int) *ListNode {
	node := new(ListNode)
	node.Val = val
	return node
}

func createList(nums []int) *ListNode {
	head := newNode(-1)
	curr := head
	for _, num := range nums {
		curr.Next = newNode((num))
		curr = curr.Next
	}
	return head.Next
}

func (head *ListNode) printList() {
	if head == nil {
		fmt.Println("]")
		return
	}
	if head.Next == nil {
		fmt.Printf("%v", head.Val)
	} else {
		fmt.Printf("%v, ", head.Val)
	}
	head.Next.printList()
}

func printFailedTest(actual *ListNode, expected *ListNode, testNum int) {
	fmt.Printf("Test %d failed:\n", testNum+1)
	fmt.Printf("    \tActual  : [")
	actual.printList()
	fmt.Printf("    \tExpected: [")
	expected.printList()
}

func main() {
	tests := []struct {
		nums1 []int
		k     int
		nums2 []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
	}

	passedAll := true
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		nums1, k, nums2 := test.nums1, test.k, test.nums2
		head := createList(nums1)
		actual := reverseKGroup(head, k)
		expected := createList(nums2)

		currA, currB := actual, expected
		broke := false
		for currA != nil && currB != nil {
			if currA.Val != currB.Val {
				passedAll = false
				printFailedTest(actual, expected, i)
				broke = true
				break
			}
			currA = currA.Next
			currB = currB.Next
		}
		if !broke {
			if currA != nil || currB != nil {
				passedAll = false
				printFailedTest(currA, currB, i)
			}
		}
	}
	if passedAll {
		fmt.Println("Passed all tests.")
	}
}

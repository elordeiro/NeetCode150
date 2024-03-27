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
	node := stk.stack[stk.len-1]
	stk.len--
	if stk.len < stk.cap/4 {
		stk.stack = stk.stack[:stk.cap/2]
		stk.cap /= 2
	}
	return node
}

func (stk *Stack) isEmpty() bool {
	return stk.len == 0
}

func (stk *Stack) peek() *ListNode {
	return stk.stack[stk.len-1]
}

func reoderList(head *ListNode) {
	stack := newStack()
	curr := head
	n := 0
	for curr != nil {
		stack.push(curr)
		curr = curr.Next
		n += 1
	}
	curr = head
	for range n / 2 {
		next := curr.Next
		curr.Next = stack.pop()
		curr = curr.Next
		curr.Next = next
		curr = curr.Next
	}
	curr.Next = nil
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
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 4, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}},
	}

	passedAll := true
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		nums, expectedNums := test.nums, test.expected
		actual := createList(nums)
		reoderList(actual)
		expected := createList(expectedNums)

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

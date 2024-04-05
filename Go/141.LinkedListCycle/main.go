package main

import "fmt"

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for {
		if fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
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

func createList(nums []int, creatCycle bool) *ListNode {
	head := newNode(-1)
	curr := head
	for _, num := range nums {
		curr.Next = newNode((num))
		curr = curr.Next
	}
	if creatCycle {
		curr.Next = head.Next
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

func printFailedTest(actual bool, expected bool, testNum int) {
	fmt.Printf("Test %d failed:\n", testNum+1)
	fmt.Printf("    \tActual  : %v\n", actual)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		nums1    []int
		expected bool
	}{
		{[]int{3, 2, 0, 4}, true},
		{[]int{1, 2}, false},
		{[]int{1}, false},
		{[]int{}, false},
	}

	passedAll := true
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		nums1, expected := test.nums1, test.expected
		head := createList(nums1, expected)
		actual := hasCycle(head)

		if actual != expected {
			passedAll = false
			printFailedTest(actual, expected, i)
		}
	}
	if passedAll {
		fmt.Println("Passed all tests.")
	}
}

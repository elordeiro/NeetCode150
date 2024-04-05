package main

import "fmt"

func findDuplicate(nums []int) int {
	slow, fast := 0, 0

	slow = nums[slow]
	fast = nums[nums[fast]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
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

func printFailedTest(actual int, expected int, testNum int) {
	fmt.Printf("Test %d failed:\n", testNum+1)
	fmt.Printf("    \tActual  : %v\n", actual)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		nums1    []int
		expected int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{3, 3, 3, 3, 3}, 3},
	}

	passedAll := true
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		nums1, expected := test.nums1, test.expected
		actual := findDuplicate(nums1)

		if actual != expected {
			passedAll = false
			printFailedTest(actual, expected, i)
		}
	}
	if passedAll {
		fmt.Println("Passed all tests.")
	}
}

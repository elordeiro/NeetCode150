package main

import "fmt"

func divmod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := &ListNode{Val: -1}
	curr := head
	for l1 != nil && l2 != nil {
		x, y := l1.Val, l2.Val
		l1, l2 = l1.Next, l2.Next
		d, m := divmod(x+y+carry, 10)
		curr.Next = &ListNode{Val: m}
		curr = curr.Next
		carry = d
	}
	for l1 != nil {
		x := l1.Val
		l1 = l1.Next
		d, m := divmod(x+carry, 10)
		curr.Next = &ListNode{Val: m}
		curr = curr.Next
		carry = d
	}
	for l2 != nil {
		y := l2.Val
		l2 = l2.Next
		d, m := divmod(y+carry, 10)
		curr.Next = &ListNode{Val: m}
		curr = curr.Next
		carry = d
	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return head.Next
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
		nums1    []int
		nums2    []int
		expected []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	passedAll := true
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		nums1, nums2, nums3 := test.nums1, test.nums2, test.expected
		list1 := createList(nums1)
		list2 := createList(nums2)
		actual := addTwoNumbers(list1, list2)
		expected := createList(nums3)

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

package main

import "fmt"

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	curr := head
	next := curr.next
	curr.next = nil
	prev := curr
	curr = next
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev
}

type ListNode struct {
	val  int
	next *ListNode
}

func newNode(val int) *ListNode {
	node := new(ListNode)
	node.val = val
	return node
}

func createList(nums []int) *ListNode {
	head := newNode(-1)
	curr := head
	for _, num := range nums {
		curr.next = newNode((num))
		curr = curr.next
	}
	return head.next
}

func (head *ListNode) printList() {
	if head == nil {
		fmt.Println("]")
		return
	}
	if head.next == nil {
		fmt.Printf("%v", head.val)
	} else {
		fmt.Printf("%v, ", head.val)
	}
	head.next.printList()
}

func print_failed_test(actual *ListNode, expected *ListNode, testNum int) {
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
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	passedAll := true
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		nums, expectedNums := test.nums, test.expected
		head := createList(nums)
		actual := reverseList(head)
		expected := createList(expectedNums)

		currA, currB := actual, expected
		broke := false
		for currA != nil && currB != nil {
			if currA.val != currB.val {
				passedAll = false
				print_failed_test(currA, currB, i)
				broke = true
				break
			}
			currA = currA.next
			currB = currB.next
		}
		if !broke {
			if currA != nil || currB != nil {
				passedAll = false
				print_failed_test(currA, currB, i)
			}
		}
	}
	if passedAll {
		fmt.Println("Passed all tests.")
	}
}

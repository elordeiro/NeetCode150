package main

import "fmt"

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	if n > 2 {
		half := n >> 1
		return mergeKLists([]*ListNode{mergeKLists(lists[:half]), mergeKLists(lists[half:])})
	}

	head := &ListNode{Val: -1000000}
	l1 := head
	l1.Next = lists[0]
	l2 := lists[1]

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			for l1.Next != nil && l1.Next.Val < l2.Val {
				l1 = l1.Next
			}
			prev := l1
			l1 = l1.Next
			prev.Next = l2
		} else {
			for l2.Next != nil && l2.Next.Val <= l1.Val {
				l2 = l2.Next
			}
			prev := l2
			l2 = l2.Next
			prev.Next = l1
		}
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
		nums1 [][]int
		nums2 []int
	}{
		{[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{[][]int{{}}, []int{}},
		{[][]int{{}}, []int{}},
		{[][]int{{-2, -1, -1, -1}, {}}, []int{-2, -1, -1, -1}},
	}

	passedAll := true
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		nums, nums2 := test.nums1, test.nums2
		lists := []*ListNode{}
		for i := 0; i < len(nums); i++ {
			list := nums[i]
			lists = append(lists, createList(list))
		}
		actual := mergeKLists(lists)
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

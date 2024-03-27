package main

import "fmt"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil || list2 == nil {
		if list1 == nil {
			return list2
		}
		return list1
	}
	var head *ListNode
	if list1.Val <= list2.Val {
		head = list1
	} else {
		head = list2
	}
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			for ; list1.Next != nil && list1.Next.Val <= list2.Val; list1 = list1.Next {
			}
			next := list1.Next
			list1.Next = list2
			list1 = next
		} else {
			for ; list2.Next != nil && list2.Next.Val <= list1.Val; list2 = list2.Next {
			}
			next := list2.Next
			list2.Next = list1
			list2 = next
		}
	}
	return head
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
		nums2 []int
		nums3 []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
	}

	passedAll := true
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		nums1, nums2, nums3 := test.nums1, test.nums2, test.nums3
		list1 := createList(nums1)
		list2 := createList(nums2)
		actual := mergeTwoLists(list1, list2)
		expected := createList(nums3)

		currA, currB := actual, expected
		broke := false
		for currA != nil && currB != nil {
			if currA.Val != currB.Val {
				passedAll = false
				printFailedTest(currA, currB, i)
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

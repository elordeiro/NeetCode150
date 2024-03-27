package main

import "fmt"

func get(m *map[*ListNode]*ListNode, key *ListNode) *ListNode {
	if val, ok := (*m)[key]; ok {
		return val
	}
	node := &ListNode{Val: 1}
	(*m)[key] = node
	return node
}

func copyRandomList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	m := make(map[*ListNode]*ListNode)
	curr := head
	for curr != nil {
		node := get(&m, curr)
		node.Val = curr.Val
		if curr.Next != nil {
			node.Next = get(&m, curr.Next)
		}
		if curr.Random != nil {
			node.Random = get(&m, curr.Random)
		}
		curr = curr.Next
	}
	return m[head]
}

type ListNode struct {
	Val    int
	Next   *ListNode
	Random *ListNode
}

func newNode(val int) *ListNode {
	node := new(ListNode)
	node.Val = val
	node.Random = nil
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
	}{
		{[]int{7, 13, 11, 10, 1}, []int{7, 13, 11, 10, 1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{3, 3, 3}, []int{3, 3, 3}},
		{[]int{}, []int{}},
	}

	passedAll := true
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		nums1, nums2 := test.nums1, test.nums2
		head := createList(nums1)
		actual := copyRandomList(head)
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

package main

import (
	"fmt"
	"math"
	"slices"
)

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

type LRUCache struct {
	cap   int
	len   int
	cache map[int]*Node
	head  *Node
	tail  *Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		cap:   capacity,
		len:   0,
		cache: make(map[int]*Node),
		tail:  &Node{Val: math.MinInt},
		head:  &Node{Val: math.MaxInt},
	}
	cache.tail.Next = cache.head
	cache.head.Prev = cache.tail
	return cache
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; ok {
		node := this.remNode(key)
		this.addNode(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		node := this.remNode(key)
		node.Val = value
		this.addNode(node)
		return
	}
	this.addNode(&Node{Key: key, Val: value})
	if this.len < this.cap {
		this.len++
		return
	}
	this.remNode(this.tail.Next.Key)
}

func (this *LRUCache) addNode(node *Node) {
	node.Next = this.head
	node.Prev = this.head.Prev
	node.Prev.Next = node
	node.Next.Prev = node
	this.cache[node.Key] = node
}

func (this *LRUCache) remNode(key int) *Node {
	node := this.cache[key]
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	delete(this.cache, key)
	return node
}

func newNode(val int) *Node {
	node := new(Node)
	node.Val = val
	return node
}

func createList(nums []int) *Node {
	head := newNode(-1)
	curr := head
	for _, num := range nums {
		curr.Next = newNode((num))
		curr = curr.Next
	}
	return head.Next
}

func (head *Node) printList() {
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

func printFailedTest(actual []int, expected []int, testNum int) {
	fmt.Printf("Test %d failed:\n", testNum+1)
	fmt.Printf("    \tActual  : %v\n", actual)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		nums1 [][]int
		nums2 []int
	}{
		{
			[][]int{
				{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4},
			},
			[]int{1, -1, -1, 3, 4},
		},
	}

	passedAll := true
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		nums1, expected := test.nums1, test.nums2

		actual := []int{}
		var constructor LRUCache
		for i, op := range nums1 {
			if i == 0 {
				constructor = Constructor(op[0])
				continue
			}
			if len(op) == 1 {
				actual = append(actual, constructor.Get(op[0]))
			} else {
				constructor.Put(op[0], op[1])
			}
		}
		if !slices.Equal(actual, expected) {
			passedAll = false
			printFailedTest(actual, expected, i)
		}
	}
	if passedAll {
		fmt.Println("Passed all tests.")
	}
}

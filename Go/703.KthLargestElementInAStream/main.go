package main

import "container/heap"

type IntHeap []int

type KthLargest struct {
	k    int
	heap *IntHeap
}

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return KthLargest{k: k, heap: h}
}

func (klg *KthLargest) Add(val int) int {
	heap.Push(klg.heap, val)
	if klg.heap.Len() > klg.k {
		heap.Pop(klg.heap)
	}
	return (*klg.heap)[0]
}
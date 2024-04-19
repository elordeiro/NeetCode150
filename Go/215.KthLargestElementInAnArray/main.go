package main

import "container/heap"

func findKthLargest(nums []int, k int) int {
	hp := IntHeap{}
	heap.Init(&hp)
	for _, num := range nums {
		heap.Push(&hp, num)
	}

	for range k - 1 {
		heap.Pop(&hp)
	}
	return heap.Pop(&hp).(int)
}

type IntHeap []int

func (heap IntHeap) Len() int           { return len(heap) }
func (heap IntHeap) Less(i, j int) bool { return heap[i] > heap[j] }
func (heap IntHeap) Swap(i, j int)      { heap[i], heap[j] = heap[j], heap[i] }
func (heap *IntHeap) Push(val any)      { *heap = append(*heap, val.(int)) }
func (heap *IntHeap) Pop() any {
	old := *heap
	n := len(old)
	x := old[n-1]
	*heap = old[:n-1]
	return x
}

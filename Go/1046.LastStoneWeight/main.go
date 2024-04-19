package main

import "container/heap"

func lastStoneWeight(stones []int) int {
	hp := IntHeap{}
	heap.Init(&hp)

	for _, stone := range stones {
		heap.Push(&hp, stone)
	}

	for hp.Len() > 1 {
		y := heap.Pop(&hp).(int)
		x := heap.Pop(&hp).(int)
		if x != y {
			heap.Push(&hp, y-x)
		}
	}
	if hp.Len() == 0 {
		return 0
	}
	return hp[0]
}

type IntHeap []int

func (heap IntHeap) Len() int           { return len(heap) }
func (heap IntHeap) Less(i, j int) bool { return heap[i] > heap[j] }
func (heap IntHeap) Swap(i, j int)      { heap[i], heap[j] = heap[j], heap[i] }
func (heap *IntHeap) Push(val any) {
	*heap = append(*heap, val.(int))
}
func (heap *IntHeap) Pop() any {
	old := *heap
	n := len(old)
	x := old[n-1]
	*heap = old[:n-1]
	return x
}

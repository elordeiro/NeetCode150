package main

import (
	"container/heap"

	. "github.com/elordeiro/NeetCode150/Go/utils"
)

type Tuple struct {
	Count    int
	NextTime int
}

func leastInterval(tasks []byte, n int) int {
	mapping := make([]int, 26, 26)
	for _, task := range tasks {
		mapping[task-65] += 1
	}
	hp := IntHeap{}
	heap.Init(&hp)
	for _, task := range mapping {
		if task > 0 {
			heap.Push(&hp, task)
		}
	}
	time := 0
	q := Deque()
	for hp.Len() > 0 || !q.IsEmpty() {
		for !q.IsEmpty() && q.PeekLeft().(Tuple).NextTime <= time {
			heap.Push(&hp, q.PopLeft().(Tuple).Count)
		}
		if hp.Len() > 0 {
			taskCount := heap.Pop(&hp).(int) - 1
			if taskCount != 0 {
				q.PushRight(Tuple{Count: taskCount, NextTime: time + n + 1})
			}
		}
		time++
	}
	return time
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

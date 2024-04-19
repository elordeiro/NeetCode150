package main

import "container/heap"

type MedianFinder struct {
	leftOfMedian IntHeap
	riteOfMedian IntHeap
	isEven       bool
}

func Constructor() MedianFinder {
	leftHeap := IntHeap{}
	riteHeap := IntHeap{}
	heap.Init(&leftHeap)
	heap.Init(&riteHeap)
	return MedianFinder{
		leftOfMedian: leftHeap,
		riteOfMedian: riteHeap,
		isEven:       true,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.isEven {
		heap.Push(&this.riteOfMedian, -this.leftOfMedian.PushPop(-num).(int))
	} else {
		heap.Push(&this.leftOfMedian, -this.riteOfMedian.PushPop(num).(int))
	}
	this.isEven = !this.isEven
}

func (this *MedianFinder) FindMedian() float64 {
	if this.isEven {
		return (-float64(this.leftOfMedian[0]) + float64(this.riteOfMedian[0])) / 2.0
	}
	return float64(this.riteOfMedian[0])
}

type IntHeap []int

func (heap IntHeap) Len() int           { return len(heap) }
func (heap IntHeap) Less(i, j int) bool { return heap[i] < heap[j] }
func (heap IntHeap) Swap(i, j int)      { heap[i], heap[j] = heap[j], heap[i] }
func (heap *IntHeap) Push(val any)      { *heap = append(*heap, val.(int)) }
func (heap *IntHeap) Pop() any {
	old := *heap
	n := len(old)
	x := old[n-1]
	*heap = old[:n-1]
	return x
}
func (hp *IntHeap) PushPop(num any) any {
	heap.Push(hp, num.(int))
	return heap.Pop(hp)
}

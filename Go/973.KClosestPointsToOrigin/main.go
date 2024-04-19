package main

import (
	"container/heap"
	"math"
)

func kClosest(points [][]int, k int) [][]int {
	hp := CustomHeap{}
	heap.Init(&hp)
	for _, point := range points {
		x, y := point[0], point[1]
		distance := math.Sqrt(float64(x*x) + float64(y*y))
		heap.Push(&hp, Tuple{Distance: distance, Point: [2]int{x, y}})
		if hp.Len() > k {
			heap.Pop(&hp)
		}
	}
	res := [][]int{}
	for _, t := range hp {
		res = append(res, t.Point[:])
	}
	return res
}

type Tuple struct {
	Distance float64
	Point    [2]int
}

type CustomHeap []Tuple

func (heap CustomHeap) Len() int           { return len(heap) }
func (heap CustomHeap) Less(i, j int) bool { return heap[i].Distance > heap[j].Distance }
func (heap CustomHeap) Swap(i, j int)      { heap[i], heap[j] = heap[j], heap[i] }
func (heap *CustomHeap) Push(tuple any)    { *heap = append(*heap, tuple.(Tuple)) }
func (heap *CustomHeap) Pop() any {
	old := *heap
	n := len(old)
	x := old[n-1]
	*heap = old[:n-1]
	return x
}

package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	dst, time int
}

type MinHeap []Edge

func (heap MinHeap) Len() int           { return len(heap) }
func (heap MinHeap) Swap(i, j int)      { heap[i], heap[j] = heap[j], heap[i] }
func (heap MinHeap) Less(i, j int) bool { return heap[i].time < heap[j].time }
func (heap *MinHeap) Push(val any)      { *heap = append(*heap, val.(Edge)) }
func (heap *MinHeap) Pop() any {
	old := *heap
	n := len(old)
	x := old[n-1]
	*heap = old[:n-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {
	graph := make(map[int][]Edge)
	for _, edge := range times {
		src, dst, time := edge[0], edge[1], edge[2]
		graph[src] = append(graph[src], Edge{dst, time})
	}

	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt
	}
	dist[k] = 0

	time := 0
	visitedArr := make([]bool, n+1)
	visitedNum := 0
	minHeap := &MinHeap{Edge{k, 0}}
	heap.Init(minHeap)

	for minHeap.Len() > 0 {
		current := heap.Pop(minHeap).(Edge)
		if visitedArr[current.dst] {
			continue
		}
		visitedArr[current.dst] = true
		visitedNum += 1
		time = current.time
		for _, edge := range graph[current.dst] {
			timeTodst := time + edge.time
			if timeTodst < dist[edge.dst] {
				dist[edge.dst] = timeTodst
				heap.Push(minHeap, Edge{edge.dst, timeTodst})
			}
		}
	}

	if visitedNum == n {
		return time
	}
	return -1
}

func main() {
	tests := []struct {
		times          [][]int
		n, k, expected int
	}{
		{[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, 2},
		{[][]int{{1, 2, 1}}, 2, 1, 1},
		{[][]int{{1, 2, 1}}, 2, 2, -1},
		{[][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 4}}, 3, 1, 3},
		{[][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 2}}, 3, 1, 2},
	}

	testOnly := 0
	failed := false
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		actual := networkDelayTime(test.times, test.n, test.k)
		if actual != test.expected {
			failed = true
			fmt.Printf("Test %v Failed\n\tActual  : %v\n\tExpected: %v\n", i+1, actual, test.expected)
		}
	}
	if !failed {
		fmt.Println("Passed All Tests.")
	}
}

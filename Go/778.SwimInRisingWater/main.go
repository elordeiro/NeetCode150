package main

import (
	"container/heap"
	"fmt"
)

type MinHeap [][3]int

func (heap MinHeap) Len() int           { return len(heap) }
func (heap MinHeap) Swap(i, j int)      { heap[i], heap[j] = heap[j], heap[i] }
func (heap MinHeap) Less(i, j int) bool { return heap[i][0] < heap[j][0] }
func (heap *MinHeap) Push(val any)      { *heap = append(*heap, val.([3]int)) }
func (heap *MinHeap) Pop() any {
	old := *heap
	n := len(old)
	x := old[n-1]
	*heap = old[:n-1]
	return x
}

func swimInWater(grid [][]int) int {
	N := len(grid)
	target := N - 1
	finalT := grid[N-1][N-1]
	minHeap := &MinHeap{{0, 0, 0}}

	for minHeap.Len() > 0 {
		top := heap.Pop(minHeap).([3]int)
		time, row, col := top[0], top[1], top[2]
		if row == target && col == target {
			return max(time, finalT)
		}
		if grid[row][col] == -1 {
			continue
		}
		time = max(time, grid[row][col])
		if row+1 < N {
			heap.Push(minHeap, [3]int{time, row + 1, col})
		}
		if col+1 < N {
			heap.Push(minHeap, [3]int{time, row, col + 1})
		}
		if row-1 > -1 {
			heap.Push(minHeap, [3]int{time, row - 1, col})
		}
		if col-1 > -1 {
			heap.Push(minHeap, [3]int{time, row, col - 1})
		}
		grid[row][col] = -1
	}
	return -1
}

func main() {
	tests := []struct {
		grid     [][]int
		expected int
	}{
		{[][]int{{0, 2}, {1, 3}}, 3},
		{[][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}, 16},
		{[][]int{{31, 28, 33, 0, 8, 57, 86, 99, 23, 98}, {25, 90, 20, 73, 34, 65, 29, 9, 42, 46}, {17, 84, 10, 4, 40, 5, 41, 21, 71, 79}, {13, 70, 69, 81, 63, 93, 77, 1, 94, 53}, {38, 87, 61, 50, 92, 2, 15, 95, 82, 68}, {44, 72, 88, 47, 27, 91, 37, 48, 83, 16}, {3, 30, 96, 66, 7, 58, 76, 54, 19, 64}, {85, 45, 60, 11, 51, 26, 6, 22, 74, 32}, {43, 12, 62, 59, 89, 52, 36, 97, 49, 78}, {75, 24, 14, 67, 56, 35, 55, 39, 80, 18}}, 78},
	}

	testOnly := 0
	failed := false
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		actual := swimInWater(test.grid)
		if actual != test.expected {
			failed = true
			fmt.Printf("Test %v Failed\n\tActual  : %v\n\tExpected: %v\n\t", i+1, actual, test.expected)
		}
	}
	if !failed {
		fmt.Println("Passed All Tests")
	}
}

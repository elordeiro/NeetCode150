package main

import (
	"container/heap"
	"fmt"
	"math"

	. "github.com/elordeiro/NeetCode150/Go/utils"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := make([][][2]int, n)
	for _, flight := range flights {
		source, dest, price := flight[0], flight[1], flight[2]
		graph[source] = append(graph[source], [2]int{dest, price})
	}

	minHeap := NewPriorityQueue(func(item1, item2 any) bool {
		return item1.([3]int)[0] < item2.([3]int)[0]
	})
	heap.Push(minHeap, [3]int{0, 0, src})

	inf := math.MaxInt
	visited := make([]int, n)
	for i := range n {
		visited[i] = inf
	}
	visited[src] = 0

	for minHeap.Len() > 0 {
		top := heap.Pop(minHeap).([3]int)
		stops, price, city := top[0], top[1], top[2]
		if stops > k {
			continue
		}
		for _, neighbor := range graph[city] {
			dest, newPrice := neighbor[0], neighbor[1]
			if totalPrice := price + newPrice; visited[dest] > totalPrice {
				visited[dest] = totalPrice
				heap.Push(minHeap, [3]int{stops + 1, totalPrice, dest})
			}
		}
	}
	if visited[dst] == inf {
		return -1
	}
	return visited[dst]
}

func main() {
	tests := []struct {
		n                     int
		flights               [][]int
		src, dst, k, expected int
	}{
		{4, [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}, 0, 3, 1, 700},
		{3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1, 200},
		{3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0, 500},
		{3, [][]int{{0, 1, 2}, {1, 2, 1}, {2, 0, 10}}, 1, 2, 1, 1},
		{4, [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}}, 0, 3, 1, 6},
		{11, [][]int{{0, 3, 3}, {3, 4, 3}, {4, 1, 3}, {0, 5, 1}, {5, 1, 100}, {0, 6, 2}, {6, 1, 100}, {0, 7, 1}, {7, 8, 1}, {8, 9, 1}, {9, 1, 1}, {1, 10, 1}, {10, 2, 1}, {1, 2, 100}}, 0, 2, 4, 11},
		{5, [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}, {3, 4, 1}}, 0, 4, 2, 7},
		{8, [][]int{{0, 1, 100}, {1, 2, 200}, {2, 3, 1}, {0, 4, 5}, {4, 5, 4}, {5, 2, 3}, {0, 6, 1}, {6, 7, 1}, {7, 2, 1}}, 0, 3, 2, 301},
	}

	testOnly := 0
	failed := false
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		actual := findCheapestPrice(test.n, test.flights, test.src, test.dst, test.k)
		if actual != test.expected {
			failed = true
			fmt.Printf("Test %v Failed\n\tActual  : %v\n\tExpected: %v\n", i+1, actual, test.expected)
		}
	}
	if !failed {
		fmt.Println("Passed All Tests.")
	}
}

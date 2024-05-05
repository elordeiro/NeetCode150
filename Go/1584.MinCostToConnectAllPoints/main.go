package main

import (
	"fmt"
	"math"
)

func minCostConnectPoints(points [][]int) int {
	N := len(points)
	cost := 0

	inf := math.MaxInt
	dist := make([]int, N)
	for i := range N {
		dist[i] = inf
	}

	var abs = func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for p1 := range N - 1 {
		next := p1 + 1
		for p2 := next; p2 < N; p2++ {
			dist[p2] = min(dist[p2], abs(points[p1][0]-points[p2][0])+abs(points[p1][1]-points[p2][1]))
			if dist[next] > dist[p2] {
				dist[next], dist[p2], points[next], points[p2] = dist[p2], dist[next], points[p2], points[next]
			}
		}
		cost += dist[next]
	}
	return cost
}

func main() {
	tests := []struct {
		points   [][]int
		expected int
	}{
		{[][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}, 20},
		{[][]int{{3, 12}, {-2, 5}, {-4, 1}}, 18},
	}

	testOnly := 2
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		actual := minCostConnectPoints(test.points)
		if actual != test.expected {
			fmt.Printf("Test %v Failed\n\tActual  : %v\n\tExpected: %v\n", i+1, actual, test.expected)
		}
	}
}

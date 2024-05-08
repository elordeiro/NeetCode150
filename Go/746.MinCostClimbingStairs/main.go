package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	twoDown, oneDown := cost[0], cost[1]
	n := len(cost)
	for i := 2; i < n; i++ {
		temp := min(oneDown, twoDown) + cost[i]
		twoDown = oneDown
		oneDown = temp
	}
	return min(oneDown, twoDown)
}

func main() {
	tests := []struct {
		cost     []int
		expected int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}
	testOnly := 0
	failed := false
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		actual := minCostClimbingStairs(test.cost)
		if actual != test.expected {
			failed = true
			fmt.Printf("Failed %v Test\n\tActual  : %v\n\tExpected: %v\n", i+1, actual, test.expected)
		}
	}
	if !failed {
		fmt.Println("All Tests Passed")
	}
}

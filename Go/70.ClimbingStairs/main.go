package main

import "fmt"

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	tests := []struct {
		n, expected int
	}{
		{2, 2},
		{3, 3},
		{44, 1134903170},
	}
	testOnly := 0
	failed := false
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		actual := climbStairs(test.n)
		if actual != test.expected {
			failed = true
			fmt.Printf("Test %v Failed\n\tActual  : %v\n\tExpected: %v\n", i+1, actual, test.expected)
		}
	}
	if !failed {
		fmt.Println("Passed All Tests")
	}
}

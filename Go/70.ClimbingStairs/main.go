package main

import "fmt"

func climbStairs(n int) int {
	twoDown, oneDown := 1, 1

	for i := 2; i < n+1; i++ {
		temp := oneDown + twoDown
		twoDown = oneDown
		oneDown = temp
	}

	return oneDown
}

func main() {
	tests := []struct {
		n, expected int
	}{
		{2, 2},
		{3, 3},
		{5, 8},
		{44, 1_134_903_170},
	}
	testOnly := 3
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

package main

import "fmt"

func maxProfit(prices []int) int {
	n := len(prices)
	i, j, profit := 0, 0, 0

	for j < n {
		if prices[j] < prices[i] {
			i = j
		} else {
			profit = max(profit, prices[j]-prices[i])
		}
		j++
	}
	return profit
}

func print_failed_test(i int, res int, expected int) {
	fmt.Printf("Test %d failed:\n", i+1)
	fmt.Printf("    \tActual  : %v\n", res)
	fmt.Printf("    \tExpected: %v\n", expected)
}

func main() {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 2, 5, 1, 7, 4}, 6},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{2, 1, 2, 1, 0, 1, 2}, 2},
		{[]int{2, 4, 1, 11, 7}, 10},
	}

	passed_all := true
	for i, test := range tests {
		nums, expected := test.nums, test.expected
		res := maxProfit(nums)
		if res != expected {
			passed_all = false
			print_failed_test(i, res, expected)
		}
	}
	if passed_all {
		fmt.Println("Passed all tests.")
	}
}

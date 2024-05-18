package main

import "fmt"

func coinChange(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for i := range amount + 1 {
		memo[i] = 10001
	}
	memo[0] = 0
	for i := 1; i < amount+1; i++ {
		min_coins := 10001
		for _, coin := range coins {
			if i == coin {
				min_coins = 1
				break
			}
			if i-coin > -1 {
				min_coins = min(min_coins, memo[i-coin]+1)
			}
		}
		memo[i] = min_coins
	}
	if memo[amount] != 10001 {
		return memo[amount]
	}
	return -1
}

func main() {
	tests := []struct {
		coins            []int
		amount, expected int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{[]int{2, 5, 10, 1}, 27, 4},
		{[]int{186, 419, 83, 408}, 6249, 20},
	}
	failed := false
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		actual := coinChange(test.coins, test.amount)
		if actual != test.expected {
			failed = true
			fmt.Printf("Test %v Failed\n\tActual  : %v\n\tExpected: %v\n", i+1, actual, test.expected)
		}
	}
	if !failed {
		fmt.Println("Passed All Tests.")
	}
}

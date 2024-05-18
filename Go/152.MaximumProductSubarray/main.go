package main

import "fmt"

func maxProduct(nums []int) int {
	res := nums[0]
	curMin, curMax := 1, 1
	var temp int
	for _, num := range nums {
		temp = curMax
		curMax = max(num, num*curMax, num*curMin)
		curMin = min(num, num*curMin, num*temp)
		res = max(res, curMin, curMax)
	}
	return res
}

func main() {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-2, 3, -4}, 24},
		{[]int{0, 2}, 2},
		{[]int{2, -5, -2, -4, 3}, 24},
		{[]int{-2}, -2},
	}
	failed := false
	testOnly := 0
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		actual := maxProduct(test.nums)
		if actual != test.expected {
			failed = true
			fmt.Printf("Test %v Failed\n\tActual  : %v\n\tExpected: %v\n", i+1, actual, test.expected)
		}
	}
	if !failed {
		fmt.Println("Passed All Tests.")
	}
}

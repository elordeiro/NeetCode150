package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	skippedPrev, robbedPrev := nums[0], nums[1]
	n := len(nums)
	for i := 2; i < n; i++ {
		temp := skippedPrev
		skippedPrev = max(robbedPrev, skippedPrev)
		robbedPrev = temp + nums[i]
	}
	return max(skippedPrev, robbedPrev)
}

func main() {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{2, 1, 1, 2}, 4},
	}

	testOnly := 0
	failed := false
	for i, test := range tests {
		if testOnly != 0 && testOnly != i+1 {
			continue
		}
		actual := rob(test.nums)
		if actual != test.expected {
			failed = true
			fmt.Printf("Failed Test %v\n\tActual  : %v\n\tExpected: %v\n", i+1, actual, test.expected)
		}
	}
	if !failed {
		fmt.Println("All Tests Passed.")
	}
}

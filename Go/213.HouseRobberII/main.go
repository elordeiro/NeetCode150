package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	skipPrev1, robPrev1 := nums[0], nums[1]
	skipPrev2, robPrev2 := 0, nums[1]
	n := len(nums) - 1
	for i := 2; i < n; i++ {
		temp1, temp2 := skipPrev1, skipPrev2
		skipPrev1, skipPrev2 = max(skipPrev1, robPrev1), max(skipPrev2, robPrev2)
		robPrev1, robPrev2 = temp1+nums[i], temp2+nums[i]
	}

	temp := skipPrev2
	skipPrev2 = max(skipPrev2, robPrev2)
	robPrev2 = temp + nums[n]

	return max(skipPrev1, skipPrev2, robPrev1, robPrev2)

}

func main() {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, 2}, 3},
		{[]int{1, 2, 3, 1}, 4},
		{[]int{1, 2, 3}, 3},
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
		fmt.Println("Passed All Tests.")
	}
}
